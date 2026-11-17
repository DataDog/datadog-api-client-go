// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentPreviewResponseData Data of the segment preview response.
type SegmentPreviewResponseData struct {
	// Members of the previewed segment.
	Attributes SegmentPreviewResponseAttributes `json:"attributes"`
	// ID of the preview response.
	Id string `json:"id"`
	// Segment preview response resource type.
	Type SegmentPreviewResponseType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentPreviewResponseData instantiates a new SegmentPreviewResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentPreviewResponseData(attributes SegmentPreviewResponseAttributes, id string, typeVar SegmentPreviewResponseType) *SegmentPreviewResponseData {
	this := SegmentPreviewResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewSegmentPreviewResponseDataWithDefaults instantiates a new SegmentPreviewResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentPreviewResponseDataWithDefaults() *SegmentPreviewResponseData {
	this := SegmentPreviewResponseData{}
	var typeVar SegmentPreviewResponseType = SEGMENTPREVIEWRESPONSETYPE_UNIFIED_SEGMENT_PREVIEW_RESPONSE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *SegmentPreviewResponseData) GetAttributes() SegmentPreviewResponseAttributes {
	if o == nil {
		var ret SegmentPreviewResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SegmentPreviewResponseData) GetAttributesOk() (*SegmentPreviewResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *SegmentPreviewResponseData) SetAttributes(v SegmentPreviewResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *SegmentPreviewResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SegmentPreviewResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SegmentPreviewResponseData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *SegmentPreviewResponseData) GetType() SegmentPreviewResponseType {
	if o == nil {
		var ret SegmentPreviewResponseType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SegmentPreviewResponseData) GetTypeOk() (*SegmentPreviewResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SegmentPreviewResponseData) SetType(v SegmentPreviewResponseType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentPreviewResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentPreviewResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SegmentPreviewResponseAttributes `json:"attributes"`
		Id         *string                           `json:"id"`
		Type       *SegmentPreviewResponseType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
