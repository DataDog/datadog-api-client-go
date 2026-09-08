// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems A static analysis rule within a ruleset, including its definition, metadata, and associated test cases.
type GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems struct {
	// The list of configurable arguments accepted by this rule.
	Arguments []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems `json:"arguments"`
	// The category classifying the type of issue this rule detects (e.g., security, style, performance).
	Category string `json:"category"`
	// A checksum of the rule definition used to detect changes.
	Checksum string `json:"checksum"`
	// The rule implementation code used by the static analysis engine.
	Code string `json:"code"`
	// The date and time when the rule was created.
	CreatedAt time.Time `json:"created_at"`
	// The identifier of the user or system that created the rule.
	CreatedBy string `json:"created_by"`
	// The CVE identifier associated with the vulnerability this rule detects, if applicable.
	Cve *string `json:"cve,omitempty"`
	// The CWE identifier associated with the weakness category this rule detects, if applicable.
	Cwe *string `json:"cwe,omitempty"`
	// A detailed explanation of what the rule detects and why it matters.
	Description string `json:"description"`
	// A URL pointing to additional documentation for this rule.
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	// The code entity type (e.g., function, class, variable) that this rule inspects.
	EntityChecked *string `json:"entity_checked,omitempty"`
	// The unique identifier of the rule, which is the same as its name.
	Id string `json:"id"`
	// Indicates whether the rule is publicly published and available to all users.
	IsPublished bool `json:"is_published"`
	// Indicates whether the rule is in testing mode and not yet promoted to production.
	IsTesting bool `json:"is_testing"`
	// The programming language this rule applies to.
	Language string `json:"language"`
	// The date and time when the rule was last modified.
	LastUpdatedAt time.Time `json:"last_updated_at"`
	// The identifier of the user or system that last updated the rule.
	LastUpdatedBy string `json:"last_updated_by"`
	// The unique name identifying this rule within its ruleset.
	Name string `json:"name"`
	// A regular expression pattern used by the rule for pattern-based detection.
	Regex *string `json:"regex,omitempty"`
	// The severity level of findings produced by this rule (e.g., ERROR, WARNING, NOTICE).
	Severity string `json:"severity"`
	// A brief summary of what the rule detects, suitable for display in listings.
	ShortDescription string `json:"short_description"`
	// Indicates whether an AI-generated fix suggestion should be offered for findings from this rule.
	ShouldUseAiFix bool `json:"should_use_ai_fix"`
	// The list of test cases used to validate the rule's behavior.
	Tests []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems `json:"tests"`
	// The Tree-sitter query expression used by the rule to match code patterns in the AST.
	TreeSitterQuery *string `json:"tree_sitter_query,omitempty"`
	// The rule type indicating the detection mechanism used (e.g., tree_sitter, regex).
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems(arguments []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems, category string, checksum string, code string, createdAt time.Time, createdBy string, description string, id string, isPublished bool, isTesting bool, language string, lastUpdatedAt time.Time, lastUpdatedBy string, name string, severity string, shortDescription string, shouldUseAiFix bool, tests []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems, typeVar string) *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems{}
	this.Arguments = arguments
	this.Category = category
	this.Checksum = checksum
	this.Code = code
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.Id = id
	this.IsPublished = isPublished
	this.IsTesting = isTesting
	this.Language = language
	this.LastUpdatedAt = lastUpdatedAt
	this.LastUpdatedBy = lastUpdatedBy
	this.Name = name
	this.Severity = severity
	this.ShortDescription = shortDescription
	this.ShouldUseAiFix = shouldUseAiFix
	this.Tests = tests
	this.Type = typeVar
	return &this
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsWithDefaults instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsWithDefaults() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems{}
	return &this
}

// GetArguments returns the Arguments field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetArguments() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems {
	if o == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetArgumentsOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// SetArguments sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetArguments(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) {
	o.Arguments = v
}

// GetCategory returns the Category field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCategory(v string) {
	o.Category = v
}

// GetChecksum returns the Checksum field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetChecksum(v string) {
	o.Checksum = v
}

// GetCode returns the Code field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCode(v string) {
	o.Code = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCve() string {
	if o == nil || o.Cve == nil {
		var ret string
		return ret
	}
	return *o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCveOk() (*string, bool) {
	if o == nil || o.Cve == nil {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCve() bool {
	return o != nil && o.Cve != nil
}

// SetCve gets a reference to the given string and assigns it to the Cve field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCve(v string) {
	o.Cve = &v
}

// GetCwe returns the Cwe field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCwe() string {
	if o == nil || o.Cwe == nil {
		var ret string
		return ret
	}
	return *o.Cwe
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetCweOk() (*string, bool) {
	if o == nil || o.Cwe == nil {
		return nil, false
	}
	return o.Cwe, true
}

// HasCwe returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasCwe() bool {
	return o != nil && o.Cwe != nil
}

// SetCwe gets a reference to the given string and assigns it to the Cwe field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetCwe(v string) {
	o.Cwe = &v
}

// GetDescription returns the Description field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetDescription(v string) {
	o.Description = v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasDocumentationUrl() bool {
	return o != nil && o.DocumentationUrl != nil
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetEntityChecked returns the EntityChecked field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetEntityChecked() string {
	if o == nil || o.EntityChecked == nil {
		var ret string
		return ret
	}
	return *o.EntityChecked
}

// GetEntityCheckedOk returns a tuple with the EntityChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetEntityCheckedOk() (*string, bool) {
	if o == nil || o.EntityChecked == nil {
		return nil, false
	}
	return o.EntityChecked, true
}

// HasEntityChecked returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasEntityChecked() bool {
	return o != nil && o.EntityChecked != nil
}

// SetEntityChecked gets a reference to the given string and assigns it to the EntityChecked field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetEntityChecked(v string) {
	o.EntityChecked = &v
}

// GetId returns the Id field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetId(v string) {
	o.Id = v
}

// GetIsPublished returns the IsPublished field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetIsTesting returns the IsTesting field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsTesting() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetIsTestingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTesting, true
}

// SetIsTesting sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetIsTesting(v bool) {
	o.IsTesting = v
}

// GetLanguage returns the Language field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetLanguage(v string) {
	o.Language = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetName returns the Name field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetName(v string) {
	o.Name = v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasRegex() bool {
	return o != nil && o.Regex != nil
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetRegex(v string) {
	o.Regex = &v
}

// GetSeverity returns the Severity field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetSeverity(v string) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetShortDescription(v string) {
	o.ShortDescription = v
}

// GetShouldUseAiFix returns the ShouldUseAiFix field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShouldUseAiFix() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ShouldUseAiFix
}

// GetShouldUseAiFixOk returns a tuple with the ShouldUseAiFix field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetShouldUseAiFixOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldUseAiFix, true
}

// SetShouldUseAiFix sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetShouldUseAiFix(v bool) {
	o.ShouldUseAiFix = v
}

// GetTests returns the Tests field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTests() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems {
	if o == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTestsOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tests, true
}

// SetTests sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetTests(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems) {
	o.Tests = v
}

// GetTreeSitterQuery returns the TreeSitterQuery field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTreeSitterQuery() string {
	if o == nil || o.TreeSitterQuery == nil {
		var ret string
		return ret
	}
	return *o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil || o.TreeSitterQuery == nil {
		return nil, false
	}
	return o.TreeSitterQuery, true
}

// HasTreeSitterQuery returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) HasTreeSitterQuery() bool {
	return o != nil && o.TreeSitterQuery != nil
}

// SetTreeSitterQuery gets a reference to the given string and assigns it to the TreeSitterQuery field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = &v
}

// GetType returns the Type field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arguments"] = o.Arguments
	toSerialize["category"] = o.Category
	toSerialize["checksum"] = o.Checksum
	toSerialize["code"] = o.Code
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	if o.Cve != nil {
		toSerialize["cve"] = o.Cve
	}
	if o.Cwe != nil {
		toSerialize["cwe"] = o.Cwe
	}
	toSerialize["description"] = o.Description
	if o.DocumentationUrl != nil {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if o.EntityChecked != nil {
		toSerialize["entity_checked"] = o.EntityChecked
	}
	toSerialize["id"] = o.Id
	toSerialize["is_published"] = o.IsPublished
	toSerialize["is_testing"] = o.IsTesting
	toSerialize["language"] = o.Language
	if o.LastUpdatedAt.Nanosecond() == 0 {
		toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["last_updated_by"] = o.LastUpdatedBy
	toSerialize["name"] = o.Name
	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}
	toSerialize["severity"] = o.Severity
	toSerialize["short_description"] = o.ShortDescription
	toSerialize["should_use_ai_fix"] = o.ShouldUseAiFix
	toSerialize["tests"] = o.Tests
	if o.TreeSitterQuery != nil {
		toSerialize["tree_sitter_query"] = o.TreeSitterQuery
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments        *[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems `json:"arguments"`
		Category         *string                                                                           `json:"category"`
		Checksum         *string                                                                           `json:"checksum"`
		Code             *string                                                                           `json:"code"`
		CreatedAt        *time.Time                                                                        `json:"created_at"`
		CreatedBy        *string                                                                           `json:"created_by"`
		Cve              *string                                                                           `json:"cve,omitempty"`
		Cwe              *string                                                                           `json:"cwe,omitempty"`
		Description      *string                                                                           `json:"description"`
		DocumentationUrl *string                                                                           `json:"documentation_url,omitempty"`
		EntityChecked    *string                                                                           `json:"entity_checked,omitempty"`
		Id               *string                                                                           `json:"id"`
		IsPublished      *bool                                                                             `json:"is_published"`
		IsTesting        *bool                                                                             `json:"is_testing"`
		Language         *string                                                                           `json:"language"`
		LastUpdatedAt    *time.Time                                                                        `json:"last_updated_at"`
		LastUpdatedBy    *string                                                                           `json:"last_updated_by"`
		Name             *string                                                                           `json:"name"`
		Regex            *string                                                                           `json:"regex,omitempty"`
		Severity         *string                                                                           `json:"severity"`
		ShortDescription *string                                                                           `json:"short_description"`
		ShouldUseAiFix   *bool                                                                             `json:"should_use_ai_fix"`
		Tests            *[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsTestsItems     `json:"tests"`
		TreeSitterQuery  *string                                                                           `json:"tree_sitter_query,omitempty"`
		Type             *string                                                                           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Arguments == nil {
		return fmt.Errorf("required field arguments missing")
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
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.IsPublished == nil {
		return fmt.Errorf("required field is_published missing")
	}
	if all.IsTesting == nil {
		return fmt.Errorf("required field is_testing missing")
	}
	if all.Language == nil {
		return fmt.Errorf("required field language missing")
	}
	if all.LastUpdatedAt == nil {
		return fmt.Errorf("required field last_updated_at missing")
	}
	if all.LastUpdatedBy == nil {
		return fmt.Errorf("required field last_updated_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	if all.ShouldUseAiFix == nil {
		return fmt.Errorf("required field should_use_ai_fix missing")
	}
	if all.Tests == nil {
		return fmt.Errorf("required field tests missing")
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
	o.Arguments = *all.Arguments
	o.Category = *all.Category
	o.Checksum = *all.Checksum
	o.Code = *all.Code
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Cve = all.Cve
	o.Cwe = all.Cwe
	o.Description = *all.Description
	o.DocumentationUrl = all.DocumentationUrl
	o.EntityChecked = all.EntityChecked
	o.Id = *all.Id
	o.IsPublished = *all.IsPublished
	o.IsTesting = *all.IsTesting
	o.Language = *all.Language
	o.LastUpdatedAt = *all.LastUpdatedAt
	o.LastUpdatedBy = *all.LastUpdatedBy
	o.Name = *all.Name
	o.Regex = all.Regex
	o.Severity = *all.Severity
	o.ShortDescription = *all.ShortDescription
	o.ShouldUseAiFix = *all.ShouldUseAiFix
	o.Tests = *all.Tests
	o.TreeSitterQuery = all.TreeSitterQuery
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
