// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TreeDataProjectionToplist Test-only generated schema for the dashboard versioning proof of concept.
type TreeDataProjectionToplist struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Dimensions []TreeDataProjectionToplistDimensionsItems `json:"dimensions"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupStyles []GroupStyleRule `json:"group_styles,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Type TreeDataProjectionToplistType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewTreeDataProjectionToplist instantiates a new TreeDataProjectionToplist object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTreeDataProjectionToplist(dimensions []TreeDataProjectionToplistDimensionsItems, typeVar TreeDataProjectionToplistType) *TreeDataProjectionToplist {
	this := TreeDataProjectionToplist{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewTreeDataProjectionToplistWithDefaults instantiates a new TreeDataProjectionToplist object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTreeDataProjectionToplistWithDefaults() *TreeDataProjectionToplist {
	this := TreeDataProjectionToplist{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *TreeDataProjectionToplist) GetDimensions() []TreeDataProjectionToplistDimensionsItems {
	if o == nil {
		var ret []TreeDataProjectionToplistDimensionsItems
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TreeDataProjectionToplist) GetDimensionsOk() (*[]TreeDataProjectionToplistDimensionsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *TreeDataProjectionToplist) SetDimensions(v []TreeDataProjectionToplistDimensionsItems) {
	o.Dimensions = v
}

// GetGroupStyles returns the GroupStyles field value if set, zero value otherwise.
func (o *TreeDataProjectionToplist) GetGroupStyles() []GroupStyleRule {
	if o == nil || o.GroupStyles == nil {
		var ret []GroupStyleRule
		return ret
	}
	return o.GroupStyles
}

// GetGroupStylesOk returns a tuple with the GroupStyles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TreeDataProjectionToplist) GetGroupStylesOk() (*[]GroupStyleRule, bool) {
	if o == nil || o.GroupStyles == nil {
		return nil, false
	}
	return &o.GroupStyles, true
}

// HasGroupStyles returns a boolean if a field has been set.
func (o *TreeDataProjectionToplist) HasGroupStyles() bool {
	return o != nil && o.GroupStyles != nil
}

// SetGroupStyles gets a reference to the given []GroupStyleRule and assigns it to the GroupStyles field.
func (o *TreeDataProjectionToplist) SetGroupStyles(v []GroupStyleRule) {
	o.GroupStyles = v
}

// GetType returns the Type field value.
func (o *TreeDataProjectionToplist) GetType() TreeDataProjectionToplistType {
	if o == nil {
		var ret TreeDataProjectionToplistType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TreeDataProjectionToplist) GetTypeOk() (*TreeDataProjectionToplistType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TreeDataProjectionToplist) SetType(v TreeDataProjectionToplistType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TreeDataProjectionToplist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	if o.GroupStyles != nil {
		toSerialize["group_styles"] = o.GroupStyles
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TreeDataProjectionToplist) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions  *[]TreeDataProjectionToplistDimensionsItems `json:"dimensions"`
		GroupStyles []GroupStyleRule                            `json:"group_styles,omitempty"`
		Type        *TreeDataProjectionToplistType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dimensions == nil {
		return fmt.Errorf("required field dimensions missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	o.Dimensions = *all.Dimensions
	o.GroupStyles = all.GroupStyles
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
