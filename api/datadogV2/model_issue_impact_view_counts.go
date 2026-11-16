// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactViewCounts Distinct users, accounts, sessions, and view instances on a view. `affected` counts only those with the issue's
// errors. `baseline` counts all users, accounts, sessions, and view instances on the view in the time window. It
// is absent when RUM has no data for the view.
type IssueImpactViewCounts struct {
	// Number of distinct accounts.
	Accounts int64 `json:"accounts"`
	// Number of distinct sessions.
	Sessions int64 `json:"sessions"`
	// Number of distinct users.
	Users int64 `json:"users"`
	// Number of distinct view instances.
	Views int64 `json:"views"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactViewCounts instantiates a new IssueImpactViewCounts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactViewCounts(accounts int64, sessions int64, users int64, views int64) *IssueImpactViewCounts {
	this := IssueImpactViewCounts{}
	this.Accounts = accounts
	this.Sessions = sessions
	this.Users = users
	this.Views = views
	return &this
}

// NewIssueImpactViewCountsWithDefaults instantiates a new IssueImpactViewCounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactViewCountsWithDefaults() *IssueImpactViewCounts {
	this := IssueImpactViewCounts{}
	return &this
}

// GetAccounts returns the Accounts field value.
func (o *IssueImpactViewCounts) GetAccounts() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactViewCounts) GetAccountsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value.
func (o *IssueImpactViewCounts) SetAccounts(v int64) {
	o.Accounts = v
}

// GetSessions returns the Sessions field value.
func (o *IssueImpactViewCounts) GetSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *IssueImpactViewCounts) GetSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *IssueImpactViewCounts) SetSessions(v int64) {
	o.Sessions = v
}

// GetUsers returns the Users field value.
func (o *IssueImpactViewCounts) GetUsers() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *IssueImpactViewCounts) GetUsersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value.
func (o *IssueImpactViewCounts) SetUsers(v int64) {
	o.Users = v
}

// GetViews returns the Views field value.
func (o *IssueImpactViewCounts) GetViews() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value
// and a boolean to check if the value has been set.
func (o *IssueImpactViewCounts) GetViewsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Views, true
}

// SetViews sets field value.
func (o *IssueImpactViewCounts) SetViews(v int64) {
	o.Views = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactViewCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["accounts"] = o.Accounts
	toSerialize["sessions"] = o.Sessions
	toSerialize["users"] = o.Users
	toSerialize["views"] = o.Views

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactViewCounts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts *int64 `json:"accounts"`
		Sessions *int64 `json:"sessions"`
		Users    *int64 `json:"users"`
		Views    *int64 `json:"views"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Accounts == nil {
		return fmt.Errorf("required field accounts missing")
	}
	if all.Sessions == nil {
		return fmt.Errorf("required field sessions missing")
	}
	if all.Users == nil {
		return fmt.Errorf("required field users missing")
	}
	if all.Views == nil {
		return fmt.Errorf("required field views missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accounts", "sessions", "users", "views"})
	} else {
		return err
	}
	o.Accounts = *all.Accounts
	o.Sessions = *all.Sessions
	o.Users = *all.Users
	o.Views = *all.Views

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
