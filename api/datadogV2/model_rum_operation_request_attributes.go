// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMOperationRequestAttributes Attributes for creating or updating a RUM operation.
type RUMOperationRequestAttributes struct {
	// The RUM application ID the operation belongs to.
	ApplicationId datadog.NullableUUID `json:"application_id,omitempty"`
	// The category of the RUM operation.
	Category datadog.NullableString `json:"category,omitempty"`
	// A description of the RUM operation.
	Description datadog.NullableString `json:"description,omitempty"`
	// A human-readable display name for the RUM operation.
	DisplayName *string `json:"display_name,omitempty"`
	// The list of feature IDs associated with the RUM operation.
	FeatureIds []string `json:"feature_ids,omitempty"`
	// The definition of a RUM operation's journey, used to detect it from RUM events.
	JourneyRum RUMOperationJourneyRum `json:"journey_rum"`
	// The unique name of the RUM operation. Must not contain spaces.
	Name string `json:"name"`
	// A list of tags associated with the RUM operation.
	Tags []string `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMOperationRequestAttributes instantiates a new RUMOperationRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMOperationRequestAttributes(journeyRum RUMOperationJourneyRum, name string, tags []string) *RUMOperationRequestAttributes {
	this := RUMOperationRequestAttributes{}
	this.JourneyRum = journeyRum
	this.Name = name
	this.Tags = tags
	return &this
}

// NewRUMOperationRequestAttributesWithDefaults instantiates a new RUMOperationRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMOperationRequestAttributesWithDefaults() *RUMOperationRequestAttributes {
	this := RUMOperationRequestAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationRequestAttributes) GetApplicationId() uuid.UUID {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationRequestAttributes) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *RUMOperationRequestAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId.IsSet()
}

// SetApplicationId gets a reference to the given datadog.NullableUUID and assigns it to the ApplicationId field.
func (o *RUMOperationRequestAttributes) SetApplicationId(v uuid.UUID) {
	o.ApplicationId.Set(&v)
}

// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil.
func (o *RUMOperationRequestAttributes) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil.
func (o *RUMOperationRequestAttributes) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationRequestAttributes) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationRequestAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *RUMOperationRequestAttributes) HasCategory() bool {
	return o != nil && o.Category.IsSet()
}

// SetCategory gets a reference to the given datadog.NullableString and assigns it to the Category field.
func (o *RUMOperationRequestAttributes) SetCategory(v string) {
	o.Category.Set(&v)
}

// SetCategoryNil sets the value for Category to be an explicit nil.
func (o *RUMOperationRequestAttributes) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil.
func (o *RUMOperationRequestAttributes) UnsetCategory() {
	o.Category.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RUMOperationRequestAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RUMOperationRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RUMOperationRequestAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *RUMOperationRequestAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *RUMOperationRequestAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *RUMOperationRequestAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RUMOperationRequestAttributes) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationRequestAttributes) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RUMOperationRequestAttributes) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RUMOperationRequestAttributes) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *RUMOperationRequestAttributes) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMOperationRequestAttributes) GetFeatureIdsOk() (*[]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return &o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *RUMOperationRequestAttributes) HasFeatureIds() bool {
	return o != nil && o.FeatureIds != nil
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *RUMOperationRequestAttributes) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetJourneyRum returns the JourneyRum field value.
func (o *RUMOperationRequestAttributes) GetJourneyRum() RUMOperationJourneyRum {
	if o == nil {
		var ret RUMOperationJourneyRum
		return ret
	}
	return o.JourneyRum
}

// GetJourneyRumOk returns a tuple with the JourneyRum field value
// and a boolean to check if the value has been set.
func (o *RUMOperationRequestAttributes) GetJourneyRumOk() (*RUMOperationJourneyRum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JourneyRum, true
}

// SetJourneyRum sets field value.
func (o *RUMOperationRequestAttributes) SetJourneyRum(v RUMOperationJourneyRum) {
	o.JourneyRum = v
}

// GetName returns the Name field value.
func (o *RUMOperationRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMOperationRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMOperationRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value.
func (o *RUMOperationRequestAttributes) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *RUMOperationRequestAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *RUMOperationRequestAttributes) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMOperationRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId.IsSet() {
		toSerialize["application_id"] = o.ApplicationId.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.FeatureIds != nil {
		toSerialize["feature_ids"] = o.FeatureIds
	}
	toSerialize["journey_rum"] = o.JourneyRum
	toSerialize["name"] = o.Name
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMOperationRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId datadog.NullableUUID    `json:"application_id,omitempty"`
		Category      datadog.NullableString  `json:"category,omitempty"`
		Description   datadog.NullableString  `json:"description,omitempty"`
		DisplayName   *string                 `json:"display_name,omitempty"`
		FeatureIds    []string                `json:"feature_ids,omitempty"`
		JourneyRum    *RUMOperationJourneyRum `json:"journey_rum"`
		Name          *string                 `json:"name"`
		Tags          *[]string               `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JourneyRum == nil {
		return fmt.Errorf("required field journey_rum missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "category", "description", "display_name", "feature_ids", "journey_rum", "name", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	o.Category = all.Category
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	o.FeatureIds = all.FeatureIds
	if all.JourneyRum.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JourneyRum = *all.JourneyRum
	o.Name = *all.Name
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
