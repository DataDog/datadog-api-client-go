// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRoutingRuleAttributes Routing rule configuration and metadata.
type AiGatewayRoutingRuleAttributes struct {
	// Time when the resource was created.
	CreatedAt time.Time `json:"created_at"`
	// Stable numeric route identifier used by scoped policies.
	RouteId *int64 `json:"route_id,omitempty"`
	//
	Tiers []AiGatewayRoutingRuleTier `json:"tiers"`
	// Time when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRoutingRuleAttributes instantiates a new AiGatewayRoutingRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRoutingRuleAttributes(createdAt time.Time, tiers []AiGatewayRoutingRuleTier, updatedAt time.Time) *AiGatewayRoutingRuleAttributes {
	this := AiGatewayRoutingRuleAttributes{}
	this.CreatedAt = createdAt
	this.Tiers = tiers
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiGatewayRoutingRuleAttributesWithDefaults instantiates a new AiGatewayRoutingRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRoutingRuleAttributesWithDefaults() *AiGatewayRoutingRuleAttributes {
	this := AiGatewayRoutingRuleAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayRoutingRuleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayRoutingRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *AiGatewayRoutingRuleAttributes) GetRouteId() int64 {
	if o == nil || o.RouteId == nil {
		var ret int64
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleAttributes) GetRouteIdOk() (*int64, bool) {
	if o == nil || o.RouteId == nil {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *AiGatewayRoutingRuleAttributes) HasRouteId() bool {
	return o != nil && o.RouteId != nil
}

// SetRouteId gets a reference to the given int64 and assigns it to the RouteId field.
func (o *AiGatewayRoutingRuleAttributes) SetRouteId(v int64) {
	o.RouteId = &v
}

// GetTiers returns the Tiers field value.
func (o *AiGatewayRoutingRuleAttributes) GetTiers() []AiGatewayRoutingRuleTier {
	if o == nil {
		var ret []AiGatewayRoutingRuleTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleAttributes) GetTiersOk() (*[]AiGatewayRoutingRuleTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tiers, true
}

// SetTiers sets field value.
func (o *AiGatewayRoutingRuleAttributes) SetTiers(v []AiGatewayRoutingRuleTier) {
	o.Tiers = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayRoutingRuleAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRoutingRuleAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayRoutingRuleAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRoutingRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.RouteId != nil {
		toSerialize["route_id"] = o.RouteId
	}
	toSerialize["tiers"] = o.Tiers
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRoutingRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                  `json:"created_at"`
		RouteId   *int64                      `json:"route_id,omitempty"`
		Tiers     *[]AiGatewayRoutingRuleTier `json:"tiers"`
		UpdatedAt *time.Time                  `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Tiers == nil {
		return fmt.Errorf("required field tiers missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "route_id", "tiers", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.RouteId = all.RouteId
	o.Tiers = *all.Tiers
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
