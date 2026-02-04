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

// FormVersionDataAttributesResponse Attributes of a form version.
type FormVersionDataAttributesResponse struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The data definition for the form.
	DataDefinition interface{} `json:"data_definition"`
	// Signature of the form definition.
	DefinitionSignature string `json:"definition_signature"`
	// The entity tag for the version.
	Etag string `json:"etag"`
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
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormVersionDataAttributesResponse instantiates a new FormVersionDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormVersionDataAttributesResponse(createdAt time.Time, dataDefinition interface{}, definitionSignature string, etag string, modifiedAt time.Time, state FormVersionState, uiDefinition interface{}, userId int64, userUuid uuid.UUID, version int64) *FormVersionDataAttributesResponse {
	this := FormVersionDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.DataDefinition = dataDefinition
	this.DefinitionSignature = definitionSignature
	this.Etag = etag
	this.ModifiedAt = modifiedAt
	this.State = state
	this.UiDefinition = uiDefinition
	this.UserId = userId
	this.UserUuid = userUuid
	this.Version = version
	return &this
}

// NewFormVersionDataAttributesResponseWithDefaults instantiates a new FormVersionDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormVersionDataAttributesResponseWithDefaults() *FormVersionDataAttributesResponse {
	this := FormVersionDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormVersionDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormVersionDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDataDefinition returns the DataDefinition field value.
func (o *FormVersionDataAttributesResponse) GetDataDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DataDefinition
}

// GetDataDefinitionOk returns a tuple with the DataDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetDataDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDefinition, true
}

// SetDataDefinition sets field value.
func (o *FormVersionDataAttributesResponse) SetDataDefinition(v interface{}) {
	o.DataDefinition = v
}

// GetDefinitionSignature returns the DefinitionSignature field value.
func (o *FormVersionDataAttributesResponse) GetDefinitionSignature() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefinitionSignature
}

// GetDefinitionSignatureOk returns a tuple with the DefinitionSignature field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetDefinitionSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionSignature, true
}

// SetDefinitionSignature sets field value.
func (o *FormVersionDataAttributesResponse) SetDefinitionSignature(v string) {
	o.DefinitionSignature = v
}

// GetEtag returns the Etag field value.
func (o *FormVersionDataAttributesResponse) GetEtag() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Etag
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etag, true
}

// SetEtag sets field value.
func (o *FormVersionDataAttributesResponse) SetEtag(v string) {
	o.Etag = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormVersionDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormVersionDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetState returns the State field value.
func (o *FormVersionDataAttributesResponse) GetState() FormVersionState {
	if o == nil {
		var ret FormVersionState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetStateOk() (*FormVersionState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *FormVersionDataAttributesResponse) SetState(v FormVersionState) {
	o.State = v
}

// GetUiDefinition returns the UiDefinition field value.
func (o *FormVersionDataAttributesResponse) GetUiDefinition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiDefinition
}

// GetUiDefinitionOk returns a tuple with the UiDefinition field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetUiDefinitionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UiDefinition, true
}

// SetUiDefinition sets field value.
func (o *FormVersionDataAttributesResponse) SetUiDefinition(v interface{}) {
	o.UiDefinition = v
}

// GetUserId returns the UserId field value.
func (o *FormVersionDataAttributesResponse) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormVersionDataAttributesResponse) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormVersionDataAttributesResponse) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormVersionDataAttributesResponse) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// GetVersion returns the Version field value.
func (o *FormVersionDataAttributesResponse) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FormVersionDataAttributesResponse) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *FormVersionDataAttributesResponse) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormVersionDataAttributesResponse) MarshalJSON() ([]byte, error) {
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
	toSerialize["definition_signature"] = o.DefinitionSignature
	toSerialize["etag"] = o.Etag
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["state"] = o.State
	toSerialize["ui_definition"] = o.UiDefinition
	toSerialize["user_id"] = o.UserId
	toSerialize["user_uuid"] = o.UserUuid
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormVersionDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt           *time.Time        `json:"created_at"`
		DataDefinition      *interface{}      `json:"data_definition"`
		DefinitionSignature *string           `json:"definition_signature"`
		Etag                *string           `json:"etag"`
		ModifiedAt          *time.Time        `json:"modified_at"`
		State               *FormVersionState `json:"state"`
		UiDefinition        *interface{}      `json:"ui_definition"`
		UserId              *int64            `json:"user_id"`
		UserUuid            *uuid.UUID        `json:"user_uuid"`
		Version             *int64            `json:"version"`
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
	if all.DefinitionSignature == nil {
		return fmt.Errorf("required field definition_signature missing")
	}
	if all.Etag == nil {
		return fmt.Errorf("required field etag missing")
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
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "data_definition", "definition_signature", "etag", "modified_at", "state", "ui_definition", "user_id", "user_uuid", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.DataDefinition = *all.DataDefinition
	o.DefinitionSignature = *all.DefinitionSignature
	o.Etag = *all.Etag
	o.ModifiedAt = *all.ModifiedAt
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	o.UiDefinition = *all.UiDefinition
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
