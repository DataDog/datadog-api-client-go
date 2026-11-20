// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupPackageUpdate A package to update in a repository, for a group of library vulnerabilities (`sca.package_update`). The group
// gathers the open vulnerabilities of direct dependencies that have a recommended fix, and groups them by
// `@git.repository_id`, `@vulnerability.stack.ecosystem`, and `@remediation.recommended.name`.
type RemediationGroupPackageUpdate struct {
	// Number of distinct values of `@advisory.id` across the group's findings.
	AdvisoryCount int64 `json:"advisory_count"`
	// Package ecosystem, the findings' `@vulnerability.stack.ecosystem`.
	Ecosystem string `json:"ecosystem"`
	// Number of distinct values of `@advisory.id` across the group's findings that have
	// `@risk.has_exploit_available` set to `true`.
	ExploitableCveCount int64 `json:"exploitable_cve_count"`
	// Number of distinct values of `@package.declaration.block.filename`, the files that declare the package.
	FileCount int64 `json:"file_count"`
	// Name of the package to update, the findings' `@remediation.recommended.name`.
	PackageName string `json:"package_name"`
	// Repository that depends on the package, the findings' `@git.repository_id`.
	RepositoryId string `json:"repository_id"`
	// Version to update the package to, the highest `@remediation.recommended.version` across the group's
	// findings. Omitted when no version is recommended or when the recommended versions can't be ordered.
	TargetVersion *string `json:"target_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupPackageUpdate instantiates a new RemediationGroupPackageUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupPackageUpdate(advisoryCount int64, ecosystem string, exploitableCveCount int64, fileCount int64, packageName string, repositoryId string) *RemediationGroupPackageUpdate {
	this := RemediationGroupPackageUpdate{}
	this.AdvisoryCount = advisoryCount
	this.Ecosystem = ecosystem
	this.ExploitableCveCount = exploitableCveCount
	this.FileCount = fileCount
	this.PackageName = packageName
	this.RepositoryId = repositoryId
	return &this
}

// NewRemediationGroupPackageUpdateWithDefaults instantiates a new RemediationGroupPackageUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupPackageUpdateWithDefaults() *RemediationGroupPackageUpdate {
	this := RemediationGroupPackageUpdate{}
	return &this
}

// GetAdvisoryCount returns the AdvisoryCount field value.
func (o *RemediationGroupPackageUpdate) GetAdvisoryCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AdvisoryCount
}

// GetAdvisoryCountOk returns a tuple with the AdvisoryCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetAdvisoryCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvisoryCount, true
}

// SetAdvisoryCount sets field value.
func (o *RemediationGroupPackageUpdate) SetAdvisoryCount(v int64) {
	o.AdvisoryCount = v
}

// GetEcosystem returns the Ecosystem field value.
func (o *RemediationGroupPackageUpdate) GetEcosystem() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ecosystem
}

// GetEcosystemOk returns a tuple with the Ecosystem field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetEcosystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ecosystem, true
}

// SetEcosystem sets field value.
func (o *RemediationGroupPackageUpdate) SetEcosystem(v string) {
	o.Ecosystem = v
}

// GetExploitableCveCount returns the ExploitableCveCount field value.
func (o *RemediationGroupPackageUpdate) GetExploitableCveCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExploitableCveCount
}

// GetExploitableCveCountOk returns a tuple with the ExploitableCveCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetExploitableCveCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableCveCount, true
}

// SetExploitableCveCount sets field value.
func (o *RemediationGroupPackageUpdate) SetExploitableCveCount(v int64) {
	o.ExploitableCveCount = v
}

// GetFileCount returns the FileCount field value.
func (o *RemediationGroupPackageUpdate) GetFileCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileCount, true
}

// SetFileCount sets field value.
func (o *RemediationGroupPackageUpdate) SetFileCount(v int64) {
	o.FileCount = v
}

// GetPackageName returns the PackageName field value.
func (o *RemediationGroupPackageUpdate) GetPackageName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetPackageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageName, true
}

// SetPackageName sets field value.
func (o *RemediationGroupPackageUpdate) SetPackageName(v string) {
	o.PackageName = v
}

// GetRepositoryId returns the RepositoryId field value.
func (o *RemediationGroupPackageUpdate) GetRepositoryId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryId, true
}

// SetRepositoryId sets field value.
func (o *RemediationGroupPackageUpdate) SetRepositoryId(v string) {
	o.RepositoryId = v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *RemediationGroupPackageUpdate) GetTargetVersion() string {
	if o == nil || o.TargetVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationGroupPackageUpdate) GetTargetVersionOk() (*string, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *RemediationGroupPackageUpdate) HasTargetVersion() bool {
	return o != nil && o.TargetVersion != nil
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *RemediationGroupPackageUpdate) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupPackageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["advisory_count"] = o.AdvisoryCount
	toSerialize["ecosystem"] = o.Ecosystem
	toSerialize["exploitable_cve_count"] = o.ExploitableCveCount
	toSerialize["file_count"] = o.FileCount
	toSerialize["package_name"] = o.PackageName
	toSerialize["repository_id"] = o.RepositoryId
	if o.TargetVersion != nil {
		toSerialize["target_version"] = o.TargetVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupPackageUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdvisoryCount       *int64  `json:"advisory_count"`
		Ecosystem           *string `json:"ecosystem"`
		ExploitableCveCount *int64  `json:"exploitable_cve_count"`
		FileCount           *int64  `json:"file_count"`
		PackageName         *string `json:"package_name"`
		RepositoryId        *string `json:"repository_id"`
		TargetVersion       *string `json:"target_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AdvisoryCount == nil {
		return fmt.Errorf("required field advisory_count missing")
	}
	if all.Ecosystem == nil {
		return fmt.Errorf("required field ecosystem missing")
	}
	if all.ExploitableCveCount == nil {
		return fmt.Errorf("required field exploitable_cve_count missing")
	}
	if all.FileCount == nil {
		return fmt.Errorf("required field file_count missing")
	}
	if all.PackageName == nil {
		return fmt.Errorf("required field package_name missing")
	}
	if all.RepositoryId == nil {
		return fmt.Errorf("required field repository_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"advisory_count", "ecosystem", "exploitable_cve_count", "file_count", "package_name", "repository_id", "target_version"})
	} else {
		return err
	}
	o.AdvisoryCount = *all.AdvisoryCount
	o.Ecosystem = *all.Ecosystem
	o.ExploitableCveCount = *all.ExploitableCveCount
	o.FileCount = *all.FileCount
	o.PackageName = *all.PackageName
	o.RepositoryId = *all.RepositoryId
	o.TargetVersion = all.TargetVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
