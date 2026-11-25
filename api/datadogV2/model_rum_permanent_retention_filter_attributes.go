// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMPermanentRetentionFilterAttributes The attributes of a permanent RUM retention filter.
type RUMPermanentRetentionFilterAttributes struct {
	// The configuration for cross-product retention filters.
	CrossProductSampling *RUMCrossProductSampling `json:"cross_product_sampling,omitempty"`
	// A description of what the filter retains.
	Description *string `json:"description,omitempty"`
	// Indicates which cross-product fields of a permanent RUM retention filter can be updated.
	Editability *RUMPermanentRetentionFilterEditability `json:"editability,omitempty"`
	// The display name of the permanent retention filter.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMPermanentRetentionFilterAttributes instantiates a new RUMPermanentRetentionFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMPermanentRetentionFilterAttributes() *RUMPermanentRetentionFilterAttributes {
	this := RUMPermanentRetentionFilterAttributes{}
	return &this
}

// NewRUMPermanentRetentionFilterAttributesWithDefaults instantiates a new RUMPermanentRetentionFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMPermanentRetentionFilterAttributesWithDefaults() *RUMPermanentRetentionFilterAttributes {
	this := RUMPermanentRetentionFilterAttributes{}
	return &this
}

// GetCrossProductSampling returns the CrossProductSampling field value if set, zero value otherwise.
func (o *RUMPermanentRetentionFilterAttributes) GetCrossProductSampling() RUMCrossProductSampling {
	if o == nil || o.CrossProductSampling == nil {
		var ret RUMCrossProductSampling
		return ret
	}
	return *o.CrossProductSampling
}

// GetCrossProductSamplingOk returns a tuple with the CrossProductSampling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMPermanentRetentionFilterAttributes) GetCrossProductSamplingOk() (*RUMCrossProductSampling, bool) {
	if o == nil || o.CrossProductSampling == nil {
		return nil, false
	}
	return o.CrossProductSampling, true
}

// HasCrossProductSampling returns a boolean if a field has been set.
func (o *RUMPermanentRetentionFilterAttributes) HasCrossProductSampling() bool {
	return o != nil && o.CrossProductSampling != nil
}

// SetCrossProductSampling gets a reference to the given RUMCrossProductSampling and assigns it to the CrossProductSampling field.
func (o *RUMPermanentRetentionFilterAttributes) SetCrossProductSampling(v RUMCrossProductSampling) {
	o.CrossProductSampling = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RUMPermanentRetentionFilterAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMPermanentRetentionFilterAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RUMPermanentRetentionFilterAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RUMPermanentRetentionFilterAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEditability returns the Editability field value if set, zero value otherwise.
func (o *RUMPermanentRetentionFilterAttributes) GetEditability() RUMPermanentRetentionFilterEditability {
	if o == nil || o.Editability == nil {
		var ret RUMPermanentRetentionFilterEditability
		return ret
	}
	return *o.Editability
}

// GetEditabilityOk returns a tuple with the Editability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMPermanentRetentionFilterAttributes) GetEditabilityOk() (*RUMPermanentRetentionFilterEditability, bool) {
	if o == nil || o.Editability == nil {
		return nil, false
	}
	return o.Editability, true
}

// HasEditability returns a boolean if a field has been set.
func (o *RUMPermanentRetentionFilterAttributes) HasEditability() bool {
	return o != nil && o.Editability != nil
}

// SetEditability gets a reference to the given RUMPermanentRetentionFilterEditability and assigns it to the Editability field.
func (o *RUMPermanentRetentionFilterAttributes) SetEditability(v RUMPermanentRetentionFilterEditability) {
	o.Editability = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RUMPermanentRetentionFilterAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMPermanentRetentionFilterAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RUMPermanentRetentionFilterAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RUMPermanentRetentionFilterAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMPermanentRetentionFilterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossProductSampling != nil {
		toSerialize["cross_product_sampling"] = o.CrossProductSampling
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Editability != nil {
		toSerialize["editability"] = o.Editability
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMPermanentRetentionFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossProductSampling *RUMCrossProductSampling                `json:"cross_product_sampling,omitempty"`
		Description          *string                                 `json:"description,omitempty"`
		Editability          *RUMPermanentRetentionFilterEditability `json:"editability,omitempty"`
		Name                 *string                                 `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cross_product_sampling", "description", "editability", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CrossProductSampling != nil && all.CrossProductSampling.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CrossProductSampling = all.CrossProductSampling
	o.Description = all.Description
	if all.Editability != nil && all.Editability.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Editability = all.Editability
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
