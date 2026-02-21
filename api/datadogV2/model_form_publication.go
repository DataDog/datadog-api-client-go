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

// FormPublication Publication information for the form.
type FormPublication struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The form identifier.
	FormId uuid.UUID `json:"form_id"`
	// The version of the form that was published.
	FormVersion *int64 `json:"form_version,omitempty"`
	// The unique identifier of the publication.
	Id string `json:"id"`
	// Last modification timestamp.
	ModifiedAt time.Time `json:"modified_at"`
	// The organization ID.
	OrgId int64 `json:"org_id"`
	// The publication sequence number.
	PublishSeq *int64 `json:"publish_seq,omitempty"`
	// The ID of the user who published.
	UserId int64 `json:"user_id"`
	// The UUID of the user who published.
	UserUuid uuid.UUID `json:"user_uuid"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormPublication instantiates a new FormPublication object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormPublication(createdAt time.Time, formId uuid.UUID, id string, modifiedAt time.Time, orgId int64, userId int64, userUuid uuid.UUID) *FormPublication {
	this := FormPublication{}
	this.CreatedAt = createdAt
	this.FormId = formId
	this.Id = id
	this.ModifiedAt = modifiedAt
	this.OrgId = orgId
	this.UserId = userId
	this.UserUuid = userUuid
	return &this
}

// NewFormPublicationWithDefaults instantiates a new FormPublication object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormPublicationWithDefaults() *FormPublication {
	this := FormPublication{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormPublication) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormPublication) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFormId returns the FormId field value.
func (o *FormPublication) GetFormId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetFormIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value.
func (o *FormPublication) SetFormId(v uuid.UUID) {
	o.FormId = v
}

// GetFormVersion returns the FormVersion field value if set, zero value otherwise.
func (o *FormPublication) GetFormVersion() int64 {
	if o == nil || o.FormVersion == nil {
		var ret int64
		return ret
	}
	return *o.FormVersion
}

// GetFormVersionOk returns a tuple with the FormVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormPublication) GetFormVersionOk() (*int64, bool) {
	if o == nil || o.FormVersion == nil {
		return nil, false
	}
	return o.FormVersion, true
}

// HasFormVersion returns a boolean if a field has been set.
func (o *FormPublication) HasFormVersion() bool {
	return o != nil && o.FormVersion != nil
}

// SetFormVersion gets a reference to the given int64 and assigns it to the FormVersion field.
func (o *FormPublication) SetFormVersion(v int64) {
	o.FormVersion = &v
}

// GetId returns the Id field value.
func (o *FormPublication) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *FormPublication) SetId(v string) {
	o.Id = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormPublication) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormPublication) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetOrgId returns the OrgId field value.
func (o *FormPublication) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *FormPublication) SetOrgId(v int64) {
	o.OrgId = v
}

// GetPublishSeq returns the PublishSeq field value if set, zero value otherwise.
func (o *FormPublication) GetPublishSeq() int64 {
	if o == nil || o.PublishSeq == nil {
		var ret int64
		return ret
	}
	return *o.PublishSeq
}

// GetPublishSeqOk returns a tuple with the PublishSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormPublication) GetPublishSeqOk() (*int64, bool) {
	if o == nil || o.PublishSeq == nil {
		return nil, false
	}
	return o.PublishSeq, true
}

// HasPublishSeq returns a boolean if a field has been set.
func (o *FormPublication) HasPublishSeq() bool {
	return o != nil && o.PublishSeq != nil
}

// SetPublishSeq gets a reference to the given int64 and assigns it to the PublishSeq field.
func (o *FormPublication) SetPublishSeq(v int64) {
	o.PublishSeq = &v
}

// GetUserId returns the UserId field value.
func (o *FormPublication) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormPublication) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormPublication) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormPublication) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormPublication) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormPublication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["form_id"] = o.FormId
	if o.FormVersion != nil {
		toSerialize["form_version"] = o.FormVersion
	}
	toSerialize["id"] = o.Id
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["org_id"] = o.OrgId
	if o.PublishSeq != nil {
		toSerialize["publish_seq"] = o.PublishSeq
	}
	toSerialize["user_id"] = o.UserId
	toSerialize["user_uuid"] = o.UserUuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormPublication) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"created_at"`
		FormId      *uuid.UUID `json:"form_id"`
		FormVersion *int64     `json:"form_version,omitempty"`
		Id          *string    `json:"id"`
		ModifiedAt  *time.Time `json:"modified_at"`
		OrgId       *int64     `json:"org_id"`
		PublishSeq  *int64     `json:"publish_seq,omitempty"`
		UserId      *int64     `json:"user_id"`
		UserUuid    *uuid.UUID `json:"user_uuid"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.FormId == nil {
		return fmt.Errorf("required field form_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.UserId == nil {
		return fmt.Errorf("required field user_id missing")
	}
	if all.UserUuid == nil {
		return fmt.Errorf("required field user_uuid missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "form_id", "form_version", "id", "modified_at", "org_id", "publish_seq", "user_id", "user_uuid"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.FormId = *all.FormId
	o.FormVersion = all.FormVersion
	o.Id = *all.Id
	o.ModifiedAt = *all.ModifiedAt
	o.OrgId = *all.OrgId
	o.PublishSeq = all.PublishSeq
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
