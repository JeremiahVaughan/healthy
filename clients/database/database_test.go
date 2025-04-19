package database

import (
    "testing"
)

func Test_calculateUnhealthy(t *testing.T) {
    currentTime := int64(5)
    t.Run("is not expired or triggered", func(t *testing.T) {
        status := HealthStatus{
            UnhealthyStartedAt: 0,
            UnhealthyDelayInSeconds: 0,
            Message: "",
            ExpiresAt: 605,
        }
        status.calculateUnhealthy(currentTime)
        if status.Unhealthy {
            t.Errorf("error, should not be unhealthy but is")
        }
    })

    t.Run("is expired but not triggered", func(t *testing.T) {
        status := HealthStatus{
            UnhealthyStartedAt: 0,
            UnhealthyDelayInSeconds: 0,
            Message: "",
            ExpiresAt: 4,
        }
        status.calculateUnhealthy(currentTime)
        if !status.Unhealthy {
            t.Errorf("error, should be unhealthy but is not")
        }
        if status.UnhealthyStartedAt != 4 {
            t.Errorf("error, should be unhealthy started at 4 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "alert status is expired" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })

    t.Run("is expired and triggered", func(t *testing.T) {
        status := HealthStatus{
            UnhealthyStartedAt: 2,
            UnhealthyDelayInSeconds: 0,
            Message: "bad bad bad",
            ExpiresAt: 4,
        }
        status.calculateUnhealthy(currentTime)
        if !status.Unhealthy {
            t.Errorf("error, should be unhealthy but is not")
        }
        if status.UnhealthyStartedAt != 4 {
            t.Errorf("error, should be unhealthy started at 4 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "alert status is expired" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })
    
    t.Run("is not expired but is triggered, but not enough delay has passed", func(t *testing.T) {
        status := HealthStatus{
            UnhealthyStartedAt: 2,
            UnhealthyDelayInSeconds: 5,
            Message: "bad bad bad",
            ExpiresAt: 605,
        }
        status.calculateUnhealthy(currentTime)
        if status.Unhealthy {
            t.Errorf("error, should not be unhealthy until enough delay has passed")
        }
        if status.UnhealthyStartedAt != 2 {
            t.Errorf("error, should be unhealthy started at 2 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })

    t.Run("is not expired but is triggered, enough delay has passed", func(t *testing.T) {
        status := HealthStatus{
            UnhealthyStartedAt: 2,
            UnhealthyDelayInSeconds: 2,
            Message: "bad bad bad",
            ExpiresAt: 605,
        }
        status.calculateUnhealthy(currentTime)
        if !status.Unhealthy {
            t.Errorf("error, should be unhealthy")
        }
        if status.UnhealthyStartedAt != 2 {
            t.Errorf("error, should be unhealthy started at 2 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "bad bad bad" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })

    t.Run("is not expired but is triggered, there is no delay", func(t *testing.T) {
        status := HealthStatus{
            UnhealthyStartedAt: 4,
            UnhealthyDelayInSeconds: 0,
            Message: "bad bad bad",
            ExpiresAt: 605,
        }
        status.calculateUnhealthy(currentTime)
        if !status.Unhealthy {
            t.Errorf("error, should be unhealthy")
        }
        if status.UnhealthyStartedAt != 4 {
            t.Errorf("error, should be unhealthy started at 4 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "bad bad bad" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })

    t.Run("expiration does not apply to unexpected internal errors", func(t *testing.T) {
        status := HealthStatus{
            StatusKey: InternalUnexpectedErrorKey,
            UnhealthyStartedAt: 1,
            UnhealthyDelayInSeconds: 0,
            Message: "bad bad bad",
            ExpiresAt: 2,
        }
        status.calculateUnhealthy(currentTime)
        if !status.Unhealthy {
            t.Errorf("error, should always be unhealthy")
        }
        if status.UnhealthyStartedAt != 1 {
            t.Errorf("error, should be unhealthy started at 1 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "bad bad bad" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })

    t.Run("expiration does not apply to unexpected external errors", func(t *testing.T) {
        status := HealthStatus{
            StatusKey: ExternalUnexpectedErrorKey,
            UnhealthyStartedAt: 1,
            UnhealthyDelayInSeconds: 0,
            Message: "bad bad bad",
            ExpiresAt: 2,
        }
        status.calculateUnhealthy(currentTime)
        if !status.Unhealthy {
            t.Errorf("error, should always be unhealthy")
        }
        if status.UnhealthyStartedAt != 1 {
            t.Errorf("error, should be unhealthy started at 1 but is %d", status.UnhealthyStartedAt)
        }
        if status.Message != "bad bad bad" {
            t.Errorf("error, incorrect message: %s", status.Message)
        }
    })

}

func Test_IsValid(t *testing.T) {
    t.Run("valid", func(t *testing.T){
        s := HealthStatus{
            Service: "test",
            StatusKey: "test", 
        }
        valid := s.IsValid()
        if !valid {
            t.Errorf("error, expected valid but did not get")
        }
    })
    t.Run("invalid 1", func(t *testing.T){
        s := HealthStatus{
            Service: "test",
            StatusKey: "", 
        }
        valid := s.IsValid()
        if valid {
            t.Errorf("error, expected invalid but did not get")
        }
    })
    t.Run("invalid 2", func(t *testing.T){
        s := HealthStatus{
            Service: "",
            StatusKey: "test", 
        }
        valid := s.IsValid()
        if valid {
            t.Errorf("error, expected invalid but did not get")
        }
    })
}
