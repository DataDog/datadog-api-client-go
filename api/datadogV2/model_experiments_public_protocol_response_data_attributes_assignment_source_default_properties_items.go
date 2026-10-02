// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems A default property supplied by the protocol's assignment source.
type ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems struct {
	// Source column that supplies this assignment property.
	ColumnName *string `json:"column_name,omitempty"`
	// Data type of the source column.
	ColumnType *string `json:"column_type,omitempty"`
	// Display name of the assignment source property.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems instantiates a new ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems() *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems {
	this := ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItemsWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItemsWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems {
	this := ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetColumnType returns the ColumnType field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) GetColumnType() string {
	if o == nil || o.ColumnType == nil {
		var ret string
		return ret
	}
	return *o.ColumnType
}

// GetColumnTypeOk returns a tuple with the ColumnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) GetColumnTypeOk() (*string, bool) {
	if o == nil || o.ColumnType == nil {
		return nil, false
	}
	return o.ColumnType, true
}

// HasColumnType returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) HasColumnType() bool {
	return o != nil && o.ColumnType != nil
}

// SetColumnType gets a reference to the given string and assigns it to the ColumnType field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) SetColumnType(v string) {
	o.ColumnType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	if o.ColumnType != nil {
		toSerialize["column_type"] = o.ColumnType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsPublicProtocolResponseDataAttributesAssignmentSourceDefaultPropertiesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName *string `json:"column_name,omitempty"`
		ColumnType *string `json:"column_type,omitempty"`
		Name       *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "column_type", "name"})
	} else {
		return err
	}
	o.ColumnName = all.ColumnName
	o.ColumnType = all.ColumnType
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
