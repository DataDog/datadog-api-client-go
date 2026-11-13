// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactAccounts Accounts whose users are affected by the issue.
type IssueImpactAccounts struct {
	// Number of distinct accounts affected by the issue in the time window.
	DistinctCount int64 `json:"distinct_count"`
	// Up to five accounts with the most errors, most affected first.
	Top []IssueImpactAccount `json:"top"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactAccounts instantiates a new IssueImpactAccounts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactAccounts(distinctCount int64, top []IssueImpactAccount) *IssueImpactAccounts {
	this := IssueImpactAccounts{}
	this.DistinctCount = distinctCount
	this.Top = top
	return &this
}

// NewIssueImpactAccountsWithDefaults instantiates a new IssueImpactAccounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactAccountsWithDefaults() *IssueImpactAccounts {
	this := IssueImpactAccounts{}
	return &this
}

// GetDistinctCount returns the DistinctCount field value.
func (o *IssueImpactAccounts) GetDistinctCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DistinctCount
}

// GetDistinctCountOk returns a tuple with the DistinctCount field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAccounts) GetDistinctCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistinctCount, true
}

// SetDistinctCount sets field value.
func (o *IssueImpactAccounts) SetDistinctCount(v int64) {
	o.DistinctCount = v
}

// GetTop returns the Top field value.
func (o *IssueImpactAccounts) GetTop() []IssueImpactAccount {
	if o == nil {
		var ret []IssueImpactAccount
		return ret
	}
	return o.Top
}

// GetTopOk returns a tuple with the Top field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAccounts) GetTopOk() (*[]IssueImpactAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Top, true
}

// SetTop sets field value.
func (o *IssueImpactAccounts) SetTop(v []IssueImpactAccount) {
	o.Top = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactAccounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["distinct_count"] = o.DistinctCount
	toSerialize["top"] = o.Top

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactAccounts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DistinctCount *int64                `json:"distinct_count"`
		Top           *[]IssueImpactAccount `json:"top"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DistinctCount == nil {
		return fmt.Errorf("required field distinct_count missing")
	}
	if all.Top == nil {
		return fmt.Errorf("required field top missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"distinct_count", "top"})
	} else {
		return err
	}
	o.DistinctCount = *all.DistinctCount
	o.Top = *all.Top

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
