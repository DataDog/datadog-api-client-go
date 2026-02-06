// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ClientToken Client token object. Client tokens (also known as public API keys) enable you to submit
// data from your browser or mobile applications to Datadog.
type ClientToken struct {
	// Creation timestamp of the client token.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Handle of the user who created the client token.
	CreatedByHandle *string `json:"created_by_handle,omitempty"`
	// UUID of the user who created the client token.
	CreatedByUuid *uuid.UUID `json:"created_by_uuid,omitempty"`
	// Timestamp when the client token was disabled.
	DisabledAt datadog.NullableTime `json:"disabled_at,omitempty"`
	// ID of the user who disabled the client token.
	DisabledBy *int64 `json:"disabled_by,omitempty"`
	// Handle of the user who disabled the client token.
	DisabledByHandle *string `json:"disabled_by_handle,omitempty"`
	// The hash value of the client token. This is the secret token value that should be
	// used in your browser or mobile application.
	Hash *string `json:"hash,omitempty"`
	// ID of the user who last modified the client token.
	ModifiedBy *int64 `json:"modified_by,omitempty"`
	// Name of the client token.
	Name *string `json:"name,omitempty"`
	// Organization ID associated with the client token.
	OrgId *int64 `json:"org_id,omitempty"`
	// List of allowed origin URLs for browser-based applications. Requests from URLs
	// not in this list will be rejected.
	OriginUrls []string `json:"origin_urls"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClientToken instantiates a new ClientToken object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClientToken(originUrls []string) *ClientToken {
	this := ClientToken{}
	this.OriginUrls = originUrls
	return &this
}

// NewClientTokenWithDefaults instantiates a new ClientToken object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClientTokenWithDefaults() *ClientToken {
	this := ClientToken{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ClientToken) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ClientToken) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ClientToken) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *ClientToken) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *ClientToken) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *ClientToken) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetCreatedByUuid returns the CreatedByUuid field value if set, zero value otherwise.
func (o *ClientToken) GetCreatedByUuid() uuid.UUID {
	if o == nil || o.CreatedByUuid == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.CreatedByUuid
}

// GetCreatedByUuidOk returns a tuple with the CreatedByUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetCreatedByUuidOk() (*uuid.UUID, bool) {
	if o == nil || o.CreatedByUuid == nil {
		return nil, false
	}
	return o.CreatedByUuid, true
}

// HasCreatedByUuid returns a boolean if a field has been set.
func (o *ClientToken) HasCreatedByUuid() bool {
	return o != nil && o.CreatedByUuid != nil
}

// SetCreatedByUuid gets a reference to the given uuid.UUID and assigns it to the CreatedByUuid field.
func (o *ClientToken) SetCreatedByUuid(v uuid.UUID) {
	o.CreatedByUuid = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientToken) GetDisabledAt() time.Time {
	if o == nil || o.DisabledAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DisabledAt.Get()
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClientToken) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisabledAt.Get(), o.DisabledAt.IsSet()
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *ClientToken) HasDisabledAt() bool {
	return o != nil && o.DisabledAt.IsSet()
}

// SetDisabledAt gets a reference to the given datadog.NullableTime and assigns it to the DisabledAt field.
func (o *ClientToken) SetDisabledAt(v time.Time) {
	o.DisabledAt.Set(&v)
}

// SetDisabledAtNil sets the value for DisabledAt to be an explicit nil.
func (o *ClientToken) SetDisabledAtNil() {
	o.DisabledAt.Set(nil)
}

// UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil.
func (o *ClientToken) UnsetDisabledAt() {
	o.DisabledAt.Unset()
}

// GetDisabledBy returns the DisabledBy field value if set, zero value otherwise.
func (o *ClientToken) GetDisabledBy() int64 {
	if o == nil || o.DisabledBy == nil {
		var ret int64
		return ret
	}
	return *o.DisabledBy
}

// GetDisabledByOk returns a tuple with the DisabledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetDisabledByOk() (*int64, bool) {
	if o == nil || o.DisabledBy == nil {
		return nil, false
	}
	return o.DisabledBy, true
}

// HasDisabledBy returns a boolean if a field has been set.
func (o *ClientToken) HasDisabledBy() bool {
	return o != nil && o.DisabledBy != nil
}

// SetDisabledBy gets a reference to the given int64 and assigns it to the DisabledBy field.
func (o *ClientToken) SetDisabledBy(v int64) {
	o.DisabledBy = &v
}

// GetDisabledByHandle returns the DisabledByHandle field value if set, zero value otherwise.
func (o *ClientToken) GetDisabledByHandle() string {
	if o == nil || o.DisabledByHandle == nil {
		var ret string
		return ret
	}
	return *o.DisabledByHandle
}

// GetDisabledByHandleOk returns a tuple with the DisabledByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetDisabledByHandleOk() (*string, bool) {
	if o == nil || o.DisabledByHandle == nil {
		return nil, false
	}
	return o.DisabledByHandle, true
}

// HasDisabledByHandle returns a boolean if a field has been set.
func (o *ClientToken) HasDisabledByHandle() bool {
	return o != nil && o.DisabledByHandle != nil
}

// SetDisabledByHandle gets a reference to the given string and assigns it to the DisabledByHandle field.
func (o *ClientToken) SetDisabledByHandle(v string) {
	o.DisabledByHandle = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ClientToken) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ClientToken) HasHash() bool {
	return o != nil && o.Hash != nil
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ClientToken) SetHash(v string) {
	o.Hash = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ClientToken) GetModifiedBy() int64 {
	if o == nil || o.ModifiedBy == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetModifiedByOk() (*int64, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ClientToken) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given int64 and assigns it to the ModifiedBy field.
func (o *ClientToken) SetModifiedBy(v int64) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientToken) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientToken) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientToken) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ClientToken) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientToken) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ClientToken) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *ClientToken) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetOriginUrls returns the OriginUrls field value.
func (o *ClientToken) GetOriginUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OriginUrls
}

// GetOriginUrlsOk returns a tuple with the OriginUrls field value
// and a boolean to check if the value has been set.
func (o *ClientToken) GetOriginUrlsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginUrls, true
}

// SetOriginUrls sets field value.
func (o *ClientToken) SetOriginUrls(v []string) {
	o.OriginUrls = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClientToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedByHandle != nil {
		toSerialize["created_by_handle"] = o.CreatedByHandle
	}
	if o.CreatedByUuid != nil {
		toSerialize["created_by_uuid"] = o.CreatedByUuid
	}
	if o.DisabledAt.IsSet() {
		toSerialize["disabled_at"] = o.DisabledAt.Get()
	}
	if o.DisabledBy != nil {
		toSerialize["disabled_by"] = o.DisabledBy
	}
	if o.DisabledByHandle != nil {
		toSerialize["disabled_by_handle"] = o.DisabledByHandle
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	toSerialize["origin_urls"] = o.OriginUrls

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClientToken) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time           `json:"created_at,omitempty"`
		CreatedByHandle  *string              `json:"created_by_handle,omitempty"`
		CreatedByUuid    *uuid.UUID           `json:"created_by_uuid,omitempty"`
		DisabledAt       datadog.NullableTime `json:"disabled_at,omitempty"`
		DisabledBy       *int64               `json:"disabled_by,omitempty"`
		DisabledByHandle *string              `json:"disabled_by_handle,omitempty"`
		Hash             *string              `json:"hash,omitempty"`
		ModifiedBy       *int64               `json:"modified_by,omitempty"`
		Name             *string              `json:"name,omitempty"`
		OrgId            *int64               `json:"org_id,omitempty"`
		OriginUrls       *[]string            `json:"origin_urls"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OriginUrls == nil {
		return fmt.Errorf("required field origin_urls missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by_handle", "created_by_uuid", "disabled_at", "disabled_by", "disabled_by_handle", "hash", "modified_by", "name", "org_id", "origin_urls"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedByHandle = all.CreatedByHandle
	o.CreatedByUuid = all.CreatedByUuid
	o.DisabledAt = all.DisabledAt
	o.DisabledBy = all.DisabledBy
	o.DisabledByHandle = all.DisabledByHandle
	o.Hash = all.Hash
	o.ModifiedBy = all.ModifiedBy
	o.Name = all.Name
	o.OrgId = all.OrgId
	o.OriginUrls = *all.OriginUrls

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
