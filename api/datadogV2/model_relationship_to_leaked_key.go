// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RelationshipToLeakedKey Relationship to the leak the access token was found in. `data` is null when the access token has not been detected as leaked.
type RelationshipToLeakedKey struct {
	// Relationship to the leak the access token was found in.
	Data NullableRelationshipToLeakedKeyData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRelationshipToLeakedKey instantiates a new RelationshipToLeakedKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationshipToLeakedKey(data NullableRelationshipToLeakedKeyData) *RelationshipToLeakedKey {
	this := RelationshipToLeakedKey{}
	this.Data = data
	return &this
}

// NewRelationshipToLeakedKeyWithDefaults instantiates a new RelationshipToLeakedKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationshipToLeakedKeyWithDefaults() *RelationshipToLeakedKey {
	this := RelationshipToLeakedKey{}
	return &this
}

// GetData returns the Data field value.
// If the value is explicit nil, the zero value for RelationshipToLeakedKeyData will be returned.
func (o *RelationshipToLeakedKey) GetData() RelationshipToLeakedKeyData {
	if o == nil || o.Data.Get() == nil {
		var ret RelationshipToLeakedKeyData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RelationshipToLeakedKey) GetDataOk() (*RelationshipToLeakedKeyData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value.
func (o *RelationshipToLeakedKey) SetData(v RelationshipToLeakedKeyData) {
	o.Data.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationshipToLeakedKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationshipToLeakedKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableRelationshipToLeakedKeyData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Data.IsSet() {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
