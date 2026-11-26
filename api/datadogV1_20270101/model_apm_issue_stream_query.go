// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APMIssueStreamQuery Test-only generated schema for the dashboard versioning proof of concept.
type APMIssueStreamQuery struct {
	// Test-only generated schema for the dashboard versioning proof of concept.
	AssigneeUuids []string `json:"assignee_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CodeGenFixStates []string `json:"code_gen_fix_states,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	CrossOrgUuids []string `json:"cross_org_uuids,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	DataSource APMIssueStreamQueryDataSource `json:"data_source"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Indexes []string `json:"indexes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Investigation *APMIssueStreamQueryInvestigation `json:"investigation,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Muted *bool `json:"muted,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Persona *IssueStreamPersona `json:"persona,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	QueryString string `json:"query_string"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	Sort *IssueStreamSort `json:"sort,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	States []IssueStreamState `json:"states,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	SuspectedCauses []string `json:"suspected_causes,omitempty"`
	// Test-only generated schema for the dashboard versioning proof of concept.
	TeamHandles []string `json:"team_handles,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewAPMIssueStreamQuery instantiates a new APMIssueStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPMIssueStreamQuery(dataSource APMIssueStreamQueryDataSource, queryString string) *APMIssueStreamQuery {
	this := APMIssueStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewAPMIssueStreamQueryWithDefaults instantiates a new APMIssueStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPMIssueStreamQueryWithDefaults() *APMIssueStreamQuery {
	this := APMIssueStreamQuery{}
	return &this
}

// GetAssigneeUuids returns the AssigneeUuids field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetAssigneeUuids() []string {
	if o == nil || o.AssigneeUuids == nil {
		var ret []string
		return ret
	}
	return o.AssigneeUuids
}

// GetAssigneeUuidsOk returns a tuple with the AssigneeUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetAssigneeUuidsOk() (*[]string, bool) {
	if o == nil || o.AssigneeUuids == nil {
		return nil, false
	}
	return &o.AssigneeUuids, true
}

// HasAssigneeUuids returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasAssigneeUuids() bool {
	return o != nil && o.AssigneeUuids != nil
}

// SetAssigneeUuids gets a reference to the given []string and assigns it to the AssigneeUuids field.
func (o *APMIssueStreamQuery) SetAssigneeUuids(v []string) {
	o.AssigneeUuids = v
}

// GetCodeGenFixStates returns the CodeGenFixStates field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetCodeGenFixStates() []string {
	if o == nil || o.CodeGenFixStates == nil {
		var ret []string
		return ret
	}
	return o.CodeGenFixStates
}

// GetCodeGenFixStatesOk returns a tuple with the CodeGenFixStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetCodeGenFixStatesOk() (*[]string, bool) {
	if o == nil || o.CodeGenFixStates == nil {
		return nil, false
	}
	return &o.CodeGenFixStates, true
}

// HasCodeGenFixStates returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasCodeGenFixStates() bool {
	return o != nil && o.CodeGenFixStates != nil
}

// SetCodeGenFixStates gets a reference to the given []string and assigns it to the CodeGenFixStates field.
func (o *APMIssueStreamQuery) SetCodeGenFixStates(v []string) {
	o.CodeGenFixStates = v
}

// GetCrossOrgUuids returns the CrossOrgUuids field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetCrossOrgUuids() []string {
	if o == nil || o.CrossOrgUuids == nil {
		var ret []string
		return ret
	}
	return o.CrossOrgUuids
}

// GetCrossOrgUuidsOk returns a tuple with the CrossOrgUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetCrossOrgUuidsOk() (*[]string, bool) {
	if o == nil || o.CrossOrgUuids == nil {
		return nil, false
	}
	return &o.CrossOrgUuids, true
}

// HasCrossOrgUuids returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasCrossOrgUuids() bool {
	return o != nil && o.CrossOrgUuids != nil
}

// SetCrossOrgUuids gets a reference to the given []string and assigns it to the CrossOrgUuids field.
func (o *APMIssueStreamQuery) SetCrossOrgUuids(v []string) {
	o.CrossOrgUuids = v
}

// GetDataSource returns the DataSource field value.
func (o *APMIssueStreamQuery) GetDataSource() APMIssueStreamQueryDataSource {
	if o == nil {
		var ret APMIssueStreamQueryDataSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetDataSourceOk() (*APMIssueStreamQueryDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *APMIssueStreamQuery) SetDataSource(v APMIssueStreamQueryDataSource) {
	o.DataSource = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *APMIssueStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetInvestigation returns the Investigation field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetInvestigation() APMIssueStreamQueryInvestigation {
	if o == nil || o.Investigation == nil {
		var ret APMIssueStreamQueryInvestigation
		return ret
	}
	return *o.Investigation
}

// GetInvestigationOk returns a tuple with the Investigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetInvestigationOk() (*APMIssueStreamQueryInvestigation, bool) {
	if o == nil || o.Investigation == nil {
		return nil, false
	}
	return o.Investigation, true
}

// HasInvestigation returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasInvestigation() bool {
	return o != nil && o.Investigation != nil
}

// SetInvestigation gets a reference to the given APMIssueStreamQueryInvestigation and assigns it to the Investigation field.
func (o *APMIssueStreamQuery) SetInvestigation(v APMIssueStreamQueryInvestigation) {
	o.Investigation = &v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetMuted() bool {
	if o == nil || o.Muted == nil {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetMutedOk() (*bool, bool) {
	if o == nil || o.Muted == nil {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasMuted() bool {
	return o != nil && o.Muted != nil
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *APMIssueStreamQuery) SetMuted(v bool) {
	o.Muted = &v
}

// GetPersona returns the Persona field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetPersona() IssueStreamPersona {
	if o == nil || o.Persona == nil {
		var ret IssueStreamPersona
		return ret
	}
	return *o.Persona
}

// GetPersonaOk returns a tuple with the Persona field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetPersonaOk() (*IssueStreamPersona, bool) {
	if o == nil || o.Persona == nil {
		return nil, false
	}
	return o.Persona, true
}

// HasPersona returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasPersona() bool {
	return o != nil && o.Persona != nil
}

// SetPersona gets a reference to the given IssueStreamPersona and assigns it to the Persona field.
func (o *APMIssueStreamQuery) SetPersona(v IssueStreamPersona) {
	o.Persona = &v
}

// GetQueryString returns the QueryString field value.
func (o *APMIssueStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *APMIssueStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetSort() IssueStreamSort {
	if o == nil || o.Sort == nil {
		var ret IssueStreamSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetSortOk() (*IssueStreamSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given IssueStreamSort and assigns it to the Sort field.
func (o *APMIssueStreamQuery) SetSort(v IssueStreamSort) {
	o.Sort = &v
}

// GetStates returns the States field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetStates() []IssueStreamState {
	if o == nil || o.States == nil {
		var ret []IssueStreamState
		return ret
	}
	return o.States
}

// GetStatesOk returns a tuple with the States field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetStatesOk() (*[]IssueStreamState, bool) {
	if o == nil || o.States == nil {
		return nil, false
	}
	return &o.States, true
}

// HasStates returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasStates() bool {
	return o != nil && o.States != nil
}

// SetStates gets a reference to the given []IssueStreamState and assigns it to the States field.
func (o *APMIssueStreamQuery) SetStates(v []IssueStreamState) {
	o.States = v
}

// GetSuspectedCauses returns the SuspectedCauses field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetSuspectedCauses() []string {
	if o == nil || o.SuspectedCauses == nil {
		var ret []string
		return ret
	}
	return o.SuspectedCauses
}

// GetSuspectedCausesOk returns a tuple with the SuspectedCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetSuspectedCausesOk() (*[]string, bool) {
	if o == nil || o.SuspectedCauses == nil {
		return nil, false
	}
	return &o.SuspectedCauses, true
}

// HasSuspectedCauses returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasSuspectedCauses() bool {
	return o != nil && o.SuspectedCauses != nil
}

// SetSuspectedCauses gets a reference to the given []string and assigns it to the SuspectedCauses field.
func (o *APMIssueStreamQuery) SetSuspectedCauses(v []string) {
	o.SuspectedCauses = v
}

// GetTeamHandles returns the TeamHandles field value if set, zero value otherwise.
func (o *APMIssueStreamQuery) GetTeamHandles() []string {
	if o == nil || o.TeamHandles == nil {
		var ret []string
		return ret
	}
	return o.TeamHandles
}

// GetTeamHandlesOk returns a tuple with the TeamHandles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APMIssueStreamQuery) GetTeamHandlesOk() (*[]string, bool) {
	if o == nil || o.TeamHandles == nil {
		return nil, false
	}
	return &o.TeamHandles, true
}

// HasTeamHandles returns a boolean if a field has been set.
func (o *APMIssueStreamQuery) HasTeamHandles() bool {
	return o != nil && o.TeamHandles != nil
}

// SetTeamHandles gets a reference to the given []string and assigns it to the TeamHandles field.
func (o *APMIssueStreamQuery) SetTeamHandles(v []string) {
	o.TeamHandles = v
}

// MarshalJSON serializes the struct using spec logic.
func (o APMIssueStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssigneeUuids != nil {
		toSerialize["assignee_uuids"] = o.AssigneeUuids
	}
	if o.CodeGenFixStates != nil {
		toSerialize["code_gen_fix_states"] = o.CodeGenFixStates
	}
	if o.CrossOrgUuids != nil {
		toSerialize["cross_org_uuids"] = o.CrossOrgUuids
	}
	toSerialize["data_source"] = o.DataSource
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	if o.Investigation != nil {
		toSerialize["investigation"] = o.Investigation
	}
	if o.Muted != nil {
		toSerialize["muted"] = o.Muted
	}
	if o.Persona != nil {
		toSerialize["persona"] = o.Persona
	}
	toSerialize["query_string"] = o.QueryString
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.States != nil {
		toSerialize["states"] = o.States
	}
	if o.SuspectedCauses != nil {
		toSerialize["suspected_causes"] = o.SuspectedCauses
	}
	if o.TeamHandles != nil {
		toSerialize["team_handles"] = o.TeamHandles
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *APMIssueStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssigneeUuids    []string                          `json:"assignee_uuids,omitempty"`
		CodeGenFixStates []string                          `json:"code_gen_fix_states,omitempty"`
		CrossOrgUuids    []string                          `json:"cross_org_uuids,omitempty"`
		DataSource       *APMIssueStreamQueryDataSource    `json:"data_source"`
		Indexes          []string                          `json:"indexes,omitempty"`
		Investigation    *APMIssueStreamQueryInvestigation `json:"investigation,omitempty"`
		Muted            *bool                             `json:"muted,omitempty"`
		Persona          *IssueStreamPersona               `json:"persona,omitempty"`
		QueryString      *string                           `json:"query_string"`
		Sort             *IssueStreamSort                  `json:"sort,omitempty"`
		States           []IssueStreamState                `json:"states,omitempty"`
		SuspectedCauses  []string                          `json:"suspected_causes,omitempty"`
		TeamHandles      []string                          `json:"team_handles,omitempty"`
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
	o.AssigneeUuids = all.AssigneeUuids
	o.CodeGenFixStates = all.CodeGenFixStates
	o.CrossOrgUuids = all.CrossOrgUuids
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	o.Indexes = all.Indexes
	if all.Investigation != nil && all.Investigation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Investigation = all.Investigation
	o.Muted = all.Muted
	if all.Persona != nil && !all.Persona.IsValid() {
		hasInvalidField = true
	} else {
		o.Persona = all.Persona
	}
	o.QueryString = *all.QueryString
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.States = all.States
	o.SuspectedCauses = all.SuspectedCauses
	o.TeamHandles = all.TeamHandles

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
