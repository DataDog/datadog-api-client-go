// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReplayAnalysisIssueSessionsResponse A paginated list of sessions related to a RUM replay analysis issue.
type ReplayAnalysisIssueSessionsResponse struct {
	// Array of session data objects related to the issue.
	Data []ReplayAnalysisIssueSessionData `json:"data"`
	// Metadata object for paginated issue list responses.
	Meta ReplayAnalysisIssueMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReplayAnalysisIssueSessionsResponse instantiates a new ReplayAnalysisIssueSessionsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReplayAnalysisIssueSessionsResponse(data []ReplayAnalysisIssueSessionData, meta ReplayAnalysisIssueMeta) *ReplayAnalysisIssueSessionsResponse {
	this := ReplayAnalysisIssueSessionsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewReplayAnalysisIssueSessionsResponseWithDefaults instantiates a new ReplayAnalysisIssueSessionsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReplayAnalysisIssueSessionsResponseWithDefaults() *ReplayAnalysisIssueSessionsResponse {
	this := ReplayAnalysisIssueSessionsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ReplayAnalysisIssueSessionsResponse) GetData() []ReplayAnalysisIssueSessionData {
	if o == nil {
		var ret []ReplayAnalysisIssueSessionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueSessionsResponse) GetDataOk() (*[]ReplayAnalysisIssueSessionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ReplayAnalysisIssueSessionsResponse) SetData(v []ReplayAnalysisIssueSessionData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *ReplayAnalysisIssueSessionsResponse) GetMeta() ReplayAnalysisIssueMeta {
	if o == nil {
		var ret ReplayAnalysisIssueMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ReplayAnalysisIssueSessionsResponse) GetMetaOk() (*ReplayAnalysisIssueMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *ReplayAnalysisIssueSessionsResponse) SetMeta(v ReplayAnalysisIssueMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReplayAnalysisIssueSessionsResponse) MarshalJSON() ([]byte, error) {
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
func (o *ReplayAnalysisIssueSessionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ReplayAnalysisIssueSessionData `json:"data"`
		Meta *ReplayAnalysisIssueMeta          `json:"meta"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
