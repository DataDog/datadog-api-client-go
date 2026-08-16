// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetAnomalyDetection Anomaly detection configuration for a timeseries widget.
type TimeseriesWidgetAnomalyDetection struct {
	// Sensitivity level for anomaly detection. Use `never_detect` to disable anomaly detection.
	DetectionSensitivity TimeseriesWidgetAnomalyDetectionSensitivity `json:"detection_sensitivity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetAnomalyDetection instantiates a new TimeseriesWidgetAnomalyDetection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetAnomalyDetection(detectionSensitivity TimeseriesWidgetAnomalyDetectionSensitivity) *TimeseriesWidgetAnomalyDetection {
	this := TimeseriesWidgetAnomalyDetection{}
	this.DetectionSensitivity = detectionSensitivity
	return &this
}

// NewTimeseriesWidgetAnomalyDetectionWithDefaults instantiates a new TimeseriesWidgetAnomalyDetection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetAnomalyDetectionWithDefaults() *TimeseriesWidgetAnomalyDetection {
	this := TimeseriesWidgetAnomalyDetection{}
	return &this
}

// GetDetectionSensitivity returns the DetectionSensitivity field value.
func (o *TimeseriesWidgetAnomalyDetection) GetDetectionSensitivity() TimeseriesWidgetAnomalyDetectionSensitivity {
	if o == nil {
		var ret TimeseriesWidgetAnomalyDetectionSensitivity
		return ret
	}
	return o.DetectionSensitivity
}

// GetDetectionSensitivityOk returns a tuple with the DetectionSensitivity field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetAnomalyDetection) GetDetectionSensitivityOk() (*TimeseriesWidgetAnomalyDetectionSensitivity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionSensitivity, true
}

// SetDetectionSensitivity sets field value.
func (o *TimeseriesWidgetAnomalyDetection) SetDetectionSensitivity(v TimeseriesWidgetAnomalyDetectionSensitivity) {
	o.DetectionSensitivity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetAnomalyDetection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["detection_sensitivity"] = o.DetectionSensitivity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesWidgetAnomalyDetection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DetectionSensitivity *TimeseriesWidgetAnomalyDetectionSensitivity `json:"detection_sensitivity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DetectionSensitivity == nil {
		return fmt.Errorf("required field detection_sensitivity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detection_sensitivity"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DetectionSensitivity.IsValid() {
		hasInvalidField = true
	} else {
		o.DetectionSensitivity = *all.DetectionSensitivity
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
