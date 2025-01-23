// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetRuleVersionHistoryResponse Response object containing the version history of a rule.
type GetRuleVersionHistoryResponse struct {
	// The number of rule versions.
	Count *int32 `json:"count,omitempty"`
	// A list of rule versions.
	Data []RuleVersionHistory `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetRuleVersionHistoryResponse instantiates a new GetRuleVersionHistoryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetRuleVersionHistoryResponse() *GetRuleVersionHistoryResponse {
	this := GetRuleVersionHistoryResponse{}
	return &this
}

// NewGetRuleVersionHistoryResponseWithDefaults instantiates a new GetRuleVersionHistoryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetRuleVersionHistoryResponseWithDefaults() *GetRuleVersionHistoryResponse {
	this := GetRuleVersionHistoryResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetRuleVersionHistoryResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleVersionHistoryResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetRuleVersionHistoryResponse) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetRuleVersionHistoryResponse) SetCount(v int32) {
	o.Count = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRuleVersionHistoryResponse) GetData() []RuleVersionHistory {
	if o == nil || o.Data == nil {
		var ret []RuleVersionHistory
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRuleVersionHistoryResponse) GetDataOk() (*[]RuleVersionHistory, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRuleVersionHistoryResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []RuleVersionHistory and assigns it to the Data field.
func (o *GetRuleVersionHistoryResponse) SetData(v []RuleVersionHistory) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetRuleVersionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetRuleVersionHistoryResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count *int32               `json:"count,omitempty"`
		Data  []RuleVersionHistory `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "data"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
