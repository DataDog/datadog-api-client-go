// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactViews RUM views where the issue occurred.
type IssueImpactViews struct {
	// Number of distinct view names where the issue occurred in the time window.
	DistinctCount int64 `json:"distinct_count"`
	// Up to five views with the most errors, most affected first.
	Top []IssueImpactView `json:"top"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactViews instantiates a new IssueImpactViews object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactViews(distinctCount int64, top []IssueImpactView) *IssueImpactViews {
	this := IssueImpactViews{}
	this.DistinctCount = distinctCount
	this.Top = top
	return &this
}

// NewIssueImpactViewsWithDefaults instantiates a new IssueImpactViews object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactViewsWithDefaults() *IssueImpactViews {
	this := IssueImpactViews{}
	return &this
}

// GetDistinctCount returns the DistinctCount field value.
func (o *IssueImpactViews) GetDistinctCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DistinctCount
}

// GetDistinctCountOk returns a tuple with the DistinctCount field value
// and a boolean to check if the value has been set.
func (o *IssueImpactViews) GetDistinctCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistinctCount, true
}

// SetDistinctCount sets field value.
func (o *IssueImpactViews) SetDistinctCount(v int64) {
	o.DistinctCount = v
}

// GetTop returns the Top field value.
func (o *IssueImpactViews) GetTop() []IssueImpactView {
	if o == nil {
		var ret []IssueImpactView
		return ret
	}
	return o.Top
}

// GetTopOk returns a tuple with the Top field value
// and a boolean to check if the value has been set.
func (o *IssueImpactViews) GetTopOk() (*[]IssueImpactView, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Top, true
}

// SetTop sets field value.
func (o *IssueImpactViews) SetTop(v []IssueImpactView) {
	o.Top = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactViews) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["distinct_count"] = o.DistinctCount
	toSerialize["top"] = o.Top

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactViews) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DistinctCount *int64             `json:"distinct_count"`
		Top           *[]IssueImpactView `json:"top"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DistinctCount == nil {
		return fmt.Errorf("required field distinct_count missing")
	}
	if all.Top == nil {
		return fmt.Errorf("required field top missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"distinct_count", "top"})
	} else {
		return err
	}
	o.DistinctCount = *all.DistinctCount
	o.Top = *all.Top

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
