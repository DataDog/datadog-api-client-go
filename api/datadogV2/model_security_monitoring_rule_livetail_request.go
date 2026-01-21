// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleLivetailRequest Request to preview a rule query with applied filters.
type SecurityMonitoringRuleLivetailRequest struct {
	// Data source for the query.
	DataSource string `json:"dataSource"`
	// The detection method.
	DetectionMethod SecurityMonitoringRuleDetectionMethod `json:"detectionMethod"`
	// Fields to apply distinct on.
	DistinctFields []string `json:"distinctFields,omitempty"`
	// Additional security filters to apply.
	Filters []SecurityMonitoringFilter `json:"filters,omitempty"`
	// Fields to group by.
	GroupByFields []string `json:"groupByFields,omitempty"`
	// The query to preview.
	Query string `json:"query"`
	// Index of the query in the rule.
	QueryIndex int32 `json:"queryIndex"`
	// The rule type.
	Type SecurityMonitoringRuleTypeRead `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleLivetailRequest instantiates a new SecurityMonitoringRuleLivetailRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleLivetailRequest(dataSource string, detectionMethod SecurityMonitoringRuleDetectionMethod, query string, queryIndex int32, typeVar SecurityMonitoringRuleTypeRead) *SecurityMonitoringRuleLivetailRequest {
	this := SecurityMonitoringRuleLivetailRequest{}
	this.DataSource = dataSource
	this.DetectionMethod = detectionMethod
	this.Query = query
	this.QueryIndex = queryIndex
	this.Type = typeVar
	return &this
}

// NewSecurityMonitoringRuleLivetailRequestWithDefaults instantiates a new SecurityMonitoringRuleLivetailRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleLivetailRequestWithDefaults() *SecurityMonitoringRuleLivetailRequest {
	this := SecurityMonitoringRuleLivetailRequest{}
	return &this
}

// GetDataSource returns the DataSource field value.
func (o *SecurityMonitoringRuleLivetailRequest) GetDataSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *SecurityMonitoringRuleLivetailRequest) SetDataSource(v string) {
	o.DataSource = v
}

// GetDetectionMethod returns the DetectionMethod field value.
func (o *SecurityMonitoringRuleLivetailRequest) GetDetectionMethod() SecurityMonitoringRuleDetectionMethod {
	if o == nil {
		var ret SecurityMonitoringRuleDetectionMethod
		return ret
	}
	return o.DetectionMethod
}

// GetDetectionMethodOk returns a tuple with the DetectionMethod field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetDetectionMethodOk() (*SecurityMonitoringRuleDetectionMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionMethod, true
}

// SetDetectionMethod sets field value.
func (o *SecurityMonitoringRuleLivetailRequest) SetDetectionMethod(v SecurityMonitoringRuleDetectionMethod) {
	o.DetectionMethod = v
}

// GetDistinctFields returns the DistinctFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleLivetailRequest) GetDistinctFields() []string {
	if o == nil || o.DistinctFields == nil {
		var ret []string
		return ret
	}
	return o.DistinctFields
}

// GetDistinctFieldsOk returns a tuple with the DistinctFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetDistinctFieldsOk() (*[]string, bool) {
	if o == nil || o.DistinctFields == nil {
		return nil, false
	}
	return &o.DistinctFields, true
}

// HasDistinctFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleLivetailRequest) HasDistinctFields() bool {
	return o != nil && o.DistinctFields != nil
}

// SetDistinctFields gets a reference to the given []string and assigns it to the DistinctFields field.
func (o *SecurityMonitoringRuleLivetailRequest) SetDistinctFields(v []string) {
	o.DistinctFields = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleLivetailRequest) GetFilters() []SecurityMonitoringFilter {
	if o == nil || o.Filters == nil {
		var ret []SecurityMonitoringFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetFiltersOk() (*[]SecurityMonitoringFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleLivetailRequest) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []SecurityMonitoringFilter and assigns it to the Filters field.
func (o *SecurityMonitoringRuleLivetailRequest) SetFilters(v []SecurityMonitoringFilter) {
	o.Filters = v
}

// GetGroupByFields returns the GroupByFields field value if set, zero value otherwise.
func (o *SecurityMonitoringRuleLivetailRequest) GetGroupByFields() []string {
	if o == nil || o.GroupByFields == nil {
		var ret []string
		return ret
	}
	return o.GroupByFields
}

// GetGroupByFieldsOk returns a tuple with the GroupByFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetGroupByFieldsOk() (*[]string, bool) {
	if o == nil || o.GroupByFields == nil {
		return nil, false
	}
	return &o.GroupByFields, true
}

// HasGroupByFields returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleLivetailRequest) HasGroupByFields() bool {
	return o != nil && o.GroupByFields != nil
}

// SetGroupByFields gets a reference to the given []string and assigns it to the GroupByFields field.
func (o *SecurityMonitoringRuleLivetailRequest) SetGroupByFields(v []string) {
	o.GroupByFields = v
}

// GetQuery returns the Query field value.
func (o *SecurityMonitoringRuleLivetailRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SecurityMonitoringRuleLivetailRequest) SetQuery(v string) {
	o.Query = v
}

// GetQueryIndex returns the QueryIndex field value.
func (o *SecurityMonitoringRuleLivetailRequest) GetQueryIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.QueryIndex
}

// GetQueryIndexOk returns a tuple with the QueryIndex field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetQueryIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryIndex, true
}

// SetQueryIndex sets field value.
func (o *SecurityMonitoringRuleLivetailRequest) SetQueryIndex(v int32) {
	o.QueryIndex = v
}

// GetType returns the Type field value.
func (o *SecurityMonitoringRuleLivetailRequest) GetType() SecurityMonitoringRuleTypeRead {
	if o == nil {
		var ret SecurityMonitoringRuleTypeRead
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringRuleLivetailRequest) GetTypeOk() (*SecurityMonitoringRuleTypeRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SecurityMonitoringRuleLivetailRequest) SetType(v SecurityMonitoringRuleTypeRead) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleLivetailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataSource"] = o.DataSource
	toSerialize["detectionMethod"] = o.DetectionMethod
	if o.DistinctFields != nil {
		toSerialize["distinctFields"] = o.DistinctFields
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.GroupByFields != nil {
		toSerialize["groupByFields"] = o.GroupByFields
	}
	toSerialize["query"] = o.Query
	toSerialize["queryIndex"] = o.QueryIndex
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleLivetailRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataSource      *string                                `json:"dataSource"`
		DetectionMethod *SecurityMonitoringRuleDetectionMethod `json:"detectionMethod"`
		DistinctFields  []string                               `json:"distinctFields,omitempty"`
		Filters         []SecurityMonitoringFilter             `json:"filters,omitempty"`
		GroupByFields   []string                               `json:"groupByFields,omitempty"`
		Query           *string                                `json:"query"`
		QueryIndex      *int32                                 `json:"queryIndex"`
		Type            *SecurityMonitoringRuleTypeRead        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field dataSource missing")
	}
	if all.DetectionMethod == nil {
		return fmt.Errorf("required field detectionMethod missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.QueryIndex == nil {
		return fmt.Errorf("required field queryIndex missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataSource", "detectionMethod", "distinctFields", "filters", "groupByFields", "query", "queryIndex", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DataSource = *all.DataSource
	if !all.DetectionMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.DetectionMethod = *all.DetectionMethod
	}
	o.DistinctFields = all.DistinctFields
	o.Filters = all.Filters
	o.GroupByFields = all.GroupByFields
	o.Query = *all.Query
	o.QueryIndex = *all.QueryIndex
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
