// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition A nullable range within a file defined by a start and end position, along with the file name.
type ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition struct {
	// A specific position (line and column) within a source file.
	End *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"end,omitempty"`
	// The name or path of the file containing this location.
	FileName *string `json:"file_name,omitempty"`
	// The semantic role associated with this file location.
	Role *string `json:"role,omitempty"`
	// A specific position (line and column) within a source file.
	Start *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition() *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePositionWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePositionWithDefaults() *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetEnd() ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition {
	if o == nil || o.End == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetEndOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition and assigns it to the End field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) SetEnd(v ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) {
	o.End = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) SetFileName(v string) {
	o.FileName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) SetRole(v string) {
	o.Role = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetStart() ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition {
	if o == nil || o.Start == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) GetStartOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition and assigns it to the Start field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) SetStart(v ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End      *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"end,omitempty"`
		FileName *string                                                          `json:"file_name,omitempty"`
		Role     *string                                                          `json:"role,omitempty"`
		Start    *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "file_name", "role", "start"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.End != nil && all.End.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.End = all.End
	o.FileName = all.FileName
	o.Role = all.Role
	if all.Start != nil && all.Start.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition handles when a null is used for ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition.
type NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition struct {
	value *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition
	isSet bool
}

// Get returns the associated value.
func (v NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) Get() *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) Set(val *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition initializes the struct as if Set has been called.
func NewNullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition(val *ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) *NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	return &NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
