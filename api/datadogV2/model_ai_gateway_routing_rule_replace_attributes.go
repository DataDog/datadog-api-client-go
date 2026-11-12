// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleReplaceAttributes Partial route update. Omit tiers to preserve the entire graph, including hidden selections. Saved target IDs retain omitted settings. Replacing tiers must explicitly retain or remove every hidden occurrence.
type AiGatewayRoutingRuleReplaceAttributes struct {
	// Saved route revision. A mismatched value returns 409 without changes.
	ExpectedUpdatedAt *time.Time `json:"expected_updated_at,omitempty"`
	// Parent-local saved occurrence IDs explicitly removed from the submitted tiers. IDs must be distinct, belong to this route, and not also be retained.
	RemovedTargetIds []int64 `json:"removed_target_ids,omitempty"`
	//
	Tiers []AiGatewayRoutingRuleTierRequest `json:"tiers,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleReplaceAttributes instantiates a new AiGatewayRoutingRuleReplaceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleReplaceAttributes() *AiGatewayRoutingRuleReplaceAttributes {
	this := AiGatewayRoutingRuleReplaceAttributes{}
	return &this
}

// NewAiGatewayRoutingRuleReplaceAttributesWithDefaults instantiates a new AiGatewayRoutingRuleReplaceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleReplaceAttributesWithDefaults() *AiGatewayRoutingRuleReplaceAttributes {
	this := AiGatewayRoutingRuleReplaceAttributes{}
	return &this
}

// GetExpectedUpdatedAt returns the ExpectedUpdatedAt field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleReplaceAttributes) GetExpectedUpdatedAt() time.Time {
	if o == nil || o.ExpectedUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpectedUpdatedAt
}

// GetExpectedUpdatedAtOk returns a tuple with the ExpectedUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleReplaceAttributes) GetExpectedUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.ExpectedUpdatedAt == nil {
		return nil, false
	}
	return o.ExpectedUpdatedAt, true
}

// HasExpectedUpdatedAt returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleReplaceAttributes) HasExpectedUpdatedAt() bool {
	return o != nil && o.ExpectedUpdatedAt != nil
}

// SetExpectedUpdatedAt gets a reference to the given time.Time and assigns it to the ExpectedUpdatedAt field.
func (o *AiGatewayRoutingRuleReplaceAttributes) SetExpectedUpdatedAt(v time.Time) {
	o.ExpectedUpdatedAt = &v
}

// GetRemovedTargetIds returns the RemovedTargetIds field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleReplaceAttributes) GetRemovedTargetIds() []int64 {
	if o == nil || o.RemovedTargetIds == nil {
		var ret []int64
		return ret
	}
	return o.RemovedTargetIds
}

// GetRemovedTargetIdsOk returns a tuple with the RemovedTargetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleReplaceAttributes) GetRemovedTargetIdsOk() (*[]int64, bool) {
	if o == nil || o.RemovedTargetIds == nil {
		return nil, false
	}
	return &o.RemovedTargetIds, true
}

// HasRemovedTargetIds returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleReplaceAttributes) HasRemovedTargetIds() bool {
	return o != nil && o.RemovedTargetIds != nil
}

// SetRemovedTargetIds gets a reference to the given []int64 and assigns it to the RemovedTargetIds field.
func (o *AiGatewayRoutingRuleReplaceAttributes) SetRemovedTargetIds(v []int64) {
	o.RemovedTargetIds = v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleReplaceAttributes) GetTiers() []AiGatewayRoutingRuleTierRequest {
	if o == nil || o.Tiers == nil {
		var ret []AiGatewayRoutingRuleTierRequest
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleReplaceAttributes) GetTiersOk() (*[]AiGatewayRoutingRuleTierRequest, bool) {
	if o == nil || o.Tiers == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleReplaceAttributes) HasTiers() bool {
	return o != nil && o.Tiers != nil
}

// SetTiers gets a reference to the given []AiGatewayRoutingRuleTierRequest and assigns it to the Tiers field.
func (o *AiGatewayRoutingRuleReplaceAttributes) SetTiers(v []AiGatewayRoutingRuleTierRequest) {
	o.Tiers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleReplaceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExpectedUpdatedAt != nil {
		if o.ExpectedUpdatedAt.Nanosecond() == 0 {
			toSerialize["expected_updated_at"] = o.ExpectedUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expected_updated_at"] = o.ExpectedUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.RemovedTargetIds != nil {
		toSerialize["removed_target_ids"] = o.RemovedTargetIds
	}
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleReplaceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExpectedUpdatedAt *time.Time                        `json:"expected_updated_at,omitempty"`
		RemovedTargetIds  []int64                           `json:"removed_target_ids,omitempty"`
		Tiers             []AiGatewayRoutingRuleTierRequest `json:"tiers,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expected_updated_at", "removed_target_ids", "tiers"})
	} else {
		return err
	}
	o.ExpectedUpdatedAt = all.ExpectedUpdatedAt
	o.RemovedTargetIds = all.RemovedTargetIds
	o.Tiers = all.Tiers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
