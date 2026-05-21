// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentMSTeamsChannel A Microsoft Teams channel associated with an incident.
type IncidentMSTeamsChannel struct {
	// The Teams channel identifier.
	MsChannelId string `json:"ms_channel_id"`
	// The name of the Teams channel.
	MsChannelName string `json:"ms_channel_name"`
	// The Teams team identifier.
	MsTeamId *string `json:"ms_team_id,omitempty"`
	// The Teams tenant identifier.
	MsTenantId *string `json:"ms_tenant_id,omitempty"`
	// The redirect URL for the channel.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// The name of the Teams team.
	TeamName *string `json:"team_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentMSTeamsChannel instantiates a new IncidentMSTeamsChannel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentMSTeamsChannel(msChannelId string, msChannelName string) *IncidentMSTeamsChannel {
	this := IncidentMSTeamsChannel{}
	this.MsChannelId = msChannelId
	this.MsChannelName = msChannelName
	return &this
}

// NewIncidentMSTeamsChannelWithDefaults instantiates a new IncidentMSTeamsChannel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentMSTeamsChannelWithDefaults() *IncidentMSTeamsChannel {
	this := IncidentMSTeamsChannel{}
	return &this
}

// GetMsChannelId returns the MsChannelId field value.
func (o *IncidentMSTeamsChannel) GetMsChannelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsChannelId
}

// GetMsChannelIdOk returns a tuple with the MsChannelId field value
// and a boolean to check if the value has been set.
func (o *IncidentMSTeamsChannel) GetMsChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsChannelId, true
}

// SetMsChannelId sets field value.
func (o *IncidentMSTeamsChannel) SetMsChannelId(v string) {
	o.MsChannelId = v
}

// GetMsChannelName returns the MsChannelName field value.
func (o *IncidentMSTeamsChannel) GetMsChannelName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsChannelName
}

// GetMsChannelNameOk returns a tuple with the MsChannelName field value
// and a boolean to check if the value has been set.
func (o *IncidentMSTeamsChannel) GetMsChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsChannelName, true
}

// SetMsChannelName sets field value.
func (o *IncidentMSTeamsChannel) SetMsChannelName(v string) {
	o.MsChannelName = v
}

// GetMsTeamId returns the MsTeamId field value if set, zero value otherwise.
func (o *IncidentMSTeamsChannel) GetMsTeamId() string {
	if o == nil || o.MsTeamId == nil {
		var ret string
		return ret
	}
	return *o.MsTeamId
}

// GetMsTeamIdOk returns a tuple with the MsTeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentMSTeamsChannel) GetMsTeamIdOk() (*string, bool) {
	if o == nil || o.MsTeamId == nil {
		return nil, false
	}
	return o.MsTeamId, true
}

// HasMsTeamId returns a boolean if a field has been set.
func (o *IncidentMSTeamsChannel) HasMsTeamId() bool {
	return o != nil && o.MsTeamId != nil
}

// SetMsTeamId gets a reference to the given string and assigns it to the MsTeamId field.
func (o *IncidentMSTeamsChannel) SetMsTeamId(v string) {
	o.MsTeamId = &v
}

// GetMsTenantId returns the MsTenantId field value if set, zero value otherwise.
func (o *IncidentMSTeamsChannel) GetMsTenantId() string {
	if o == nil || o.MsTenantId == nil {
		var ret string
		return ret
	}
	return *o.MsTenantId
}

// GetMsTenantIdOk returns a tuple with the MsTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentMSTeamsChannel) GetMsTenantIdOk() (*string, bool) {
	if o == nil || o.MsTenantId == nil {
		return nil, false
	}
	return o.MsTenantId, true
}

// HasMsTenantId returns a boolean if a field has been set.
func (o *IncidentMSTeamsChannel) HasMsTenantId() bool {
	return o != nil && o.MsTenantId != nil
}

// SetMsTenantId gets a reference to the given string and assigns it to the MsTenantId field.
func (o *IncidentMSTeamsChannel) SetMsTenantId(v string) {
	o.MsTenantId = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *IncidentMSTeamsChannel) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentMSTeamsChannel) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *IncidentMSTeamsChannel) HasRedirectUrl() bool {
	return o != nil && o.RedirectUrl != nil
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *IncidentMSTeamsChannel) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *IncidentMSTeamsChannel) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentMSTeamsChannel) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *IncidentMSTeamsChannel) HasTeamName() bool {
	return o != nil && o.TeamName != nil
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *IncidentMSTeamsChannel) SetTeamName(v string) {
	o.TeamName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentMSTeamsChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ms_channel_id"] = o.MsChannelId
	toSerialize["ms_channel_name"] = o.MsChannelName
	if o.MsTeamId != nil {
		toSerialize["ms_team_id"] = o.MsTeamId
	}
	if o.MsTenantId != nil {
		toSerialize["ms_tenant_id"] = o.MsTenantId
	}
	if o.RedirectUrl != nil {
		toSerialize["redirect_url"] = o.RedirectUrl
	}
	if o.TeamName != nil {
		toSerialize["team_name"] = o.TeamName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentMSTeamsChannel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MsChannelId   *string `json:"ms_channel_id"`
		MsChannelName *string `json:"ms_channel_name"`
		MsTeamId      *string `json:"ms_team_id,omitempty"`
		MsTenantId    *string `json:"ms_tenant_id,omitempty"`
		RedirectUrl   *string `json:"redirect_url,omitempty"`
		TeamName      *string `json:"team_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MsChannelId == nil {
		return fmt.Errorf("required field ms_channel_id missing")
	}
	if all.MsChannelName == nil {
		return fmt.Errorf("required field ms_channel_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ms_channel_id", "ms_channel_name", "ms_team_id", "ms_tenant_id", "redirect_url", "team_name"})
	} else {
		return err
	}
	o.MsChannelId = *all.MsChannelId
	o.MsChannelName = *all.MsChannelName
	o.MsTeamId = all.MsTeamId
	o.MsTenantId = all.MsTenantId
	o.RedirectUrl = all.RedirectUrl
	o.TeamName = all.TeamName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
