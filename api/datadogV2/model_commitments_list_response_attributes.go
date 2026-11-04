// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsListResponseAttributes Attributes of a commitments list response.
type CommitmentsListResponseAttributes struct {
	// Array of commitment items.
	Commitments []CommitmentsListItem `json:"commitments"`
	// Metadata for a commitments list response.
	Meta *CommitmentsListMeta `json:"meta,omitempty"`
	// The query that produces the response values, in the same format as dashboard and notebook widget
	// requests, so a client can graph the values in a widget. Each formula's `id` names the response value
	// that it produces.
	QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsListResponseAttributes instantiates a new CommitmentsListResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsListResponseAttributes(commitments []CommitmentsListItem) *CommitmentsListResponseAttributes {
	this := CommitmentsListResponseAttributes{}
	this.Commitments = commitments
	return &this
}

// NewCommitmentsListResponseAttributesWithDefaults instantiates a new CommitmentsListResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsListResponseAttributesWithDefaults() *CommitmentsListResponseAttributes {
	this := CommitmentsListResponseAttributes{}
	return &this
}

// GetCommitments returns the Commitments field value.
func (o *CommitmentsListResponseAttributes) GetCommitments() []CommitmentsListItem {
	if o == nil {
		var ret []CommitmentsListItem
		return ret
	}
	return o.Commitments
}

// GetCommitmentsOk returns a tuple with the Commitments field value
// and a boolean to check if the value has been set.
func (o *CommitmentsListResponseAttributes) GetCommitmentsOk() (*[]CommitmentsListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commitments, true
}

// SetCommitments sets field value.
func (o *CommitmentsListResponseAttributes) SetCommitments(v []CommitmentsListItem) {
	o.Commitments = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsListResponseAttributes) GetMeta() CommitmentsListMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsListMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsListResponseAttributes) GetMetaOk() (*CommitmentsListMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsListResponseAttributes) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsListMeta and assigns it to the Meta field.
func (o *CommitmentsListResponseAttributes) SetMeta(v CommitmentsListMeta) {
	o.Meta = &v
}

// GetQueryDefinition returns the QueryDefinition field value if set, zero value otherwise.
func (o *CommitmentsListResponseAttributes) GetQueryDefinition() CommitmentsQueryDefinition {
	if o == nil || o.QueryDefinition == nil {
		var ret CommitmentsQueryDefinition
		return ret
	}
	return *o.QueryDefinition
}

// GetQueryDefinitionOk returns a tuple with the QueryDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsListResponseAttributes) GetQueryDefinitionOk() (*CommitmentsQueryDefinition, bool) {
	if o == nil || o.QueryDefinition == nil {
		return nil, false
	}
	return o.QueryDefinition, true
}

// HasQueryDefinition returns a boolean if a field has been set.
func (o *CommitmentsListResponseAttributes) HasQueryDefinition() bool {
	return o != nil && o.QueryDefinition != nil
}

// SetQueryDefinition gets a reference to the given CommitmentsQueryDefinition and assigns it to the QueryDefinition field.
func (o *CommitmentsListResponseAttributes) SetQueryDefinition(v CommitmentsQueryDefinition) {
	o.QueryDefinition = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsListResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commitments"] = o.Commitments
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.QueryDefinition != nil {
		toSerialize["query_definition"] = o.QueryDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsListResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Commitments     *[]CommitmentsListItem      `json:"commitments"`
		Meta            *CommitmentsListMeta        `json:"meta,omitempty"`
		QueryDefinition *CommitmentsQueryDefinition `json:"query_definition,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Commitments == nil {
		return fmt.Errorf("required field commitments missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commitments", "meta", "query_definition"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Commitments = *all.Commitments
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if all.QueryDefinition != nil && all.QueryDefinition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.QueryDefinition = all.QueryDefinition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
