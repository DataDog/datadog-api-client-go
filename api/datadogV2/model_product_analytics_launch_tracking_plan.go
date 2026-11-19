// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchTrackingPlan Tracking plan of a launch, listing the events and properties to instrument.
type ProductAnalyticsLaunchTrackingPlan struct {
	// Latest inspection of the repository for the tracking plan instrumentation. A new analysis replaces the stored
	// one.
	// Events and properties that it does not list are not verified.
	CodeAnalysis *ProductAnalyticsLaunchCodeAnalysis `json:"code_analysis,omitempty"`
	// ID of the code generation session that implemented the plan.
	CodegenSessionId *string `json:"codegen_session_id,omitempty"`
	// Time the tracking plan was created.
	CreatedAt time.Time `json:"created_at"`
	// UUID of the user who created the tracking plan.
	CreatedBy string `json:"created_by"`
	// Events to instrument.
	Events []ProductAnalyticsLaunchEvent `json:"events,omitempty"`
	// ID of the tracking plan.
	Id string `json:"id"`
	// ID of the Jira ticket that tracks the instrumentation work.
	JiraTicketId *string `json:"jira_ticket_id,omitempty"`
	// Link to the pull request that implements the plan.
	PrLink *string `json:"pr_link,omitempty"`
	// Event properties of the tracking plan, keyed by property ID.
	Properties map[string]ProductAnalyticsLaunchProperty `json:"properties,omitempty"`
	// Time the tracking plan was last updated.
	UpdatedAt datadog.NullableTime `json:"updated_at,omitempty"`
	// UUID of the user who last updated the tracking plan.
	UpdatedBy datadog.NullableString `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchTrackingPlan instantiates a new ProductAnalyticsLaunchTrackingPlan object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchTrackingPlan(createdAt time.Time, createdBy string, id string) *ProductAnalyticsLaunchTrackingPlan {
	this := ProductAnalyticsLaunchTrackingPlan{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Id = id
	return &this
}

// NewProductAnalyticsLaunchTrackingPlanWithDefaults instantiates a new ProductAnalyticsLaunchTrackingPlan object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchTrackingPlanWithDefaults() *ProductAnalyticsLaunchTrackingPlan {
	this := ProductAnalyticsLaunchTrackingPlan{}
	return &this
}

// GetCodeAnalysis returns the CodeAnalysis field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCodeAnalysis() ProductAnalyticsLaunchCodeAnalysis {
	if o == nil || o.CodeAnalysis == nil {
		var ret ProductAnalyticsLaunchCodeAnalysis
		return ret
	}
	return *o.CodeAnalysis
}

// GetCodeAnalysisOk returns a tuple with the CodeAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCodeAnalysisOk() (*ProductAnalyticsLaunchCodeAnalysis, bool) {
	if o == nil || o.CodeAnalysis == nil {
		return nil, false
	}
	return o.CodeAnalysis, true
}

// HasCodeAnalysis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasCodeAnalysis() bool {
	return o != nil && o.CodeAnalysis != nil
}

// SetCodeAnalysis gets a reference to the given ProductAnalyticsLaunchCodeAnalysis and assigns it to the CodeAnalysis field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetCodeAnalysis(v ProductAnalyticsLaunchCodeAnalysis) {
	o.CodeAnalysis = &v
}

// GetCodegenSessionId returns the CodegenSessionId field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCodegenSessionId() string {
	if o == nil || o.CodegenSessionId == nil {
		var ret string
		return ret
	}
	return *o.CodegenSessionId
}

// GetCodegenSessionIdOk returns a tuple with the CodegenSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCodegenSessionIdOk() (*string, bool) {
	if o == nil || o.CodegenSessionId == nil {
		return nil, false
	}
	return o.CodegenSessionId, true
}

// HasCodegenSessionId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasCodegenSessionId() bool {
	return o != nil && o.CodegenSessionId != nil
}

// SetCodegenSessionId gets a reference to the given string and assigns it to the CodegenSessionId field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetCodegenSessionId(v string) {
	o.CodegenSessionId = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ProductAnalyticsLaunchTrackingPlan) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *ProductAnalyticsLaunchTrackingPlan) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlan) GetEvents() []ProductAnalyticsLaunchEvent {
	if o == nil || o.Events == nil {
		var ret []ProductAnalyticsLaunchEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetEventsOk() (*[]ProductAnalyticsLaunchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []ProductAnalyticsLaunchEvent and assigns it to the Events field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetEvents(v []ProductAnalyticsLaunchEvent) {
	o.Events = v
}

// GetId returns the Id field value.
func (o *ProductAnalyticsLaunchTrackingPlan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ProductAnalyticsLaunchTrackingPlan) SetId(v string) {
	o.Id = v
}

// GetJiraTicketId returns the JiraTicketId field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlan) GetJiraTicketId() string {
	if o == nil || o.JiraTicketId == nil {
		var ret string
		return ret
	}
	return *o.JiraTicketId
}

// GetJiraTicketIdOk returns a tuple with the JiraTicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetJiraTicketIdOk() (*string, bool) {
	if o == nil || o.JiraTicketId == nil {
		return nil, false
	}
	return o.JiraTicketId, true
}

// HasJiraTicketId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasJiraTicketId() bool {
	return o != nil && o.JiraTicketId != nil
}

// SetJiraTicketId gets a reference to the given string and assigns it to the JiraTicketId field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetJiraTicketId(v string) {
	o.JiraTicketId = &v
}

// GetPrLink returns the PrLink field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlan) GetPrLink() string {
	if o == nil || o.PrLink == nil {
		var ret string
		return ret
	}
	return *o.PrLink
}

// GetPrLinkOk returns a tuple with the PrLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetPrLinkOk() (*string, bool) {
	if o == nil || o.PrLink == nil {
		return nil, false
	}
	return o.PrLink, true
}

// HasPrLink returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasPrLink() bool {
	return o != nil && o.PrLink != nil
}

// SetPrLink gets a reference to the given string and assigns it to the PrLink field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetPrLink(v string) {
	o.PrLink = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlan) GetProperties() map[string]ProductAnalyticsLaunchProperty {
	if o == nil || o.Properties == nil {
		var ret map[string]ProductAnalyticsLaunchProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) GetPropertiesOk() (*map[string]ProductAnalyticsLaunchProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given map[string]ProductAnalyticsLaunchProperty and assigns it to the Properties field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetProperties(v map[string]ProductAnalyticsLaunchProperty) {
	o.Properties = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchTrackingPlan) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchTrackingPlan) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given datadog.NullableTime and assigns it to the UpdatedAt field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *ProductAnalyticsLaunchTrackingPlan) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *ProductAnalyticsLaunchTrackingPlan) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchTrackingPlan) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchTrackingPlan) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlan) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy.IsSet()
}

// SetUpdatedBy gets a reference to the given datadog.NullableString and assigns it to the UpdatedBy field.
func (o *ProductAnalyticsLaunchTrackingPlan) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil.
func (o *ProductAnalyticsLaunchTrackingPlan) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil.
func (o *ProductAnalyticsLaunchTrackingPlan) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchTrackingPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CodeAnalysis != nil {
		toSerialize["code_analysis"] = o.CodeAnalysis
	}
	if o.CodegenSessionId != nil {
		toSerialize["codegen_session_id"] = o.CodegenSessionId
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	toSerialize["id"] = o.Id
	if o.JiraTicketId != nil {
		toSerialize["jira_ticket_id"] = o.JiraTicketId
	}
	if o.PrLink != nil {
		toSerialize["pr_link"] = o.PrLink
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchTrackingPlan) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeAnalysis     *ProductAnalyticsLaunchCodeAnalysis       `json:"code_analysis,omitempty"`
		CodegenSessionId *string                                   `json:"codegen_session_id,omitempty"`
		CreatedAt        *time.Time                                `json:"created_at"`
		CreatedBy        *string                                   `json:"created_by"`
		Events           []ProductAnalyticsLaunchEvent             `json:"events,omitempty"`
		Id               *string                                   `json:"id"`
		JiraTicketId     *string                                   `json:"jira_ticket_id,omitempty"`
		PrLink           *string                                   `json:"pr_link,omitempty"`
		Properties       map[string]ProductAnalyticsLaunchProperty `json:"properties,omitempty"`
		UpdatedAt        datadog.NullableTime                      `json:"updated_at,omitempty"`
		UpdatedBy        datadog.NullableString                    `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code_analysis", "codegen_session_id", "created_at", "created_by", "events", "id", "jira_ticket_id", "pr_link", "properties", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CodeAnalysis != nil && all.CodeAnalysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CodeAnalysis = all.CodeAnalysis
	o.CodegenSessionId = all.CodegenSessionId
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Events = all.Events
	o.Id = *all.Id
	o.JiraTicketId = all.JiraTicketId
	o.PrLink = all.PrLink
	o.Properties = all.Properties
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
