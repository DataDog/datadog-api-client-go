// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRuleQueryCondition A query-based condition for an incident rule.
type IncidentRuleQueryCondition struct {
	// The normalized query string.
	NormalizedQuery datadog.NullableString `json:"normalized_query,omitempty"`
	// The raw query string.
	RawQuery datadog.NullableString `json:"raw_query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRuleQueryCondition instantiates a new IncidentRuleQueryCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRuleQueryCondition() *IncidentRuleQueryCondition {
	this := IncidentRuleQueryCondition{}
	return &this
}

// NewIncidentRuleQueryConditionWithDefaults instantiates a new IncidentRuleQueryCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRuleQueryConditionWithDefaults() *IncidentRuleQueryCondition {
	this := IncidentRuleQueryCondition{}
	return &this
}

// GetNormalizedQuery returns the NormalizedQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleQueryCondition) GetNormalizedQuery() string {
	if o == nil || o.NormalizedQuery.Get() == nil {
		var ret string
		return ret
	}
	return *o.NormalizedQuery.Get()
}

// GetNormalizedQueryOk returns a tuple with the NormalizedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleQueryCondition) GetNormalizedQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NormalizedQuery.Get(), o.NormalizedQuery.IsSet()
}

// HasNormalizedQuery returns a boolean if a field has been set.
func (o *IncidentRuleQueryCondition) HasNormalizedQuery() bool {
	return o != nil && o.NormalizedQuery.IsSet()
}

// SetNormalizedQuery gets a reference to the given datadog.NullableString and assigns it to the NormalizedQuery field.
func (o *IncidentRuleQueryCondition) SetNormalizedQuery(v string) {
	o.NormalizedQuery.Set(&v)
}

// SetNormalizedQueryNil sets the value for NormalizedQuery to be an explicit nil.
func (o *IncidentRuleQueryCondition) SetNormalizedQueryNil() {
	o.NormalizedQuery.Set(nil)
}

// UnsetNormalizedQuery ensures that no value is present for NormalizedQuery, not even an explicit nil.
func (o *IncidentRuleQueryCondition) UnsetNormalizedQuery() {
	o.NormalizedQuery.Unset()
}

// GetRawQuery returns the RawQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentRuleQueryCondition) GetRawQuery() string {
	if o == nil || o.RawQuery.Get() == nil {
		var ret string
		return ret
	}
	return *o.RawQuery.Get()
}

// GetRawQueryOk returns a tuple with the RawQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentRuleQueryCondition) GetRawQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawQuery.Get(), o.RawQuery.IsSet()
}

// HasRawQuery returns a boolean if a field has been set.
func (o *IncidentRuleQueryCondition) HasRawQuery() bool {
	return o != nil && o.RawQuery.IsSet()
}

// SetRawQuery gets a reference to the given datadog.NullableString and assigns it to the RawQuery field.
func (o *IncidentRuleQueryCondition) SetRawQuery(v string) {
	o.RawQuery.Set(&v)
}

// SetRawQueryNil sets the value for RawQuery to be an explicit nil.
func (o *IncidentRuleQueryCondition) SetRawQueryNil() {
	o.RawQuery.Set(nil)
}

// UnsetRawQuery ensures that no value is present for RawQuery, not even an explicit nil.
func (o *IncidentRuleQueryCondition) UnsetRawQuery() {
	o.RawQuery.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRuleQueryCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NormalizedQuery.IsSet() {
		toSerialize["normalized_query"] = o.NormalizedQuery.Get()
	}
	if o.RawQuery.IsSet() {
		toSerialize["raw_query"] = o.RawQuery.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRuleQueryCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NormalizedQuery datadog.NullableString `json:"normalized_query,omitempty"`
		RawQuery        datadog.NullableString `json:"raw_query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"normalized_query", "raw_query"})
	} else {
		return err
	}
	o.NormalizedQuery = all.NormalizedQuery
	o.RawQuery = all.RawQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
