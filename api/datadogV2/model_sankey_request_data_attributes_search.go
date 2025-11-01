// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyRequestDataAttributesSearch
type SankeyRequestDataAttributesSearch struct {
	//
	AudienceFilters *SankeyRequestDataAttributesSearchAudienceFilters `json:"audience_filters,omitempty"`
	//
	Occurrences *SankeyRequestDataAttributesSearchOccurrences `json:"occurrences,omitempty"`
	//
	Query *string `json:"query,omitempty"`
	//
	SubqueryId *string `json:"subquery_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSankeyRequestDataAttributesSearch instantiates a new SankeyRequestDataAttributesSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyRequestDataAttributesSearch() *SankeyRequestDataAttributesSearch {
	this := SankeyRequestDataAttributesSearch{}
	return &this
}

// NewSankeyRequestDataAttributesSearchWithDefaults instantiates a new SankeyRequestDataAttributesSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyRequestDataAttributesSearchWithDefaults() *SankeyRequestDataAttributesSearch {
	this := SankeyRequestDataAttributesSearch{}
	return &this
}

// GetAudienceFilters returns the AudienceFilters field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearch) GetAudienceFilters() SankeyRequestDataAttributesSearchAudienceFilters {
	if o == nil || o.AudienceFilters == nil {
		var ret SankeyRequestDataAttributesSearchAudienceFilters
		return ret
	}
	return *o.AudienceFilters
}

// GetAudienceFiltersOk returns a tuple with the AudienceFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearch) GetAudienceFiltersOk() (*SankeyRequestDataAttributesSearchAudienceFilters, bool) {
	if o == nil || o.AudienceFilters == nil {
		return nil, false
	}
	return o.AudienceFilters, true
}

// HasAudienceFilters returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearch) HasAudienceFilters() bool {
	return o != nil && o.AudienceFilters != nil
}

// SetAudienceFilters gets a reference to the given SankeyRequestDataAttributesSearchAudienceFilters and assigns it to the AudienceFilters field.
func (o *SankeyRequestDataAttributesSearch) SetAudienceFilters(v SankeyRequestDataAttributesSearchAudienceFilters) {
	o.AudienceFilters = &v
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearch) GetOccurrences() SankeyRequestDataAttributesSearchOccurrences {
	if o == nil || o.Occurrences == nil {
		var ret SankeyRequestDataAttributesSearchOccurrences
		return ret
	}
	return *o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearch) GetOccurrencesOk() (*SankeyRequestDataAttributesSearchOccurrences, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearch) HasOccurrences() bool {
	return o != nil && o.Occurrences != nil
}

// SetOccurrences gets a reference to the given SankeyRequestDataAttributesSearchOccurrences and assigns it to the Occurrences field.
func (o *SankeyRequestDataAttributesSearch) SetOccurrences(v SankeyRequestDataAttributesSearchOccurrences) {
	o.Occurrences = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearch) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearch) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearch) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SankeyRequestDataAttributesSearch) SetQuery(v string) {
	o.Query = &v
}

// GetSubqueryId returns the SubqueryId field value if set, zero value otherwise.
func (o *SankeyRequestDataAttributesSearch) GetSubqueryId() string {
	if o == nil || o.SubqueryId == nil {
		var ret string
		return ret
	}
	return *o.SubqueryId
}

// GetSubqueryIdOk returns a tuple with the SubqueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyRequestDataAttributesSearch) GetSubqueryIdOk() (*string, bool) {
	if o == nil || o.SubqueryId == nil {
		return nil, false
	}
	return o.SubqueryId, true
}

// HasSubqueryId returns a boolean if a field has been set.
func (o *SankeyRequestDataAttributesSearch) HasSubqueryId() bool {
	return o != nil && o.SubqueryId != nil
}

// SetSubqueryId gets a reference to the given string and assigns it to the SubqueryId field.
func (o *SankeyRequestDataAttributesSearch) SetSubqueryId(v string) {
	o.SubqueryId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyRequestDataAttributesSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AudienceFilters != nil {
		toSerialize["audience_filters"] = o.AudienceFilters
	}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
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
func (o *SankeyRequestDataAttributesSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AudienceFilters *SankeyRequestDataAttributesSearchAudienceFilters `json:"audience_filters,omitempty"`
		Occurrences     *SankeyRequestDataAttributesSearchOccurrences     `json:"occurrences,omitempty"`
		Query           *string                                           `json:"query,omitempty"`
		SubqueryId      *string                                           `json:"subquery_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"audience_filters", "occurrences", "query", "subquery_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AudienceFilters != nil && all.AudienceFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AudienceFilters = all.AudienceFilters
	if all.Occurrences != nil && all.Occurrences.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Occurrences = all.Occurrences
	o.Query = all.Query
	o.SubqueryId = all.SubqueryId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
