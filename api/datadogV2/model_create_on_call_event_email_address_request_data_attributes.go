// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateOnCallEventEmailAddressRequestDataAttributes
type CreateOnCallEventEmailAddressRequestDataAttributes struct {
	//
	AlertType *string `json:"alert_type,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	Format string `json:"format"`
	//
	Tags []string `json:"tags"`
	//
	TeamHandle *string `json:"team_handle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateOnCallEventEmailAddressRequestDataAttributes instantiates a new CreateOnCallEventEmailAddressRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateOnCallEventEmailAddressRequestDataAttributes(format string, tags []string) *CreateOnCallEventEmailAddressRequestDataAttributes {
	this := CreateOnCallEventEmailAddressRequestDataAttributes{}
	this.Format = format
	this.Tags = tags
	return &this
}

// NewCreateOnCallEventEmailAddressRequestDataAttributesWithDefaults instantiates a new CreateOnCallEventEmailAddressRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateOnCallEventEmailAddressRequestDataAttributesWithDefaults() *CreateOnCallEventEmailAddressRequestDataAttributes {
	this := CreateOnCallEventEmailAddressRequestDataAttributes{}
	return &this
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetAlertType() string {
	if o == nil || o.AlertType == nil {
		var ret string
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetAlertTypeOk() (*string, bool) {
	if o == nil || o.AlertType == nil {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) HasAlertType() bool {
	return o != nil && o.AlertType != nil
}

// SetAlertType gets a reference to the given string and assigns it to the AlertType field.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) SetAlertType(v string) {
	o.AlertType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFormat returns the Format field value.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) SetFormat(v string) {
	o.Format = v
}

// GetTags returns the Tags field value.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTeamHandle returns the TeamHandle field value if set, zero value otherwise.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetTeamHandle() string {
	if o == nil || o.TeamHandle == nil {
		var ret string
		return ret
	}
	return *o.TeamHandle
}

// GetTeamHandleOk returns a tuple with the TeamHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) GetTeamHandleOk() (*string, bool) {
	if o == nil || o.TeamHandle == nil {
		return nil, false
	}
	return o.TeamHandle, true
}

// HasTeamHandle returns a boolean if a field has been set.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) HasTeamHandle() bool {
	return o != nil && o.TeamHandle != nil
}

// SetTeamHandle gets a reference to the given string and assigns it to the TeamHandle field.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) SetTeamHandle(v string) {
	o.TeamHandle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateOnCallEventEmailAddressRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["tags"] = o.Tags
	if o.TeamHandle != nil {
		toSerialize["team_handle"] = o.TeamHandle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateOnCallEventEmailAddressRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertType   *string   `json:"alert_type,omitempty"`
		Description *string   `json:"description,omitempty"`
		Format      *string   `json:"format"`
		Tags        *[]string `json:"tags"`
		TeamHandle  *string   `json:"team_handle,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alert_type", "description", "format", "tags", "team_handle"})
	} else {
		return err
	}
	o.AlertType = all.AlertType
	o.Description = all.Description
	o.Format = *all.Format
	o.Tags = *all.Tags
	o.TeamHandle = all.TeamHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
