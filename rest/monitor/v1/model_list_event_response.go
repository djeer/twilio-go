/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEventResponse struct for ListEventResponse
type ListEventResponse struct {
	Events []MonitorV1Event      `json:"events,omitempty"`
	Meta   ListAlertResponseMeta `json:"meta,omitempty"`
}
