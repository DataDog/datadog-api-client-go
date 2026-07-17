// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardLiveTimeframe A live dashboard timeframe.
type DashboardLiveTimeframe struct {
	// Type of live timeframe.
	Type DashboardLiveTimeframeType `json:"type"`
	// Unit of the time span.
	Unit WidgetLiveSpanUnit `json:"unit"`
	// Value of the live timeframe span.
	Value int64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardLiveTimeframe instantiates a new DashboardLiveTimeframe object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardLiveTimeframe(typeVar DashboardLiveTimeframeType, unit WidgetLiveSpanUnit, value int64) *DashboardLiveTimeframe {
	this := DashboardLiveTimeframe{}
	this.Type = typeVar
	this.Unit = unit
	this.Value = value
	return &this
}

// NewDashboardLiveTimeframeWithDefaults instantiates a new DashboardLiveTimeframe object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardLiveTimeframeWithDefaults() *DashboardLiveTimeframe {
	this := DashboardLiveTimeframe{}
	return &this
}

// GetType returns the Type field value.
func (o *DashboardLiveTimeframe) GetType() DashboardLiveTimeframeType {
	if o == nil {
		var ret DashboardLiveTimeframeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardLiveTimeframe) GetTypeOk() (*DashboardLiveTimeframeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardLiveTimeframe) SetType(v DashboardLiveTimeframeType) {
	o.Type = v
}

// GetUnit returns the Unit field value.
func (o *DashboardLiveTimeframe) GetUnit() WidgetLiveSpanUnit {
	if o == nil {
		var ret WidgetLiveSpanUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *DashboardLiveTimeframe) GetUnitOk() (*WidgetLiveSpanUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value.
func (o *DashboardLiveTimeframe) SetUnit(v WidgetLiveSpanUnit) {
	o.Unit = v
}

// GetValue returns the Value field value.
func (o *DashboardLiveTimeframe) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DashboardLiveTimeframe) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *DashboardLiveTimeframe) SetValue(v int64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardLiveTimeframe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardLiveTimeframe) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type  *DashboardLiveTimeframeType `json:"type"`
		Unit  *WidgetLiveSpanUnit         `json:"unit"`
		Value *int64                      `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Unit == nil {
		return fmt.Errorf("required field unit missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "unit", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.Unit.IsValid() {
		hasInvalidField = true
	} else {
		o.Unit = *all.Unit
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
