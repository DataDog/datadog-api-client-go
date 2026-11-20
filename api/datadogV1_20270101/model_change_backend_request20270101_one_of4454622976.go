// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeBackendRequest20270101OneOf4454622976 Test-only generated schema for the dashboard versioning proof of concept.
type ChangeBackendRequest20270101OneOf4454622976 struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryOptions map[string]DashboardInlineSchema8 `json:"_query_options,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ChangeType *ChangeTypeOptions `json:"change_type,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CompareTo *CompareToOptions `json:"compare_to,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	IncreaseGood *bool `json:"increase_good,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Limit *float64 `json:"limit,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderBy *OrderByOptions `json:"order_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	OrderDir *QuerySortOrder `json:"order_dir,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Q string `json:"q"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	ShowPresent *bool `json:"show_present,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewChangeBackendRequest20270101OneOf4454622976 instantiates a new ChangeBackendRequest20270101OneOf4454622976 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeBackendRequest20270101OneOf4454622976(q string) *ChangeBackendRequest20270101OneOf4454622976 {
	this := ChangeBackendRequest20270101OneOf4454622976{}
	this.Q = q
	return &this
}

// NewChangeBackendRequest20270101OneOf4454622976WithDefaults instantiates a new ChangeBackendRequest20270101OneOf4454622976 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeBackendRequest20270101OneOf4454622976WithDefaults() *ChangeBackendRequest20270101OneOf4454622976 {
	this := ChangeBackendRequest20270101OneOf4454622976{}
	return &this
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetQueryOptions() map[string]DashboardInlineSchema8 {
	if o == nil || o.QueryOptions == nil {
		var ret map[string]DashboardInlineSchema8
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetQueryOptionsOk() (*map[string]DashboardInlineSchema8, bool) {
	if o == nil || o.QueryOptions == nil {
		return nil, false
	}
	return &o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasQueryOptions() bool {
	return o != nil && o.QueryOptions != nil
}

// SetQueryOptions gets a reference to the given map[string]DashboardInlineSchema8 and assigns it to the QueryOptions field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetQueryOptions(v map[string]DashboardInlineSchema8) {
	o.QueryOptions = v
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetChangeType() ChangeTypeOptions {
	if o == nil || o.ChangeType == nil {
		var ret ChangeTypeOptions
		return ret
	}
	return *o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetChangeTypeOk() (*ChangeTypeOptions, bool) {
	if o == nil || o.ChangeType == nil {
		return nil, false
	}
	return o.ChangeType, true
}

// HasChangeType returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasChangeType() bool {
	return o != nil && o.ChangeType != nil
}

// SetChangeType gets a reference to the given ChangeTypeOptions and assigns it to the ChangeType field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetChangeType(v ChangeTypeOptions) {
	o.ChangeType = &v
}

// GetCompareTo returns the CompareTo field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetCompareTo() CompareToOptions {
	if o == nil || o.CompareTo == nil {
		var ret CompareToOptions
		return ret
	}
	return *o.CompareTo
}

// GetCompareToOk returns a tuple with the CompareTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetCompareToOk() (*CompareToOptions, bool) {
	if o == nil || o.CompareTo == nil {
		return nil, false
	}
	return o.CompareTo, true
}

// HasCompareTo returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasCompareTo() bool {
	return o != nil && o.CompareTo != nil
}

// SetCompareTo gets a reference to the given CompareToOptions and assigns it to the CompareTo field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetCompareTo(v CompareToOptions) {
	o.CompareTo = &v
}

// GetIncreaseGood returns the IncreaseGood field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetIncreaseGood() bool {
	if o == nil || o.IncreaseGood == nil {
		var ret bool
		return ret
	}
	return *o.IncreaseGood
}

// GetIncreaseGoodOk returns a tuple with the IncreaseGood field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetIncreaseGoodOk() (*bool, bool) {
	if o == nil || o.IncreaseGood == nil {
		return nil, false
	}
	return o.IncreaseGood, true
}

// HasIncreaseGood returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasIncreaseGood() bool {
	return o != nil && o.IncreaseGood != nil
}

// SetIncreaseGood gets a reference to the given bool and assigns it to the IncreaseGood field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetIncreaseGood(v bool) {
	o.IncreaseGood = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetLimit() float64 {
	if o == nil || o.Limit == nil {
		var ret float64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetLimitOk() (*float64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given float64 and assigns it to the Limit field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetLimit(v float64) {
	o.Limit = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetOrderBy() OrderByOptions {
	if o == nil || o.OrderBy == nil {
		var ret OrderByOptions
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetOrderByOk() (*OrderByOptions, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given OrderByOptions and assigns it to the OrderBy field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetOrderBy(v OrderByOptions) {
	o.OrderBy = &v
}

// GetOrderDir returns the OrderDir field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetOrderDir() QuerySortOrder {
	if o == nil || o.OrderDir == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.OrderDir
}

// GetOrderDirOk returns a tuple with the OrderDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetOrderDirOk() (*QuerySortOrder, bool) {
	if o == nil || o.OrderDir == nil {
		return nil, false
	}
	return o.OrderDir, true
}

// HasOrderDir returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasOrderDir() bool {
	return o != nil && o.OrderDir != nil
}

// SetOrderDir gets a reference to the given QuerySortOrder and assigns it to the OrderDir field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetOrderDir(v QuerySortOrder) {
	o.OrderDir = &v
}

// GetQ returns the Q field value.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetQ() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Q
}

// GetQOk returns a tuple with the Q field value
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetQOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Q, true
}

// SetQ sets field value.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetQ(v string) {
	o.Q = v
}

// GetShowPresent returns the ShowPresent field value if set, zero value otherwise.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetShowPresent() bool {
	if o == nil || o.ShowPresent == nil {
		var ret bool
		return ret
	}
	return *o.ShowPresent
}

// GetShowPresentOk returns a tuple with the ShowPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) GetShowPresentOk() (*bool, bool) {
	if o == nil || o.ShowPresent == nil {
		return nil, false
	}
	return o.ShowPresent, true
}

// HasShowPresent returns a boolean if a field has been set.
func (o *ChangeBackendRequest20270101OneOf4454622976) HasShowPresent() bool {
	return o != nil && o.ShowPresent != nil
}

// SetShowPresent gets a reference to the given bool and assigns it to the ShowPresent field.
func (o *ChangeBackendRequest20270101OneOf4454622976) SetShowPresent(v bool) {
	o.ShowPresent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeBackendRequest20270101OneOf4454622976) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryOptions != nil {
		toSerialize["_query_options"] = o.QueryOptions
	}
	if o.ChangeType != nil {
		toSerialize["change_type"] = o.ChangeType
	}
	if o.CompareTo != nil {
		toSerialize["compare_to"] = o.CompareTo
	}
	if o.IncreaseGood != nil {
		toSerialize["increase_good"] = o.IncreaseGood
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.OrderDir != nil {
		toSerialize["order_dir"] = o.OrderDir
	}
	toSerialize["q"] = o.Q
	if o.ShowPresent != nil {
		toSerialize["show_present"] = o.ShowPresent
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeBackendRequest20270101OneOf4454622976) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryOptions map[string]DashboardInlineSchema8 `json:"_query_options,omitempty"`
		ChangeType   *ChangeTypeOptions                `json:"change_type,omitempty"`
		CompareTo    *CompareToOptions                 `json:"compare_to,omitempty"`
		IncreaseGood *bool                             `json:"increase_good,omitempty"`
		Limit        *float64                          `json:"limit,omitempty"`
		OrderBy      *OrderByOptions                   `json:"order_by,omitempty"`
		OrderDir     *QuerySortOrder                   `json:"order_dir,omitempty"`
		Q            *string                           `json:"q"`
		ShowPresent  *bool                             `json:"show_present,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Q == nil {
		return fmt.Errorf("required field q missing")
	}

	hasInvalidField := false
	o.QueryOptions = all.QueryOptions
	if all.ChangeType != nil && !all.ChangeType.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeType = all.ChangeType
	}
	if all.CompareTo != nil && !all.CompareTo.IsValid() {
		hasInvalidField = true
	} else {
		o.CompareTo = all.CompareTo
	}
	o.IncreaseGood = all.IncreaseGood
	o.Limit = all.Limit
	if all.OrderBy != nil && !all.OrderBy.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderBy = all.OrderBy
	}
	if all.OrderDir != nil && !all.OrderDir.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderDir = all.OrderDir
	}
	o.Q = *all.Q
	o.ShowPresent = all.ShowPresent

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
