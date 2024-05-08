package errors

import (
	"encoding/xml"
	"fmt"
)

type ApiError struct {
	XMLName  xml.Name `json:"-" xml:"Error"`
	Type     string   `json:"type" xml:",attr"`
	Title    string   `json:"title"`
	Instance string   `json:"instance"`
	Detail   string   `json:"detail,omitempty"`
	Code     int      `json:"status" xml:"-"`
}

func (err ApiError) Error() string {
	return fmt.Sprintf("%v(%v) %v: %v", err.Code, err.Instance, err.Title, err.Detail)
}
