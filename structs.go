package main

const (
    ACTION_NEW_MESSAGE   = iota
    ACTION_UPDATE_LISTEN = iota
)

type PushjetApiCall struct {
    Message PushjetMessage
    Listen  PushjetListen
}

type PushjetService struct {
    Created int
    Icon    string
    Name    string
    Public  string
}

type PushjetListen struct {
    Uuid              string
    Timestamp         int
    Timestamp_checked int
    Service           PushjetService
}

type PushjetMessage struct {
    Level     int
    Link      string
    Message   string
    Service   PushjetService
    Timestamp int
    Title     string
}
