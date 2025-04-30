// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FullCustomFrameworkDataAttributes Full Framework Data Attributes.
type FullCustomFrameworkDataAttributes struct {
	// Framework Description
	Description string `json:"description"`
	// Framework Handle
	Handle string `json:"handle"`
	// Framework Icon URL
	IconUrl string `json:"icon_url"`
	// Framework Name
	Name string `json:"name"`
	// Framework Requirements
	Requirements []CustomFrameworkRequirement `json:"requirements"`
	// Framework Version
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFullCustomFrameworkDataAttributes instantiates a new FullCustomFrameworkDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFullCustomFrameworkDataAttributes(description string, handle string, iconUrl string, name string, requirements []CustomFrameworkRequirement, version string) *FullCustomFrameworkDataAttributes {
	this := FullCustomFrameworkDataAttributes{}
	this.Description = description
	this.Handle = handle
	this.IconUrl = iconUrl
	this.Name = name
	this.Requirements = requirements
	this.Version = version
	return &this
}

// NewFullCustomFrameworkDataAttributesWithDefaults instantiates a new FullCustomFrameworkDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFullCustomFrameworkDataAttributesWithDefaults() *FullCustomFrameworkDataAttributes {
	this := FullCustomFrameworkDataAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *FullCustomFrameworkDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FullCustomFrameworkDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetHandle returns the Handle field value.
func (o *FullCustomFrameworkDataAttributes) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value.
func (o *FullCustomFrameworkDataAttributes) SetHandle(v string) {
	o.Handle = v
}

// GetIconUrl returns the IconUrl field value.
func (o *FullCustomFrameworkDataAttributes) GetIconUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconUrl, true
}

// SetIconUrl sets field value.
func (o *FullCustomFrameworkDataAttributes) SetIconUrl(v string) {
	o.IconUrl = v
}

// GetName returns the Name field value.
func (o *FullCustomFrameworkDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FullCustomFrameworkDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRequirements returns the Requirements field value.
func (o *FullCustomFrameworkDataAttributes) GetRequirements() []CustomFrameworkRequirement {
	if o == nil {
		var ret []CustomFrameworkRequirement
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetRequirementsOk() (*[]CustomFrameworkRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value.
func (o *FullCustomFrameworkDataAttributes) SetRequirements(v []CustomFrameworkRequirement) {
	o.Requirements = v
}

// GetVersion returns the Version field value.
func (o *FullCustomFrameworkDataAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FullCustomFrameworkDataAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *FullCustomFrameworkDataAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FullCustomFrameworkDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["handle"] = o.Handle
	toSerialize["icon_url"] = o.IconUrl
	toSerialize["name"] = o.Name
	toSerialize["requirements"] = o.Requirements
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FullCustomFrameworkDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description  *string                       `json:"description"`
		Handle       *string                       `json:"handle"`
		IconUrl      *string                       `json:"icon_url"`
		Name         *string                       `json:"name"`
		Requirements *[]CustomFrameworkRequirement `json:"requirements"`
		Version      *string                       `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Handle == nil {
		return fmt.Errorf("required field handle missing")
	}
	if all.IconUrl == nil {
		return fmt.Errorf("required field icon_url missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Requirements == nil {
		return fmt.Errorf("required field requirements missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "handle", "icon_url", "name", "requirements", "version"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Handle = *all.Handle
	o.IconUrl = *all.IconUrl
	o.Name = *all.Name
	o.Requirements = *all.Requirements
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
