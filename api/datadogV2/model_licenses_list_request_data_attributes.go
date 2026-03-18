// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LicensesListRequestDataAttributes
type LicensesListRequestDataAttributes struct {
	// List of available licenses
	Licenses []LicensesListRequestLicense `json:"licenses,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicensesListRequestDataAttributes instantiates a new LicensesListRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicensesListRequestDataAttributes() *LicensesListRequestDataAttributes {
	this := LicensesListRequestDataAttributes{}
	return &this
}

// NewLicensesListRequestDataAttributesWithDefaults instantiates a new LicensesListRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicensesListRequestDataAttributesWithDefaults() *LicensesListRequestDataAttributes {
	this := LicensesListRequestDataAttributes{}
	return &this
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *LicensesListRequestDataAttributes) GetLicenses() []LicensesListRequestLicense {
	if o == nil || o.Licenses == nil {
		var ret []LicensesListRequestLicense
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicensesListRequestDataAttributes) GetLicensesOk() (*[]LicensesListRequestLicense, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return &o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *LicensesListRequestDataAttributes) HasLicenses() bool {
	return o != nil && o.Licenses != nil
}

// SetLicenses gets a reference to the given []LicensesListRequestLicense and assigns it to the Licenses field.
func (o *LicensesListRequestDataAttributes) SetLicenses(v []LicensesListRequestLicense) {
	o.Licenses = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LicensesListRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicensesListRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Licenses []LicensesListRequestLicense `json:"licenses,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"licenses"})
	} else {
		return err
	}
	o.Licenses = all.Licenses

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
