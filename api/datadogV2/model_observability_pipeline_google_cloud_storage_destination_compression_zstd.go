// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd Zstd compression.
type ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd struct {
	// The compression type. Always `zstd`.
	Algorithm ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType `json:"algorithm"`
	// Zstd compression level. Valid values range from 1 to 21.
	Level int64 `json:"level"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd instantiates a new ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd(algorithm ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType, level int64) *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd {
	this := ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd{}
	this.Algorithm = algorithm
	this.Level = level
	return &this
}

// NewObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdWithDefaults instantiates a new ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdWithDefaults() *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd {
	this := ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd{}
	var algorithm ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType = OBSERVABILITYPIPELINEGOOGLECLOUDSTORAGEDESTINATIONCOMPRESSIONZSTDTYPE_ZSTD
	this.Algorithm = algorithm
	return &this
}

// GetAlgorithm returns the Algorithm field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) GetAlgorithm() ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType {
	if o == nil {
		var ret ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType
		return ret
	}
	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) GetAlgorithmOk() (*ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) SetAlgorithm(v ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType) {
	o.Algorithm = v
}

// GetLevel returns the Level field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) GetLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) GetLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) SetLevel(v int64) {
	o.Level = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["level"] = o.Level

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstd) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Algorithm *ObservabilityPipelineGoogleCloudStorageDestinationCompressionZstdType `json:"algorithm"`
		Level     *int64                                                                 `json:"level"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Algorithm == nil {
		return fmt.Errorf("required field algorithm missing")
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"algorithm", "level"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Algorithm.IsValid() {
		hasInvalidField = true
	} else {
		o.Algorithm = *all.Algorithm
	}
	o.Level = *all.Level

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
