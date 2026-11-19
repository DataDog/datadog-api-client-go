// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchFunnelStep Funnel linked to the launch.
type ProductAnalyticsLaunchFunnelStep struct {
	// Definition of the funnel widget, as a JSON object.
	FunnelTileDef map[string]interface{} `json:"funnel_tile_def,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchFunnelStep instantiates a new ProductAnalyticsLaunchFunnelStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchFunnelStep() *ProductAnalyticsLaunchFunnelStep {
	this := ProductAnalyticsLaunchFunnelStep{}
	return &this
}

// NewProductAnalyticsLaunchFunnelStepWithDefaults instantiates a new ProductAnalyticsLaunchFunnelStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchFunnelStepWithDefaults() *ProductAnalyticsLaunchFunnelStep {
	this := ProductAnalyticsLaunchFunnelStep{}
	return &this
}

// GetFunnelTileDef returns the FunnelTileDef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchFunnelStep) GetFunnelTileDef() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FunnelTileDef
}

// GetFunnelTileDefOk returns a tuple with the FunnelTileDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchFunnelStep) GetFunnelTileDefOk() (*map[string]interface{}, bool) {
	if o == nil || o.FunnelTileDef == nil {
		return nil, false
	}
	return &o.FunnelTileDef, true
}

// HasFunnelTileDef returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchFunnelStep) HasFunnelTileDef() bool {
	return o != nil && o.FunnelTileDef != nil
}

// SetFunnelTileDef gets a reference to the given map[string]interface{} and assigns it to the FunnelTileDef field.
func (o *ProductAnalyticsLaunchFunnelStep) SetFunnelTileDef(v map[string]interface{}) {
	o.FunnelTileDef = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchFunnelStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FunnelTileDef != nil {
		toSerialize["funnel_tile_def"] = o.FunnelTileDef
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchFunnelStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FunnelTileDef map[string]interface{} `json:"funnel_tile_def,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"funnel_tile_def"})
	} else {
		return err
	}
	o.FunnelTileDef = all.FunnelTileDef

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
