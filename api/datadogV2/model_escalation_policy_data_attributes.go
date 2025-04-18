// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyDataAttributes Defines the main attributes of an escalation policy, such as its description, name, and behavior on policy end.
type EscalationPolicyDataAttributes struct {
	// Provides a detailed text description of the escalation policy.
	Description *string `json:"description,omitempty"`
	// Specifies the name of the escalation policy.
	Name *string `json:"name,omitempty"`
	// Indicates whether the page is automatically resolved when the policy ends.
	ResolvePageOnPolicyEnd *bool `json:"resolve_page_on_policy_end,omitempty"`
	// Specifies how many times the escalation sequence is retried if there is no response.
	Retries *int64 `json:"retries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyDataAttributes instantiates a new EscalationPolicyDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyDataAttributes() *EscalationPolicyDataAttributes {
	this := EscalationPolicyDataAttributes{}
	return &this
}

// NewEscalationPolicyDataAttributesWithDefaults instantiates a new EscalationPolicyDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyDataAttributesWithDefaults() *EscalationPolicyDataAttributes {
	this := EscalationPolicyDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EscalationPolicyDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EscalationPolicyDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EscalationPolicyDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EscalationPolicyDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EscalationPolicyDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EscalationPolicyDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetResolvePageOnPolicyEnd returns the ResolvePageOnPolicyEnd field value if set, zero value otherwise.
func (o *EscalationPolicyDataAttributes) GetResolvePageOnPolicyEnd() bool {
	if o == nil || o.ResolvePageOnPolicyEnd == nil {
		var ret bool
		return ret
	}
	return *o.ResolvePageOnPolicyEnd
}

// GetResolvePageOnPolicyEndOk returns a tuple with the ResolvePageOnPolicyEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataAttributes) GetResolvePageOnPolicyEndOk() (*bool, bool) {
	if o == nil || o.ResolvePageOnPolicyEnd == nil {
		return nil, false
	}
	return o.ResolvePageOnPolicyEnd, true
}

// HasResolvePageOnPolicyEnd returns a boolean if a field has been set.
func (o *EscalationPolicyDataAttributes) HasResolvePageOnPolicyEnd() bool {
	return o != nil && o.ResolvePageOnPolicyEnd != nil
}

// SetResolvePageOnPolicyEnd gets a reference to the given bool and assigns it to the ResolvePageOnPolicyEnd field.
func (o *EscalationPolicyDataAttributes) SetResolvePageOnPolicyEnd(v bool) {
	o.ResolvePageOnPolicyEnd = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *EscalationPolicyDataAttributes) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataAttributes) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *EscalationPolicyDataAttributes) HasRetries() bool {
	return o != nil && o.Retries != nil
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *EscalationPolicyDataAttributes) SetRetries(v int64) {
	o.Retries = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ResolvePageOnPolicyEnd != nil {
		toSerialize["resolve_page_on_policy_end"] = o.ResolvePageOnPolicyEnd
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description            *string `json:"description,omitempty"`
		Name                   *string `json:"name,omitempty"`
		ResolvePageOnPolicyEnd *bool   `json:"resolve_page_on_policy_end,omitempty"`
		Retries                *int64  `json:"retries,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "resolve_page_on_policy_end", "retries"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = all.Name
	o.ResolvePageOnPolicyEnd = all.ResolvePageOnPolicyEnd
	o.Retries = all.Retries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
