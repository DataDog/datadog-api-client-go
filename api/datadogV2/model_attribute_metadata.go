// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttributeMetadata Attribute Metadata.
type AttributeMetadata struct {
	// Silence monitors for expected GCE instance shutdowns.
	Automute *bool `json:"automute,omitempty"`
	// Your client email.
	ClientEmail *string `json:"client_email,omitempty"`
	// Your host filters.
	HostFilters []string `json:"host_filters,omitempty"`
	// When enabled, Datadog performs configuration checks across your Google Cloud environment by continuously scanning every resource.
	IsCspmEnabled *bool `json:"is_cspm_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAttributeMetadata instantiates a new AttributeMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttributeMetadata() *AttributeMetadata {
	this := AttributeMetadata{}
	return &this
}

// NewAttributeMetadataWithDefaults instantiates a new AttributeMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttributeMetadataWithDefaults() *AttributeMetadata {
	this := AttributeMetadata{}
	return &this
}

// GetAutomute returns the Automute field value if set, zero value otherwise.
func (o *AttributeMetadata) GetAutomute() bool {
	if o == nil || o.Automute == nil {
		var ret bool
		return ret
	}
	return *o.Automute
}

// GetAutomuteOk returns a tuple with the Automute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMetadata) GetAutomuteOk() (*bool, bool) {
	if o == nil || o.Automute == nil {
		return nil, false
	}
	return o.Automute, true
}

// HasAutomute returns a boolean if a field has been set.
func (o *AttributeMetadata) HasAutomute() bool {
	return o != nil && o.Automute != nil
}

// SetAutomute gets a reference to the given bool and assigns it to the Automute field.
func (o *AttributeMetadata) SetAutomute(v bool) {
	o.Automute = &v
}

// GetClientEmail returns the ClientEmail field value if set, zero value otherwise.
func (o *AttributeMetadata) GetClientEmail() string {
	if o == nil || o.ClientEmail == nil {
		var ret string
		return ret
	}
	return *o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMetadata) GetClientEmailOk() (*string, bool) {
	if o == nil || o.ClientEmail == nil {
		return nil, false
	}
	return o.ClientEmail, true
}

// HasClientEmail returns a boolean if a field has been set.
func (o *AttributeMetadata) HasClientEmail() bool {
	return o != nil && o.ClientEmail != nil
}

// SetClientEmail gets a reference to the given string and assigns it to the ClientEmail field.
func (o *AttributeMetadata) SetClientEmail(v string) {
	o.ClientEmail = &v
}

// GetHostFilters returns the HostFilters field value if set, zero value otherwise.
func (o *AttributeMetadata) GetHostFilters() []string {
	if o == nil || o.HostFilters == nil {
		var ret []string
		return ret
	}
	return o.HostFilters
}

// GetHostFiltersOk returns a tuple with the HostFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMetadata) GetHostFiltersOk() (*[]string, bool) {
	if o == nil || o.HostFilters == nil {
		return nil, false
	}
	return &o.HostFilters, true
}

// HasHostFilters returns a boolean if a field has been set.
func (o *AttributeMetadata) HasHostFilters() bool {
	return o != nil && o.HostFilters != nil
}

// SetHostFilters gets a reference to the given []string and assigns it to the HostFilters field.
func (o *AttributeMetadata) SetHostFilters(v []string) {
	o.HostFilters = v
}

// GetIsCspmEnabled returns the IsCspmEnabled field value if set, zero value otherwise.
func (o *AttributeMetadata) GetIsCspmEnabled() bool {
	if o == nil || o.IsCspmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCspmEnabled
}

// GetIsCspmEnabledOk returns a tuple with the IsCspmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeMetadata) GetIsCspmEnabledOk() (*bool, bool) {
	if o == nil || o.IsCspmEnabled == nil {
		return nil, false
	}
	return o.IsCspmEnabled, true
}

// HasIsCspmEnabled returns a boolean if a field has been set.
func (o *AttributeMetadata) HasIsCspmEnabled() bool {
	return o != nil && o.IsCspmEnabled != nil
}

// SetIsCspmEnabled gets a reference to the given bool and assigns it to the IsCspmEnabled field.
func (o *AttributeMetadata) SetIsCspmEnabled(v bool) {
	o.IsCspmEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttributeMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Automute != nil {
		toSerialize["automute"] = o.Automute
	}
	if o.ClientEmail != nil {
		toSerialize["client_email"] = o.ClientEmail
	}
	if o.HostFilters != nil {
		toSerialize["host_filters"] = o.HostFilters
	}
	if o.IsCspmEnabled != nil {
		toSerialize["is_cspm_enabled"] = o.IsCspmEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AttributeMetadata) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Automute      *bool    `json:"automute,omitempty"`
		ClientEmail   *string  `json:"client_email,omitempty"`
		HostFilters   []string `json:"host_filters,omitempty"`
		IsCspmEnabled *bool    `json:"is_cspm_enabled,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"automute", "client_email", "host_filters", "is_cspm_enabled"})
	} else {
		return err
	}
	o.Automute = all.Automute
	o.ClientEmail = all.ClientEmail
	o.HostFilters = all.HostFilters
	o.IsCspmEnabled = all.IsCspmEnabled
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
