// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

// ApmRecommendationsStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type ApmRecommendationsStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CustomServiceTags []string `json:"custom_service_tags,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource ApmRecommendationsStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Env *string `json:"env,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	GroupBy *datadogV1.ApmRecommendationsStreamQuery20270101GroupBy `json:"group_by,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Priorities []string `json:"priorities,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	RecommendationTypes []string `json:"recommendation_types,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Services []string `json:"services,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *ApmRecommendationsStreamQuery20270101Sort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Statuses []string `json:"statuses,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Teams []string `json:"teams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewApmRecommendationsStreamQuery instantiates a new ApmRecommendationsStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApmRecommendationsStreamQuery(dataSource ApmRecommendationsStreamQueryDataSource, queryString string) *ApmRecommendationsStreamQuery {
	this := ApmRecommendationsStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewApmRecommendationsStreamQueryWithDefaults instantiates a new ApmRecommendationsStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApmRecommendationsStreamQueryWithDefaults() *ApmRecommendationsStreamQuery {
	this := ApmRecommendationsStreamQuery{}
	return &this
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *ApmRecommendationsStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetCustomServiceTags returns the CustomServiceTags field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetCustomServiceTags() []string {
	if o == nil || o.CustomServiceTags == nil {
		var ret []string
		return ret
	}
	return o.CustomServiceTags
}

// GetCustomServiceTagsOk returns a tuple with the CustomServiceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetCustomServiceTagsOk() (*[]string, bool) {
	if o == nil || o.CustomServiceTags == nil {
		return nil, false
	}
	return &o.CustomServiceTags, true
}

// HasCustomServiceTags returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasCustomServiceTags() bool {
	return o != nil && o.CustomServiceTags != nil
}

// SetCustomServiceTags gets a reference to the given []string and assigns it to the CustomServiceTags field.
func (o *ApmRecommendationsStreamQuery) SetCustomServiceTags(v []string) {
	o.CustomServiceTags = v
}

// GetDataSource returns the DataSource field value.
func (o *ApmRecommendationsStreamQuery) GetDataSource() ApmRecommendationsStreamQueryDataSource {
	if o == nil {
		var ret ApmRecommendationsStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetDataSourceOk() (*ApmRecommendationsStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ApmRecommendationsStreamQuery) SetDataSource(v ApmRecommendationsStreamQueryDataSource) {
	o.DataSource = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *ApmRecommendationsStreamQuery) SetEnv(v string) {
	o.Env = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetGroupBy() datadogV1.ApmRecommendationsStreamQuery20270101GroupBy {
	if o == nil || o.GroupBy == nil {
		var ret datadogV1.ApmRecommendationsStreamQuery20270101GroupBy
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetGroupByOk() (*datadogV1.ApmRecommendationsStreamQuery20270101GroupBy, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given datadogV1.ApmRecommendationsStreamQuery20270101GroupBy and assigns it to the GroupBy field.
func (o *ApmRecommendationsStreamQuery) SetGroupBy(v datadogV1.ApmRecommendationsStreamQuery20270101GroupBy) {
	o.GroupBy = &v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *ApmRecommendationsStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetPriorities returns the Priorities field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetPriorities() []string {
	if o == nil || o.Priorities == nil {
		var ret []string
		return ret
	}
	return o.Priorities
}

// GetPrioritiesOk returns a tuple with the Priorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetPrioritiesOk() (*[]string, bool) {
	if o == nil || o.Priorities == nil {
		return nil, false
	}
	return &o.Priorities, true
}

// HasPriorities returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasPriorities() bool {
	return o != nil && o.Priorities != nil
}

// SetPriorities gets a reference to the given []string and assigns it to the Priorities field.
func (o *ApmRecommendationsStreamQuery) SetPriorities(v []string) {
	o.Priorities = v
}

// GetQueryString returns the QueryString field value.
func (o *ApmRecommendationsStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *ApmRecommendationsStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetRecommendationTypes returns the RecommendationTypes field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetRecommendationTypes() []string {
	if o == nil || o.RecommendationTypes == nil {
		var ret []string
		return ret
	}
	return o.RecommendationTypes
}

// GetRecommendationTypesOk returns a tuple with the RecommendationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetRecommendationTypesOk() (*[]string, bool) {
	if o == nil || o.RecommendationTypes == nil {
		return nil, false
	}
	return &o.RecommendationTypes, true
}

// HasRecommendationTypes returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasRecommendationTypes() bool {
	return o != nil && o.RecommendationTypes != nil
}

// SetRecommendationTypes gets a reference to the given []string and assigns it to the RecommendationTypes field.
func (o *ApmRecommendationsStreamQuery) SetRecommendationTypes(v []string) {
	o.RecommendationTypes = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *ApmRecommendationsStreamQuery) SetServices(v []string) {
	o.Services = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetSort() ApmRecommendationsStreamQuery20270101Sort {
	if o == nil || o.Sort == nil {
		var ret ApmRecommendationsStreamQuery20270101Sort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetSortOk() (*ApmRecommendationsStreamQuery20270101Sort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given ApmRecommendationsStreamQuery20270101Sort and assigns it to the Sort field.
func (o *ApmRecommendationsStreamQuery) SetSort(v ApmRecommendationsStreamQuery20270101Sort) {
	o.Sort = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetStatuses() []string {
	if o == nil || o.Statuses == nil {
		var ret []string
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetStatusesOk() (*[]string, bool) {
	if o == nil || o.Statuses == nil {
		return nil, false
	}
	return &o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasStatuses() bool {
	return o != nil && o.Statuses != nil
}

// SetStatuses gets a reference to the given []string and assigns it to the Statuses field.
func (o *ApmRecommendationsStreamQuery) SetStatuses(v []string) {
	o.Statuses = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ApmRecommendationsStreamQuery) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApmRecommendationsStreamQuery) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ApmRecommendationsStreamQuery) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ApmRecommendationsStreamQuery) SetTeams(v []string) {
	o.Teams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApmRecommendationsStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	if o.CustomServiceTags != nil {
		toSerialize["custom_service_tags"] = o.CustomServiceTags
	}
	toSerialize["data_source"] = o.DataSource
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.Priorities != nil {
		toSerialize["priorities"] = o.Priorities
	}
	toSerialize["query_string"] = o.QueryString
	if o.RecommendationTypes != nil {
		toSerialize["recommendation_types"] = o.RecommendationTypes
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApmRecommendationsStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossOrgUuids       []string                                                `json:"cross_org_uuids,omitempty"`
		CustomServiceTags   []string                                                `json:"custom_service_tags,omitempty"`
		DataSource          *ApmRecommendationsStreamQueryDataSource                `json:"data_source"`
		Env                 *string                                                 `json:"env,omitempty"`
		GroupBy             *datadogV1.ApmRecommendationsStreamQuery20270101GroupBy `json:"group_by,omitempty"`
		Indexes             []string                                                `json:"indexes,omitempty"`
		Priorities          []string                                                `json:"priorities,omitempty"`
		QueryString         *string                                                 `json:"query_string"`
		RecommendationTypes []string                                                `json:"recommendation_types,omitempty"`
		Services            []string                                                `json:"services,omitempty"`
		Sort                *ApmRecommendationsStreamQuery20270101Sort              `json:"sort,omitempty"`
		Statuses            []string                                                `json:"statuses,omitempty"`
		Teams               []string                                                `json:"teams,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}

	hasInvalidField := false
	o.CrossOrgUuids = all.CrossOrgUuids
	o.CustomServiceTags = all.CustomServiceTags
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Env = all.Env
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.Priorities = all.Priorities
	o.QueryString = *all.QueryString
	o.RecommendationTypes = all.RecommendationTypes
	o.Services = all.Services
	o.Sort = all.Sort
	o.Statuses = all.Statuses
	o.Teams = all.Teams

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
