// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationSchemaFileSpecV2 A configuration file specification for an integration.
type FleetIntegrationSchemaFileSpecV2 struct {
	// The name of the example configuration file.
	ExampleName string `json:"example_name"`
	// The name of the configuration file.
	Name string `json:"name"`
	// The configuration options declared in the file.
	Options []FleetIntegrationSchemaSpecOptionV2 `json:"options"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationSchemaFileSpecV2 instantiates a new FleetIntegrationSchemaFileSpecV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationSchemaFileSpecV2(exampleName string, name string, options []FleetIntegrationSchemaSpecOptionV2) *FleetIntegrationSchemaFileSpecV2 {
	this := FleetIntegrationSchemaFileSpecV2{}
	this.ExampleName = exampleName
	this.Name = name
	this.Options = options
	return &this
}

// NewFleetIntegrationSchemaFileSpecV2WithDefaults instantiates a new FleetIntegrationSchemaFileSpecV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationSchemaFileSpecV2WithDefaults() *FleetIntegrationSchemaFileSpecV2 {
	this := FleetIntegrationSchemaFileSpecV2{}
	return &this
}

// GetExampleName returns the ExampleName field value.
func (o *FleetIntegrationSchemaFileSpecV2) GetExampleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExampleName
}

// GetExampleNameOk returns a tuple with the ExampleName field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaFileSpecV2) GetExampleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExampleName, true
}

// SetExampleName sets field value.
func (o *FleetIntegrationSchemaFileSpecV2) SetExampleName(v string) {
	o.ExampleName = v
}

// GetName returns the Name field value.
func (o *FleetIntegrationSchemaFileSpecV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaFileSpecV2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FleetIntegrationSchemaFileSpecV2) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value.
func (o *FleetIntegrationSchemaFileSpecV2) GetOptions() []FleetIntegrationSchemaSpecOptionV2 {
	if o == nil {
		var ret []FleetIntegrationSchemaSpecOptionV2
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaFileSpecV2) GetOptionsOk() (*[]FleetIntegrationSchemaSpecOptionV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value.
func (o *FleetIntegrationSchemaFileSpecV2) SetOptions(v []FleetIntegrationSchemaSpecOptionV2) {
	o.Options = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationSchemaFileSpecV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["example_name"] = o.ExampleName
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationSchemaFileSpecV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExampleName *string                               `json:"example_name"`
		Name        *string                               `json:"name"`
		Options     *[]FleetIntegrationSchemaSpecOptionV2 `json:"options"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExampleName == nil {
		return fmt.Errorf("required field example_name missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Options == nil {
		return fmt.Errorf("required field options missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"example_name", "name", "options"})
	} else {
		return err
	}
	o.ExampleName = *all.ExampleName
	o.Name = *all.Name
	o.Options = *all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
