// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityModifierRulesResponse A list of severity modifier rules with pagination metadata.
type SeverityModifierRulesResponse struct {
	// A list of severity modifier rule data objects.
	Data []SeverityModifierRuleDataResponse `json:"data"`
	// Pagination links for the list of automation rules.
	Links SecurityAutomationRulesLinks `json:"links"`
	// Metadata for the list of automation rules.
	Meta SecurityAutomationRulesMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityModifierRulesResponse instantiates a new SeverityModifierRulesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityModifierRulesResponse(data []SeverityModifierRuleDataResponse, links SecurityAutomationRulesLinks, meta SecurityAutomationRulesMeta) *SeverityModifierRulesResponse {
	this := SeverityModifierRulesResponse{}
	this.Data = data
	this.Links = links
	this.Meta = meta
	return &this
}

// NewSeverityModifierRulesResponseWithDefaults instantiates a new SeverityModifierRulesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityModifierRulesResponseWithDefaults() *SeverityModifierRulesResponse {
	this := SeverityModifierRulesResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *SeverityModifierRulesResponse) GetData() []SeverityModifierRuleDataResponse {
	if o == nil {
		var ret []SeverityModifierRuleDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRulesResponse) GetDataOk() (*[]SeverityModifierRuleDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *SeverityModifierRulesResponse) SetData(v []SeverityModifierRuleDataResponse) {
	o.Data = v
}

// GetLinks returns the Links field value.
func (o *SeverityModifierRulesResponse) GetLinks() SecurityAutomationRulesLinks {
	if o == nil {
		var ret SecurityAutomationRulesLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRulesResponse) GetLinksOk() (*SecurityAutomationRulesLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value.
func (o *SeverityModifierRulesResponse) SetLinks(v SecurityAutomationRulesLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value.
func (o *SeverityModifierRulesResponse) GetMeta() SecurityAutomationRulesMeta {
	if o == nil {
		var ret SecurityAutomationRulesMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *SeverityModifierRulesResponse) GetMetaOk() (*SecurityAutomationRulesMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *SeverityModifierRulesResponse) SetMeta(v SecurityAutomationRulesMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityModifierRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityModifierRulesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]SeverityModifierRuleDataResponse `json:"data"`
		Links *SecurityAutomationRulesLinks       `json:"links"`
		Meta  *SecurityAutomationRulesMeta        `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Links == nil {
		return fmt.Errorf("required field links missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = *all.Links
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
