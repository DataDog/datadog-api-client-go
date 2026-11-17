// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsMonitorUnlinkData Data of the Product Analytics monitor unlink request.
type ProductAnalyticsMonitorUnlinkData struct {
	// The chart link to remove. It must match the stored link exactly:
	// `chart_id` alone for a saved chart, or `chart_id` and `dashboard_id` for a dashboard widget.
	Attributes ProductAnalyticsMonitorUnlinkAttributes `json:"attributes"`
	// Product Analytics monitor unlink request resource type.
	Type ProductAnalyticsMonitorUnlinkRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsMonitorUnlinkData instantiates a new ProductAnalyticsMonitorUnlinkData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsMonitorUnlinkData(attributes ProductAnalyticsMonitorUnlinkAttributes, typeVar ProductAnalyticsMonitorUnlinkRequestType) *ProductAnalyticsMonitorUnlinkData {
	this := ProductAnalyticsMonitorUnlinkData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsMonitorUnlinkDataWithDefaults instantiates a new ProductAnalyticsMonitorUnlinkData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsMonitorUnlinkDataWithDefaults() *ProductAnalyticsMonitorUnlinkData {
	this := ProductAnalyticsMonitorUnlinkData{}
	var typeVar ProductAnalyticsMonitorUnlinkRequestType = PRODUCTANALYTICSMONITORUNLINKREQUESTTYPE_PA_MONITOR_UNLINK_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ProductAnalyticsMonitorUnlinkData) GetAttributes() ProductAnalyticsMonitorUnlinkAttributes {
	if o == nil {
		var ret ProductAnalyticsMonitorUnlinkAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUnlinkData) GetAttributesOk() (*ProductAnalyticsMonitorUnlinkAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ProductAnalyticsMonitorUnlinkData) SetAttributes(v ProductAnalyticsMonitorUnlinkAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsMonitorUnlinkData) GetType() ProductAnalyticsMonitorUnlinkRequestType {
	if o == nil {
		var ret ProductAnalyticsMonitorUnlinkRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsMonitorUnlinkData) GetTypeOk() (*ProductAnalyticsMonitorUnlinkRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsMonitorUnlinkData) SetType(v ProductAnalyticsMonitorUnlinkRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsMonitorUnlinkData) MarshalJSON() ([]byte, error) {
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
func (o *ProductAnalyticsMonitorUnlinkData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ProductAnalyticsMonitorUnlinkAttributes  `json:"attributes"`
		Type       *ProductAnalyticsMonitorUnlinkRequestType `json:"type"`
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
