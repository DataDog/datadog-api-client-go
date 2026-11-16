// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactUser A user affected by the issue.
type IssueImpactUser struct {
	// Number of errors reported by each source. Several sources can report the same error, such as an APM span and
	// a log, so the counts overlap: don't add them together.
	ErrorCounts IssueImpactErrorCounts `json:"error_counts"`
	// User ID, from the `usr.id` attribute.
	Id string `json:"id"`
	// User name or email, if available.
	Name *string `json:"name,omitempty"`
	// Whether the entry stands out as a likely driver of the issue, because it accounts for most of the issue's
	// errors or for far more errors than the other listed entries.
	Significant bool `json:"significant"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactUser instantiates a new IssueImpactUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactUser(errorCounts IssueImpactErrorCounts, id string, significant bool) *IssueImpactUser {
	this := IssueImpactUser{}
	this.ErrorCounts = errorCounts
	this.Id = id
	this.Significant = significant
	return &this
}

// NewIssueImpactUserWithDefaults instantiates a new IssueImpactUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactUserWithDefaults() *IssueImpactUser {
	this := IssueImpactUser{}
	return &this
}

// GetErrorCounts returns the ErrorCounts field value.
func (o *IssueImpactUser) GetErrorCounts() IssueImpactErrorCounts {
	if o == nil {
		var ret IssueImpactErrorCounts
		return ret
	}
	return o.ErrorCounts
}

// GetErrorCountsOk returns a tuple with the ErrorCounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactUser) GetErrorCountsOk() (*IssueImpactErrorCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCounts, true
}

// SetErrorCounts sets field value.
func (o *IssueImpactUser) SetErrorCounts(v IssueImpactErrorCounts) {
	o.ErrorCounts = v
}

// GetId returns the Id field value.
func (o *IssueImpactUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IssueImpactUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IssueImpactUser) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueImpactUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueImpactUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueImpactUser) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueImpactUser) SetName(v string) {
	o.Name = &v
}

// GetSignificant returns the Significant field value.
func (o *IssueImpactUser) GetSignificant() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Significant
}

// GetSignificantOk returns a tuple with the Significant field value
// and a boolean to check if the value has been set.
func (o *IssueImpactUser) GetSignificantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Significant, true
}

// SetSignificant sets field value.
func (o *IssueImpactUser) SetSignificant(v bool) {
	o.Significant = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactUser) MarshalJSON() ([]byte, error) {
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
func (o *IssueImpactUser) UnmarshalJSON(bytes []byte) (err error) {
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
