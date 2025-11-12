// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EmailAttributes Attributes for an on-call email.
type EmailAttributes struct {
	// Whether the email is currently active.
	Active *bool `json:"active,omitempty"`
	// Email address.
	Address *string `json:"address,omitempty"`
	// Optional display alias for the email.
	Alias *string `json:"alias,omitempty"`
	// Preferred content formats for notifications.
	Formats []EmailFormatType `json:"formats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEmailAttributes instantiates a new EmailAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEmailAttributes() *EmailAttributes {
	this := EmailAttributes{}
	return &this
}

// NewEmailAttributesWithDefaults instantiates a new EmailAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEmailAttributesWithDefaults() *EmailAttributes {
	this := EmailAttributes{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *EmailAttributes) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *EmailAttributes) HasActive() bool {
	return o != nil && o.Active != nil
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *EmailAttributes) SetActive(v bool) {
	o.Active = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *EmailAttributes) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *EmailAttributes) HasAddress() bool {
	return o != nil && o.Address != nil
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *EmailAttributes) SetAddress(v string) {
	o.Address = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *EmailAttributes) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *EmailAttributes) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *EmailAttributes) SetAlias(v string) {
	o.Alias = &v
}

// GetFormats returns the Formats field value if set, zero value otherwise.
func (o *EmailAttributes) GetFormats() []EmailFormatType {
	if o == nil || o.Formats == nil {
		var ret []EmailFormatType
		return ret
	}
	return o.Formats
}

// GetFormatsOk returns a tuple with the Formats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAttributes) GetFormatsOk() (*[]EmailFormatType, bool) {
	if o == nil || o.Formats == nil {
		return nil, false
	}
	return &o.Formats, true
}

// HasFormats returns a boolean if a field has been set.
func (o *EmailAttributes) HasFormats() bool {
	return o != nil && o.Formats != nil
}

// SetFormats gets a reference to the given []EmailFormatType and assigns it to the Formats field.
func (o *EmailAttributes) SetFormats(v []EmailFormatType) {
	o.Formats = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EmailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Formats != nil {
		toSerialize["formats"] = o.Formats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EmailAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Active  *bool             `json:"active,omitempty"`
		Address *string           `json:"address,omitempty"`
		Alias   *string           `json:"alias,omitempty"`
		Formats []EmailFormatType `json:"formats,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active", "address", "alias", "formats"})
	} else {
		return err
	}
	o.Active = all.Active
	o.Address = all.Address
	o.Alias = all.Alias
	o.Formats = all.Formats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
