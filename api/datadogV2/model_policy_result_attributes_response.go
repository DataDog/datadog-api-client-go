// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PolicyResultAttributesResponse Attributes of a policy evaluation result.
type PolicyResultAttributesResponse struct {
	// Whether the policy is active.
	Active bool `json:"active"`
	// The policy configuration payload.
	Payload interface{} `json:"payload"`
	// The type of policy being evaluated.
	PolicyType string `json:"policy_type"`
	// The organization UUID reference.
	ReferenceOrgUuid uuid.UUID `json:"reference_org_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPolicyResultAttributesResponse instantiates a new PolicyResultAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPolicyResultAttributesResponse(active bool, payload interface{}, policyType string, referenceOrgUuid uuid.UUID) *PolicyResultAttributesResponse {
	this := PolicyResultAttributesResponse{}
	this.Active = active
	this.Payload = payload
	this.PolicyType = policyType
	this.ReferenceOrgUuid = referenceOrgUuid
	return &this
}

// NewPolicyResultAttributesResponseWithDefaults instantiates a new PolicyResultAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPolicyResultAttributesResponseWithDefaults() *PolicyResultAttributesResponse {
	this := PolicyResultAttributesResponse{}
	return &this
}

// GetActive returns the Active field value.
func (o *PolicyResultAttributesResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *PolicyResultAttributesResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value.
func (o *PolicyResultAttributesResponse) SetActive(v bool) {
	o.Active = v
}

// GetPayload returns the Payload field value.
func (o *PolicyResultAttributesResponse) GetPayload() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *PolicyResultAttributesResponse) GetPayloadOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value.
func (o *PolicyResultAttributesResponse) SetPayload(v interface{}) {
	o.Payload = v
}

// GetPolicyType returns the PolicyType field value.
func (o *PolicyResultAttributesResponse) GetPolicyType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
func (o *PolicyResultAttributesResponse) GetPolicyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyType, true
}

// SetPolicyType sets field value.
func (o *PolicyResultAttributesResponse) SetPolicyType(v string) {
	o.PolicyType = v
}

// GetReferenceOrgUuid returns the ReferenceOrgUuid field value.
func (o *PolicyResultAttributesResponse) GetReferenceOrgUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ReferenceOrgUuid
}

// GetReferenceOrgUuidOk returns a tuple with the ReferenceOrgUuid field value
// and a boolean to check if the value has been set.
func (o *PolicyResultAttributesResponse) GetReferenceOrgUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceOrgUuid, true
}

// SetReferenceOrgUuid sets field value.
func (o *PolicyResultAttributesResponse) SetReferenceOrgUuid(v uuid.UUID) {
	o.ReferenceOrgUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PolicyResultAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["active"] = o.Active
	toSerialize["payload"] = o.Payload
	toSerialize["policy_type"] = o.PolicyType
	toSerialize["reference_org_uuid"] = o.ReferenceOrgUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PolicyResultAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Active           *bool        `json:"active"`
		Payload          *interface{} `json:"payload"`
		PolicyType       *string      `json:"policy_type"`
		ReferenceOrgUuid *uuid.UUID   `json:"reference_org_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Active == nil {
		return fmt.Errorf("required field active missing")
	}
	if all.Payload == nil {
		return fmt.Errorf("required field payload missing")
	}
	if all.PolicyType == nil {
		return fmt.Errorf("required field policy_type missing")
	}
	if all.ReferenceOrgUuid == nil {
		return fmt.Errorf("required field reference_org_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active", "payload", "policy_type", "reference_org_uuid"})
	} else {
		return err
	}
	o.Active = *all.Active
	o.Payload = *all.Payload
	o.PolicyType = *all.PolicyType
	o.ReferenceOrgUuid = *all.ReferenceOrgUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
