// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLink A link between a Page and another Datadog resource, such as an incident.
type OnCallPageLink struct {
	// The attributes of a Page link. A Page can be the child of another Datadog resource
	// (for example, when an incident triggers it), or the parent of one (for example, when the
	// Page is escalated to an incident).
	Attributes *OnCallPageLinkAttributes `json:"attributes,omitempty"`
	// The ID of the link.
	Id string `json:"id"`
	// Page link resource type.
	Type OnCallPageLinkType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageLink instantiates a new OnCallPageLink object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageLink(id string, typeVar OnCallPageLinkType) *OnCallPageLink {
	this := OnCallPageLink{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOnCallPageLinkWithDefaults instantiates a new OnCallPageLink object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageLinkWithDefaults() *OnCallPageLink {
	this := OnCallPageLink{}
	var typeVar OnCallPageLinkType = ONCALLPAGELINKTYPE_LINKS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *OnCallPageLink) GetAttributes() OnCallPageLinkAttributes {
	if o == nil || o.Attributes == nil {
		var ret OnCallPageLinkAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageLink) GetAttributesOk() (*OnCallPageLinkAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *OnCallPageLink) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given OnCallPageLinkAttributes and assigns it to the Attributes field.
func (o *OnCallPageLink) SetAttributes(v OnCallPageLinkAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *OnCallPageLink) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLink) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OnCallPageLink) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OnCallPageLink) GetType() OnCallPageLinkType {
	if o == nil {
		var ret OnCallPageLinkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnCallPageLink) GetTypeOk() (*OnCallPageLinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnCallPageLink) SetType(v OnCallPageLinkType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageLink) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *OnCallPageLinkAttributes `json:"attributes,omitempty"`
		Id         *string                   `json:"id"`
		Type       *OnCallPageLinkType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
