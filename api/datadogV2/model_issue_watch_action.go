// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueWatchAction Action applied to the issue when the watch condition is met.
type IssueWatchAction struct {
	// State of the issue
	From *IssueState `json:"from,omitempty"`
	// Kind of watch action.
	Kind IssueWatchActionKind `json:"kind"`
	// State of the issue
	To *IssueState `json:"to,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueWatchAction instantiates a new IssueWatchAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueWatchAction(kind IssueWatchActionKind) *IssueWatchAction {
	this := IssueWatchAction{}
	this.Kind = kind
	return &this
}

// NewIssueWatchActionWithDefaults instantiates a new IssueWatchAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueWatchActionWithDefaults() *IssueWatchAction {
	this := IssueWatchAction{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *IssueWatchAction) GetFrom() IssueState {
	if o == nil || o.From == nil {
		var ret IssueState
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchAction) GetFromOk() (*IssueState, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *IssueWatchAction) HasFrom() bool {
	return o != nil && o.From != nil
}

// SetFrom gets a reference to the given IssueState and assigns it to the From field.
func (o *IssueWatchAction) SetFrom(v IssueState) {
	o.From = &v
}

// GetKind returns the Kind field value.
func (o *IssueWatchAction) GetKind() IssueWatchActionKind {
	if o == nil {
		var ret IssueWatchActionKind
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *IssueWatchAction) GetKindOk() (*IssueWatchActionKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *IssueWatchAction) SetKind(v IssueWatchActionKind) {
	o.Kind = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *IssueWatchAction) GetTo() IssueState {
	if o == nil || o.To == nil {
		var ret IssueState
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueWatchAction) GetToOk() (*IssueState, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *IssueWatchAction) HasTo() bool {
	return o != nil && o.To != nil
}

// SetTo gets a reference to the given IssueState and assigns it to the To field.
func (o *IssueWatchAction) SetTo(v IssueState) {
	o.To = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueWatchAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	toSerialize["kind"] = o.Kind
	if o.To != nil {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueWatchAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *IssueState           `json:"from,omitempty"`
		Kind *IssueWatchActionKind `json:"kind"`
		To   *IssueState           `json:"to,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "kind", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.From != nil && !all.From.IsValid() {
		hasInvalidField = true
	} else {
		o.From = all.From
	}
	if !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = *all.Kind
	}
	if all.To != nil && !all.To.IsValid() {
		hasInvalidField = true
	} else {
		o.To = all.To
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
