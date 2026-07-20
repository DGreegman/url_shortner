package models

import "time"

type ClickEvent struct {
	EventID    int     `json:"event_id"`
	LinkID     int     `json:"link_id"`
	Code 	   string  `json:"code"`
	IP		   string  `json:"ip_address"`
	UserAgent  string  `json:"user_agent"`
	Referer    string  `json:"referrer"`
	DeviceType string  `json:"device_type"`
	Country    string  `json:"country"`
	Timestamp  time.Time `json:"timestamp"`
}

type AnalyticsResponse struct {
	Code        string       `json:"code"`
	TargetUrl   string       `json:"target_url"`
	TotalClicks string       `json:"redirect_type"`
	Events	    []ClickEvent `json:"events"`
}