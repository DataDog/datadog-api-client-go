// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesTags A map of tags providing additional metadata for the SCA scan.
type ScaRequestDataAttributesTags struct {
	// Tool metadata included in SCA tags.
	Tool *ScaRequestDataAttributesTagsTool `json:"tool,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]string      `json:"-"`
}

// NewScaRequestDataAttributesTags instantiates a new ScaRequestDataAttributesTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesTags() *ScaRequestDataAttributesTags {
	this := ScaRequestDataAttributesTags{}
	return &this
}

// NewScaRequestDataAttributesTagsWithDefaults instantiates a new ScaRequestDataAttributesTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesTagsWithDefaults() *ScaRequestDataAttributesTags {
	this := ScaRequestDataAttributesTags{}
	return &this
}

// GetTool returns the Tool field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesTags) GetTool() ScaRequestDataAttributesTagsTool {
	if o == nil || o.Tool == nil {
		var ret ScaRequestDataAttributesTagsTool
		return ret
	}
	return *o.Tool
}

// GetToolOk returns a tuple with the Tool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesTags) GetToolOk() (*ScaRequestDataAttributesTagsTool, bool) {
	if o == nil || o.Tool == nil {
		return nil, false
	}
	return o.Tool, true
}

// HasTool returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesTags) HasTool() bool {
	return o != nil && o.Tool != nil
}

// SetTool gets a reference to the given ScaRequestDataAttributesTagsTool and assigns it to the Tool field.
func (o *ScaRequestDataAttributesTags) SetTool(v ScaRequestDataAttributesTagsTool) {
	o.Tool = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Tool != nil {
		toSerialize["tool"] = o.Tool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Tool *ScaRequestDataAttributesTagsTool `json:"tool,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]string)
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"tool"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Tool != nil && all.Tool.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tool = all.Tool

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
