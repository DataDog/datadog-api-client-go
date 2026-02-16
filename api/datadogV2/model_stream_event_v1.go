// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StreamEventV1
type StreamEventV1 struct {
	// The event data payload.
	Data interface{} `json:"data,omitempty"`
	// The type of stream event.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStreamEventV1 instantiates a new StreamEventV1 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStreamEventV1(typeVar string) *StreamEventV1 {
	this := StreamEventV1{}
	this.Type = typeVar
	return &this
}

// NewStreamEventV1WithDefaults instantiates a new StreamEventV1 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStreamEventV1WithDefaults() *StreamEventV1 {
	this := StreamEventV1{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StreamEventV1) GetData() interface{} {
	if o == nil || o.Data == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamEventV1) GetDataOk() (*interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StreamEventV1) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *StreamEventV1) SetData(v interface{}) {
	o.Data = v
}

// GetType returns the Type field value.
func (o *StreamEventV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StreamEventV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StreamEventV1) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StreamEventV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StreamEventV1) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data interface{} `json:"data,omitempty"`
		Type *string     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "type"})
	} else {
		return err
	}
	o.Data = all.Data
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
