// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseJiraIssue Jira issue associated with the case.
type CaseJiraIssue struct {
	// The error message if the Jira issue creation failed.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Result of the Jira issue creation.
	Result *CaseJiraIssueResult `json:"result,omitempty"`
	// The status of the Jira issue creation. Can be "COMPLETED" if the Jira issue was created successfully, or "FAILED" if the Jira issue creation failed.
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseJiraIssue instantiates a new CaseJiraIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseJiraIssue() *CaseJiraIssue {
	this := CaseJiraIssue{}
	return &this
}

// NewCaseJiraIssueWithDefaults instantiates a new CaseJiraIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseJiraIssueWithDefaults() *CaseJiraIssue {
	this := CaseJiraIssue{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CaseJiraIssue) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseJiraIssue) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CaseJiraIssue) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CaseJiraIssue) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CaseJiraIssue) GetResult() CaseJiraIssueResult {
	if o == nil || o.Result == nil {
		var ret CaseJiraIssueResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseJiraIssue) GetResultOk() (*CaseJiraIssueResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CaseJiraIssue) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given CaseJiraIssueResult and assigns it to the Result field.
func (o *CaseJiraIssue) SetResult(v CaseJiraIssueResult) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CaseJiraIssue) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseJiraIssue) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CaseJiraIssue) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CaseJiraIssue) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseJiraIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseJiraIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorMessage *string              `json:"error_message,omitempty"`
		Result       *CaseJiraIssueResult `json:"result,omitempty"`
		Status       *string              `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_message", "result", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ErrorMessage = all.ErrorMessage
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
