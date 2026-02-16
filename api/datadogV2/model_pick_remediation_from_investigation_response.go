// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PickRemediationFromInvestigationResponse
type PickRemediationFromInvestigationResponse struct {
	// The keywords generated and used for finding actions.
	KeywordsUsed []string `json:"keywords_used"`
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

// NewPickRemediationFromInvestigationResponse instantiates a new PickRemediationFromInvestigationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPickRemediationFromInvestigationResponse(keywordsUsed []string, matches []ActionMatch, requestId string, totalMatches int64) *PickRemediationFromInvestigationResponse {
	this := PickRemediationFromInvestigationResponse{}
	this.KeywordsUsed = keywordsUsed
	this.Matches = matches
	this.RequestId = requestId
	this.TotalMatches = totalMatches
	return &this
}

// NewPickRemediationFromInvestigationResponseWithDefaults instantiates a new PickRemediationFromInvestigationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPickRemediationFromInvestigationResponseWithDefaults() *PickRemediationFromInvestigationResponse {
	this := PickRemediationFromInvestigationResponse{}
	return &this
}

// GetKeywordsUsed returns the KeywordsUsed field value.
func (o *PickRemediationFromInvestigationResponse) GetKeywordsUsed() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.KeywordsUsed
}

// GetKeywordsUsedOk returns a tuple with the KeywordsUsed field value
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationResponse) GetKeywordsUsedOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordsUsed, true
}

// SetKeywordsUsed sets field value.
func (o *PickRemediationFromInvestigationResponse) SetKeywordsUsed(v []string) {
	o.KeywordsUsed = v
}

// GetMatches returns the Matches field value.
func (o *PickRemediationFromInvestigationResponse) GetMatches() []ActionMatch {
	if o == nil {
		var ret []ActionMatch
		return ret
	}
	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationResponse) GetMatchesOk() (*[]ActionMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matches, true
}

// SetMatches sets field value.
func (o *PickRemediationFromInvestigationResponse) SetMatches(v []ActionMatch) {
	o.Matches = v
}

// GetRequestId returns the RequestId field value.
func (o *PickRemediationFromInvestigationResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *PickRemediationFromInvestigationResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetTotalMatches returns the TotalMatches field value.
func (o *PickRemediationFromInvestigationResponse) GetTotalMatches() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalMatches
}

// GetTotalMatchesOk returns a tuple with the TotalMatches field value
// and a boolean to check if the value has been set.
func (o *PickRemediationFromInvestigationResponse) GetTotalMatchesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMatches, true
}

// SetTotalMatches sets field value.
func (o *PickRemediationFromInvestigationResponse) SetTotalMatches(v int64) {
	o.TotalMatches = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PickRemediationFromInvestigationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["keywords_used"] = o.KeywordsUsed
	toSerialize["matches"] = o.Matches
	toSerialize["request_id"] = o.RequestId
	toSerialize["total_matches"] = o.TotalMatches

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PickRemediationFromInvestigationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KeywordsUsed *[]string      `json:"keywords_used"`
		Matches      *[]ActionMatch `json:"matches"`
		RequestId    *string        `json:"request_id"`
		TotalMatches *int64         `json:"total_matches"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.KeywordsUsed == nil {
		return fmt.Errorf("required field keywords_used missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"keywords_used", "matches", "request_id", "total_matches"})
	} else {
		return err
	}
	o.KeywordsUsed = *all.KeywordsUsed
	o.Matches = *all.Matches
	o.RequestId = *all.RequestId
	o.TotalMatches = *all.TotalMatches

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
