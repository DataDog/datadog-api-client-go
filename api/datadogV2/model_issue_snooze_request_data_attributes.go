// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSnoozeRequestDataAttributes Object describing an issue snooze request.
type IssueSnoozeRequestDataAttributes struct {
	// Condition that triggers the watch action. The fields that are set depend on `kind`:
	// `OCCURRENCES` uses `threshold`, `CARDINALITY` uses `threshold` and `projection`,
	// `ELAPSED_TIME` uses `duration`, and `ANY` and `ALL` use `operands`.
	Condition IssueWatchCondition `json:"condition"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSnoozeRequestDataAttributes instantiates a new IssueSnoozeRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSnoozeRequestDataAttributes(condition IssueWatchCondition) *IssueSnoozeRequestDataAttributes {
	this := IssueSnoozeRequestDataAttributes{}
	this.Condition = condition
	return &this
}

// NewIssueSnoozeRequestDataAttributesWithDefaults instantiates a new IssueSnoozeRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSnoozeRequestDataAttributesWithDefaults() *IssueSnoozeRequestDataAttributes {
	this := IssueSnoozeRequestDataAttributes{}
	return &this
}

// GetCondition returns the Condition field value.
func (o *IssueSnoozeRequestDataAttributes) GetCondition() IssueWatchCondition {
	if o == nil {
		var ret IssueWatchCondition
		return ret
	}
	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *IssueSnoozeRequestDataAttributes) GetConditionOk() (*IssueWatchCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value.
func (o *IssueSnoozeRequestDataAttributes) SetCondition(v IssueWatchCondition) {
	o.Condition = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSnoozeRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["condition"] = o.Condition

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSnoozeRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition *IssueWatchCondition `json:"condition"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Condition == nil {
		return fmt.Errorf("required field condition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Condition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Condition = *all.Condition

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
