// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleRevisionInput A revision of a custom static analysis rule as embedded in a rule supplied by a create
// or update request. Nested revisions are sent flat, without a `data`/`type`/`attributes`
// envelope. `id`, `version_id`, `checksum`, `created_at` and `created_by` are server-assigned
// and read-only; they are declared so that a ruleset previously read back can be supplied
// unchanged.
type CustomRuleRevisionInput struct {
	// Rule arguments
	Arguments []Argument `json:"arguments,omitempty"`
	// Rule category
	Category *CustomRuleRevisionAttributesCategory `json:"category,omitempty"`
	// Code checksum
	Checksum *string `json:"checksum,omitempty"`
	// Rule code
	Code *string `json:"code,omitempty"`
	// Creation timestamp
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Creator identifier
	CreatedBy *string `json:"created_by,omitempty"`
	// Revision creation message
	CreationMessage *string `json:"creation_message,omitempty"`
	// Associated CVE
	Cve datadog.NullableString `json:"cve,omitempty"`
	// Associated CWE
	Cwe datadog.NullableString `json:"cwe,omitempty"`
	// Base64-encoded full description
	Description *string `json:"description,omitempty"`
	// Documentation URL
	DocumentationUrl datadog.NullableString `json:"documentation_url,omitempty"`
	// Revision identifier
	Id *string `json:"id,omitempty"`
	// Whether the revision should be published
	IsPublished *bool `json:"is_published,omitempty"`
	// Whether this is a testing revision
	IsTesting *bool `json:"is_testing,omitempty"`
	// Programming language
	Language *Language `json:"language,omitempty"`
	// Rule severity
	Severity *CustomRuleRevisionAttributesSeverity `json:"severity,omitempty"`
	// Base64-encoded short description
	ShortDescription *string `json:"short_description,omitempty"`
	// Whether to use AI for fixes
	ShouldUseAiFix *bool `json:"should_use_ai_fix,omitempty"`
	// Rule tags
	Tags datadog.NullableList[string] `json:"tags,omitempty"`
	// Rule tests
	Tests []CustomRuleRevisionTest `json:"tests,omitempty"`
	// Tree-sitter query
	TreeSitterQuery *string `json:"tree_sitter_query,omitempty"`
	// Monotonically increasing version number of the revision.
	VersionId *int64 `json:"version_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewCustomRuleRevisionInput instantiates a new CustomRuleRevisionInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRuleRevisionInput() *CustomRuleRevisionInput {
	this := CustomRuleRevisionInput{}
	return &this
}

// NewCustomRuleRevisionInputWithDefaults instantiates a new CustomRuleRevisionInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleRevisionInputWithDefaults() *CustomRuleRevisionInput {
	this := CustomRuleRevisionInput{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRuleRevisionInput) GetArguments() []Argument {
	if o == nil {
		var ret []Argument
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInput) GetArgumentsOk() (*[]Argument, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasArguments() bool {
	return o != nil && o.Arguments != nil
}

// SetArguments gets a reference to the given []Argument and assigns it to the Arguments field.
func (o *CustomRuleRevisionInput) SetArguments(v []Argument) {
	o.Arguments = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetCategory() CustomRuleRevisionAttributesCategory {
	if o == nil || o.Category == nil {
		var ret CustomRuleRevisionAttributesCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetCategoryOk() (*CustomRuleRevisionAttributesCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given CustomRuleRevisionAttributesCategory and assigns it to the Category field.
func (o *CustomRuleRevisionInput) SetCategory(v CustomRuleRevisionAttributesCategory) {
	o.Category = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetChecksum() string {
	if o == nil || o.Checksum == nil {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetChecksumOk() (*string, bool) {
	if o == nil || o.Checksum == nil {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasChecksum() bool {
	return o != nil && o.Checksum != nil
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *CustomRuleRevisionInput) SetChecksum(v string) {
	o.Checksum = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CustomRuleRevisionInput) SetCode(v string) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CustomRuleRevisionInput) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CustomRuleRevisionInput) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreationMessage returns the CreationMessage field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetCreationMessage() string {
	if o == nil || o.CreationMessage == nil {
		var ret string
		return ret
	}
	return *o.CreationMessage
}

// GetCreationMessageOk returns a tuple with the CreationMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetCreationMessageOk() (*string, bool) {
	if o == nil || o.CreationMessage == nil {
		return nil, false
	}
	return o.CreationMessage, true
}

// HasCreationMessage returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCreationMessage() bool {
	return o != nil && o.CreationMessage != nil
}

// SetCreationMessage gets a reference to the given string and assigns it to the CreationMessage field.
func (o *CustomRuleRevisionInput) SetCreationMessage(v string) {
	o.CreationMessage = &v
}

// GetCve returns the Cve field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRuleRevisionInput) GetCve() string {
	if o == nil || o.Cve.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cve.Get()
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInput) GetCveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cve.Get(), o.Cve.IsSet()
}

// HasCve returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCve() bool {
	return o != nil && o.Cve.IsSet()
}

// SetCve gets a reference to the given datadog.NullableString and assigns it to the Cve field.
func (o *CustomRuleRevisionInput) SetCve(v string) {
	o.Cve.Set(&v)
}

// SetCveNil sets the value for Cve to be an explicit nil.
func (o *CustomRuleRevisionInput) SetCveNil() {
	o.Cve.Set(nil)
}

// UnsetCve ensures that no value is present for Cve, not even an explicit nil.
func (o *CustomRuleRevisionInput) UnsetCve() {
	o.Cve.Unset()
}

// GetCwe returns the Cwe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRuleRevisionInput) GetCwe() string {
	if o == nil || o.Cwe.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cwe.Get()
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInput) GetCweOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cwe.Get(), o.Cwe.IsSet()
}

// HasCwe returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasCwe() bool {
	return o != nil && o.Cwe.IsSet()
}

// SetCwe gets a reference to the given datadog.NullableString and assigns it to the Cwe field.
func (o *CustomRuleRevisionInput) SetCwe(v string) {
	o.Cwe.Set(&v)
}

// SetCweNil sets the value for Cwe to be an explicit nil.
func (o *CustomRuleRevisionInput) SetCweNil() {
	o.Cwe.Set(nil)
}

// UnsetCwe ensures that no value is present for Cwe, not even an explicit nil.
func (o *CustomRuleRevisionInput) UnsetCwe() {
	o.Cwe.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomRuleRevisionInput) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRuleRevisionInput) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl.Get()
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInput) GetDocumentationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentationUrl.Get(), o.DocumentationUrl.IsSet()
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasDocumentationUrl() bool {
	return o != nil && o.DocumentationUrl.IsSet()
}

// SetDocumentationUrl gets a reference to the given datadog.NullableString and assigns it to the DocumentationUrl field.
func (o *CustomRuleRevisionInput) SetDocumentationUrl(v string) {
	o.DocumentationUrl.Set(&v)
}

// SetDocumentationUrlNil sets the value for DocumentationUrl to be an explicit nil.
func (o *CustomRuleRevisionInput) SetDocumentationUrlNil() {
	o.DocumentationUrl.Set(nil)
}

// UnsetDocumentationUrl ensures that no value is present for DocumentationUrl, not even an explicit nil.
func (o *CustomRuleRevisionInput) UnsetDocumentationUrl() {
	o.DocumentationUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomRuleRevisionInput) SetId(v string) {
	o.Id = &v
}

// GetIsPublished returns the IsPublished field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetIsPublished() bool {
	if o == nil || o.IsPublished == nil {
		var ret bool
		return ret
	}
	return *o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetIsPublishedOk() (*bool, bool) {
	if o == nil || o.IsPublished == nil {
		return nil, false
	}
	return o.IsPublished, true
}

// HasIsPublished returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasIsPublished() bool {
	return o != nil && o.IsPublished != nil
}

// SetIsPublished gets a reference to the given bool and assigns it to the IsPublished field.
func (o *CustomRuleRevisionInput) SetIsPublished(v bool) {
	o.IsPublished = &v
}

// GetIsTesting returns the IsTesting field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetIsTesting() bool {
	if o == nil || o.IsTesting == nil {
		var ret bool
		return ret
	}
	return *o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetIsTestingOk() (*bool, bool) {
	if o == nil || o.IsTesting == nil {
		return nil, false
	}
	return o.IsTesting, true
}

// HasIsTesting returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasIsTesting() bool {
	return o != nil && o.IsTesting != nil
}

// SetIsTesting gets a reference to the given bool and assigns it to the IsTesting field.
func (o *CustomRuleRevisionInput) SetIsTesting(v bool) {
	o.IsTesting = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetLanguage() Language {
	if o == nil || o.Language == nil {
		var ret Language
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetLanguageOk() (*Language, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given Language and assigns it to the Language field.
func (o *CustomRuleRevisionInput) SetLanguage(v Language) {
	o.Language = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetSeverity() CustomRuleRevisionAttributesSeverity {
	if o == nil || o.Severity == nil {
		var ret CustomRuleRevisionAttributesSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetSeverityOk() (*CustomRuleRevisionAttributesSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given CustomRuleRevisionAttributesSeverity and assigns it to the Severity field.
func (o *CustomRuleRevisionInput) SetSeverity(v CustomRuleRevisionAttributesSeverity) {
	o.Severity = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetShortDescription() string {
	if o == nil || o.ShortDescription == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetShortDescriptionOk() (*string, bool) {
	if o == nil || o.ShortDescription == nil {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasShortDescription() bool {
	return o != nil && o.ShortDescription != nil
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *CustomRuleRevisionInput) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetShouldUseAiFix returns the ShouldUseAiFix field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetShouldUseAiFix() bool {
	if o == nil || o.ShouldUseAiFix == nil {
		var ret bool
		return ret
	}
	return *o.ShouldUseAiFix
}

// GetShouldUseAiFixOk returns a tuple with the ShouldUseAiFix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetShouldUseAiFixOk() (*bool, bool) {
	if o == nil || o.ShouldUseAiFix == nil {
		return nil, false
	}
	return o.ShouldUseAiFix, true
}

// HasShouldUseAiFix returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasShouldUseAiFix() bool {
	return o != nil && o.ShouldUseAiFix != nil
}

// SetShouldUseAiFix gets a reference to the given bool and assigns it to the ShouldUseAiFix field.
func (o *CustomRuleRevisionInput) SetShouldUseAiFix(v bool) {
	o.ShouldUseAiFix = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRuleRevisionInput) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInput) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasTags() bool {
	return o != nil && o.Tags.IsSet()
}

// SetTags gets a reference to the given datadog.NullableList[string] and assigns it to the Tags field.
func (o *CustomRuleRevisionInput) SetTags(v []string) {
	o.Tags.Set(&v)
}

// SetTagsNil sets the value for Tags to be an explicit nil.
func (o *CustomRuleRevisionInput) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil.
func (o *CustomRuleRevisionInput) UnsetTags() {
	o.Tags.Unset()
}

// GetTests returns the Tests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomRuleRevisionInput) GetTests() []CustomRuleRevisionTest {
	if o == nil {
		var ret []CustomRuleRevisionTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleRevisionInput) GetTestsOk() (*[]CustomRuleRevisionTest, bool) {
	if o == nil || o.Tests == nil {
		return nil, false
	}
	return &o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasTests() bool {
	return o != nil && o.Tests != nil
}

// SetTests gets a reference to the given []CustomRuleRevisionTest and assigns it to the Tests field.
func (o *CustomRuleRevisionInput) SetTests(v []CustomRuleRevisionTest) {
	o.Tests = v
}

// GetTreeSitterQuery returns the TreeSitterQuery field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetTreeSitterQuery() string {
	if o == nil || o.TreeSitterQuery == nil {
		var ret string
		return ret
	}
	return *o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil || o.TreeSitterQuery == nil {
		return nil, false
	}
	return o.TreeSitterQuery, true
}

// HasTreeSitterQuery returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasTreeSitterQuery() bool {
	return o != nil && o.TreeSitterQuery != nil
}

// SetTreeSitterQuery gets a reference to the given string and assigns it to the TreeSitterQuery field.
func (o *CustomRuleRevisionInput) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *CustomRuleRevisionInput) GetVersionId() int64 {
	if o == nil || o.VersionId == nil {
		var ret int64
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleRevisionInput) GetVersionIdOk() (*int64, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *CustomRuleRevisionInput) HasVersionId() bool {
	return o != nil && o.VersionId != nil
}

// SetVersionId gets a reference to the given int64 and assigns it to the VersionId field.
func (o *CustomRuleRevisionInput) SetVersionId(v int64) {
	o.VersionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRuleRevisionInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Checksum != nil {
		toSerialize["checksum"] = o.Checksum
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
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
	if o.CreationMessage != nil {
		toSerialize["creation_message"] = o.CreationMessage
	}
	if o.Cve.IsSet() {
		toSerialize["cve"] = o.Cve.Get()
	}
	if o.Cwe.IsSet() {
		toSerialize["cwe"] = o.Cwe.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentationUrl.IsSet() {
		toSerialize["documentation_url"] = o.DocumentationUrl.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsPublished != nil {
		toSerialize["is_published"] = o.IsPublished
	}
	if o.IsTesting != nil {
		toSerialize["is_testing"] = o.IsTesting
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.ShortDescription != nil {
		toSerialize["short_description"] = o.ShortDescription
	}
	if o.ShouldUseAiFix != nil {
		toSerialize["should_use_ai_fix"] = o.ShouldUseAiFix
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Tests != nil {
		toSerialize["tests"] = o.Tests
	}
	if o.TreeSitterQuery != nil {
		toSerialize["tree_sitter_query"] = o.TreeSitterQuery
	}
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRuleRevisionInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments        []Argument                            `json:"arguments,omitempty"`
		Category         *CustomRuleRevisionAttributesCategory `json:"category,omitempty"`
		Checksum         *string                               `json:"checksum,omitempty"`
		Code             *string                               `json:"code,omitempty"`
		CreatedAt        *time.Time                            `json:"created_at,omitempty"`
		CreatedBy        *string                               `json:"created_by,omitempty"`
		CreationMessage  *string                               `json:"creation_message,omitempty"`
		Cve              datadog.NullableString                `json:"cve,omitempty"`
		Cwe              datadog.NullableString                `json:"cwe,omitempty"`
		Description      *string                               `json:"description,omitempty"`
		DocumentationUrl datadog.NullableString                `json:"documentation_url,omitempty"`
		Id               *string                               `json:"id,omitempty"`
		IsPublished      *bool                                 `json:"is_published,omitempty"`
		IsTesting        *bool                                 `json:"is_testing,omitempty"`
		Language         *Language                             `json:"language,omitempty"`
		Severity         *CustomRuleRevisionAttributesSeverity `json:"severity,omitempty"`
		ShortDescription *string                               `json:"short_description,omitempty"`
		ShouldUseAiFix   *bool                                 `json:"should_use_ai_fix,omitempty"`
		Tags             datadog.NullableList[string]          `json:"tags,omitempty"`
		Tests            []CustomRuleRevisionTest              `json:"tests,omitempty"`
		TreeSitterQuery  *string                               `json:"tree_sitter_query,omitempty"`
		VersionId        *int64                                `json:"version_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.Arguments = all.Arguments
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.Checksum = all.Checksum
	o.Code = all.Code
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.CreationMessage = all.CreationMessage
	o.Cve = all.Cve
	o.Cwe = all.Cwe
	o.Description = all.Description
	o.DocumentationUrl = all.DocumentationUrl
	o.Id = all.Id
	o.IsPublished = all.IsPublished
	o.IsTesting = all.IsTesting
	if all.Language != nil && !all.Language.IsValid() {
		hasInvalidField = true
	} else {
		o.Language = all.Language
	}
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.ShortDescription = all.ShortDescription
	o.ShouldUseAiFix = all.ShouldUseAiFix
	o.Tags = all.Tags
	o.Tests = all.Tests
	o.TreeSitterQuery = all.TreeSitterQuery
	o.VersionId = all.VersionId

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
