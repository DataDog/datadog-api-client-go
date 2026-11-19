// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchEventFinding Code analysis result for an event of the tracking plan.
type ProductAnalyticsLaunchEventFinding struct {
	// Name of the event in the tracking plan.
	Name string `json:"name"`
	// Results for the properties of the event. Only allowed when the event is `found`.
	Properties []ProductAnalyticsLaunchPropertyFinding `json:"properties,omitempty"`
	// Whether the inspected code sends the event, or attaches the property to the event.
	Status ProductAnalyticsLaunchCodeStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchEventFinding instantiates a new ProductAnalyticsLaunchEventFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchEventFinding(name string, status ProductAnalyticsLaunchCodeStatus) *ProductAnalyticsLaunchEventFinding {
	this := ProductAnalyticsLaunchEventFinding{}
	this.Name = name
	this.Status = status
	return &this
}

// NewProductAnalyticsLaunchEventFindingWithDefaults instantiates a new ProductAnalyticsLaunchEventFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchEventFindingWithDefaults() *ProductAnalyticsLaunchEventFinding {
	this := ProductAnalyticsLaunchEventFinding{}
	return &this
}

// GetName returns the Name field value.
func (o *ProductAnalyticsLaunchEventFinding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEventFinding) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsLaunchEventFinding) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEventFinding) GetProperties() []ProductAnalyticsLaunchPropertyFinding {
	if o == nil || o.Properties == nil {
		var ret []ProductAnalyticsLaunchPropertyFinding
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEventFinding) GetPropertiesOk() (*[]ProductAnalyticsLaunchPropertyFinding, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEventFinding) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []ProductAnalyticsLaunchPropertyFinding and assigns it to the Properties field.
func (o *ProductAnalyticsLaunchEventFinding) SetProperties(v []ProductAnalyticsLaunchPropertyFinding) {
	o.Properties = v
}

// GetStatus returns the Status field value.
func (o *ProductAnalyticsLaunchEventFinding) GetStatus() ProductAnalyticsLaunchCodeStatus {
	if o == nil {
		var ret ProductAnalyticsLaunchCodeStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEventFinding) GetStatusOk() (*ProductAnalyticsLaunchCodeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ProductAnalyticsLaunchEventFinding) SetStatus(v ProductAnalyticsLaunchCodeStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchEventFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchEventFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string                                 `json:"name"`
		Properties []ProductAnalyticsLaunchPropertyFinding `json:"properties,omitempty"`
		Status     *ProductAnalyticsLaunchCodeStatus       `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "properties", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Properties = all.Properties
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
