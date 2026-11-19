// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchKPI A KPI that the launch is meant to move.
type ProductAnalyticsLaunchKPI struct {
	// ID of the KPI, unique within the launch. Tracking plan events reference it.
	Id string `json:"id"`
	// Name of the KPI.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchKPI instantiates a new ProductAnalyticsLaunchKPI object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchKPI(id string, title string) *ProductAnalyticsLaunchKPI {
	this := ProductAnalyticsLaunchKPI{}
	this.Id = id
	this.Title = title
	return &this
}

// NewProductAnalyticsLaunchKPIWithDefaults instantiates a new ProductAnalyticsLaunchKPI object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchKPIWithDefaults() *ProductAnalyticsLaunchKPI {
	this := ProductAnalyticsLaunchKPI{}
	return &this
}

// GetId returns the Id field value.
func (o *ProductAnalyticsLaunchKPI) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchKPI) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsLaunchKPI) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value.
func (o *ProductAnalyticsLaunchKPI) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchKPI) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ProductAnalyticsLaunchKPI) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchKPI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchKPI) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id    *string `json:"id"`
		Title *string `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "title"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
