// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageAuthorRelationship Reference to the entity that performed an action on the Page.
type OnCallPageAuthorRelationship struct {
	// Reference to the author of a Page action.
	Data NullableOnCallPageAuthorRelationshipData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageAuthorRelationship instantiates a new OnCallPageAuthorRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageAuthorRelationship() *OnCallPageAuthorRelationship {
	this := OnCallPageAuthorRelationship{}
	return &this
}

// NewOnCallPageAuthorRelationshipWithDefaults instantiates a new OnCallPageAuthorRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageAuthorRelationshipWithDefaults() *OnCallPageAuthorRelationship {
	this := OnCallPageAuthorRelationship{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnCallPageAuthorRelationship) GetData() OnCallPageAuthorRelationshipData {
	if o == nil || o.Data.Get() == nil {
		var ret OnCallPageAuthorRelationshipData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OnCallPageAuthorRelationship) GetDataOk() (*OnCallPageAuthorRelationshipData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *OnCallPageAuthorRelationship) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given NullableOnCallPageAuthorRelationshipData and assigns it to the Data field.
func (o *OnCallPageAuthorRelationship) SetData(v OnCallPageAuthorRelationshipData) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *OnCallPageAuthorRelationship) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil.
func (o *OnCallPageAuthorRelationship) UnsetData() {
	o.Data.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageAuthorRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageAuthorRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableOnCallPageAuthorRelationshipData `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
