// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactErrorCounts Number of errors reported by each source. Several sources can report the same error, such as an APM span and
// a log, so the counts overlap: don't add them together.
type IssueImpactErrorCounts struct {
	// Number of errors reported by APM.
	Apm int64 `json:"apm"`
	// Number of errors reported by Agent Observability.
	Llmobs int64 `json:"llmobs"`
	// Number of errors reported by Log Management.
	Logs int64 `json:"logs"`
	// Number of errors reported by RUM.
	Rum int64 `json:"rum"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactErrorCounts instantiates a new IssueImpactErrorCounts object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactErrorCounts(apm int64, llmobs int64, logs int64, rum int64) *IssueImpactErrorCounts {
	this := IssueImpactErrorCounts{}
	this.Apm = apm
	this.Llmobs = llmobs
	this.Logs = logs
	this.Rum = rum
	return &this
}

// NewIssueImpactErrorCountsWithDefaults instantiates a new IssueImpactErrorCounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactErrorCountsWithDefaults() *IssueImpactErrorCounts {
	this := IssueImpactErrorCounts{}
	return &this
}

// GetApm returns the Apm field value.
func (o *IssueImpactErrorCounts) GetApm() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Apm
}

// GetApmOk returns a tuple with the Apm field value
// and a boolean to check if the value has been set.
func (o *IssueImpactErrorCounts) GetApmOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apm, true
}

// SetApm sets field value.
func (o *IssueImpactErrorCounts) SetApm(v int64) {
	o.Apm = v
}

// GetLlmobs returns the Llmobs field value.
func (o *IssueImpactErrorCounts) GetLlmobs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Llmobs
}

// GetLlmobsOk returns a tuple with the Llmobs field value
// and a boolean to check if the value has been set.
func (o *IssueImpactErrorCounts) GetLlmobsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Llmobs, true
}

// SetLlmobs sets field value.
func (o *IssueImpactErrorCounts) SetLlmobs(v int64) {
	o.Llmobs = v
}

// GetLogs returns the Logs field value.
func (o *IssueImpactErrorCounts) GetLogs() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value
// and a boolean to check if the value has been set.
func (o *IssueImpactErrorCounts) GetLogsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logs, true
}

// SetLogs sets field value.
func (o *IssueImpactErrorCounts) SetLogs(v int64) {
	o.Logs = v
}

// GetRum returns the Rum field value.
func (o *IssueImpactErrorCounts) GetRum() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Rum
}

// GetRumOk returns a tuple with the Rum field value
// and a boolean to check if the value has been set.
func (o *IssueImpactErrorCounts) GetRumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rum, true
}

// SetRum sets field value.
func (o *IssueImpactErrorCounts) SetRum(v int64) {
	o.Rum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactErrorCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["apm"] = o.Apm
	toSerialize["llmobs"] = o.Llmobs
	toSerialize["logs"] = o.Logs
	toSerialize["rum"] = o.Rum

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactErrorCounts) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Apm    *int64 `json:"apm"`
		Llmobs *int64 `json:"llmobs"`
		Logs   *int64 `json:"logs"`
		Rum    *int64 `json:"rum"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Apm == nil {
		return fmt.Errorf("required field apm missing")
	}
	if all.Llmobs == nil {
		return fmt.Errorf("required field llmobs missing")
	}
	if all.Logs == nil {
		return fmt.Errorf("required field logs missing")
	}
	if all.Rum == nil {
		return fmt.Errorf("required field rum missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm", "llmobs", "logs", "rum"})
	} else {
		return err
	}
	o.Apm = *all.Apm
	o.Llmobs = *all.Llmobs
	o.Logs = *all.Logs
	o.Rum = *all.Rum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
