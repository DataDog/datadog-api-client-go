// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicies A paginated list of escalation policies, with any related resources requested through `include`.
type EscalationPolicies struct {
	// A list of escalation policies.
	Data []EscalationPolicyListItem `json:"data"`
	// Provides any included related resources, such as teams, steps, or targets, returned with the policies.
	Included []EscalationPolicyIncluded `json:"included,omitempty"`
	// Metadata included in the response when listing escalation policies.
	Meta *EscalationPoliciesResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicies instantiates a new EscalationPolicies object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicies(data []EscalationPolicyListItem) *EscalationPolicies {
	this := EscalationPolicies{}
	this.Data = data
	return &this
}

// NewEscalationPoliciesWithDefaults instantiates a new EscalationPolicies object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPoliciesWithDefaults() *EscalationPolicies {
	this := EscalationPolicies{}
	return &this
}

// GetData returns the Data field value.
func (o *EscalationPolicies) GetData() []EscalationPolicyListItem {
	if o == nil {
		var ret []EscalationPolicyListItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicies) GetDataOk() (*[]EscalationPolicyListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *EscalationPolicies) SetData(v []EscalationPolicyListItem) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *EscalationPolicies) GetIncluded() []EscalationPolicyIncluded {
	if o == nil || o.Included == nil {
		var ret []EscalationPolicyIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicies) GetIncludedOk() (*[]EscalationPolicyIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *EscalationPolicies) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []EscalationPolicyIncluded and assigns it to the Included field.
func (o *EscalationPolicies) SetIncluded(v []EscalationPolicyIncluded) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EscalationPolicies) GetMeta() EscalationPoliciesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret EscalationPoliciesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicies) GetMetaOk() (*EscalationPoliciesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EscalationPolicies) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given EscalationPoliciesResponseMeta and assigns it to the Meta field.
func (o *EscalationPolicies) SetMeta(v EscalationPoliciesResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicies) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]EscalationPolicyListItem     `json:"data"`
		Included []EscalationPolicyIncluded      `json:"included,omitempty"`
		Meta     *EscalationPoliciesResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
