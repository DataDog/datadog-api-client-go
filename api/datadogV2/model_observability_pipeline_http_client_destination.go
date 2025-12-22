// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineHttpClientDestination The `http_client` destination sends data to an HTTP endpoint.
//
// **Supported pipeline types:** logs, metrics
type ObservabilityPipelineHttpClientDestination struct {
	// HTTP authentication strategy.
	AuthStrategy *ObservabilityPipelineHttpClientDestinationAuthStrategy `json:"auth_strategy,omitempty"`
	// Compression configuration for HTTP requests.
	Compression *ObservabilityPipelineHttpClientDestinationCompression `json:"compression,omitempty"`
	// Encoding format for log events.
	Encoding ObservabilityPipelineHttpClientDestinationEncoding `json:"encoding"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the input for this component.
	Inputs []string `json:"inputs"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. The value should always be `http_client`.
	Type ObservabilityPipelineHttpClientDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineHttpClientDestination instantiates a new ObservabilityPipelineHttpClientDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineHttpClientDestination(encoding ObservabilityPipelineHttpClientDestinationEncoding, id string, inputs []string, typeVar ObservabilityPipelineHttpClientDestinationType) *ObservabilityPipelineHttpClientDestination {
	this := ObservabilityPipelineHttpClientDestination{}
	this.Encoding = encoding
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineHttpClientDestinationWithDefaults instantiates a new ObservabilityPipelineHttpClientDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineHttpClientDestinationWithDefaults() *ObservabilityPipelineHttpClientDestination {
	this := ObservabilityPipelineHttpClientDestination{}
	var typeVar ObservabilityPipelineHttpClientDestinationType = OBSERVABILITYPIPELINEHTTPCLIENTDESTINATIONTYPE_HTTP_CLIENT
	this.Type = typeVar
	return &this
}

// GetAuthStrategy returns the AuthStrategy field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetAuthStrategy() ObservabilityPipelineHttpClientDestinationAuthStrategy {
	if o == nil || o.AuthStrategy == nil {
		var ret ObservabilityPipelineHttpClientDestinationAuthStrategy
		return ret
	}
	return *o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetAuthStrategyOk() (*ObservabilityPipelineHttpClientDestinationAuthStrategy, bool) {
	if o == nil || o.AuthStrategy == nil {
		return nil, false
	}
	return o.AuthStrategy, true
}

// HasAuthStrategy returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasAuthStrategy() bool {
	return o != nil && o.AuthStrategy != nil
}

// SetAuthStrategy gets a reference to the given ObservabilityPipelineHttpClientDestinationAuthStrategy and assigns it to the AuthStrategy field.
func (o *ObservabilityPipelineHttpClientDestination) SetAuthStrategy(v ObservabilityPipelineHttpClientDestinationAuthStrategy) {
	o.AuthStrategy = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetCompression() ObservabilityPipelineHttpClientDestinationCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineHttpClientDestinationCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetCompressionOk() (*ObservabilityPipelineHttpClientDestinationCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineHttpClientDestinationCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineHttpClientDestination) SetCompression(v ObservabilityPipelineHttpClientDestinationCompression) {
	o.Compression = &v
}

// GetEncoding returns the Encoding field value.
func (o *ObservabilityPipelineHttpClientDestination) GetEncoding() ObservabilityPipelineHttpClientDestinationEncoding {
	if o == nil {
		var ret ObservabilityPipelineHttpClientDestinationEncoding
		return ret
	}
	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetEncodingOk() (*ObservabilityPipelineHttpClientDestinationEncoding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetEncoding(v ObservabilityPipelineHttpClientDestinationEncoding) {
	o.Encoding = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineHttpClientDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineHttpClientDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineHttpClientDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineHttpClientDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineHttpClientDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineHttpClientDestination) GetType() ObservabilityPipelineHttpClientDestinationType {
	if o == nil {
		var ret ObservabilityPipelineHttpClientDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineHttpClientDestination) GetTypeOk() (*ObservabilityPipelineHttpClientDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineHttpClientDestination) SetType(v ObservabilityPipelineHttpClientDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineHttpClientDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthStrategy != nil {
		toSerialize["auth_strategy"] = o.AuthStrategy
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	toSerialize["encoding"] = o.Encoding
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
func (o *ObservabilityPipelineHttpClientDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthStrategy *ObservabilityPipelineHttpClientDestinationAuthStrategy `json:"auth_strategy,omitempty"`
		Compression  *ObservabilityPipelineHttpClientDestinationCompression  `json:"compression,omitempty"`
		Encoding     *ObservabilityPipelineHttpClientDestinationEncoding     `json:"encoding"`
		Id           *string                                                 `json:"id"`
		Inputs       *[]string                                               `json:"inputs"`
		Tls          *ObservabilityPipelineTls                               `json:"tls,omitempty"`
		Type         *ObservabilityPipelineHttpClientDestinationType         `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Encoding == nil {
		return fmt.Errorf("required field encoding missing")
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_strategy", "compression", "encoding", "id", "inputs", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuthStrategy != nil && !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = all.AuthStrategy
	}
	if all.Compression != nil && all.Compression.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compression = all.Compression
	if !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = *all.Encoding
	}
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
