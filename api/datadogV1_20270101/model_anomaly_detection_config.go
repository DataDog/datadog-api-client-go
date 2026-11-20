// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnomalyDetectionConfig Test-only generated schema for the dashboard versioning proof of concept.
type AnomalyDetectionConfig struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	DetectionSensitivity AnomalyDetectionConfigDetectionSensitivity `json:"detection_sensitivity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAnomalyDetectionConfig instantiates a new AnomalyDetectionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnomalyDetectionConfig(detectionSensitivity AnomalyDetectionConfigDetectionSensitivity) *AnomalyDetectionConfig {
	this := AnomalyDetectionConfig{}
	this.DetectionSensitivity = detectionSensitivity
	return &this
}

// NewAnomalyDetectionConfigWithDefaults instantiates a new AnomalyDetectionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnomalyDetectionConfigWithDefaults() *AnomalyDetectionConfig {
	this := AnomalyDetectionConfig{}
	return &this
}

// GetDetectionSensitivity returns the DetectionSensitivity field value.
func (o *AnomalyDetectionConfig) GetDetectionSensitivity() AnomalyDetectionConfigDetectionSensitivity {
	if o == nil {
		var ret AnomalyDetectionConfigDetectionSensitivity
		return ret
	}
	return o.DetectionSensitivity
}

// GetDetectionSensitivityOk returns a tuple with the DetectionSensitivity field value
// and a boolean to check if the value has been set.
func (o *AnomalyDetectionConfig) GetDetectionSensitivityOk() (*AnomalyDetectionConfigDetectionSensitivity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionSensitivity, true
}

// SetDetectionSensitivity sets field value.
func (o *AnomalyDetectionConfig) SetDetectionSensitivity(v AnomalyDetectionConfigDetectionSensitivity) {
	o.DetectionSensitivity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["detection_sensitivity"] = o.DetectionSensitivity
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnomalyDetectionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DetectionSensitivity *AnomalyDetectionConfigDetectionSensitivity `json:"detection_sensitivity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DetectionSensitivity == nil {
		return fmt.Errorf("required field detection_sensitivity missing")
	}

	hasInvalidField := false
	if !all.DetectionSensitivity.IsValid() {
		hasInvalidField = true
	} else {
		o.DetectionSensitivity = *all.DetectionSensitivity
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
