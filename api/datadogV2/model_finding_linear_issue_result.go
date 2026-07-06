// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingLinearIssueResult Result of the Linear issue creation.
type FindingLinearIssueResult struct {
	// Account ID of the Linear workspace.
	AccountId *string `json:"account_id,omitempty"`
	// Unique identifier of the Linear issue.
	IssueId *string `json:"issue_id,omitempty"`
	// Key of the Linear issue.
	IssueKey *string `json:"issue_key,omitempty"`
	// Team ID of the Linear issue.
	TeamId *string `json:"team_id,omitempty"`
	// URL of the Linear issue.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFindingLinearIssueResult instantiates a new FindingLinearIssueResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingLinearIssueResult() *FindingLinearIssueResult {
	this := FindingLinearIssueResult{}
	return &this
}

// NewFindingLinearIssueResultWithDefaults instantiates a new FindingLinearIssueResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingLinearIssueResultWithDefaults() *FindingLinearIssueResult {
	this := FindingLinearIssueResult{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *FindingLinearIssueResult) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingLinearIssueResult) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *FindingLinearIssueResult) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *FindingLinearIssueResult) SetAccountId(v string) {
	o.AccountId = &v
}

// GetIssueId returns the IssueId field value if set, zero value otherwise.
func (o *FindingLinearIssueResult) GetIssueId() string {
	if o == nil || o.IssueId == nil {
		var ret string
		return ret
	}
	return *o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingLinearIssueResult) GetIssueIdOk() (*string, bool) {
	if o == nil || o.IssueId == nil {
		return nil, false
	}
	return o.IssueId, true
}

// HasIssueId returns a boolean if a field has been set.
func (o *FindingLinearIssueResult) HasIssueId() bool {
	return o != nil && o.IssueId != nil
}

// SetIssueId gets a reference to the given string and assigns it to the IssueId field.
func (o *FindingLinearIssueResult) SetIssueId(v string) {
	o.IssueId = &v
}

// GetIssueKey returns the IssueKey field value if set, zero value otherwise.
func (o *FindingLinearIssueResult) GetIssueKey() string {
	if o == nil || o.IssueKey == nil {
		var ret string
		return ret
	}
	return *o.IssueKey
}

// GetIssueKeyOk returns a tuple with the IssueKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingLinearIssueResult) GetIssueKeyOk() (*string, bool) {
	if o == nil || o.IssueKey == nil {
		return nil, false
	}
	return o.IssueKey, true
}

// HasIssueKey returns a boolean if a field has been set.
func (o *FindingLinearIssueResult) HasIssueKey() bool {
	return o != nil && o.IssueKey != nil
}

// SetIssueKey gets a reference to the given string and assigns it to the IssueKey field.
func (o *FindingLinearIssueResult) SetIssueKey(v string) {
	o.IssueKey = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *FindingLinearIssueResult) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingLinearIssueResult) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *FindingLinearIssueResult) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *FindingLinearIssueResult) SetTeamId(v string) {
	o.TeamId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FindingLinearIssueResult) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingLinearIssueResult) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FindingLinearIssueResult) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FindingLinearIssueResult) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingLinearIssueResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.IssueId != nil {
		toSerialize["issue_id"] = o.IssueId
	}
	if o.IssueKey != nil {
		toSerialize["issue_key"] = o.IssueKey
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingLinearIssueResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string `json:"account_id,omitempty"`
		IssueId   *string `json:"issue_id,omitempty"`
		IssueKey  *string `json:"issue_key,omitempty"`
		TeamId    *string `json:"team_id,omitempty"`
		Url       *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "issue_id", "issue_key", "team_id", "url"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.IssueId = all.IssueId
	o.IssueKey = all.IssueKey
	o.TeamId = all.TeamId
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
