// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationBatch Event batching settings for Azure Data Explorer ingestion.
type ObservabilityPipelineAzureDataExplorerDestinationBatch struct {
	// Maximum batch size in bytes.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// Maximum number of events per batch before it is flushed.
	MaxEvents *int64 `json:"max_events,omitempty"`
	// Maximum number of seconds to wait before flushing a partial batch.
	TimeoutSecs *int64 `json:"timeout_secs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestinationBatch instantiates a new ObservabilityPipelineAzureDataExplorerDestinationBatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestinationBatch() *ObservabilityPipelineAzureDataExplorerDestinationBatch {
	this := ObservabilityPipelineAzureDataExplorerDestinationBatch{}
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationBatchWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestinationBatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationBatchWithDefaults() *ObservabilityPipelineAzureDataExplorerDestinationBatch {
	this := ObservabilityPipelineAzureDataExplorerDestinationBatch{}
	return &this
}

// GetMaxBytes returns the MaxBytes field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) GetMaxBytes() int64 {
	if o == nil || o.MaxBytes == nil {
		var ret int64
		return ret
	}
	return *o.MaxBytes
}

// GetMaxBytesOk returns a tuple with the MaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) GetMaxBytesOk() (*int64, bool) {
	if o == nil || o.MaxBytes == nil {
		return nil, false
	}
	return o.MaxBytes, true
}

// HasMaxBytes returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) HasMaxBytes() bool {
	return o != nil && o.MaxBytes != nil
}

// SetMaxBytes gets a reference to the given int64 and assigns it to the MaxBytes field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) SetMaxBytes(v int64) {
	o.MaxBytes = &v
}

// GetMaxEvents returns the MaxEvents field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) GetMaxEvents() int64 {
	if o == nil || o.MaxEvents == nil {
		var ret int64
		return ret
	}
	return *o.MaxEvents
}

// GetMaxEventsOk returns a tuple with the MaxEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) GetMaxEventsOk() (*int64, bool) {
	if o == nil || o.MaxEvents == nil {
		return nil, false
	}
	return o.MaxEvents, true
}

// HasMaxEvents returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) HasMaxEvents() bool {
	return o != nil && o.MaxEvents != nil
}

// SetMaxEvents gets a reference to the given int64 and assigns it to the MaxEvents field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) SetMaxEvents(v int64) {
	o.MaxEvents = &v
}

// GetTimeoutSecs returns the TimeoutSecs field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) GetTimeoutSecs() int64 {
	if o == nil || o.TimeoutSecs == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutSecs
}

// GetTimeoutSecsOk returns a tuple with the TimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) GetTimeoutSecsOk() (*int64, bool) {
	if o == nil || o.TimeoutSecs == nil {
		return nil, false
	}
	return o.TimeoutSecs, true
}

// HasTimeoutSecs returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) HasTimeoutSecs() bool {
	return o != nil && o.TimeoutSecs != nil
}

// SetTimeoutSecs gets a reference to the given int64 and assigns it to the TimeoutSecs field.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) SetTimeoutSecs(v int64) {
	o.TimeoutSecs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestinationBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MaxBytes != nil {
		toSerialize["max_bytes"] = o.MaxBytes
	}
	if o.MaxEvents != nil {
		toSerialize["max_events"] = o.MaxEvents
	}
	if o.TimeoutSecs != nil {
		toSerialize["timeout_secs"] = o.TimeoutSecs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerDestinationBatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxBytes    *int64 `json:"max_bytes,omitempty"`
		MaxEvents   *int64 `json:"max_events,omitempty"`
		TimeoutSecs *int64 `json:"timeout_secs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_bytes", "max_events", "timeout_secs"})
	} else {
		return err
	}
	o.MaxBytes = all.MaxBytes
	o.MaxEvents = all.MaxEvents
	o.TimeoutSecs = all.TimeoutSecs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
