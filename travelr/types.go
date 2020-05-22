package travelr

import "fmt"

var (
	invalidParams = fmt.Errorf("Invalid parameters")
	dataNotFound  = fmt.Errorf("Not found")
)

const (
	host      = "localhost"
	locale    = "en"
	tripClass = "Y"
	token     = "0088b3ed5e8e56d471d7f49ac4e2ee6f"
)

type DefaultRequest struct {
	Data map[string]interface{} `json:"data"`
}
