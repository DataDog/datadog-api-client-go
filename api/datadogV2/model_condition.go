// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Condition Targeting condition details. A condition is either an inline
// predicate with `operator`, `attribute`, and `value`, or a reference to a
// saved filter with `saved_filter_id`. The inline fields are omitted for saved-filter
// references.
type Condition struct {
	// The user or request attribute to evaluate. Omitted for saved-filter references.
	Attribute *string `json:"attribute,omitempty"`
	// The timestamp when the condition was created.
	CreatedAt time.Time `json:"created_at"`
	// The unique identifier of the condition.
	Id uuid.UUID `json:"id"`
	// The operator used in a targeting condition.
	Operator *ConditionOperator `json:"operator,omitempty"`
	// The ID of the saved filter referenced by this condition, or null for inline conditions.
	SavedFilterId datadog.NullableUUID `json:"saved_filter_id,omitempty"`
	// The timestamp when the condition was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// Values used by the selected operator. Omitted for saved-filter references.
	Value []string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCondition instantiates a new Condition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCondition(createdAt time.Time, id uuid.UUID, updatedAt time.Time) *Condition {
	this := Condition{}
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewConditionWithDefaults instantiates a new Condition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConditionWithDefaults() *Condition {
	this := Condition{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *Condition) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *Condition) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *Condition) SetAttribute(v string) {
	o.Attribute = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Condition) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Condition) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Condition) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value.
func (o *Condition) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Condition) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Condition) SetId(v uuid.UUID) {
	o.Id = v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Condition) GetOperator() ConditionOperator {
	if o == nil || o.Operator == nil {
		var ret ConditionOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetOperatorOk() (*ConditionOperator, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Condition) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given ConditionOperator and assigns it to the Operator field.
func (o *Condition) SetOperator(v ConditionOperator) {
	o.Operator = &v
}

// GetSavedFilterId returns the SavedFilterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Condition) GetSavedFilterId() uuid.UUID {
	if o == nil || o.SavedFilterId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.SavedFilterId.Get()
}

// GetSavedFilterIdOk returns a tuple with the SavedFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Condition) GetSavedFilterIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.SavedFilterId.Get(), o.SavedFilterId.IsSet()
}

// HasSavedFilterId returns a boolean if a field has been set.
func (o *Condition) HasSavedFilterId() bool {
	return o != nil && o.SavedFilterId.IsSet()
}

// SetSavedFilterId gets a reference to the given datadog.NullableUUID and assigns it to the SavedFilterId field.
func (o *Condition) SetSavedFilterId(v uuid.UUID) {
	o.SavedFilterId.Set(&v)
}

// SetSavedFilterIdNil sets the value for SavedFilterId to be an explicit nil.
func (o *Condition) SetSavedFilterIdNil() {
	o.SavedFilterId.Set(nil)
}

// UnsetSavedFilterId ensures that no value is present for SavedFilterId, not even an explicit nil.
func (o *Condition) UnsetSavedFilterId() {
	o.SavedFilterId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Condition) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Condition) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Condition) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Condition) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Condition) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Condition) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *Condition) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Condition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["id"] = o.Id
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.SavedFilterId.IsSet() {
		toSerialize["saved_filter_id"] = o.SavedFilterId.Get()
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Condition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute     *string              `json:"attribute,omitempty"`
		CreatedAt     *time.Time           `json:"created_at"`
		Id            *uuid.UUID           `json:"id"`
		Operator      *ConditionOperator   `json:"operator,omitempty"`
		SavedFilterId datadog.NullableUUID `json:"saved_filter_id,omitempty"`
		UpdatedAt     *time.Time           `json:"updated_at"`
		Value         []string             `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "created_at", "id", "operator", "saved_filter_id", "updated_at", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = all.Attribute
	o.CreatedAt = *all.CreatedAt
	o.Id = *all.Id
	if all.Operator != nil && !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = all.Operator
	}
	o.SavedFilterId = all.SavedFilterId
	o.UpdatedAt = *all.UpdatedAt
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
