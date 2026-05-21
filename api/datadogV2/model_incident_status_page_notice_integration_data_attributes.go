// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatusPageNoticeIntegrationDataAttributes Attributes of a status page notice integration.
type IncidentStatusPageNoticeIntegrationDataAttributes struct {
	// Timestamp when the integration was created.
	Created time.Time `json:"created"`
	// The type of integration.
	IntegrationType string `json:"integration_type"`
	// Timestamp when the integration was last modified.
	Modified time.Time `json:"modified"`
	// The status of the integration.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatusPageNoticeIntegrationDataAttributes instantiates a new IncidentStatusPageNoticeIntegrationDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatusPageNoticeIntegrationDataAttributes(created time.Time, integrationType string, modified time.Time, status string) *IncidentStatusPageNoticeIntegrationDataAttributes {
	this := IncidentStatusPageNoticeIntegrationDataAttributes{}
	this.Created = created
	this.IntegrationType = integrationType
	this.Modified = modified
	this.Status = status
	return &this
}

// NewIncidentStatusPageNoticeIntegrationDataAttributesWithDefaults instantiates a new IncidentStatusPageNoticeIntegrationDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatusPageNoticeIntegrationDataAttributesWithDefaults() *IncidentStatusPageNoticeIntegrationDataAttributes {
	this := IncidentStatusPageNoticeIntegrationDataAttributes{}
	return &this
}

// GetCreated returns the Created field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) SetCreated(v time.Time) {
	o.Created = v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) SetIntegrationType(v string) {
	o.IntegrationType = v
}

// GetModified returns the Modified field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) SetModified(v time.Time) {
	o.Modified = v
}

// GetStatus returns the Status field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatusPageNoticeIntegrationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["integration_type"] = o.IntegrationType
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatusPageNoticeIntegrationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created         *time.Time `json:"created"`
		IntegrationType *string    `json:"integration_type"`
		Modified        *time.Time `json:"modified"`
		Status          *string    `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "integration_type", "modified", "status"})
	} else {
		return err
	}
	o.Created = *all.Created
	o.IntegrationType = *all.IntegrationType
	o.Modified = *all.Modified
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
