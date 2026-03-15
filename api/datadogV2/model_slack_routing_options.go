// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackRoutingOptions Slack routing options for report delivery.
type SlackRoutingOptions struct {
	// Slack channel ID.
	ChannelId *string `json:"channel_id,omitempty"`
	// Slack channel name.
	ChannelName *string `json:"channel_name,omitempty"`
	// Slack workspace ID.
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// Slack workspace name.
	WorkspaceName *string `json:"workspace_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackRoutingOptions instantiates a new SlackRoutingOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackRoutingOptions() *SlackRoutingOptions {
	this := SlackRoutingOptions{}
	return &this
}

// NewSlackRoutingOptionsWithDefaults instantiates a new SlackRoutingOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackRoutingOptionsWithDefaults() *SlackRoutingOptions {
	this := SlackRoutingOptions{}
	return &this
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *SlackRoutingOptions) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackRoutingOptions) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *SlackRoutingOptions) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *SlackRoutingOptions) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *SlackRoutingOptions) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackRoutingOptions) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *SlackRoutingOptions) HasChannelName() bool {
	return o != nil && o.ChannelName != nil
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *SlackRoutingOptions) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *SlackRoutingOptions) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackRoutingOptions) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *SlackRoutingOptions) HasWorkspaceId() bool {
	return o != nil && o.WorkspaceId != nil
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *SlackRoutingOptions) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetWorkspaceName returns the WorkspaceName field value if set, zero value otherwise.
func (o *SlackRoutingOptions) GetWorkspaceName() string {
	if o == nil || o.WorkspaceName == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceName
}

// GetWorkspaceNameOk returns a tuple with the WorkspaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackRoutingOptions) GetWorkspaceNameOk() (*string, bool) {
	if o == nil || o.WorkspaceName == nil {
		return nil, false
	}
	return o.WorkspaceName, true
}

// HasWorkspaceName returns a boolean if a field has been set.
func (o *SlackRoutingOptions) HasWorkspaceName() bool {
	return o != nil && o.WorkspaceName != nil
}

// SetWorkspaceName gets a reference to the given string and assigns it to the WorkspaceName field.
func (o *SlackRoutingOptions) SetWorkspaceName(v string) {
	o.WorkspaceName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackRoutingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.ChannelName != nil {
		toSerialize["channel_name"] = o.ChannelName
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	if o.WorkspaceName != nil {
		toSerialize["workspace_name"] = o.WorkspaceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackRoutingOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChannelId     *string `json:"channel_id,omitempty"`
		ChannelName   *string `json:"channel_name,omitempty"`
		WorkspaceId   *string `json:"workspace_id,omitempty"`
		WorkspaceName *string `json:"workspace_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel_id", "channel_name", "workspace_id", "workspace_name"})
	} else {
		return err
	}
	o.ChannelId = all.ChannelId
	o.ChannelName = all.ChannelName
	o.WorkspaceId = all.WorkspaceId
	o.WorkspaceName = all.WorkspaceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
