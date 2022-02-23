/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TrunkingV1Recording struct for TrunkingV1Recording
type TrunkingV1Recording struct {
	// The recording mode for the trunk.
	Mode *string `json:"mode,omitempty"`
	// The recording trim setting for the trunk.
	Trim *string `json:"trim,omitempty"`
}
