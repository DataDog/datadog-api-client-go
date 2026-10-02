// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationSchemaSpecOptionV2 A single configuration option within an integration's configuration file.
type FleetIntegrationSchemaSpecOptionV2 struct {
	// Deprecation information for a configuration option. Currently carries no fields and is always emitted as an empty object or `null`.
	Deprecation map[string]interface{} `json:"deprecation"`
	// A human-readable description of the option.
	Description string `json:"description"`
	// The display order priority of the option relative to other options.
	DisplayPriority int64 `json:"display_priority"`
	// Whether the option is enabled by default.
	Enabled bool `json:"enabled"`
	// An example value for the option. Can be any JSON type. Absent from the response when not set.
	Example interface{} `json:"example,omitempty"`
	// Whether the option is hidden from the default configuration UI.
	Hidden bool `json:"hidden"`
	// Metadata tags associated with the option. Returned as an empty array when the option has no tags.
	MetadataTags []string `json:"metadata_tags"`
	// Whether the option accepts multiple values.
	Multiple bool `json:"multiple"`
	// Whether multiple instances of this option are defined in the configuration file.
	MultipleInstancesDefined bool `json:"multiple_instances_defined"`
	// The option name.
	Name string `json:"name"`
	// Nested options. Absent from the response when the option has no nested options.
	Options []FleetIntegrationSchemaSpecOptionV2 `json:"options,omitempty"`
	// A prefill value for the option. Can be any JSON type. Absent from the response when not set.
	Prefill interface{} `json:"prefill,omitempty"`
	// Whether the option is required.
	Required bool `json:"required"`
	// Whether the option is a secret that should be masked. Absent from the response when not set, distinct from being explicitly set to `false`.
	Secret *bool `json:"secret,omitempty"`
	// A JSON Schema-like specification for a configuration value.
	//
	// Object-typed values always include a `properties` array, even when empty.
	// Non-object-typed values never include `properties`. Throughout this schema,
	// an empty array is meaningfully different from an absent field.
	//
	// Three further JSON Schema keywords can appear directly on this object but are
	// not listed among its properties below to avoid clashing with this document's
	// own schema composition keywords: `enum` (an array of allowed values, present
	// only when there are enum constraints), `required` (an array of required
	// property names, present only when `type` is `object`), and `oneOf` (an array
	// of exclusive alternative value specifications this value can match, present
	// only when there are alternatives).
	Value *FleetIntegrationSchemaSpecValueV2 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationSchemaSpecOptionV2 instantiates a new FleetIntegrationSchemaSpecOptionV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationSchemaSpecOptionV2(deprecation map[string]interface{}, description string, displayPriority int64, enabled bool, hidden bool, metadataTags []string, multiple bool, multipleInstancesDefined bool, name string, required bool) *FleetIntegrationSchemaSpecOptionV2 {
	this := FleetIntegrationSchemaSpecOptionV2{}
	this.Deprecation = deprecation
	this.Description = description
	this.DisplayPriority = displayPriority
	this.Enabled = enabled
	this.Hidden = hidden
	this.MetadataTags = metadataTags
	this.Multiple = multiple
	this.MultipleInstancesDefined = multipleInstancesDefined
	this.Name = name
	this.Required = required
	return &this
}

// NewFleetIntegrationSchemaSpecOptionV2WithDefaults instantiates a new FleetIntegrationSchemaSpecOptionV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationSchemaSpecOptionV2WithDefaults() *FleetIntegrationSchemaSpecOptionV2 {
	this := FleetIntegrationSchemaSpecOptionV2{}
	return &this
}

// GetDeprecation returns the Deprecation field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *FleetIntegrationSchemaSpecOptionV2) GetDeprecation() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Deprecation
}

// GetDeprecationOk returns a tuple with the Deprecation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FleetIntegrationSchemaSpecOptionV2) GetDeprecationOk() (*map[string]interface{}, bool) {
	if o == nil || o.Deprecation == nil {
		return nil, false
	}
	return &o.Deprecation, true
}

// SetDeprecation sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetDeprecation(v map[string]interface{}) {
	o.Deprecation = v
}

// GetDescription returns the Description field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetDescription(v string) {
	o.Description = v
}

// GetDisplayPriority returns the DisplayPriority field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetDisplayPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DisplayPriority
}

// GetDisplayPriorityOk returns a tuple with the DisplayPriority field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetDisplayPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayPriority, true
}

// SetDisplayPriority sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetDisplayPriority(v int64) {
	o.DisplayPriority = v
}

// GetEnabled returns the Enabled field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecOptionV2) GetExample() interface{} {
	if o == nil || o.Example == nil {
		var ret interface{}
		return ret
	}
	return o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetExampleOk() (*interface{}, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return &o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) HasExample() bool {
	return o != nil && o.Example != nil
}

// SetExample gets a reference to the given interface{} and assigns it to the Example field.
func (o *FleetIntegrationSchemaSpecOptionV2) SetExample(v interface{}) {
	o.Example = v
}

// GetHidden returns the Hidden field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetHidden(v bool) {
	o.Hidden = v
}

// GetMetadataTags returns the MetadataTags field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetMetadataTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MetadataTags
}

// GetMetadataTagsOk returns a tuple with the MetadataTags field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetMetadataTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataTags, true
}

// SetMetadataTags sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetMetadataTags(v []string) {
	o.MetadataTags = v
}

// GetMultiple returns the Multiple field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetMultiple() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Multiple
}

// GetMultipleOk returns a tuple with the Multiple field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetMultipleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Multiple, true
}

// SetMultiple sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetMultiple(v bool) {
	o.Multiple = v
}

// GetMultipleInstancesDefined returns the MultipleInstancesDefined field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetMultipleInstancesDefined() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.MultipleInstancesDefined
}

// GetMultipleInstancesDefinedOk returns a tuple with the MultipleInstancesDefined field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetMultipleInstancesDefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultipleInstancesDefined, true
}

// SetMultipleInstancesDefined sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetMultipleInstancesDefined(v bool) {
	o.MultipleInstancesDefined = v
}

// GetName returns the Name field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecOptionV2) GetOptions() []FleetIntegrationSchemaSpecOptionV2 {
	if o == nil || o.Options == nil {
		var ret []FleetIntegrationSchemaSpecOptionV2
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetOptionsOk() (*[]FleetIntegrationSchemaSpecOptionV2, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given []FleetIntegrationSchemaSpecOptionV2 and assigns it to the Options field.
func (o *FleetIntegrationSchemaSpecOptionV2) SetOptions(v []FleetIntegrationSchemaSpecOptionV2) {
	o.Options = v
}

// GetPrefill returns the Prefill field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecOptionV2) GetPrefill() interface{} {
	if o == nil || o.Prefill == nil {
		var ret interface{}
		return ret
	}
	return o.Prefill
}

// GetPrefillOk returns a tuple with the Prefill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetPrefillOk() (*interface{}, bool) {
	if o == nil || o.Prefill == nil {
		return nil, false
	}
	return &o.Prefill, true
}

// HasPrefill returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) HasPrefill() bool {
	return o != nil && o.Prefill != nil
}

// SetPrefill gets a reference to the given interface{} and assigns it to the Prefill field.
func (o *FleetIntegrationSchemaSpecOptionV2) SetPrefill(v interface{}) {
	o.Prefill = v
}

// GetRequired returns the Required field value.
func (o *FleetIntegrationSchemaSpecOptionV2) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *FleetIntegrationSchemaSpecOptionV2) SetRequired(v bool) {
	o.Required = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecOptionV2) GetSecret() bool {
	if o == nil || o.Secret == nil {
		var ret bool
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetSecretOk() (*bool, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) HasSecret() bool {
	return o != nil && o.Secret != nil
}

// SetSecret gets a reference to the given bool and assigns it to the Secret field.
func (o *FleetIntegrationSchemaSpecOptionV2) SetSecret(v bool) {
	o.Secret = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaSpecOptionV2) GetValue() FleetIntegrationSchemaSpecValueV2 {
	if o == nil || o.Value == nil {
		var ret FleetIntegrationSchemaSpecValueV2
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) GetValueOk() (*FleetIntegrationSchemaSpecValueV2, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaSpecOptionV2) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given FleetIntegrationSchemaSpecValueV2 and assigns it to the Value field.
func (o *FleetIntegrationSchemaSpecOptionV2) SetValue(v FleetIntegrationSchemaSpecValueV2) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationSchemaSpecOptionV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deprecation != nil {
		toSerialize["deprecation"] = o.Deprecation
	}
	toSerialize["description"] = o.Description
	toSerialize["display_priority"] = o.DisplayPriority
	toSerialize["enabled"] = o.Enabled
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	toSerialize["hidden"] = o.Hidden
	toSerialize["metadata_tags"] = o.MetadataTags
	toSerialize["multiple"] = o.Multiple
	toSerialize["multiple_instances_defined"] = o.MultipleInstancesDefined
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Prefill != nil {
		toSerialize["prefill"] = o.Prefill
	}
	toSerialize["required"] = o.Required
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
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
func (o *FleetIntegrationSchemaSpecOptionV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deprecation              map[string]interface{}               `json:"deprecation"`
		Description              *string                              `json:"description"`
		DisplayPriority          *int64                               `json:"display_priority"`
		Enabled                  *bool                                `json:"enabled"`
		Example                  interface{}                          `json:"example,omitempty"`
		Hidden                   *bool                                `json:"hidden"`
		MetadataTags             *[]string                            `json:"metadata_tags"`
		Multiple                 *bool                                `json:"multiple"`
		MultipleInstancesDefined *bool                                `json:"multiple_instances_defined"`
		Name                     *string                              `json:"name"`
		Options                  []FleetIntegrationSchemaSpecOptionV2 `json:"options,omitempty"`
		Prefill                  interface{}                          `json:"prefill,omitempty"`
		Required                 *bool                                `json:"required"`
		Secret                   *bool                                `json:"secret,omitempty"`
		Value                    *FleetIntegrationSchemaSpecValueV2   `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Deprecation == nil {
		return fmt.Errorf("required field deprecation missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayPriority == nil {
		return fmt.Errorf("required field display_priority missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Hidden == nil {
		return fmt.Errorf("required field hidden missing")
	}
	if all.MetadataTags == nil {
		return fmt.Errorf("required field metadata_tags missing")
	}
	if all.Multiple == nil {
		return fmt.Errorf("required field multiple missing")
	}
	if all.MultipleInstancesDefined == nil {
		return fmt.Errorf("required field multiple_instances_defined missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deprecation", "description", "display_priority", "enabled", "example", "hidden", "metadata_tags", "multiple", "multiple_instances_defined", "name", "options", "prefill", "required", "secret", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Deprecation = all.Deprecation
	o.Description = *all.Description
	o.DisplayPriority = *all.DisplayPriority
	o.Enabled = *all.Enabled
	o.Example = all.Example
	o.Hidden = *all.Hidden
	o.MetadataTags = *all.MetadataTags
	o.Multiple = *all.Multiple
	o.MultipleInstancesDefined = *all.MultipleInstancesDefined
	o.Name = *all.Name
	o.Options = all.Options
	o.Prefill = all.Prefill
	o.Required = *all.Required
	o.Secret = all.Secret
	if all.Value != nil && all.Value.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
