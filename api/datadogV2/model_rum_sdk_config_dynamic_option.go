// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMSdkConfigDynamicOption A dynamic configuration option that extracts a value at runtime using a specified strategy.
type RUMSdkConfigDynamicOption struct {
	// The element attribute to read. Used when `strategy` is `dom`.
	Attribute *string `json:"attribute,omitempty"`
	// A serialized regex used as an extractor in dynamic options.
	Extractor *RUMSdkConfigSerializedRegex `json:"extractor,omitempty"`
	// The `localStorage` key to read. Required when `strategy` is `localStorage`.
	Key *string `json:"key,omitempty"`
	// The cookie name to read. Required when `strategy` is `cookie`.
	Name *string `json:"name,omitempty"`
	// The JavaScript path used to extract the value. Required when `strategy` is `js`.
	Path *string `json:"path,omitempty"`
	// The type identifier for a dynamic option. Always `dynamic`.
	RcSerializedType RUMSdkConfigDynamicOptionSerializedType `json:"rc_serialized_type"`
	// The CSS selector to read from the page. Required when `strategy` is `dom`.
	Selector *string `json:"selector,omitempty"`
	// The strategy used to extract the dynamic value.
	Strategy RUMSdkConfigDynamicOptionStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMSdkConfigDynamicOption instantiates a new RUMSdkConfigDynamicOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMSdkConfigDynamicOption(rcSerializedType RUMSdkConfigDynamicOptionSerializedType, strategy RUMSdkConfigDynamicOptionStrategy) *RUMSdkConfigDynamicOption {
	this := RUMSdkConfigDynamicOption{}
	this.RcSerializedType = rcSerializedType
	this.Strategy = strategy
	return &this
}

// NewRUMSdkConfigDynamicOptionWithDefaults instantiates a new RUMSdkConfigDynamicOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMSdkConfigDynamicOptionWithDefaults() *RUMSdkConfigDynamicOption {
	this := RUMSdkConfigDynamicOption{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *RUMSdkConfigDynamicOption) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *RUMSdkConfigDynamicOption) HasAttribute() bool {
	return o != nil && o.Attribute != nil
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *RUMSdkConfigDynamicOption) SetAttribute(v string) {
	o.Attribute = &v
}

// GetExtractor returns the Extractor field value if set, zero value otherwise.
func (o *RUMSdkConfigDynamicOption) GetExtractor() RUMSdkConfigSerializedRegex {
	if o == nil || o.Extractor == nil {
		var ret RUMSdkConfigSerializedRegex
		return ret
	}
	return *o.Extractor
}

// GetExtractorOk returns a tuple with the Extractor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetExtractorOk() (*RUMSdkConfigSerializedRegex, bool) {
	if o == nil || o.Extractor == nil {
		return nil, false
	}
	return o.Extractor, true
}

// HasExtractor returns a boolean if a field has been set.
func (o *RUMSdkConfigDynamicOption) HasExtractor() bool {
	return o != nil && o.Extractor != nil
}

// SetExtractor gets a reference to the given RUMSdkConfigSerializedRegex and assigns it to the Extractor field.
func (o *RUMSdkConfigDynamicOption) SetExtractor(v RUMSdkConfigSerializedRegex) {
	o.Extractor = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RUMSdkConfigDynamicOption) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RUMSdkConfigDynamicOption) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RUMSdkConfigDynamicOption) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RUMSdkConfigDynamicOption) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RUMSdkConfigDynamicOption) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RUMSdkConfigDynamicOption) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RUMSdkConfigDynamicOption) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RUMSdkConfigDynamicOption) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RUMSdkConfigDynamicOption) SetPath(v string) {
	o.Path = &v
}

// GetRcSerializedType returns the RcSerializedType field value.
func (o *RUMSdkConfigDynamicOption) GetRcSerializedType() RUMSdkConfigDynamicOptionSerializedType {
	if o == nil {
		var ret RUMSdkConfigDynamicOptionSerializedType
		return ret
	}
	return o.RcSerializedType
}

// GetRcSerializedTypeOk returns a tuple with the RcSerializedType field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetRcSerializedTypeOk() (*RUMSdkConfigDynamicOptionSerializedType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RcSerializedType, true
}

// SetRcSerializedType sets field value.
func (o *RUMSdkConfigDynamicOption) SetRcSerializedType(v RUMSdkConfigDynamicOptionSerializedType) {
	o.RcSerializedType = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *RUMSdkConfigDynamicOption) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *RUMSdkConfigDynamicOption) HasSelector() bool {
	return o != nil && o.Selector != nil
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *RUMSdkConfigDynamicOption) SetSelector(v string) {
	o.Selector = &v
}

// GetStrategy returns the Strategy field value.
func (o *RUMSdkConfigDynamicOption) GetStrategy() RUMSdkConfigDynamicOptionStrategy {
	if o == nil {
		var ret RUMSdkConfigDynamicOptionStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *RUMSdkConfigDynamicOption) GetStrategyOk() (*RUMSdkConfigDynamicOptionStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *RUMSdkConfigDynamicOption) SetStrategy(v RUMSdkConfigDynamicOptionStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMSdkConfigDynamicOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Extractor != nil {
		toSerialize["extractor"] = o.Extractor
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	toSerialize["rc_serialized_type"] = o.RcSerializedType
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMSdkConfigDynamicOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute        *string                                  `json:"attribute,omitempty"`
		Extractor        *RUMSdkConfigSerializedRegex             `json:"extractor,omitempty"`
		Key              *string                                  `json:"key,omitempty"`
		Name             *string                                  `json:"name,omitempty"`
		Path             *string                                  `json:"path,omitempty"`
		RcSerializedType *RUMSdkConfigDynamicOptionSerializedType `json:"rc_serialized_type"`
		Selector         *string                                  `json:"selector,omitempty"`
		Strategy         *RUMSdkConfigDynamicOptionStrategy       `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RcSerializedType == nil {
		return fmt.Errorf("required field rc_serialized_type missing")
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "extractor", "key", "name", "path", "rc_serialized_type", "selector", "strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = all.Attribute
	if all.Extractor != nil && all.Extractor.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Extractor = all.Extractor
	o.Key = all.Key
	o.Name = all.Name
	o.Path = all.Path
	if !all.RcSerializedType.IsValid() {
		hasInvalidField = true
	} else {
		o.RcSerializedType = *all.RcSerializedType
	}
	o.Selector = all.Selector
	if !all.Strategy.IsValid() {
		hasInvalidField = true
	} else {
		o.Strategy = *all.Strategy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
