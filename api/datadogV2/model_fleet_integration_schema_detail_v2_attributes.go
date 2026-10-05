// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationSchemaDetailV2Attributes Attributes for a single integration's configuration schema.
type FleetIntegrationSchemaDetailV2Attributes struct {
	// The configuration file specifications for the integration. Always present, returned as an empty array when there are none.
	Files []FleetIntegrationSchemaFileSpecV2 `json:"files"`
	// The integration folder key. Absent from the response when empty.
	Folder *string `json:"folder,omitempty"`
	// The display name of the integration. Absent from the response when empty.
	Name *string `json:"name,omitempty"`
	// The integration version. Absent from the response when empty.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationSchemaDetailV2Attributes instantiates a new FleetIntegrationSchemaDetailV2Attributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationSchemaDetailV2Attributes(files []FleetIntegrationSchemaFileSpecV2) *FleetIntegrationSchemaDetailV2Attributes {
	this := FleetIntegrationSchemaDetailV2Attributes{}
	this.Files = files
	return &this
}

// NewFleetIntegrationSchemaDetailV2AttributesWithDefaults instantiates a new FleetIntegrationSchemaDetailV2Attributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationSchemaDetailV2AttributesWithDefaults() *FleetIntegrationSchemaDetailV2Attributes {
	this := FleetIntegrationSchemaDetailV2Attributes{}
	return &this
}

// GetFiles returns the Files field value.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetFiles() []FleetIntegrationSchemaFileSpecV2 {
	if o == nil {
		var ret []FleetIntegrationSchemaFileSpecV2
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetFilesOk() (*[]FleetIntegrationSchemaFileSpecV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *FleetIntegrationSchemaDetailV2Attributes) SetFiles(v []FleetIntegrationSchemaFileSpecV2) {
	o.Files = v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) HasFolder() bool {
	return o != nil && o.Folder != nil
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *FleetIntegrationSchemaDetailV2Attributes) SetFolder(v string) {
	o.Folder = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetIntegrationSchemaDetailV2Attributes) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FleetIntegrationSchemaDetailV2Attributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FleetIntegrationSchemaDetailV2Attributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationSchemaDetailV2Attributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["files"] = o.Files
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationSchemaDetailV2Attributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Files   *[]FleetIntegrationSchemaFileSpecV2 `json:"files"`
		Folder  *string                             `json:"folder,omitempty"`
		Name    *string                             `json:"name,omitempty"`
		Version *string                             `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"files", "folder", "name", "version"})
	} else {
		return err
	}
	o.Files = *all.Files
	o.Folder = all.Folder
	o.Name = all.Name
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
