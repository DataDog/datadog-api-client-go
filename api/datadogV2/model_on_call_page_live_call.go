// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageLiveCall Live call information for a Page triggered by a live call route.
type OnCallPageLiveCall struct {
	// Information about a live call that has ended.
	CallOver *OnCallPageLiveCallOver `json:"call_over,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageLiveCall instantiates a new OnCallPageLiveCall object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageLiveCall() *OnCallPageLiveCall {
	this := OnCallPageLiveCall{}
	return &this
}

// NewOnCallPageLiveCallWithDefaults instantiates a new OnCallPageLiveCall object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageLiveCallWithDefaults() *OnCallPageLiveCall {
	this := OnCallPageLiveCall{}
	return &this
}

// GetCallOver returns the CallOver field value if set, zero value otherwise.
func (o *OnCallPageLiveCall) GetCallOver() OnCallPageLiveCallOver {
	if o == nil || o.CallOver == nil {
		var ret OnCallPageLiveCallOver
		return ret
	}
	return *o.CallOver
}

// GetCallOverOk returns a tuple with the CallOver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageLiveCall) GetCallOverOk() (*OnCallPageLiveCallOver, bool) {
	if o == nil || o.CallOver == nil {
		return nil, false
	}
	return o.CallOver, true
}

// HasCallOver returns a boolean if a field has been set.
func (o *OnCallPageLiveCall) HasCallOver() bool {
	return o != nil && o.CallOver != nil
}

// SetCallOver gets a reference to the given OnCallPageLiveCallOver and assigns it to the CallOver field.
func (o *OnCallPageLiveCall) SetCallOver(v OnCallPageLiveCallOver) {
	o.CallOver = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageLiveCall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CallOver != nil {
		toSerialize["call_over"] = o.CallOver
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageLiveCall) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CallOver *OnCallPageLiveCallOver `json:"call_over,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"call_over"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CallOver != nil && all.CallOver.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CallOver = all.CallOver

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
