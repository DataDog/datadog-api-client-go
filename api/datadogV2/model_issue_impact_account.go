// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactAccount An account whose users are affected by the issue.
type IssueImpactAccount struct {
	// Number of errors reported by each source. Several sources can report the same error, such as an APM span and
	// a log, so the counts overlap: don't add them together.
	ErrorCounts IssueImpactErrorCounts `json:"error_counts"`
	// Account ID, from the `account.id` attribute.
	Id string `json:"id"`
	// Account name, if available.
	Name *string `json:"name,omitempty"`
	// Whether the entry stands out as a likely driver of the issue, because it accounts for most of the issue's
	// errors or for far more errors than the other listed entries.
	Significant bool `json:"significant"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactAccount instantiates a new IssueImpactAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactAccount(errorCounts IssueImpactErrorCounts, id string, significant bool) *IssueImpactAccount {
	this := IssueImpactAccount{}
	this.ErrorCounts = errorCounts
	this.Id = id
	this.Significant = significant
	return &this
}

// NewIssueImpactAccountWithDefaults instantiates a new IssueImpactAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactAccountWithDefaults() *IssueImpactAccount {
	this := IssueImpactAccount{}
	return &this
}

// GetErrorCounts returns the ErrorCounts field value.
func (o *IssueImpactAccount) GetErrorCounts() IssueImpactErrorCounts {
	if o == nil {
		var ret IssueImpactErrorCounts
		return ret
	}
	return o.ErrorCounts
}

// GetErrorCountsOk returns a tuple with the ErrorCounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAccount) GetErrorCountsOk() (*IssueImpactErrorCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCounts, true
}

// SetErrorCounts sets field value.
func (o *IssueImpactAccount) SetErrorCounts(v IssueImpactErrorCounts) {
	o.ErrorCounts = v
}

// GetId returns the Id field value.
func (o *IssueImpactAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IssueImpactAccount) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueImpactAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueImpactAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueImpactAccount) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueImpactAccount) SetName(v string) {
	o.Name = &v
}

// GetSignificant returns the Significant field value.
func (o *IssueImpactAccount) GetSignificant() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Significant
}

// GetSignificantOk returns a tuple with the Significant field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAccount) GetSignificantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Significant, true
}

// SetSignificant sets field value.
func (o *IssueImpactAccount) SetSignificant(v bool) {
	o.Significant = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["error_counts"] = o.ErrorCounts
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["significant"] = o.Significant

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorCounts *IssueImpactErrorCounts `json:"error_counts"`
		Id          *string                 `json:"id"`
		Name        *string                 `json:"name,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"error_counts", "id", "name", "significant"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ErrorCounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorCounts = *all.ErrorCounts
	o.Id = *all.Id
	o.Name = all.Name
	o.Significant = *all.Significant

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
