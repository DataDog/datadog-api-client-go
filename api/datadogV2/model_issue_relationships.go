// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueRelationships Relationship between the issue and other resources, such as an assignee, case, teams, pull requests, watches, samples, or a Bits AI investigation.
type IssueRelationships struct {
	// Relationship between the issue and assignee.
	Assignee *IssueAssigneeRelationship `json:"assignee,omitempty"`
	// Bits AI investigation linked to the issue, returned when `bits_investigation` is requested
	// in the `include` query parameter.
	BitsInvestigation *IssueBitsInvestigationRelationship `json:"bits_investigation,omitempty"`
	// Relationship between the issue and case.
	Case *IssueCaseRelationship `json:"case,omitempty"`
	// Relationship between the issue and its linked pull requests. Returned when `pull_requests` is included and the issue has linked pull requests.
	PullRequests *IssuePullRequestsRelationship `json:"pull_requests,omitempty"`
	// Relationship between the issue and its curated error samples.
	Samples *IssueSamplesRelationship `json:"samples,omitempty"`
	// Relationship between the issue and teams.
	TeamOwners *IssueTeamOwnersRelationship `json:"team_owners,omitempty"`
	// Relationship between the issue and its active watches. Returned when `watches` is included and the issue has an active watch.
	Watches *IssueWatchesRelationship `json:"watches,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueRelationships instantiates a new IssueRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueRelationships() *IssueRelationships {
	this := IssueRelationships{}
	return &this
}

// NewIssueRelationshipsWithDefaults instantiates a new IssueRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueRelationshipsWithDefaults() *IssueRelationships {
	this := IssueRelationships{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *IssueRelationships) GetAssignee() IssueAssigneeRelationship {
	if o == nil || o.Assignee == nil {
		var ret IssueAssigneeRelationship
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetAssigneeOk() (*IssueAssigneeRelationship, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *IssueRelationships) HasAssignee() bool {
	return o != nil && o.Assignee != nil
}

// SetAssignee gets a reference to the given IssueAssigneeRelationship and assigns it to the Assignee field.
func (o *IssueRelationships) SetAssignee(v IssueAssigneeRelationship) {
	o.Assignee = &v
}

// GetBitsInvestigation returns the BitsInvestigation field value if set, zero value otherwise.
func (o *IssueRelationships) GetBitsInvestigation() IssueBitsInvestigationRelationship {
	if o == nil || o.BitsInvestigation == nil {
		var ret IssueBitsInvestigationRelationship
		return ret
	}
	return *o.BitsInvestigation
}

// GetBitsInvestigationOk returns a tuple with the BitsInvestigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetBitsInvestigationOk() (*IssueBitsInvestigationRelationship, bool) {
	if o == nil || o.BitsInvestigation == nil {
		return nil, false
	}
	return o.BitsInvestigation, true
}

// HasBitsInvestigation returns a boolean if a field has been set.
func (o *IssueRelationships) HasBitsInvestigation() bool {
	return o != nil && o.BitsInvestigation != nil
}

// SetBitsInvestigation gets a reference to the given IssueBitsInvestigationRelationship and assigns it to the BitsInvestigation field.
func (o *IssueRelationships) SetBitsInvestigation(v IssueBitsInvestigationRelationship) {
	o.BitsInvestigation = &v
}

// GetCase returns the Case field value if set, zero value otherwise.
func (o *IssueRelationships) GetCase() IssueCaseRelationship {
	if o == nil || o.Case == nil {
		var ret IssueCaseRelationship
		return ret
	}
	return *o.Case
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetCaseOk() (*IssueCaseRelationship, bool) {
	if o == nil || o.Case == nil {
		return nil, false
	}
	return o.Case, true
}

// HasCase returns a boolean if a field has been set.
func (o *IssueRelationships) HasCase() bool {
	return o != nil && o.Case != nil
}

// SetCase gets a reference to the given IssueCaseRelationship and assigns it to the Case field.
func (o *IssueRelationships) SetCase(v IssueCaseRelationship) {
	o.Case = &v
}

// GetPullRequests returns the PullRequests field value if set, zero value otherwise.
func (o *IssueRelationships) GetPullRequests() IssuePullRequestsRelationship {
	if o == nil || o.PullRequests == nil {
		var ret IssuePullRequestsRelationship
		return ret
	}
	return *o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetPullRequestsOk() (*IssuePullRequestsRelationship, bool) {
	if o == nil || o.PullRequests == nil {
		return nil, false
	}
	return o.PullRequests, true
}

// HasPullRequests returns a boolean if a field has been set.
func (o *IssueRelationships) HasPullRequests() bool {
	return o != nil && o.PullRequests != nil
}

// SetPullRequests gets a reference to the given IssuePullRequestsRelationship and assigns it to the PullRequests field.
func (o *IssueRelationships) SetPullRequests(v IssuePullRequestsRelationship) {
	o.PullRequests = &v
}

// GetSamples returns the Samples field value if set, zero value otherwise.
func (o *IssueRelationships) GetSamples() IssueSamplesRelationship {
	if o == nil || o.Samples == nil {
		var ret IssueSamplesRelationship
		return ret
	}
	return *o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetSamplesOk() (*IssueSamplesRelationship, bool) {
	if o == nil || o.Samples == nil {
		return nil, false
	}
	return o.Samples, true
}

// HasSamples returns a boolean if a field has been set.
func (o *IssueRelationships) HasSamples() bool {
	return o != nil && o.Samples != nil
}

// SetSamples gets a reference to the given IssueSamplesRelationship and assigns it to the Samples field.
func (o *IssueRelationships) SetSamples(v IssueSamplesRelationship) {
	o.Samples = &v
}

// GetTeamOwners returns the TeamOwners field value if set, zero value otherwise.
func (o *IssueRelationships) GetTeamOwners() IssueTeamOwnersRelationship {
	if o == nil || o.TeamOwners == nil {
		var ret IssueTeamOwnersRelationship
		return ret
	}
	return *o.TeamOwners
}

// GetTeamOwnersOk returns a tuple with the TeamOwners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetTeamOwnersOk() (*IssueTeamOwnersRelationship, bool) {
	if o == nil || o.TeamOwners == nil {
		return nil, false
	}
	return o.TeamOwners, true
}

// HasTeamOwners returns a boolean if a field has been set.
func (o *IssueRelationships) HasTeamOwners() bool {
	return o != nil && o.TeamOwners != nil
}

// SetTeamOwners gets a reference to the given IssueTeamOwnersRelationship and assigns it to the TeamOwners field.
func (o *IssueRelationships) SetTeamOwners(v IssueTeamOwnersRelationship) {
	o.TeamOwners = &v
}

// GetWatches returns the Watches field value if set, zero value otherwise.
func (o *IssueRelationships) GetWatches() IssueWatchesRelationship {
	if o == nil || o.Watches == nil {
		var ret IssueWatchesRelationship
		return ret
	}
	return *o.Watches
}

// GetWatchesOk returns a tuple with the Watches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRelationships) GetWatchesOk() (*IssueWatchesRelationship, bool) {
	if o == nil || o.Watches == nil {
		return nil, false
	}
	return o.Watches, true
}

// HasWatches returns a boolean if a field has been set.
func (o *IssueRelationships) HasWatches() bool {
	return o != nil && o.Watches != nil
}

// SetWatches gets a reference to the given IssueWatchesRelationship and assigns it to the Watches field.
func (o *IssueRelationships) SetWatches(v IssueWatchesRelationship) {
	o.Watches = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.BitsInvestigation != nil {
		toSerialize["bits_investigation"] = o.BitsInvestigation
	}
	if o.Case != nil {
		toSerialize["case"] = o.Case
	}
	if o.PullRequests != nil {
		toSerialize["pull_requests"] = o.PullRequests
	}
	if o.Samples != nil {
		toSerialize["samples"] = o.Samples
	}
	if o.TeamOwners != nil {
		toSerialize["team_owners"] = o.TeamOwners
	}
	if o.Watches != nil {
		toSerialize["watches"] = o.Watches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignee          *IssueAssigneeRelationship          `json:"assignee,omitempty"`
		BitsInvestigation *IssueBitsInvestigationRelationship `json:"bits_investigation,omitempty"`
		Case              *IssueCaseRelationship              `json:"case,omitempty"`
		PullRequests      *IssuePullRequestsRelationship      `json:"pull_requests,omitempty"`
		Samples           *IssueSamplesRelationship           `json:"samples,omitempty"`
		TeamOwners        *IssueTeamOwnersRelationship        `json:"team_owners,omitempty"`
		Watches           *IssueWatchesRelationship           `json:"watches,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee", "bits_investigation", "case", "pull_requests", "samples", "team_owners", "watches"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assignee != nil && all.Assignee.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Assignee = all.Assignee
	if all.BitsInvestigation != nil && all.BitsInvestigation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BitsInvestigation = all.BitsInvestigation
	if all.Case != nil && all.Case.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Case = all.Case
	if all.PullRequests != nil && all.PullRequests.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PullRequests = all.PullRequests
	if all.Samples != nil && all.Samples.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Samples = all.Samples
	if all.TeamOwners != nil && all.TeamOwners.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TeamOwners = all.TeamOwners
	if all.Watches != nil && all.Watches.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Watches = all.Watches

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
