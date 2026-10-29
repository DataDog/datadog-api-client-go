// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageDataRelationships The relationships of an On-Call Page.
type OnCallPageDataRelationships struct {
	// Reference to the entity that performed an action on the Page.
	Author *OnCallPageAuthorRelationship `json:"author,omitempty"`
	// Reference to the entity that performed an action on the Page.
	ClosedBy *OnCallPageAuthorRelationship `json:"closed_by,omitempty"`
	// The Jira issue linked to the Page.
	JiraIssue *OnCallPageJiraIssueRelationship `json:"jira_issue,omitempty"`
	// Links between the Page and other Datadog resources, such as incidents. Always empty when listing Pages.
	Links *OnCallPageLinksRelationship `json:"links,omitempty"`
	// The users responding to the Page.
	Responders *OnCallPageRespondersRelationship `json:"responders,omitempty"`
	// The next action scheduled on the Page, such as an escalation.
	ScheduledAction *OnCallPageScheduledActionRelationship `json:"scheduled_action,omitempty"`
	// The ServiceNow incident linked to the Page.
	ServiceNowIncident *OnCallPageServiceNowIncidentRelationship `json:"service_now_incident,omitempty"`
	// The team, user, or cross-organization handle that was paged.
	Target *OnCallPageTargetRelationship `json:"target,omitempty"`
	// The team routing rule that matched the Page.
	TeamRoutingRule *OnCallPageTeamRoutingRuleRelationship `json:"team_routing_rule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageDataRelationships instantiates a new OnCallPageDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageDataRelationships() *OnCallPageDataRelationships {
	this := OnCallPageDataRelationships{}
	return &this
}

// NewOnCallPageDataRelationshipsWithDefaults instantiates a new OnCallPageDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageDataRelationshipsWithDefaults() *OnCallPageDataRelationships {
	this := OnCallPageDataRelationships{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetAuthor() OnCallPageAuthorRelationship {
	if o == nil || o.Author == nil {
		var ret OnCallPageAuthorRelationship
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetAuthorOk() (*OnCallPageAuthorRelationship, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given OnCallPageAuthorRelationship and assigns it to the Author field.
func (o *OnCallPageDataRelationships) SetAuthor(v OnCallPageAuthorRelationship) {
	o.Author = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetClosedBy() OnCallPageAuthorRelationship {
	if o == nil || o.ClosedBy == nil {
		var ret OnCallPageAuthorRelationship
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetClosedByOk() (*OnCallPageAuthorRelationship, bool) {
	if o == nil || o.ClosedBy == nil {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasClosedBy() bool {
	return o != nil && o.ClosedBy != nil
}

// SetClosedBy gets a reference to the given OnCallPageAuthorRelationship and assigns it to the ClosedBy field.
func (o *OnCallPageDataRelationships) SetClosedBy(v OnCallPageAuthorRelationship) {
	o.ClosedBy = &v
}

// GetJiraIssue returns the JiraIssue field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetJiraIssue() OnCallPageJiraIssueRelationship {
	if o == nil || o.JiraIssue == nil {
		var ret OnCallPageJiraIssueRelationship
		return ret
	}
	return *o.JiraIssue
}

// GetJiraIssueOk returns a tuple with the JiraIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetJiraIssueOk() (*OnCallPageJiraIssueRelationship, bool) {
	if o == nil || o.JiraIssue == nil {
		return nil, false
	}
	return o.JiraIssue, true
}

// HasJiraIssue returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasJiraIssue() bool {
	return o != nil && o.JiraIssue != nil
}

// SetJiraIssue gets a reference to the given OnCallPageJiraIssueRelationship and assigns it to the JiraIssue field.
func (o *OnCallPageDataRelationships) SetJiraIssue(v OnCallPageJiraIssueRelationship) {
	o.JiraIssue = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetLinks() OnCallPageLinksRelationship {
	if o == nil || o.Links == nil {
		var ret OnCallPageLinksRelationship
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetLinksOk() (*OnCallPageLinksRelationship, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given OnCallPageLinksRelationship and assigns it to the Links field.
func (o *OnCallPageDataRelationships) SetLinks(v OnCallPageLinksRelationship) {
	o.Links = &v
}

// GetResponders returns the Responders field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetResponders() OnCallPageRespondersRelationship {
	if o == nil || o.Responders == nil {
		var ret OnCallPageRespondersRelationship
		return ret
	}
	return *o.Responders
}

// GetRespondersOk returns a tuple with the Responders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetRespondersOk() (*OnCallPageRespondersRelationship, bool) {
	if o == nil || o.Responders == nil {
		return nil, false
	}
	return o.Responders, true
}

// HasResponders returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasResponders() bool {
	return o != nil && o.Responders != nil
}

// SetResponders gets a reference to the given OnCallPageRespondersRelationship and assigns it to the Responders field.
func (o *OnCallPageDataRelationships) SetResponders(v OnCallPageRespondersRelationship) {
	o.Responders = &v
}

// GetScheduledAction returns the ScheduledAction field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetScheduledAction() OnCallPageScheduledActionRelationship {
	if o == nil || o.ScheduledAction == nil {
		var ret OnCallPageScheduledActionRelationship
		return ret
	}
	return *o.ScheduledAction
}

// GetScheduledActionOk returns a tuple with the ScheduledAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetScheduledActionOk() (*OnCallPageScheduledActionRelationship, bool) {
	if o == nil || o.ScheduledAction == nil {
		return nil, false
	}
	return o.ScheduledAction, true
}

// HasScheduledAction returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasScheduledAction() bool {
	return o != nil && o.ScheduledAction != nil
}

// SetScheduledAction gets a reference to the given OnCallPageScheduledActionRelationship and assigns it to the ScheduledAction field.
func (o *OnCallPageDataRelationships) SetScheduledAction(v OnCallPageScheduledActionRelationship) {
	o.ScheduledAction = &v
}

// GetServiceNowIncident returns the ServiceNowIncident field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetServiceNowIncident() OnCallPageServiceNowIncidentRelationship {
	if o == nil || o.ServiceNowIncident == nil {
		var ret OnCallPageServiceNowIncidentRelationship
		return ret
	}
	return *o.ServiceNowIncident
}

// GetServiceNowIncidentOk returns a tuple with the ServiceNowIncident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetServiceNowIncidentOk() (*OnCallPageServiceNowIncidentRelationship, bool) {
	if o == nil || o.ServiceNowIncident == nil {
		return nil, false
	}
	return o.ServiceNowIncident, true
}

// HasServiceNowIncident returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasServiceNowIncident() bool {
	return o != nil && o.ServiceNowIncident != nil
}

// SetServiceNowIncident gets a reference to the given OnCallPageServiceNowIncidentRelationship and assigns it to the ServiceNowIncident field.
func (o *OnCallPageDataRelationships) SetServiceNowIncident(v OnCallPageServiceNowIncidentRelationship) {
	o.ServiceNowIncident = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetTarget() OnCallPageTargetRelationship {
	if o == nil || o.Target == nil {
		var ret OnCallPageTargetRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetTargetOk() (*OnCallPageTargetRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given OnCallPageTargetRelationship and assigns it to the Target field.
func (o *OnCallPageDataRelationships) SetTarget(v OnCallPageTargetRelationship) {
	o.Target = &v
}

// GetTeamRoutingRule returns the TeamRoutingRule field value if set, zero value otherwise.
func (o *OnCallPageDataRelationships) GetTeamRoutingRule() OnCallPageTeamRoutingRuleRelationship {
	if o == nil || o.TeamRoutingRule == nil {
		var ret OnCallPageTeamRoutingRuleRelationship
		return ret
	}
	return *o.TeamRoutingRule
}

// GetTeamRoutingRuleOk returns a tuple with the TeamRoutingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallPageDataRelationships) GetTeamRoutingRuleOk() (*OnCallPageTeamRoutingRuleRelationship, bool) {
	if o == nil || o.TeamRoutingRule == nil {
		return nil, false
	}
	return o.TeamRoutingRule, true
}

// HasTeamRoutingRule returns a boolean if a field has been set.
func (o *OnCallPageDataRelationships) HasTeamRoutingRule() bool {
	return o != nil && o.TeamRoutingRule != nil
}

// SetTeamRoutingRule gets a reference to the given OnCallPageTeamRoutingRuleRelationship and assigns it to the TeamRoutingRule field.
func (o *OnCallPageDataRelationships) SetTeamRoutingRule(v OnCallPageTeamRoutingRuleRelationship) {
	o.TeamRoutingRule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.ClosedBy != nil {
		toSerialize["closed_by"] = o.ClosedBy
	}
	if o.JiraIssue != nil {
		toSerialize["jira_issue"] = o.JiraIssue
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Responders != nil {
		toSerialize["responders"] = o.Responders
	}
	if o.ScheduledAction != nil {
		toSerialize["scheduled_action"] = o.ScheduledAction
	}
	if o.ServiceNowIncident != nil {
		toSerialize["service_now_incident"] = o.ServiceNowIncident
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.TeamRoutingRule != nil {
		toSerialize["team_routing_rule"] = o.TeamRoutingRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author             *OnCallPageAuthorRelationship             `json:"author,omitempty"`
		ClosedBy           *OnCallPageAuthorRelationship             `json:"closed_by,omitempty"`
		JiraIssue          *OnCallPageJiraIssueRelationship          `json:"jira_issue,omitempty"`
		Links              *OnCallPageLinksRelationship              `json:"links,omitempty"`
		Responders         *OnCallPageRespondersRelationship         `json:"responders,omitempty"`
		ScheduledAction    *OnCallPageScheduledActionRelationship    `json:"scheduled_action,omitempty"`
		ServiceNowIncident *OnCallPageServiceNowIncidentRelationship `json:"service_now_incident,omitempty"`
		Target             *OnCallPageTargetRelationship             `json:"target,omitempty"`
		TeamRoutingRule    *OnCallPageTeamRoutingRuleRelationship    `json:"team_routing_rule,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "closed_by", "jira_issue", "links", "responders", "scheduled_action", "service_now_incident", "target", "team_routing_rule"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	if all.ClosedBy != nil && all.ClosedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ClosedBy = all.ClosedBy
	if all.JiraIssue != nil && all.JiraIssue.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JiraIssue = all.JiraIssue
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
	if all.Responders != nil && all.Responders.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Responders = all.Responders
	if all.ScheduledAction != nil && all.ScheduledAction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScheduledAction = all.ScheduledAction
	if all.ServiceNowIncident != nil && all.ServiceNowIncident.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServiceNowIncident = all.ServiceNowIncident
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	if all.TeamRoutingRule != nil && all.TeamRoutingRule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TeamRoutingRule = all.TeamRoutingRule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
