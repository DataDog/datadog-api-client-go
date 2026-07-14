// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedRolePolicy Policy configuration for a user-defined role.
type IncidentUserDefinedRolePolicy struct {
	// Whether this role can only be assigned to one responder at a time.
	IsSingle bool `json:"is_single"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedRolePolicy instantiates a new IncidentUserDefinedRolePolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedRolePolicy(isSingle bool) *IncidentUserDefinedRolePolicy {
	this := IncidentUserDefinedRolePolicy{}
	this.IsSingle = isSingle
	return &this
}

// NewIncidentUserDefinedRolePolicyWithDefaults instantiates a new IncidentUserDefinedRolePolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedRolePolicyWithDefaults() *IncidentUserDefinedRolePolicy {
	this := IncidentUserDefinedRolePolicy{}
	return &this
}

// GetIsSingle returns the IsSingle field value.
func (o *IncidentUserDefinedRolePolicy) GetIsSingle() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsSingle
}

// GetIsSingleOk returns a tuple with the IsSingle field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedRolePolicy) GetIsSingleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSingle, true
}

// SetIsSingle sets field value.
func (o *IncidentUserDefinedRolePolicy) SetIsSingle(v bool) {
	o.IsSingle = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedRolePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_single"] = o.IsSingle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedRolePolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsSingle *bool `json:"is_single"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsSingle == nil {
		return fmt.Errorf("required field is_single missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_single"})
	} else {
		return err
	}
	o.IsSingle = *all.IsSingle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
