// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveSessionsResponse JSON:API document.
type RumArchiveSessionsResponse struct {
	// Archived sessions.
	Data []RumArchiveSession `json:"data"`
	// Search metadata.
	Meta RumArchiveSessionsMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumArchiveSessionsResponse instantiates a new RumArchiveSessionsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumArchiveSessionsResponse(data []RumArchiveSession, meta RumArchiveSessionsMeta) *RumArchiveSessionsResponse {
	this := RumArchiveSessionsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewRumArchiveSessionsResponseWithDefaults instantiates a new RumArchiveSessionsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumArchiveSessionsResponseWithDefaults() *RumArchiveSessionsResponse {
	this := RumArchiveSessionsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *RumArchiveSessionsResponse) GetData() []RumArchiveSession {
	if o == nil {
		var ret []RumArchiveSession
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsResponse) GetDataOk() (*[]RumArchiveSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *RumArchiveSessionsResponse) SetData(v []RumArchiveSession) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *RumArchiveSessionsResponse) GetMeta() RumArchiveSessionsMeta {
	if o == nil {
		var ret RumArchiveSessionsMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionsResponse) GetMetaOk() (*RumArchiveSessionsMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *RumArchiveSessionsResponse) SetMeta(v RumArchiveSessionsMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumArchiveSessionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumArchiveSessionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]RumArchiveSession    `json:"data"`
		Meta *RumArchiveSessionsMeta `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
