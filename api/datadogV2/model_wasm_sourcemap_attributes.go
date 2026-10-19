// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WASMSourcemapAttributes Attributes of a WebAssembly symbol file.
type WASMSourcemapAttributes struct {
	// The build identifier.
	BuildId *string `json:"build_id,omitempty"`
	// The timestamp when the symbol file was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The file name of the symbol file.
	FileName *string `json:"file_name,omitempty"`
	// The type of source map. Always `wasm` for WebAssembly symbol files.
	Mapkind WASMSourcemapMapKind `json:"mapkind"`
	// The size of the symbol file in bytes.
	Size *int64 `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWASMSourcemapAttributes instantiates a new WASMSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWASMSourcemapAttributes(mapkind WASMSourcemapMapKind) *WASMSourcemapAttributes {
	this := WASMSourcemapAttributes{}
	this.Mapkind = mapkind
	return &this
}

// NewWASMSourcemapAttributesWithDefaults instantiates a new WASMSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWASMSourcemapAttributesWithDefaults() *WASMSourcemapAttributes {
	this := WASMSourcemapAttributes{}
	return &this
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *WASMSourcemapAttributes) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WASMSourcemapAttributes) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *WASMSourcemapAttributes) HasBuildId() bool {
	return o != nil && o.BuildId != nil
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *WASMSourcemapAttributes) SetBuildId(v string) {
	o.BuildId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WASMSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WASMSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WASMSourcemapAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WASMSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *WASMSourcemapAttributes) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WASMSourcemapAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *WASMSourcemapAttributes) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *WASMSourcemapAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetMapkind returns the Mapkind field value.
func (o *WASMSourcemapAttributes) GetMapkind() WASMSourcemapMapKind {
	if o == nil {
		var ret WASMSourcemapMapKind
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *WASMSourcemapAttributes) GetMapkindOk() (*WASMSourcemapMapKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *WASMSourcemapAttributes) SetMapkind(v WASMSourcemapMapKind) {
	o.Mapkind = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *WASMSourcemapAttributes) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WASMSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *WASMSourcemapAttributes) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *WASMSourcemapAttributes) SetSize(v int64) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WASMSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BuildId != nil {
		toSerialize["build_id"] = o.BuildId
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	toSerialize["mapkind"] = o.Mapkind
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WASMSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BuildId   *string               `json:"build_id,omitempty"`
		CreatedAt *time.Time            `json:"created_at,omitempty"`
		FileName  *string               `json:"file_name,omitempty"`
		Mapkind   *WASMSourcemapMapKind `json:"mapkind"`
		Size      *int64                `json:"size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mapkind == nil {
		return fmt.Errorf("required field mapkind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"build_id", "created_at", "file_name", "mapkind", "size"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BuildId = all.BuildId
	o.CreatedAt = all.CreatedAt
	o.FileName = all.FileName
	if !all.Mapkind.IsValid() {
		hasInvalidField = true
	} else {
		o.Mapkind = *all.Mapkind
	}
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
