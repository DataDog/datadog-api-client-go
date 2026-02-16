// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PickActionResponse
type PickActionResponse struct {
	// The matching actions.
	Matches []ActionMatch `json:"matches"`
	// The request ID.
	RequestId string `json:"request_id"`
	// The total number of matches.
	TotalMatches int64 `json:"total_matches"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPickActionResponse instantiates a new PickActionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPickActionResponse(matches []ActionMatch, requestId string, totalMatches int64) *PickActionResponse {
	this := PickActionResponse{}
	this.Matches = matches
	this.RequestId = requestId
	this.TotalMatches = totalMatches
	return &this
}

// NewPickActionResponseWithDefaults instantiates a new PickActionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPickActionResponseWithDefaults() *PickActionResponse {
	this := PickActionResponse{}
	return &this
}

// GetMatches returns the Matches field value.
func (o *PickActionResponse) GetMatches() []ActionMatch {
	if o == nil {
		var ret []ActionMatch
		return ret
	}
	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value
// and a boolean to check if the value has been set.
func (o *PickActionResponse) GetMatchesOk() (*[]ActionMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matches, true
}

// SetMatches sets field value.
func (o *PickActionResponse) SetMatches(v []ActionMatch) {
	o.Matches = v
}

// GetRequestId returns the RequestId field value.
func (o *PickActionResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PickActionResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *PickActionResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetTotalMatches returns the TotalMatches field value.
func (o *PickActionResponse) GetTotalMatches() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalMatches
}

// GetTotalMatchesOk returns a tuple with the TotalMatches field value
// and a boolean to check if the value has been set.
func (o *PickActionResponse) GetTotalMatchesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMatches, true
}

// SetTotalMatches sets field value.
func (o *PickActionResponse) SetTotalMatches(v int64) {
	o.TotalMatches = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PickActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["matches"] = o.Matches
	toSerialize["request_id"] = o.RequestId
	toSerialize["total_matches"] = o.TotalMatches

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PickActionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Matches      *[]ActionMatch `json:"matches"`
		RequestId    *string        `json:"request_id"`
		TotalMatches *int64         `json:"total_matches"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Matches == nil {
		return fmt.Errorf("required field matches missing")
	}
	if all.RequestId == nil {
		return fmt.Errorf("required field request_id missing")
	}
	if all.TotalMatches == nil {
		return fmt.Errorf("required field total_matches missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"matches", "request_id", "total_matches"})
	} else {
		return err
	}
	o.Matches = *all.Matches
	o.RequestId = *all.RequestId
	o.TotalMatches = *all.TotalMatches

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
