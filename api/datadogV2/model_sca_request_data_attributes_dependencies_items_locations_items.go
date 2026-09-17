// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItemsLocationsItems The source code location where a dependency is declared, including block, name, namespace, and version positions within the file.
type ScaRequestDataAttributesDependenciesItemsLocationsItems struct {
	// A range within a file defined by a start and end position, along with the file name.
	Block *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"block,omitempty"`
	// A nullable range within a file defined by a start and end position, along with the file name.
	Name NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition `json:"name,omitempty"`
	// A nullable range within a file defined by a start and end position, along with the file name.
	Namespace NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition `json:"namespace,omitempty"`
	// A nullable range within a file defined by a start and end position, along with the file name.
	Version NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItems instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItemsLocationsItems() *ScaRequestDataAttributesDependenciesItemsLocationsItems {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItems{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsWithDefaults() *ScaRequestDataAttributesDependenciesItemsLocationsItems {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItems{}
	return &this
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetBlock() ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	if o == nil || o.Block == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetBlockOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasBlock() bool {
	return o != nil && o.Block != nil
}

// SetBlock gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition and assigns it to the Block field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetBlock(v ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) {
	o.Block = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetName() ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	if o == nil || o.Name.Get() == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetNameOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasName() bool {
	return o != nil && o.Name.IsSet()
}

// SetName gets a reference to the given NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition and assigns it to the Name field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetName(v ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) UnsetName() {
	o.Name.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetNamespace() ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	if o == nil || o.Namespace.Get() == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetNamespaceOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasNamespace() bool {
	return o != nil && o.Namespace.IsSet()
}

// SetNamespace gets a reference to the given NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition and assigns it to the Namespace field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetNamespace(v ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetVersion() ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition {
	if o == nil || o.Version.Get() == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetVersionOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition and assigns it to the Version field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetVersion(v ScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) UnsetVersion() {
	o.Version.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItemsLocationsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Block != nil {
		toSerialize["block"] = o.Block
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Block     *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition                `json:"block,omitempty"`
		Name      NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition `json:"name,omitempty"`
		Namespace NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition `json:"namespace,omitempty"`
		Version   NullableScaRequestDataAttributesDependenciesItemsLocationsItemsNullableFilePosition `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"block", "name", "namespace", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Block != nil && all.Block.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Block = all.Block
	o.Name = all.Name
	o.Namespace = all.Namespace
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
