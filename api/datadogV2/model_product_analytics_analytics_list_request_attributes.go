// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAnalyticsListRequestAttributes Attributes for an analytics list request.
type ProductAnalyticsAnalyticsListRequestAttributes struct {
	// Start time in epoch milliseconds. Must be less than `to`.
	From int64 `json:"from"`
	// The analytics list query definition. It selects the events to return with `query`, then
	// chooses the columns on each event row, the sort applied to those rows, and a row limit.
	// Unlike the scalar and timeseries queries, a list query returns raw event rows rather than
	// aggregates, so it takes no compute or group-by rule.
	Query ProductAnalyticsAnalyticsListQuery `json:"query"`
	// End time in epoch milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAnalyticsListRequestAttributes instantiates a new ProductAnalyticsAnalyticsListRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAnalyticsListRequestAttributes(from int64, query ProductAnalyticsAnalyticsListQuery, to int64) *ProductAnalyticsAnalyticsListRequestAttributes {
	this := ProductAnalyticsAnalyticsListRequestAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewProductAnalyticsAnalyticsListRequestAttributesWithDefaults instantiates a new ProductAnalyticsAnalyticsListRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAnalyticsListRequestAttributesWithDefaults() *ProductAnalyticsAnalyticsListRequestAttributes {
	this := ProductAnalyticsAnalyticsListRequestAttributes{}
	return &this
}

// GetFrom returns the From field value.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) GetQuery() ProductAnalyticsAnalyticsListQuery {
	if o == nil {
		var ret ProductAnalyticsAnalyticsListQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) GetQueryOk() (*ProductAnalyticsAnalyticsListQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) SetQuery(v ProductAnalyticsAnalyticsListQuery) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAnalyticsListRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAnalyticsListRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From  *int64                              `json:"from"`
		Query *ProductAnalyticsAnalyticsListQuery `json:"query"`
		To    *int64                              `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "query", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	o.From = *all.From
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
