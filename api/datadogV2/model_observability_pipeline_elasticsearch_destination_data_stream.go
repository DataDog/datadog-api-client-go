// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationDataStream Configuration options for writing to Elasticsearch Data Streams instead of a fixed index.
type ObservabilityPipelineElasticsearchDestinationDataStream struct {
	// The data stream dataset for your logs. This groups logs by their source or application.
	Dataset *string `json:"dataset,omitempty"`
	// The data stream type for your logs. This determines how logs are categorized within the data stream.
	Dtype *string `json:"dtype,omitempty"`
	// The data stream namespace for your logs. This separates logs into different environments or domains.
	Namespace *string `json:"namespace,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineElasticsearchDestinationDataStream instantiates a new ObservabilityPipelineElasticsearchDestinationDataStream object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineElasticsearchDestinationDataStream() *ObservabilityPipelineElasticsearchDestinationDataStream {
	this := ObservabilityPipelineElasticsearchDestinationDataStream{}
	return &this
}

// NewObservabilityPipelineElasticsearchDestinationDataStreamWithDefaults instantiates a new ObservabilityPipelineElasticsearchDestinationDataStream object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineElasticsearchDestinationDataStreamWithDefaults() *ObservabilityPipelineElasticsearchDestinationDataStream {
	this := ObservabilityPipelineElasticsearchDestinationDataStream{}
	return &this
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasDataset() bool {
	return o != nil && o.Dataset != nil
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetDataset(v string) {
	o.Dataset = &v
}

// GetDtype returns the Dtype field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDtype() string {
	if o == nil || o.Dtype == nil {
		var ret string
		return ret
	}
	return *o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDtypeOk() (*string, bool) {
	if o == nil || o.Dtype == nil {
		return nil, false
	}
	return o.Dtype, true
}

// HasDtype returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasDtype() bool {
	return o != nil && o.Dtype != nil
}

// SetDtype gets a reference to the given string and assigns it to the Dtype field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetDtype(v string) {
	o.Dtype = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetNamespace(v string) {
	o.Namespace = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineElasticsearchDestinationDataStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.Dtype != nil {
		toSerialize["dtype"] = o.Dtype
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dataset   *string `json:"dataset,omitempty"`
		Dtype     *string `json:"dtype,omitempty"`
		Namespace *string `json:"namespace,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset", "dtype", "namespace"})
	} else {
		return err
	}
	o.Dataset = all.Dataset
	o.Dtype = all.Dtype
	o.Namespace = all.Namespace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
