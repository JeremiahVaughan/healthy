package database

import (
    "fmt"
    "os"
	"database/sql"

    "github.com/JeremiahVaughan/healthy/config"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type Client struct {
    conn *sql.DB
}

func New(config config.Database) (*Client, error) {
    var err error
    _, err = os.Stat(config.DirectoryLocation)
    if os.IsNotExist(err) {
        err = os.MkdirAll(config.DirectoryLocation, 0700)
        if err != nil {
            return nil, fmt.Errorf("error, when creating database data directory. Error: %v", err)
        }
    }
    var c Client
    dbFile := fmt.Sprintf("%s/data", config.DirectoryLocation)
    c.conn, err = sql.Open("sqlite3", dbFile)
    if err != nil {
        return nil, fmt.Errorf("error, when entablishing database connection. Error: %v", err)
    }
    return &c, nil
}
