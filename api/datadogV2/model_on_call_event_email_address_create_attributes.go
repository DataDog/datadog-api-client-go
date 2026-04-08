// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallEventEmailAddressCreateAttributes Attributes for creating an on-call event email address.
type OnCallEventEmailAddressCreateAttributes struct {
	// The alert type of events generated from the email address.
	AlertType *EventEmailAddressAlertType `json:"alert_type,omitempty"`
	// A description of the on-call event email address.
	Description *string `json:"description,omitempty"`
	// The format of events ingested through the email address.
	Format EventEmailAddressFormat `json:"format"`
	// A list of tags to apply to events generated from the email address.
	Tags []string `json:"tags,omitempty"`
	// The team handle associated with the on-call email address.
	// Must contain only alphanumeric characters, hyphens, and underscores.
	TeamHandle string `json:"team_handle"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallEventEmailAddressCreateAttributes instantiates a new OnCallEventEmailAddressCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallEventEmailAddressCreateAttributes(format EventEmailAddressFormat, teamHandle string) *OnCallEventEmailAddressCreateAttributes {
	this := OnCallEventEmailAddressCreateAttributes{}
	this.Format = format
	this.TeamHandle = teamHandle
	return &this
}

// NewOnCallEventEmailAddressCreateAttributesWithDefaults instantiates a new OnCallEventEmailAddressCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallEventEmailAddressCreateAttributesWithDefaults() *OnCallEventEmailAddressCreateAttributes {
	this := OnCallEventEmailAddressCreateAttributes{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *OnCallEventEmailAddressCreateAttributes) GetAlertType() EventEmailAddressAlertType {
	if o == nil || o.AlertType == nil {
		var ret EventEmailAddressAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallEventEmailAddressCreateAttributes) GetAlertTypeOk() (*EventEmailAddressAlertType, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *OnCallEventEmailAddressCreateAttributes) HasAlertType() bool {
	return o != nil && o.AlertType != nil
}

// SetAlertType gets a reference to the given EventEmailAddressAlertType and assigns it to the AlertType field.
func (o *OnCallEventEmailAddressCreateAttributes) SetAlertType(v EventEmailAddressAlertType) {
	o.AlertType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OnCallEventEmailAddressCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallEventEmailAddressCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OnCallEventEmailAddressCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OnCallEventEmailAddressCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFormat returns the Format field value.
func (o *OnCallEventEmailAddressCreateAttributes) GetFormat() EventEmailAddressFormat {
	if o == nil {
		var ret EventEmailAddressFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *OnCallEventEmailAddressCreateAttributes) GetFormatOk() (*EventEmailAddressFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *OnCallEventEmailAddressCreateAttributes) SetFormat(v EventEmailAddressFormat) {
	o.Format = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *OnCallEventEmailAddressCreateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallEventEmailAddressCreateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *OnCallEventEmailAddressCreateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *OnCallEventEmailAddressCreateAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeamHandle returns the TeamHandle field value.
func (o *OnCallEventEmailAddressCreateAttributes) GetTeamHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamHandle
}

// GetTeamHandleOk returns a tuple with the TeamHandle field value
// and a boolean to check if the value has been set.
func (o *OnCallEventEmailAddressCreateAttributes) GetTeamHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamHandle, true
}

// SetTeamHandle sets field value.
func (o *OnCallEventEmailAddressCreateAttributes) SetTeamHandle(v string) {
	o.TeamHandle = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallEventEmailAddressCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlertType != nil {
		toSerialize["alert_type"] = o.AlertType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["format"] = o.Format
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["team_handle"] = o.TeamHandle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallEventEmailAddressCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertType   *EventEmailAddressAlertType `json:"alert_type,omitempty"`
		Description *string                     `json:"description,omitempty"`
		Format      *EventEmailAddressFormat    `json:"format"`
		Tags        []string                    `json:"tags,omitempty"`
		TeamHandle  *string                     `json:"team_handle"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.TeamHandle == nil {
		return fmt.Errorf("required field team_handle missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_type", "description", "format", "tags", "team_handle"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AlertType != nil && !all.AlertType.IsValid() {
		hasInvalidField = true
	} else {
		o.AlertType = all.AlertType
	}
	o.Description = all.Description
	if !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = *all.Format
	}
	o.Tags = all.Tags
	o.TeamHandle = *all.TeamHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
