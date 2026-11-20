// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionBackendRequest20270101OneOf4455772256 Test-only generated schema for the dashboard versioning proof of concept.
type DistributionBackendRequest20270101OneOf4455772256 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	ApmStatsQuery ApmStatsQuery `json:"apm_stats_query"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Style *RequestStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDistributionBackendRequest20270101OneOf4455772256 instantiates a new DistributionBackendRequest20270101OneOf4455772256 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionBackendRequest20270101OneOf4455772256(apmStatsQuery ApmStatsQuery) *DistributionBackendRequest20270101OneOf4455772256 {
	this := DistributionBackendRequest20270101OneOf4455772256{}
	this.ApmStatsQuery = apmStatsQuery
	return &this
}

// NewDistributionBackendRequest20270101OneOf4455772256WithDefaults instantiates a new DistributionBackendRequest20270101OneOf4455772256 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionBackendRequest20270101OneOf4455772256WithDefaults() *DistributionBackendRequest20270101OneOf4455772256 {
	this := DistributionBackendRequest20270101OneOf4455772256{}
	return &this
}

// GetApmStatsQuery returns the ApmStatsQuery field value.
func (o *DistributionBackendRequest20270101OneOf4455772256) GetApmStatsQuery() ApmStatsQuery {
	if o == nil {
		var ret ApmStatsQuery
		return ret
	}
	return o.ApmStatsQuery
}

// GetApmStatsQueryOk returns a tuple with the ApmStatsQuery field value
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455772256) GetApmStatsQueryOk() (*ApmStatsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApmStatsQuery, true
}

// SetApmStatsQuery sets field value.
func (o *DistributionBackendRequest20270101OneOf4455772256) SetApmStatsQuery(v ApmStatsQuery) {
	o.ApmStatsQuery = v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *DistributionBackendRequest20270101OneOf4455772256) GetStyle() RequestStyle {
	if o == nil || o.Style == nil {
		var ret RequestStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionBackendRequest20270101OneOf4455772256) GetStyleOk() (*RequestStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *DistributionBackendRequest20270101OneOf4455772256) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given RequestStyle and assigns it to the Style field.
func (o *DistributionBackendRequest20270101OneOf4455772256) SetStyle(v RequestStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionBackendRequest20270101OneOf4455772256) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["apm_stats_query"] = o.ApmStatsQuery
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionBackendRequest20270101OneOf4455772256) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApmStatsQuery *ApmStatsQuery `json:"apm_stats_query"`
		Style         *RequestStyle  `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApmStatsQuery == nil {
		return fmt.Errorf("required field apm_stats_query missing")
	}

	hasInvalidField := false
	if all.ApmStatsQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApmStatsQuery = *all.ApmStatsQuery
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
