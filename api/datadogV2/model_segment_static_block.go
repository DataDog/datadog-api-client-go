// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentStaticBlock A fixed list of users uploaded to Datadog.
type SegmentStaticBlock struct {
	// ID of the uploaded list.
	Id *string `json:"id,omitempty"`
	// Name of the block, used in `combination`.
	Name *string `json:"name,omitempty"`
	// Number of users in the list.
	UserCount *int64 `json:"user_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentStaticBlock instantiates a new SegmentStaticBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentStaticBlock() *SegmentStaticBlock {
	this := SegmentStaticBlock{}
	return &this
}

// NewSegmentStaticBlockWithDefaults instantiates a new SegmentStaticBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentStaticBlockWithDefaults() *SegmentStaticBlock {
	this := SegmentStaticBlock{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SegmentStaticBlock) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentStaticBlock) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SegmentStaticBlock) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SegmentStaticBlock) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentStaticBlock) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentStaticBlock) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentStaticBlock) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentStaticBlock) SetName(v string) {
	o.Name = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *SegmentStaticBlock) GetUserCount() int64 {
	if o == nil || o.UserCount == nil {
		var ret int64
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentStaticBlock) GetUserCountOk() (*int64, bool) {
	if o == nil || o.UserCount == nil {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *SegmentStaticBlock) HasUserCount() bool {
	return o != nil && o.UserCount != nil
}

// SetUserCount gets a reference to the given int64 and assigns it to the UserCount field.
func (o *SegmentStaticBlock) SetUserCount(v int64) {
	o.UserCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentStaticBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UserCount != nil {
		toSerialize["user_count"] = o.UserCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentStaticBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *string `json:"id,omitempty"`
		Name      *string `json:"name,omitempty"`
		UserCount *int64  `json:"user_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "user_count"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Name = all.Name
	o.UserCount = all.UserCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
