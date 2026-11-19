// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorUnlinkAttributes The chart link to remove. It must match the stored link exactly:
// `chart_id` alone for a saved chart, or `chart_id` and `dashboard_id` for a dashboard widget.
type ProductAnalyticsMonitorUnlinkAttributes struct {
	// ID of the chart the monitor is linked to.
	ChartId string `json:"chart_id"`
	// ID of the dashboard the monitor is linked to.
	DashboardId *string `json:"dashboard_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorUnlinkAttributes instantiates a new ProductAnalyticsMonitorUnlinkAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorUnlinkAttributes(chartId string) *ProductAnalyticsMonitorUnlinkAttributes {
	this := ProductAnalyticsMonitorUnlinkAttributes{}
	this.ChartId = chartId
	return &this
}

// NewProductAnalyticsMonitorUnlinkAttributesWithDefaults instantiates a new ProductAnalyticsMonitorUnlinkAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorUnlinkAttributesWithDefaults() *ProductAnalyticsMonitorUnlinkAttributes {
	this := ProductAnalyticsMonitorUnlinkAttributes{}
	return &this
}

// GetChartId returns the ChartId field value.
func (o *ProductAnalyticsMonitorUnlinkAttributes) GetChartId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChartId
}

// GetChartIdOk returns a tuple with the ChartId field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUnlinkAttributes) GetChartIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartId, true
}

// SetChartId sets field value.
func (o *ProductAnalyticsMonitorUnlinkAttributes) SetChartId(v string) {
	o.ChartId = v
}

// GetDashboardId returns the DashboardId field value if set, zero value otherwise.
func (o *ProductAnalyticsMonitorUnlinkAttributes) GetDashboardId() string {
	if o == nil || o.DashboardId == nil {
		var ret string
		return ret
	}
	return *o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUnlinkAttributes) GetDashboardIdOk() (*string, bool) {
	if o == nil || o.DashboardId == nil {
		return nil, false
	}
	return o.DashboardId, true
}

// HasDashboardId returns a boolean if a field has been set.
func (o *ProductAnalyticsMonitorUnlinkAttributes) HasDashboardId() bool {
	return o != nil && o.DashboardId != nil
}

// SetDashboardId gets a reference to the given string and assigns it to the DashboardId field.
func (o *ProductAnalyticsMonitorUnlinkAttributes) SetDashboardId(v string) {
	o.DashboardId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorUnlinkAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["chart_id"] = o.ChartId
	if o.DashboardId != nil {
		toSerialize["dashboard_id"] = o.DashboardId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorUnlinkAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChartId     *string `json:"chart_id"`
		DashboardId *string `json:"dashboard_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChartId == nil {
		return fmt.Errorf("required field chart_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"chart_id", "dashboard_id"})
	} else {
		return err
	}
	o.ChartId = *all.ChartId
	o.DashboardId = all.DashboardId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
