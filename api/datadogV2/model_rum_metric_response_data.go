// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMMetricResponseData The RUM-based metric properties.
type RUMMetricResponseData struct {
	// The object describing a Datadog RUM-based metric.
	Attributes *RUMMetricResponseAttributes `json:"attributes,omitempty"`
	// The name of the RUM-based metric.
	// On creation, it must be 1 to 200 characters long, contain only ASCII letters, digits, underscores, and periods, start with a letter, and not start with `rum.measure` or `rum.investigate`.
	// It cannot be changed after creation.
	Id *string `json:"id,omitempty"`
	// The type of the resource. The value must be `rum_metrics`.
	Type *RUMMetricType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMMetricResponseData instantiates a new RUMMetricResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMMetricResponseData() *RUMMetricResponseData {
	this := RUMMetricResponseData{}
	var typeVar RUMMetricType = RUMMETRICTYPE_RUM_METRICS
	this.Type = &typeVar
	return &this
}

// NewRUMMetricResponseDataWithDefaults instantiates a new RUMMetricResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMMetricResponseDataWithDefaults() *RUMMetricResponseData {
	this := RUMMetricResponseData{}
	var typeVar RUMMetricType = RUMMETRICTYPE_RUM_METRICS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RUMMetricResponseData) GetAttributes() RUMMetricResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret RUMMetricResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseData) GetAttributesOk() (*RUMMetricResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RUMMetricResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given RUMMetricResponseAttributes and assigns it to the Attributes field.
func (o *RUMMetricResponseData) SetAttributes(v RUMMetricResponseAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RUMMetricResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RUMMetricResponseData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RUMMetricResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RUMMetricResponseData) GetType() RUMMetricType {
	if o == nil || o.Type == nil {
		var ret RUMMetricType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMMetricResponseData) GetTypeOk() (*RUMMetricType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RUMMetricResponseData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given RUMMetricType and assigns it to the Type field.
func (o *RUMMetricResponseData) SetType(v RUMMetricType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMMetricResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMMetricResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *RUMMetricResponseAttributes `json:"attributes,omitempty"`
		Id         *string                      `json:"id,omitempty"`
		Type       *RUMMetricType               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
