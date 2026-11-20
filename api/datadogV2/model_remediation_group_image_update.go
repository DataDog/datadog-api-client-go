// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupImageUpdate An image or package to upgrade, for a group of host and container vulnerabilities (`cloudvm.image_update`).
// The group gathers the open vulnerabilities of running production images that have a fix available, and groups
// them by `grouping_location`, `remediation_kind`, and `action`.
type RemediationGroupImageUpdate struct {
	// Name of the container image, base image, or package to upgrade. The value comes from the first of these
	// attributes that the findings have: `@remediation.container_image.latest_major.name`,
	// `@remediation.base_image.latest_major.name`, or `@remediation.recommended.name`.
	Action string `json:"action"`
	// Number of distinct values of `@advisory.id` across the group's findings.
	AdvisoryCount int64 `json:"advisory_count"`
	// Number of distinct values of `@advisory.id` across the group's findings that have
	// `@risk.has_exploit_available` set to `true`.
	ExploitableCveCount int64 `json:"exploitable_cve_count"`
	// Number of distinct values of `@code_location.filename` across the group's findings.
	FileCount int64 `json:"file_count"`
	// File that contains the findings, from their `@code_location.filename` attribute. When no file is known,
	// the value is the findings' `@container_image.name`.
	GroupingLocation string `json:"grouping_location"`
	// Whether `grouping_location` is a `@code_location.filename` rather than a `@container_image.name`.
	LocationIsFilename bool `json:"location_is_filename"`
	// Kind of upgrade: `image_update` when the findings have a `@remediation.container_image.latest_major.name`,
	// `base_image_bump` when they have `@vulnerability.is_inherited_from_base_image` set to `true`, and
	// `package_update` otherwise.
	RemediationKind string `json:"remediation_kind"`
	// Version to upgrade to: the `@remediation.container_image.latest_major.tag` or
	// `@remediation.base_image.latest_major.tag` of the image in `action`, or else the highest
	// `@remediation.recommended.version` across the group's findings. Omitted when no version is recommended
	// or when the recommended versions can't be ordered.
	TargetVersion *string `json:"target_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupImageUpdate instantiates a new RemediationGroupImageUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupImageUpdate(action string, advisoryCount int64, exploitableCveCount int64, fileCount int64, groupingLocation string, locationIsFilename bool, remediationKind string) *RemediationGroupImageUpdate {
	this := RemediationGroupImageUpdate{}
	this.Action = action
	this.AdvisoryCount = advisoryCount
	this.ExploitableCveCount = exploitableCveCount
	this.FileCount = fileCount
	this.GroupingLocation = groupingLocation
	this.LocationIsFilename = locationIsFilename
	this.RemediationKind = remediationKind
	return &this
}

// NewRemediationGroupImageUpdateWithDefaults instantiates a new RemediationGroupImageUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupImageUpdateWithDefaults() *RemediationGroupImageUpdate {
	this := RemediationGroupImageUpdate{}
	return &this
}

// GetAction returns the Action field value.
func (o *RemediationGroupImageUpdate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *RemediationGroupImageUpdate) SetAction(v string) {
	o.Action = v
}

// GetAdvisoryCount returns the AdvisoryCount field value.
func (o *RemediationGroupImageUpdate) GetAdvisoryCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AdvisoryCount
}

// GetAdvisoryCountOk returns a tuple with the AdvisoryCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetAdvisoryCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdvisoryCount, true
}

// SetAdvisoryCount sets field value.
func (o *RemediationGroupImageUpdate) SetAdvisoryCount(v int64) {
	o.AdvisoryCount = v
}

// GetExploitableCveCount returns the ExploitableCveCount field value.
func (o *RemediationGroupImageUpdate) GetExploitableCveCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ExploitableCveCount
}

// GetExploitableCveCountOk returns a tuple with the ExploitableCveCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetExploitableCveCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExploitableCveCount, true
}

// SetExploitableCveCount sets field value.
func (o *RemediationGroupImageUpdate) SetExploitableCveCount(v int64) {
	o.ExploitableCveCount = v
}

// GetFileCount returns the FileCount field value.
func (o *RemediationGroupImageUpdate) GetFileCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileCount, true
}

// SetFileCount sets field value.
func (o *RemediationGroupImageUpdate) SetFileCount(v int64) {
	o.FileCount = v
}

// GetGroupingLocation returns the GroupingLocation field value.
func (o *RemediationGroupImageUpdate) GetGroupingLocation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.GroupingLocation
}

// GetGroupingLocationOk returns a tuple with the GroupingLocation field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetGroupingLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupingLocation, true
}

// SetGroupingLocation sets field value.
func (o *RemediationGroupImageUpdate) SetGroupingLocation(v string) {
	o.GroupingLocation = v
}

// GetLocationIsFilename returns the LocationIsFilename field value.
func (o *RemediationGroupImageUpdate) GetLocationIsFilename() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.LocationIsFilename
}

// GetLocationIsFilenameOk returns a tuple with the LocationIsFilename field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetLocationIsFilenameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationIsFilename, true
}

// SetLocationIsFilename sets field value.
func (o *RemediationGroupImageUpdate) SetLocationIsFilename(v bool) {
	o.LocationIsFilename = v
}

// GetRemediationKind returns the RemediationKind field value.
func (o *RemediationGroupImageUpdate) GetRemediationKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RemediationKind
}

// GetRemediationKindOk returns a tuple with the RemediationKind field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetRemediationKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationKind, true
}

// SetRemediationKind sets field value.
func (o *RemediationGroupImageUpdate) SetRemediationKind(v string) {
	o.RemediationKind = v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *RemediationGroupImageUpdate) GetTargetVersion() string {
	if o == nil || o.TargetVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationGroupImageUpdate) GetTargetVersionOk() (*string, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *RemediationGroupImageUpdate) HasTargetVersion() bool {
	return o != nil && o.TargetVersion != nil
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *RemediationGroupImageUpdate) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupImageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["advisory_count"] = o.AdvisoryCount
	toSerialize["exploitable_cve_count"] = o.ExploitableCveCount
	toSerialize["file_count"] = o.FileCount
	toSerialize["grouping_location"] = o.GroupingLocation
	toSerialize["location_is_filename"] = o.LocationIsFilename
	toSerialize["remediation_kind"] = o.RemediationKind
	if o.TargetVersion != nil {
		toSerialize["target_version"] = o.TargetVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupImageUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action              *string `json:"action"`
		AdvisoryCount       *int64  `json:"advisory_count"`
		ExploitableCveCount *int64  `json:"exploitable_cve_count"`
		FileCount           *int64  `json:"file_count"`
		GroupingLocation    *string `json:"grouping_location"`
		LocationIsFilename  *bool   `json:"location_is_filename"`
		RemediationKind     *string `json:"remediation_kind"`
		TargetVersion       *string `json:"target_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.AdvisoryCount == nil {
		return fmt.Errorf("required field advisory_count missing")
	}
	if all.ExploitableCveCount == nil {
		return fmt.Errorf("required field exploitable_cve_count missing")
	}
	if all.FileCount == nil {
		return fmt.Errorf("required field file_count missing")
	}
	if all.GroupingLocation == nil {
		return fmt.Errorf("required field grouping_location missing")
	}
	if all.LocationIsFilename == nil {
		return fmt.Errorf("required field location_is_filename missing")
	}
	if all.RemediationKind == nil {
		return fmt.Errorf("required field remediation_kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "advisory_count", "exploitable_cve_count", "file_count", "grouping_location", "location_is_filename", "remediation_kind", "target_version"})
	} else {
		return err
	}
	o.Action = *all.Action
	o.AdvisoryCount = *all.AdvisoryCount
	o.ExploitableCveCount = *all.ExploitableCveCount
	o.FileCount = *all.FileCount
	o.GroupingLocation = *all.GroupingLocation
	o.LocationIsFilename = *all.LocationIsFilename
	o.RemediationKind = *all.RemediationKind
	o.TargetVersion = all.TargetVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
