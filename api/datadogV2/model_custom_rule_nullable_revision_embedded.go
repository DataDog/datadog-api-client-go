// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRuleNullableRevisionEmbedded The most recent revision of the rule, or null when the rule has no revisions.
type CustomRuleNullableRevisionEmbedded struct {
	// Rule arguments
	Arguments datadog.NullableList[Argument] `json:"arguments"`
	// Rule category
	Category CustomRuleRevisionAttributesCategory `json:"category"`
	// Code checksum
	Checksum string `json:"checksum"`
	// Rule code
	Code string `json:"code"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Creator identifier
	CreatedBy string `json:"created_by"`
	// Revision creation message
	CreationMessage string `json:"creation_message"`
	// Associated CVE. Omitted when the revision has no associated CVE.
	Cve *string `json:"cve,omitempty"`
	// Associated CWE. Omitted when the revision has no associated CWE.
	Cwe *string `json:"cwe,omitempty"`
	// Full description
	Description string `json:"description"`
	// Documentation URL. Omitted when the revision has no documentation URL.
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	// Revision identifier
	Id string `json:"id"`
	// Whether the revision is published
	IsPublished bool `json:"is_published"`
	// Whether this is a testing revision
	IsTesting bool `json:"is_testing"`
	// Programming language
	Language Language `json:"language"`
	// Rule severity
	Severity CustomRuleRevisionAttributesSeverity `json:"severity"`
	// Short description
	ShortDescription string `json:"short_description"`
	// Whether to use AI for fixes
	ShouldUseAiFix bool `json:"should_use_ai_fix"`
	// Rule tags
	Tags datadog.NullableList[string] `json:"tags"`
	// Rule tests
	Tests datadog.NullableList[CustomRuleRevisionTest] `json:"tests"`
	// Tree-sitter query
	TreeSitterQuery string `json:"tree_sitter_query"`
	// Monotonically increasing version number of the revision.
	VersionId int64 `json:"version_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRuleNullableRevisionEmbedded instantiates a new CustomRuleNullableRevisionEmbedded object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRuleNullableRevisionEmbedded(arguments datadog.NullableList[Argument], category CustomRuleRevisionAttributesCategory, checksum string, code string, createdAt time.Time, createdBy string, creationMessage string, description string, id string, isPublished bool, isTesting bool, language Language, severity CustomRuleRevisionAttributesSeverity, shortDescription string, shouldUseAiFix bool, tags datadog.NullableList[string], tests datadog.NullableList[CustomRuleRevisionTest], treeSitterQuery string, versionId int64) *CustomRuleNullableRevisionEmbedded {
	this := CustomRuleNullableRevisionEmbedded{}
	this.Arguments = arguments
	this.Category = category
	this.Checksum = checksum
	this.Code = code
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.CreationMessage = creationMessage
	this.Description = description
	this.Id = id
	this.IsPublished = isPublished
	this.IsTesting = isTesting
	this.Language = language
	this.Severity = severity
	this.ShortDescription = shortDescription
	this.ShouldUseAiFix = shouldUseAiFix
	this.Tags = tags
	this.Tests = tests
	this.TreeSitterQuery = treeSitterQuery
	this.VersionId = versionId
	return &this
}

// NewCustomRuleNullableRevisionEmbeddedWithDefaults instantiates a new CustomRuleNullableRevisionEmbedded object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRuleNullableRevisionEmbeddedWithDefaults() *CustomRuleNullableRevisionEmbedded {
	this := CustomRuleNullableRevisionEmbedded{}
	return &this
}

// GetArguments returns the Arguments field value.
// If the value is explicit nil, the zero value for []Argument will be returned.
func (o *CustomRuleNullableRevisionEmbedded) GetArguments() []Argument {
	if o == nil || o.Arguments.Get() == nil {
		var ret []Argument
		return ret
	}
	return *o.Arguments.Get()
}

// GetArgumentsOk returns a tuple with the Arguments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleNullableRevisionEmbedded) GetArgumentsOk() (*[]Argument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arguments.Get(), o.Arguments.IsSet()
}

// SetArguments sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetArguments(v []Argument) {
	o.Arguments.Set(&v)
}

// GetCategory returns the Category field value.
func (o *CustomRuleNullableRevisionEmbedded) GetCategory() CustomRuleRevisionAttributesCategory {
	if o == nil {
		var ret CustomRuleRevisionAttributesCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCategoryOk() (*CustomRuleRevisionAttributesCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetCategory(v CustomRuleRevisionAttributesCategory) {
	o.Category = v
}

// GetChecksum returns the Checksum field value.
func (o *CustomRuleNullableRevisionEmbedded) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetChecksum(v string) {
	o.Checksum = v
}

// GetCode returns the Code field value.
func (o *CustomRuleNullableRevisionEmbedded) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetCode(v string) {
	o.Code = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CustomRuleNullableRevisionEmbedded) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CustomRuleNullableRevisionEmbedded) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreationMessage returns the CreationMessage field value.
func (o *CustomRuleNullableRevisionEmbedded) GetCreationMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreationMessage
}

// GetCreationMessageOk returns a tuple with the CreationMessage field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCreationMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationMessage, true
}

// SetCreationMessage sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetCreationMessage(v string) {
	o.CreationMessage = v
}

// GetCve returns the Cve field value if set, zero value otherwise.
func (o *CustomRuleNullableRevisionEmbedded) GetCve() string {
	if o == nil || o.Cve == nil {
		var ret string
		return ret
	}
	return *o.Cve
}

// GetCveOk returns a tuple with the Cve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCveOk() (*string, bool) {
	if o == nil || o.Cve == nil {
		return nil, false
	}
	return o.Cve, true
}

// HasCve returns a boolean if a field has been set.
func (o *CustomRuleNullableRevisionEmbedded) HasCve() bool {
	return o != nil && o.Cve != nil
}

// SetCve gets a reference to the given string and assigns it to the Cve field.
func (o *CustomRuleNullableRevisionEmbedded) SetCve(v string) {
	o.Cve = &v
}

// GetCwe returns the Cwe field value if set, zero value otherwise.
func (o *CustomRuleNullableRevisionEmbedded) GetCwe() string {
	if o == nil || o.Cwe == nil {
		var ret string
		return ret
	}
	return *o.Cwe
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetCweOk() (*string, bool) {
	if o == nil || o.Cwe == nil {
		return nil, false
	}
	return o.Cwe, true
}

// HasCwe returns a boolean if a field has been set.
func (o *CustomRuleNullableRevisionEmbedded) HasCwe() bool {
	return o != nil && o.Cwe != nil
}

// SetCwe gets a reference to the given string and assigns it to the Cwe field.
func (o *CustomRuleNullableRevisionEmbedded) SetCwe(v string) {
	o.Cwe = &v
}

// GetDescription returns the Description field value.
func (o *CustomRuleNullableRevisionEmbedded) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetDescription(v string) {
	o.Description = v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *CustomRuleNullableRevisionEmbedded) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *CustomRuleNullableRevisionEmbedded) HasDocumentationUrl() bool {
	return o != nil && o.DocumentationUrl != nil
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *CustomRuleNullableRevisionEmbedded) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetId returns the Id field value.
func (o *CustomRuleNullableRevisionEmbedded) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetId(v string) {
	o.Id = v
}

// GetIsPublished returns the IsPublished field value.
func (o *CustomRuleNullableRevisionEmbedded) GetIsPublished() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPublished
}

// GetIsPublishedOk returns a tuple with the IsPublished field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetIsPublishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublished, true
}

// SetIsPublished sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetIsPublished(v bool) {
	o.IsPublished = v
}

// GetIsTesting returns the IsTesting field value.
func (o *CustomRuleNullableRevisionEmbedded) GetIsTesting() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsTesting
}

// GetIsTestingOk returns a tuple with the IsTesting field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetIsTestingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsTesting, true
}

// SetIsTesting sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetIsTesting(v bool) {
	o.IsTesting = v
}

// GetLanguage returns the Language field value.
func (o *CustomRuleNullableRevisionEmbedded) GetLanguage() Language {
	if o == nil {
		var ret Language
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetLanguageOk() (*Language, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetLanguage(v Language) {
	o.Language = v
}

// GetSeverity returns the Severity field value.
func (o *CustomRuleNullableRevisionEmbedded) GetSeverity() CustomRuleRevisionAttributesSeverity {
	if o == nil {
		var ret CustomRuleRevisionAttributesSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetSeverityOk() (*CustomRuleRevisionAttributesSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetSeverity(v CustomRuleRevisionAttributesSeverity) {
	o.Severity = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *CustomRuleNullableRevisionEmbedded) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetShortDescription(v string) {
	o.ShortDescription = v
}

// GetShouldUseAiFix returns the ShouldUseAiFix field value.
func (o *CustomRuleNullableRevisionEmbedded) GetShouldUseAiFix() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ShouldUseAiFix
}

// GetShouldUseAiFixOk returns a tuple with the ShouldUseAiFix field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetShouldUseAiFixOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldUseAiFix, true
}

// SetShouldUseAiFix sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetShouldUseAiFix(v bool) {
	o.ShouldUseAiFix = v
}

// GetTags returns the Tags field value.
// If the value is explicit nil, the zero value for []string will be returned.
func (o *CustomRuleNullableRevisionEmbedded) GetTags() []string {
	if o == nil || o.Tags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleNullableRevisionEmbedded) GetTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// SetTags sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetTags(v []string) {
	o.Tags.Set(&v)
}

// GetTests returns the Tests field value.
// If the value is explicit nil, the zero value for []CustomRuleRevisionTest will be returned.
func (o *CustomRuleNullableRevisionEmbedded) GetTests() []CustomRuleRevisionTest {
	if o == nil || o.Tests.Get() == nil {
		var ret []CustomRuleRevisionTest
		return ret
	}
	return *o.Tests.Get()
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRuleNullableRevisionEmbedded) GetTestsOk() (*[]CustomRuleRevisionTest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tests.Get(), o.Tests.IsSet()
}

// SetTests sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetTests(v []CustomRuleRevisionTest) {
	o.Tests.Set(&v)
}

// GetTreeSitterQuery returns the TreeSitterQuery field value.
func (o *CustomRuleNullableRevisionEmbedded) GetTreeSitterQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TreeSitterQuery
}

// GetTreeSitterQueryOk returns a tuple with the TreeSitterQuery field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetTreeSitterQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TreeSitterQuery, true
}

// SetTreeSitterQuery sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetTreeSitterQuery(v string) {
	o.TreeSitterQuery = v
}

// GetVersionId returns the VersionId field value.
func (o *CustomRuleNullableRevisionEmbedded) GetVersionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *CustomRuleNullableRevisionEmbedded) GetVersionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value.
func (o *CustomRuleNullableRevisionEmbedded) SetVersionId(v int64) {
	o.VersionId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRuleNullableRevisionEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arguments"] = o.Arguments.Get()
	toSerialize["category"] = o.Category
	toSerialize["checksum"] = o.Checksum
	toSerialize["code"] = o.Code
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["creation_message"] = o.CreationMessage
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
	toSerialize["id"] = o.Id
	toSerialize["is_published"] = o.IsPublished
	toSerialize["is_testing"] = o.IsTesting
	toSerialize["language"] = o.Language
	toSerialize["severity"] = o.Severity
	toSerialize["short_description"] = o.ShortDescription
	toSerialize["should_use_ai_fix"] = o.ShouldUseAiFix
	toSerialize["tags"] = o.Tags.Get()
	toSerialize["tests"] = o.Tests.Get()
	toSerialize["tree_sitter_query"] = o.TreeSitterQuery
	toSerialize["version_id"] = o.VersionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRuleNullableRevisionEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arguments        datadog.NullableList[Argument]               `json:"arguments"`
		Category         *CustomRuleRevisionAttributesCategory        `json:"category"`
		Checksum         *string                                      `json:"checksum"`
		Code             *string                                      `json:"code"`
		CreatedAt        *time.Time                                   `json:"created_at"`
		CreatedBy        *string                                      `json:"created_by"`
		CreationMessage  *string                                      `json:"creation_message"`
		Cve              *string                                      `json:"cve,omitempty"`
		Cwe              *string                                      `json:"cwe,omitempty"`
		Description      *string                                      `json:"description"`
		DocumentationUrl *string                                      `json:"documentation_url,omitempty"`
		Id               *string                                      `json:"id"`
		IsPublished      *bool                                        `json:"is_published"`
		IsTesting        *bool                                        `json:"is_testing"`
		Language         *Language                                    `json:"language"`
		Severity         *CustomRuleRevisionAttributesSeverity        `json:"severity"`
		ShortDescription *string                                      `json:"short_description"`
		ShouldUseAiFix   *bool                                        `json:"should_use_ai_fix"`
		Tags             datadog.NullableList[string]                 `json:"tags"`
		Tests            datadog.NullableList[CustomRuleRevisionTest] `json:"tests"`
		TreeSitterQuery  *string                                      `json:"tree_sitter_query"`
		VersionId        *int64                                       `json:"version_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Arguments.IsSet() {
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
	if all.CreationMessage == nil {
		return fmt.Errorf("required field creation_message missing")
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
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	if all.ShouldUseAiFix == nil {
		return fmt.Errorf("required field should_use_ai_fix missing")
	}
	if !all.Tags.IsSet() {
		return fmt.Errorf("required field tags missing")
	}
	if !all.Tests.IsSet() {
		return fmt.Errorf("required field tests missing")
	}
	if all.TreeSitterQuery == nil {
		return fmt.Errorf("required field tree_sitter_query missing")
	}
	if all.VersionId == nil {
		return fmt.Errorf("required field version_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arguments", "category", "checksum", "code", "created_at", "created_by", "creation_message", "cve", "cwe", "description", "documentation_url", "id", "is_published", "is_testing", "language", "severity", "short_description", "should_use_ai_fix", "tags", "tests", "tree_sitter_query", "version_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Arguments = all.Arguments
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Checksum = *all.Checksum
	o.Code = *all.Code
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.CreationMessage = *all.CreationMessage
	o.Cve = all.Cve
	o.Cwe = all.Cwe
	o.Description = *all.Description
	o.DocumentationUrl = all.DocumentationUrl
	o.Id = *all.Id
	o.IsPublished = *all.IsPublished
	o.IsTesting = *all.IsTesting
	if !all.Language.IsValid() {
		hasInvalidField = true
	} else {
		o.Language = *all.Language
	}
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.ShortDescription = *all.ShortDescription
	o.ShouldUseAiFix = *all.ShouldUseAiFix
	o.Tags = all.Tags
	o.Tests = all.Tests
	o.TreeSitterQuery = *all.TreeSitterQuery
	o.VersionId = *all.VersionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableCustomRuleNullableRevisionEmbedded handles when a null is used for CustomRuleNullableRevisionEmbedded.
type NullableCustomRuleNullableRevisionEmbedded struct {
	value *CustomRuleNullableRevisionEmbedded
	isSet bool
}

// Get returns the associated value.
func (v NullableCustomRuleNullableRevisionEmbedded) Get() *CustomRuleNullableRevisionEmbedded {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCustomRuleNullableRevisionEmbedded) Set(val *CustomRuleNullableRevisionEmbedded) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCustomRuleNullableRevisionEmbedded) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCustomRuleNullableRevisionEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCustomRuleNullableRevisionEmbedded initializes the struct as if Set has been called.
func NewNullableCustomRuleNullableRevisionEmbedded(val *CustomRuleNullableRevisionEmbedded) *NullableCustomRuleNullableRevisionEmbedded {
	return &NullableCustomRuleNullableRevisionEmbedded{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCustomRuleNullableRevisionEmbedded) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCustomRuleNullableRevisionEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
