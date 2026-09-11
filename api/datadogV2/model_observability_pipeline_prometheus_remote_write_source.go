// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelinePrometheusRemoteWriteSource The `prometheus_remote_write` source ingests metrics pushed over the Prometheus Remote Write protocol.
//
// **Supported pipeline types:** metrics
type ObservabilityPipelinePrometheusRemoteWriteSource struct {
	// Name of the environment variable or secret that holds the listen address for the Prometheus Remote Write endpoint.
	AddressKey *string `json:"address_key,omitempty"`
	// HTTP authentication method.
	AuthStrategy ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy `json:"auth_strategy"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// Name of the environment variable or secret that holds the password (used when `auth_strategy` is `plain`).
	PasswordKey *string `json:"password_key,omitempty"`
	// The HTTP path on which the source listens for incoming Prometheus Remote Write requests.
	Path *string `json:"path,omitempty"`
	// Configuration for enabling TLS encryption between the pipeline component and external connecting clients.
	Tls *ObservabilityPipelineMtlsServerTls `json:"tls,omitempty"`
	// The source type. The value should always be `prometheus_remote_write`.
	Type ObservabilityPipelinePrometheusRemoteWriteSourceType `json:"type"`
	// Name of the environment variable or secret that holds the username (used when `auth_strategy` is `plain`).
	UsernameKey *string `json:"username_key,omitempty"`
	// A list of tokens that are accepted for authenticating incoming requests. When set,
	// the source rejects any request whose token does not match an enabled entry in this list.
	ValidTokens []ObservabilityPipelinePrometheusRemoteWriteSourceValidToken `json:"valid_tokens,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelinePrometheusRemoteWriteSource instantiates a new ObservabilityPipelinePrometheusRemoteWriteSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelinePrometheusRemoteWriteSource(authStrategy ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy, id string, typeVar ObservabilityPipelinePrometheusRemoteWriteSourceType) *ObservabilityPipelinePrometheusRemoteWriteSource {
	this := ObservabilityPipelinePrometheusRemoteWriteSource{}
	this.AuthStrategy = authStrategy
	this.Id = id
	var path string = "/api/v1/write"
	this.Path = &path
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelinePrometheusRemoteWriteSourceWithDefaults instantiates a new ObservabilityPipelinePrometheusRemoteWriteSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelinePrometheusRemoteWriteSourceWithDefaults() *ObservabilityPipelinePrometheusRemoteWriteSource {
	this := ObservabilityPipelinePrometheusRemoteWriteSource{}
	var path string = "/api/v1/write"
	this.Path = &path
	var typeVar ObservabilityPipelinePrometheusRemoteWriteSourceType = OBSERVABILITYPIPELINEPROMETHEUSREMOTEWRITESOURCETYPE_PROMETHEUS_REMOTE_WRITE
	this.Type = typeVar
	return &this
}

// GetAddressKey returns the AddressKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetAddressKey() string {
	if o == nil || o.AddressKey == nil {
		var ret string
		return ret
	}
	return *o.AddressKey
}

// GetAddressKeyOk returns a tuple with the AddressKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetAddressKeyOk() (*string, bool) {
	if o == nil || o.AddressKey == nil {
		return nil, false
	}
	return o.AddressKey, true
}

// HasAddressKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) HasAddressKey() bool {
	return o != nil && o.AddressKey != nil
}

// SetAddressKey gets a reference to the given string and assigns it to the AddressKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetAddressKey(v string) {
	o.AddressKey = &v
}

// GetAuthStrategy returns the AuthStrategy field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetAuthStrategy() ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy {
	if o == nil {
		var ret ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy
		return ret
	}
	return o.AuthStrategy
}

// GetAuthStrategyOk returns a tuple with the AuthStrategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetAuthStrategyOk() (*ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthStrategy, true
}

// SetAuthStrategy sets field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetAuthStrategy(v ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy) {
	o.AuthStrategy = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetId(v string) {
	o.Id = v
}

// GetPasswordKey returns the PasswordKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetPasswordKey() string {
	if o == nil || o.PasswordKey == nil {
		var ret string
		return ret
	}
	return *o.PasswordKey
}

// GetPasswordKeyOk returns a tuple with the PasswordKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetPasswordKeyOk() (*string, bool) {
	if o == nil || o.PasswordKey == nil {
		return nil, false
	}
	return o.PasswordKey, true
}

// HasPasswordKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) HasPasswordKey() bool {
	return o != nil && o.PasswordKey != nil
}

// SetPasswordKey gets a reference to the given string and assigns it to the PasswordKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetPasswordKey(v string) {
	o.PasswordKey = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetPath(v string) {
	o.Path = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetTls() ObservabilityPipelineMtlsServerTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineMtlsServerTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetTlsOk() (*ObservabilityPipelineMtlsServerTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineMtlsServerTls and assigns it to the Tls field.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetTls(v ObservabilityPipelineMtlsServerTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetType() ObservabilityPipelinePrometheusRemoteWriteSourceType {
	if o == nil {
		var ret ObservabilityPipelinePrometheusRemoteWriteSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetTypeOk() (*ObservabilityPipelinePrometheusRemoteWriteSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetType(v ObservabilityPipelinePrometheusRemoteWriteSourceType) {
	o.Type = v
}

// GetUsernameKey returns the UsernameKey field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetUsernameKey() string {
	if o == nil || o.UsernameKey == nil {
		var ret string
		return ret
	}
	return *o.UsernameKey
}

// GetUsernameKeyOk returns a tuple with the UsernameKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetUsernameKeyOk() (*string, bool) {
	if o == nil || o.UsernameKey == nil {
		return nil, false
	}
	return o.UsernameKey, true
}

// HasUsernameKey returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) HasUsernameKey() bool {
	return o != nil && o.UsernameKey != nil
}

// SetUsernameKey gets a reference to the given string and assigns it to the UsernameKey field.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetUsernameKey(v string) {
	o.UsernameKey = &v
}

// GetValidTokens returns the ValidTokens field value if set, zero value otherwise.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetValidTokens() []ObservabilityPipelinePrometheusRemoteWriteSourceValidToken {
	if o == nil || o.ValidTokens == nil {
		var ret []ObservabilityPipelinePrometheusRemoteWriteSourceValidToken
		return ret
	}
	return o.ValidTokens
}

// GetValidTokensOk returns a tuple with the ValidTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) GetValidTokensOk() (*[]ObservabilityPipelinePrometheusRemoteWriteSourceValidToken, bool) {
	if o == nil || o.ValidTokens == nil {
		return nil, false
	}
	return &o.ValidTokens, true
}

// HasValidTokens returns a boolean if a field has been set.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) HasValidTokens() bool {
	return o != nil && o.ValidTokens != nil
}

// SetValidTokens gets a reference to the given []ObservabilityPipelinePrometheusRemoteWriteSourceValidToken and assigns it to the ValidTokens field.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) SetValidTokens(v []ObservabilityPipelinePrometheusRemoteWriteSourceValidToken) {
	o.ValidTokens = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelinePrometheusRemoteWriteSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddressKey != nil {
		toSerialize["address_key"] = o.AddressKey
	}
	toSerialize["auth_strategy"] = o.AuthStrategy
	toSerialize["id"] = o.Id
	if o.PasswordKey != nil {
		toSerialize["password_key"] = o.PasswordKey
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type
	if o.UsernameKey != nil {
		toSerialize["username_key"] = o.UsernameKey
	}
	if o.ValidTokens != nil {
		toSerialize["valid_tokens"] = o.ValidTokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelinePrometheusRemoteWriteSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddressKey   *string                                                       `json:"address_key,omitempty"`
		AuthStrategy *ObservabilityPipelinePrometheusRemoteWriteSourceAuthStrategy `json:"auth_strategy"`
		Id           *string                                                       `json:"id"`
		PasswordKey  *string                                                       `json:"password_key,omitempty"`
		Path         *string                                                       `json:"path,omitempty"`
		Tls          *ObservabilityPipelineMtlsServerTls                           `json:"tls,omitempty"`
		Type         *ObservabilityPipelinePrometheusRemoteWriteSourceType         `json:"type"`
		UsernameKey  *string                                                       `json:"username_key,omitempty"`
		ValidTokens  []ObservabilityPipelinePrometheusRemoteWriteSourceValidToken  `json:"valid_tokens,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthStrategy == nil {
		return fmt.Errorf("required field auth_strategy missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"address_key", "auth_strategy", "id", "password_key", "path", "tls", "type", "username_key", "valid_tokens"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AddressKey = all.AddressKey
	if !all.AuthStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.AuthStrategy = *all.AuthStrategy
	}
	o.Id = *all.Id
	o.PasswordKey = all.PasswordKey
	o.Path = all.Path
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UsernameKey = all.UsernameKey
	o.ValidTokens = all.ValidTokens

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
