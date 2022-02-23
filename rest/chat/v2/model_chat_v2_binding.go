/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV2Binding struct for ChatV2Binding
type ChatV2Binding struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The push technology to use for the binding
	BindingType *string `json:"binding_type,omitempty"`
	// The SID of the Credential for the binding
	CredentialSid *string `json:"credential_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique endpoint identifier for the Binding
	Endpoint *string `json:"endpoint,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"identity,omitempty"`
	// The absolute URLs of the Binding's User
	Links *map[string]interface{} `json:"links,omitempty"`
	// The Programmable Chat message types the binding is subscribed to
	MessageTypes *[]string `json:"message_types,omitempty"`
	// The SID of the Service that the Binding resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Binding resource
	Url *string `json:"url,omitempty"`
}
