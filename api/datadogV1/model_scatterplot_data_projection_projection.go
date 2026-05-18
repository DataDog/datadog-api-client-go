// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotDataProjectionProjection The projection configuration for a scatterplot data projection request.
type ScatterplotDataProjectionProjection struct {
	// Dimension mappings for the scatterplot axes.
	Dimensions []ScatterplotDataProjectionDimension `json:"dimensions"`
	// The type of the scatterplot data projection.
	Type ScatterplotDataProjectionProjectionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScatterplotDataProjectionProjection instantiates a new ScatterplotDataProjectionProjection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotDataProjectionProjection(dimensions []ScatterplotDataProjectionDimension, typeVar ScatterplotDataProjectionProjectionType) *ScatterplotDataProjectionProjection {
	this := ScatterplotDataProjectionProjection{}
	this.Dimensions = dimensions
	this.Type = typeVar
	return &this
}

// NewScatterplotDataProjectionProjectionWithDefaults instantiates a new ScatterplotDataProjectionProjection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotDataProjectionProjectionWithDefaults() *ScatterplotDataProjectionProjection {
	this := ScatterplotDataProjectionProjection{}
	return &this
}

// GetDimensions returns the Dimensions field value.
func (o *ScatterplotDataProjectionProjection) GetDimensions() []ScatterplotDataProjectionDimension {
	if o == nil {
		var ret []ScatterplotDataProjectionDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *ScatterplotDataProjectionProjection) GetDimensionsOk() (*[]ScatterplotDataProjectionDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value.
func (o *ScatterplotDataProjectionProjection) SetDimensions(v []ScatterplotDataProjectionDimension) {
	o.Dimensions = v
}

// GetType returns the Type field value.
func (o *ScatterplotDataProjectionProjection) GetType() ScatterplotDataProjectionProjectionType {
	if o == nil {
		var ret ScatterplotDataProjectionProjectionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScatterplotDataProjectionProjection) GetTypeOk() (*ScatterplotDataProjectionProjectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScatterplotDataProjectionProjection) SetType(v ScatterplotDataProjectionProjectionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotDataProjectionProjection) MarshalJSON() ([]byte, error) {
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
func (o *ScatterplotDataProjectionProjection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dimensions *[]ScatterplotDataProjectionDimension    `json:"dimensions"`
		Type       *ScatterplotDataProjectionProjectionType `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
