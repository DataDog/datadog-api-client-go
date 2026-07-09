// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetProjection Projection for the DDSQL host map request. Maps dataset columns to map dimensions: `node` identifies the entity, repeated `group` entries define the grouping hierarchy (outermost first), and `fill`/`size` drive the tile color and size.
type HostMapWidgetProjection struct {
	// List of column-to-dimension mappings for the projection.
	Dimensions []HostMapWidgetProjectionDimensionMapping `json:"dimensions"`
	// Type of the host map projection.
	Type HostMapWidgetProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetProjection instantiates a new HostMapWidgetProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetProjection(dimensions []HostMapWidgetProjectionDimensionMapping, typeVar HostMapWidgetProjectionType) *HostMapWidgetProjection {
	this := HostMapWidgetProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewHostMapWidgetProjectionWithDefaults instantiates a new HostMapWidgetProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetProjectionWithDefaults() *HostMapWidgetProjection {
	this := HostMapWidgetProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *HostMapWidgetProjection) GetDimensions() []HostMapWidgetProjectionDimensionMapping {
	if o == nil {
		var ret []HostMapWidgetProjectionDimensionMapping
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetProjection) GetDimensionsOk() (*[]HostMapWidgetProjectionDimensionMapping, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *HostMapWidgetProjection) SetDimensions(v []HostMapWidgetProjectionDimensionMapping) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *HostMapWidgetProjection) GetType() HostMapWidgetProjectionType {
	if o == nil {
		var ret HostMapWidgetProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetProjection) GetTypeOk() (*HostMapWidgetProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HostMapWidgetProjection) SetType(v HostMapWidgetProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]HostMapWidgetProjectionDimensionMapping `json:"dimensions"`
		Type       *HostMapWidgetProjectionType               `json:"type"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dimensions", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Dimensions = *all.Dimensions
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
