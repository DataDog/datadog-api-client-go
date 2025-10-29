// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionRequestDataAttributesSearch
type FunnelSuggestionRequestDataAttributesSearch struct {
	//
	CrossSessionFilter *string `json:"cross_session_filter,omitempty"`
	//
	QueryString *string `json:"query_string,omitempty"`
	//
	Steps []FunnelSuggestionRequestDataAttributesSearchStepsItems `json:"steps,omitempty"`
	//
	SubqueryId *string `json:"subquery_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelSuggestionRequestDataAttributesSearch instantiates a new FunnelSuggestionRequestDataAttributesSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelSuggestionRequestDataAttributesSearch() *FunnelSuggestionRequestDataAttributesSearch {
	this := FunnelSuggestionRequestDataAttributesSearch{}
	return &this
}

// NewFunnelSuggestionRequestDataAttributesSearchWithDefaults instantiates a new FunnelSuggestionRequestDataAttributesSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelSuggestionRequestDataAttributesSearchWithDefaults() *FunnelSuggestionRequestDataAttributesSearch {
	this := FunnelSuggestionRequestDataAttributesSearch{}
	return &this
}

// GetCrossSessionFilter returns the CrossSessionFilter field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetCrossSessionFilter() string {
	if o == nil || o.CrossSessionFilter == nil {
		var ret string
		return ret
	}
	return *o.CrossSessionFilter
}

// GetCrossSessionFilterOk returns a tuple with the CrossSessionFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetCrossSessionFilterOk() (*string, bool) {
	if o == nil || o.CrossSessionFilter == nil {
		return nil, false
	}
	return o.CrossSessionFilter, true
}

// HasCrossSessionFilter returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) HasCrossSessionFilter() bool {
	return o != nil && o.CrossSessionFilter != nil
}

// SetCrossSessionFilter gets a reference to the given string and assigns it to the CrossSessionFilter field.
func (o *FunnelSuggestionRequestDataAttributesSearch) SetCrossSessionFilter(v string) {
	o.CrossSessionFilter = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) HasQueryString() bool {
	return o != nil && o.QueryString != nil
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *FunnelSuggestionRequestDataAttributesSearch) SetQueryString(v string) {
	o.QueryString = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetSteps() []FunnelSuggestionRequestDataAttributesSearchStepsItems {
	if o == nil || o.Steps == nil {
		var ret []FunnelSuggestionRequestDataAttributesSearchStepsItems
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetStepsOk() (*[]FunnelSuggestionRequestDataAttributesSearchStepsItems, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []FunnelSuggestionRequestDataAttributesSearchStepsItems and assigns it to the Steps field.
func (o *FunnelSuggestionRequestDataAttributesSearch) SetSteps(v []FunnelSuggestionRequestDataAttributesSearchStepsItems) {
	o.Steps = v
}

// GetSubqueryId returns the SubqueryId field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetSubqueryId() string {
	if o == nil || o.SubqueryId == nil {
		var ret string
		return ret
	}
	return *o.SubqueryId
}

// GetSubqueryIdOk returns a tuple with the SubqueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) GetSubqueryIdOk() (*string, bool) {
	if o == nil || o.SubqueryId == nil {
		return nil, false
	}
	return o.SubqueryId, true
}

// HasSubqueryId returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributesSearch) HasSubqueryId() bool {
	return o != nil && o.SubqueryId != nil
}

// SetSubqueryId gets a reference to the given string and assigns it to the SubqueryId field.
func (o *FunnelSuggestionRequestDataAttributesSearch) SetSubqueryId(v string) {
	o.SubqueryId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelSuggestionRequestDataAttributesSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossSessionFilter != nil {
		toSerialize["cross_session_filter"] = o.CrossSessionFilter
	}
	if o.QueryString != nil {
		toSerialize["query_string"] = o.QueryString
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.SubqueryId != nil {
		toSerialize["subquery_id"] = o.SubqueryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelSuggestionRequestDataAttributesSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossSessionFilter *string                                                 `json:"cross_session_filter,omitempty"`
		QueryString        *string                                                 `json:"query_string,omitempty"`
		Steps              []FunnelSuggestionRequestDataAttributesSearchStepsItems `json:"steps,omitempty"`
		SubqueryId         *string                                                 `json:"subquery_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cross_session_filter", "query_string", "steps", "subquery_id"})
	} else {
		return err
	}
	o.CrossSessionFilter = all.CrossSessionFilter
	o.QueryString = all.QueryString
	o.Steps = all.Steps
	o.SubqueryId = all.SubqueryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
