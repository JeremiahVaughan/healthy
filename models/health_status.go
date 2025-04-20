package models

import (
    "fmt"
    "sync"
    "time"
    "log"

    "github.com/JeremiahVaughan/healthy/clients"
    "github.com/JeremiahVaughan/healthy/clients/database"
    "github.com/JeremiahVaughan/healthy/clients/nats"
    "github.com/JeremiahVaughan/healthy/config"
)

type HealthStatusModel struct {
    m sync.Mutex
    db *database.Client
    nats *nats.Client
    healthStatusExpiresDurationInSeconds int64
}

func NewHealthStatusModel(clients *clients.Clients, config config.Config) *HealthStatusModel {
    return &HealthStatusModel{
        db: clients.Database,
        nats: clients.Nats,
        healthStatusExpiresDurationInSeconds: config.HealthStatusExpiresDurationInSeconds,
    }
}

func (m *HealthStatusModel) UpdateHealthStatus(newStatus database.HealthStatus) error {
    valid := newStatus.IsValid()
    if !valid {
        return fmt.Errorf(
            "error, invalid status provided by service: %s of key: %s",
            newStatus.Service,
            newStatus.StatusKey,
        )
    }
    m.m.Lock()
    defer m.m.Unlock()
    existing, err := m.db.FetchExistingHealthStatus(newStatus)    
    if err != nil {
        return fmt.Errorf("error, when fetching existing health status for HealthStatusModel.UpdateHealthStatus(). Error: %v", err)
    }
    merged := m.mergeHealthStatuses(
        existing,
        newStatus,
        time.Now().Unix(),
    )    
    if existing == nil {
        err = m.db.InsertHealthStatus(merged)
        if err != nil {
            return fmt.Errorf("error, when inserting new health status for HealthStatusModel.UpdateHealthStatus(). Error: %v", err)
        }
    } else {
        err = m.db.UpdateHealthStatus(merged)
        if err != nil {
            return fmt.Errorf("error, when updating existing health status for HealthStatusModel.UpdateHealthStatus(). Error: %v", err)
        }
    }
    return nil
}

func (m *HealthStatusModel) DeleteHealthStatus(status database.HealthStatus) error {
    m.m.Lock()
    defer m.m.Unlock()
    err := m.db.DeleteHealthStatus(status) 
    if err != nil {
        return fmt.Errorf("error, when deleting health status. Error: %v", err)
    }
    return nil
}

func (m *HealthStatusModel) mergeHealthStatuses(
    existingStatus *database.HealthStatus,
    newStatus database.HealthStatus,
    currentTime int64,
) database.HealthStatus {
    mergedStatus := database.HealthStatus{
        Service: newStatus.Service, 
        StatusKey: newStatus.StatusKey, 
        Message: newStatus.Message,
        UnhealthyDelayInSeconds: newStatus.UnhealthyDelayInSeconds,
        ExpiresAt: currentTime + m.healthStatusExpiresDurationInSeconds,
    }
    if newStatus.Unhealthy {
        if existingStatus == nil || existingStatus.UnhealthyStartedAt == 0 {
            mergedStatus.UnhealthyStartedAt = currentTime
        } else {
            mergedStatus.UnhealthyStartedAt = existingStatus.UnhealthyStartedAt
        }
    } else {
        mergedStatus.UnhealthyStartedAt = 0
    }
    return mergedStatus
}

func (m *HealthStatusModel) IsHealthy() (bool, error) {
    result, err := m.db.FetchAllHealthStatuses()
    if err != nil {
        return false, fmt.Errorf("error, when fetching all health statuses for HealthStatusModel.FetchAllHealthStatuses(). Error: %v", err)
    }
    for _, r := range result {
        if r.Unhealthy {
            return false, nil
        }
    }
    return true, nil
}

func (m *HealthStatusModel) GetAllStatuses() ([]database.HealthStatus, error) {
    result, err := m.db.FetchAllHealthStatuses()
    if err != nil {
        return nil, fmt.Errorf("error, when fetching all health statuses for HealthStatusModel.FetchAllHealthStatuses(). Error: %v", err)
    }
    return result, nil
}


func (m *HealthStatusModel) ExternalUnexpectedError(status database.HealthStatus) {
    s := database.HealthStatus{
        Service: status.Service,
        Message: status.Message,
        StatusKey: database.ExternalUnexpectedErrorKey,
        UnhealthyStartedAt: time.Now().Unix(),
        UnhealthyDelayInSeconds: 0,
        Unhealthy: true,
    }
    err := m.UpdateHealthStatus(s)
    if err != nil {
        err = fmt.Errorf("error, when updating health status for external error. Error: %v", err)
        m.InternalUnexpectedError(err)
        return
    }
    return 
}

func (m *HealthStatusModel) InternalUnexpectedError(err error) {
    status := database.HealthStatus{
        Service: "healthy",
        StatusKey: database.InternalUnexpectedErrorKey,
        UnhealthyStartedAt: time.Now().Unix(),
        UnhealthyDelayInSeconds: 0,
        Message: err.Error(),
    }
    err = m.UpdateHealthStatus(status)
    if err != nil {
        log.Fatalf("error, when updating health status for internal error. Error: %v", err)
    }
    return 
}

func (m *HealthStatusModel) ClearUnexpectedErrorStatuses() error {
    err := m.db.ClearUnexpected()
    if err != nil {
        return fmt.Errorf("error, when clearing unexpected errors. Error: %v", err)
    }
    return nil
}

func (m *HealthStatusModel) RefreshStatus(subject string, bytes []byte) error {
    err := m.nats.Conn.Publish(subject, bytes)
    if err != nil {
        return fmt.Errorf("error, when publishing refresh status. Error: %v", err)
    }
    return nil
}
