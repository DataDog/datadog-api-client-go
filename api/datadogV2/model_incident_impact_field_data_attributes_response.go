// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImpactFieldDataAttributesResponse Attributes of an impact field in a response.
type IncidentImpactFieldDataAttributesResponse struct {
	// The display name of the impact field.
	DisplayName string `json:"display_name"`
	// The choices for dropdown or multiselect fields.
	FieldChoices []IncidentImpactFieldChoice `json:"field_choices,omitempty"`
	// The type of an impact field.
	FieldType IncidentImpactFieldValueType `json:"field_type"`
	// The normalized name of the impact field.
	Name string `json:"name"`
	// The tag key associated with the field.
	TagKey datadog.NullableString `json:"tag_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImpactFieldDataAttributesResponse instantiates a new IncidentImpactFieldDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImpactFieldDataAttributesResponse(displayName string, fieldType IncidentImpactFieldValueType, name string) *IncidentImpactFieldDataAttributesResponse {
	this := IncidentImpactFieldDataAttributesResponse{}
	this.DisplayName = displayName
	this.FieldType = fieldType
	this.Name = name
	return &this
}

// NewIncidentImpactFieldDataAttributesResponseWithDefaults instantiates a new IncidentImpactFieldDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImpactFieldDataAttributesResponseWithDefaults() *IncidentImpactFieldDataAttributesResponse {
	this := IncidentImpactFieldDataAttributesResponse{}
	return &this
}

// GetDisplayName returns the DisplayName field value.
func (o *IncidentImpactFieldDataAttributesResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *IncidentImpactFieldDataAttributesResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *IncidentImpactFieldDataAttributesResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFieldChoices returns the FieldChoices field value if set, zero value otherwise.
func (o *IncidentImpactFieldDataAttributesResponse) GetFieldChoices() []IncidentImpactFieldChoice {
	if o == nil || o.FieldChoices == nil {
		var ret []IncidentImpactFieldChoice
		return ret
	}
	return o.FieldChoices
}

// GetFieldChoicesOk returns a tuple with the FieldChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImpactFieldDataAttributesResponse) GetFieldChoicesOk() (*[]IncidentImpactFieldChoice, bool) {
	if o == nil || o.FieldChoices == nil {
		return nil, false
	}
	return &o.FieldChoices, true
}

// HasFieldChoices returns a boolean if a field has been set.
func (o *IncidentImpactFieldDataAttributesResponse) HasFieldChoices() bool {
	return o != nil && o.FieldChoices != nil
}

// SetFieldChoices gets a reference to the given []IncidentImpactFieldChoice and assigns it to the FieldChoices field.
func (o *IncidentImpactFieldDataAttributesResponse) SetFieldChoices(v []IncidentImpactFieldChoice) {
	o.FieldChoices = v
}

// GetFieldType returns the FieldType field value.
func (o *IncidentImpactFieldDataAttributesResponse) GetFieldType() IncidentImpactFieldValueType {
	if o == nil {
		var ret IncidentImpactFieldValueType
		return ret
	}
	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *IncidentImpactFieldDataAttributesResponse) GetFieldTypeOk() (*IncidentImpactFieldValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value.
func (o *IncidentImpactFieldDataAttributesResponse) SetFieldType(v IncidentImpactFieldValueType) {
	o.FieldType = v
}

// GetName returns the Name field value.
func (o *IncidentImpactFieldDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentImpactFieldDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentImpactFieldDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetTagKey returns the TagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncidentImpactFieldDataAttributesResponse) GetTagKey() string {
	if o == nil || o.TagKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TagKey.Get()
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IncidentImpactFieldDataAttributesResponse) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagKey.Get(), o.TagKey.IsSet()
}

// HasTagKey returns a boolean if a field has been set.
func (o *IncidentImpactFieldDataAttributesResponse) HasTagKey() bool {
	return o != nil && o.TagKey.IsSet()
}

// SetTagKey gets a reference to the given datadog.NullableString and assigns it to the TagKey field.
func (o *IncidentImpactFieldDataAttributesResponse) SetTagKey(v string) {
	o.TagKey.Set(&v)
}

// SetTagKeyNil sets the value for TagKey to be an explicit nil.
func (o *IncidentImpactFieldDataAttributesResponse) SetTagKeyNil() {
	o.TagKey.Set(nil)
}

// UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil.
func (o *IncidentImpactFieldDataAttributesResponse) UnsetTagKey() {
	o.TagKey.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImpactFieldDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["display_name"] = o.DisplayName
	if o.FieldChoices != nil {
		toSerialize["field_choices"] = o.FieldChoices
	}
	toSerialize["field_type"] = o.FieldType
	toSerialize["name"] = o.Name
	if o.TagKey.IsSet() {
		toSerialize["tag_key"] = o.TagKey.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentImpactFieldDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName  *string                       `json:"display_name"`
		FieldChoices []IncidentImpactFieldChoice   `json:"field_choices,omitempty"`
		FieldType    *IncidentImpactFieldValueType `json:"field_type"`
		Name         *string                       `json:"name"`
		TagKey       datadog.NullableString        `json:"tag_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.FieldType == nil {
		return fmt.Errorf("required field field_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "field_choices", "field_type", "name", "tag_key"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = *all.DisplayName
	o.FieldChoices = all.FieldChoices
	if !all.FieldType.IsValid() {
		hasInvalidField = true
	} else {
		o.FieldType = *all.FieldType
	}
	o.Name = *all.Name
	o.TagKey = all.TagKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
