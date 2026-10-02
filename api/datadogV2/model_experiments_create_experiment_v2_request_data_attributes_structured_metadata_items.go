// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems Metadata field key and values to set on the experiment.
type ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems struct {
	// Selected values for an enumerated metadata field.
	EnumValues []string `json:"enum_values,omitempty"`
	// Key that identifies the metadata field.
	FieldKey string `json:"field_key"`
	// Text value for a free-text metadata field.
	FreetextValue *string `json:"freetext_value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems(fieldKey string) *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems{}
	this.FieldKey = fieldKey
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems{}
	return &this
}

// GetEnumValues returns the EnumValues field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) GetEnumValues() []string {
	if o == nil || o.EnumValues == nil {
		var ret []string
		return ret
	}
	return o.EnumValues
}

// GetEnumValuesOk returns a tuple with the EnumValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) GetEnumValuesOk() (*[]string, bool) {
	if o == nil || o.EnumValues == nil {
		return nil, false
	}
	return &o.EnumValues, true
}

// HasEnumValues returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) HasEnumValues() bool {
	return o != nil && o.EnumValues != nil
}

// SetEnumValues gets a reference to the given []string and assigns it to the EnumValues field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) SetEnumValues(v []string) {
	o.EnumValues = v
}

// GetFieldKey returns the FieldKey field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) GetFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) GetFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldKey, true
}

// SetFieldKey sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) SetFieldKey(v string) {
	o.FieldKey = v
}

// GetFreetextValue returns the FreetextValue field value if set, zero value otherwise.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) GetFreetextValue() string {
	if o == nil || o.FreetextValue == nil {
		var ret string
		return ret
	}
	return *o.FreetextValue
}

// GetFreetextValueOk returns a tuple with the FreetextValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) GetFreetextValueOk() (*string, bool) {
	if o == nil || o.FreetextValue == nil {
		return nil, false
	}
	return o.FreetextValue, true
}

// HasFreetextValue returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) HasFreetextValue() bool {
	return o != nil && o.FreetextValue != nil
}

// SetFreetextValue gets a reference to the given string and assigns it to the FreetextValue field.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) SetFreetextValue(v string) {
	o.FreetextValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EnumValues != nil {
		toSerialize["enum_values"] = o.EnumValues
	}
	toSerialize["field_key"] = o.FieldKey
	if o.FreetextValue != nil {
		toSerialize["freetext_value"] = o.FreetextValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesStructuredMetadataItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnumValues    []string `json:"enum_values,omitempty"`
		FieldKey      *string  `json:"field_key"`
		FreetextValue *string  `json:"freetext_value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FieldKey == nil {
		return fmt.Errorf("required field field_key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enum_values", "field_key", "freetext_value"})
	} else {
		return err
	}
	o.EnumValues = all.EnumValues
	o.FieldKey = *all.FieldKey
	o.FreetextValue = all.FreetextValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
