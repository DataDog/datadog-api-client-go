// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListFormVersionsDataAttributes The attributes for a list of form versions.
type ListFormVersionsDataAttributes struct {
	// The list of versions for the form.
	Versions []FormVersionAttributes `json:"versions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListFormVersionsDataAttributes instantiates a new ListFormVersionsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListFormVersionsDataAttributes(versions []FormVersionAttributes) *ListFormVersionsDataAttributes {
	this := ListFormVersionsDataAttributes{}
	this.Versions = versions
	return &this
}

// NewListFormVersionsDataAttributesWithDefaults instantiates a new ListFormVersionsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListFormVersionsDataAttributesWithDefaults() *ListFormVersionsDataAttributes {
	this := ListFormVersionsDataAttributes{}
	return &this
}

// GetVersions returns the Versions field value.
func (o *ListFormVersionsDataAttributes) GetVersions() []FormVersionAttributes {
	if o == nil {
		var ret []FormVersionAttributes
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ListFormVersionsDataAttributes) GetVersionsOk() (*[]FormVersionAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versions, true
}

// SetVersions sets field value.
func (o *ListFormVersionsDataAttributes) SetVersions(v []FormVersionAttributes) {
	o.Versions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListFormVersionsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["versions"] = o.Versions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListFormVersionsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Versions *[]FormVersionAttributes `json:"versions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Versions == nil {
		return fmt.Errorf("required field versions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"versions"})
	} else {
		return err
	}
	o.Versions = *all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
