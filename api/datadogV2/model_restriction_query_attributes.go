// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryAttributes Attributes of the restriction query.
type RestrictionQueryAttributes struct {
	// Creation time of the restriction query.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Time of last restriction query modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The query that defines the restriction. Only the content matching the query can be returned.
	RestrictionQuery *string `json:"restriction_query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionQueryAttributes instantiates a new RestrictionQueryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionQueryAttributes() *RestrictionQueryAttributes {
	this := RestrictionQueryAttributes{}
	return &this
}

// NewRestrictionQueryAttributesWithDefaults instantiates a new RestrictionQueryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionQueryAttributesWithDefaults() *RestrictionQueryAttributes {
	this := RestrictionQueryAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RestrictionQueryAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *RestrictionQueryAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRestrictionQuery returns the RestrictionQuery field value if set, zero value otherwise.
func (o *RestrictionQueryAttributes) GetRestrictionQuery() string {
	if o == nil || o.RestrictionQuery == nil {
		var ret string
		return ret
	}
	return *o.RestrictionQuery
}

// GetRestrictionQueryOk returns a tuple with the RestrictionQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryAttributes) GetRestrictionQueryOk() (*string, bool) {
	if o == nil || o.RestrictionQuery == nil {
		return nil, false
	}
	return o.RestrictionQuery, true
}

// HasRestrictionQuery returns a boolean if a field has been set.
func (o *RestrictionQueryAttributes) HasRestrictionQuery() bool {
	return o != nil && o.RestrictionQuery != nil
}

// SetRestrictionQuery gets a reference to the given string and assigns it to the RestrictionQuery field.
func (o *RestrictionQueryAttributes) SetRestrictionQuery(v string) {
	o.RestrictionQuery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionQueryAttributes) MarshalJSON() ([]byte, error) {
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
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.RestrictionQuery != nil {
		toSerialize["restriction_query"] = o.RestrictionQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestrictionQueryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time `json:"created_at,omitempty"`
		ModifiedAt       *time.Time `json:"modified_at,omitempty"`
		RestrictionQuery *string    `json:"restriction_query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "restriction_query"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.RestrictionQuery = all.RestrictionQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
