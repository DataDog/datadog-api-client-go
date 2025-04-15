// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Budget A Budget.
type Budget struct {
	// The timestamp when the budget was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The id of the user that created the budget.
	CreatedBy *string `json:"created_by,omitempty"`
	// The month when the budget ends.
	EndMonth *int64 `json:"end_month,omitempty"`
	// The id of the budget.
	Id *string `json:"id,omitempty"`
	// The cost query for to track against the budget.
	MetricsQuery *string `json:"metrics_query,omitempty"`
	// The name of the budget.
	Name *string `json:"name,omitempty"`
	// The Id of the Org the budget belongs to.
	OrgId *int64 `json:"org_id,omitempty"`
	// The month when the budget starts.
	StartMonth *int64 `json:"start_month,omitempty"`
	// The tags for which the budget is created.
	Tags []string `json:"tags,omitempty"`
	// The sum of all budget entries amounts.
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The timestamp when the budget was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The id of the user that created the budget.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudget instantiates a new Budget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudget() *Budget {
	this := Budget{}
	return &this
}

// NewBudgetWithDefaults instantiates a new Budget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetWithDefaults() *Budget {
	this := Budget{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Budget) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Budget) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Budget) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Budget) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Budget) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Budget) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEndMonth returns the EndMonth field value if set, zero value otherwise.
func (o *Budget) GetEndMonth() int64 {
	if o == nil || o.EndMonth == nil {
		var ret int64
		return ret
	}
	return *o.EndMonth
}

// GetEndMonthOk returns a tuple with the EndMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetEndMonthOk() (*int64, bool) {
	if o == nil || o.EndMonth == nil {
		return nil, false
	}
	return o.EndMonth, true
}

// HasEndMonth returns a boolean if a field has been set.
func (o *Budget) HasEndMonth() bool {
	return o != nil && o.EndMonth != nil
}

// SetEndMonth gets a reference to the given int64 and assigns it to the EndMonth field.
func (o *Budget) SetEndMonth(v int64) {
	o.EndMonth = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Budget) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Budget) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Budget) SetId(v string) {
	o.Id = &v
}

// GetMetricsQuery returns the MetricsQuery field value if set, zero value otherwise.
func (o *Budget) GetMetricsQuery() string {
	if o == nil || o.MetricsQuery == nil {
		var ret string
		return ret
	}
	return *o.MetricsQuery
}

// GetMetricsQueryOk returns a tuple with the MetricsQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetMetricsQueryOk() (*string, bool) {
	if o == nil || o.MetricsQuery == nil {
		return nil, false
	}
	return o.MetricsQuery, true
}

// HasMetricsQuery returns a boolean if a field has been set.
func (o *Budget) HasMetricsQuery() bool {
	return o != nil && o.MetricsQuery != nil
}

// SetMetricsQuery gets a reference to the given string and assigns it to the MetricsQuery field.
func (o *Budget) SetMetricsQuery(v string) {
	o.MetricsQuery = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Budget) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Budget) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Budget) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Budget) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Budget) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *Budget) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetStartMonth returns the StartMonth field value if set, zero value otherwise.
func (o *Budget) GetStartMonth() int64 {
	if o == nil || o.StartMonth == nil {
		var ret int64
		return ret
	}
	return *o.StartMonth
}

// GetStartMonthOk returns a tuple with the StartMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetStartMonthOk() (*int64, bool) {
	if o == nil || o.StartMonth == nil {
		return nil, false
	}
	return o.StartMonth, true
}

// HasStartMonth returns a boolean if a field has been set.
func (o *Budget) HasStartMonth() bool {
	return o != nil && o.StartMonth != nil
}

// SetStartMonth gets a reference to the given int64 and assigns it to the StartMonth field.
func (o *Budget) SetStartMonth(v int64) {
	o.StartMonth = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Budget) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Budget) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Budget) SetTags(v []string) {
	o.Tags = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *Budget) GetTotalAmount() float64 {
	if o == nil || o.TotalAmount == nil {
		var ret float64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetTotalAmountOk() (*float64, bool) {
	if o == nil || o.TotalAmount == nil {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *Budget) HasTotalAmount() bool {
	return o != nil && o.TotalAmount != nil
}

// SetTotalAmount gets a reference to the given float64 and assigns it to the TotalAmount field.
func (o *Budget) SetTotalAmount(v float64) {
	o.TotalAmount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Budget) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Budget) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Budget) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Budget) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Budget) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Budget) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *Budget) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Budget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.EndMonth != nil {
		toSerialize["end_month"] = o.EndMonth
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetricsQuery != nil {
		toSerialize["metrics_query"] = o.MetricsQuery
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.StartMonth != nil {
		toSerialize["start_month"] = o.StartMonth
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TotalAmount != nil {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Budget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *int64   `json:"created_at,omitempty"`
		CreatedBy    *string  `json:"created_by,omitempty"`
		EndMonth     *int64   `json:"end_month,omitempty"`
		Id           *string  `json:"id,omitempty"`
		MetricsQuery *string  `json:"metrics_query,omitempty"`
		Name         *string  `json:"name,omitempty"`
		OrgId        *int64   `json:"org_id,omitempty"`
		StartMonth   *int64   `json:"start_month,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		TotalAmount  *float64 `json:"total_amount,omitempty"`
		UpdatedAt    *int64   `json:"updated_at,omitempty"`
		UpdatedBy    *string  `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "end_month", "id", "metrics_query", "name", "org_id", "start_month", "tags", "total_amount", "updated_at", "updated_by"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.EndMonth = all.EndMonth
	o.Id = all.Id
	o.MetricsQuery = all.MetricsQuery
	o.Name = all.Name
	o.OrgId = all.OrgId
	o.StartMonth = all.StartMonth
	o.Tags = all.Tags
	o.TotalAmount = all.TotalAmount
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
