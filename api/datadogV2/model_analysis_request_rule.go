// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnalysisRequestRule A static analysis rule to apply during code analysis. Clients forward complete rule
// objects returned by the rulesets endpoints, so every member of that resource is
// declared here; only `id`, `category`, `checksum`, `language`, `severity`,
// `tree_sitter_query`, `entity_checked`, `regex`, `type` and `code` are read by this
// operation and the rest are ignored. The schema stays open so that any member beyond
// the forwarded rule resource is reported as a promotion candidate rather than
// rejected; it can be closed once that telemetry confirms none remain.
type AnalysisRequestRule struct {
	// The configurable arguments accepted by the rule. Forwarded from the rulesets endpoints; ignored by this operation.
	Arguments []AnalysisRequestRuleArgument `json:"arguments,omitempty"`
	// The category of the rule (for example, `BEST_PRACTICES`, `SECURITY`).
	Category string `json:"category"`
	// A checksum of the rule definition.
	Checksum string `json:"checksum"`
	// The base64-encoded rule implementation code.
	Code string `json:"code"`
	// The date and time when the rule was created. Server-assigned by the rulesets endpoints; ignored by this operation.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The identifier of the user or system that created the rule. Server-assigned by the rulesets endpoints; ignored by this operation.
	CreatedBy *string `json:"created_by,omitempty"`
	// The CVE identifier associated with the rule. Forwarded from the rulesets endpoints; ignored by this operation.
	Cve *string `json:"cve,omitempty"`
	// The CWE identifier associated with the rule. Forwarded from the rulesets endpoints; ignored by this operation.
	Cwe *string `json:"cwe,omitempty"`
	// A detailed explanation of what the rule detects. Forwarded from the rulesets endpoints; ignored by this operation.
	Description *string `json:"description,omitempty"`
	// A URL pointing to the rule documentation. Forwarded from the rulesets endpoints; ignored by this operation.
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	// The code entity type checked by the rule, applicable when rule type is `AST_CHECK`.
	EntityChecked datadog.NullableString `json:"entity_checked,omitempty"`
	// The unique identifier of the rule.
	Id string `json:"id"`
	// Whether the rule is published. Forwarded from the rulesets endpoints; ignored by this operation.
	IsPublished *bool `json:"is_published,omitempty"`
	// Whether the rule is in testing mode. Forwarded from the rulesets endpoints; ignored by this operation.
	IsTesting *bool `json:"is_testing,omitempty"`
	// The programming language this rule targets.
	Language string `json:"language"`
	// The date and time when the rule was last modified. Server-assigned by the rulesets endpoints; ignored by this operation.
	LastUpdatedAt *time.Time `json:"last_updated_at,omitempty"`
	// The identifier of the user or system that last updated the rule. Server-assigned by the rulesets endpoints; ignored by this operation.
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// The name of the rule. Forwarded from the rulesets endpoints; ignored by this operation.
	Name *string `json:"name,omitempty"`
	// A base64-encoded regex pattern used by the rule, applicable when rule type is `REGEX`.
	Regex datadog.NullableString `json:"regex,omitempty"`
	// The severity of findings from this rule (for example, `ERROR`, `WARNING`).
	Severity string `json:"severity"`
	// A brief summary of what the rule detects. Forwarded from the rulesets endpoints; ignored by this operation.
	ShortDescription *string `json:"short_description,omitempty"`
	// Whether an AI-generated fix should be offered. Forwarded from the rulesets endpoints; ignored by this operation.
	ShouldUseAiFix *bool `json:"should_use_ai_fix,omitempty"`
	// The test cases associated with the rule. Forwarded from the rulesets endpoints; ignored by this operation.
	Tests []AnalysisRequestRuleTest `json:"tests,omitempty"`
	// The base64-encoded tree-sitter query used by the rule.
	TreeSitterQuery string `json:"tree_sitter_query"`
	// The rule type indicating the detection mechanism (for example, `TREE_SITTER_QUERY`).
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisRequestRule instantiates a new AnalysisRequestRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisRequestRule(category string, checksum string, code string, id string, language string, severity string, treeSitterQuery string, typeVar string) *AnalysisRequestRule {
	this := AnalysisRequestRule{}
	this.Category = category
	this.Checksum = checksum
	this.Code = code
	this.Id = id
	this.Language = language
	this.Severity = severity
	this.TreeSitterQuery = treeSitterQuery
	this.Type = typeVar
	return &this
}

// NewAnalysisRequestRuleWithDefaults instantiates a new AnalysisRequestRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisRequestRuleWithDefaults() *AnalysisRequestRule {
	this := AnalysisRequestRule{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetArguments() []AnalysisRequestRuleArgument {
	if o == nil || o.Arguments == nil {
		var ret []AnalysisRequestRuleArgument
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetArgumentsOk() (*[]AnalysisRequestRuleArgument, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasArguments() bool {
	return o != nil && o.Arguments != nil
}

// SetArguments gets a reference to the given []AnalysisRequestRuleArgument and assigns it to the Arguments field.
func (o *AnalysisRequestRule) SetArguments(v []AnalysisRequestRuleArgument) {
	o.Arguments = v
}

// GetCategory returns the Category field value.
func (o *AnalysisRequestRule) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *AnalysisRequestRule) SetCategory(v string) {
	o.Category = v
}

// GetChecksum returns the Checksum field value.
func (o *AnalysisRequestRule) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *AnalysisRequestRule) SetChecksum(v string) {
	o.Checksum = v
}

// GetCode returns the Code field value.
func (o *AnalysisRequestRule) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *AnalysisRequestRule) SetCode(v string) {
	o.Code = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AnalysisRequestRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AnalysisRequestRule) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetCve() string {
	if o == nil || o.Cve == nil {
		var ret string
		return ret
	}
	return *o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCveOk() (*string, bool) {
	if o == nil || o.Cve == nil {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasCve() bool {
	return o != nil && o.Cve != nil
}

// SetCve gets a reference to the given string and assigns it to the Cve field.
func (o *AnalysisRequestRule) SetCve(v string) {
	o.Cve = &v
}

// GetCwe returns the Cwe field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetCwe() string {
	if o == nil || o.Cwe == nil {
		var ret string
		return ret
	}
	return *o.Cwe
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetCweOk() (*string, bool) {
	if o == nil || o.Cwe == nil {
		return nil, false
	}
	return o.Cwe, true
}

// HasCwe returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasCwe() bool {
	return o != nil && o.Cwe != nil
}

// SetCwe gets a reference to the given string and assigns it to the Cwe field.
func (o *AnalysisRequestRule) SetCwe(v string) {
	o.Cwe = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AnalysisRequestRule) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasDocumentationUrl() bool {
	return o != nil && o.DocumentationUrl != nil
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *AnalysisRequestRule) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetEntityChecked returns the EntityChecked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisRequestRule) GetEntityChecked() string {
	if o == nil || o.EntityChecked.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityChecked.Get()
}

// GetEntityCheckedOk returns a tuple with the EntityChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnalysisRequestRule) GetEntityCheckedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityChecked.Get(), o.EntityChecked.IsSet()
}

// HasEntityChecked returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasEntityChecked() bool {
	return o != nil && o.EntityChecked.IsSet()
}

// SetEntityChecked gets a reference to the given datadog.NullableString and assigns it to the EntityChecked field.
func (o *AnalysisRequestRule) SetEntityChecked(v string) {
	o.EntityChecked.Set(&v)
}

// SetEntityCheckedNil sets the value for EntityChecked to be an explicit nil.
func (o *AnalysisRequestRule) SetEntityCheckedNil() {
	o.EntityChecked.Set(nil)
}

// UnsetEntityChecked ensures that no value is present for EntityChecked, not even an explicit nil.
func (o *AnalysisRequestRule) UnsetEntityChecked() {
	o.EntityChecked.Unset()
}

// GetId returns the Id field value.
func (o *AnalysisRequestRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AnalysisRequestRule) SetId(v string) {
	o.Id = v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetIsPublished() bool {
	if o == nil || o.IsPublished == nil {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetIsPublishedOk() (*bool, bool) {
	if o == nil || o.IsPublished == nil {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasIsPublished() bool {
	return o != nil && o.IsPublished != nil
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *AnalysisRequestRule) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetIsTesting returns the IsTesting field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetIsTesting() bool {
	if o == nil || o.IsTesting == nil {
		var ret bool
		return ret
	}
	return *o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetIsTestingOk() (*bool, bool) {
	if o == nil || o.IsTesting == nil {
		return nil, false
	}
	return o.IsTesting, true
}

// HasIsTesting returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasIsTesting() bool {
	return o != nil && o.IsTesting != nil
}

// SetIsTesting gets a reference to the given bool and assigns it to the IsTesting field.
func (o *AnalysisRequestRule) SetIsTesting(v bool) {
	o.IsTesting = &v
}

// GetLanguage returns the Language field value.
func (o *AnalysisRequestRule) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *AnalysisRequestRule) SetLanguage(v string) {
	o.Language = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetLastUpdatedAt() time.Time {
	if o == nil || o.LastUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUpdatedAt == nil {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasLastUpdatedAt() bool {
	return o != nil && o.LastUpdatedAt != nil
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *AnalysisRequestRule) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasLastUpdatedBy() bool {
	return o != nil && o.LastUpdatedBy != nil
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *AnalysisRequestRule) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnalysisRequestRule) SetName(v string) {
	o.Name = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnalysisRequestRule) GetRegex() string {
	if o == nil || o.Regex.Get() == nil {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnalysisRequestRule) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasRegex() bool {
	return o != nil && o.Regex.IsSet()
}

// SetRegex gets a reference to the given datadog.NullableString and assigns it to the Regex field.
func (o *AnalysisRequestRule) SetRegex(v string) {
	o.Regex.Set(&v)
}

// SetRegexNil sets the value for Regex to be an explicit nil.
func (o *AnalysisRequestRule) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil.
func (o *AnalysisRequestRule) UnsetRegex() {
	o.Regex.Unset()
}

// GetSeverity returns the Severity field value.
func (o *AnalysisRequestRule) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *AnalysisRequestRule) SetSeverity(v string) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasShortDescription() bool {
	return o != nil && o.ShortDescription != nil
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *AnalysisRequestRule) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetShouldUseAiFix returns the ShouldUseAiFix field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetShouldUseAiFix() bool {
	if o == nil || o.ShouldUseAiFix == nil {
		var ret bool
		return ret
	}
	return *o.ShouldUseAiFix
}

// GetShouldUseAiFixOk returns a tuple with the ShouldUseAiFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetShouldUseAiFixOk() (*bool, bool) {
	if o == nil || o.ShouldUseAiFix == nil {
		return nil, false
	}
	return o.ShouldUseAiFix, true
}

// HasShouldUseAiFix returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasShouldUseAiFix() bool {
	return o != nil && o.ShouldUseAiFix != nil
}

// SetShouldUseAiFix gets a reference to the given bool and assigns it to the ShouldUseAiFix field.
func (o *AnalysisRequestRule) SetShouldUseAiFix(v bool) {
	o.ShouldUseAiFix = &v
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *AnalysisRequestRule) GetTests() []AnalysisRequestRuleTest {
	if o == nil || o.Tests == nil {
		var ret []AnalysisRequestRuleTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetTestsOk() (*[]AnalysisRequestRuleTest, bool) {
	if o == nil || o.Tests == nil {
		return nil, false
	}
	return &o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *AnalysisRequestRule) HasTests() bool {
	return o != nil && o.Tests != nil
}

// SetTests gets a reference to the given []AnalysisRequestRuleTest and assigns it to the Tests field.
func (o *AnalysisRequestRule) SetTests(v []AnalysisRequestRuleTest) {
	o.Tests = v
}

// GetTreeSitterQuery returns the TreeSitterQuery field value.
func (o *AnalysisRequestRule) GetTreeSitterQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TreeSitterQuery, true
}

// SetTreeSitterQuery sets field value.
func (o *AnalysisRequestRule) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = v
}

// GetType returns the Type field value.
func (o *AnalysisRequestRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnalysisRequestRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AnalysisRequestRule) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisRequestRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	toSerialize["category"] = o.Category
	toSerialize["checksum"] = o.Checksum
	toSerialize["code"] = o.Code
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Cve != nil {
		toSerialize["cve"] = o.Cve
	}
	if o.Cwe != nil {
		toSerialize["cwe"] = o.Cwe
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentationUrl != nil {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if o.EntityChecked.IsSet() {
		toSerialize["entity_checked"] = o.EntityChecked.Get()
	}
	toSerialize["id"] = o.Id
	if o.IsPublished != nil {
		toSerialize["is_published"] = o.IsPublished
	}
	if o.IsTesting != nil {
		toSerialize["is_testing"] = o.IsTesting
	}
	toSerialize["language"] = o.Language
	if o.LastUpdatedAt != nil {
		if o.LastUpdatedAt.Nanosecond() == 0 {
			toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	toSerialize["severity"] = o.Severity
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}
	if o.ShouldUseAiFix != nil {
		toSerialize["should_use_ai_fix"] = o.ShouldUseAiFix
	}
	if o.Tests != nil {
		toSerialize["tests"] = o.Tests
	}
	toSerialize["tree_sitter_query"] = o.TreeSitterQuery
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisRequestRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments        []AnalysisRequestRuleArgument `json:"arguments,omitempty"`
		Category         *string                       `json:"category"`
		Checksum         *string                       `json:"checksum"`
		Code             *string                       `json:"code"`
		CreatedAt        *time.Time                    `json:"created_at,omitempty"`
		CreatedBy        *string                       `json:"created_by,omitempty"`
		Cve              *string                       `json:"cve,omitempty"`
		Cwe              *string                       `json:"cwe,omitempty"`
		Description      *string                       `json:"description,omitempty"`
		DocumentationUrl *string                       `json:"documentation_url,omitempty"`
		EntityChecked    datadog.NullableString        `json:"entity_checked,omitempty"`
		Id               *string                       `json:"id"`
		IsPublished      *bool                         `json:"is_published,omitempty"`
		IsTesting        *bool                         `json:"is_testing,omitempty"`
		Language         *string                       `json:"language"`
		LastUpdatedAt    *time.Time                    `json:"last_updated_at,omitempty"`
		LastUpdatedBy    *string                       `json:"last_updated_by,omitempty"`
		Name             *string                       `json:"name,omitempty"`
		Regex            datadog.NullableString        `json:"regex,omitempty"`
		Severity         *string                       `json:"severity"`
		ShortDescription *string                       `json:"short_description,omitempty"`
		ShouldUseAiFix   *bool                         `json:"should_use_ai_fix,omitempty"`
		Tests            []AnalysisRequestRuleTest     `json:"tests,omitempty"`
		TreeSitterQuery  *string                       `json:"tree_sitter_query"`
		Type             *string                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Checksum == nil {
		return fmt.Errorf("required field checksum missing")
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Language == nil {
		return fmt.Errorf("required field language missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.TreeSitterQuery == nil {
		return fmt.Errorf("required field tree_sitter_query missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "category", "checksum", "code", "created_at", "created_by", "cve", "cwe", "description", "documentation_url", "entity_checked", "id", "is_published", "is_testing", "language", "last_updated_at", "last_updated_by", "name", "regex", "severity", "short_description", "should_use_ai_fix", "tests", "tree_sitter_query", "type"})
	} else {
		return err
	}
	o.Arguments = all.Arguments
	o.Category = *all.Category
	o.Checksum = *all.Checksum
	o.Code = *all.Code
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Cve = all.Cve
	o.Cwe = all.Cwe
	o.Description = all.Description
	o.DocumentationUrl = all.DocumentationUrl
	o.EntityChecked = all.EntityChecked
	o.Id = *all.Id
	o.IsPublished = all.IsPublished
	o.IsTesting = all.IsTesting
	o.Language = *all.Language
	o.LastUpdatedAt = all.LastUpdatedAt
	o.LastUpdatedBy = all.LastUpdatedBy
	o.Name = all.Name
	o.Regex = all.Regex
	o.Severity = *all.Severity
	o.ShortDescription = all.ShortDescription
	o.ShouldUseAiFix = all.ShouldUseAiFix
	o.Tests = all.Tests
	o.TreeSitterQuery = *all.TreeSitterQuery
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
