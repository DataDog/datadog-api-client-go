// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupFindingsBySeverity Number of the group's findings for each value of their `@severity` attribute. The response includes every
// level, including levels with a count of zero.
type RemediationGroupFindingsBySeverity struct {
	// Number of critical findings.
	Critical int64 `json:"critical"`
	// Number of high findings.
	High int64 `json:"high"`
	// Number of info findings.
	Info int64 `json:"info"`
	// Number of low findings.
	Low int64 `json:"low"`
	// Number of medium findings.
	Medium int64 `json:"medium"`
	// Number of findings with no severity.
	None int64 `json:"none"`
	// Number of findings with an unknown severity.
	Unknown int64 `json:"unknown"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupFindingsBySeverity instantiates a new RemediationGroupFindingsBySeverity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupFindingsBySeverity(critical int64, high int64, info int64, low int64, medium int64, none int64, unknown int64) *RemediationGroupFindingsBySeverity {
	this := RemediationGroupFindingsBySeverity{}
	this.Critical = critical
	this.High = high
	this.Info = info
	this.Low = low
	this.Medium = medium
	this.None = none
	this.Unknown = unknown
	return &this
}

// NewRemediationGroupFindingsBySeverityWithDefaults instantiates a new RemediationGroupFindingsBySeverity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupFindingsBySeverityWithDefaults() *RemediationGroupFindingsBySeverity {
	this := RemediationGroupFindingsBySeverity{}
	return &this
}

// GetCritical returns the Critical field value.
func (o *RemediationGroupFindingsBySeverity) GetCritical() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetCriticalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Critical, true
}

// SetCritical sets field value.
func (o *RemediationGroupFindingsBySeverity) SetCritical(v int64) {
	o.Critical = v
}

// GetHigh returns the High field value.
func (o *RemediationGroupFindingsBySeverity) GetHigh() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.High
}

// GetHighOk returns a tuple with the High field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetHighOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.High, true
}

// SetHigh sets field value.
func (o *RemediationGroupFindingsBySeverity) SetHigh(v int64) {
	o.High = v
}

// GetInfo returns the Info field value.
func (o *RemediationGroupFindingsBySeverity) GetInfo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetInfoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value.
func (o *RemediationGroupFindingsBySeverity) SetInfo(v int64) {
	o.Info = v
}

// GetLow returns the Low field value.
func (o *RemediationGroupFindingsBySeverity) GetLow() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Low
}

// GetLowOk returns a tuple with the Low field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetLowOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Low, true
}

// SetLow sets field value.
func (o *RemediationGroupFindingsBySeverity) SetLow(v int64) {
	o.Low = v
}

// GetMedium returns the Medium field value.
func (o *RemediationGroupFindingsBySeverity) GetMedium() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetMediumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Medium, true
}

// SetMedium sets field value.
func (o *RemediationGroupFindingsBySeverity) SetMedium(v int64) {
	o.Medium = v
}

// GetNone returns the None field value.
func (o *RemediationGroupFindingsBySeverity) GetNone() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.None
}

// GetNoneOk returns a tuple with the None field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetNoneOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.None, true
}

// SetNone sets field value.
func (o *RemediationGroupFindingsBySeverity) SetNone(v int64) {
	o.None = v
}

// GetUnknown returns the Unknown field value.
func (o *RemediationGroupFindingsBySeverity) GetUnknown() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupFindingsBySeverity) GetUnknownOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unknown, true
}

// SetUnknown sets field value.
func (o *RemediationGroupFindingsBySeverity) SetUnknown(v int64) {
	o.Unknown = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupFindingsBySeverity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["critical"] = o.Critical
	toSerialize["high"] = o.High
	toSerialize["info"] = o.Info
	toSerialize["low"] = o.Low
	toSerialize["medium"] = o.Medium
	toSerialize["none"] = o.None
	toSerialize["unknown"] = o.Unknown

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupFindingsBySeverity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Critical *int64 `json:"critical"`
		High     *int64 `json:"high"`
		Info     *int64 `json:"info"`
		Low      *int64 `json:"low"`
		Medium   *int64 `json:"medium"`
		None     *int64 `json:"none"`
		Unknown  *int64 `json:"unknown"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Critical == nil {
		return fmt.Errorf("required field critical missing")
	}
	if all.High == nil {
		return fmt.Errorf("required field high missing")
	}
	if all.Info == nil {
		return fmt.Errorf("required field info missing")
	}
	if all.Low == nil {
		return fmt.Errorf("required field low missing")
	}
	if all.Medium == nil {
		return fmt.Errorf("required field medium missing")
	}
	if all.None == nil {
		return fmt.Errorf("required field none missing")
	}
	if all.Unknown == nil {
		return fmt.Errorf("required field unknown missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"critical", "high", "info", "low", "medium", "none", "unknown"})
	} else {
		return err
	}
	o.Critical = *all.Critical
	o.High = *all.High
	o.Info = *all.Info
	o.Low = *all.Low
	o.Medium = *all.Medium
	o.None = *all.None
	o.Unknown = *all.Unknown

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
