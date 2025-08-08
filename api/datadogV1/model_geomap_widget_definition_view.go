// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GeomapWidgetDefinitionView The view of the world that the map should render.
type GeomapWidgetDefinitionView struct {
	// A custom extent of the map defined by an array of four numbers in the order `[minLongitude, minLatitude, maxLongitude, maxLatitude]`.
	CustomExtent []float64 `json:"custom_extent,omitempty"`
	// The ISO code of a country, sub-division, or region to focus the map on. Or `WORLD`. Mutually exclusive with `custom_extent`.
	Focus string `json:"focus"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGeomapWidgetDefinitionView instantiates a new GeomapWidgetDefinitionView object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGeomapWidgetDefinitionView(focus string) *GeomapWidgetDefinitionView {
	this := GeomapWidgetDefinitionView{}
	this.Focus = focus
	return &this
}

// NewGeomapWidgetDefinitionViewWithDefaults instantiates a new GeomapWidgetDefinitionView object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGeomapWidgetDefinitionViewWithDefaults() *GeomapWidgetDefinitionView {
	this := GeomapWidgetDefinitionView{}
	return &this
}

// GetCustomExtent returns the CustomExtent field value if set, zero value otherwise.
func (o *GeomapWidgetDefinitionView) GetCustomExtent() []float64 {
	if o == nil || o.CustomExtent == nil {
		var ret []float64
		return ret
	}
	return o.CustomExtent
}

// GetCustomExtentOk returns a tuple with the CustomExtent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeomapWidgetDefinitionView) GetCustomExtentOk() (*[]float64, bool) {
	if o == nil || o.CustomExtent == nil {
		return nil, false
	}
	return &o.CustomExtent, true
}

// HasCustomExtent returns a boolean if a field has been set.
func (o *GeomapWidgetDefinitionView) HasCustomExtent() bool {
	return o != nil && o.CustomExtent != nil
}

// SetCustomExtent gets a reference to the given []float64 and assigns it to the CustomExtent field.
func (o *GeomapWidgetDefinitionView) SetCustomExtent(v []float64) {
	o.CustomExtent = v
}

// GetFocus returns the Focus field value.
func (o *GeomapWidgetDefinitionView) GetFocus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Focus
}

// GetFocusOk returns a tuple with the Focus field value
// and a boolean to check if the value has been set.
func (o *GeomapWidgetDefinitionView) GetFocusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Focus, true
}

// SetFocus sets field value.
func (o *GeomapWidgetDefinitionView) SetFocus(v string) {
	o.Focus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GeomapWidgetDefinitionView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomExtent != nil {
		toSerialize["custom_extent"] = o.CustomExtent
	}
	toSerialize["focus"] = o.Focus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GeomapWidgetDefinitionView) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomExtent []float64 `json:"custom_extent,omitempty"`
		Focus        *string   `json:"focus"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Focus == nil {
		return fmt.Errorf("required field focus missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_extent", "focus"})
	} else {
		return err
	}
	o.CustomExtent = all.CustomExtent
	o.Focus = *all.Focus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
