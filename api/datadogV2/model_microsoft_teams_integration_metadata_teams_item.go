// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsIntegrationMetadataTeamsItem Item in the Microsoft Teams integration metadata teams array.
type MicrosoftTeamsIntegrationMetadataTeamsItem struct {
	// Team ID of the Microsoft Teams team
	MsTeamId string `json:"ms_team_id"`
	// Microsoft Teams tenant ID
	MsTenantId string `json:"ms_tenant_id"`
	// URL redirecting to the Microsoft Teams team
	RedirectUrl string `json:"redirect_url"`
	// Name of the Microsoft Teams team
	TeamName string `json:"team_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMicrosoftTeamsIntegrationMetadataTeamsItem instantiates a new MicrosoftTeamsIntegrationMetadataTeamsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsIntegrationMetadataTeamsItem(msTeamId string, msTenantId string, redirectUrl string, teamName string) *MicrosoftTeamsIntegrationMetadataTeamsItem {
	this := MicrosoftTeamsIntegrationMetadataTeamsItem{}
	this.MsTeamId = msTeamId
	this.MsTenantId = msTenantId
	this.RedirectUrl = redirectUrl
	this.TeamName = teamName
	return &this
}

// NewMicrosoftTeamsIntegrationMetadataTeamsItemWithDefaults instantiates a new MicrosoftTeamsIntegrationMetadataTeamsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsIntegrationMetadataTeamsItemWithDefaults() *MicrosoftTeamsIntegrationMetadataTeamsItem {
	this := MicrosoftTeamsIntegrationMetadataTeamsItem{}
	return &this
}

// GetMsTeamId returns the MsTeamId field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetMsTeamId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsTeamId
}

// GetMsTeamIdOk returns a tuple with the MsTeamId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetMsTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsTeamId, true
}

// SetMsTeamId sets field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) SetMsTeamId(v string) {
	o.MsTeamId = v
}

// GetMsTenantId returns the MsTenantId field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetMsTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MsTenantId
}

// GetMsTenantIdOk returns a tuple with the MsTenantId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetMsTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsTenantId, true
}

// SetMsTenantId sets field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) SetMsTenantId(v string) {
	o.MsTenantId = v
}

// GetRedirectUrl returns the RedirectUrl field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

// GetTeamName returns the TeamName field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetTeamName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) GetTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamName, true
}

// SetTeamName sets field value.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) SetTeamName(v string) {
	o.TeamName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsIntegrationMetadataTeamsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ms_team_id"] = o.MsTeamId
	toSerialize["ms_tenant_id"] = o.MsTenantId
	toSerialize["redirect_url"] = o.RedirectUrl
	toSerialize["team_name"] = o.TeamName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsIntegrationMetadataTeamsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MsTeamId    *string `json:"ms_team_id"`
		MsTenantId  *string `json:"ms_tenant_id"`
		RedirectUrl *string `json:"redirect_url"`
		TeamName    *string `json:"team_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MsTeamId == nil {
		return fmt.Errorf("required field ms_team_id missing")
	}
	if all.MsTenantId == nil {
		return fmt.Errorf("required field ms_tenant_id missing")
	}
	if all.RedirectUrl == nil {
		return fmt.Errorf("required field redirect_url missing")
	}
	if all.TeamName == nil {
		return fmt.Errorf("required field team_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ms_team_id", "ms_tenant_id", "redirect_url", "team_name"})
	} else {
		return err
	}
	o.MsTeamId = *all.MsTeamId
	o.MsTenantId = *all.MsTenantId
	o.RedirectUrl = *all.RedirectUrl
	o.TeamName = *all.TeamName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
