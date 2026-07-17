// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardFixedTimeframe A fixed dashboard timeframe.
type DashboardFixedTimeframe struct {
	// Start time in milliseconds since epoch.
	From int64 `json:"from"`
	// End time in milliseconds since epoch.
	To int64 `json:"to"`
	// Type of fixed timeframe.
	Type DashboardFixedTimeframeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardFixedTimeframe instantiates a new DashboardFixedTimeframe object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardFixedTimeframe(from int64, to int64, typeVar DashboardFixedTimeframeType) *DashboardFixedTimeframe {
	this := DashboardFixedTimeframe{}
	this.From = from
	this.To = to
	this.Type = typeVar
	return &this
}

// NewDashboardFixedTimeframeWithDefaults instantiates a new DashboardFixedTimeframe object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardFixedTimeframeWithDefaults() *DashboardFixedTimeframe {
	this := DashboardFixedTimeframe{}
	return &this
}

// GetFrom returns the From field value.
func (o *DashboardFixedTimeframe) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *DashboardFixedTimeframe) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *DashboardFixedTimeframe) SetFrom(v int64) {
	o.From = v
}

// GetTo returns the To field value.
func (o *DashboardFixedTimeframe) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *DashboardFixedTimeframe) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *DashboardFixedTimeframe) SetTo(v int64) {
	o.To = v
}

// GetType returns the Type field value.
func (o *DashboardFixedTimeframe) GetType() DashboardFixedTimeframeType {
	if o == nil {
		var ret DashboardFixedTimeframeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardFixedTimeframe) GetTypeOk() (*DashboardFixedTimeframeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardFixedTimeframe) SetType(v DashboardFixedTimeframeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardFixedTimeframe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardFixedTimeframe) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *int64                       `json:"from"`
		To   *int64                       `json:"to"`
		Type *DashboardFixedTimeframeType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "to", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.From = *all.From
	o.To = *all.To
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
