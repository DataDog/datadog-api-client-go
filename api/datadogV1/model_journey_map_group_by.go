// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JourneyMapGroupBy Group by configuration for the Journey Paths widget.
type JourneyMapGroupBy struct {
	// Facet to group by. Must be a valid tag name.
	Facet string `json:"facet"`
	// Audience the group by facet is read from.
	Source *JourneyMapGroupByAudience `json:"source,omitempty"`
	// Facet values to keep.
	ValueFilters []string `json:"value_filters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewJourneyMapGroupBy instantiates a new JourneyMapGroupBy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJourneyMapGroupBy(facet string, valueFilters []string) *JourneyMapGroupBy {
	this := JourneyMapGroupBy{}
	this.Facet = facet
	this.ValueFilters = valueFilters
	return &this
}

// NewJourneyMapGroupByWithDefaults instantiates a new JourneyMapGroupBy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJourneyMapGroupByWithDefaults() *JourneyMapGroupBy {
	this := JourneyMapGroupBy{}
	return &this
}

// GetFacet returns the Facet field value.
func (o *JourneyMapGroupBy) GetFacet() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *JourneyMapGroupBy) GetFacetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value.
func (o *JourneyMapGroupBy) SetFacet(v string) {
	o.Facet = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *JourneyMapGroupBy) GetSource() JourneyMapGroupByAudience {
	if o == nil || o.Source == nil {
		var ret JourneyMapGroupByAudience
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMapGroupBy) GetSourceOk() (*JourneyMapGroupByAudience, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *JourneyMapGroupBy) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given JourneyMapGroupByAudience and assigns it to the Source field.
func (o *JourneyMapGroupBy) SetSource(v JourneyMapGroupByAudience) {
	o.Source = &v
}

// GetValueFilters returns the ValueFilters field value.
func (o *JourneyMapGroupBy) GetValueFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ValueFilters
}

// GetValueFiltersOk returns a tuple with the ValueFilters field value
// and a boolean to check if the value has been set.
func (o *JourneyMapGroupBy) GetValueFiltersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueFilters, true
}

// SetValueFilters sets field value.
func (o *JourneyMapGroupBy) SetValueFilters(v []string) {
	o.ValueFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JourneyMapGroupBy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet"] = o.Facet
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	toSerialize["value_filters"] = o.ValueFilters
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JourneyMapGroupBy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet        *string                    `json:"facet"`
		Source       *JourneyMapGroupByAudience `json:"source,omitempty"`
		ValueFilters *[]string                  `json:"value_filters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Facet == nil {
		return fmt.Errorf("required field facet missing")
	}
	if all.ValueFilters == nil {
		return fmt.Errorf("required field value_filters missing")
	}

	hasInvalidField := false
	o.Facet = *all.Facet
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.ValueFilters = *all.ValueFilters

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
