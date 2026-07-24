// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationStrongLinkResponseAttributes Attributes of a RUM operation strong link response.
type RUMOperationStrongLinkResponseAttributes struct {
	// The timestamp when the strong link was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A description of the strong link.
	Description datadog.NullableString `json:"description,omitempty"`
	// The unique identifier of the linked feature.
	FeatureId string `json:"feature_id"`
	// The unique identifier of the linked RUM operation.
	OperationId string `json:"operation_id"`
	// The status of a RUM operation strong link.
	Status RUMOperationStrongLinkStatus `json:"status"`
	// A list of tags associated with the strong link.
	Tags []string `json:"tags,omitempty"`
	// The timestamp when the strong link was last updated.
	UpdatedAt datadog.NullableTime `json:"updated_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMOperationStrongLinkResponseAttributes instantiates a new RUMOperationStrongLinkResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMOperationStrongLinkResponseAttributes(featureId string, operationId string, status RUMOperationStrongLinkStatus) *RUMOperationStrongLinkResponseAttributes {
	this := RUMOperationStrongLinkResponseAttributes{}
	this.FeatureId = featureId
	this.OperationId = operationId
	this.Status = status
	return &this
}

// NewRUMOperationStrongLinkResponseAttributesWithDefaults instantiates a new RUMOperationStrongLinkResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMOperationStrongLinkResponseAttributesWithDefaults() *RUMOperationStrongLinkResponseAttributes {
	this := RUMOperationStrongLinkResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RUMOperationStrongLinkResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationStrongLinkResponseAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationStrongLinkResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkResponseAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *RUMOperationStrongLinkResponseAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *RUMOperationStrongLinkResponseAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *RUMOperationStrongLinkResponseAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetFeatureId returns the FeatureId field value.
func (o *RUMOperationStrongLinkResponseAttributes) GetFeatureId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkResponseAttributes) GetFeatureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureId, true
}

// SetFeatureId sets field value.
func (o *RUMOperationStrongLinkResponseAttributes) SetFeatureId(v string) {
	o.FeatureId = v
}

// GetOperationId returns the OperationId field value.
func (o *RUMOperationStrongLinkResponseAttributes) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkResponseAttributes) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value.
func (o *RUMOperationStrongLinkResponseAttributes) SetOperationId(v string) {
	o.OperationId = v
}

// GetStatus returns the Status field value.
func (o *RUMOperationStrongLinkResponseAttributes) GetStatus() RUMOperationStrongLinkStatus {
	if o == nil {
		var ret RUMOperationStrongLinkStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkResponseAttributes) GetStatusOk() (*RUMOperationStrongLinkStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RUMOperationStrongLinkResponseAttributes) SetStatus(v RUMOperationStrongLinkStatus) {
	o.Status = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkResponseAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkResponseAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkResponseAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RUMOperationStrongLinkResponseAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationStrongLinkResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationStrongLinkResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkResponseAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given datadog.NullableTime and assigns it to the UpdatedAt field.
func (o *RUMOperationStrongLinkResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *RUMOperationStrongLinkResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *RUMOperationStrongLinkResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMOperationStrongLinkResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["feature_id"] = o.FeatureId
	toSerialize["operation_id"] = o.OperationId
	toSerialize["status"] = o.Status
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMOperationStrongLinkResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time                    `json:"created_at,omitempty"`
		Description datadog.NullableString        `json:"description,omitempty"`
		FeatureId   *string                       `json:"feature_id"`
		OperationId *string                       `json:"operation_id"`
		Status      *RUMOperationStrongLinkStatus `json:"status"`
		Tags        []string                      `json:"tags,omitempty"`
		UpdatedAt   datadog.NullableTime          `json:"updated_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FeatureId == nil {
		return fmt.Errorf("required field feature_id missing")
	}
	if all.OperationId == nil {
		return fmt.Errorf("required field operation_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "feature_id", "operation_id", "status", "tags", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.FeatureId = *all.FeatureId
	o.OperationId = *all.OperationId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
