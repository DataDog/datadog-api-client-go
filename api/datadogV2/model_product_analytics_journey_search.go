// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneySearch Defines the steps of the journey and the filters applied to it.
type ProductAnalyticsJourneySearch struct {
	// Expression combining the node aliases in order, for example `A -> B -> C`.
	Expression string `json:"expression"`
	// Filters applied on top of the journey step expression.
	Filters *ProductAnalyticsJourneySearchFilters `json:"filters,omitempty"`
	// Identity join keys used to stitch events belonging to the same user or session.
	JoinKeys *ProductAnalyticsJoinKeys `json:"join_keys,omitempty"`
	// Map of node alias to the query matching that step of the journey.
	// Every alias used in `expression` must have an entry here.
	NodeObjects map[string]ProductAnalyticsBaseQuery `json:"node_objects"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneySearch instantiates a new ProductAnalyticsJourneySearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneySearch(expression string, nodeObjects map[string]ProductAnalyticsBaseQuery) *ProductAnalyticsJourneySearch {
	this := ProductAnalyticsJourneySearch{}
	this.Expression = expression
	this.NodeObjects = nodeObjects
	return &this
}

// NewProductAnalyticsJourneySearchWithDefaults instantiates a new ProductAnalyticsJourneySearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneySearchWithDefaults() *ProductAnalyticsJourneySearch {
	this := ProductAnalyticsJourneySearch{}
	return &this
}

// GetExpression returns the Expression field value.
func (o *ProductAnalyticsJourneySearch) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearch) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *ProductAnalyticsJourneySearch) SetExpression(v string) {
	o.Expression = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneySearch) GetFilters() ProductAnalyticsJourneySearchFilters {
	if o == nil || o.Filters == nil {
		var ret ProductAnalyticsJourneySearchFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearch) GetFiltersOk() (*ProductAnalyticsJourneySearchFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneySearch) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given ProductAnalyticsJourneySearchFilters and assigns it to the Filters field.
func (o *ProductAnalyticsJourneySearch) SetFilters(v ProductAnalyticsJourneySearchFilters) {
	o.Filters = &v
}

// GetJoinKeys returns the JoinKeys field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneySearch) GetJoinKeys() ProductAnalyticsJoinKeys {
	if o == nil || o.JoinKeys == nil {
		var ret ProductAnalyticsJoinKeys
		return ret
	}
	return *o.JoinKeys
}

// GetJoinKeysOk returns a tuple with the JoinKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearch) GetJoinKeysOk() (*ProductAnalyticsJoinKeys, bool) {
	if o == nil || o.JoinKeys == nil {
		return nil, false
	}
	return o.JoinKeys, true
}

// HasJoinKeys returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneySearch) HasJoinKeys() bool {
	return o != nil && o.JoinKeys != nil
}

// SetJoinKeys gets a reference to the given ProductAnalyticsJoinKeys and assigns it to the JoinKeys field.
func (o *ProductAnalyticsJourneySearch) SetJoinKeys(v ProductAnalyticsJoinKeys) {
	o.JoinKeys = &v
}

// GetNodeObjects returns the NodeObjects field value.
func (o *ProductAnalyticsJourneySearch) GetNodeObjects() map[string]ProductAnalyticsBaseQuery {
	if o == nil {
		var ret map[string]ProductAnalyticsBaseQuery
		return ret
	}
	return o.NodeObjects
}

// GetNodeObjectsOk returns a tuple with the NodeObjects field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneySearch) GetNodeObjectsOk() (*map[string]ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeObjects, true
}

// SetNodeObjects sets field value.
func (o *ProductAnalyticsJourneySearch) SetNodeObjects(v map[string]ProductAnalyticsBaseQuery) {
	o.NodeObjects = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneySearch) MarshalJSON() ([]byte, error) {
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
	toSerialize["node_objects"] = o.NodeObjects

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneySearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression  *string                               `json:"expression"`
		Filters     *ProductAnalyticsJourneySearchFilters `json:"filters,omitempty"`
		JoinKeys    *ProductAnalyticsJoinKeys             `json:"join_keys,omitempty"`
		NodeObjects *map[string]ProductAnalyticsBaseQuery `json:"node_objects"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.NodeObjects == nil {
		return fmt.Errorf("required field node_objects missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expression", "filters", "join_keys", "node_objects"})
	} else {
		return err
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
	o.NodeObjects = *all.NodeObjects

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
