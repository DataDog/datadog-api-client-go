// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupAttributes Attributes of a remediation group. Each group includes one kind object, depending on its grouping definition.
type RemediationGroupAttributes struct {
	// Counts of the group's open findings. The counts include every finding in the group, including findings that
	// don't match `filter[query]`.
	Findings RemediationGroupFindings `json:"findings"`
	// Grouping definition that formed the group, such as `sca.package_update`.
	GroupingDefinitionId string `json:"grouping_definition_id"`
	// Who should fix the group: the code owners of the group's findings, from their `@git.codeowners` attribute,
	// and their Datadog teams, from their `team` tag. A group with no owners has empty `items` and a `total` of 0.
	Owners RemediationGroupOwners `json:"owners"`
	// A package to update in a repository, for a group of library vulnerabilities (`sca.package_update`). The group
	// gathers the open vulnerabilities of direct dependencies that have a recommended fix, and groups them by
	// `@git.repository_id`, `@vulnerability.stack.ecosystem`, and `@remediation.recommended.name`.
	PackageUpdate *RemediationGroupPackageUpdate `json:"package_update,omitempty"`
	// Title of the group, which describes the change to make.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupAttributes instantiates a new RemediationGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupAttributes(findings RemediationGroupFindings, groupingDefinitionId string, owners RemediationGroupOwners, title string) *RemediationGroupAttributes {
	this := RemediationGroupAttributes{}
	this.Findings = findings
	this.GroupingDefinitionId = groupingDefinitionId
	this.Owners = owners
	this.Title = title
	return &this
}

// NewRemediationGroupAttributesWithDefaults instantiates a new RemediationGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupAttributesWithDefaults() *RemediationGroupAttributes {
	this := RemediationGroupAttributes{}
	return &this
}

// GetFindings returns the Findings field value.
func (o *RemediationGroupAttributes) GetFindings() RemediationGroupFindings {
	if o == nil {
		var ret RemediationGroupFindings
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupAttributes) GetFindingsOk() (*RemediationGroupFindings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *RemediationGroupAttributes) SetFindings(v RemediationGroupFindings) {
	o.Findings = v
}

// GetGroupingDefinitionId returns the GroupingDefinitionId field value.
func (o *RemediationGroupAttributes) GetGroupingDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupingDefinitionId
}

// GetGroupingDefinitionIdOk returns a tuple with the GroupingDefinitionId field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupAttributes) GetGroupingDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupingDefinitionId, true
}

// SetGroupingDefinitionId sets field value.
func (o *RemediationGroupAttributes) SetGroupingDefinitionId(v string) {
	o.GroupingDefinitionId = v
}

// GetOwners returns the Owners field value.
func (o *RemediationGroupAttributes) GetOwners() RemediationGroupOwners {
	if o == nil {
		var ret RemediationGroupOwners
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupAttributes) GetOwnersOk() (*RemediationGroupOwners, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value.
func (o *RemediationGroupAttributes) SetOwners(v RemediationGroupOwners) {
	o.Owners = v
}

// GetPackageUpdate returns the PackageUpdate field value if set, zero value otherwise.
func (o *RemediationGroupAttributes) GetPackageUpdate() RemediationGroupPackageUpdate {
	if o == nil || o.PackageUpdate == nil {
		var ret RemediationGroupPackageUpdate
		return ret
	}
	return *o.PackageUpdate
}

// GetPackageUpdateOk returns a tuple with the PackageUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationGroupAttributes) GetPackageUpdateOk() (*RemediationGroupPackageUpdate, bool) {
	if o == nil || o.PackageUpdate == nil {
		return nil, false
	}
	return o.PackageUpdate, true
}

// HasPackageUpdate returns a boolean if a field has been set.
func (o *RemediationGroupAttributes) HasPackageUpdate() bool {
	return o != nil && o.PackageUpdate != nil
}

// SetPackageUpdate gets a reference to the given RemediationGroupPackageUpdate and assigns it to the PackageUpdate field.
func (o *RemediationGroupAttributes) SetPackageUpdate(v RemediationGroupPackageUpdate) {
	o.PackageUpdate = &v
}

// GetTitle returns the Title field value.
func (o *RemediationGroupAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *RemediationGroupAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["findings"] = o.Findings
	toSerialize["grouping_definition_id"] = o.GroupingDefinitionId
	toSerialize["owners"] = o.Owners
	if o.PackageUpdate != nil {
		toSerialize["package_update"] = o.PackageUpdate
	}
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Findings             *RemediationGroupFindings      `json:"findings"`
		GroupingDefinitionId *string                        `json:"grouping_definition_id"`
		Owners               *RemediationGroupOwners        `json:"owners"`
		PackageUpdate        *RemediationGroupPackageUpdate `json:"package_update,omitempty"`
		Title                *string                        `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Findings == nil {
		return fmt.Errorf("required field findings missing")
	}
	if all.GroupingDefinitionId == nil {
		return fmt.Errorf("required field grouping_definition_id missing")
	}
	if all.Owners == nil {
		return fmt.Errorf("required field owners missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"findings", "grouping_definition_id", "owners", "package_update", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Findings.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Findings = *all.Findings
	o.GroupingDefinitionId = *all.GroupingDefinitionId
	if all.Owners.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Owners = *all.Owners
	if all.PackageUpdate != nil && all.PackageUpdate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PackageUpdate = all.PackageUpdate
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
