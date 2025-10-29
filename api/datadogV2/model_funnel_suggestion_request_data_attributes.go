// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionRequestDataAttributes
type FunnelSuggestionRequestDataAttributes struct {
	//
	DataSource *string `json:"data_source,omitempty"`
	//
	Search *FunnelSuggestionRequestDataAttributesSearch `json:"search,omitempty"`
	//
	TermSearch *FunnelSuggestionRequestDataAttributesTermSearch `json:"term_search,omitempty"`
	//
	Time *FunnelSuggestionRequestDataAttributesTime `json:"time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelSuggestionRequestDataAttributes instantiates a new FunnelSuggestionRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelSuggestionRequestDataAttributes() *FunnelSuggestionRequestDataAttributes {
	this := FunnelSuggestionRequestDataAttributes{}
	return &this
}

// NewFunnelSuggestionRequestDataAttributesWithDefaults instantiates a new FunnelSuggestionRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelSuggestionRequestDataAttributesWithDefaults() *FunnelSuggestionRequestDataAttributes {
	this := FunnelSuggestionRequestDataAttributes{}
	return &this
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributes) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *FunnelSuggestionRequestDataAttributes) SetDataSource(v string) {
	o.DataSource = &v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributes) GetSearch() FunnelSuggestionRequestDataAttributesSearch {
	if o == nil || o.Search == nil {
		var ret FunnelSuggestionRequestDataAttributesSearch
		return ret
	}
	return *o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributes) GetSearchOk() (*FunnelSuggestionRequestDataAttributesSearch, bool) {
	if o == nil || o.Search == nil {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributes) HasSearch() bool {
	return o != nil && o.Search != nil
}

// SetSearch gets a reference to the given FunnelSuggestionRequestDataAttributesSearch and assigns it to the Search field.
func (o *FunnelSuggestionRequestDataAttributes) SetSearch(v FunnelSuggestionRequestDataAttributesSearch) {
	o.Search = &v
}

// GetTermSearch returns the TermSearch field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributes) GetTermSearch() FunnelSuggestionRequestDataAttributesTermSearch {
	if o == nil || o.TermSearch == nil {
		var ret FunnelSuggestionRequestDataAttributesTermSearch
		return ret
	}
	return *o.TermSearch
}

// GetTermSearchOk returns a tuple with the TermSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributes) GetTermSearchOk() (*FunnelSuggestionRequestDataAttributesTermSearch, bool) {
	if o == nil || o.TermSearch == nil {
		return nil, false
	}
	return o.TermSearch, true
}

// HasTermSearch returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributes) HasTermSearch() bool {
	return o != nil && o.TermSearch != nil
}

// SetTermSearch gets a reference to the given FunnelSuggestionRequestDataAttributesTermSearch and assigns it to the TermSearch field.
func (o *FunnelSuggestionRequestDataAttributes) SetTermSearch(v FunnelSuggestionRequestDataAttributesTermSearch) {
	o.TermSearch = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *FunnelSuggestionRequestDataAttributes) GetTime() FunnelSuggestionRequestDataAttributesTime {
	if o == nil || o.Time == nil {
		var ret FunnelSuggestionRequestDataAttributesTime
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionRequestDataAttributes) GetTimeOk() (*FunnelSuggestionRequestDataAttributesTime, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *FunnelSuggestionRequestDataAttributes) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given FunnelSuggestionRequestDataAttributesTime and assigns it to the Time field.
func (o *FunnelSuggestionRequestDataAttributes) SetTime(v FunnelSuggestionRequestDataAttributesTime) {
	o.Time = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelSuggestionRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataSource != nil {
		toSerialize["data_source"] = o.DataSource
	}
	if o.Search != nil {
		toSerialize["search"] = o.Search
	}
	if o.TermSearch != nil {
		toSerialize["term_search"] = o.TermSearch
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelSuggestionRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource *string                                          `json:"data_source,omitempty"`
		Search     *FunnelSuggestionRequestDataAttributesSearch     `json:"search,omitempty"`
		TermSearch *FunnelSuggestionRequestDataAttributesTermSearch `json:"term_search,omitempty"`
		Time       *FunnelSuggestionRequestDataAttributesTime       `json:"time,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_source", "search", "term_search", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataSource = all.DataSource
	if all.Search != nil && all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = all.Search
	if all.TermSearch != nil && all.TermSearch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TermSearch = all.TermSearch
	if all.Time != nil && all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
