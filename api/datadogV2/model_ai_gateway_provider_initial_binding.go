// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayProviderInitialBinding Reserved initial provider access binding. Initial provider sharing is unavailable.
type AiGatewayProviderInitialBinding struct {
	// Access principals in type:id form.
	Principals []string `json:"principals"`
	// Provider access relation. Editor also grants View.
	Relation AiGatewayProviderInitialBindingRelation `json:"relation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayProviderInitialBinding instantiates a new AiGatewayProviderInitialBinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayProviderInitialBinding(principals []string, relation AiGatewayProviderInitialBindingRelation) *AiGatewayProviderInitialBinding {
	this := AiGatewayProviderInitialBinding{}
	this.Principals = principals
	this.Relation = relation
	return &this
}

// NewAiGatewayProviderInitialBindingWithDefaults instantiates a new AiGatewayProviderInitialBinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayProviderInitialBindingWithDefaults() *AiGatewayProviderInitialBinding {
	this := AiGatewayProviderInitialBinding{}
	return &this
}

// GetPrincipals returns the Principals field value.
func (o *AiGatewayProviderInitialBinding) GetPrincipals() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderInitialBinding) GetPrincipalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principals, true
}

// SetPrincipals sets field value.
func (o *AiGatewayProviderInitialBinding) SetPrincipals(v []string) {
	o.Principals = v
}

// GetRelation returns the Relation field value.
func (o *AiGatewayProviderInitialBinding) GetRelation() AiGatewayProviderInitialBindingRelation {
	if o == nil {
		var ret AiGatewayProviderInitialBindingRelation
		return ret
	}
	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *AiGatewayProviderInitialBinding) GetRelationOk() (*AiGatewayProviderInitialBindingRelation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value.
func (o *AiGatewayProviderInitialBinding) SetRelation(v AiGatewayProviderInitialBindingRelation) {
	o.Relation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayProviderInitialBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["principals"] = o.Principals
	toSerialize["relation"] = o.Relation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayProviderInitialBinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Principals *[]string                                `json:"principals"`
		Relation   *AiGatewayProviderInitialBindingRelation `json:"relation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Principals == nil {
		return fmt.Errorf("required field principals missing")
	}
	if all.Relation == nil {
		return fmt.Errorf("required field relation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"principals", "relation"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Principals = *all.Principals
	if !all.Relation.IsValid() {
		hasInvalidField = true
	} else {
		o.Relation = *all.Relation
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
