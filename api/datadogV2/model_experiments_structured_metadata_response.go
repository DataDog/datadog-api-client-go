// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsStructuredMetadataResponse Metadata field values with their display name and type.
type ExperimentsStructuredMetadataResponse struct {
	// Selected values for an enumerated metadata field.
	EnumValues []string `json:"enum_values,omitempty"`
	// Display name of the metadata field.
	FieldDisplayName *string `json:"field_display_name,omitempty"`
	// Key that identifies the metadata field.
	FieldKey string `json:"field_key"`
	// Type of value stored in the structured metadata field.
	FieldType *ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType `json:"field_type,omitempty"`
	// Text value for a free-text metadata field.
	FreetextValue *string `json:"freetext_value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsStructuredMetadataResponse instantiates a new ExperimentsStructuredMetadataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsStructuredMetadataResponse(fieldKey string) *ExperimentsStructuredMetadataResponse {
	this := ExperimentsStructuredMetadataResponse{}
	this.FieldKey = fieldKey
	return &this
}

// NewExperimentsStructuredMetadataResponseWithDefaults instantiates a new ExperimentsStructuredMetadataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsStructuredMetadataResponseWithDefaults() *ExperimentsStructuredMetadataResponse {
	this := ExperimentsStructuredMetadataResponse{}
	return &this
}

// GetEnumValues returns the EnumValues field value if set, zero value otherwise.
func (o *ExperimentsStructuredMetadataResponse) GetEnumValues() []string {
	if o == nil || o.EnumValues == nil {
		var ret []string
		return ret
	}
	return o.EnumValues
}

// GetEnumValuesOk returns a tuple with the EnumValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsStructuredMetadataResponse) GetEnumValuesOk() (*[]string, bool) {
	if o == nil || o.EnumValues == nil {
		return nil, false
	}
	return &o.EnumValues, true
}

// HasEnumValues returns a boolean if a field has been set.
func (o *ExperimentsStructuredMetadataResponse) HasEnumValues() bool {
	return o != nil && o.EnumValues != nil
}

// SetEnumValues gets a reference to the given []string and assigns it to the EnumValues field.
func (o *ExperimentsStructuredMetadataResponse) SetEnumValues(v []string) {
	o.EnumValues = v
}

// GetFieldDisplayName returns the FieldDisplayName field value if set, zero value otherwise.
func (o *ExperimentsStructuredMetadataResponse) GetFieldDisplayName() string {
	if o == nil || o.FieldDisplayName == nil {
		var ret string
		return ret
	}
	return *o.FieldDisplayName
}

// GetFieldDisplayNameOk returns a tuple with the FieldDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsStructuredMetadataResponse) GetFieldDisplayNameOk() (*string, bool) {
	if o == nil || o.FieldDisplayName == nil {
		return nil, false
	}
	return o.FieldDisplayName, true
}

// HasFieldDisplayName returns a boolean if a field has been set.
func (o *ExperimentsStructuredMetadataResponse) HasFieldDisplayName() bool {
	return o != nil && o.FieldDisplayName != nil
}

// SetFieldDisplayName gets a reference to the given string and assigns it to the FieldDisplayName field.
func (o *ExperimentsStructuredMetadataResponse) SetFieldDisplayName(v string) {
	o.FieldDisplayName = &v
}

// GetFieldKey returns the FieldKey field value.
func (o *ExperimentsStructuredMetadataResponse) GetFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value
// and a boolean to check if the value has been set.
func (o *ExperimentsStructuredMetadataResponse) GetFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldKey, true
}

// SetFieldKey sets field value.
func (o *ExperimentsStructuredMetadataResponse) SetFieldKey(v string) {
	o.FieldKey = v
}

// GetFieldType returns the FieldType field value if set, zero value otherwise.
func (o *ExperimentsStructuredMetadataResponse) GetFieldType() ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType {
	if o == nil || o.FieldType == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType
		return ret
	}
	return *o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsStructuredMetadataResponse) GetFieldTypeOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType, bool) {
	if o == nil || o.FieldType == nil {
		return nil, false
	}
	return o.FieldType, true
}

// HasFieldType returns a boolean if a field has been set.
func (o *ExperimentsStructuredMetadataResponse) HasFieldType() bool {
	return o != nil && o.FieldType != nil
}

// SetFieldType gets a reference to the given ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType and assigns it to the FieldType field.
func (o *ExperimentsStructuredMetadataResponse) SetFieldType(v ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType) {
	o.FieldType = &v
}

// GetFreetextValue returns the FreetextValue field value if set, zero value otherwise.
func (o *ExperimentsStructuredMetadataResponse) GetFreetextValue() string {
	if o == nil || o.FreetextValue == nil {
		var ret string
		return ret
	}
	return *o.FreetextValue
}

// GetFreetextValueOk returns a tuple with the FreetextValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsStructuredMetadataResponse) GetFreetextValueOk() (*string, bool) {
	if o == nil || o.FreetextValue == nil {
		return nil, false
	}
	return o.FreetextValue, true
}

// HasFreetextValue returns a boolean if a field has been set.
func (o *ExperimentsStructuredMetadataResponse) HasFreetextValue() bool {
	return o != nil && o.FreetextValue != nil
}

// SetFreetextValue gets a reference to the given string and assigns it to the FreetextValue field.
func (o *ExperimentsStructuredMetadataResponse) SetFreetextValue(v string) {
	o.FreetextValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsStructuredMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnumValues != nil {
		toSerialize["enum_values"] = o.EnumValues
	}
	if o.FieldDisplayName != nil {
		toSerialize["field_display_name"] = o.FieldDisplayName
	}
	toSerialize["field_key"] = o.FieldKey
	if o.FieldType != nil {
		toSerialize["field_type"] = o.FieldType
	}
	if o.FreetextValue != nil {
		toSerialize["freetext_value"] = o.FreetextValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsStructuredMetadataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnumValues       []string                                                                            `json:"enum_values,omitempty"`
		FieldDisplayName *string                                                                             `json:"field_display_name,omitempty"`
		FieldKey         *string                                                                             `json:"field_key"`
		FieldType        *ExperimentsPatchExperimentV2ResponseDataAttributesStructuredMetadataItemsFieldType `json:"field_type,omitempty"`
		FreetextValue    *string                                                                             `json:"freetext_value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FieldKey == nil {
		return fmt.Errorf("required field field_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enum_values", "field_display_name", "field_key", "field_type", "freetext_value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EnumValues = all.EnumValues
	o.FieldDisplayName = all.FieldDisplayName
	o.FieldKey = *all.FieldKey
	if all.FieldType != nil && !all.FieldType.IsValid() {
		hasInvalidField = true
	} else {
		o.FieldType = all.FieldType
	}
	o.FreetextValue = all.FreetextValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
