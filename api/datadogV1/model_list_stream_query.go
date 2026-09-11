// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamQuery Updated list stream widget.
type ListStreamQuery struct {
	// Filter by assignee UUIDs. Usable only with `issue_stream`.
	AssigneeUuids []string `json:"assignee_uuids,omitempty"`
	// Specifies the field for logs pattern clustering. Usable only with logs_pattern_stream.
	ClusteringPatternFieldPath *string `json:"clustering_pattern_field_path,omitempty"`
	// Compute configuration for the List Stream Widget. Compute can be used only with the logs_transaction_stream (from 1 to 5 items) list stream source.
	Compute []ListStreamComputeItems `json:"compute,omitempty"`
	// Source from which to query items to display in the stream. apm_issue_stream, rum_issue_stream, and logs_issue_stream are deprecated. Use issue_stream instead. apm_recommendations_stream is used to query APM recommendations, and supports filtering by environment, services, teams, recommendation types, and status.
	DataSource ListStreamSource `json:"data_source"`
	// Filter by APM environment. Usable only with `apm_recommendations_stream`.
	Env *string `json:"env,omitempty"`
	// Size to use to display an event.
	EventSize *WidgetEventSize `json:"event_size,omitempty"`
	// Group by configuration for the List Stream Widget. Group by can be used only with logs_pattern_stream (up to 4 items) or logs_transaction_stream (one group by item is required) list stream source.
	GroupBy []ListStreamGroupByItems `json:"group_by,omitempty"`
	// List of indexes.
	Indexes []string `json:"indexes,omitempty"`
	// Persona filter for the `issue_stream` data source.
	Persona *ListStreamIssuePersona `json:"persona,omitempty"`
	// Widget query.
	QueryString string `json:"query_string"`
	// Filter by recommendation types. Usable only with `apm_recommendations_stream`.
	RecommendationTypes []string `json:"recommendation_types,omitempty"`
	// Filter by service names. Usable only with `apm_recommendations_stream`.
	Services []string `json:"services,omitempty"`
	// Which column and order to sort by
	Sort *WidgetFieldSort `json:"sort,omitempty"`
	// Filter by issue states. Usable only with `issue_stream`.
	States []ListStreamIssueState `json:"states,omitempty"`
	// Filter by recommendation statuses. Usable only with `apm_recommendations_stream`.
	Statuses []string `json:"statuses,omitempty"`
	// Option for storage location. Feature in Private Beta.
	Storage *string `json:"storage,omitempty"`
	// Filter by suspected causes. Usable only with `issue_stream`.
	SuspectedCauses []string `json:"suspected_causes,omitempty"`
	// Filter by team handles. Usable only with `issue_stream`.
	TeamHandles []string `json:"team_handles,omitempty"`
	// Filter by team handles. Usable only with `apm_recommendations_stream`.
	Teams []string `json:"teams,omitempty"`
	// Version of the query for the logs transaction stream widget. When omitted, v1 query behavior is
	// preserved. Set to `sequential_query` to use v2 behavior. **This feature is in Preview.**
	Version *ListStreamQueryVersion `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListStreamQuery instantiates a new ListStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListStreamQuery(dataSource ListStreamSource, queryString string) *ListStreamQuery {
	this := ListStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewListStreamQueryWithDefaults instantiates a new ListStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListStreamQueryWithDefaults() *ListStreamQuery {
	this := ListStreamQuery{}
	var dataSource ListStreamSource = LISTSTREAMSOURCE_LOGS_STREAM
	this.DataSource = dataSource
	return &this
}

// GetAssigneeUuids returns the AssigneeUuids field value if set, zero value otherwise.
func (o *ListStreamQuery) GetAssigneeUuids() []string {
	if o == nil || o.AssigneeUuids == nil {
		var ret []string
		return ret
	}
	return o.AssigneeUuids
}

// GetAssigneeUuidsOk returns a tuple with the AssigneeUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetAssigneeUuidsOk() (*[]string, bool) {
	if o == nil || o.AssigneeUuids == nil {
		return nil, false
	}
	return &o.AssigneeUuids, true
}

// HasAssigneeUuids returns a boolean if a field has been set.
func (o *ListStreamQuery) HasAssigneeUuids() bool {
	return o != nil && o.AssigneeUuids != nil
}

// SetAssigneeUuids gets a reference to the given []string and assigns it to the AssigneeUuids field.
func (o *ListStreamQuery) SetAssigneeUuids(v []string) {
	o.AssigneeUuids = v
}

// GetClusteringPatternFieldPath returns the ClusteringPatternFieldPath field value if set, zero value otherwise.
func (o *ListStreamQuery) GetClusteringPatternFieldPath() string {
	if o == nil || o.ClusteringPatternFieldPath == nil {
		var ret string
		return ret
	}
	return *o.ClusteringPatternFieldPath
}

// GetClusteringPatternFieldPathOk returns a tuple with the ClusteringPatternFieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetClusteringPatternFieldPathOk() (*string, bool) {
	if o == nil || o.ClusteringPatternFieldPath == nil {
		return nil, false
	}
	return o.ClusteringPatternFieldPath, true
}

// HasClusteringPatternFieldPath returns a boolean if a field has been set.
func (o *ListStreamQuery) HasClusteringPatternFieldPath() bool {
	return o != nil && o.ClusteringPatternFieldPath != nil
}

// SetClusteringPatternFieldPath gets a reference to the given string and assigns it to the ClusteringPatternFieldPath field.
func (o *ListStreamQuery) SetClusteringPatternFieldPath(v string) {
	o.ClusteringPatternFieldPath = &v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *ListStreamQuery) GetCompute() []ListStreamComputeItems {
	if o == nil || o.Compute == nil {
		var ret []ListStreamComputeItems
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetComputeOk() (*[]ListStreamComputeItems, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ListStreamQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []ListStreamComputeItems and assigns it to the Compute field.
func (o *ListStreamQuery) SetCompute(v []ListStreamComputeItems) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *ListStreamQuery) GetDataSource() ListStreamSource {
	if o == nil {
		var ret ListStreamSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetDataSourceOk() (*ListStreamSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ListStreamQuery) SetDataSource(v ListStreamSource) {
	o.DataSource = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ListStreamQuery) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ListStreamQuery) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *ListStreamQuery) SetEnv(v string) {
	o.Env = &v
}

// GetEventSize returns the EventSize field value if set, zero value otherwise.
func (o *ListStreamQuery) GetEventSize() WidgetEventSize {
	if o == nil || o.EventSize == nil {
		var ret WidgetEventSize
		return ret
	}
	return *o.EventSize
}

// GetEventSizeOk returns a tuple with the EventSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetEventSizeOk() (*WidgetEventSize, bool) {
	if o == nil || o.EventSize == nil {
		return nil, false
	}
	return o.EventSize, true
}

// HasEventSize returns a boolean if a field has been set.
func (o *ListStreamQuery) HasEventSize() bool {
	return o != nil && o.EventSize != nil
}

// SetEventSize gets a reference to the given WidgetEventSize and assigns it to the EventSize field.
func (o *ListStreamQuery) SetEventSize(v WidgetEventSize) {
	o.EventSize = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ListStreamQuery) GetGroupBy() []ListStreamGroupByItems {
	if o == nil || o.GroupBy == nil {
		var ret []ListStreamGroupByItems
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetGroupByOk() (*[]ListStreamGroupByItems, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ListStreamQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ListStreamGroupByItems and assigns it to the GroupBy field.
func (o *ListStreamQuery) SetGroupBy(v []ListStreamGroupByItems) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *ListStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *ListStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *ListStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetPersona returns the Persona field value if set, zero value otherwise.
func (o *ListStreamQuery) GetPersona() ListStreamIssuePersona {
	if o == nil || o.Persona == nil {
		var ret ListStreamIssuePersona
		return ret
	}
	return *o.Persona
}

// GetPersonaOk returns a tuple with the Persona field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetPersonaOk() (*ListStreamIssuePersona, bool) {
	if o == nil || o.Persona == nil {
		return nil, false
	}
	return o.Persona, true
}

// HasPersona returns a boolean if a field has been set.
func (o *ListStreamQuery) HasPersona() bool {
	return o != nil && o.Persona != nil
}

// SetPersona gets a reference to the given ListStreamIssuePersona and assigns it to the Persona field.
func (o *ListStreamQuery) SetPersona(v ListStreamIssuePersona) {
	o.Persona = &v
}

// GetQueryString returns the QueryString field value.
func (o *ListStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *ListStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetRecommendationTypes returns the RecommendationTypes field value if set, zero value otherwise.
func (o *ListStreamQuery) GetRecommendationTypes() []string {
	if o == nil || o.RecommendationTypes == nil {
		var ret []string
		return ret
	}
	return o.RecommendationTypes
}

// GetRecommendationTypesOk returns a tuple with the RecommendationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetRecommendationTypesOk() (*[]string, bool) {
	if o == nil || o.RecommendationTypes == nil {
		return nil, false
	}
	return &o.RecommendationTypes, true
}

// HasRecommendationTypes returns a boolean if a field has been set.
func (o *ListStreamQuery) HasRecommendationTypes() bool {
	return o != nil && o.RecommendationTypes != nil
}

// SetRecommendationTypes gets a reference to the given []string and assigns it to the RecommendationTypes field.
func (o *ListStreamQuery) SetRecommendationTypes(v []string) {
	o.RecommendationTypes = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ListStreamQuery) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ListStreamQuery) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *ListStreamQuery) SetServices(v []string) {
	o.Services = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ListStreamQuery) GetSort() WidgetFieldSort {
	if o == nil || o.Sort == nil {
		var ret WidgetFieldSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetSortOk() (*WidgetFieldSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ListStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given WidgetFieldSort and assigns it to the Sort field.
func (o *ListStreamQuery) SetSort(v WidgetFieldSort) {
	o.Sort = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *ListStreamQuery) GetStates() []ListStreamIssueState {
	if o == nil || o.States == nil {
		var ret []ListStreamIssueState
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetStatesOk() (*[]ListStreamIssueState, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return &o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *ListStreamQuery) HasStates() bool {
	return o != nil && o.States != nil
}

// SetStates gets a reference to the given []ListStreamIssueState and assigns it to the States field.
func (o *ListStreamQuery) SetStates(v []ListStreamIssueState) {
	o.States = v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *ListStreamQuery) GetStatuses() []string {
	if o == nil || o.Statuses == nil {
		var ret []string
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetStatusesOk() (*[]string, bool) {
	if o == nil || o.Statuses == nil {
		return nil, false
	}
	return &o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *ListStreamQuery) HasStatuses() bool {
	return o != nil && o.Statuses != nil
}

// SetStatuses gets a reference to the given []string and assigns it to the Statuses field.
func (o *ListStreamQuery) SetStatuses(v []string) {
	o.Statuses = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ListStreamQuery) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ListStreamQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *ListStreamQuery) SetStorage(v string) {
	o.Storage = &v
}

// GetSuspectedCauses returns the SuspectedCauses field value if set, zero value otherwise.
func (o *ListStreamQuery) GetSuspectedCauses() []string {
	if o == nil || o.SuspectedCauses == nil {
		var ret []string
		return ret
	}
	return o.SuspectedCauses
}

// GetSuspectedCausesOk returns a tuple with the SuspectedCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetSuspectedCausesOk() (*[]string, bool) {
	if o == nil || o.SuspectedCauses == nil {
		return nil, false
	}
	return &o.SuspectedCauses, true
}

// HasSuspectedCauses returns a boolean if a field has been set.
func (o *ListStreamQuery) HasSuspectedCauses() bool {
	return o != nil && o.SuspectedCauses != nil
}

// SetSuspectedCauses gets a reference to the given []string and assigns it to the SuspectedCauses field.
func (o *ListStreamQuery) SetSuspectedCauses(v []string) {
	o.SuspectedCauses = v
}

// GetTeamHandles returns the TeamHandles field value if set, zero value otherwise.
func (o *ListStreamQuery) GetTeamHandles() []string {
	if o == nil || o.TeamHandles == nil {
		var ret []string
		return ret
	}
	return o.TeamHandles
}

// GetTeamHandlesOk returns a tuple with the TeamHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetTeamHandlesOk() (*[]string, bool) {
	if o == nil || o.TeamHandles == nil {
		return nil, false
	}
	return &o.TeamHandles, true
}

// HasTeamHandles returns a boolean if a field has been set.
func (o *ListStreamQuery) HasTeamHandles() bool {
	return o != nil && o.TeamHandles != nil
}

// SetTeamHandles gets a reference to the given []string and assigns it to the TeamHandles field.
func (o *ListStreamQuery) SetTeamHandles(v []string) {
	o.TeamHandles = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ListStreamQuery) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ListStreamQuery) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *ListStreamQuery) SetTeams(v []string) {
	o.Teams = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ListStreamQuery) GetVersion() ListStreamQueryVersion {
	if o == nil || o.Version == nil {
		var ret ListStreamQueryVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetVersionOk() (*ListStreamQueryVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ListStreamQuery) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given ListStreamQueryVersion and assigns it to the Version field.
func (o *ListStreamQuery) SetVersion(v ListStreamQueryVersion) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeUuids != nil {
		toSerialize["assignee_uuids"] = o.AssigneeUuids
	}
	if o.ClusteringPatternFieldPath != nil {
		toSerialize["clustering_pattern_field_path"] = o.ClusteringPatternFieldPath
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["data_source"] = o.DataSource
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.EventSize != nil {
		toSerialize["event_size"] = o.EventSize
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.Persona != nil {
		toSerialize["persona"] = o.Persona
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
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.SuspectedCauses != nil {
		toSerialize["suspected_causes"] = o.SuspectedCauses
	}
	if o.TeamHandles != nil {
		toSerialize["team_handles"] = o.TeamHandles
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeUuids              []string                 `json:"assignee_uuids,omitempty"`
		ClusteringPatternFieldPath *string                  `json:"clustering_pattern_field_path,omitempty"`
		Compute                    []ListStreamComputeItems `json:"compute,omitempty"`
		DataSource                 *ListStreamSource        `json:"data_source"`
		Env                        *string                  `json:"env,omitempty"`
		EventSize                  *WidgetEventSize         `json:"event_size,omitempty"`
		GroupBy                    []ListStreamGroupByItems `json:"group_by,omitempty"`
		Indexes                    []string                 `json:"indexes,omitempty"`
		Persona                    *ListStreamIssuePersona  `json:"persona,omitempty"`
		QueryString                *string                  `json:"query_string"`
		RecommendationTypes        []string                 `json:"recommendation_types,omitempty"`
		Services                   []string                 `json:"services,omitempty"`
		Sort                       *WidgetFieldSort         `json:"sort,omitempty"`
		States                     []ListStreamIssueState   `json:"states,omitempty"`
		Statuses                   []string                 `json:"statuses,omitempty"`
		Storage                    *string                  `json:"storage,omitempty"`
		SuspectedCauses            []string                 `json:"suspected_causes,omitempty"`
		TeamHandles                []string                 `json:"team_handles,omitempty"`
		Teams                      []string                 `json:"teams,omitempty"`
		Version                    *ListStreamQueryVersion  `json:"version,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee_uuids", "clustering_pattern_field_path", "compute", "data_source", "env", "event_size", "group_by", "indexes", "persona", "query_string", "recommendation_types", "services", "sort", "states", "statuses", "storage", "suspected_causes", "team_handles", "teams", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssigneeUuids = all.AssigneeUuids
	o.ClusteringPatternFieldPath = all.ClusteringPatternFieldPath
	o.Compute = all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Env = all.Env
	if all.EventSize != nil && !all.EventSize.IsValid() {
		hasInvalidField = true
	} else {
		o.EventSize = all.EventSize
	}
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	if all.Persona != nil && !all.Persona.IsValid() {
		hasInvalidField = true
	} else {
		o.Persona = all.Persona
	}
	o.QueryString = *all.QueryString
	o.RecommendationTypes = all.RecommendationTypes
	o.Services = all.Services
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.States = all.States
	o.Statuses = all.Statuses
	o.Storage = all.Storage
	o.SuspectedCauses = all.SuspectedCauses
	o.TeamHandles = all.TeamHandles
	o.Teams = all.Teams
	if all.Version != nil && !all.Version.IsValid() {
		hasInvalidField = true
	} else {
		o.Version = all.Version
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
