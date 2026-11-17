// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchEventProperty A reference from an event to a property of the tracking plan.
type ProductAnalyticsLaunchEventProperty struct {
	// ID of a property in the tracking plan `properties`.
	Id string `json:"id"`
	// Whether the event must always include the property.
	Required *bool `json:"required,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchEventProperty instantiates a new ProductAnalyticsLaunchEventProperty object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchEventProperty(id string) *ProductAnalyticsLaunchEventProperty {
	this := ProductAnalyticsLaunchEventProperty{}
	this.Id = id
	return &this
}

// NewProductAnalyticsLaunchEventPropertyWithDefaults instantiates a new ProductAnalyticsLaunchEventProperty object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchEventPropertyWithDefaults() *ProductAnalyticsLaunchEventProperty {
	this := ProductAnalyticsLaunchEventProperty{}
	return &this
}

// GetId returns the Id field value.
func (o *ProductAnalyticsLaunchEventProperty) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEventProperty) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsLaunchEventProperty) SetId(v string) {
	o.Id = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEventProperty) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEventProperty) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEventProperty) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ProductAnalyticsLaunchEventProperty) SetRequired(v bool) {
	o.Required = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchEventProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchEventProperty) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *string `json:"id"`
		Required *bool   `json:"required,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "required"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Required = all.Required

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
