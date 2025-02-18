// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterMetadata Extra information about the exclusion filter.
type ApplicationSecurityExclusionFilterMetadata struct {
	// The creation date of the exclusion filter.
	AddedAt *time.Time `json:"added_at,omitempty"`
	// The handle of the user who created the exclusion filter.
	AddedBy *string `json:"added_by,omitempty"`
	// The last modification date of the exclusion filter.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The handle of the user who last modified the exclusion filter.
	ModifiedBy *string `json:"modified_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterMetadata instantiates a new ApplicationSecurityExclusionFilterMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterMetadata() *ApplicationSecurityExclusionFilterMetadata {
	this := ApplicationSecurityExclusionFilterMetadata{}
	return &this
}

// NewApplicationSecurityExclusionFilterMetadataWithDefaults instantiates a new ApplicationSecurityExclusionFilterMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterMetadataWithDefaults() *ApplicationSecurityExclusionFilterMetadata {
	this := ApplicationSecurityExclusionFilterMetadata{}
	return &this
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterMetadata) GetAddedAt() time.Time {
	if o == nil || o.AddedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || o.AddedAt == nil {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) HasAddedAt() bool {
	return o != nil && o.AddedAt != nil
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *ApplicationSecurityExclusionFilterMetadata) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterMetadata) GetAddedBy() string {
	if o == nil || o.AddedBy == nil {
		var ret string
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) GetAddedByOk() (*string, bool) {
	if o == nil || o.AddedBy == nil {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) HasAddedBy() bool {
	return o != nil && o.AddedBy != nil
}

// SetAddedBy gets a reference to the given string and assigns it to the AddedBy field.
func (o *ApplicationSecurityExclusionFilterMetadata) SetAddedBy(v string) {
	o.AddedBy = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterMetadata) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *ApplicationSecurityExclusionFilterMetadata) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterMetadata) GetModifiedBy() string {
	if o == nil || o.ModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) GetModifiedByOk() (*string, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterMetadata) HasModifiedBy() bool {
	return o != nil && o.ModifiedBy != nil
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *ApplicationSecurityExclusionFilterMetadata) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AddedAt != nil {
		if o.AddedAt.Nanosecond() == 0 {
			toSerialize["added_at"] = o.AddedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["added_at"] = o.AddedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.AddedBy != nil {
		toSerialize["added_by"] = o.AddedBy
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedBy != nil {
		toSerialize["modified_by"] = o.ModifiedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityExclusionFilterMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AddedAt    *time.Time `json:"added_at,omitempty"`
		AddedBy    *string    `json:"added_by,omitempty"`
		ModifiedAt *time.Time `json:"modified_at,omitempty"`
		ModifiedBy *string    `json:"modified_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"added_at", "added_by", "modified_at", "modified_by"})
	} else {
		return err
	}
	o.AddedAt = all.AddedAt
	o.AddedBy = all.AddedBy
	o.ModifiedAt = all.ModifiedAt
	o.ModifiedBy = all.ModifiedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
