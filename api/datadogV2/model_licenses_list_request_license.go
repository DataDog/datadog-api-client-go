// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LicensesListRequestLicense
type LicensesListRequestLicense struct {
	// The display name of the license
	DisplayName *string `json:"display_name,omitempty"`
	// The SPDX identifier of the license
	Identifier *string `json:"identifier,omitempty"`
	// The short name of the license
	ShortName *string `json:"short_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicensesListRequestLicense instantiates a new LicensesListRequestLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicensesListRequestLicense() *LicensesListRequestLicense {
	this := LicensesListRequestLicense{}
	return &this
}

// NewLicensesListRequestLicenseWithDefaults instantiates a new LicensesListRequestLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicensesListRequestLicenseWithDefaults() *LicensesListRequestLicense {
	this := LicensesListRequestLicense{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LicensesListRequestLicense) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensesListRequestLicense) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LicensesListRequestLicense) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LicensesListRequestLicense) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *LicensesListRequestLicense) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensesListRequestLicense) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *LicensesListRequestLicense) HasIdentifier() bool {
	return o != nil && o.Identifier != nil
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *LicensesListRequestLicense) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *LicensesListRequestLicense) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensesListRequestLicense) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *LicensesListRequestLicense) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *LicensesListRequestLicense) SetShortName(v string) {
	o.ShortName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LicensesListRequestLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicensesListRequestLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string `json:"display_name,omitempty"`
		Identifier  *string `json:"identifier,omitempty"`
		ShortName   *string `json:"short_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "identifier", "short_name"})
	} else {
		return err
	}
	o.DisplayName = all.DisplayName
	o.Identifier = all.Identifier
	o.ShortName = all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
