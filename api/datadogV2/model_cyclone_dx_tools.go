// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXTools Tools used to generate the BOM.
type CycloneDXTools struct {
	// List of tool components. Only one tool is supported.
	Components []CycloneDXToolComponent `json:"components"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXTools instantiates a new CycloneDXTools object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXTools(components []CycloneDXToolComponent) *CycloneDXTools {
	this := CycloneDXTools{}
	this.Components = components
	return &this
}

// NewCycloneDXToolsWithDefaults instantiates a new CycloneDXTools object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXToolsWithDefaults() *CycloneDXTools {
	this := CycloneDXTools{}
	return &this
}

// GetComponents returns the Components field value.
func (o *CycloneDXTools) GetComponents() []CycloneDXToolComponent {
	if o == nil {
		var ret []CycloneDXToolComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *CycloneDXTools) GetComponentsOk() (*[]CycloneDXToolComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *CycloneDXTools) SetComponents(v []CycloneDXToolComponent) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXTools) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["components"] = o.Components

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXTools) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components *[]CycloneDXToolComponent `json:"components"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components"})
	} else {
		return err
	}
	o.Components = *all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
