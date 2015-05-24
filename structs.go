package main


type PushjetApiCall struct {
    Message PushjetMessage
    Subscription  PushjetSubscription
}

type PushjetService struct {
    Created int
    Icon    string
    Name    string
    Public  string
}

type PushjetSubscription struct {
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
