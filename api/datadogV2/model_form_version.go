// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormVersion Version information for the form.
type FormVersion struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The data definition for the form.
	DataDefinition interface{} `json:"data_definition"`
	// Signature of the form definition.
	DefinitionSignature *string `json:"definition_signature,omitempty"`
	// The entity tag for the version.
	Etag *string `json:"etag,omitempty"`
	// The unique identifier of the form version.
	Id string `json:"id"`
	// Last modification timestamp.
	ModifiedAt time.Time `json:"modified_at"`
	// The state of the form version.
	State FormVersionState `json:"state"`
	// The UI definition for the form.
	UiDefinition interface{} `json:"ui_definition"`
	// The ID of the user who created the version.
	UserId int64 `json:"user_id"`
	// The UUID of the user who created the version.
	UserUuid uuid.UUID `json:"user_uuid"`
	// The version number.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormVersion instantiates a new FormVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormVersion(createdAt time.Time, dataDefinition interface{}, id string, modifiedAt time.Time, state FormVersionState, uiDefinition interface{}, userId int64, userUuid uuid.UUID) *FormVersion {
	this := FormVersion{}
	this.CreatedAt = createdAt
	this.DataDefinition = dataDefinition
	this.Id = id
	this.ModifiedAt = modifiedAt
	this.State = state
	this.UiDefinition = uiDefinition
	this.UserId = userId
	this.UserUuid = userUuid
	return &this
}

// NewFormVersionWithDefaults instantiates a new FormVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormVersionWithDefaults() *FormVersion {
	this := FormVersion{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormVersion) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDataDefinition returns the DataDefinition field value.
func (o *FormVersion) GetDataDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetDataDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *FormVersion) SetDataDefinition(v interface{}) {
	o.DataDefinition = v
}

// GetDefinitionSignature returns the DefinitionSignature field value if set, zero value otherwise.
func (o *FormVersion) GetDefinitionSignature() string {
	if o == nil || o.DefinitionSignature == nil {
		var ret string
		return ret
	}
	return *o.DefinitionSignature
}

// GetDefinitionSignatureOk returns a tuple with the DefinitionSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersion) GetDefinitionSignatureOk() (*string, bool) {
	if o == nil || o.DefinitionSignature == nil {
		return nil, false
	}
	return o.DefinitionSignature, true
}

// HasDefinitionSignature returns a boolean if a field has been set.
func (o *FormVersion) HasDefinitionSignature() bool {
	return o != nil && o.DefinitionSignature != nil
}

// SetDefinitionSignature gets a reference to the given string and assigns it to the DefinitionSignature field.
func (o *FormVersion) SetDefinitionSignature(v string) {
	o.DefinitionSignature = &v
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *FormVersion) GetEtag() string {
	if o == nil || o.Etag == nil {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersion) GetEtagOk() (*string, bool) {
	if o == nil || o.Etag == nil {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *FormVersion) HasEtag() bool {
	return o != nil && o.Etag != nil
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *FormVersion) SetEtag(v string) {
	o.Etag = &v
}

// GetId returns the Id field value.
func (o *FormVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *FormVersion) SetId(v string) {
	o.Id = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormVersion) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormVersion) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetState returns the State field value.
func (o *FormVersion) GetState() FormVersionState {
	if o == nil {
		var ret FormVersionState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetStateOk() (*FormVersionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *FormVersion) SetState(v FormVersionState) {
	o.State = v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *FormVersion) GetUiDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetUiDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *FormVersion) SetUiDefinition(v interface{}) {
	o.UiDefinition = v
}

// GetUserId returns the UserId field value.
func (o *FormVersion) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormVersion) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormVersion) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormVersion) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormVersion) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FormVersion) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormVersion) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FormVersion) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *FormVersion) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["data_definition"] = o.DataDefinition
	if o.DefinitionSignature != nil {
		toSerialize["definition_signature"] = o.DefinitionSignature
	}
	if o.Etag != nil {
		toSerialize["etag"] = o.Etag
	}
	toSerialize["id"] = o.Id
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["state"] = o.State
	toSerialize["ui_definition"] = o.UiDefinition
	toSerialize["user_id"] = o.UserId
	toSerialize["user_uuid"] = o.UserUuid
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt           *time.Time        `json:"created_at"`
		DataDefinition      *interface{}      `json:"data_definition"`
		DefinitionSignature *string           `json:"definition_signature,omitempty"`
		Etag                *string           `json:"etag,omitempty"`
		Id                  *string           `json:"id"`
		ModifiedAt          *time.Time        `json:"modified_at"`
		State               *FormVersionState `json:"state"`
		UiDefinition        *interface{}      `json:"ui_definition"`
		UserId              *int64            `json:"user_id"`
		UserUuid            *uuid.UUID        `json:"user_uuid"`
		Version             *int64            `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DataDefinition == nil {
		return fmt.Errorf("required field data_definition missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.UiDefinition == nil {
		return fmt.Errorf("required field ui_definition missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field user_id missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field user_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "data_definition", "definition_signature", "etag", "id", "modified_at", "state", "ui_definition", "user_id", "user_uuid", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.DataDefinition = *all.DataDefinition
	o.DefinitionSignature = all.DefinitionSignature
	o.Etag = all.Etag
	o.Id = *all.Id
	o.ModifiedAt = *all.ModifiedAt
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	o.UiDefinition = *all.UiDefinition
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
