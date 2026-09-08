// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationTagAnalysis Summary of optional influential-tag enrichment. Count and key fields are present only when analysis completes; enrichment availability does not affect completion of the investigation result.
type TimeseriesAnomalyInvestigationTagAnalysis struct {
	// Tag keys analyzed. Present only when analysis completes.
	AnalyzedTagKeys []string `json:"analyzed_tag_keys,omitempty"`
	// Outcome of optional influential-tag enrichment.
	Status TimeseriesAnomalyInvestigationTagAnalysisStatus `json:"status"`
	// Number of tag keys analyzed. Present only when analysis completes.
	TagKeysAnalyzed *int64 `json:"tag_keys_analyzed,omitempty"`
	// Number of tag values analyzed. Present only when analysis completes.
	TagValuesAnalyzed *int64 `json:"tag_values_analyzed,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationTagAnalysis instantiates a new TimeseriesAnomalyInvestigationTagAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationTagAnalysis(status TimeseriesAnomalyInvestigationTagAnalysisStatus) *TimeseriesAnomalyInvestigationTagAnalysis {
	this := TimeseriesAnomalyInvestigationTagAnalysis{}
	this.Status = status
	return &this
}

// NewTimeseriesAnomalyInvestigationTagAnalysisWithDefaults instantiates a new TimeseriesAnomalyInvestigationTagAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationTagAnalysisWithDefaults() *TimeseriesAnomalyInvestigationTagAnalysis {
	this := TimeseriesAnomalyInvestigationTagAnalysis{}
	return &this
}

// GetAnalyzedTagKeys returns the AnalyzedTagKeys field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetAnalyzedTagKeys() []string {
	if o == nil || o.AnalyzedTagKeys == nil {
		var ret []string
		return ret
	}
	return o.AnalyzedTagKeys
}

// GetAnalyzedTagKeysOk returns a tuple with the AnalyzedTagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetAnalyzedTagKeysOk() (*[]string, bool) {
	if o == nil || o.AnalyzedTagKeys == nil {
		return nil, false
	}
	return &o.AnalyzedTagKeys, true
}

// HasAnalyzedTagKeys returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) HasAnalyzedTagKeys() bool {
	return o != nil && o.AnalyzedTagKeys != nil
}

// SetAnalyzedTagKeys gets a reference to the given []string and assigns it to the AnalyzedTagKeys field.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) SetAnalyzedTagKeys(v []string) {
	o.AnalyzedTagKeys = v
}

// GetStatus returns the Status field value.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetStatus() TimeseriesAnomalyInvestigationTagAnalysisStatus {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationTagAnalysisStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetStatusOk() (*TimeseriesAnomalyInvestigationTagAnalysisStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) SetStatus(v TimeseriesAnomalyInvestigationTagAnalysisStatus) {
	o.Status = v
}

// GetTagKeysAnalyzed returns the TagKeysAnalyzed field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetTagKeysAnalyzed() int64 {
	if o == nil || o.TagKeysAnalyzed == nil {
		var ret int64
		return ret
	}
	return *o.TagKeysAnalyzed
}

// GetTagKeysAnalyzedOk returns a tuple with the TagKeysAnalyzed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetTagKeysAnalyzedOk() (*int64, bool) {
	if o == nil || o.TagKeysAnalyzed == nil {
		return nil, false
	}
	return o.TagKeysAnalyzed, true
}

// HasTagKeysAnalyzed returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) HasTagKeysAnalyzed() bool {
	return o != nil && o.TagKeysAnalyzed != nil
}

// SetTagKeysAnalyzed gets a reference to the given int64 and assigns it to the TagKeysAnalyzed field.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) SetTagKeysAnalyzed(v int64) {
	o.TagKeysAnalyzed = &v
}

// GetTagValuesAnalyzed returns the TagValuesAnalyzed field value if set, zero value otherwise.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetTagValuesAnalyzed() int64 {
	if o == nil || o.TagValuesAnalyzed == nil {
		var ret int64
		return ret
	}
	return *o.TagValuesAnalyzed
}

// GetTagValuesAnalyzedOk returns a tuple with the TagValuesAnalyzed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) GetTagValuesAnalyzedOk() (*int64, bool) {
	if o == nil || o.TagValuesAnalyzed == nil {
		return nil, false
	}
	return o.TagValuesAnalyzed, true
}

// HasTagValuesAnalyzed returns a boolean if a field has been set.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) HasTagValuesAnalyzed() bool {
	return o != nil && o.TagValuesAnalyzed != nil
}

// SetTagValuesAnalyzed gets a reference to the given int64 and assigns it to the TagValuesAnalyzed field.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) SetTagValuesAnalyzed(v int64) {
	o.TagValuesAnalyzed = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationTagAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnalyzedTagKeys != nil {
		toSerialize["analyzed_tag_keys"] = o.AnalyzedTagKeys
	}
	toSerialize["status"] = o.Status
	if o.TagKeysAnalyzed != nil {
		toSerialize["tag_keys_analyzed"] = o.TagKeysAnalyzed
	}
	if o.TagValuesAnalyzed != nil {
		toSerialize["tag_values_analyzed"] = o.TagValuesAnalyzed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationTagAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalyzedTagKeys   []string                                         `json:"analyzed_tag_keys,omitempty"`
		Status            *TimeseriesAnomalyInvestigationTagAnalysisStatus `json:"status"`
		TagKeysAnalyzed   *int64                                           `json:"tag_keys_analyzed,omitempty"`
		TagValuesAnalyzed *int64                                           `json:"tag_values_analyzed,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analyzed_tag_keys", "status", "tag_keys_analyzed", "tag_values_analyzed"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AnalyzedTagKeys = all.AnalyzedTagKeys
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.TagKeysAnalyzed = all.TagKeysAnalyzed
	o.TagValuesAnalyzed = all.TagValuesAnalyzed

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
