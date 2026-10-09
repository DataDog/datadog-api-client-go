// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideResult Per-finding outcome of a severity override operation.
type SeverityOverrideResult struct {
	// Human-readable explanation of why the finding was skipped.
	Detail string `json:"detail"`
	// Unique identifier of the security finding.
	FindingId string `json:"finding_id"`
	// HTTP-like status code describing the outcome for this finding.
	Status int32 `json:"status"`
	// Short label describing the outcome for this finding.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityOverrideResult instantiates a new SeverityOverrideResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityOverrideResult(detail string, findingId string, status int32, title string) *SeverityOverrideResult {
	this := SeverityOverrideResult{}
	this.Detail = detail
	this.FindingId = findingId
	this.Status = status
	this.Title = title
	return &this
}

// NewSeverityOverrideResultWithDefaults instantiates a new SeverityOverrideResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityOverrideResultWithDefaults() *SeverityOverrideResult {
	this := SeverityOverrideResult{}
	return &this
}

// GetDetail returns the Detail field value.
func (o *SeverityOverrideResult) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResult) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value.
func (o *SeverityOverrideResult) SetDetail(v string) {
	o.Detail = v
}

// GetFindingId returns the FindingId field value.
func (o *SeverityOverrideResult) GetFindingId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FindingId
}

// GetFindingIdOk returns a tuple with the FindingId field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResult) GetFindingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FindingId, true
}

// SetFindingId sets field value.
func (o *SeverityOverrideResult) SetFindingId(v string) {
	o.FindingId = v
}

// GetStatus returns the Status field value.
func (o *SeverityOverrideResult) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResult) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SeverityOverrideResult) SetStatus(v int32) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *SeverityOverrideResult) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResult) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *SeverityOverrideResult) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityOverrideResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["detail"] = o.Detail
	toSerialize["finding_id"] = o.FindingId
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityOverrideResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Detail    *string `json:"detail"`
		FindingId *string `json:"finding_id"`
		Status    *int32  `json:"status"`
		Title     *string `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Detail == nil {
		return fmt.Errorf("required field detail missing")
	}
	if all.FindingId == nil {
		return fmt.Errorf("required field finding_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detail", "finding_id", "status", "title"})
	} else {
		return err
	}
	o.Detail = *all.Detail
	o.FindingId = *all.FindingId
	o.Status = *all.Status
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
