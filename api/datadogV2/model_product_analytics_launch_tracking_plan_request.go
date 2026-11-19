// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchTrackingPlanRequest Tracking plan of a launch, listing the events and properties to instrument.
type ProductAnalyticsLaunchTrackingPlanRequest struct {
	// Latest inspection of the repository for the tracking plan instrumentation. A new analysis replaces the stored
	// one.
	// Events and properties that it does not list are not verified.
	CodeAnalysis *ProductAnalyticsLaunchCodeAnalysis `json:"code_analysis,omitempty"`
	// ID of the code generation session that implemented the plan.
	CodegenSessionId *string `json:"codegen_session_id,omitempty"`
	// Events to instrument.
	Events []ProductAnalyticsLaunchEvent `json:"events,omitempty"`
	// ID of the Jira ticket that tracks the instrumentation work.
	JiraTicketId *string `json:"jira_ticket_id,omitempty"`
	// Link to the pull request that implements the plan.
	PrLink *string `json:"pr_link,omitempty"`
	// Event properties of the tracking plan, keyed by property ID.
	Properties map[string]ProductAnalyticsLaunchProperty `json:"properties,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchTrackingPlanRequest instantiates a new ProductAnalyticsLaunchTrackingPlanRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchTrackingPlanRequest() *ProductAnalyticsLaunchTrackingPlanRequest {
	this := ProductAnalyticsLaunchTrackingPlanRequest{}
	return &this
}

// NewProductAnalyticsLaunchTrackingPlanRequestWithDefaults instantiates a new ProductAnalyticsLaunchTrackingPlanRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchTrackingPlanRequestWithDefaults() *ProductAnalyticsLaunchTrackingPlanRequest {
	this := ProductAnalyticsLaunchTrackingPlanRequest{}
	return &this
}

// GetCodeAnalysis returns the CodeAnalysis field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetCodeAnalysis() ProductAnalyticsLaunchCodeAnalysis {
	if o == nil || o.CodeAnalysis == nil {
		var ret ProductAnalyticsLaunchCodeAnalysis
		return ret
	}
	return *o.CodeAnalysis
}

// GetCodeAnalysisOk returns a tuple with the CodeAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetCodeAnalysisOk() (*ProductAnalyticsLaunchCodeAnalysis, bool) {
	if o == nil || o.CodeAnalysis == nil {
		return nil, false
	}
	return o.CodeAnalysis, true
}

// HasCodeAnalysis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) HasCodeAnalysis() bool {
	return o != nil && o.CodeAnalysis != nil
}

// SetCodeAnalysis gets a reference to the given ProductAnalyticsLaunchCodeAnalysis and assigns it to the CodeAnalysis field.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) SetCodeAnalysis(v ProductAnalyticsLaunchCodeAnalysis) {
	o.CodeAnalysis = &v
}

// GetCodegenSessionId returns the CodegenSessionId field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetCodegenSessionId() string {
	if o == nil || o.CodegenSessionId == nil {
		var ret string
		return ret
	}
	return *o.CodegenSessionId
}

// GetCodegenSessionIdOk returns a tuple with the CodegenSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetCodegenSessionIdOk() (*string, bool) {
	if o == nil || o.CodegenSessionId == nil {
		return nil, false
	}
	return o.CodegenSessionId, true
}

// HasCodegenSessionId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) HasCodegenSessionId() bool {
	return o != nil && o.CodegenSessionId != nil
}

// SetCodegenSessionId gets a reference to the given string and assigns it to the CodegenSessionId field.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) SetCodegenSessionId(v string) {
	o.CodegenSessionId = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetEvents() []ProductAnalyticsLaunchEvent {
	if o == nil || o.Events == nil {
		var ret []ProductAnalyticsLaunchEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetEventsOk() (*[]ProductAnalyticsLaunchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []ProductAnalyticsLaunchEvent and assigns it to the Events field.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) SetEvents(v []ProductAnalyticsLaunchEvent) {
	o.Events = v
}

// GetJiraTicketId returns the JiraTicketId field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetJiraTicketId() string {
	if o == nil || o.JiraTicketId == nil {
		var ret string
		return ret
	}
	return *o.JiraTicketId
}

// GetJiraTicketIdOk returns a tuple with the JiraTicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetJiraTicketIdOk() (*string, bool) {
	if o == nil || o.JiraTicketId == nil {
		return nil, false
	}
	return o.JiraTicketId, true
}

// HasJiraTicketId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) HasJiraTicketId() bool {
	return o != nil && o.JiraTicketId != nil
}

// SetJiraTicketId gets a reference to the given string and assigns it to the JiraTicketId field.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) SetJiraTicketId(v string) {
	o.JiraTicketId = &v
}

// GetPrLink returns the PrLink field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetPrLink() string {
	if o == nil || o.PrLink == nil {
		var ret string
		return ret
	}
	return *o.PrLink
}

// GetPrLinkOk returns a tuple with the PrLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetPrLinkOk() (*string, bool) {
	if o == nil || o.PrLink == nil {
		return nil, false
	}
	return o.PrLink, true
}

// HasPrLink returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) HasPrLink() bool {
	return o != nil && o.PrLink != nil
}

// SetPrLink gets a reference to the given string and assigns it to the PrLink field.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) SetPrLink(v string) {
	o.PrLink = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetProperties() map[string]ProductAnalyticsLaunchProperty {
	if o == nil || o.Properties == nil {
		var ret map[string]ProductAnalyticsLaunchProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) GetPropertiesOk() (*map[string]ProductAnalyticsLaunchProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given map[string]ProductAnalyticsLaunchProperty and assigns it to the Properties field.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) SetProperties(v map[string]ProductAnalyticsLaunchProperty) {
	o.Properties = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchTrackingPlanRequest) MarshalJSON() ([]byte, error) {
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
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.JiraTicketId != nil {
		toSerialize["jira_ticket_id"] = o.JiraTicketId
	}
	if o.PrLink != nil {
		toSerialize["pr_link"] = o.PrLink
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchTrackingPlanRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeAnalysis     *ProductAnalyticsLaunchCodeAnalysis       `json:"code_analysis,omitempty"`
		CodegenSessionId *string                                   `json:"codegen_session_id,omitempty"`
		Events           []ProductAnalyticsLaunchEvent             `json:"events,omitempty"`
		JiraTicketId     *string                                   `json:"jira_ticket_id,omitempty"`
		PrLink           *string                                   `json:"pr_link,omitempty"`
		Properties       map[string]ProductAnalyticsLaunchProperty `json:"properties,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code_analysis", "codegen_session_id", "events", "jira_ticket_id", "pr_link", "properties"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CodeAnalysis != nil && all.CodeAnalysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CodeAnalysis = all.CodeAnalysis
	o.CodegenSessionId = all.CodegenSessionId
	o.Events = all.Events
	o.JiraTicketId = all.JiraTicketId
	o.PrLink = all.PrLink
	o.Properties = all.Properties

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
