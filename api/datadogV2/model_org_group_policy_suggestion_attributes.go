// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicySuggestionAttributes Attributes of an org group policy suggestion.
type OrgGroupPolicySuggestionAttributes struct {
	// The ratio of member orgs whose configuration agrees on the recommended value.
	ConsensusRatio float64 `json:"consensus_ratio"`
	// The name of the suggested policy.
	PolicyName string `json:"policy_name"`
	// The recommended value for the policy, based on member org consensus.
	RecommendedValue interface{} `json:"recommended_value"`
	// The status of the policy suggestion.
	Status OrgGroupPolicySuggestionStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicySuggestionAttributes instantiates a new OrgGroupPolicySuggestionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicySuggestionAttributes(consensusRatio float64, policyName string, recommendedValue interface{}, status OrgGroupPolicySuggestionStatus) *OrgGroupPolicySuggestionAttributes {
	this := OrgGroupPolicySuggestionAttributes{}
	this.ConsensusRatio = consensusRatio
	this.PolicyName = policyName
	this.RecommendedValue = recommendedValue
	this.Status = status
	return &this
}

// NewOrgGroupPolicySuggestionAttributesWithDefaults instantiates a new OrgGroupPolicySuggestionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicySuggestionAttributesWithDefaults() *OrgGroupPolicySuggestionAttributes {
	this := OrgGroupPolicySuggestionAttributes{}
	return &this
}

// GetConsensusRatio returns the ConsensusRatio field value.
func (o *OrgGroupPolicySuggestionAttributes) GetConsensusRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ConsensusRatio
}

// GetConsensusRatioOk returns a tuple with the ConsensusRatio field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicySuggestionAttributes) GetConsensusRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsensusRatio, true
}

// SetConsensusRatio sets field value.
func (o *OrgGroupPolicySuggestionAttributes) SetConsensusRatio(v float64) {
	o.ConsensusRatio = v
}

// GetPolicyName returns the PolicyName field value.
func (o *OrgGroupPolicySuggestionAttributes) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicySuggestionAttributes) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value.
func (o *OrgGroupPolicySuggestionAttributes) SetPolicyName(v string) {
	o.PolicyName = v
}

// GetRecommendedValue returns the RecommendedValue field value.
func (o *OrgGroupPolicySuggestionAttributes) GetRecommendedValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RecommendedValue
}

// GetRecommendedValueOk returns a tuple with the RecommendedValue field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicySuggestionAttributes) GetRecommendedValueOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedValue, true
}

// SetRecommendedValue sets field value.
func (o *OrgGroupPolicySuggestionAttributes) SetRecommendedValue(v interface{}) {
	o.RecommendedValue = v
}

// GetStatus returns the Status field value.
func (o *OrgGroupPolicySuggestionAttributes) GetStatus() OrgGroupPolicySuggestionStatus {
	if o == nil {
		var ret OrgGroupPolicySuggestionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicySuggestionAttributes) GetStatusOk() (*OrgGroupPolicySuggestionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OrgGroupPolicySuggestionAttributes) SetStatus(v OrgGroupPolicySuggestionStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicySuggestionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["consensus_ratio"] = o.ConsensusRatio
	toSerialize["policy_name"] = o.PolicyName
	toSerialize["recommended_value"] = o.RecommendedValue
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicySuggestionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConsensusRatio   *float64                        `json:"consensus_ratio"`
		PolicyName       *string                         `json:"policy_name"`
		RecommendedValue *interface{}                    `json:"recommended_value"`
		Status           *OrgGroupPolicySuggestionStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConsensusRatio == nil {
		return fmt.Errorf("required field consensus_ratio missing")
	}
	if all.PolicyName == nil {
		return fmt.Errorf("required field policy_name missing")
	}
	if all.RecommendedValue == nil {
		return fmt.Errorf("required field recommended_value missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"consensus_ratio", "policy_name", "recommended_value", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConsensusRatio = *all.ConsensusRatio
	o.PolicyName = *all.PolicyName
	o.RecommendedValue = *all.RecommendedValue
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
