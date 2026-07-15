// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleOnCallResponderData Represents one position's (previous, current, or next) group of on-call responder shifts. Positions with no matching shift are omitted entirely from the response.
type ScheduleOnCallResponderData struct {
	// Attributes for one position's (previous, current, or next) group of on-call responder shifts.
	Attributes *ScheduleOnCallResponderDataAttributes `json:"attributes,omitempty"`
	// Unique identifier of this responder group.
	Id *string `json:"id,omitempty"`
	// Relationships for a single position's (previous, current, or next) responder group.
	Relationships *ScheduleOnCallResponderDataRelationships `json:"relationships,omitempty"`
	// Represents the resource type for a single position's (previous, current, or next) group of on-call responder shifts.
	Type ScheduleOnCallResponderDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleOnCallResponderData instantiates a new ScheduleOnCallResponderData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleOnCallResponderData(typeVar ScheduleOnCallResponderDataType) *ScheduleOnCallResponderData {
	this := ScheduleOnCallResponderData{}
	this.Type = typeVar
	return &this
}

// NewScheduleOnCallResponderDataWithDefaults instantiates a new ScheduleOnCallResponderData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleOnCallResponderDataWithDefaults() *ScheduleOnCallResponderData {
	this := ScheduleOnCallResponderData{}
	var typeVar ScheduleOnCallResponderDataType = SCHEDULEONCALLRESPONDERDATATYPE_SCHEDULE_ONCALL_RESPONDER
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScheduleOnCallResponderData) GetAttributes() ScheduleOnCallResponderDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret ScheduleOnCallResponderDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponderData) GetAttributesOk() (*ScheduleOnCallResponderDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ScheduleOnCallResponderData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given ScheduleOnCallResponderDataAttributes and assigns it to the Attributes field.
func (o *ScheduleOnCallResponderData) SetAttributes(v ScheduleOnCallResponderDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScheduleOnCallResponderData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponderData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScheduleOnCallResponderData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScheduleOnCallResponderData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScheduleOnCallResponderData) GetRelationships() ScheduleOnCallResponderDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ScheduleOnCallResponderDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponderData) GetRelationshipsOk() (*ScheduleOnCallResponderDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScheduleOnCallResponderData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given ScheduleOnCallResponderDataRelationships and assigns it to the Relationships field.
func (o *ScheduleOnCallResponderData) SetRelationships(v ScheduleOnCallResponderDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *ScheduleOnCallResponderData) GetType() ScheduleOnCallResponderDataType {
	if o == nil {
		var ret ScheduleOnCallResponderDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleOnCallResponderData) GetTypeOk() (*ScheduleOnCallResponderDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ScheduleOnCallResponderData) SetType(v ScheduleOnCallResponderDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleOnCallResponderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleOnCallResponderData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ScheduleOnCallResponderDataAttributes    `json:"attributes,omitempty"`
		Id            *string                                   `json:"id,omitempty"`
		Relationships *ScheduleOnCallResponderDataRelationships `json:"relationships,omitempty"`
		Type          *ScheduleOnCallResponderDataType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
