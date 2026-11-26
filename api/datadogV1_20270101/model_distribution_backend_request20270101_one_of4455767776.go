// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455767776 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455767776 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ProfileMetricsQuery LogQueryDefinition `json:"profile_metrics_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455767776 instantiates a new DistributionBackendRequest20270101OneOf4455767776 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455767776(profileMetricsQuery LogQueryDefinition) *DistributionBackendRequest20270101OneOf4455767776 {
	this := DistributionBackendRequest20270101OneOf4455767776{}
	this.ProfileMetricsQuery = profileMetricsQuery
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455767776WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455767776 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455767776WithDefaults() *DistributionBackendRequest20270101OneOf4455767776 {
	this := DistributionBackendRequest20270101OneOf4455767776{}
	return &this
}

// GetProfileMetricsQuery returns the ProfileMetricsQuery field value.
func (o *DistributionBackendRequest20270101OneOf4455767776) GetProfileMetricsQuery() LogQueryDefinition {
	if o == nil {
		var ret LogQueryDefinition
		return ret
	}
	return o.ProfileMetricsQuery
}

// GetProfileMetricsQueryOk returns a tuple with the ProfileMetricsQuery field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455767776) GetProfileMetricsQueryOk() (*LogQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileMetricsQuery, true
}

// SetProfileMetricsQuery sets field value.
func (o *DistributionBackendRequest20270101OneOf4455767776) SetProfileMetricsQuery(v LogQueryDefinition) {
	o.ProfileMetricsQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455767776) GetStyle() RequestStyle {
	if o == nil || o.Style == nil {
		var ret RequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455767776) GetStyleOk() (*RequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455767776) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455767776) SetStyle(v RequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455767776) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["profile_metrics_query"] = o.ProfileMetricsQuery
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455767776) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ProfileMetricsQuery *LogQueryDefinition `json:"profile_metrics_query"`
		Style               *RequestStyle       `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ProfileMetricsQuery == nil {
		return fmt.Errorf("required field profile_metrics_query missing")
	}

	hasInvalidField := false
	if all.ProfileMetricsQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProfileMetricsQuery = *all.ProfileMetricsQuery
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
