// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentOrgSettingsDataAttributesResponse Attributes of an incident org settings resource in a response.
type IncidentOrgSettingsDataAttributesResponse struct {
	// Timestamp when the settings were created.
	Created time.Time `json:"created"`
	// Timestamp when the settings were last modified.
	Modified time.Time `json:"modified"`
	// The settings configuration for an incident org settings resource.
	Settings map[string]interface{} `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentOrgSettingsDataAttributesResponse instantiates a new IncidentOrgSettingsDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentOrgSettingsDataAttributesResponse(created time.Time, modified time.Time, settings map[string]interface{}) *IncidentOrgSettingsDataAttributesResponse {
	this := IncidentOrgSettingsDataAttributesResponse{}
	this.Created = created
	this.Modified = modified
	this.Settings = settings
	return &this
}

// NewIncidentOrgSettingsDataAttributesResponseWithDefaults instantiates a new IncidentOrgSettingsDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentOrgSettingsDataAttributesResponseWithDefaults() *IncidentOrgSettingsDataAttributesResponse {
	this := IncidentOrgSettingsDataAttributesResponse{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentOrgSettingsDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentOrgSettingsDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentOrgSettingsDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value.
func (o *IncidentOrgSettingsDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentOrgSettingsDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentOrgSettingsDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// GetSettings returns the Settings field value.
func (o *IncidentOrgSettingsDataAttributesResponse) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *IncidentOrgSettingsDataAttributesResponse) GetSettingsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *IncidentOrgSettingsDataAttributesResponse) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentOrgSettingsDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentOrgSettingsDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created  *time.Time              `json:"created"`
		Modified *time.Time              `json:"modified"`
		Settings *map[string]interface{} `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "modified", "settings"})
	} else {
		return err
	}
	o.Created = *all.Created
	o.Modified = *all.Modified
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
