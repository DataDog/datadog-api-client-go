// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyDropOffResponseAttributes Attributes of a journey drop-off analysis response.
type ProductAnalyticsJourneyDropOffResponseAttributes struct {
	// Number of sessions that reached the step before the target but not the target step.
	DropOffCount int64 `json:"drop_off_count"`
	// The issues most associated with the drop-off.
	Issues []ProductAnalyticsJourneyDropOffIssue `json:"issues"`
	// Number of dropped-off sessions that contained an error.
	SessionsWithErrors int64 `json:"sessions_with_errors"`
	// Number of dropped-off sessions that contained a frustration signal.
	SessionsWithFrustrationSignals int64 `json:"sessions_with_frustration_signals"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyDropOffResponseAttributes instantiates a new ProductAnalyticsJourneyDropOffResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyDropOffResponseAttributes(dropOffCount int64, issues []ProductAnalyticsJourneyDropOffIssue, sessionsWithErrors int64, sessionsWithFrustrationSignals int64) *ProductAnalyticsJourneyDropOffResponseAttributes {
	this := ProductAnalyticsJourneyDropOffResponseAttributes{}
	this.DropOffCount = dropOffCount
	this.Issues = issues
	this.SessionsWithErrors = sessionsWithErrors
	this.SessionsWithFrustrationSignals = sessionsWithFrustrationSignals
	return &this
}

// NewProductAnalyticsJourneyDropOffResponseAttributesWithDefaults instantiates a new ProductAnalyticsJourneyDropOffResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyDropOffResponseAttributesWithDefaults() *ProductAnalyticsJourneyDropOffResponseAttributes {
	this := ProductAnalyticsJourneyDropOffResponseAttributes{}
	return &this
}

// GetDropOffCount returns the DropOffCount field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetDropOffCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DropOffCount
}

// GetDropOffCountOk returns a tuple with the DropOffCount field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetDropOffCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DropOffCount, true
}

// SetDropOffCount sets field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) SetDropOffCount(v int64) {
	o.DropOffCount = v
}

// GetIssues returns the Issues field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetIssues() []ProductAnalyticsJourneyDropOffIssue {
	if o == nil {
		var ret []ProductAnalyticsJourneyDropOffIssue
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetIssuesOk() (*[]ProductAnalyticsJourneyDropOffIssue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issues, true
}

// SetIssues sets field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) SetIssues(v []ProductAnalyticsJourneyDropOffIssue) {
	o.Issues = v
}

// GetSessionsWithErrors returns the SessionsWithErrors field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetSessionsWithErrors() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionsWithErrors
}

// GetSessionsWithErrorsOk returns a tuple with the SessionsWithErrors field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetSessionsWithErrorsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsWithErrors, true
}

// SetSessionsWithErrors sets field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) SetSessionsWithErrors(v int64) {
	o.SessionsWithErrors = v
}

// GetSessionsWithFrustrationSignals returns the SessionsWithFrustrationSignals field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetSessionsWithFrustrationSignals() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionsWithFrustrationSignals
}

// GetSessionsWithFrustrationSignalsOk returns a tuple with the SessionsWithFrustrationSignals field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) GetSessionsWithFrustrationSignalsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionsWithFrustrationSignals, true
}

// SetSessionsWithFrustrationSignals sets field value.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) SetSessionsWithFrustrationSignals(v int64) {
	o.SessionsWithFrustrationSignals = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyDropOffResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["drop_off_count"] = o.DropOffCount
	toSerialize["issues"] = o.Issues
	toSerialize["sessions_with_errors"] = o.SessionsWithErrors
	toSerialize["sessions_with_frustration_signals"] = o.SessionsWithFrustrationSignals

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyDropOffResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DropOffCount                   *int64                                 `json:"drop_off_count"`
		Issues                         *[]ProductAnalyticsJourneyDropOffIssue `json:"issues"`
		SessionsWithErrors             *int64                                 `json:"sessions_with_errors"`
		SessionsWithFrustrationSignals *int64                                 `json:"sessions_with_frustration_signals"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DropOffCount == nil {
		return fmt.Errorf("required field drop_off_count missing")
	}
	if all.Issues == nil {
		return fmt.Errorf("required field issues missing")
	}
	if all.SessionsWithErrors == nil {
		return fmt.Errorf("required field sessions_with_errors missing")
	}
	if all.SessionsWithFrustrationSignals == nil {
		return fmt.Errorf("required field sessions_with_frustration_signals missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"drop_off_count", "issues", "sessions_with_errors", "sessions_with_frustration_signals"})
	} else {
		return err
	}
	o.DropOffCount = *all.DropOffCount
	o.Issues = *all.Issues
	o.SessionsWithErrors = *all.SessionsWithErrors
	o.SessionsWithFrustrationSignals = *all.SessionsWithFrustrationSignals

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
