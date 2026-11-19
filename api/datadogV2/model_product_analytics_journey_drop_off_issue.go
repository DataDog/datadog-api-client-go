// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDropOffIssue An issue contributing to drop-off at the target step.
type ProductAnalyticsJourneyDropOffIssue struct {
	// Number of errors attributed to this issue.
	ErrorCount int64 `json:"error_count"`
	// Identifier of the issue.
	IssueId string `json:"issue_id"`
	// Number of sessions affected by this issue.
	SessionCount int64 `json:"session_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyDropOffIssue instantiates a new ProductAnalyticsJourneyDropOffIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyDropOffIssue(errorCount int64, issueId string, sessionCount int64) *ProductAnalyticsJourneyDropOffIssue {
	this := ProductAnalyticsJourneyDropOffIssue{}
	this.ErrorCount = errorCount
	this.IssueId = issueId
	this.SessionCount = sessionCount
	return &this
}

// NewProductAnalyticsJourneyDropOffIssueWithDefaults instantiates a new ProductAnalyticsJourneyDropOffIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyDropOffIssueWithDefaults() *ProductAnalyticsJourneyDropOffIssue {
	this := ProductAnalyticsJourneyDropOffIssue{}
	return &this
}

// GetErrorCount returns the ErrorCount field value.
func (o *ProductAnalyticsJourneyDropOffIssue) GetErrorCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffIssue) GetErrorCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCount, true
}

// SetErrorCount sets field value.
func (o *ProductAnalyticsJourneyDropOffIssue) SetErrorCount(v int64) {
	o.ErrorCount = v
}

// GetIssueId returns the IssueId field value.
func (o *ProductAnalyticsJourneyDropOffIssue) GetIssueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffIssue) GetIssueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value.
func (o *ProductAnalyticsJourneyDropOffIssue) SetIssueId(v string) {
	o.IssueId = v
}

// GetSessionCount returns the SessionCount field value.
func (o *ProductAnalyticsJourneyDropOffIssue) GetSessionCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffIssue) GetSessionCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionCount, true
}

// SetSessionCount sets field value.
func (o *ProductAnalyticsJourneyDropOffIssue) SetSessionCount(v int64) {
	o.SessionCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyDropOffIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_count"] = o.ErrorCount
	toSerialize["issue_id"] = o.IssueId
	toSerialize["session_count"] = o.SessionCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyDropOffIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorCount   *int64  `json:"error_count"`
		IssueId      *string `json:"issue_id"`
		SessionCount *int64  `json:"session_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ErrorCount == nil {
		return fmt.Errorf("required field error_count missing")
	}
	if all.IssueId == nil {
		return fmt.Errorf("required field issue_id missing")
	}
	if all.SessionCount == nil {
		return fmt.Errorf("required field session_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_count", "issue_id", "session_count"})
	} else {
		return err
	}
	o.ErrorCount = *all.ErrorCount
	o.IssueId = *all.IssueId
	o.SessionCount = *all.SessionCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
