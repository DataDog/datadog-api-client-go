// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayPolicyTarget Resolved routing target in a policy response.
type AiGatewayPolicyTarget struct {
	// Saved target row ID. On replacement, retain this ID only within its original route and at most once. A stale or foreign ID is rejected.
	Id int64 `json:"id"`
	// Model name or pattern. Omitted when retaining a saved target by ID or when its provider is hidden.
	Model *string `json:"model,omitempty"`
	// Routing priority.
	Priority int64 `json:"priority"`
	// Provider slug. Omitted for dynamic resolution or when its provider is hidden.
	Provider datadog.NullableString `json:"provider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayPolicyTarget instantiates a new AiGatewayPolicyTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayPolicyTarget(id int64, priority int64) *AiGatewayPolicyTarget {
	this := AiGatewayPolicyTarget{}
	this.Id = id
	this.Priority = priority
	return &this
}

// NewAiGatewayPolicyTargetWithDefaults instantiates a new AiGatewayPolicyTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayPolicyTargetWithDefaults() *AiGatewayPolicyTarget {
	this := AiGatewayPolicyTarget{}
	return &this
}

// GetId returns the Id field value.
func (o *AiGatewayPolicyTarget) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTarget) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayPolicyTarget) SetId(v int64) {
	o.Id = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiGatewayPolicyTarget) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTarget) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiGatewayPolicyTarget) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiGatewayPolicyTarget) SetModel(v string) {
	o.Model = &v
}

// GetPriority returns the Priority field value.
func (o *AiGatewayPolicyTarget) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *AiGatewayPolicyTarget) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *AiGatewayPolicyTarget) SetPriority(v int64) {
	o.Priority = v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiGatewayPolicyTarget) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AiGatewayPolicyTarget) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *AiGatewayPolicyTarget) HasProvider() bool {
	return o != nil && o.Provider.IsSet()
}

// SetProvider gets a reference to the given datadog.NullableString and assigns it to the Provider field.
func (o *AiGatewayPolicyTarget) SetProvider(v string) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil.
func (o *AiGatewayPolicyTarget) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil.
func (o *AiGatewayPolicyTarget) UnsetProvider() {
	o.Provider.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayPolicyTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	toSerialize["priority"] = o.Priority
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayPolicyTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id       *int64                 `json:"id"`
		Model    *string                `json:"model,omitempty"`
		Priority *int64                 `json:"priority"`
		Provider datadog.NullableString `json:"provider,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "model", "priority", "provider"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Model = all.Model
	o.Priority = *all.Priority
	o.Provider = all.Provider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
