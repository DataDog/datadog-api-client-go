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

// FormDataAttributesResponse Attributes of a form.
type FormDataAttributesResponse struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// Configuration for the form's associated datastore.
	DatastoreConfig FormDatastoreConfig `json:"datastore_config"`
	// The description of the form.
	Description string `json:"description"`
	// Last modification timestamp.
	ModifiedAt time.Time `json:"modified_at"`
	// The name of the form.
	Name string `json:"name"`
	// The organization ID.
	OrgId int64 `json:"org_id"`
	// Publication information for the form.
	Publication *FormPublication `json:"publication,omitempty"`
	// The ID of the user who created the form.
	UserId int64 `json:"user_id"`
	// The UUID of the user who created the form.
	UserUuid uuid.UUID `json:"user_uuid"`
	// Version information for the form.
	Version *FormVersion `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormDataAttributesResponse instantiates a new FormDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormDataAttributesResponse(createdAt time.Time, datastoreConfig FormDatastoreConfig, description string, modifiedAt time.Time, name string, orgId int64, userId int64, userUuid uuid.UUID) *FormDataAttributesResponse {
	this := FormDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.DatastoreConfig = datastoreConfig
	this.Description = description
	this.ModifiedAt = modifiedAt
	this.Name = name
	this.OrgId = orgId
	this.UserId = userId
	this.UserUuid = userUuid
	return &this
}

// NewFormDataAttributesResponseWithDefaults instantiates a new FormDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormDataAttributesResponseWithDefaults() *FormDataAttributesResponse {
	this := FormDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FormDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FormDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDatastoreConfig returns the DatastoreConfig field value.
func (o *FormDataAttributesResponse) GetDatastoreConfig() FormDatastoreConfig {
	if o == nil {
		var ret FormDatastoreConfig
		return ret
	}
	return o.DatastoreConfig
}

// GetDatastoreConfigOk returns a tuple with the DatastoreConfig field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetDatastoreConfigOk() (*FormDatastoreConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatastoreConfig, true
}

// SetDatastoreConfig sets field value.
func (o *FormDataAttributesResponse) SetDatastoreConfig(v FormDatastoreConfig) {
	o.DatastoreConfig = v
}

// GetDescription returns the Description field value.
func (o *FormDataAttributesResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *FormDataAttributesResponse) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *FormDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *FormDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *FormDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *FormDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *FormDataAttributesResponse) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *FormDataAttributesResponse) SetOrgId(v int64) {
	o.OrgId = v
}

// GetPublication returns the Publication field value if set, zero value otherwise.
func (o *FormDataAttributesResponse) GetPublication() FormPublication {
	if o == nil || o.Publication == nil {
		var ret FormPublication
		return ret
	}
	return *o.Publication
}

// GetPublicationOk returns a tuple with the Publication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetPublicationOk() (*FormPublication, bool) {
	if o == nil || o.Publication == nil {
		return nil, false
	}
	return o.Publication, true
}

// HasPublication returns a boolean if a field has been set.
func (o *FormDataAttributesResponse) HasPublication() bool {
	return o != nil && o.Publication != nil
}

// SetPublication gets a reference to the given FormPublication and assigns it to the Publication field.
func (o *FormDataAttributesResponse) SetPublication(v FormPublication) {
	o.Publication = &v
}

// GetUserId returns the UserId field value.
func (o *FormDataAttributesResponse) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value.
func (o *FormDataAttributesResponse) SetUserId(v int64) {
	o.UserId = v
}

// GetUserUuid returns the UserUuid field value.
func (o *FormDataAttributesResponse) GetUserUuid() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetUserUuidOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value.
func (o *FormDataAttributesResponse) SetUserUuid(v uuid.UUID) {
	o.UserUuid = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FormDataAttributesResponse) GetVersion() FormVersion {
	if o == nil || o.Version == nil {
		var ret FormVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataAttributesResponse) GetVersionOk() (*FormVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FormDataAttributesResponse) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given FormVersion and assigns it to the Version field.
func (o *FormDataAttributesResponse) SetVersion(v FormVersion) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["datastore_config"] = o.DatastoreConfig
	toSerialize["description"] = o.Description
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["org_id"] = o.OrgId
	if o.Publication != nil {
		toSerialize["publication"] = o.Publication
	}
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
func (o *FormDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *time.Time           `json:"created_at"`
		DatastoreConfig *FormDatastoreConfig `json:"datastore_config"`
		Description     *string              `json:"description"`
		ModifiedAt      *time.Time           `json:"modified_at"`
		Name            *string              `json:"name"`
		OrgId           *int64               `json:"org_id"`
		Publication     *FormPublication     `json:"publication,omitempty"`
		UserId          *int64               `json:"user_id"`
		UserUuid        *uuid.UUID           `json:"user_uuid"`
		Version         *FormVersion         `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.DatastoreConfig == nil {
		return fmt.Errorf("required field datastore_config missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "datastore_config", "description", "modified_at", "name", "org_id", "publication", "user_id", "user_uuid", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.DatastoreConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DatastoreConfig = *all.DatastoreConfig
	o.Description = *all.Description
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	if all.Publication != nil && all.Publication.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Publication = all.Publication
	o.UserId = *all.UserId
	o.UserUuid = *all.UserUuid
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
