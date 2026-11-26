// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LegacyUserJourneySearch Test-only generated schema for the dashboard versioning proof of concept.
type LegacyUserJourneySearch struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	Expression string `json:"expression"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Filters *LegacyUserJourneySearchFilters `json:"filters,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	JoinKeys *LegacyUserJourneySearchJoinKeys `json:"join_keys,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Nodes map[string]string `json:"nodes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewLegacyUserJourneySearch instantiates a new LegacyUserJourneySearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLegacyUserJourneySearch(expression string, nodes map[string]string) *LegacyUserJourneySearch {
	this := LegacyUserJourneySearch{}
	this.Expression = expression
	this.Nodes = nodes
	return &this
}

// NewLegacyUserJourneySearchWithDefaults instantiates a new LegacyUserJourneySearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLegacyUserJourneySearchWithDefaults() *LegacyUserJourneySearch {
	this := LegacyUserJourneySearch{}
	return &this
}

// GetExpression returns the Expression field value.
func (o *LegacyUserJourneySearch) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *LegacyUserJourneySearch) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *LegacyUserJourneySearch) SetExpression(v string) {
	o.Expression = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *LegacyUserJourneySearch) GetFilters() LegacyUserJourneySearchFilters {
	if o == nil || o.Filters == nil {
		var ret LegacyUserJourneySearchFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyUserJourneySearch) GetFiltersOk() (*LegacyUserJourneySearchFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *LegacyUserJourneySearch) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given LegacyUserJourneySearchFilters and assigns it to the Filters field.
func (o *LegacyUserJourneySearch) SetFilters(v LegacyUserJourneySearchFilters) {
	o.Filters = &v
}

// GetJoinKeys returns the JoinKeys field value if set, zero value otherwise.
func (o *LegacyUserJourneySearch) GetJoinKeys() LegacyUserJourneySearchJoinKeys {
	if o == nil || o.JoinKeys == nil {
		var ret LegacyUserJourneySearchJoinKeys
		return ret
	}
	return *o.JoinKeys
}

// GetJoinKeysOk returns a tuple with the JoinKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyUserJourneySearch) GetJoinKeysOk() (*LegacyUserJourneySearchJoinKeys, bool) {
	if o == nil || o.JoinKeys == nil {
		return nil, false
	}
	return o.JoinKeys, true
}

// HasJoinKeys returns a boolean if a field has been set.
func (o *LegacyUserJourneySearch) HasJoinKeys() bool {
	return o != nil && o.JoinKeys != nil
}

// SetJoinKeys gets a reference to the given LegacyUserJourneySearchJoinKeys and assigns it to the JoinKeys field.
func (o *LegacyUserJourneySearch) SetJoinKeys(v LegacyUserJourneySearchJoinKeys) {
	o.JoinKeys = &v
}

// GetNodes returns the Nodes field value.
func (o *LegacyUserJourneySearch) GetNodes() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *LegacyUserJourneySearch) GetNodesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *LegacyUserJourneySearch) SetNodes(v map[string]string) {
	o.Nodes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LegacyUserJourneySearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["expression"] = o.Expression
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.JoinKeys != nil {
		toSerialize["join_keys"] = o.JoinKeys
	}
	toSerialize["nodes"] = o.Nodes
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LegacyUserJourneySearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression *string                          `json:"expression"`
		Filters    *LegacyUserJourneySearchFilters  `json:"filters,omitempty"`
		JoinKeys   *LegacyUserJourneySearchJoinKeys `json:"join_keys,omitempty"`
		Nodes      *map[string]string               `json:"nodes"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}

	hasInvalidField := false
	o.Expression = *all.Expression
	if all.Filters != nil && all.Filters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filters = all.Filters
	if all.JoinKeys != nil && all.JoinKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinKeys = all.JoinKeys
	o.Nodes = *all.Nodes

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
