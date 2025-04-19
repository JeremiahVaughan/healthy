package database

import (
    "fmt"
    "os"
	"database/sql"
    "time"
    "errors"
    "log"

    "github.com/JeremiahVaughan/healthy/config"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const (
    InternalUnexpectedErrorKey = "internal_unexpected_error"
    ExternalUnexpectedErrorKey = "external_unexpected_error"
)

type Client struct {
    conn *sql.DB
    migrationDir string
}

type HealthStatus struct {
    Service string `json:"service"`
    StatusKey string `json:"statusKey"`
    Unhealthy bool `json:"unhealthy"`
    UnhealthyStartedAt int64 `json:"unhealthyStartedAt"`
    // UnhealthyDelayInSeconds this many seconds will pass with an unhealthy status of true before status cake is triggered
    UnhealthyDelayInSeconds int64 `json:"unhealthyDelayInSeconds"` 
    Message string `json:"message"`
    ExpiresAt int64
}

func (s *HealthStatus) IsValid() bool {
    return s.Service != "" && s.StatusKey != ""
}

func New(config config.Database) (*Client, error) {
    var err error
    _, err = os.Stat(config.DataDirectory)
    if os.IsNotExist(err) {
        err = os.MkdirAll(config.DataDirectory, 0700)
        if err != nil {
            return nil, fmt.Errorf("error, when creating database data directory. Error: %v", err)
        }
    }
    c := Client{
        migrationDir: config.MigrationDirectory,
    }
    dbFile := fmt.Sprintf("%s/data", config.DataDirectory)
    c.conn, err = sql.Open("sqlite3", dbFile)
    if err != nil {
        return nil, fmt.Errorf("error, when entablishing database connection. Error: %v", err)
    }
    err = c.migrate()
    if err != nil {
        return nil, fmt.Errorf("error, when migrating database files. Error: %v", err)
    }
    return &c, nil
}

func (c *Client) InsertHealthStatus(status HealthStatus) error {
    // ignore is in case the status is being inserted for the very first time by two different goroutines
    // it is unlikely but even if it did occure the second insert would just be redundant so it is safe to ignore
    _, err := c.conn.Exec(
        `INSERT OR IGNORE INTO health_status (
            service,
            status_key,
            unhealthy_started_at,
            unhealthy_delay_in_seconds,
            message,
            expires_at
        ) VALUES (
            ?,
            ?,
            ?,
            ?,
            ?,
            ?
        )`,
    )
    if err != nil {
        return fmt.Errorf("error, when creating health status. Error: %v", err)
    }
    return nil
}

func (c *Client) UpdateHealthStatus(status HealthStatus) error {
    _, err := c.conn.Exec(
        `UPDATE health_status
        SET unhealthy_started_at = ?,
            unhealthy_delay_in_seconds = ?,
            message = ?,
            expires_at = ?
        WHERE service = ? 
            AND status_key = ?`,
        status.UnhealthyStartedAt,
        status.UnhealthyDelayInSeconds,
        status.Message,
        status.ExpiresAt,
    )
    if err != nil {
        return fmt.Errorf("error, when creating health status. Error: %v", err)
    }
    return nil
}

func (c *Client) DeleteHealthStatus(status HealthStatus) error {
    _, err := c.conn.Exec(
        `DELETE FROM health_status
        WHERE service = ? 
            AND status_key = ?`,
    )
    if err != nil {
        return fmt.Errorf("error, when creating health status. Error: %v", err)
    }
    return nil
}

func (c *Client) ClearUnexpected() error {
    _, err := c.conn.Exec(
        `UPDATE health_status
        SET unhealthy_started_at = 0
        WHERE status_key IN (?, ?)`,
        InternalUnexpectedErrorKey,
        ExternalUnexpectedErrorKey,
    )
    if err != nil {
        return fmt.Errorf("error, when executing sql statement. Error: %v", err)
    }
    return nil
}

func (c *Client) FetchExistingHealthStatus(status HealthStatus) (*HealthStatus, error) {
    var result HealthStatus
    err := c.conn.QueryRow(
        `SELECT service,
                status_key,
                unhealthy_started_at,
                unhealthy_delay_in_seconds,
                message,
                expires_at
        FROM health_status
        WHERE service = ?
            AND status_key = ?`,
            status.Service,
            status.StatusKey,
    ).Scan(
        &result.Service,
        &result.StatusKey,
        &result.UnhealthyStartedAt,
        &result.UnhealthyDelayInSeconds,
        &result.Message,
        &result.ExpiresAt,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        } else {
            return nil, fmt.Errorf("error, when attempting to execute sql statement. Error: %v", err)
        }
    }
    return &result, nil
}

func (c *Client) FetchAllHealthStatuses() ([]HealthStatus, error) {
    rows, err := c.conn.Query(
        `SELECT service,
                status_key,
                unhealthy_started_at,
                unhealthy_delay_in_seconds,
                message,
                expires_at
        FROM health_status`,
    )

    defer func(rows *sql.Rows) {
        if rows != nil {
            closeRowsError := rows.Close()
            if closeRowsError != nil {
                // no choice but to log the error since defer doesn't let us return errors
                // defer is needed though because it ensures a cleanup attempt is made even if we should return early due to an error
                log.Printf("error, when attempting to close database rows: %v", closeRowsError)
            }
        }
    }(rows)

    if err != nil {
        return nil, fmt.Errorf("error, when attempting to retrieve records. Error: %v", err)
    }

    currentTime := time.Now().Unix()
    var result []HealthStatus
    for rows.Next() {
        var r HealthStatus
        err = rows.Scan(
            &r.Service, 
            &r.StatusKey, 
            &r.UnhealthyStartedAt, 
            &r.UnhealthyDelayInSeconds, 
            &r.Message, 
            &r.ExpiresAt,
        )
        if err != nil {
            return nil, fmt.Errorf("error, when scanning database rows. Error: %v", err)
        }
        r.calculateUnhealthy(currentTime)
        result = append(result, r)
    }

    err = rows.Err()
    if err != nil {
        return nil, fmt.Errorf("error, when iterating through database rows. Error: %v", err)
    }
    return result, nil
}

func (s *HealthStatus) calculateUnhealthy(currentTime int64) {
    if s.ExpiresAt <= currentTime && (s.StatusKey != ExternalUnexpectedErrorKey && s.StatusKey != InternalUnexpectedErrorKey) {
        s.Message = "alert status is expired"
        s.UnhealthyStartedAt = s.ExpiresAt
        s.Unhealthy = true
    } else if s.UnhealthyStartedAt != 0 {
        if s.UnhealthyStartedAt + s.UnhealthyDelayInSeconds < currentTime {
            s.Unhealthy = true
        } else {
            s.Message = "" // nothing to report until unhealthy is determined
        }
    }
}
