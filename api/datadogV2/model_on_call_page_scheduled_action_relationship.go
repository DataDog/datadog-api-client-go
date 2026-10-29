// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageScheduledActionRelationship The next action scheduled on the Page, such as an escalation.
type OnCallPageScheduledActionRelationship struct {
	// Reference to the next scheduled action on the Page.
	Data NullableOnCallPageScheduledActionRelationshipData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageScheduledActionRelationship instantiates a new OnCallPageScheduledActionRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageScheduledActionRelationship() *OnCallPageScheduledActionRelationship {
	this := OnCallPageScheduledActionRelationship{}
	return &this
}

// NewOnCallPageScheduledActionRelationshipWithDefaults instantiates a new OnCallPageScheduledActionRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageScheduledActionRelationshipWithDefaults() *OnCallPageScheduledActionRelationship {
	this := OnCallPageScheduledActionRelationship{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnCallPageScheduledActionRelationship) GetData() OnCallPageScheduledActionRelationshipData {
	if o == nil || o.Data.Get() == nil {
		var ret OnCallPageScheduledActionRelationshipData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OnCallPageScheduledActionRelationship) GetDataOk() (*OnCallPageScheduledActionRelationshipData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *OnCallPageScheduledActionRelationship) HasData() bool {
	return o != nil && o.Data.IsSet()
}

// SetData gets a reference to the given NullableOnCallPageScheduledActionRelationshipData and assigns it to the Data field.
func (o *OnCallPageScheduledActionRelationship) SetData(v OnCallPageScheduledActionRelationshipData) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil.
func (o *OnCallPageScheduledActionRelationship) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil.
func (o *OnCallPageScheduledActionRelationship) UnsetData() {
	o.Data.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageScheduledActionRelationship) MarshalJSON() ([]byte, error) {
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
func (o *OnCallPageScheduledActionRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data NullableOnCallPageScheduledActionRelationshipData `json:"data,omitempty"`
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
