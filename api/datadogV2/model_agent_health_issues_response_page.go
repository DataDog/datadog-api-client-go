// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AgentHealthIssuesResponsePage Pagination details and counts for the matching Agent Health issues.
type AgentHealthIssuesResponsePage struct {
	// The number of open issues matching the filters in each issue-type group.
	CountsByIssueType map[string]int64 `json:"counts_by_issue_type,omitempty"`
	// The maximum number of issues returned, echoing the requested `page[limit]`.
	Limit int64 `json:"limit"`
	// The number of issues skipped, echoing the requested `page[offset]`.
	Offset int64 `json:"offset"`
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
func NewAgentHealthIssuesResponsePage(limit int64, offset int64, totalCount int64) *AgentHealthIssuesResponsePage {
	this := AgentHealthIssuesResponsePage{}
	this.Limit = limit
	this.Offset = offset
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

// GetCountsByIssueType returns the CountsByIssueType field value if set, zero value otherwise.
func (o *AgentHealthIssuesResponsePage) GetCountsByIssueType() map[string]int64 {
	if o == nil || o.CountsByIssueType == nil {
		var ret map[string]int64
		return ret
	}
	return o.CountsByIssueType
}

// GetCountsByIssueTypeOk returns a tuple with the CountsByIssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentHealthIssuesResponsePage) GetCountsByIssueTypeOk() (*map[string]int64, bool) {
	if o == nil || o.CountsByIssueType == nil {
		return nil, false
	}
	return &o.CountsByIssueType, true
}

// HasCountsByIssueType returns a boolean if a field has been set.
func (o *AgentHealthIssuesResponsePage) HasCountsByIssueType() bool {
	return o != nil && o.CountsByIssueType != nil
}

// SetCountsByIssueType gets a reference to the given map[string]int64 and assigns it to the CountsByIssueType field.
func (o *AgentHealthIssuesResponsePage) SetCountsByIssueType(v map[string]int64) {
	o.CountsByIssueType = v
}

// GetLimit returns the Limit field value.
func (o *AgentHealthIssuesResponsePage) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssuesResponsePage) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *AgentHealthIssuesResponsePage) SetLimit(v int64) {
	o.Limit = v
}

// GetOffset returns the Offset field value.
func (o *AgentHealthIssuesResponsePage) GetOffset() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *AgentHealthIssuesResponsePage) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value.
func (o *AgentHealthIssuesResponsePage) SetOffset(v int64) {
	o.Offset = v
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
	if o.CountsByIssueType != nil {
		toSerialize["counts_by_issue_type"] = o.CountsByIssueType
	}
	toSerialize["limit"] = o.Limit
	toSerialize["offset"] = o.Offset
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AgentHealthIssuesResponsePage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CountsByIssueType map[string]int64 `json:"counts_by_issue_type,omitempty"`
		Limit             *int64           `json:"limit"`
		Offset            *int64           `json:"offset"`
		TotalCount        *int64           `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.Offset == nil {
		return fmt.Errorf("required field offset missing")
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"counts_by_issue_type", "limit", "offset", "total_count"})
	} else {
		return err
	}
	o.CountsByIssueType = all.CountsByIssueType
	o.Limit = *all.Limit
	o.Offset = *all.Offset
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
