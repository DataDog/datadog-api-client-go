// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryResponseMeta Top-level JSON:API meta block accompanying every DDSQL tabular query response.
// Carries standard observability handles for client-side correlation.
type DdsqlTabularQueryResponseMeta struct {
	// Server-side time spent serving this request, in milliseconds.
	Elapsed int64 `json:"elapsed"`
	// Echo of the `DD-Request-ID` header assigned by Datadog's edge to this request,
	// for support correlation.
	RequestId string `json:"request_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDdsqlTabularQueryResponseMeta instantiates a new DdsqlTabularQueryResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDdsqlTabularQueryResponseMeta(elapsed int64, requestId string) *DdsqlTabularQueryResponseMeta {
	this := DdsqlTabularQueryResponseMeta{}
	this.Elapsed = elapsed
	this.RequestId = requestId
	return &this
}

// NewDdsqlTabularQueryResponseMetaWithDefaults instantiates a new DdsqlTabularQueryResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDdsqlTabularQueryResponseMetaWithDefaults() *DdsqlTabularQueryResponseMeta {
	this := DdsqlTabularQueryResponseMeta{}
	return &this
}

// GetElapsed returns the Elapsed field value.
func (o *DdsqlTabularQueryResponseMeta) GetElapsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryResponseMeta) GetElapsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Elapsed, true
}

// SetElapsed sets field value.
func (o *DdsqlTabularQueryResponseMeta) SetElapsed(v int64) {
	o.Elapsed = v
}

// GetRequestId returns the RequestId field value.
func (o *DdsqlTabularQueryResponseMeta) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryResponseMeta) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *DdsqlTabularQueryResponseMeta) SetRequestId(v string) {
	o.RequestId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DdsqlTabularQueryResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["elapsed"] = o.Elapsed
	toSerialize["request_id"] = o.RequestId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DdsqlTabularQueryResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Elapsed   *int64  `json:"elapsed"`
		RequestId *string `json:"request_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Elapsed == nil {
		return fmt.Errorf("required field elapsed missing")
	}
	if all.RequestId == nil {
		return fmt.Errorf("required field request_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed", "request_id"})
	} else {
		return err
	}
	o.Elapsed = *all.Elapsed
	o.RequestId = *all.RequestId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
