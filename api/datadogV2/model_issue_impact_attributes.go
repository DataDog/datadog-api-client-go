// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactAttributes The issue's errors in the time window, and the users, sessions, accounts, views, and endpoints they affect.
type IssueImpactAttributes struct {
	// Accounts whose users are affected by the issue.
	Accounts IssueImpactAccounts `json:"accounts"`
	// Endpoints whose traces have the issue's errors, identified by each trace's root span.
	Endpoints IssueImpactEndpoints `json:"endpoints"`
	// Number of errors reported by each source. Several sources can report the same error, such as an APM span and
	// a log, so the counts overlap: don't add them together.
	ErrorCounts IssueImpactErrorCounts `json:"error_counts"`
	// RUM sessions where the issue occurred.
	Sessions IssueImpactSessions `json:"sessions"`
	// Users affected by the issue.
	Users IssueImpactUsers `json:"users"`
	// RUM views where the issue occurred.
	Views IssueImpactViews `json:"views"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactAttributes instantiates a new IssueImpactAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactAttributes(accounts IssueImpactAccounts, endpoints IssueImpactEndpoints, errorCounts IssueImpactErrorCounts, sessions IssueImpactSessions, users IssueImpactUsers, views IssueImpactViews) *IssueImpactAttributes {
	this := IssueImpactAttributes{}
	this.Accounts = accounts
	this.Endpoints = endpoints
	this.ErrorCounts = errorCounts
	this.Sessions = sessions
	this.Users = users
	this.Views = views
	return &this
}

// NewIssueImpactAttributesWithDefaults instantiates a new IssueImpactAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactAttributesWithDefaults() *IssueImpactAttributes {
	this := IssueImpactAttributes{}
	return &this
}

// GetAccounts returns the Accounts field value.
func (o *IssueImpactAttributes) GetAccounts() IssueImpactAccounts {
	if o == nil {
		var ret IssueImpactAccounts
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAttributes) GetAccountsOk() (*IssueImpactAccounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value.
func (o *IssueImpactAttributes) SetAccounts(v IssueImpactAccounts) {
	o.Accounts = v
}

// GetEndpoints returns the Endpoints field value.
func (o *IssueImpactAttributes) GetEndpoints() IssueImpactEndpoints {
	if o == nil {
		var ret IssueImpactEndpoints
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAttributes) GetEndpointsOk() (*IssueImpactEndpoints, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoints, true
}

// SetEndpoints sets field value.
func (o *IssueImpactAttributes) SetEndpoints(v IssueImpactEndpoints) {
	o.Endpoints = v
}

// GetErrorCounts returns the ErrorCounts field value.
func (o *IssueImpactAttributes) GetErrorCounts() IssueImpactErrorCounts {
	if o == nil {
		var ret IssueImpactErrorCounts
		return ret
	}
	return o.ErrorCounts
}

// GetErrorCountsOk returns a tuple with the ErrorCounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAttributes) GetErrorCountsOk() (*IssueImpactErrorCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCounts, true
}

// SetErrorCounts sets field value.
func (o *IssueImpactAttributes) SetErrorCounts(v IssueImpactErrorCounts) {
	o.ErrorCounts = v
}

// GetSessions returns the Sessions field value.
func (o *IssueImpactAttributes) GetSessions() IssueImpactSessions {
	if o == nil {
		var ret IssueImpactSessions
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAttributes) GetSessionsOk() (*IssueImpactSessions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *IssueImpactAttributes) SetSessions(v IssueImpactSessions) {
	o.Sessions = v
}

// GetUsers returns the Users field value.
func (o *IssueImpactAttributes) GetUsers() IssueImpactUsers {
	if o == nil {
		var ret IssueImpactUsers
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAttributes) GetUsersOk() (*IssueImpactUsers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value.
func (o *IssueImpactAttributes) SetUsers(v IssueImpactUsers) {
	o.Users = v
}

// GetViews returns the Views field value.
func (o *IssueImpactAttributes) GetViews() IssueImpactViews {
	if o == nil {
		var ret IssueImpactViews
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value
// and a boolean to check if the value has been set.
func (o *IssueImpactAttributes) GetViewsOk() (*IssueImpactViews, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Views, true
}

// SetViews sets field value.
func (o *IssueImpactAttributes) SetViews(v IssueImpactViews) {
	o.Views = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["accounts"] = o.Accounts
	toSerialize["endpoints"] = o.Endpoints
	toSerialize["error_counts"] = o.ErrorCounts
	toSerialize["sessions"] = o.Sessions
	toSerialize["users"] = o.Users
	toSerialize["views"] = o.Views

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Accounts    *IssueImpactAccounts    `json:"accounts"`
		Endpoints   *IssueImpactEndpoints   `json:"endpoints"`
		ErrorCounts *IssueImpactErrorCounts `json:"error_counts"`
		Sessions    *IssueImpactSessions    `json:"sessions"`
		Users       *IssueImpactUsers       `json:"users"`
		Views       *IssueImpactViews       `json:"views"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Accounts == nil {
		return fmt.Errorf("required field accounts missing")
	}
	if all.Endpoints == nil {
		return fmt.Errorf("required field endpoints missing")
	}
	if all.ErrorCounts == nil {
		return fmt.Errorf("required field error_counts missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"accounts", "endpoints", "error_counts", "sessions", "users", "views"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Accounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Accounts = *all.Accounts
	if all.Endpoints.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Endpoints = *all.Endpoints
	if all.ErrorCounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorCounts = *all.ErrorCounts
	if all.Sessions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sessions = *all.Sessions
	if all.Users.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Users = *all.Users
	if all.Views.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Views = *all.Views

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
