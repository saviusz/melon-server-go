package errors

import (
	"encoding/xml"
	"fmt"
)

type ApiError struct {
	XMLName  xml.Name `json:"-" xml:"Error"`
	Type     string   `json:"type" xml:",attr"` // Uri to documentation
	Title    string   `json:"title"`
	Instance string   `json:"instance"` // Identifier (or uri with details)
	Detail   string   `json:"detail,omitempty"`
	Code     int      `json:"status" xml:"-"`
}

func (err ApiError) Error() string {
	return fmt.Sprintf("%v(%v) %v: %v", err.Code, err.Instance, err.Title, err.Detail)
}

func UnsupportedMediaError(mediatype string) *ApiError {

	detail := fmt.Sprintf("Provided mediatype %q is not supported", mediatype)

	return &ApiError{
		Code:     415,
		Title:    "Unsupported Media Type",
		Detail:   detail,
		Type:     "about:blank",
		Instance: "elo", // TODO: Change me to random
	}
}

func InternalServerError(err error) *ApiError {

	// TODO: Insert logging

	return &ApiError{
		Code:     500,
		Title:    "Internal Server Error",
		Detail:   err.Error(),
		Type:     "about:blank",
		Instance: "elo", //TODO: Change me to random
	}
}

func BadRequestError(desc string) *ApiError {
	detail := &desc
	if desc == "" {
		detail = nil
	}

	return &ApiError{
		Code:     400,
		Title:    "Bad Request",
		Detail:   *detail,
		Type:     "about:blank",
		Instance: "elo", //TODO: Change me to random
	}
}
