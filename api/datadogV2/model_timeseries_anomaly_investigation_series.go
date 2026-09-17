// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationSeries Logical series on which the anomaly was detected.
type TimeseriesAnomalyInvestigationSeries struct {
	// Tags identifying the selected group. Empty for a query without grouping.
	GroupTags []string `json:"group_tags"`
	// Display label for the selected series.
	Label string `json:"label"`
	// Zero-based index of the caller's formula that produced the series.
	QueryIndex int64 `json:"query_index"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationSeries instantiates a new TimeseriesAnomalyInvestigationSeries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationSeries(groupTags []string, label string, queryIndex int64) *TimeseriesAnomalyInvestigationSeries {
	this := TimeseriesAnomalyInvestigationSeries{}
	this.GroupTags = groupTags
	this.Label = label
	this.QueryIndex = queryIndex
	return &this
}

// NewTimeseriesAnomalyInvestigationSeriesWithDefaults instantiates a new TimeseriesAnomalyInvestigationSeries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationSeriesWithDefaults() *TimeseriesAnomalyInvestigationSeries {
	this := TimeseriesAnomalyInvestigationSeries{}
	return &this
}

// GetGroupTags returns the GroupTags field value.
func (o *TimeseriesAnomalyInvestigationSeries) GetGroupTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupTags
}

// GetGroupTagsOk returns a tuple with the GroupTags field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationSeries) GetGroupTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupTags, true
}

// SetGroupTags sets field value.
func (o *TimeseriesAnomalyInvestigationSeries) SetGroupTags(v []string) {
	o.GroupTags = v
}

// GetLabel returns the Label field value.
func (o *TimeseriesAnomalyInvestigationSeries) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationSeries) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *TimeseriesAnomalyInvestigationSeries) SetLabel(v string) {
	o.Label = v
}

// GetQueryIndex returns the QueryIndex field value.
func (o *TimeseriesAnomalyInvestigationSeries) GetQueryIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.QueryIndex
}

// GetQueryIndexOk returns a tuple with the QueryIndex field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationSeries) GetQueryIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryIndex, true
}

// SetQueryIndex sets field value.
func (o *TimeseriesAnomalyInvestigationSeries) SetQueryIndex(v int64) {
	o.QueryIndex = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["group_tags"] = o.GroupTags
	toSerialize["label"] = o.Label
	toSerialize["query_index"] = o.QueryIndex

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationSeries) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupTags  *[]string `json:"group_tags"`
		Label      *string   `json:"label"`
		QueryIndex *int64    `json:"query_index"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GroupTags == nil {
		return fmt.Errorf("required field group_tags missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.QueryIndex == nil {
		return fmt.Errorf("required field query_index missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_tags", "label", "query_index"})
	} else {
		return err
	}
	o.GroupTags = *all.GroupTags
	o.Label = *all.Label
	o.QueryIndex = *all.QueryIndex

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
