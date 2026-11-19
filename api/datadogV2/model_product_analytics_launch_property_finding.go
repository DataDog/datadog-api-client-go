// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchPropertyFinding Code analysis result for a property of an event.
type ProductAnalyticsLaunchPropertyFinding struct {
	// ID of the property in the tracking plan.
	Id string `json:"id"`
	// Whether the inspected code sends the event, or attaches the property to the event.
	Status ProductAnalyticsLaunchCodeStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchPropertyFinding instantiates a new ProductAnalyticsLaunchPropertyFinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchPropertyFinding(id string, status ProductAnalyticsLaunchCodeStatus) *ProductAnalyticsLaunchPropertyFinding {
	this := ProductAnalyticsLaunchPropertyFinding{}
	this.Id = id
	this.Status = status
	return &this
}

// NewProductAnalyticsLaunchPropertyFindingWithDefaults instantiates a new ProductAnalyticsLaunchPropertyFinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchPropertyFindingWithDefaults() *ProductAnalyticsLaunchPropertyFinding {
	this := ProductAnalyticsLaunchPropertyFinding{}
	return &this
}

// GetId returns the Id field value.
func (o *ProductAnalyticsLaunchPropertyFinding) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPropertyFinding) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsLaunchPropertyFinding) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value.
func (o *ProductAnalyticsLaunchPropertyFinding) GetStatus() ProductAnalyticsLaunchCodeStatus {
	if o == nil {
		var ret ProductAnalyticsLaunchCodeStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPropertyFinding) GetStatusOk() (*ProductAnalyticsLaunchCodeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ProductAnalyticsLaunchPropertyFinding) SetStatus(v ProductAnalyticsLaunchCodeStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchPropertyFinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchPropertyFinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                           `json:"id"`
		Status *ProductAnalyticsLaunchCodeStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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
