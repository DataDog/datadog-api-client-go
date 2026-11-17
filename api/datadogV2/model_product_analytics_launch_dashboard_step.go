// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchDashboardStep Dashboard linked to the launch.
type ProductAnalyticsLaunchDashboardStep struct {
	// ID of the linked dashboard.
	DashboardId datadog.NullableString `json:"dashboard_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchDashboardStep instantiates a new ProductAnalyticsLaunchDashboardStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchDashboardStep() *ProductAnalyticsLaunchDashboardStep {
	this := ProductAnalyticsLaunchDashboardStep{}
	return &this
}

// NewProductAnalyticsLaunchDashboardStepWithDefaults instantiates a new ProductAnalyticsLaunchDashboardStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchDashboardStepWithDefaults() *ProductAnalyticsLaunchDashboardStep {
	this := ProductAnalyticsLaunchDashboardStep{}
	return &this
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchDashboardStep) GetDashboardId() string {
	if o == nil || o.DashboardId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DashboardId.Get()
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchDashboardStep) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DashboardId.Get(), o.DashboardId.IsSet()
}

// HasDashboardId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchDashboardStep) HasDashboardId() bool {
	return o != nil && o.DashboardId.IsSet()
}

// SetDashboardId gets a reference to the given datadog.NullableString and assigns it to the DashboardId field.
func (o *ProductAnalyticsLaunchDashboardStep) SetDashboardId(v string) {
	o.DashboardId.Set(&v)
}

// SetDashboardIdNil sets the value for DashboardId to be an explicit nil.
func (o *ProductAnalyticsLaunchDashboardStep) SetDashboardIdNil() {
	o.DashboardId.Set(nil)
}

// UnsetDashboardId ensures that no value is present for DashboardId, not even an explicit nil.
func (o *ProductAnalyticsLaunchDashboardStep) UnsetDashboardId() {
	o.DashboardId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchDashboardStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DashboardId.IsSet() {
		toSerialize["dashboard_id"] = o.DashboardId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchDashboardStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DashboardId datadog.NullableString `json:"dashboard_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dashboard_id"})
	} else {
		return err
	}
	o.DashboardId = all.DashboardId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
