// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRuleTargetConfiguration Configuration for a notification target. Which fields apply depends on the target's `type`.
type NotificationRuleTargetConfiguration struct {
	// Slack channel name, for a `SLACK_CHANNEL` target.
	Channel *string `json:"channel,omitempty"`
	// Slack channel ID for a `SLACK_CHANNEL` target, or Microsoft Teams channel ID
	// for an `MS_TEAMS_CHANNEL` target.
	ChannelId *string `json:"channel_id,omitempty"`
	// Microsoft Teams channel name, for an `MS_TEAMS_CHANNEL` target.
	ChannelName *string `json:"channel_name,omitempty"`
	// Microsoft Teams connector name, for an `MS_TEAMS_CHANNEL` target.
	ConnectorName *string `json:"connector_name,omitempty"`
	// Recipient for an `EMAIL` target.
	Recipient *NotificationRuleTargetConfigurationRecipient `json:"recipient,omitempty"`
	// PagerDuty service name, for a `PAGERDUTY_SERVICE` target.
	ServiceName *string `json:"service_name,omitempty"`
	// Microsoft Teams team ID, for an `MS_TEAMS_CHANNEL` target.
	TeamId *string `json:"team_id,omitempty"`
	// Microsoft Teams team name, for an `MS_TEAMS_CHANNEL` target.
	TeamName *string `json:"team_name,omitempty"`
	// Microsoft Teams tenant ID, for an `MS_TEAMS_CHANNEL` target.
	TenantId *string `json:"tenant_id,omitempty"`
	// Microsoft Teams tenant name, for an `MS_TEAMS_CHANNEL` target.
	TenantName *string `json:"tenant_name,omitempty"`
	// Slack username, for a `SLACK_USER` target.
	Username *string `json:"username,omitempty"`
	// Name of the configured webhook, for a `WEBHOOK` target.
	WebhookName *string `json:"webhook_name,omitempty"`
	// Slack workspace name, for a `SLACK_CHANNEL` or `SLACK_USER` target.
	Workspace *string `json:"workspace,omitempty"`
	// Slack workspace ID, for a `SLACK_CHANNEL` target.
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationRuleTargetConfiguration instantiates a new NotificationRuleTargetConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationRuleTargetConfiguration() *NotificationRuleTargetConfiguration {
	this := NotificationRuleTargetConfiguration{}
	return &this
}

// NewNotificationRuleTargetConfigurationWithDefaults instantiates a new NotificationRuleTargetConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationRuleTargetConfigurationWithDefaults() *NotificationRuleTargetConfiguration {
	this := NotificationRuleTargetConfiguration{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasChannel() bool {
	return o != nil && o.Channel != nil
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *NotificationRuleTargetConfiguration) SetChannel(v string) {
	o.Channel = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *NotificationRuleTargetConfiguration) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasChannelName() bool {
	return o != nil && o.ChannelName != nil
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *NotificationRuleTargetConfiguration) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetConnectorName returns the ConnectorName field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetConnectorName() string {
	if o == nil || o.ConnectorName == nil {
		var ret string
		return ret
	}
	return *o.ConnectorName
}

// GetConnectorNameOk returns a tuple with the ConnectorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetConnectorNameOk() (*string, bool) {
	if o == nil || o.ConnectorName == nil {
		return nil, false
	}
	return o.ConnectorName, true
}

// HasConnectorName returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasConnectorName() bool {
	return o != nil && o.ConnectorName != nil
}

// SetConnectorName gets a reference to the given string and assigns it to the ConnectorName field.
func (o *NotificationRuleTargetConfiguration) SetConnectorName(v string) {
	o.ConnectorName = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetRecipient() NotificationRuleTargetConfigurationRecipient {
	if o == nil || o.Recipient == nil {
		var ret NotificationRuleTargetConfigurationRecipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetRecipientOk() (*NotificationRuleTargetConfigurationRecipient, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasRecipient() bool {
	return o != nil && o.Recipient != nil
}

// SetRecipient gets a reference to the given NotificationRuleTargetConfigurationRecipient and assigns it to the Recipient field.
func (o *NotificationRuleTargetConfiguration) SetRecipient(v NotificationRuleTargetConfigurationRecipient) {
	o.Recipient = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *NotificationRuleTargetConfiguration) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *NotificationRuleTargetConfiguration) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasTeamName() bool {
	return o != nil && o.TeamName != nil
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *NotificationRuleTargetConfiguration) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *NotificationRuleTargetConfiguration) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasTenantName() bool {
	return o != nil && o.TenantName != nil
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *NotificationRuleTargetConfiguration) SetTenantName(v string) {
	o.TenantName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NotificationRuleTargetConfiguration) SetUsername(v string) {
	o.Username = &v
}

// GetWebhookName returns the WebhookName field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetWebhookName() string {
	if o == nil || o.WebhookName == nil {
		var ret string
		return ret
	}
	return *o.WebhookName
}

// GetWebhookNameOk returns a tuple with the WebhookName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetWebhookNameOk() (*string, bool) {
	if o == nil || o.WebhookName == nil {
		return nil, false
	}
	return o.WebhookName, true
}

// HasWebhookName returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasWebhookName() bool {
	return o != nil && o.WebhookName != nil
}

// SetWebhookName gets a reference to the given string and assigns it to the WebhookName field.
func (o *NotificationRuleTargetConfiguration) SetWebhookName(v string) {
	o.WebhookName = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasWorkspace() bool {
	return o != nil && o.Workspace != nil
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *NotificationRuleTargetConfiguration) SetWorkspace(v string) {
	o.Workspace = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *NotificationRuleTargetConfiguration) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRuleTargetConfiguration) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *NotificationRuleTargetConfiguration) HasWorkspaceId() bool {
	return o != nil && o.WorkspaceId != nil
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *NotificationRuleTargetConfiguration) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationRuleTargetConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.ChannelName != nil {
		toSerialize["channel_name"] = o.ChannelName
	}
	if o.ConnectorName != nil {
		toSerialize["connector_name"] = o.ConnectorName
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TeamName != nil {
		toSerialize["team_name"] = o.TeamName
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.TenantName != nil {
		toSerialize["tenant_name"] = o.TenantName
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.WebhookName != nil {
		toSerialize["webhook_name"] = o.WebhookName
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationRuleTargetConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel       *string                                       `json:"channel,omitempty"`
		ChannelId     *string                                       `json:"channel_id,omitempty"`
		ChannelName   *string                                       `json:"channel_name,omitempty"`
		ConnectorName *string                                       `json:"connector_name,omitempty"`
		Recipient     *NotificationRuleTargetConfigurationRecipient `json:"recipient,omitempty"`
		ServiceName   *string                                       `json:"service_name,omitempty"`
		TeamId        *string                                       `json:"team_id,omitempty"`
		TeamName      *string                                       `json:"team_name,omitempty"`
		TenantId      *string                                       `json:"tenant_id,omitempty"`
		TenantName    *string                                       `json:"tenant_name,omitempty"`
		Username      *string                                       `json:"username,omitempty"`
		WebhookName   *string                                       `json:"webhook_name,omitempty"`
		Workspace     *string                                       `json:"workspace,omitempty"`
		WorkspaceId   *string                                       `json:"workspace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel", "channel_id", "channel_name", "connector_name", "recipient", "service_name", "team_id", "team_name", "tenant_id", "tenant_name", "username", "webhook_name", "workspace", "workspace_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Channel = all.Channel
	o.ChannelId = all.ChannelId
	o.ChannelName = all.ChannelName
	o.ConnectorName = all.ConnectorName
	if all.Recipient != nil && all.Recipient.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Recipient = all.Recipient
	o.ServiceName = all.ServiceName
	o.TeamId = all.TeamId
	o.TeamName = all.TeamName
	o.TenantId = all.TenantId
	o.TenantName = all.TenantName
	o.Username = all.Username
	o.WebhookName = all.WebhookName
	o.Workspace = all.Workspace
	o.WorkspaceId = all.WorkspaceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
