// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpentelemetryMetricsDestination The `opentelemetry` destination forwards metrics using the OpenTelemetry Protocol (OTLP) over HTTP.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelineOpentelemetryMetricsDestination struct {
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// Environment variable name containing the URI of the OTLP HTTP endpoint to send metrics to.
	HttpClientUriKey *string `json:"http_client_uri_key,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. Always `opentelemetry`.
	Type ObservabilityPipelineOpentelemetryMetricsDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOpentelemetryMetricsDestination instantiates a new ObservabilityPipelineOpentelemetryMetricsDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOpentelemetryMetricsDestination(id string, inputs []string, typeVar ObservabilityPipelineOpentelemetryMetricsDestinationType) *ObservabilityPipelineOpentelemetryMetricsDestination {
	this := ObservabilityPipelineOpentelemetryMetricsDestination{}
	var httpClientUriKey string = "DESTINATION_OTEL_HTTP_CLIENT_URI"
	this.HttpClientUriKey = &httpClientUriKey
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineOpentelemetryMetricsDestinationWithDefaults instantiates a new ObservabilityPipelineOpentelemetryMetricsDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOpentelemetryMetricsDestinationWithDefaults() *ObservabilityPipelineOpentelemetryMetricsDestination {
	this := ObservabilityPipelineOpentelemetryMetricsDestination{}
	var httpClientUriKey string = "DESTINATION_OTEL_HTTP_CLIENT_URI"
	this.HttpClientUriKey = &httpClientUriKey
	var typeVar ObservabilityPipelineOpentelemetryMetricsDestinationType = OBSERVABILITYPIPELINEOPENTELEMETRYMETRICSDESTINATIONTYPE_OPENTELEMETRY
	this.Type = typeVar
	return &this
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetHttpClientUriKey returns the HttpClientUriKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetHttpClientUriKey() string {
	if o == nil || o.HttpClientUriKey == nil {
		var ret string
		return ret
	}
	return *o.HttpClientUriKey
}

// GetHttpClientUriKeyOk returns a tuple with the HttpClientUriKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetHttpClientUriKeyOk() (*string, bool) {
	if o == nil || o.HttpClientUriKey == nil {
		return nil, false
	}
	return o.HttpClientUriKey, true
}

// HasHttpClientUriKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) HasHttpClientUriKey() bool {
	return o != nil && o.HttpClientUriKey != nil
}

// SetHttpClientUriKey gets a reference to the given string and assigns it to the HttpClientUriKey field.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) SetHttpClientUriKey(v string) {
	o.HttpClientUriKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetType() ObservabilityPipelineOpentelemetryMetricsDestinationType {
	if o == nil {
		var ret ObservabilityPipelineOpentelemetryMetricsDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) GetTypeOk() (*ObservabilityPipelineOpentelemetryMetricsDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) SetType(v ObservabilityPipelineOpentelemetryMetricsDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOpentelemetryMetricsDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.HttpClientUriKey != nil {
		toSerialize["http_client_uri_key"] = o.HttpClientUriKey
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOpentelemetryMetricsDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Buffer           *ObservabilityPipelineBufferOptions                       `json:"buffer,omitempty"`
		HttpClientUriKey *string                                                   `json:"http_client_uri_key,omitempty"`
		Id               *string                                                   `json:"id"`
		Inputs           *[]string                                                 `json:"inputs"`
		Tls              *ObservabilityPipelineTls                                 `json:"tls,omitempty"`
		Type             *ObservabilityPipelineOpentelemetryMetricsDestinationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"buffer", "http_client_uri_key", "id", "inputs", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Buffer = all.Buffer
	o.HttpClientUriKey = all.HttpClientUriKey
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
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
