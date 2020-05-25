package travelr

import (
	"fmt"
	"time"
)

var (
	invalidParams = fmt.Errorf("Invalid parameters")
	dataNotFound  = fmt.Errorf("Not found")
	invalidData   = fmt.Errorf("Invalid data")
)

const (
	host      = "localhost"
	locale    = "en"
	tripClass = "Y"
	userIP    = "127.0.0.1"

	timeFormat = "2006-01-02"
	url        = "http://api.travelpayouts.com"
)

// endpoint ...
func endpoint(suffix string) string {
	return fmt.Sprintf("%v/%v", url, suffix)
}

// convertInt ...
func convertInt(i int) *int {
	return &i
}

// bodySetting ...
type bodySetting struct {
	Token       string      `json:"token"`
	Host        string      `json:"host"`
	Marker      string      `json:"marker"`
	UserIP      string      `json:"user_ip"`
	Locale      string      `json:"locale"`
	TripClass   string      `json:"trip_class"`
	KnowEnglish string      `json:"know_english"`
	Passenger   *Passengers `json:"passengers"`
}

// Passengers ...
type Passengers struct {
	Adults   *int `json:"adults"`
	Children *int `json:"children"`
	Infants  *int `json:"infants"`
}

// Segment ...
type Segment struct {
	Origin      *string `json:"origin"`
	Destination *string `json:"destination"`
	Date        *string `json:"date"`
}

// RoundTripBody ...
type RoundTripBody struct {
	Signature   *string     `json:"signature"`
	Marker      *string     `json:"marker"`
	Host        *string     `json:"host"`
	UserIP      *string     `json:"user_ip"`
	Locale      *string     `json:"locale"`
	TripClass   *string     `json:"trip_class"`
	Passengers  *Passengers `json:"passengers"`
	Segments    *[]Segment  `json:"segments"`
	KnowEnglish *string     `json:"know_english"`
}

// SegmentRequest ...
type SegmentRequest struct {
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Date        time.Time `json:"date"`
}
