// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationResult Completed result for one timeseries request. The anomalies array is empty when no qualifying anomaly is found.
type TimeseriesAnomalyInvestigationResult struct {
	// Detected anomalies. This API version returns at most one anomaly.
	Anomalies []TimeseriesAnomalyInvestigationAnomaly `json:"anomalies"`
	// Status value indicating successful completion.
	Status TimeseriesAnomalyInvestigationCompleteStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationResult instantiates a new TimeseriesAnomalyInvestigationResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationResult(anomalies []TimeseriesAnomalyInvestigationAnomaly, status TimeseriesAnomalyInvestigationCompleteStatus) *TimeseriesAnomalyInvestigationResult {
	this := TimeseriesAnomalyInvestigationResult{}
	this.Anomalies = anomalies
	this.Status = status
	return &this
}

// NewTimeseriesAnomalyInvestigationResultWithDefaults instantiates a new TimeseriesAnomalyInvestigationResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationResultWithDefaults() *TimeseriesAnomalyInvestigationResult {
	this := TimeseriesAnomalyInvestigationResult{}
	return &this
}

// GetAnomalies returns the Anomalies field value.
func (o *TimeseriesAnomalyInvestigationResult) GetAnomalies() []TimeseriesAnomalyInvestigationAnomaly {
	if o == nil {
		var ret []TimeseriesAnomalyInvestigationAnomaly
		return ret
	}
	return o.Anomalies
}

// GetAnomaliesOk returns a tuple with the Anomalies field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResult) GetAnomaliesOk() (*[]TimeseriesAnomalyInvestigationAnomaly, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Anomalies, true
}

// SetAnomalies sets field value.
func (o *TimeseriesAnomalyInvestigationResult) SetAnomalies(v []TimeseriesAnomalyInvestigationAnomaly) {
	o.Anomalies = v
}

// GetStatus returns the Status field value.
func (o *TimeseriesAnomalyInvestigationResult) GetStatus() TimeseriesAnomalyInvestigationCompleteStatus {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationCompleteStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResult) GetStatusOk() (*TimeseriesAnomalyInvestigationCompleteStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TimeseriesAnomalyInvestigationResult) SetStatus(v TimeseriesAnomalyInvestigationCompleteStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["anomalies"] = o.Anomalies
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Anomalies *[]TimeseriesAnomalyInvestigationAnomaly      `json:"anomalies"`
		Status    *TimeseriesAnomalyInvestigationCompleteStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Anomalies == nil {
		return fmt.Errorf("required field anomalies missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anomalies", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Anomalies = *all.Anomalies
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
