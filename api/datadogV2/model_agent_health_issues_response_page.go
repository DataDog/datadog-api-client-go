// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssuesResponsePage Counts for the matching Agent Health issues.
type AgentHealthIssuesResponsePage struct {
	// The number of open issues matching the filters in each issue-type group.
	CountsByIssueType map[string]int64 `json:"counts_by_issue_type"`
	// The total number of open issues matching the filters.
	TotalCount int64 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAgentHealthIssuesResponsePage instantiates a new AgentHealthIssuesResponsePage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAgentHealthIssuesResponsePage(countsByIssueType map[string]int64, totalCount int64) *AgentHealthIssuesResponsePage {
	this := AgentHealthIssuesResponsePage{}
	this.CountsByIssueType = countsByIssueType
	this.TotalCount = totalCount
	return &this
}

// NewAgentHealthIssuesResponsePageWithDefaults instantiates a new AgentHealthIssuesResponsePage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAgentHealthIssuesResponsePageWithDefaults() *AgentHealthIssuesResponsePage {
	this := AgentHealthIssuesResponsePage{}
	return &this
}

// GetCountsByIssueType returns the CountsByIssueType field value.
func (o *AgentHealthIssuesResponsePage) GetCountsByIssueType() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.CountsByIssueType
}

// GetCountsByIssueTypeOk returns a tuple with the CountsByIssueType field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssuesResponsePage) GetCountsByIssueTypeOk() (*map[string]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountsByIssueType, true
}

// SetCountsByIssueType sets field value.
func (o *AgentHealthIssuesResponsePage) SetCountsByIssueType(v map[string]int64) {
	o.CountsByIssueType = v
}

// GetTotalCount returns the TotalCount field value.
func (o *AgentHealthIssuesResponsePage) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssuesResponsePage) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *AgentHealthIssuesResponsePage) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AgentHealthIssuesResponsePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["counts_by_issue_type"] = o.CountsByIssueType
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthIssuesResponsePage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CountsByIssueType *map[string]int64 `json:"counts_by_issue_type"`
		TotalCount        *int64            `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CountsByIssueType == nil {
		return fmt.Errorf("required field counts_by_issue_type missing")
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"counts_by_issue_type", "total_count"})
	} else {
		return err
	}
	o.CountsByIssueType = *all.CountsByIssueType
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
