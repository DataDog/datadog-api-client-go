// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsSankeyRequestAttributes Attributes of a Sankey request.
type ProductAnalyticsSankeyRequestAttributes struct {
	// The shape of the Sankey diagram, expressed as the facets to flow between and how many steps to show.
	Definition ProductAnalyticsSankeyDefinition `json:"definition"`
	// Selects the sessions a Sankey diagram is built from.
	Search ProductAnalyticsSankeySearch `json:"search"`
	// The time window a Sankey query covers.
	Time ProductAnalyticsSankeyTime `json:"time"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsSankeyRequestAttributes instantiates a new ProductAnalyticsSankeyRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsSankeyRequestAttributes(definition ProductAnalyticsSankeyDefinition, search ProductAnalyticsSankeySearch, time ProductAnalyticsSankeyTime) *ProductAnalyticsSankeyRequestAttributes {
	this := ProductAnalyticsSankeyRequestAttributes{}
	this.Definition = definition
	this.Search = search
	this.Time = time
	return &this
}

// NewProductAnalyticsSankeyRequestAttributesWithDefaults instantiates a new ProductAnalyticsSankeyRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsSankeyRequestAttributesWithDefaults() *ProductAnalyticsSankeyRequestAttributes {
	this := ProductAnalyticsSankeyRequestAttributes{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetDefinition() ProductAnalyticsSankeyDefinition {
	if o == nil {
		var ret ProductAnalyticsSankeyDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetDefinitionOk() (*ProductAnalyticsSankeyDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetDefinition(v ProductAnalyticsSankeyDefinition) {
	o.Definition = v
}

// GetSearch returns the Search field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetSearch() ProductAnalyticsSankeySearch {
	if o == nil {
		var ret ProductAnalyticsSankeySearch
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetSearchOk() (*ProductAnalyticsSankeySearch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetSearch(v ProductAnalyticsSankeySearch) {
	o.Search = v
}

// GetTime returns the Time field value.
func (o *ProductAnalyticsSankeyRequestAttributes) GetTime() ProductAnalyticsSankeyTime {
	if o == nil {
		var ret ProductAnalyticsSankeyTime
		return ret
	}
	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsSankeyRequestAttributes) GetTimeOk() (*ProductAnalyticsSankeyTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value.
func (o *ProductAnalyticsSankeyRequestAttributes) SetTime(v ProductAnalyticsSankeyTime) {
	o.Time = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsSankeyRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	toSerialize["search"] = o.Search
	toSerialize["time"] = o.Time

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsSankeyRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition *ProductAnalyticsSankeyDefinition `json:"definition"`
		Search     *ProductAnalyticsSankeySearch     `json:"search"`
		Time       *ProductAnalyticsSankeyTime       `json:"time"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	if all.Search == nil {
		return fmt.Errorf("required field search missing")
	}
	if all.Time == nil {
		return fmt.Errorf("required field time missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "search", "time"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	if all.Search.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Search = *all.Search
	if all.Time.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Time = *all.Time

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
