// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchProperty An event property of the tracking plan.
type ProductAnalyticsLaunchProperty struct {
	// What the property captures.
	Description string `json:"description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchProperty instantiates a new ProductAnalyticsLaunchProperty object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchProperty(description string) *ProductAnalyticsLaunchProperty {
	this := ProductAnalyticsLaunchProperty{}
	this.Description = description
	return &this
}

// NewProductAnalyticsLaunchPropertyWithDefaults instantiates a new ProductAnalyticsLaunchProperty object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchPropertyWithDefaults() *ProductAnalyticsLaunchProperty {
	this := ProductAnalyticsLaunchProperty{}
	return &this
}

// GetDescription returns the Description field value.
func (o *ProductAnalyticsLaunchProperty) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchProperty) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ProductAnalyticsLaunchProperty) SetDescription(v string) {
	o.Description = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchProperty) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description"})
	} else {
		return err
	}
	o.Description = *all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
