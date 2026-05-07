// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostSettingDataAttributes Attributes for a cost setting.
type CostSettingDataAttributes struct {
	// The timestamp when the setting was created.
	CreatedAt string `json:"created_at"`
	// The UUID of the user who created the setting.
	CreatedBy string `json:"created_by"`
	// The setting data as a flexible key-value map.
	Data map[string]interface{} `json:"data,omitempty"`
	// A human-readable description of the setting.
	Description string `json:"description"`
	// The UUID of the user who last modified the setting.
	LastModifiedBy string `json:"last_modified_by"`
	// The name of the setting.
	SettingName string `json:"setting_name"`
	// The timestamp when the setting was last updated.
	UpdatedAt string `json:"updated_at"`
	// The version of the setting.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostSettingDataAttributes instantiates a new CostSettingDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostSettingDataAttributes(createdAt string, createdBy string, description string, lastModifiedBy string, settingName string, updatedAt string, version string) *CostSettingDataAttributes {
	this := CostSettingDataAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.LastModifiedBy = lastModifiedBy
	this.SettingName = settingName
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewCostSettingDataAttributesWithDefaults instantiates a new CostSettingDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostSettingDataAttributesWithDefaults() *CostSettingDataAttributes {
	this := CostSettingDataAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CostSettingDataAttributes) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CostSettingDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CostSettingDataAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CostSettingDataAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostSettingDataAttributes) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CostSettingDataAttributes) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CostSettingDataAttributes) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CostSettingDataAttributes) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetDescription returns the Description field value.
func (o *CostSettingDataAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CostSettingDataAttributes) SetDescription(v string) {
	o.Description = v
}

// GetLastModifiedBy returns the LastModifiedBy field value.
func (o *CostSettingDataAttributes) GetLastModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// SetLastModifiedBy sets field value.
func (o *CostSettingDataAttributes) SetLastModifiedBy(v string) {
	o.LastModifiedBy = v
}

// GetSettingName returns the SettingName field value.
func (o *CostSettingDataAttributes) GetSettingName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SettingName
}

// GetSettingNameOk returns a tuple with the SettingName field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetSettingNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettingName, true
}

// SetSettingName sets field value.
func (o *CostSettingDataAttributes) SetSettingName(v string) {
	o.SettingName = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *CostSettingDataAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *CostSettingDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value.
func (o *CostSettingDataAttributes) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CostSettingDataAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *CostSettingDataAttributes) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostSettingDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by"] = o.CreatedBy
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["description"] = o.Description
	toSerialize["last_modified_by"] = o.LastModifiedBy
	toSerialize["setting_name"] = o.SettingName
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostSettingDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *string                `json:"created_at"`
		CreatedBy      *string                `json:"created_by"`
		Data           map[string]interface{} `json:"data,omitempty"`
		Description    *string                `json:"description"`
		LastModifiedBy *string                `json:"last_modified_by"`
		SettingName    *string                `json:"setting_name"`
		UpdatedAt      *string                `json:"updated_at"`
		Version        *string                `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.LastModifiedBy == nil {
		return fmt.Errorf("required field last_modified_by missing")
	}
	if all.SettingName == nil {
		return fmt.Errorf("required field setting_name missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "data", "description", "last_modified_by", "setting_name", "updated_at", "version"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Data = all.Data
	o.Description = *all.Description
	o.LastModifiedBy = *all.LastModifiedBy
	o.SettingName = *all.SettingName
	o.UpdatedAt = *all.UpdatedAt
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
