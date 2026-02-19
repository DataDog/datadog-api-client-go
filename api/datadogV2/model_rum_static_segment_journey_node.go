// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumStaticSegmentJourneyNode A node in a journey query object.
type RumStaticSegmentJourneyNode struct {
	// The list of filters for this node.
	Filters []RumStaticSegmentJourneyFilter `json:"filters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumStaticSegmentJourneyNode instantiates a new RumStaticSegmentJourneyNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumStaticSegmentJourneyNode(filters []RumStaticSegmentJourneyFilter) *RumStaticSegmentJourneyNode {
	this := RumStaticSegmentJourneyNode{}
	this.Filters = filters
	return &this
}

// NewRumStaticSegmentJourneyNodeWithDefaults instantiates a new RumStaticSegmentJourneyNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumStaticSegmentJourneyNodeWithDefaults() *RumStaticSegmentJourneyNode {
	this := RumStaticSegmentJourneyNode{}
	return &this
}

// GetFilters returns the Filters field value.
func (o *RumStaticSegmentJourneyNode) GetFilters() []RumStaticSegmentJourneyFilter {
	if o == nil {
		var ret []RumStaticSegmentJourneyFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *RumStaticSegmentJourneyNode) GetFiltersOk() (*[]RumStaticSegmentJourneyFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *RumStaticSegmentJourneyNode) SetFilters(v []RumStaticSegmentJourneyFilter) {
	o.Filters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumStaticSegmentJourneyNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filters"] = o.Filters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumStaticSegmentJourneyNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filters *[]RumStaticSegmentJourneyFilter `json:"filters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filters"})
	} else {
		return err
	}
	o.Filters = *all.Filters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
