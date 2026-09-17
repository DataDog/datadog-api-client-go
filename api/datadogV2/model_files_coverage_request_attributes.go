// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FilesCoverageRequestAttributes Attributes for requesting per-file code coverage data. Exactly one of `commit_sha`, `branch`, or `pr_number` must be provided. At most one of `service`, `codeowner`, or `flag` may be provided.
type FilesCoverageRequestAttributes struct {
	// The branch name.
	Branch *string `json:"branch,omitempty"`
	// When true, return coverage data only for files that were changed in the specified scope.
	ChangedOnly *bool `json:"changed_only,omitempty"`
	// Filter coverage by code owner. At most one of `service`, `codeowner`, or `flag` may be provided.
	Codeowner *string `json:"codeowner,omitempty"`
	// The commit SHA (40-character hexadecimal string).
	CommitSha *string `json:"commit_sha,omitempty"`
	// Filter coverage by coverage flag. At most one of `service`, `codeowner`, or `flag` may be provided.
	Flag *string `json:"flag,omitempty"`
	// The pull request number. Must be a positive integer.
	PrNumber *int64 `json:"pr_number,omitempty"`
	// Deprecated: use `repository_url` instead. The repository URL.
	// Deprecated
	RepositoryId *string `json:"repository_id,omitempty"`
	// The repository URL. Accepts a full URL with or without a scheme (for example, `https://github.com/org/repo` or `github.com/org/repo`).
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// Filter coverage by service name. At most one of `service`, `codeowner`, or `flag` may be provided.
	Service *string `json:"service,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFilesCoverageRequestAttributes instantiates a new FilesCoverageRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFilesCoverageRequestAttributes() *FilesCoverageRequestAttributes {
	this := FilesCoverageRequestAttributes{}
	return &this
}

// NewFilesCoverageRequestAttributesWithDefaults instantiates a new FilesCoverageRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFilesCoverageRequestAttributesWithDefaults() *FilesCoverageRequestAttributes {
	this := FilesCoverageRequestAttributes{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasBranch() bool {
	return o != nil && o.Branch != nil
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *FilesCoverageRequestAttributes) SetBranch(v string) {
	o.Branch = &v
}

// GetChangedOnly returns the ChangedOnly field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetChangedOnly() bool {
	if o == nil || o.ChangedOnly == nil {
		var ret bool
		return ret
	}
	return *o.ChangedOnly
}

// GetChangedOnlyOk returns a tuple with the ChangedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetChangedOnlyOk() (*bool, bool) {
	if o == nil || o.ChangedOnly == nil {
		return nil, false
	}
	return o.ChangedOnly, true
}

// HasChangedOnly returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasChangedOnly() bool {
	return o != nil && o.ChangedOnly != nil
}

// SetChangedOnly gets a reference to the given bool and assigns it to the ChangedOnly field.
func (o *FilesCoverageRequestAttributes) SetChangedOnly(v bool) {
	o.ChangedOnly = &v
}

// GetCodeowner returns the Codeowner field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetCodeowner() string {
	if o == nil || o.Codeowner == nil {
		var ret string
		return ret
	}
	return *o.Codeowner
}

// GetCodeownerOk returns a tuple with the Codeowner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetCodeownerOk() (*string, bool) {
	if o == nil || o.Codeowner == nil {
		return nil, false
	}
	return o.Codeowner, true
}

// HasCodeowner returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasCodeowner() bool {
	return o != nil && o.Codeowner != nil
}

// SetCodeowner gets a reference to the given string and assigns it to the Codeowner field.
func (o *FilesCoverageRequestAttributes) SetCodeowner(v string) {
	o.Codeowner = &v
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetCommitSha() string {
	if o == nil || o.CommitSha == nil {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetCommitShaOk() (*string, bool) {
	if o == nil || o.CommitSha == nil {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasCommitSha() bool {
	return o != nil && o.CommitSha != nil
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *FilesCoverageRequestAttributes) SetCommitSha(v string) {
	o.CommitSha = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasFlag() bool {
	return o != nil && o.Flag != nil
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *FilesCoverageRequestAttributes) SetFlag(v string) {
	o.Flag = &v
}

// GetPrNumber returns the PrNumber field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetPrNumber() int64 {
	if o == nil || o.PrNumber == nil {
		var ret int64
		return ret
	}
	return *o.PrNumber
}

// GetPrNumberOk returns a tuple with the PrNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetPrNumberOk() (*int64, bool) {
	if o == nil || o.PrNumber == nil {
		return nil, false
	}
	return o.PrNumber, true
}

// HasPrNumber returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasPrNumber() bool {
	return o != nil && o.PrNumber != nil
}

// SetPrNumber gets a reference to the given int64 and assigns it to the PrNumber field.
func (o *FilesCoverageRequestAttributes) SetPrNumber(v int64) {
	o.PrNumber = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
// Deprecated
func (o *FilesCoverageRequestAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *FilesCoverageRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
// Deprecated
func (o *FilesCoverageRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasRepositoryUrl() bool {
	return o != nil && o.RepositoryUrl != nil
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *FilesCoverageRequestAttributes) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FilesCoverageRequestAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesCoverageRequestAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FilesCoverageRequestAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FilesCoverageRequestAttributes) SetService(v string) {
	o.Service = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FilesCoverageRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.ChangedOnly != nil {
		toSerialize["changed_only"] = o.ChangedOnly
	}
	if o.Codeowner != nil {
		toSerialize["codeowner"] = o.Codeowner
	}
	if o.CommitSha != nil {
		toSerialize["commit_sha"] = o.CommitSha
	}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	if o.PrNumber != nil {
		toSerialize["pr_number"] = o.PrNumber
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if o.RepositoryUrl != nil {
		toSerialize["repository_url"] = o.RepositoryUrl
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FilesCoverageRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Branch        *string `json:"branch,omitempty"`
		ChangedOnly   *bool   `json:"changed_only,omitempty"`
		Codeowner     *string `json:"codeowner,omitempty"`
		CommitSha     *string `json:"commit_sha,omitempty"`
		Flag          *string `json:"flag,omitempty"`
		PrNumber      *int64  `json:"pr_number,omitempty"`
		RepositoryId  *string `json:"repository_id,omitempty"`
		RepositoryUrl *string `json:"repository_url,omitempty"`
		Service       *string `json:"service,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"branch", "changed_only", "codeowner", "commit_sha", "flag", "pr_number", "repository_id", "repository_url", "service"})
	} else {
		return err
	}
	o.Branch = all.Branch
	o.ChangedOnly = all.ChangedOnly
	o.Codeowner = all.Codeowner
	o.CommitSha = all.CommitSha
	o.Flag = all.Flag
	o.PrNumber = all.PrNumber
	o.RepositoryId = all.RepositoryId
	o.RepositoryUrl = all.RepositoryUrl
	o.Service = all.Service

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
