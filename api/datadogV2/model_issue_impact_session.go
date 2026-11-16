// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactSession A RUM session where the issue occurred.
type IssueImpactSession struct {
	// Number of errors reported by each source. Several sources can report the same error, such as an APM span and
	// a log, so the counts overlap: don't add them together.
	ErrorCounts IssueImpactErrorCounts `json:"error_counts"`
	// Session identifier.
	Id string `json:"id"`
	// Whether the entry stands out as a likely driver of the issue, because it accounts for most of the issue's
	// errors or for far more errors than the other listed entries.
	Significant bool `json:"significant"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactSession instantiates a new IssueImpactSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactSession(errorCounts IssueImpactErrorCounts, id string, significant bool) *IssueImpactSession {
	this := IssueImpactSession{}
	this.ErrorCounts = errorCounts
	this.Id = id
	this.Significant = significant
	return &this
}

// NewIssueImpactSessionWithDefaults instantiates a new IssueImpactSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactSessionWithDefaults() *IssueImpactSession {
	this := IssueImpactSession{}
	return &this
}

// GetErrorCounts returns the ErrorCounts field value.
func (o *IssueImpactSession) GetErrorCounts() IssueImpactErrorCounts {
	if o == nil {
		var ret IssueImpactErrorCounts
		return ret
	}
	return o.ErrorCounts
}

// GetErrorCountsOk returns a tuple with the ErrorCounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactSession) GetErrorCountsOk() (*IssueImpactErrorCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCounts, true
}

// SetErrorCounts sets field value.
func (o *IssueImpactSession) SetErrorCounts(v IssueImpactErrorCounts) {
	o.ErrorCounts = v
}

// GetId returns the Id field value.
func (o *IssueImpactSession) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IssueImpactSession) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IssueImpactSession) SetId(v string) {
	o.Id = v
}

// GetSignificant returns the Significant field value.
func (o *IssueImpactSession) GetSignificant() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Significant
}

// GetSignificantOk returns a tuple with the Significant field value
// and a boolean to check if the value has been set.
func (o *IssueImpactSession) GetSignificantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Significant, true
}

// SetSignificant sets field value.
func (o *IssueImpactSession) SetSignificant(v bool) {
	o.Significant = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_counts"] = o.ErrorCounts
	toSerialize["id"] = o.Id
	toSerialize["significant"] = o.Significant

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorCounts *IssueImpactErrorCounts `json:"error_counts"`
		Id          *string                 `json:"id"`
		Significant *bool                   `json:"significant"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ErrorCounts == nil {
		return fmt.Errorf("required field error_counts missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Significant == nil {
		return fmt.Errorf("required field significant missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_counts", "id", "significant"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ErrorCounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorCounts = *all.ErrorCounts
	o.Id = *all.Id
	o.Significant = *all.Significant

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
