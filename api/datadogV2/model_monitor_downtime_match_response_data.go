// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorDowntimeMatchResponseData A downtime match.
type MonitorDowntimeMatchResponseData struct {
	// Downtime match details.
	Attributes *MonitorDowntimeMatchAttributeResponse `json:"attributes,omitempty"`
	// The downtime ID.
	Id *string `json:"id,omitempty"`
	// Monitor Downtime Match resource type.
	Type *MonitorDowntimeMatchResourceType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMonitorDowntimeMatchResponseData instantiates a new MonitorDowntimeMatchResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorDowntimeMatchResponseData() *MonitorDowntimeMatchResponseData {
	this := MonitorDowntimeMatchResponseData{}
	var typeVar MonitorDowntimeMatchResourceType = MONITORDOWNTIMEMATCHRESOURCETYPE_DOWNTIME_MATCH
	this.Type = &typeVar
	return &this
}

// NewMonitorDowntimeMatchResponseDataWithDefaults instantiates a new MonitorDowntimeMatchResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorDowntimeMatchResponseDataWithDefaults() *MonitorDowntimeMatchResponseData {
	this := MonitorDowntimeMatchResponseData{}
	var typeVar MonitorDowntimeMatchResourceType = MONITORDOWNTIMEMATCHRESOURCETYPE_DOWNTIME_MATCH
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseData) GetAttributes() MonitorDowntimeMatchAttributeResponse {
	if o == nil || o.Attributes == nil {
		var ret MonitorDowntimeMatchAttributeResponse
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseData) GetAttributesOk() (*MonitorDowntimeMatchAttributeResponse, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given MonitorDowntimeMatchAttributeResponse and assigns it to the Attributes field.
func (o *MonitorDowntimeMatchResponseData) SetAttributes(v MonitorDowntimeMatchAttributeResponse) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MonitorDowntimeMatchResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorDowntimeMatchResponseData) GetType() MonitorDowntimeMatchResourceType {
	if o == nil || o.Type == nil {
		var ret MonitorDowntimeMatchResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorDowntimeMatchResponseData) GetTypeOk() (*MonitorDowntimeMatchResourceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorDowntimeMatchResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MonitorDowntimeMatchResourceType and assigns it to the Type field.
func (o *MonitorDowntimeMatchResponseData) SetType(v MonitorDowntimeMatchResourceType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorDowntimeMatchResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorDowntimeMatchResponseData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Attributes *MonitorDowntimeMatchAttributeResponse `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Type       *MonitorDowntimeMatchResourceType      `json:"type,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	o.Type = all.Type
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
