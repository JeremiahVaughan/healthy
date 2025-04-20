package models

import (
    "testing"

    "github.com/JeremiahVaughan/healthy/clients/database"
)

func Test_mergeHealthStatuses(t *testing.T) {
    currentTime := int64(5)
    
    m := HealthStatusModel{
        healthStatusExpiresDurationInSeconds: 3,
    }
    t.Run("old status healthy, new status unhealthy", func(t *testing.T){
        oldStatus := database.HealthStatus{
            Message: "nothing to report",
            UnhealthyDelayInSeconds: 4,
            Unhealthy: false,
            UnhealthyStartedAt: 0,
        }
        newStatus := database.HealthStatus{
            Message: "oh noes",
            UnhealthyDelayInSeconds: 5,
            Unhealthy: true,
        }
        mergedStatus := m.mergeHealthStatuses(&oldStatus, newStatus, currentTime)
        if mergedStatus.Message != "oh noes" {
            t.Errorf("error, expected message not found. Message: %s", mergedStatus.Message)
        }
        if mergedStatus.UnhealthyDelayInSeconds != 5 {
            t.Errorf("error, delay was not updated. Found delay: %d", mergedStatus.UnhealthyDelayInSeconds)
        }
        if mergedStatus.ExpiresAt != 8 {
            t.Errorf("error, expected expires at to be 8, but it was: %d", mergedStatus.ExpiresAt)
        }
        if mergedStatus.UnhealthyStartedAt != currentTime {
            t.Errorf("error, expected unhealthy started at of %d but got %d", currentTime, mergedStatus.UnhealthyStartedAt)
        }
    })

    t.Run("old status healthy, new status healthy", func(t *testing.T){
        oldStatus := database.HealthStatus{
            Message: "nothing to report",
            UnhealthyDelayInSeconds: 4,
            Unhealthy: false,
            UnhealthyStartedAt: 0,
        }
        newStatus := database.HealthStatus{
            Message: "nothing to report ok?",
            UnhealthyDelayInSeconds: 5,
            Unhealthy: false,
        }
        mergedStatus := m.mergeHealthStatuses(&oldStatus, newStatus, currentTime)
        if mergedStatus.Message != "nothing to report ok?" {
            t.Errorf("error, expected message not found. Message: %s", mergedStatus.Message)
        }
        if mergedStatus.UnhealthyDelayInSeconds != 5 {
            t.Errorf("error, delay was not updated. Found delay: %d", mergedStatus.UnhealthyDelayInSeconds)
        }
        if mergedStatus.ExpiresAt != 8 {
            t.Errorf("error, expected expires at to be 8, but it was: %d", mergedStatus.ExpiresAt)
        }
        if mergedStatus.UnhealthyStartedAt != 0 {
            t.Errorf("error, expected unhealthy started at of 0 but got %d", mergedStatus.UnhealthyStartedAt)
        }
    })

    t.Run("old status unhealthy, new status unhealthy", func(t *testing.T){
        oldStatus := database.HealthStatus{
            Message: "not good",
            UnhealthyDelayInSeconds: 4,
            Unhealthy: true,
            UnhealthyStartedAt: 3,
        }
        newStatus := database.HealthStatus{
            Message: "really not good",
            UnhealthyDelayInSeconds: 5,
            Unhealthy: true,
        }
        mergedStatus := m.mergeHealthStatuses(&oldStatus, newStatus, currentTime)
        if mergedStatus.Message != "really not good" {
            t.Errorf("error, expected message not found. Message: %s", mergedStatus.Message)
        }
        if mergedStatus.UnhealthyDelayInSeconds != 5 {
            t.Errorf("error, delay was not updated. Found delay: %d", mergedStatus.UnhealthyDelayInSeconds)
        }
        if mergedStatus.ExpiresAt != 8 {
            t.Errorf("error, expected expires at to be 8, but it was: %d", mergedStatus.ExpiresAt)
        }
        if mergedStatus.UnhealthyStartedAt != 3 {
            t.Errorf("error, expected unhealthy started at of 3 but got %d", mergedStatus.UnhealthyStartedAt)
        }
    })

    t.Run("old status unhealthy, new status healthy", func(t *testing.T){
        oldStatus := database.HealthStatus{
            Message: "not good",
            UnhealthyDelayInSeconds: 4,
            Unhealthy: true,
            UnhealthyStartedAt: 3,
        }
        newStatus := database.HealthStatus{
            Message: "good again",
            UnhealthyDelayInSeconds: 5,
            Unhealthy: false,
        }
        mergedStatus := m.mergeHealthStatuses(&oldStatus, newStatus, currentTime)
        if mergedStatus.Message != "good again" {
            t.Errorf("error, expected message not found. Message: %s", mergedStatus.Message)
        }
        if mergedStatus.UnhealthyDelayInSeconds != 5 {
            t.Errorf("error, delay was not updated. Found delay: %d", mergedStatus.UnhealthyDelayInSeconds)
        }
        if mergedStatus.ExpiresAt != 8 {
            t.Errorf("error, expected expires at to be 8, but it was: %d", mergedStatus.ExpiresAt)
        }
        if mergedStatus.UnhealthyStartedAt != 0 {
            t.Errorf("error, expected unhealthy started at of 0 but got %d", mergedStatus.UnhealthyStartedAt)
        }
    })

    t.Run("old status doesn't exist, new status not healthy", func(t *testing.T){
        newStatus := database.HealthStatus{
            Message: "not good",
            UnhealthyDelayInSeconds: 5,
            Unhealthy: true,
        }
        mergedStatus := m.mergeHealthStatuses(nil, newStatus, currentTime)
        if mergedStatus.Message != "not good" {
            t.Errorf("error, expected message not found. Message: %s", mergedStatus.Message)
        }
        if mergedStatus.UnhealthyDelayInSeconds != 5 {
            t.Errorf("error, delay was not updated. Found delay: %d", mergedStatus.UnhealthyDelayInSeconds)
        }
        if mergedStatus.ExpiresAt != 8 {
            t.Errorf("error, expected expires at to be 8, but it was: %d", mergedStatus.ExpiresAt)
        }
        if mergedStatus.UnhealthyStartedAt != currentTime {
            t.Errorf("error, expected unhealthy started at of %d but got %d", currentTime, mergedStatus.UnhealthyStartedAt)
        }
    })

    t.Run("old status doesn't exist, new status is healthy", func(t *testing.T){
        newStatus := database.HealthStatus{
            Message: "good",
            UnhealthyDelayInSeconds: 5,
            Unhealthy: false,
        }
        mergedStatus := m.mergeHealthStatuses(nil, newStatus, currentTime)
        if mergedStatus.Message != "good" {
            t.Errorf("error, expected message not found. Message: %s", mergedStatus.Message)
        }
        if mergedStatus.UnhealthyDelayInSeconds != 5 {
            t.Errorf("error, delay was not updated. Found delay: %d", mergedStatus.UnhealthyDelayInSeconds)
        }
        if mergedStatus.ExpiresAt != 8 {
            t.Errorf("error, expected expires at to be 8, but it was: %d", mergedStatus.ExpiresAt)
        }
        if mergedStatus.UnhealthyStartedAt != 0 {
            t.Errorf("error, expected unhealthy started at of 0 but got %d", mergedStatus.UnhealthyStartedAt)
        }
    })

}
