/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateConversationScopedWebhook'
type CreateConversationScopedWebhookParams struct {
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	// The message index for which and it's successors the webhook will be replayed. Not set by default
	ConfigurationReplayAfter *int `json:"Configuration.ReplayAfter,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
	// The target of this webhook: `webhook`, `studio`, `trigger`
	Target *string `json:"Target,omitempty"`
}

func (params *CreateConversationScopedWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *CreateConversationScopedWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *CreateConversationScopedWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *CreateConversationScopedWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationReplayAfter(ConfigurationReplayAfter int) *CreateConversationScopedWebhookParams {
	params.ConfigurationReplayAfter = &ConfigurationReplayAfter
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *CreateConversationScopedWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *CreateConversationScopedWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}
func (params *CreateConversationScopedWebhookParams) SetTarget(Target string) *CreateConversationScopedWebhookParams {
	params.Target = &Target
	return params
}

// Create a new webhook scoped to the conversation
func (c *ApiService) CreateConversationScopedWebhook(ConversationSid string, params *CreateConversationScopedWebhookParams) (*ConversationsV1ConversationScopedWebhook, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationReplayAfter != nil {
		data.Set("Configuration.ReplayAfter", fmt.Sprint(*params.ConfigurationReplayAfter))
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an existing webhook scoped to the conversation
func (c *ApiService) DeleteConversationScopedWebhook(ConversationSid string, Sid string) error {
	path := "/v1/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch the configuration of a conversation-scoped webhook
func (c *ApiService) FetchConversationScopedWebhook(ConversationSid string, Sid string) (*ConversationsV1ConversationScopedWebhook, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversationScopedWebhook'
type ListConversationScopedWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConversationScopedWebhookParams) SetPageSize(PageSize int) *ListConversationScopedWebhookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConversationScopedWebhookParams) SetLimit(Limit int) *ListConversationScopedWebhookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConversationScopedWebhook records from the API. Request is executed immediately.
func (c *ApiService) PageConversationScopedWebhook(ConversationSid string, params *ListConversationScopedWebhookParams, pageToken, pageNumber string) (*ListConversationScopedWebhookResponse, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks"

	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationScopedWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConversationScopedWebhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConversationScopedWebhook(ConversationSid string, params *ListConversationScopedWebhookParams) ([]ConversationsV1ConversationScopedWebhook, error) {
	if params == nil {
		params = &ListConversationScopedWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationScopedWebhook(ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1ConversationScopedWebhook

	for response != nil {
		records = append(records, response.Webhooks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationScopedWebhookResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListConversationScopedWebhookResponse)
	}

	return records, err
}

// Streams ConversationScopedWebhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConversationScopedWebhook(ConversationSid string, params *ListConversationScopedWebhookParams) (chan ConversationsV1ConversationScopedWebhook, error) {
	if params == nil {
		params = &ListConversationScopedWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageConversationScopedWebhook(ConversationSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1ConversationScopedWebhook, 1)

	go func() {
		for response != nil {
			for item := range response.Webhooks {
				channel <- response.Webhooks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListConversationScopedWebhookResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListConversationScopedWebhookResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListConversationScopedWebhookResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationScopedWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConversationScopedWebhook'
type UpdateConversationScopedWebhookParams struct {
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
}

func (params *UpdateConversationScopedWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}

// Update an existing conversation-scoped webhook
func (c *ApiService) UpdateConversationScopedWebhook(ConversationSid string, Sid string, params *UpdateConversationScopedWebhookParams) (*ConversationsV1ConversationScopedWebhook, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
