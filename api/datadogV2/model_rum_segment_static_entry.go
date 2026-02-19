// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentStaticEntry A static user list entry within a segment data query.
type RumSegmentStaticEntry struct {
	// The identifier of the static list.
	Id string `json:"id"`
	// The name of the static list.
	Name string `json:"name"`
	// The number of users in the static list.
	UserCount int64 `json:"user_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentStaticEntry instantiates a new RumSegmentStaticEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentStaticEntry(id string, name string, userCount int64) *RumSegmentStaticEntry {
	this := RumSegmentStaticEntry{}
	this.Id = id
	this.Name = name
	this.UserCount = userCount
	return &this
}

// NewRumSegmentStaticEntryWithDefaults instantiates a new RumSegmentStaticEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentStaticEntryWithDefaults() *RumSegmentStaticEntry {
	this := RumSegmentStaticEntry{}
	return &this
}

// GetId returns the Id field value.
func (o *RumSegmentStaticEntry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RumSegmentStaticEntry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RumSegmentStaticEntry) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *RumSegmentStaticEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentStaticEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentStaticEntry) SetName(v string) {
	o.Name = v
}

// GetUserCount returns the UserCount field value.
func (o *RumSegmentStaticEntry) GetUserCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value
// and a boolean to check if the value has been set.
func (o *RumSegmentStaticEntry) GetUserCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserCount, true
}

// SetUserCount sets field value.
func (o *RumSegmentStaticEntry) SetUserCount(v int64) {
	o.UserCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentStaticEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["user_count"] = o.UserCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentStaticEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string `json:"id"`
		Name      *string `json:"name"`
		UserCount *int64  `json:"user_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.UserCount == nil {
		return fmt.Errorf("required field user_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "user_count"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name
	o.UserCount = *all.UserCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
