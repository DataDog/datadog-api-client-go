// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotRequests20270101OneOf Test-only generated schema for the dashboard versioning proof of concept.
type ScatterplotRequests20270101OneOf struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	X ScatterplotBackendAxisRequest `json:"x"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Y ScatterplotBackendAxisRequest `json:"y"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewScatterplotRequests20270101OneOf instantiates a new ScatterplotRequests20270101OneOf object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotRequests20270101OneOf(x ScatterplotBackendAxisRequest, y ScatterplotBackendAxisRequest) *ScatterplotRequests20270101OneOf {
	this := ScatterplotRequests20270101OneOf{}
	this.X = x
	this.Y = y
	return &this
}

// NewScatterplotRequests20270101OneOfWithDefaults instantiates a new ScatterplotRequests20270101OneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotRequests20270101OneOfWithDefaults() *ScatterplotRequests20270101OneOf {
	this := ScatterplotRequests20270101OneOf{}
	return &this
}

// GetX returns the X field value.
func (o *ScatterplotRequests20270101OneOf) GetX() ScatterplotBackendAxisRequest {
	if o == nil {
		var ret ScatterplotBackendAxisRequest
		return ret
	}
	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ScatterplotRequests20270101OneOf) GetXOk() (*ScatterplotBackendAxisRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value.
func (o *ScatterplotRequests20270101OneOf) SetX(v ScatterplotBackendAxisRequest) {
	o.X = v
}

// GetY returns the Y field value.
func (o *ScatterplotRequests20270101OneOf) GetY() ScatterplotBackendAxisRequest {
	if o == nil {
		var ret ScatterplotBackendAxisRequest
		return ret
	}
	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ScatterplotRequests20270101OneOf) GetYOk() (*ScatterplotBackendAxisRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value.
func (o *ScatterplotRequests20270101OneOf) SetY(v ScatterplotBackendAxisRequest) {
	o.Y = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotRequests20270101OneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotRequests20270101OneOf) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		X *ScatterplotBackendAxisRequest `json:"x"`
		Y *ScatterplotBackendAxisRequest `json:"y"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.X == nil {
		return fmt.Errorf("required field x missing")
	}
	if all.Y == nil {
		return fmt.Errorf("required field y missing")
	}

	hasInvalidField := false
	if all.X.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.X = *all.X
	if all.Y.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Y = *all.Y

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
