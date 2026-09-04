// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationAnomaly Most significant anomaly detected in the request.
type TimeseriesAnomalyInvestigationAnomaly struct {
	// Anomaly detection configuration used for the result.
	AnomalyDetection TimeseriesAnomalyInvestigationDetection `json:"anomaly_detection"`
	// Half-open time interval in milliseconds since the Unix epoch.
	DetectedInterval TimeseriesAnomalyInvestigationInterval `json:"detected_interval"`
	// Half-open time interval in milliseconds since the Unix epoch.
	DisplayInterval TimeseriesAnomalyInvestigationInterval `json:"display_interval"`
	// Deterministic explanations for the anomaly, ordered by importance.
	Findings []TimeseriesAnomalyInvestigationFinding `json:"findings"`
	// Most anomalous point within the detected interval.
	MaximumDeviation TimeseriesAnomalyInvestigationMaximumDeviation `json:"maximum_deviation"`
	// Logical series on which the anomaly was detected.
	Series TimeseriesAnomalyInvestigationSeries `json:"series"`
	// Summary of optional influential-tag enrichment. Count and key fields are present only when analysis completes; enrichment availability does not affect completion of the investigation result.
	TagAnalysis TimeseriesAnomalyInvestigationTagAnalysis `json:"tag_analysis"`
	// Direction of an anomaly relative to its expected range.
	Type TimeseriesAnomalyInvestigationAnomalyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationAnomaly instantiates a new TimeseriesAnomalyInvestigationAnomaly object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationAnomaly(anomalyDetection TimeseriesAnomalyInvestigationDetection, detectedInterval TimeseriesAnomalyInvestigationInterval, displayInterval TimeseriesAnomalyInvestigationInterval, findings []TimeseriesAnomalyInvestigationFinding, maximumDeviation TimeseriesAnomalyInvestigationMaximumDeviation, series TimeseriesAnomalyInvestigationSeries, tagAnalysis TimeseriesAnomalyInvestigationTagAnalysis, typeVar TimeseriesAnomalyInvestigationAnomalyType) *TimeseriesAnomalyInvestigationAnomaly {
	this := TimeseriesAnomalyInvestigationAnomaly{}
	this.AnomalyDetection = anomalyDetection
	this.DetectedInterval = detectedInterval
	this.DisplayInterval = displayInterval
	this.Findings = findings
	this.MaximumDeviation = maximumDeviation
	this.Series = series
	this.TagAnalysis = tagAnalysis
	this.Type = typeVar
	return &this
}

// NewTimeseriesAnomalyInvestigationAnomalyWithDefaults instantiates a new TimeseriesAnomalyInvestigationAnomaly object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationAnomalyWithDefaults() *TimeseriesAnomalyInvestigationAnomaly {
	this := TimeseriesAnomalyInvestigationAnomaly{}
	return &this
}

// GetAnomalyDetection returns the AnomalyDetection field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetAnomalyDetection() TimeseriesAnomalyInvestigationDetection {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationDetection
		return ret
	}
	return o.AnomalyDetection
}

// GetAnomalyDetectionOk returns a tuple with the AnomalyDetection field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetAnomalyDetectionOk() (*TimeseriesAnomalyInvestigationDetection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnomalyDetection, true
}

// SetAnomalyDetection sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetAnomalyDetection(v TimeseriesAnomalyInvestigationDetection) {
	o.AnomalyDetection = v
}

// GetDetectedInterval returns the DetectedInterval field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetDetectedInterval() TimeseriesAnomalyInvestigationInterval {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationInterval
		return ret
	}
	return o.DetectedInterval
}

// GetDetectedIntervalOk returns a tuple with the DetectedInterval field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetDetectedIntervalOk() (*TimeseriesAnomalyInvestigationInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectedInterval, true
}

// SetDetectedInterval sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetDetectedInterval(v TimeseriesAnomalyInvestigationInterval) {
	o.DetectedInterval = v
}

// GetDisplayInterval returns the DisplayInterval field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetDisplayInterval() TimeseriesAnomalyInvestigationInterval {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationInterval
		return ret
	}
	return o.DisplayInterval
}

// GetDisplayIntervalOk returns a tuple with the DisplayInterval field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetDisplayIntervalOk() (*TimeseriesAnomalyInvestigationInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayInterval, true
}

// SetDisplayInterval sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetDisplayInterval(v TimeseriesAnomalyInvestigationInterval) {
	o.DisplayInterval = v
}

// GetFindings returns the Findings field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetFindings() []TimeseriesAnomalyInvestigationFinding {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationFinding
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetFindingsOk() (*[]TimeseriesAnomalyInvestigationFinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetFindings(v []TimeseriesAnomalyInvestigationFinding) {
	o.Findings = v
}

// GetMaximumDeviation returns the MaximumDeviation field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetMaximumDeviation() TimeseriesAnomalyInvestigationMaximumDeviation {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationMaximumDeviation
		return ret
	}
	return o.MaximumDeviation
}

// GetMaximumDeviationOk returns a tuple with the MaximumDeviation field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetMaximumDeviationOk() (*TimeseriesAnomalyInvestigationMaximumDeviation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumDeviation, true
}

// SetMaximumDeviation sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetMaximumDeviation(v TimeseriesAnomalyInvestigationMaximumDeviation) {
	o.MaximumDeviation = v
}

// GetSeries returns the Series field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetSeries() TimeseriesAnomalyInvestigationSeries {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationSeries
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetSeriesOk() (*TimeseriesAnomalyInvestigationSeries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Series, true
}

// SetSeries sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetSeries(v TimeseriesAnomalyInvestigationSeries) {
	o.Series = v
}

// GetTagAnalysis returns the TagAnalysis field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetTagAnalysis() TimeseriesAnomalyInvestigationTagAnalysis {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationTagAnalysis
		return ret
	}
	return o.TagAnalysis
}

// GetTagAnalysisOk returns a tuple with the TagAnalysis field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetTagAnalysisOk() (*TimeseriesAnomalyInvestigationTagAnalysis, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagAnalysis, true
}

// SetTagAnalysis sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetTagAnalysis(v TimeseriesAnomalyInvestigationTagAnalysis) {
	o.TagAnalysis = v
}

// GetType returns the Type field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetType() TimeseriesAnomalyInvestigationAnomalyType {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationAnomalyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationAnomaly) GetTypeOk() (*TimeseriesAnomalyInvestigationAnomalyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TimeseriesAnomalyInvestigationAnomaly) SetType(v TimeseriesAnomalyInvestigationAnomalyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationAnomaly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["anomaly_detection"] = o.AnomalyDetection
	toSerialize["detected_interval"] = o.DetectedInterval
	toSerialize["display_interval"] = o.DisplayInterval
	toSerialize["findings"] = o.Findings
	toSerialize["maximum_deviation"] = o.MaximumDeviation
	toSerialize["series"] = o.Series
	toSerialize["tag_analysis"] = o.TagAnalysis
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationAnomaly) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnomalyDetection *TimeseriesAnomalyInvestigationDetection        `json:"anomaly_detection"`
		DetectedInterval *TimeseriesAnomalyInvestigationInterval         `json:"detected_interval"`
		DisplayInterval  *TimeseriesAnomalyInvestigationInterval         `json:"display_interval"`
		Findings         *[]TimeseriesAnomalyInvestigationFinding        `json:"findings"`
		MaximumDeviation *TimeseriesAnomalyInvestigationMaximumDeviation `json:"maximum_deviation"`
		Series           *TimeseriesAnomalyInvestigationSeries           `json:"series"`
		TagAnalysis      *TimeseriesAnomalyInvestigationTagAnalysis      `json:"tag_analysis"`
		Type             *TimeseriesAnomalyInvestigationAnomalyType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnomalyDetection == nil {
		return fmt.Errorf("required field anomaly_detection missing")
	}
	if all.DetectedInterval == nil {
		return fmt.Errorf("required field detected_interval missing")
	}
	if all.DisplayInterval == nil {
		return fmt.Errorf("required field display_interval missing")
	}
	if all.Findings == nil {
		return fmt.Errorf("required field findings missing")
	}
	if all.MaximumDeviation == nil {
		return fmt.Errorf("required field maximum_deviation missing")
	}
	if all.Series == nil {
		return fmt.Errorf("required field series missing")
	}
	if all.TagAnalysis == nil {
		return fmt.Errorf("required field tag_analysis missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anomaly_detection", "detected_interval", "display_interval", "findings", "maximum_deviation", "series", "tag_analysis", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AnomalyDetection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AnomalyDetection = *all.AnomalyDetection
	if all.DetectedInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DetectedInterval = *all.DetectedInterval
	if all.DisplayInterval.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayInterval = *all.DisplayInterval
	o.Findings = *all.Findings
	if all.MaximumDeviation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MaximumDeviation = *all.MaximumDeviation
	if all.Series.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Series = *all.Series
	if all.TagAnalysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TagAnalysis = *all.TagAnalysis
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
