// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentTemplateResponseAttributes Attributes of a segment template in a response.
type RumSegmentTemplateResponseAttributes struct {
	// The category of the template.
	Category string `json:"category"`
	// The creation timestamp in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// A description of the template.
	Description string `json:"description"`
	// The last modification timestamp in RFC 3339 format.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the template.
	Name string `json:"name"`
	// The template parameter definitions.
	Parameters map[string]RumSegmentTemplateParameterDef `json:"parameters"`
	// The status of a segment template.
	Status RumSegmentTemplateStatus `json:"status"`
	// The version number of the template.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentTemplateResponseAttributes instantiates a new RumSegmentTemplateResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentTemplateResponseAttributes(category string, createdAt time.Time, description string, modifiedAt time.Time, name string, parameters map[string]RumSegmentTemplateParameterDef, status RumSegmentTemplateStatus, version int64) *RumSegmentTemplateResponseAttributes {
	this := RumSegmentTemplateResponseAttributes{}
	this.Category = category
	this.CreatedAt = createdAt
	this.Description = description
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.Parameters = parameters
	this.Status = status
	this.Version = version
	return &this
}

// NewRumSegmentTemplateResponseAttributesWithDefaults instantiates a new RumSegmentTemplateResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentTemplateResponseAttributesWithDefaults() *RumSegmentTemplateResponseAttributes {
	this := RumSegmentTemplateResponseAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *RumSegmentTemplateResponseAttributes) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetCategory(v string) {
	o.Category = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RumSegmentTemplateResponseAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *RumSegmentTemplateResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *RumSegmentTemplateResponseAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *RumSegmentTemplateResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value.
func (o *RumSegmentTemplateResponseAttributes) GetParameters() map[string]RumSegmentTemplateParameterDef {
	if o == nil {
		var ret map[string]RumSegmentTemplateParameterDef
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetParametersOk() (*map[string]RumSegmentTemplateParameterDef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetParameters(v map[string]RumSegmentTemplateParameterDef) {
	o.Parameters = v
}

// GetStatus returns the Status field value.
func (o *RumSegmentTemplateResponseAttributes) GetStatus() RumSegmentTemplateStatus {
	if o == nil {
		var ret RumSegmentTemplateStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetStatusOk() (*RumSegmentTemplateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetStatus(v RumSegmentTemplateStatus) {
	o.Status = v
}

// GetVersion returns the Version field value.
func (o *RumSegmentTemplateResponseAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RumSegmentTemplateResponseAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *RumSegmentTemplateResponseAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentTemplateResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["parameters"] = o.Parameters
	toSerialize["status"] = o.Status
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentTemplateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category    *string                                    `json:"category"`
		CreatedAt   *time.Time                                 `json:"created_at"`
		Description *string                                    `json:"description"`
		ModifiedAt  *time.Time                                 `json:"modified_at"`
		Name        *string                                    `json:"name"`
		Parameters  *map[string]RumSegmentTemplateParameterDef `json:"parameters"`
		Status      *RumSegmentTemplateStatus                  `json:"status"`
		Version     *int64                                     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Parameters == nil {
		return fmt.Errorf("required field parameters missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "created_at", "description", "modified_at", "name", "parameters", "status", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = *all.Category
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.Parameters = *all.Parameters
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
