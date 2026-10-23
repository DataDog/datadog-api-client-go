// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLinksRelationship Links between the page and other Datadog resources, such as incidents. Always empty when listing pages.
type OnCallPageLinksRelationship struct {
	// The list of page links.
	Data []OnCallPageLinkRelationshipData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageLinksRelationship instantiates a new OnCallPageLinksRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageLinksRelationship() *OnCallPageLinksRelationship {
	this := OnCallPageLinksRelationship{}
	return &this
}

// NewOnCallPageLinksRelationshipWithDefaults instantiates a new OnCallPageLinksRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageLinksRelationshipWithDefaults() *OnCallPageLinksRelationship {
	this := OnCallPageLinksRelationship{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OnCallPageLinksRelationship) GetData() []OnCallPageLinkRelationshipData {
	if o == nil || o.Data == nil {
		var ret []OnCallPageLinkRelationshipData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageLinksRelationship) GetDataOk() (*[]OnCallPageLinkRelationshipData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OnCallPageLinksRelationship) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []OnCallPageLinkRelationshipData and assigns it to the Data field.
func (o *OnCallPageLinksRelationship) SetData(v []OnCallPageLinkRelationshipData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageLinksRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageLinksRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []OnCallPageLinkRelationshipData `json:"data,omitempty"`
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
