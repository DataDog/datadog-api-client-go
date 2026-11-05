// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationsListMeta Metadata of a page of investigations.
type SecurityAgentInvestigationsListMeta struct {
	// Why `filter[signal_id]` returned no investigation. Set only on the first page of an empty result. The reason is
	// based on the organization's current settings. Poll again only on `investigation_not_started`, after the number
	// of seconds in the `Retry-After` header; treat the others as final for this signal.
	// - `investigation_not_started`: no investigation has started yet, but one may still start.
	// - `not_investigated`: no investigation started within an hour of the signal, so none is expected. The signal
	//   was skipped for a reason this API doesn't report, such as its severity or a rate limit.
	// - `rule_not_enabled`: the signal's detection rule doesn't have Bits Security Analyst enabled, so no
	//   investigation starts automatically.
	// - `signal_not_found`: the ID isn't a signal ID, or no signal with this ID exists in the organization.
	// - `org_not_activated`: the organization hasn't activated Bits Security Analyst.
	Reason *SecurityAgentInvestigationNoRunReason `json:"reason,omitempty"`
	// Without `filter[signal_id]`, the number of signals that match the filters. With `filter[signal_id]`, the
	// number of investigations of that signal. Pages past the last page return 0.
	TotalCount int64 `json:"total_count"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationsListMeta instantiates a new SecurityAgentInvestigationsListMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationsListMeta(totalCount int64) *SecurityAgentInvestigationsListMeta {
	this := SecurityAgentInvestigationsListMeta{}
	this.TotalCount = totalCount
	return &this
}

// NewSecurityAgentInvestigationsListMetaWithDefaults instantiates a new SecurityAgentInvestigationsListMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationsListMetaWithDefaults() *SecurityAgentInvestigationsListMeta {
	this := SecurityAgentInvestigationsListMeta{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SecurityAgentInvestigationsListMeta) GetReason() SecurityAgentInvestigationNoRunReason {
	if o == nil || o.Reason == nil {
		var ret SecurityAgentInvestigationNoRunReason
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationsListMeta) GetReasonOk() (*SecurityAgentInvestigationNoRunReason, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SecurityAgentInvestigationsListMeta) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given SecurityAgentInvestigationNoRunReason and assigns it to the Reason field.
func (o *SecurityAgentInvestigationsListMeta) SetReason(v SecurityAgentInvestigationNoRunReason) {
	o.Reason = &v
}

// GetTotalCount returns the TotalCount field value.
func (o *SecurityAgentInvestigationsListMeta) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationsListMeta) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *SecurityAgentInvestigationsListMeta) SetTotalCount(v int64) {
	o.TotalCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationsListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["total_count"] = o.TotalCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationsListMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason     *SecurityAgentInvestigationNoRunReason `json:"reason,omitempty"`
		TotalCount *int64                                 `json:"total_count"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field total_count missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"reason", "total_count"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Reason != nil && !all.Reason.IsValid() {
		hasInvalidField = true
	} else {
		o.Reason = all.Reason
	}
	o.TotalCount = *all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
