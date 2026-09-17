// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateFeatureFlagAttributes Attributes for creating a new feature flag.
type CreateFeatureFlagAttributes struct {
	// The key of the default variant.
	DefaultVariantKey datadog.NullableString `json:"default_variant_key,omitempty"`
	// The description of the feature flag.
	Description *string `json:"description,omitempty"`
	// The distribution channel for the feature flag.
	DistributionChannel *FeatureFlagDistributionChannel `json:"distribution_channel,omitempty"`
	// JSON schema for validation when value_type is JSON.
	JsonSchema datadog.NullableString `json:"json_schema,omitempty"`
	// The unique key of the feature flag.
	Key string `json:"key"`
	// The name of the feature flag.
	Name string `json:"name"`
	// Indicates whether this feature flag requires approval for changes.
	RequireApproval *bool `json:"require_approval,omitempty"`
	// The staleness status for the feature flag at creation.
	StalenessStatus *CreateFeatureFlagStalenessStatus `json:"staleness_status,omitempty"`
	// Tags associated with the feature flag.
	Tags []string `json:"tags,omitempty"`
	// The type of values for the feature flag variants.
	ValueType ValueType `json:"value_type"`
	// The variants of the feature flag.
	Variants []CreateVariant `json:"variants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateFeatureFlagAttributes instantiates a new CreateFeatureFlagAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateFeatureFlagAttributes(key string, name string, valueType ValueType, variants []CreateVariant) *CreateFeatureFlagAttributes {
	this := CreateFeatureFlagAttributes{}
	this.Key = key
	this.Name = name
	this.ValueType = valueType
	this.Variants = variants
	return &this
}

// NewCreateFeatureFlagAttributesWithDefaults instantiates a new CreateFeatureFlagAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateFeatureFlagAttributesWithDefaults() *CreateFeatureFlagAttributes {
	this := CreateFeatureFlagAttributes{}
	return &this
}

// GetDefaultVariantKey returns the DefaultVariantKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureFlagAttributes) GetDefaultVariantKey() string {
	if o == nil || o.DefaultVariantKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultVariantKey.Get()
}

// GetDefaultVariantKeyOk returns a tuple with the DefaultVariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CreateFeatureFlagAttributes) GetDefaultVariantKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultVariantKey.Get(), o.DefaultVariantKey.IsSet()
}

// HasDefaultVariantKey returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasDefaultVariantKey() bool {
	return o != nil && o.DefaultVariantKey.IsSet()
}

// SetDefaultVariantKey gets a reference to the given datadog.NullableString and assigns it to the DefaultVariantKey field.
func (o *CreateFeatureFlagAttributes) SetDefaultVariantKey(v string) {
	o.DefaultVariantKey.Set(&v)
}

// SetDefaultVariantKeyNil sets the value for DefaultVariantKey to be an explicit nil.
func (o *CreateFeatureFlagAttributes) SetDefaultVariantKeyNil() {
	o.DefaultVariantKey.Set(nil)
}

// UnsetDefaultVariantKey ensures that no value is present for DefaultVariantKey, not even an explicit nil.
func (o *CreateFeatureFlagAttributes) UnsetDefaultVariantKey() {
	o.DefaultVariantKey.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFeatureFlagAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFeatureFlagAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDistributionChannel returns the DistributionChannel field value if set, zero value otherwise.
func (o *CreateFeatureFlagAttributes) GetDistributionChannel() FeatureFlagDistributionChannel {
	if o == nil || o.DistributionChannel == nil {
		var ret FeatureFlagDistributionChannel
		return ret
	}
	return *o.DistributionChannel
}

// GetDistributionChannelOk returns a tuple with the DistributionChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetDistributionChannelOk() (*FeatureFlagDistributionChannel, bool) {
	if o == nil || o.DistributionChannel == nil {
		return nil, false
	}
	return o.DistributionChannel, true
}

// HasDistributionChannel returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasDistributionChannel() bool {
	return o != nil && o.DistributionChannel != nil
}

// SetDistributionChannel gets a reference to the given FeatureFlagDistributionChannel and assigns it to the DistributionChannel field.
func (o *CreateFeatureFlagAttributes) SetDistributionChannel(v FeatureFlagDistributionChannel) {
	o.DistributionChannel = &v
}

// GetJsonSchema returns the JsonSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureFlagAttributes) GetJsonSchema() string {
	if o == nil || o.JsonSchema.Get() == nil {
		var ret string
		return ret
	}
	return *o.JsonSchema.Get()
}

// GetJsonSchemaOk returns a tuple with the JsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CreateFeatureFlagAttributes) GetJsonSchemaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JsonSchema.Get(), o.JsonSchema.IsSet()
}

// HasJsonSchema returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasJsonSchema() bool {
	return o != nil && o.JsonSchema.IsSet()
}

// SetJsonSchema gets a reference to the given datadog.NullableString and assigns it to the JsonSchema field.
func (o *CreateFeatureFlagAttributes) SetJsonSchema(v string) {
	o.JsonSchema.Set(&v)
}

// SetJsonSchemaNil sets the value for JsonSchema to be an explicit nil.
func (o *CreateFeatureFlagAttributes) SetJsonSchemaNil() {
	o.JsonSchema.Set(nil)
}

// UnsetJsonSchema ensures that no value is present for JsonSchema, not even an explicit nil.
func (o *CreateFeatureFlagAttributes) UnsetJsonSchema() {
	o.JsonSchema.Unset()
}

// GetKey returns the Key field value.
func (o *CreateFeatureFlagAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *CreateFeatureFlagAttributes) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value.
func (o *CreateFeatureFlagAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateFeatureFlagAttributes) SetName(v string) {
	o.Name = v
}

// GetRequireApproval returns the RequireApproval field value if set, zero value otherwise.
func (o *CreateFeatureFlagAttributes) GetRequireApproval() bool {
	if o == nil || o.RequireApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireApproval
}

// GetRequireApprovalOk returns a tuple with the RequireApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetRequireApprovalOk() (*bool, bool) {
	if o == nil || o.RequireApproval == nil {
		return nil, false
	}
	return o.RequireApproval, true
}

// HasRequireApproval returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasRequireApproval() bool {
	return o != nil && o.RequireApproval != nil
}

// SetRequireApproval gets a reference to the given bool and assigns it to the RequireApproval field.
func (o *CreateFeatureFlagAttributes) SetRequireApproval(v bool) {
	o.RequireApproval = &v
}

// GetStalenessStatus returns the StalenessStatus field value if set, zero value otherwise.
func (o *CreateFeatureFlagAttributes) GetStalenessStatus() CreateFeatureFlagStalenessStatus {
	if o == nil || o.StalenessStatus == nil {
		var ret CreateFeatureFlagStalenessStatus
		return ret
	}
	return *o.StalenessStatus
}

// GetStalenessStatusOk returns a tuple with the StalenessStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetStalenessStatusOk() (*CreateFeatureFlagStalenessStatus, bool) {
	if o == nil || o.StalenessStatus == nil {
		return nil, false
	}
	return o.StalenessStatus, true
}

// HasStalenessStatus returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasStalenessStatus() bool {
	return o != nil && o.StalenessStatus != nil
}

// SetStalenessStatus gets a reference to the given CreateFeatureFlagStalenessStatus and assigns it to the StalenessStatus field.
func (o *CreateFeatureFlagAttributes) SetStalenessStatus(v CreateFeatureFlagStalenessStatus) {
	o.StalenessStatus = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateFeatureFlagAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateFeatureFlagAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateFeatureFlagAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetValueType returns the ValueType field value.
func (o *CreateFeatureFlagAttributes) GetValueType() ValueType {
	if o == nil {
		var ret ValueType
		return ret
	}
	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetValueTypeOk() (*ValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value.
func (o *CreateFeatureFlagAttributes) SetValueType(v ValueType) {
	o.ValueType = v
}

// GetVariants returns the Variants field value.
func (o *CreateFeatureFlagAttributes) GetVariants() []CreateVariant {
	if o == nil {
		var ret []CreateVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureFlagAttributes) GetVariantsOk() (*[]CreateVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variants, true
}

// SetVariants sets field value.
func (o *CreateFeatureFlagAttributes) SetVariants(v []CreateVariant) {
	o.Variants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateFeatureFlagAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultVariantKey.IsSet() {
		toSerialize["default_variant_key"] = o.DefaultVariantKey.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DistributionChannel != nil {
		toSerialize["distribution_channel"] = o.DistributionChannel
	}
	if o.JsonSchema.IsSet() {
		toSerialize["json_schema"] = o.JsonSchema.Get()
	}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if o.RequireApproval != nil {
		toSerialize["require_approval"] = o.RequireApproval
	}
	if o.StalenessStatus != nil {
		toSerialize["staleness_status"] = o.StalenessStatus
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["value_type"] = o.ValueType
	toSerialize["variants"] = o.Variants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateFeatureFlagAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultVariantKey   datadog.NullableString            `json:"default_variant_key,omitempty"`
		Description         *string                           `json:"description,omitempty"`
		DistributionChannel *FeatureFlagDistributionChannel   `json:"distribution_channel,omitempty"`
		JsonSchema          datadog.NullableString            `json:"json_schema,omitempty"`
		Key                 *string                           `json:"key"`
		Name                *string                           `json:"name"`
		RequireApproval     *bool                             `json:"require_approval,omitempty"`
		StalenessStatus     *CreateFeatureFlagStalenessStatus `json:"staleness_status,omitempty"`
		Tags                []string                          `json:"tags,omitempty"`
		ValueType           *ValueType                        `json:"value_type"`
		Variants            *[]CreateVariant                  `json:"variants"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ValueType == nil {
		return fmt.Errorf("required field value_type missing")
	}
	if all.Variants == nil {
		return fmt.Errorf("required field variants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_variant_key", "description", "distribution_channel", "json_schema", "key", "name", "require_approval", "staleness_status", "tags", "value_type", "variants"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DefaultVariantKey = all.DefaultVariantKey
	o.Description = all.Description
	if all.DistributionChannel != nil && !all.DistributionChannel.IsValid() {
		hasInvalidField = true
	} else {
		o.DistributionChannel = all.DistributionChannel
	}
	o.JsonSchema = all.JsonSchema
	o.Key = *all.Key
	o.Name = *all.Name
	o.RequireApproval = all.RequireApproval
	if all.StalenessStatus != nil && !all.StalenessStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.StalenessStatus = all.StalenessStatus
	}
	o.Tags = all.Tags
	if !all.ValueType.IsValid() {
		hasInvalidField = true
	} else {
		o.ValueType = *all.ValueType
	}
	o.Variants = *all.Variants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
