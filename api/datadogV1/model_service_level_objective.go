// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceLevelObjective A service level objective object includes a service level indicator, thresholds
// for one or more timeframes, and metadata (`name`, `description`, `tags`, etc.).
type ServiceLevelObjective struct {
	// Creation timestamp (UNIX time in seconds)
	//
	// Always included in service level objective responses.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Object describing the creator of the shared element.
	Creator *Creator `json:"creator,omitempty"`
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
	// The union of monitor tags for all monitors referenced by the `monitor_ids`
	// field.
	// Always included in service level objective responses for monitor-based service level
	// objectives (but may be empty). Ignored in create/update requests. Does not
	// affect which monitors are included in the service level objective (that is
	// determined entirely by the `monitor_ids` field).
	MonitorTags []string `json:"monitor_tags,omitempty"`
	// The name of the service level objective object.
	Name string `json:"name"`
	// A metric-based SLO. **Required if type is `metric`**. Note that Datadog only allows the sum by aggregator
	// to be used because this will sum up all request counts instead of averaging them, or taking the max or
	// min of all of those requests.
	Query *ServiceLevelObjectiveQuery `json:"query,omitempty"`
	// A list of tags associated with this service level objective.
	// Always included in service level objective responses (but may be empty).
	// Optional in create/update requests.
	Tags []string `json:"tags,omitempty"`
	// The target threshold such that when the service level indicator is above this
	// threshold over the given timeframe, the objective is being met.
	TargetThreshold *float64 `json:"target_threshold,omitempty"`
	// The thresholds (timeframes and associated targets) for this service level
	// objective object.
	Thresholds []SLOThreshold `json:"thresholds"`
	// The SLO time window options.
	Timeframe *SLOTimeframe `json:"timeframe,omitempty"`
	// The type of the service level objective.
	Type SLOType `json:"type"`
	// The optional warning threshold such that when the service level indicator is
	// below this value for the given threshold, but above the target threshold, the
	// objective appears in a "warning" state. This value must be greater than the target
	// threshold.
	WarningThreshold *float64 `json:"warning_threshold,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewServiceLevelObjective instantiates a new ServiceLevelObjective object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceLevelObjective(name string, thresholds []SLOThreshold, typeVar SLOType) *ServiceLevelObjective {
	this := ServiceLevelObjective{}
	this.Name = name
	this.Thresholds = thresholds
	this.Type = typeVar
	return &this
}

// NewServiceLevelObjectiveWithDefaults instantiates a new ServiceLevelObjective object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceLevelObjectiveWithDefaults() *ServiceLevelObjective {
	this := ServiceLevelObjective{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ServiceLevelObjective) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetCreator() Creator {
	if o == nil || o.Creator == nil {
		var ret Creator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetCreatorOk() (*Creator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given Creator and assigns it to the Creator field.
func (o *ServiceLevelObjective) SetCreator(v Creator) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLevelObjective) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ServiceLevelObjective) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ServiceLevelObjective) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ServiceLevelObjective) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ServiceLevelObjective) UnsetDescription() {
	o.Description.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasGroups() bool {
	return o != nil && o.Groups != nil
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ServiceLevelObjective) SetGroups(v []string) {
	o.Groups = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceLevelObjective) SetId(v string) {
	o.Id = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetModifiedAt() int64 {
	if o == nil || o.ModifiedAt == nil {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetModifiedAtOk() (*int64, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *ServiceLevelObjective) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetMonitorIds returns the MonitorIds field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetMonitorIds() []int64 {
	if o == nil || o.MonitorIds == nil {
		var ret []int64
		return ret
	}
	return o.MonitorIds
}

// GetMonitorIdsOk returns a tuple with the MonitorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetMonitorIdsOk() (*[]int64, bool) {
	if o == nil || o.MonitorIds == nil {
		return nil, false
	}
	return &o.MonitorIds, true
}

// HasMonitorIds returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasMonitorIds() bool {
	return o != nil && o.MonitorIds != nil
}

// SetMonitorIds gets a reference to the given []int64 and assigns it to the MonitorIds field.
func (o *ServiceLevelObjective) SetMonitorIds(v []int64) {
	o.MonitorIds = v
}

// GetMonitorTags returns the MonitorTags field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetMonitorTags() []string {
	if o == nil || o.MonitorTags == nil {
		var ret []string
		return ret
	}
	return o.MonitorTags
}

// GetMonitorTagsOk returns a tuple with the MonitorTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetMonitorTagsOk() (*[]string, bool) {
	if o == nil || o.MonitorTags == nil {
		return nil, false
	}
	return &o.MonitorTags, true
}

// HasMonitorTags returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasMonitorTags() bool {
	return o != nil && o.MonitorTags != nil
}

// SetMonitorTags gets a reference to the given []string and assigns it to the MonitorTags field.
func (o *ServiceLevelObjective) SetMonitorTags(v []string) {
	o.MonitorTags = v
}

// GetName returns the Name field value.
func (o *ServiceLevelObjective) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ServiceLevelObjective) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetQuery() ServiceLevelObjectiveQuery {
	if o == nil || o.Query == nil {
		var ret ServiceLevelObjectiveQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetQueryOk() (*ServiceLevelObjectiveQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given ServiceLevelObjectiveQuery and assigns it to the Query field.
func (o *ServiceLevelObjective) SetQuery(v ServiceLevelObjectiveQuery) {
	o.Query = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceLevelObjective) SetTags(v []string) {
	o.Tags = v
}

// GetTargetThreshold returns the TargetThreshold field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetTargetThreshold() float64 {
	if o == nil || o.TargetThreshold == nil {
		var ret float64
		return ret
	}
	return *o.TargetThreshold
}

// GetTargetThresholdOk returns a tuple with the TargetThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetTargetThresholdOk() (*float64, bool) {
	if o == nil || o.TargetThreshold == nil {
		return nil, false
	}
	return o.TargetThreshold, true
}

// HasTargetThreshold returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasTargetThreshold() bool {
	return o != nil && o.TargetThreshold != nil
}

// SetTargetThreshold gets a reference to the given float64 and assigns it to the TargetThreshold field.
func (o *ServiceLevelObjective) SetTargetThreshold(v float64) {
	o.TargetThreshold = &v
}

// GetThresholds returns the Thresholds field value.
func (o *ServiceLevelObjective) GetThresholds() []SLOThreshold {
	if o == nil {
		var ret []SLOThreshold
		return ret
	}
	return o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetThresholdsOk() (*[]SLOThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thresholds, true
}

// SetThresholds sets field value.
func (o *ServiceLevelObjective) SetThresholds(v []SLOThreshold) {
	o.Thresholds = v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetTimeframe() SLOTimeframe {
	if o == nil || o.Timeframe == nil {
		var ret SLOTimeframe
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetTimeframeOk() (*SLOTimeframe, bool) {
	if o == nil || o.Timeframe == nil {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasTimeframe() bool {
	return o != nil && o.Timeframe != nil
}

// SetTimeframe gets a reference to the given SLOTimeframe and assigns it to the Timeframe field.
func (o *ServiceLevelObjective) SetTimeframe(v SLOTimeframe) {
	o.Timeframe = &v
}

// GetType returns the Type field value.
func (o *ServiceLevelObjective) GetType() SLOType {
	if o == nil {
		var ret SLOType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetTypeOk() (*SLOType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceLevelObjective) SetType(v SLOType) {
	o.Type = v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *ServiceLevelObjective) GetWarningThreshold() float64 {
	if o == nil || o.WarningThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjective) GetWarningThresholdOk() (*float64, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *ServiceLevelObjective) HasWarningThreshold() bool {
	return o != nil && o.WarningThreshold != nil
}

// SetWarningThreshold gets a reference to the given float64 and assigns it to the WarningThreshold field.
func (o *ServiceLevelObjective) SetWarningThreshold(v float64) {
	o.WarningThreshold = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceLevelObjective) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
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
	if o.MonitorTags != nil {
		toSerialize["monitor_tags"] = o.MonitorTags
	}
	toSerialize["name"] = o.Name
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetThreshold != nil {
		toSerialize["target_threshold"] = o.TargetThreshold
	}
	toSerialize["thresholds"] = o.Thresholds
	if o.Timeframe != nil {
		toSerialize["timeframe"] = o.Timeframe
	}
	toSerialize["type"] = o.Type
	if o.WarningThreshold != nil {
		toSerialize["warning_threshold"] = o.WarningThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceLevelObjective) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		CreatedAt        *int64                      `json:"created_at,omitempty"`
		Creator          *Creator                    `json:"creator,omitempty"`
		Description      datadog.NullableString      `json:"description,omitempty"`
		Groups           []string                    `json:"groups,omitempty"`
		Id               *string                     `json:"id,omitempty"`
		ModifiedAt       *int64                      `json:"modified_at,omitempty"`
		MonitorIds       []int64                     `json:"monitor_ids,omitempty"`
		MonitorTags      []string                    `json:"monitor_tags,omitempty"`
		Name             *string                     `json:"name"`
		Query            *ServiceLevelObjectiveQuery `json:"query,omitempty"`
		Tags             []string                    `json:"tags,omitempty"`
		TargetThreshold  *float64                    `json:"target_threshold,omitempty"`
		Thresholds       *[]SLOThreshold             `json:"thresholds"`
		Timeframe        *SLOTimeframe               `json:"timeframe,omitempty"`
		Type             *SLOType                    `json:"type"`
		WarningThreshold *float64                    `json:"warning_threshold,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Thresholds == nil {
		return fmt.Errorf("required field thresholds missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "creator", "description", "groups", "id", "modified_at", "monitor_ids", "monitor_tags", "name", "query", "tags", "target_threshold", "thresholds", "timeframe", "type", "warning_threshold"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Description = all.Description
	o.Groups = all.Groups
	o.Id = all.Id
	o.ModifiedAt = all.ModifiedAt
	o.MonitorIds = all.MonitorIds
	o.MonitorTags = all.MonitorTags
	o.Name = *all.Name
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
	o.Tags = all.Tags
	o.TargetThreshold = all.TargetThreshold
	o.Thresholds = *all.Thresholds
	if v := all.Timeframe; v != nil && !v.IsValid() {
		hasInvalidField = true
	} else {
		o.Timeframe = all.Timeframe
	}
	if v := all.Type; !v.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.WarningThreshold = all.WarningThreshold

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}

	return nil
}
