// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompressionZstd Zstd compression.
type ObservabilityPipelineAmazonS3GenericCompressionZstd struct {
	// Zstd compression level.
	Level int64 `json:"level"`
	// The compression type. Always `zstd`.
	Type ObservabilityPipelineAmazonS3GenericCompressionZstdType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonS3GenericCompressionZstd instantiates a new ObservabilityPipelineAmazonS3GenericCompressionZstd object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonS3GenericCompressionZstd(level int64, typeVar ObservabilityPipelineAmazonS3GenericCompressionZstdType) *ObservabilityPipelineAmazonS3GenericCompressionZstd {
	this := ObservabilityPipelineAmazonS3GenericCompressionZstd{}
	this.Level = level
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAmazonS3GenericCompressionZstdWithDefaults instantiates a new ObservabilityPipelineAmazonS3GenericCompressionZstd object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonS3GenericCompressionZstdWithDefaults() *ObservabilityPipelineAmazonS3GenericCompressionZstd {
	this := ObservabilityPipelineAmazonS3GenericCompressionZstd{}
	var typeVar ObservabilityPipelineAmazonS3GenericCompressionZstdType = OBSERVABILITYPIPELINEAMAZONS3GENERICCOMPRESSIONZSTDTYPE_ZSTD
	this.Type = typeVar
	return &this
}

// GetLevel returns the Level field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) GetLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) GetLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) SetLevel(v int64) {
	o.Level = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) GetType() ObservabilityPipelineAmazonS3GenericCompressionZstdType {
	if o == nil {
		var ret ObservabilityPipelineAmazonS3GenericCompressionZstdType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) GetTypeOk() (*ObservabilityPipelineAmazonS3GenericCompressionZstdType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) SetType(v ObservabilityPipelineAmazonS3GenericCompressionZstdType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonS3GenericCompressionZstd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["level"] = o.Level
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonS3GenericCompressionZstd) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Level *int64                                                   `json:"level"`
		Type  *ObservabilityPipelineAmazonS3GenericCompressionZstdType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"level", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Level = *all.Level
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
