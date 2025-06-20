// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Monitor Object describing a monitor.
type Monitor struct {
	// Timestamp of the monitor creation.
	Created *time.Time `json:"created,omitempty"`
	// Object describing the creator of the shared element.
	Creator *Creator `json:"creator,omitempty"`
	// Whether or not the monitor is deleted. (Always `null`)
	Deleted datadog.NullableTime `json:"deleted,omitempty"`
	// Indicates whether the monitor is in a draft or published state.
	//
	// `draft`: The monitor appears as Draft and does not send notifications.
	// `published`: The monitor is active and evaluates conditions and notify as configured.
	//
	// This field is in preview. The draft value is only available to customers with the feature enabled.
	//
	DraftStatus *MonitorDraftStatus `json:"draft_status,omitempty"`
	// ID of this monitor.
	Id *int64 `json:"id,omitempty"`
	// A list of active v1 downtimes that match this monitor.
	MatchingDowntimes []MatchingDowntime `json:"matching_downtimes,omitempty"`
	// A message to include with notifications for this monitor.
	Message *string `json:"message,omitempty"`
	// Last timestamp when the monitor was edited.
	Modified *time.Time `json:"modified,omitempty"`
	// Whether or not the monitor is broken down on different groups.
	Multi *bool `json:"multi,omitempty"`
	// The monitor name.
	Name *string `json:"name,omitempty"`
	// List of options associated with your monitor.
	Options *MonitorOptions `json:"options,omitempty"`
	// The different states your monitor can be in.
	OverallState *MonitorOverallStates `json:"overall_state,omitempty"`
	// Integer from 1 (high) to 5 (low) indicating alert severity.
	Priority datadog.NullableInt64 `json:"priority,omitempty"`
	// The monitor query.
	Query string `json:"query"`
	// A list of unique role identifiers to define which roles are allowed to edit the monitor. The unique identifiers for all roles can be pulled from the [Roles API](https://docs.datadoghq.com/api/latest/roles/#list-roles) and are located in the `data.id` field. Editing a monitor includes any updates to the monitor configuration, monitor deletion, and muting of the monitor for any amount of time. You can use the [Restriction Policies API](https://docs.datadoghq.com/api/latest/restriction-policies/) to manage write authorization for individual monitors by teams and users, in addition to roles.
	RestrictedRoles datadog.NullableList[string] `json:"restricted_roles,omitempty"`
	// Wrapper object with the different monitor states.
	State *MonitorState `json:"state,omitempty"`
	// Tags associated to your monitor.
	Tags []string `json:"tags,omitempty"`
	// The type of the monitor. For more information about `type`, see the [monitor options](https://docs.datadoghq.com/monitors/guide/monitor_api_options/) docs.
	Type MonitorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitor instantiates a new Monitor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitor(query string, typeVar MonitorType) *Monitor {
	this := Monitor{}
	var draftStatus MonitorDraftStatus = MONITORDRAFTSTATUS_PUBLISHED
	this.DraftStatus = &draftStatus
	this.Query = query
	this.Type = typeVar
	return &this
}

// NewMonitorWithDefaults instantiates a new Monitor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorWithDefaults() *Monitor {
	this := Monitor{}
	var draftStatus MonitorDraftStatus = MONITORDRAFTSTATUS_PUBLISHED
	this.DraftStatus = &draftStatus
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Monitor) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Monitor) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Monitor) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Monitor) GetCreator() Creator {
	if o == nil || o.Creator == nil {
		var ret Creator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetCreatorOk() (*Creator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Monitor) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given Creator and assigns it to the Creator field.
func (o *Monitor) SetCreator(v Creator) {
	o.Creator = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Monitor) GetDeleted() time.Time {
	if o == nil || o.Deleted.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Monitor) GetDeletedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// HasDeleted returns a boolean if a field has been set.
func (o *Monitor) HasDeleted() bool {
	return o != nil && o.Deleted.IsSet()
}

// SetDeleted gets a reference to the given datadog.NullableTime and assigns it to the Deleted field.
func (o *Monitor) SetDeleted(v time.Time) {
	o.Deleted.Set(&v)
}

// SetDeletedNil sets the value for Deleted to be an explicit nil.
func (o *Monitor) SetDeletedNil() {
	o.Deleted.Set(nil)
}

// UnsetDeleted ensures that no value is present for Deleted, not even an explicit nil.
func (o *Monitor) UnsetDeleted() {
	o.Deleted.Unset()
}

// GetDraftStatus returns the DraftStatus field value if set, zero value otherwise.
func (o *Monitor) GetDraftStatus() MonitorDraftStatus {
	if o == nil || o.DraftStatus == nil {
		var ret MonitorDraftStatus
		return ret
	}
	return *o.DraftStatus
}

// GetDraftStatusOk returns a tuple with the DraftStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetDraftStatusOk() (*MonitorDraftStatus, bool) {
	if o == nil || o.DraftStatus == nil {
		return nil, false
	}
	return o.DraftStatus, true
}

// HasDraftStatus returns a boolean if a field has been set.
func (o *Monitor) HasDraftStatus() bool {
	return o != nil && o.DraftStatus != nil
}

// SetDraftStatus gets a reference to the given MonitorDraftStatus and assigns it to the DraftStatus field.
func (o *Monitor) SetDraftStatus(v MonitorDraftStatus) {
	o.DraftStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Monitor) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Monitor) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Monitor) SetId(v int64) {
	o.Id = &v
}

// GetMatchingDowntimes returns the MatchingDowntimes field value if set, zero value otherwise.
func (o *Monitor) GetMatchingDowntimes() []MatchingDowntime {
	if o == nil || o.MatchingDowntimes == nil {
		var ret []MatchingDowntime
		return ret
	}
	return o.MatchingDowntimes
}

// GetMatchingDowntimesOk returns a tuple with the MatchingDowntimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetMatchingDowntimesOk() (*[]MatchingDowntime, bool) {
	if o == nil || o.MatchingDowntimes == nil {
		return nil, false
	}
	return &o.MatchingDowntimes, true
}

// HasMatchingDowntimes returns a boolean if a field has been set.
func (o *Monitor) HasMatchingDowntimes() bool {
	return o != nil && o.MatchingDowntimes != nil
}

// SetMatchingDowntimes gets a reference to the given []MatchingDowntime and assigns it to the MatchingDowntimes field.
func (o *Monitor) SetMatchingDowntimes(v []MatchingDowntime) {
	o.MatchingDowntimes = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Monitor) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Monitor) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Monitor) SetMessage(v string) {
	o.Message = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *Monitor) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *Monitor) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *Monitor) SetModified(v time.Time) {
	o.Modified = &v
}

// GetMulti returns the Multi field value if set, zero value otherwise.
func (o *Monitor) GetMulti() bool {
	if o == nil || o.Multi == nil {
		var ret bool
		return ret
	}
	return *o.Multi
}

// GetMultiOk returns a tuple with the Multi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetMultiOk() (*bool, bool) {
	if o == nil || o.Multi == nil {
		return nil, false
	}
	return o.Multi, true
}

// HasMulti returns a boolean if a field has been set.
func (o *Monitor) HasMulti() bool {
	return o != nil && o.Multi != nil
}

// SetMulti gets a reference to the given bool and assigns it to the Multi field.
func (o *Monitor) SetMulti(v bool) {
	o.Multi = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Monitor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Monitor) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Monitor) SetName(v string) {
	o.Name = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Monitor) GetOptions() MonitorOptions {
	if o == nil || o.Options == nil {
		var ret MonitorOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetOptionsOk() (*MonitorOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Monitor) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given MonitorOptions and assigns it to the Options field.
func (o *Monitor) SetOptions(v MonitorOptions) {
	o.Options = &v
}

// GetOverallState returns the OverallState field value if set, zero value otherwise.
func (o *Monitor) GetOverallState() MonitorOverallStates {
	if o == nil || o.OverallState == nil {
		var ret MonitorOverallStates
		return ret
	}
	return *o.OverallState
}

// GetOverallStateOk returns a tuple with the OverallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetOverallStateOk() (*MonitorOverallStates, bool) {
	if o == nil || o.OverallState == nil {
		return nil, false
	}
	return o.OverallState, true
}

// HasOverallState returns a boolean if a field has been set.
func (o *Monitor) HasOverallState() bool {
	return o != nil && o.OverallState != nil
}

// SetOverallState gets a reference to the given MonitorOverallStates and assigns it to the OverallState field.
func (o *Monitor) SetOverallState(v MonitorOverallStates) {
	o.OverallState = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Monitor) GetPriority() int64 {
	if o == nil || o.Priority.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Monitor) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *Monitor) HasPriority() bool {
	return o != nil && o.Priority.IsSet()
}

// SetPriority gets a reference to the given datadog.NullableInt64 and assigns it to the Priority field.
func (o *Monitor) SetPriority(v int64) {
	o.Priority.Set(&v)
}

// SetPriorityNil sets the value for Priority to be an explicit nil.
func (o *Monitor) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil.
func (o *Monitor) UnsetPriority() {
	o.Priority.Unset()
}

// GetQuery returns the Query field value.
func (o *Monitor) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *Monitor) SetQuery(v string) {
	o.Query = v
}

// GetRestrictedRoles returns the RestrictedRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Monitor) GetRestrictedRoles() []string {
	if o == nil || o.RestrictedRoles.Get() == nil {
		var ret []string
		return ret
	}
	return *o.RestrictedRoles.Get()
}

// GetRestrictedRolesOk returns a tuple with the RestrictedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Monitor) GetRestrictedRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RestrictedRoles.Get(), o.RestrictedRoles.IsSet()
}

// HasRestrictedRoles returns a boolean if a field has been set.
func (o *Monitor) HasRestrictedRoles() bool {
	return o != nil && o.RestrictedRoles.IsSet()
}

// SetRestrictedRoles gets a reference to the given datadog.NullableList[string] and assigns it to the RestrictedRoles field.
func (o *Monitor) SetRestrictedRoles(v []string) {
	o.RestrictedRoles.Set(&v)
}

// SetRestrictedRolesNil sets the value for RestrictedRoles to be an explicit nil.
func (o *Monitor) SetRestrictedRolesNil() {
	o.RestrictedRoles.Set(nil)
}

// UnsetRestrictedRoles ensures that no value is present for RestrictedRoles, not even an explicit nil.
func (o *Monitor) UnsetRestrictedRoles() {
	o.RestrictedRoles.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Monitor) GetState() MonitorState {
	if o == nil || o.State == nil {
		var ret MonitorState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetStateOk() (*MonitorState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Monitor) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given MonitorState and assigns it to the State field.
func (o *Monitor) SetState(v MonitorState) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Monitor) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitor) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Monitor) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Monitor) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *Monitor) GetType() MonitorType {
	if o == nil {
		var ret MonitorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetTypeOk() (*MonitorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Monitor) SetType(v MonitorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Monitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Deleted.IsSet() {
		toSerialize["deleted"] = o.Deleted.Get()
	}
	if o.DraftStatus != nil {
		toSerialize["draft_status"] = o.DraftStatus
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MatchingDowntimes != nil {
		toSerialize["matching_downtimes"] = o.MatchingDowntimes
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Multi != nil {
		toSerialize["multi"] = o.Multi
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.OverallState != nil {
		toSerialize["overall_state"] = o.OverallState
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	toSerialize["query"] = o.Query
	if o.RestrictedRoles.IsSet() {
		toSerialize["restricted_roles"] = o.RestrictedRoles.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Monitor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created           *time.Time                   `json:"created,omitempty"`
		Creator           *Creator                     `json:"creator,omitempty"`
		Deleted           datadog.NullableTime         `json:"deleted,omitempty"`
		DraftStatus       *MonitorDraftStatus          `json:"draft_status,omitempty"`
		Id                *int64                       `json:"id,omitempty"`
		MatchingDowntimes []MatchingDowntime           `json:"matching_downtimes,omitempty"`
		Message           *string                      `json:"message,omitempty"`
		Modified          *time.Time                   `json:"modified,omitempty"`
		Multi             *bool                        `json:"multi,omitempty"`
		Name              *string                      `json:"name,omitempty"`
		Options           *MonitorOptions              `json:"options,omitempty"`
		OverallState      *MonitorOverallStates        `json:"overall_state,omitempty"`
		Priority          datadog.NullableInt64        `json:"priority,omitempty"`
		Query             *string                      `json:"query"`
		RestrictedRoles   datadog.NullableList[string] `json:"restricted_roles,omitempty"`
		State             *MonitorState                `json:"state,omitempty"`
		Tags              []string                     `json:"tags,omitempty"`
		Type              *MonitorType                 `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "creator", "deleted", "draft_status", "id", "matching_downtimes", "message", "modified", "multi", "name", "options", "overall_state", "priority", "query", "restricted_roles", "state", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Created = all.Created
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Deleted = all.Deleted
	if all.DraftStatus != nil && !all.DraftStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.DraftStatus = all.DraftStatus
	}
	o.Id = all.Id
	o.MatchingDowntimes = all.MatchingDowntimes
	o.Message = all.Message
	o.Modified = all.Modified
	o.Multi = all.Multi
	o.Name = all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	if all.OverallState != nil && !all.OverallState.IsValid() {
		hasInvalidField = true
	} else {
		o.OverallState = all.OverallState
	}
	o.Priority = all.Priority
	o.Query = *all.Query
	o.RestrictedRoles = all.RestrictedRoles
	if all.State != nil && all.State.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.State = all.State
	o.Tags = all.Tags
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
