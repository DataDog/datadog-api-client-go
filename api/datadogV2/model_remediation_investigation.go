// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationInvestigation A single ECS remediation investigation: a detected issue together with its proposed plan, history, and ECS workload metadata.
type RemediationInvestigation struct {
	// A linked code session (for example, a pull request) for the investigation.
	CodeSession *RemediationCodeSession `json:"code_session,omitempty"`
	// Creation time in epoch milliseconds (64-bit integer encoded as a string).
	CreatedAtMs *string `json:"created_at_ms,omitempty"`
	// The guardrail decision applied to a plan or investigation.
	GuardrailDecision *RemediationGuardrailDecision `json:"guardrail_decision,omitempty"`
	// Ordered list of history events for the investigation.
	History []RemediationHistoryEvent `json:"history,omitempty"`
	// Unique investigation ID.
	Id *string `json:"id,omitempty"`
	// The detected issue type.
	IssueType *string `json:"issue_type,omitempty"`
	// Time of the last action in epoch milliseconds (64-bit integer encoded as a string).
	LastActionAtMs *string `json:"last_action_at_ms,omitempty"`
	// ECS-specific context for the investigation.
	Metadata *RemediationEcsMetadata `json:"metadata,omitempty"`
	// Datadog organization ID that owns the investigation (64-bit integer encoded as a string).
	OrgId *string `json:"org_id,omitempty"`
	// The remediation plan proposed by the ECS remediation agent.
	Plan *RemediationPlan `json:"plan,omitempty"`
	// ARN of the ECS resource the investigation is about.
	ResourceArn *string `json:"resource_arn,omitempty"`
	// Investigation status.
	Status *RemediationInvestigationStatus `json:"status,omitempty"`
	// Last update time in epoch milliseconds (64-bit integer encoded as a string).
	UpdatedAtMs *string `json:"updated_at_ms,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationInvestigation instantiates a new RemediationInvestigation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationInvestigation() *RemediationInvestigation {
	this := RemediationInvestigation{}
	return &this
}

// NewRemediationInvestigationWithDefaults instantiates a new RemediationInvestigation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationInvestigationWithDefaults() *RemediationInvestigation {
	this := RemediationInvestigation{}
	return &this
}

// GetCodeSession returns the CodeSession field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetCodeSession() RemediationCodeSession {
	if o == nil || o.CodeSession == nil {
		var ret RemediationCodeSession
		return ret
	}
	return *o.CodeSession
}

// GetCodeSessionOk returns a tuple with the CodeSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetCodeSessionOk() (*RemediationCodeSession, bool) {
	if o == nil || o.CodeSession == nil {
		return nil, false
	}
	return o.CodeSession, true
}

// HasCodeSession returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasCodeSession() bool {
	return o != nil && o.CodeSession != nil
}

// SetCodeSession gets a reference to the given RemediationCodeSession and assigns it to the CodeSession field.
func (o *RemediationInvestigation) SetCodeSession(v RemediationCodeSession) {
	o.CodeSession = &v
}

// GetCreatedAtMs returns the CreatedAtMs field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetCreatedAtMs() string {
	if o == nil || o.CreatedAtMs == nil {
		var ret string
		return ret
	}
	return *o.CreatedAtMs
}

// GetCreatedAtMsOk returns a tuple with the CreatedAtMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetCreatedAtMsOk() (*string, bool) {
	if o == nil || o.CreatedAtMs == nil {
		return nil, false
	}
	return o.CreatedAtMs, true
}

// HasCreatedAtMs returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasCreatedAtMs() bool {
	return o != nil && o.CreatedAtMs != nil
}

// SetCreatedAtMs gets a reference to the given string and assigns it to the CreatedAtMs field.
func (o *RemediationInvestigation) SetCreatedAtMs(v string) {
	o.CreatedAtMs = &v
}

// GetGuardrailDecision returns the GuardrailDecision field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetGuardrailDecision() RemediationGuardrailDecision {
	if o == nil || o.GuardrailDecision == nil {
		var ret RemediationGuardrailDecision
		return ret
	}
	return *o.GuardrailDecision
}

// GetGuardrailDecisionOk returns a tuple with the GuardrailDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetGuardrailDecisionOk() (*RemediationGuardrailDecision, bool) {
	if o == nil || o.GuardrailDecision == nil {
		return nil, false
	}
	return o.GuardrailDecision, true
}

// HasGuardrailDecision returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasGuardrailDecision() bool {
	return o != nil && o.GuardrailDecision != nil
}

// SetGuardrailDecision gets a reference to the given RemediationGuardrailDecision and assigns it to the GuardrailDecision field.
func (o *RemediationInvestigation) SetGuardrailDecision(v RemediationGuardrailDecision) {
	o.GuardrailDecision = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetHistory() []RemediationHistoryEvent {
	if o == nil || o.History == nil {
		var ret []RemediationHistoryEvent
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetHistoryOk() (*[]RemediationHistoryEvent, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return &o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasHistory() bool {
	return o != nil && o.History != nil
}

// SetHistory gets a reference to the given []RemediationHistoryEvent and assigns it to the History field.
func (o *RemediationInvestigation) SetHistory(v []RemediationHistoryEvent) {
	o.History = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemediationInvestigation) SetId(v string) {
	o.Id = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetIssueType() string {
	if o == nil || o.IssueType == nil {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetIssueTypeOk() (*string, bool) {
	if o == nil || o.IssueType == nil {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasIssueType() bool {
	return o != nil && o.IssueType != nil
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *RemediationInvestigation) SetIssueType(v string) {
	o.IssueType = &v
}

// GetLastActionAtMs returns the LastActionAtMs field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetLastActionAtMs() string {
	if o == nil || o.LastActionAtMs == nil {
		var ret string
		return ret
	}
	return *o.LastActionAtMs
}

// GetLastActionAtMsOk returns a tuple with the LastActionAtMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetLastActionAtMsOk() (*string, bool) {
	if o == nil || o.LastActionAtMs == nil {
		return nil, false
	}
	return o.LastActionAtMs, true
}

// HasLastActionAtMs returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasLastActionAtMs() bool {
	return o != nil && o.LastActionAtMs != nil
}

// SetLastActionAtMs gets a reference to the given string and assigns it to the LastActionAtMs field.
func (o *RemediationInvestigation) SetLastActionAtMs(v string) {
	o.LastActionAtMs = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetMetadata() RemediationEcsMetadata {
	if o == nil || o.Metadata == nil {
		var ret RemediationEcsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetMetadataOk() (*RemediationEcsMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given RemediationEcsMetadata and assigns it to the Metadata field.
func (o *RemediationInvestigation) SetMetadata(v RemediationEcsMetadata) {
	o.Metadata = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *RemediationInvestigation) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetPlan() RemediationPlan {
	if o == nil || o.Plan == nil {
		var ret RemediationPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetPlanOk() (*RemediationPlan, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasPlan() bool {
	return o != nil && o.Plan != nil
}

// SetPlan gets a reference to the given RemediationPlan and assigns it to the Plan field.
func (o *RemediationInvestigation) SetPlan(v RemediationPlan) {
	o.Plan = &v
}

// GetResourceArn returns the ResourceArn field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetResourceArn() string {
	if o == nil || o.ResourceArn == nil {
		var ret string
		return ret
	}
	return *o.ResourceArn
}

// GetResourceArnOk returns a tuple with the ResourceArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetResourceArnOk() (*string, bool) {
	if o == nil || o.ResourceArn == nil {
		return nil, false
	}
	return o.ResourceArn, true
}

// HasResourceArn returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasResourceArn() bool {
	return o != nil && o.ResourceArn != nil
}

// SetResourceArn gets a reference to the given string and assigns it to the ResourceArn field.
func (o *RemediationInvestigation) SetResourceArn(v string) {
	o.ResourceArn = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetStatus() RemediationInvestigationStatus {
	if o == nil || o.Status == nil {
		var ret RemediationInvestigationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetStatusOk() (*RemediationInvestigationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given RemediationInvestigationStatus and assigns it to the Status field.
func (o *RemediationInvestigation) SetStatus(v RemediationInvestigationStatus) {
	o.Status = &v
}

// GetUpdatedAtMs returns the UpdatedAtMs field value if set, zero value otherwise.
func (o *RemediationInvestigation) GetUpdatedAtMs() string {
	if o == nil || o.UpdatedAtMs == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAtMs
}

// GetUpdatedAtMsOk returns a tuple with the UpdatedAtMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationInvestigation) GetUpdatedAtMsOk() (*string, bool) {
	if o == nil || o.UpdatedAtMs == nil {
		return nil, false
	}
	return o.UpdatedAtMs, true
}

// HasUpdatedAtMs returns a boolean if a field has been set.
func (o *RemediationInvestigation) HasUpdatedAtMs() bool {
	return o != nil && o.UpdatedAtMs != nil
}

// SetUpdatedAtMs gets a reference to the given string and assigns it to the UpdatedAtMs field.
func (o *RemediationInvestigation) SetUpdatedAtMs(v string) {
	o.UpdatedAtMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationInvestigation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CodeSession != nil {
		toSerialize["code_session"] = o.CodeSession
	}
	if o.CreatedAtMs != nil {
		toSerialize["created_at_ms"] = o.CreatedAtMs
	}
	if o.GuardrailDecision != nil {
		toSerialize["guardrail_decision"] = o.GuardrailDecision
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IssueType != nil {
		toSerialize["issue_type"] = o.IssueType
	}
	if o.LastActionAtMs != nil {
		toSerialize["last_action_at_ms"] = o.LastActionAtMs
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.ResourceArn != nil {
		toSerialize["resource_arn"] = o.ResourceArn
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAtMs != nil {
		toSerialize["updated_at_ms"] = o.UpdatedAtMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationInvestigation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CodeSession       *RemediationCodeSession         `json:"code_session,omitempty"`
		CreatedAtMs       *string                         `json:"created_at_ms,omitempty"`
		GuardrailDecision *RemediationGuardrailDecision   `json:"guardrail_decision,omitempty"`
		History           []RemediationHistoryEvent       `json:"history,omitempty"`
		Id                *string                         `json:"id,omitempty"`
		IssueType         *string                         `json:"issue_type,omitempty"`
		LastActionAtMs    *string                         `json:"last_action_at_ms,omitempty"`
		Metadata          *RemediationEcsMetadata         `json:"metadata,omitempty"`
		OrgId             *string                         `json:"org_id,omitempty"`
		Plan              *RemediationPlan                `json:"plan,omitempty"`
		ResourceArn       *string                         `json:"resource_arn,omitempty"`
		Status            *RemediationInvestigationStatus `json:"status,omitempty"`
		UpdatedAtMs       *string                         `json:"updated_at_ms,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"code_session", "created_at_ms", "guardrail_decision", "history", "id", "issue_type", "last_action_at_ms", "metadata", "org_id", "plan", "resource_arn", "status", "updated_at_ms"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CodeSession != nil && all.CodeSession.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CodeSession = all.CodeSession
	o.CreatedAtMs = all.CreatedAtMs
	if all.GuardrailDecision != nil && all.GuardrailDecision.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GuardrailDecision = all.GuardrailDecision
	o.History = all.History
	o.Id = all.Id
	o.IssueType = all.IssueType
	o.LastActionAtMs = all.LastActionAtMs
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.OrgId = all.OrgId
	if all.Plan != nil && all.Plan.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Plan = all.Plan
	o.ResourceArn = all.ResourceArn
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.UpdatedAtMs = all.UpdatedAtMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
