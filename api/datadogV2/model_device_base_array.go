// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceBaseArray Response body for the list devices endpoint.
type DeviceBaseArray struct {
	// List of devices matching the request.
	Data []DeviceBaseData `json:"data"`
	// Links related to the current page of devices.
	Links *DeviceBaseArrayLinks `json:"links,omitempty"`
	// Metadata about the list devices response.
	Meta *DeviceBaseArrayMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeviceBaseArray instantiates a new DeviceBaseArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceBaseArray(data []DeviceBaseData) *DeviceBaseArray {
	this := DeviceBaseArray{}
	this.Data = data
	return &this
}

// NewDeviceBaseArrayWithDefaults instantiates a new DeviceBaseArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceBaseArrayWithDefaults() *DeviceBaseArray {
	this := DeviceBaseArray{}
	return &this
}

// GetData returns the Data field value.
func (o *DeviceBaseArray) GetData() []DeviceBaseData {
	if o == nil {
		var ret []DeviceBaseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeviceBaseArray) GetDataOk() (*[]DeviceBaseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *DeviceBaseArray) SetData(v []DeviceBaseData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceBaseArray) GetLinks() DeviceBaseArrayLinks {
	if o == nil || o.Links == nil {
		var ret DeviceBaseArrayLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseArray) GetLinksOk() (*DeviceBaseArrayLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceBaseArray) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given DeviceBaseArrayLinks and assigns it to the Links field.
func (o *DeviceBaseArray) SetLinks(v DeviceBaseArrayLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DeviceBaseArray) GetMeta() DeviceBaseArrayMeta {
	if o == nil || o.Meta == nil {
		var ret DeviceBaseArrayMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseArray) GetMetaOk() (*DeviceBaseArrayMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DeviceBaseArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DeviceBaseArrayMeta and assigns it to the Meta field.
func (o *DeviceBaseArray) SetMeta(v DeviceBaseArrayMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceBaseArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeviceBaseArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data  *[]DeviceBaseData     `json:"data"`
		Links *DeviceBaseArrayLinks `json:"links,omitempty"`
		Meta  *DeviceBaseArrayMeta  `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "links", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
