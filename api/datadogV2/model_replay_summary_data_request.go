// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplaySummaryDataRequest Data object for a RUM replay summary request.
type ReplaySummaryDataRequest struct {
	// RUM replay summary request resource type.
	Type ReplaySummaryRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplaySummaryDataRequest instantiates a new ReplaySummaryDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplaySummaryDataRequest(typeVar ReplaySummaryRequestType) *ReplaySummaryDataRequest {
	this := ReplaySummaryDataRequest{}
	this.Type = typeVar
	return &this
}

// NewReplaySummaryDataRequestWithDefaults instantiates a new ReplaySummaryDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplaySummaryDataRequestWithDefaults() *ReplaySummaryDataRequest {
	this := ReplaySummaryDataRequest{}
	var typeVar ReplaySummaryRequestType = REPLAYSUMMARYREQUESTTYPE_REPLAY_SUMMARY_REQUEST
	this.Type = typeVar
	return &this
}

// GetType returns the Type field value.
func (o *ReplaySummaryDataRequest) GetType() ReplaySummaryRequestType {
	if o == nil {
		var ret ReplaySummaryRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReplaySummaryDataRequest) GetTypeOk() (*ReplaySummaryRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ReplaySummaryDataRequest) SetType(v ReplaySummaryRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplaySummaryDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReplaySummaryDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *ReplaySummaryRequestType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type"})
	} else {
		return err
	}

	hasInvalidField := false
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
