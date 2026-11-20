// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupFindings Counts of the group's open findings. The counts include every finding in the group, including findings that
// don't match `filter[query]`.
type RemediationGroupFindings struct {
	// Number of the group's findings for each value of their `@severity` attribute. The response includes every
	// level, including levels with a count of zero.
	BySeverity RemediationGroupFindingsBySeverity `json:"by_severity"`
	// Number of open findings in the group.
	Total int64 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupFindings instantiates a new RemediationGroupFindings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupFindings(bySeverity RemediationGroupFindingsBySeverity, total int64) *RemediationGroupFindings {
	this := RemediationGroupFindings{}
	this.BySeverity = bySeverity
	this.Total = total
	return &this
}

// NewRemediationGroupFindingsWithDefaults instantiates a new RemediationGroupFindings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupFindingsWithDefaults() *RemediationGroupFindings {
	this := RemediationGroupFindings{}
	return &this
}

// GetBySeverity returns the BySeverity field value.
func (o *RemediationGroupFindings) GetBySeverity() RemediationGroupFindingsBySeverity {
	if o == nil {
		var ret RemediationGroupFindingsBySeverity
		return ret
	}
	return o.BySeverity
}

// GetBySeverityOk returns a tuple with the BySeverity field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindings) GetBySeverityOk() (*RemediationGroupFindingsBySeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BySeverity, true
}

// SetBySeverity sets field value.
func (o *RemediationGroupFindings) SetBySeverity(v RemediationGroupFindingsBySeverity) {
	o.BySeverity = v
}

// GetTotal returns the Total field value.
func (o *RemediationGroupFindings) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindings) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *RemediationGroupFindings) SetTotal(v int64) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupFindings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["by_severity"] = o.BySeverity
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupFindings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BySeverity *RemediationGroupFindingsBySeverity `json:"by_severity"`
		Total      *int64                              `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BySeverity == nil {
		return fmt.Errorf("required field by_severity missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"by_severity", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.BySeverity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BySeverity = *all.BySeverity
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
