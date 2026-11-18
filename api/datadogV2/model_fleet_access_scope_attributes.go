// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAccessScopeAttributes Mutable fleet access scope attributes. PUT replaces all mutable attributes; omitted strings become empty and omitted feature flags become false.
type FleetAccessScopeAttributes struct {
	// Human-readable description of the access scope. Omission clears the description on update.
	Description *string `json:"description,omitempty"`
	// Fleet actions authorized on agents matching the access scope query.
	Features *FleetAccessScopeFeatures `json:"features,omitempty"`
	// Agent query selecting hosts covered by the scope. An empty query matches all hosts. Omission clears the query on update.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAccessScopeAttributes instantiates a new FleetAccessScopeAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAccessScopeAttributes() *FleetAccessScopeAttributes {
	this := FleetAccessScopeAttributes{}
	var description string = ""
	this.Description = &description
	var query string = ""
	this.Query = &query
	return &this
}

// NewFleetAccessScopeAttributesWithDefaults instantiates a new FleetAccessScopeAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAccessScopeAttributesWithDefaults() *FleetAccessScopeAttributes {
	this := FleetAccessScopeAttributes{}
	var description string = ""
	this.Description = &description
	var query string = ""
	this.Query = &query
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FleetAccessScopeAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FleetAccessScopeAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FleetAccessScopeAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FleetAccessScopeAttributes) GetFeatures() FleetAccessScopeFeatures {
	if o == nil || o.Features == nil {
		var ret FleetAccessScopeFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeAttributes) GetFeaturesOk() (*FleetAccessScopeFeatures, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FleetAccessScopeAttributes) HasFeatures() bool {
	return o != nil && o.Features != nil
}

// SetFeatures gets a reference to the given FleetAccessScopeFeatures and assigns it to the Features field.
func (o *FleetAccessScopeAttributes) SetFeatures(v FleetAccessScopeFeatures) {
	o.Features = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *FleetAccessScopeAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetAccessScopeAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *FleetAccessScopeAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *FleetAccessScopeAttributes) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAccessScopeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAccessScopeAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                   `json:"description,omitempty"`
		Features    *FleetAccessScopeFeatures `json:"features,omitempty"`
		Query       *string                   `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "features", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.Features != nil && all.Features.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Features = all.Features
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
