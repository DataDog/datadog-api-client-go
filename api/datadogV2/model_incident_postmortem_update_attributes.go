// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentPostmortemUpdateAttributes The postmortem's attributes for an update request.
type IncidentPostmortemUpdateAttributes struct {
	// The status of the postmortem.
	Status *PostmortemStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentPostmortemUpdateAttributes instantiates a new IncidentPostmortemUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentPostmortemUpdateAttributes() *IncidentPostmortemUpdateAttributes {
	this := IncidentPostmortemUpdateAttributes{}
	return &this
}

// NewIncidentPostmortemUpdateAttributesWithDefaults instantiates a new IncidentPostmortemUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentPostmortemUpdateAttributesWithDefaults() *IncidentPostmortemUpdateAttributes {
	this := IncidentPostmortemUpdateAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IncidentPostmortemUpdateAttributes) GetStatus() PostmortemStatus {
	if o == nil || o.Status == nil {
		var ret PostmortemStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentPostmortemUpdateAttributes) GetStatusOk() (*PostmortemStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IncidentPostmortemUpdateAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PostmortemStatus and assigns it to the Status field.
func (o *IncidentPostmortemUpdateAttributes) SetStatus(v PostmortemStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentPostmortemUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentPostmortemUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *PostmortemStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
