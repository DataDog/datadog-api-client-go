// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentProtocolRequestDataAttributes Create a draft protocol with the organization default subject type.
type ExperimentsCreateExperimentProtocolRequestDataAttributes struct {
	// Optional text that explains the protocol.
	Description datadog.NullableString `json:"description,omitempty"`
	// Metadata for imported resources. Must encode to at most 4096 bytes. Null clears it; omission preserves it on update.
	MigrationMetadata map[string]interface{} `json:"migration_metadata,omitempty"`
	// Protocol name. Leading and trailing spaces are removed. The trimmed name must contain 1 to 255 characters.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentProtocolRequestDataAttributes instantiates a new ExperimentsCreateExperimentProtocolRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentProtocolRequestDataAttributes(name string) *ExperimentsCreateExperimentProtocolRequestDataAttributes {
	this := ExperimentsCreateExperimentProtocolRequestDataAttributes{}
	this.Name = name
	return &this
}

// NewExperimentsCreateExperimentProtocolRequestDataAttributesWithDefaults instantiates a new ExperimentsCreateExperimentProtocolRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentProtocolRequestDataAttributesWithDefaults() *ExperimentsCreateExperimentProtocolRequestDataAttributes {
	this := ExperimentsCreateExperimentProtocolRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetMigrationMetadata returns the MigrationMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) GetMigrationMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MigrationMetadata
}

// GetMigrationMetadataOk returns a tuple with the MigrationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) GetMigrationMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.MigrationMetadata == nil {
		return nil, false
	}
	return &o.MigrationMetadata, true
}

// HasMigrationMetadata returns a boolean if a field has been set.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) HasMigrationMetadata() bool {
	return o != nil && o.MigrationMetadata != nil
}

// SetMigrationMetadata gets a reference to the given map[string]interface{} and assigns it to the MigrationMetadata field.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) SetMigrationMetadata(v map[string]interface{}) {
	o.MigrationMetadata = v
}

// GetName returns the Name field value.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentProtocolRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.MigrationMetadata != nil {
		toSerialize["migration_metadata"] = o.MigrationMetadata
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentProtocolRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       datadog.NullableString `json:"description,omitempty"`
		MigrationMetadata map[string]interface{} `json:"migration_metadata,omitempty"`
		Name              *string                `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "migration_metadata", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.MigrationMetadata = all.MigrationMetadata
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
