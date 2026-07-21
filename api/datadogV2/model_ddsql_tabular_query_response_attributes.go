// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DdsqlTabularQueryResponseAttributes Attributes of a DDSQL tabular query response. `query_id` is set when
// `state` is `running`; `columns` is set when `state` is `completed`.
type DdsqlTabularQueryResponseAttributes struct {
	// Column-major result set. Each element carries one column's name, type, and values,
	// with one value per row of the result. Set when `state` is `completed`.
	Columns []DdsqlTabularQueryColumn `json:"columns,omitempty"`
	// Opaque token to pass to the fetch endpoint to poll for results.
	// Set when `state` is `running` and absent when `state` is `completed`.
	QueryId *string `json:"query_id,omitempty"`
	// Lifecycle state of a DDSQL tabular query response.
	// `running` means the query is still executing and the client should poll
	// the fetch endpoint with the returned `query_id`. `completed` means the
	// result set is inlined in `columns` and no further polling is required.
	State DdsqlTabularQueryState `json:"state"`
	// Non-fatal messages emitted by the query engine while serving this response.
	Warnings []string `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDdsqlTabularQueryResponseAttributes instantiates a new DdsqlTabularQueryResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDdsqlTabularQueryResponseAttributes(state DdsqlTabularQueryState) *DdsqlTabularQueryResponseAttributes {
	this := DdsqlTabularQueryResponseAttributes{}
	this.State = state
	return &this
}

// NewDdsqlTabularQueryResponseAttributesWithDefaults instantiates a new DdsqlTabularQueryResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDdsqlTabularQueryResponseAttributesWithDefaults() *DdsqlTabularQueryResponseAttributes {
	this := DdsqlTabularQueryResponseAttributes{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *DdsqlTabularQueryResponseAttributes) GetColumns() []DdsqlTabularQueryColumn {
	if o == nil || o.Columns == nil {
		var ret []DdsqlTabularQueryColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryResponseAttributes) GetColumnsOk() (*[]DdsqlTabularQueryColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *DdsqlTabularQueryResponseAttributes) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []DdsqlTabularQueryColumn and assigns it to the Columns field.
func (o *DdsqlTabularQueryResponseAttributes) SetColumns(v []DdsqlTabularQueryColumn) {
	o.Columns = v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *DdsqlTabularQueryResponseAttributes) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryResponseAttributes) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *DdsqlTabularQueryResponseAttributes) HasQueryId() bool {
	return o != nil && o.QueryId != nil
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *DdsqlTabularQueryResponseAttributes) SetQueryId(v string) {
	o.QueryId = &v
}

// GetState returns the State field value.
func (o *DdsqlTabularQueryResponseAttributes) GetState() DdsqlTabularQueryState {
	if o == nil {
		var ret DdsqlTabularQueryState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryResponseAttributes) GetStateOk() (*DdsqlTabularQueryState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DdsqlTabularQueryResponseAttributes) SetState(v DdsqlTabularQueryState) {
	o.State = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DdsqlTabularQueryResponseAttributes) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdsqlTabularQueryResponseAttributes) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DdsqlTabularQueryResponseAttributes) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *DdsqlTabularQueryResponseAttributes) SetWarnings(v []string) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DdsqlTabularQueryResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.QueryId != nil {
		toSerialize["query_id"] = o.QueryId
	}
	toSerialize["state"] = o.State
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DdsqlTabularQueryResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns  []DdsqlTabularQueryColumn `json:"columns,omitempty"`
		QueryId  *string                   `json:"query_id,omitempty"`
		State    *DdsqlTabularQueryState   `json:"state"`
		Warnings []string                  `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "query_id", "state", "warnings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.QueryId = all.QueryId
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
