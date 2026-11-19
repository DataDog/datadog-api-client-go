// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorListAttributes A list of Product Analytics monitors.
type ProductAnalyticsMonitorListAttributes struct {
	// The monitors.
	Monitors []ProductAnalyticsMonitorListItem `json:"monitors"`
	// Total number of monitors that match, across all pages.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorListAttributes instantiates a new ProductAnalyticsMonitorListAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorListAttributes(monitors []ProductAnalyticsMonitorListItem, total int64) *ProductAnalyticsMonitorListAttributes {
	this := ProductAnalyticsMonitorListAttributes{}
	this.Monitors = monitors
	this.Total = total
	return &this
}

// NewProductAnalyticsMonitorListAttributesWithDefaults instantiates a new ProductAnalyticsMonitorListAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorListAttributesWithDefaults() *ProductAnalyticsMonitorListAttributes {
	this := ProductAnalyticsMonitorListAttributes{}
	return &this
}

// GetMonitors returns the Monitors field value.
func (o *ProductAnalyticsMonitorListAttributes) GetMonitors() []ProductAnalyticsMonitorListItem {
	if o == nil {
		var ret []ProductAnalyticsMonitorListItem
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorListAttributes) GetMonitorsOk() (*[]ProductAnalyticsMonitorListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Monitors, true
}

// SetMonitors sets field value.
func (o *ProductAnalyticsMonitorListAttributes) SetMonitors(v []ProductAnalyticsMonitorListItem) {
	o.Monitors = v
}

// GetTotal returns the Total field value.
func (o *ProductAnalyticsMonitorListAttributes) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorListAttributes) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *ProductAnalyticsMonitorListAttributes) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorListAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["monitors"] = o.Monitors
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorListAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Monitors *[]ProductAnalyticsMonitorListItem `json:"monitors"`
		Total    *int64                             `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Monitors == nil {
		return fmt.Errorf("required field monitors missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"monitors", "total"})
	} else {
		return err
	}
	o.Monitors = *all.Monitors
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
