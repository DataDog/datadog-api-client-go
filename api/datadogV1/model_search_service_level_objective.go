// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SearchServiceLevelObjective A service level objective object includes a service level indicator, thresholds
// for one or more timeframes, and metadata (`name`, `description`, `tags`, etc.).
type SearchServiceLevelObjective struct {
	// A list of tags associated with this service level objective.
	// Always included in service level objective responses (but may be empty).
	// Optional in create/update requests.
	AllTags []string `json:"all_tags,omitempty"`
	// Creation timestamp (UNIX time in seconds)
	//
	// Always included in service level objective responses.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The creator of the SLO
	Creator NullableSLOCreator `json:"creator,omitempty"`
	// A user-defined description of the service level objective.
	//
	// Always included in service level objective responses (but may be `null`).
	// Optional in create/update requests.
	Description datadog.NullableString `json:"description,omitempty"`
	// A list of (up to 100) monitor groups that narrow the scope of a monitor service level objective.
	//
	// Included in service level objective responses if it is not empty. Optional in
	// create/update requests for monitor service level objectives, but may only be
	// used when then length of the `monitor_ids` field is one.
	Groups []string `json:"groups,omitempty"`
	// A unique identifier for the service level objective object.
	//
	// Always included in service level objective responses.
	Id *string `json:"id,omitempty"`
	// Modification timestamp (UNIX time in seconds)
	//
	// Always included in service level objective responses.
	ModifiedAt *int64 `json:"modified_at,omitempty"`
	// A list of monitor ids that defines the scope of a monitor service level
	// objective. **Required if type is `monitor`**.
	MonitorIds []int64 `json:"monitor_ids,omitempty"`
	// The name of the service level objective object.
	Name *string `json:"name,omitempty"`
	// calculated status and error budget remaining.
	OverallStatus []SLOOverallStatuses `json:"overall_status,omitempty"`
	// A metric-based SLO. **Required if type is `metric`**. Note that Datadog only allows the sum by aggregator
	// to be used because this will sum up all request counts instead of averaging them, or taking the max or
	// min of all of those requests.
	Query NullableSearchSLOQuery `json:"query,omitempty"`
	// The thresholds (timeframes and associated targets) for this service level
	// objective object.
	Thresholds []SearchSLOThreshold `json:"thresholds,omitempty"`
	// The type of the service level objective.
	Type *SLOType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewSearchServiceLevelObjective instantiates a new SearchServiceLevelObjective object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSearchServiceLevelObjective() *SearchServiceLevelObjective {
	this := SearchServiceLevelObjective{}
	return &this
}

// NewSearchServiceLevelObjectiveWithDefaults instantiates a new SearchServiceLevelObjective object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSearchServiceLevelObjectiveWithDefaults() *SearchServiceLevelObjective {
	this := SearchServiceLevelObjective{}
	return &this
}

// GetAllTags returns the AllTags field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetAllTags() []string {
	if o == nil || o.AllTags == nil {
		var ret []string
		return ret
	}
	return o.AllTags
}

// GetAllTagsOk returns a tuple with the AllTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetAllTagsOk() (*[]string, bool) {
	if o == nil || o.AllTags == nil {
		return nil, false
	}
	return &o.AllTags, true
}

// HasAllTags returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasAllTags() bool {
	if o != nil && o.AllTags != nil {
		return true
	}

	return false
}

// SetAllTags gets a reference to the given []string and assigns it to the AllTags field.
func (o *SearchServiceLevelObjective) SetAllTags(v []string) {
	o.AllTags = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *SearchServiceLevelObjective) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchServiceLevelObjective) GetCreator() SLOCreator {
	if o == nil || o.Creator.Get() == nil {
		var ret SLOCreator
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchServiceLevelObjective) GetCreatorOk() (*SLOCreator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given NullableSLOCreator and assigns it to the Creator field.
func (o *SearchServiceLevelObjective) SetCreator(v SLOCreator) {
	o.Creator.Set(&v)
}

// SetCreatorNil sets the value for Creator to be an explicit nil.
func (o *SearchServiceLevelObjective) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil.
func (o *SearchServiceLevelObjective) UnsetCreator() {
	o.Creator.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchServiceLevelObjective) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchServiceLevelObjective) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *SearchServiceLevelObjective) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *SearchServiceLevelObjective) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *SearchServiceLevelObjective) UnsetDescription() {
	o.Description.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchServiceLevelObjective) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchServiceLevelObjective) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *SearchServiceLevelObjective) SetGroups(v []string) {
	o.Groups = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchServiceLevelObjective) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetModifiedAt() int64 {
	if o == nil || o.ModifiedAt == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetModifiedAtOk() (*int64, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *SearchServiceLevelObjective) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetMonitorIds returns the MonitorIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchServiceLevelObjective) GetMonitorIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.MonitorIds
}

// GetMonitorIdsOk returns a tuple with the MonitorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchServiceLevelObjective) GetMonitorIdsOk() (*[]int64, bool) {
	if o == nil || o.MonitorIds == nil {
		return nil, false
	}
	return &o.MonitorIds, true
}

// HasMonitorIds returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasMonitorIds() bool {
	if o != nil && o.MonitorIds != nil {
		return true
	}

	return false
}

// SetMonitorIds gets a reference to the given []int64 and assigns it to the MonitorIds field.
func (o *SearchServiceLevelObjective) SetMonitorIds(v []int64) {
	o.MonitorIds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchServiceLevelObjective) SetName(v string) {
	o.Name = &v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetOverallStatus() []SLOOverallStatuses {
	if o == nil || o.OverallStatus == nil {
		var ret []SLOOverallStatuses
		return ret
	}
	return o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetOverallStatusOk() (*[]SLOOverallStatuses, bool) {
	if o == nil || o.OverallStatus == nil {
		return nil, false
	}
	return &o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasOverallStatus() bool {
	if o != nil && o.OverallStatus != nil {
		return true
	}

	return false
}

// SetOverallStatus gets a reference to the given []SLOOverallStatuses and assigns it to the OverallStatus field.
func (o *SearchServiceLevelObjective) SetOverallStatus(v []SLOOverallStatuses) {
	o.OverallStatus = v
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchServiceLevelObjective) GetQuery() SearchSLOQuery {
	if o == nil || o.Query.Get() == nil {
		var ret SearchSLOQuery
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SearchServiceLevelObjective) GetQueryOk() (*SearchSLOQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableSearchSLOQuery and assigns it to the Query field.
func (o *SearchServiceLevelObjective) SetQuery(v SearchSLOQuery) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil.
func (o *SearchServiceLevelObjective) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil.
func (o *SearchServiceLevelObjective) UnsetQuery() {
	o.Query.Unset()
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetThresholds() []SearchSLOThreshold {
	if o == nil || o.Thresholds == nil {
		var ret []SearchSLOThreshold
		return ret
	}
	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetThresholdsOk() (*[]SearchSLOThreshold, bool) {
	if o == nil || o.Thresholds == nil {
		return nil, false
	}
	return &o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasThresholds() bool {
	if o != nil && o.Thresholds != nil {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given []SearchSLOThreshold and assigns it to the Thresholds field.
func (o *SearchServiceLevelObjective) SetThresholds(v []SearchSLOThreshold) {
	o.Thresholds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SearchServiceLevelObjective) GetType() SLOType {
	if o == nil || o.Type == nil {
		var ret SLOType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchServiceLevelObjective) GetTypeOk() (*SLOType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchServiceLevelObjective) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given SLOType and assigns it to the Type field.
func (o *SearchServiceLevelObjective) SetType(v SLOType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SearchServiceLevelObjective) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AllTags != nil {
		toSerialize["all_tags"] = o.AllTags
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Creator.IsSet() {
		toSerialize["creator"] = o.Creator.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.MonitorIds != nil {
		toSerialize["monitor_ids"] = o.MonitorIds
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverallStatus != nil {
		toSerialize["overall_status"] = o.OverallStatus
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Thresholds != nil {
		toSerialize["thresholds"] = o.Thresholds
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SearchServiceLevelObjective) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AllTags       []string               `json:"all_tags,omitempty"`
		CreatedAt     *int64                 `json:"created_at,omitempty"`
		Creator       NullableSLOCreator     `json:"creator,omitempty"`
		Description   datadog.NullableString `json:"description,omitempty"`
		Groups        []string               `json:"groups,omitempty"`
		Id            *string                `json:"id,omitempty"`
		ModifiedAt    *int64                 `json:"modified_at,omitempty"`
		MonitorIds    []int64                `json:"monitor_ids,omitempty"`
		Name          *string                `json:"name,omitempty"`
		OverallStatus []SLOOverallStatuses   `json:"overall_status,omitempty"`
		Query         NullableSearchSLOQuery `json:"query,omitempty"`
		Thresholds    []SearchSLOThreshold   `json:"thresholds,omitempty"`
		Type          *SLOType               `json:"type,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; v != nil && !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.AllTags = all.AllTags
	o.CreatedAt = all.CreatedAt
	o.Creator = all.Creator
	o.Description = all.Description
	o.Groups = all.Groups
	o.Id = all.Id
	o.ModifiedAt = all.ModifiedAt
	o.MonitorIds = all.MonitorIds
	o.Name = all.Name
	o.OverallStatus = all.OverallStatus
	o.Query = all.Query
	o.Thresholds = all.Thresholds
	o.Type = all.Type
	return nil
}
