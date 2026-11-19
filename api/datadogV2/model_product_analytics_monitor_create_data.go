// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorCreateData Data of the Product Analytics monitor to create.
type ProductAnalyticsMonitorCreateData struct {
	// Attributes of the Product Analytics monitor to create.
	Attributes ProductAnalyticsMonitorCreateAttributes `json:"attributes"`
	// Product Analytics monitor request resource type.
	Type ProductAnalyticsMonitorRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorCreateData instantiates a new ProductAnalyticsMonitorCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorCreateData(attributes ProductAnalyticsMonitorCreateAttributes, typeVar ProductAnalyticsMonitorRequestType) *ProductAnalyticsMonitorCreateData {
	this := ProductAnalyticsMonitorCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsMonitorCreateDataWithDefaults instantiates a new ProductAnalyticsMonitorCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorCreateDataWithDefaults() *ProductAnalyticsMonitorCreateData {
	this := ProductAnalyticsMonitorCreateData{}
	var typeVar ProductAnalyticsMonitorRequestType = PRODUCTANALYTICSMONITORREQUESTTYPE_PA_MONITOR_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ProductAnalyticsMonitorCreateData) GetAttributes() ProductAnalyticsMonitorCreateAttributes {
	if o == nil {
		var ret ProductAnalyticsMonitorCreateAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateData) GetAttributesOk() (*ProductAnalyticsMonitorCreateAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ProductAnalyticsMonitorCreateData) SetAttributes(v ProductAnalyticsMonitorCreateAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsMonitorCreateData) GetType() ProductAnalyticsMonitorRequestType {
	if o == nil {
		var ret ProductAnalyticsMonitorRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorCreateData) GetTypeOk() (*ProductAnalyticsMonitorRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsMonitorCreateData) SetType(v ProductAnalyticsMonitorRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsMonitorCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProductAnalyticsMonitorCreateAttributes `json:"attributes"`
		Type       *ProductAnalyticsMonitorRequestType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
