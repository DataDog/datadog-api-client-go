// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationStrongLinkCreateRequestAttributes Attributes for creating a RUM operation strong link.
type RUMOperationStrongLinkCreateRequestAttributes struct {
	// The RUM application ID used when creating a stub operation from `operation_name`.
	ApplicationId *uuid.UUID `json:"application_id,omitempty"`
	// A description of the strong link.
	Description datadog.NullableString `json:"description,omitempty"`
	// The unique identifier of the feature to link.
	FeatureId string `json:"feature_id"`
	// The unique identifier of the RUM operation to link. Either `operation_id` or
	// `operation_name` is required.
	OperationId *string `json:"operation_id,omitempty"`
	// The name of the RUM operation to link. Either `operation_id` or `operation_name` is
	// required. If no operation with this name exists, a stub operation is created.
	OperationName *string `json:"operation_name,omitempty"`
	// The status of a RUM operation strong link.
	Status *RUMOperationStrongLinkStatus `json:"status,omitempty"`
	// A list of tags associated with the strong link.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMOperationStrongLinkCreateRequestAttributes instantiates a new RUMOperationStrongLinkCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMOperationStrongLinkCreateRequestAttributes(featureId string) *RUMOperationStrongLinkCreateRequestAttributes {
	this := RUMOperationStrongLinkCreateRequestAttributes{}
	this.FeatureId = featureId
	return &this
}

// NewRUMOperationStrongLinkCreateRequestAttributesWithDefaults instantiates a new RUMOperationStrongLinkCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMOperationStrongLinkCreateRequestAttributesWithDefaults() *RUMOperationStrongLinkCreateRequestAttributes {
	this := RUMOperationStrongLinkCreateRequestAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetApplicationId() uuid.UUID {
	if o == nil || o.ApplicationId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given uuid.UUID and assigns it to the ApplicationId field.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetApplicationId(v uuid.UUID) {
	o.ApplicationId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *RUMOperationStrongLinkCreateRequestAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetFeatureId returns the FeatureId field value.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetFeatureId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetFeatureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureId, true
}

// SetFeatureId sets field value.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetFeatureId(v string) {
	o.FeatureId = v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) HasOperationId() bool {
	return o != nil && o.OperationId != nil
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetOperationId(v string) {
	o.OperationId = &v
}

// GetOperationName returns the OperationName field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetOperationName() string {
	if o == nil || o.OperationName == nil {
		var ret string
		return ret
	}
	return *o.OperationName
}

// GetOperationNameOk returns a tuple with the OperationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetOperationNameOk() (*string, bool) {
	if o == nil || o.OperationName == nil {
		return nil, false
	}
	return o.OperationName, true
}

// HasOperationName returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) HasOperationName() bool {
	return o != nil && o.OperationName != nil
}

// SetOperationName gets a reference to the given string and assigns it to the OperationName field.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetOperationName(v string) {
	o.OperationName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetStatus() RUMOperationStrongLinkStatus {
	if o == nil || o.Status == nil {
		var ret RUMOperationStrongLinkStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetStatusOk() (*RUMOperationStrongLinkStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given RUMOperationStrongLinkStatus and assigns it to the Status field.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetStatus(v RUMOperationStrongLinkStatus) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RUMOperationStrongLinkCreateRequestAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RUMOperationStrongLinkCreateRequestAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMOperationStrongLinkCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["feature_id"] = o.FeatureId
	if o.OperationId != nil {
		toSerialize["operation_id"] = o.OperationId
	}
	if o.OperationName != nil {
		toSerialize["operation_name"] = o.OperationName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMOperationStrongLinkCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *uuid.UUID                    `json:"application_id,omitempty"`
		Description   datadog.NullableString        `json:"description,omitempty"`
		FeatureId     *string                       `json:"feature_id"`
		OperationId   *string                       `json:"operation_id,omitempty"`
		OperationName *string                       `json:"operation_name,omitempty"`
		Status        *RUMOperationStrongLinkStatus `json:"status,omitempty"`
		Tags          []string                      `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FeatureId == nil {
		return fmt.Errorf("required field feature_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "description", "feature_id", "operation_id", "operation_name", "status", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	o.Description = all.Description
	o.FeatureId = *all.FeatureId
	o.OperationId = all.OperationId
	o.OperationName = all.OperationName
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
