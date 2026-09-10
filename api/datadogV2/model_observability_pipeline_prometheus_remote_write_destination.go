// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePrometheusRemoteWriteDestination The `prometheus_remote_write` destination forwards metrics to an endpoint that supports the Prometheus Remote Write protocol.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelinePrometheusRemoteWriteDestination struct {
	// The authentication strategy to use for outgoing Prometheus Remote Write requests.
	AuthStrategy *ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy `json:"auth_strategy,omitempty"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// The default namespace to add as a prefix to metric names that do not already have one.
	DefaultNamespace *string `json:"default_namespace,omitempty"`
	// Name of the environment variable or secret that holds the Prometheus Remote Write endpoint URL.
	// Defaults to `DESTINATION_PROMETHEUS_REMOTE_WRITE_ENDPOINT_URL` (prefixed with `DD_OP_` at runtime).
	EndpointUrlKey *string `json:"endpoint_url_key,omitempty"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// Name of the environment variable or secret that holds the password (used when `auth_strategy` is `basic`).
	PasswordKey *string `json:"password_key,omitempty"`
	// The tenant ID to include with outgoing requests. Used by multi-tenant Prometheus Remote Write receivers.
	TenantId *string `json:"tenant_id,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineClientTls `json:"tls,omitempty"`
	// Name of the environment variable or secret that holds the bearer token (used when `auth_strategy` is `bearer`).
	TokenKey *string `json:"token_key,omitempty"`
	// The destination type. The value should always be `prometheus_remote_write`.
	Type ObservabilityPipelinePrometheusRemoteWriteDestinationType `json:"type"`
	// Name of the environment variable or secret that holds the username (used when `auth_strategy` is `basic`).
	UsernameKey *string `json:"username_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelinePrometheusRemoteWriteDestination instantiates a new ObservabilityPipelinePrometheusRemoteWriteDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelinePrometheusRemoteWriteDestination(id string, inputs []string, typeVar ObservabilityPipelinePrometheusRemoteWriteDestinationType) *ObservabilityPipelinePrometheusRemoteWriteDestination {
	this := ObservabilityPipelinePrometheusRemoteWriteDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelinePrometheusRemoteWriteDestinationWithDefaults instantiates a new ObservabilityPipelinePrometheusRemoteWriteDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelinePrometheusRemoteWriteDestinationWithDefaults() *ObservabilityPipelinePrometheusRemoteWriteDestination {
	this := ObservabilityPipelinePrometheusRemoteWriteDestination{}
	var typeVar ObservabilityPipelinePrometheusRemoteWriteDestinationType = OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITEDESTINATIONTYPE_PROMETHEUS_REMOTE_WRITE
	this.Type = typeVar
	return &this
}

// GetAuthStrategy returns the AuthStrategy field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetAuthStrategy() ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy {
	if o == nil || o.AuthStrategy == nil {
		var ret ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy
		return ret
	}
	return *o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetAuthStrategyOk() (*ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy, bool) {
	if o == nil || o.AuthStrategy == nil {
		return nil, false
	}
	return o.AuthStrategy, true
}

// HasAuthStrategy returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasAuthStrategy() bool {
	return o != nil && o.AuthStrategy != nil
}

// SetAuthStrategy gets a reference to the given ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy and assigns it to the AuthStrategy field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetAuthStrategy(v ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy) {
	o.AuthStrategy = &v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetDefaultNamespace returns the DefaultNamespace field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetDefaultNamespace() string {
	if o == nil || o.DefaultNamespace == nil {
		var ret string
		return ret
	}
	return *o.DefaultNamespace
}

// GetDefaultNamespaceOk returns a tuple with the DefaultNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetDefaultNamespaceOk() (*string, bool) {
	if o == nil || o.DefaultNamespace == nil {
		return nil, false
	}
	return o.DefaultNamespace, true
}

// HasDefaultNamespace returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasDefaultNamespace() bool {
	return o != nil && o.DefaultNamespace != nil
}

// SetDefaultNamespace gets a reference to the given string and assigns it to the DefaultNamespace field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetDefaultNamespace(v string) {
	o.DefaultNamespace = &v
}

// GetEndpointUrlKey returns the EndpointUrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetEndpointUrlKey() string {
	if o == nil || o.EndpointUrlKey == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrlKey
}

// GetEndpointUrlKeyOk returns a tuple with the EndpointUrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetEndpointUrlKeyOk() (*string, bool) {
	if o == nil || o.EndpointUrlKey == nil {
		return nil, false
	}
	return o.EndpointUrlKey, true
}

// HasEndpointUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasEndpointUrlKey() bool {
	return o != nil && o.EndpointUrlKey != nil
}

// SetEndpointUrlKey gets a reference to the given string and assigns it to the EndpointUrlKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetEndpointUrlKey(v string) {
	o.EndpointUrlKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTls() ObservabilityPipelineClientTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineClientTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTlsOk() (*ObservabilityPipelineClientTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineClientTls and assigns it to the Tls field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetTls(v ObservabilityPipelineClientTls) {
	o.Tls = &v
}

// GetTokenKey returns the TokenKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTokenKey() string {
	if o == nil || o.TokenKey == nil {
		var ret string
		return ret
	}
	return *o.TokenKey
}

// GetTokenKeyOk returns a tuple with the TokenKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTokenKeyOk() (*string, bool) {
	if o == nil || o.TokenKey == nil {
		return nil, false
	}
	return o.TokenKey, true
}

// HasTokenKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasTokenKey() bool {
	return o != nil && o.TokenKey != nil
}

// SetTokenKey gets a reference to the given string and assigns it to the TokenKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetTokenKey(v string) {
	o.TokenKey = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetType() ObservabilityPipelinePrometheusRemoteWriteDestinationType {
	if o == nil {
		var ret ObservabilityPipelinePrometheusRemoteWriteDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetTypeOk() (*ObservabilityPipelinePrometheusRemoteWriteDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetType(v ObservabilityPipelinePrometheusRemoteWriteDestinationType) {
	o.Type = v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelinePrometheusRemoteWriteDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthStrategy != nil {
		toSerialize["auth_strategy"] = o.AuthStrategy
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.DefaultNamespace != nil {
		toSerialize["default_namespace"] = o.DefaultNamespace
	}
	if o.EndpointUrlKey != nil {
		toSerialize["endpoint_url_key"] = o.EndpointUrlKey
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.PasswordKey != nil {
		toSerialize["password_key"] = o.PasswordKey
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	if o.TokenKey != nil {
		toSerialize["token_key"] = o.TokenKey
	}
	toSerialize["type"] = o.Type
	if o.UsernameKey != nil {
		toSerialize["username_key"] = o.UsernameKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelinePrometheusRemoteWriteDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthStrategy     *ObservabilityPipelinePrometheusRemoteWriteDestinationAuthStrategy `json:"auth_strategy,omitempty"`
		Buffer           *ObservabilityPipelineBufferOptions                                `json:"buffer,omitempty"`
		DefaultNamespace *string                                                            `json:"default_namespace,omitempty"`
		EndpointUrlKey   *string                                                            `json:"endpoint_url_key,omitempty"`
		Id               *string                                                            `json:"id"`
		Inputs           *[]string                                                          `json:"inputs"`
		PasswordKey      *string                                                            `json:"password_key,omitempty"`
		TenantId         *string                                                            `json:"tenant_id,omitempty"`
		Tls              *ObservabilityPipelineClientTls                                    `json:"tls,omitempty"`
		TokenKey         *string                                                            `json:"token_key,omitempty"`
		Type             *ObservabilityPipelinePrometheusRemoteWriteDestinationType         `json:"type"`
		UsernameKey      *string                                                            `json:"username_key,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"auth_strategy", "buffer", "default_namespace", "endpoint_url_key", "id", "inputs", "password_key", "tenant_id", "tls", "token_key", "type", "username_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AuthStrategy != nil && !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = all.AuthStrategy
	}
	o.Buffer = all.Buffer
	o.DefaultNamespace = all.DefaultNamespace
	o.EndpointUrlKey = all.EndpointUrlKey
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.PasswordKey = all.PasswordKey
	o.TenantId = all.TenantId
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	o.TokenKey = all.TokenKey
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UsernameKey = all.UsernameKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
