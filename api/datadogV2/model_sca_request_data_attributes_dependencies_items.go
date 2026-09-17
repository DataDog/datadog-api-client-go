// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItems A dependency found in the repository, including its identity, location, and reachability metadata.
type ScaRequestDataAttributesDependenciesItems struct {
	// A list of patterns or identifiers that should be excluded from analysis for this dependency.
	Exclusions []string `json:"exclusions,omitempty"`
	// The group or organization namespace of the dependency (e.g., Maven group ID).
	Group datadog.NullableString `json:"group,omitempty"`
	// Indicates whether this is a development-only dependency not used in production.
	IsDev *bool `json:"is_dev,omitempty"`
	// Indicates whether this is a direct dependency (as opposed to a transitive one).
	IsDirect datadog.NullableBool `json:"is_direct,omitempty"`
	// The programming language ecosystem of this dependency (e.g., java, python, javascript).
	Language *string `json:"language,omitempty"`
	// The list of source file locations where this dependency is declared.
	Locations []ScaRequestDataAttributesDependenciesItemsLocationsItems `json:"locations,omitempty"`
	// The name of the dependency package.
	Name *string `json:"name,omitempty"`
	// Indicates whether dependency details are intentionally opaque.
	Opaque *bool `json:"opaque,omitempty"`
	// The package manager responsible for this dependency (e.g., maven, pip, npm).
	PackageManager *string `json:"package_manager,omitempty"`
	// The Package URL (PURL) uniquely identifying this dependency.
	Purl *string `json:"purl,omitempty"`
	// Properties describing symbols from this dependency that are reachable in the application code.
	ReachableSymbolProperties []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems `json:"reachable_symbol_properties,omitempty"`
	// Indicates whether this dependency requires transitive dependency enrichment.
	RequiresTransitiveEnrichment *bool `json:"requires_transitive_enrichment,omitempty"`
	// The target framework identifiers associated with this dependency.
	TargetFrameworks []string `json:"target_frameworks,omitempty"`
	// The version of the dependency.
	Version datadog.NullableString `json:"version,omitempty"`
	// Indicates whether the version value represents a version constraint.
	VersionConstraint *bool `json:"version_constraint,omitempty"`
	// The version range associated with this dependency when a manifest declares a range.
	VersionRange *string `json:"version_range,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItems instantiates a new ScaRequestDataAttributesDependenciesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItems() *ScaRequestDataAttributesDependenciesItems {
	this := ScaRequestDataAttributesDependenciesItems{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsWithDefaults() *ScaRequestDataAttributesDependenciesItems {
	this := ScaRequestDataAttributesDependenciesItems{}
	return &this
}

// GetExclusions returns the Exclusions field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetExclusions() []string {
	if o == nil || o.Exclusions == nil {
		var ret []string
		return ret
	}
	return o.Exclusions
}

// GetExclusionsOk returns a tuple with the Exclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetExclusionsOk() (*[]string, bool) {
	if o == nil || o.Exclusions == nil {
		return nil, false
	}
	return &o.Exclusions, true
}

// HasExclusions returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasExclusions() bool {
	return o != nil && o.Exclusions != nil
}

// SetExclusions gets a reference to the given []string and assigns it to the Exclusions field.
func (o *ScaRequestDataAttributesDependenciesItems) SetExclusions(v []string) {
	o.Exclusions = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItems) GetGroup() string {
	if o == nil || o.Group.Get() == nil {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItems) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasGroup() bool {
	return o != nil && o.Group.IsSet()
}

// SetGroup gets a reference to the given datadog.NullableString and assigns it to the Group field.
func (o *ScaRequestDataAttributesDependenciesItems) SetGroup(v string) {
	o.Group.Set(&v)
}

// SetGroupNil sets the value for Group to be an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItems) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItems) UnsetGroup() {
	o.Group.Unset()
}

// GetIsDev returns the IsDev field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDev() bool {
	if o == nil || o.IsDev == nil {
		var ret bool
		return ret
	}
	return *o.IsDev
}

// GetIsDevOk returns a tuple with the IsDev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDevOk() (*bool, bool) {
	if o == nil || o.IsDev == nil {
		return nil, false
	}
	return o.IsDev, true
}

// HasIsDev returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasIsDev() bool {
	return o != nil && o.IsDev != nil
}

// SetIsDev gets a reference to the given bool and assigns it to the IsDev field.
func (o *ScaRequestDataAttributesDependenciesItems) SetIsDev(v bool) {
	o.IsDev = &v
}

// GetIsDirect returns the IsDirect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDirect() bool {
	if o == nil || o.IsDirect.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDirect.Get()
}

// GetIsDirectOk returns a tuple with the IsDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDirectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDirect.Get(), o.IsDirect.IsSet()
}

// HasIsDirect returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasIsDirect() bool {
	return o != nil && o.IsDirect.IsSet()
}

// SetIsDirect gets a reference to the given datadog.NullableBool and assigns it to the IsDirect field.
func (o *ScaRequestDataAttributesDependenciesItems) SetIsDirect(v bool) {
	o.IsDirect.Set(&v)
}

// SetIsDirectNil sets the value for IsDirect to be an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItems) SetIsDirectNil() {
	o.IsDirect.Set(nil)
}

// UnsetIsDirect ensures that no value is present for IsDirect, not even an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItems) UnsetIsDirect() {
	o.IsDirect.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ScaRequestDataAttributesDependenciesItems) SetLanguage(v string) {
	o.Language = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItems) GetLocations() []ScaRequestDataAttributesDependenciesItemsLocationsItems {
	if o == nil {
		var ret []ScaRequestDataAttributesDependenciesItemsLocationsItems
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItems) GetLocationsOk() (*[]ScaRequestDataAttributesDependenciesItemsLocationsItems, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return &o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasLocations() bool {
	return o != nil && o.Locations != nil
}

// SetLocations gets a reference to the given []ScaRequestDataAttributesDependenciesItemsLocationsItems and assigns it to the Locations field.
func (o *ScaRequestDataAttributesDependenciesItems) SetLocations(v []ScaRequestDataAttributesDependenciesItemsLocationsItems) {
	o.Locations = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScaRequestDataAttributesDependenciesItems) SetName(v string) {
	o.Name = &v
}

// GetOpaque returns the Opaque field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetOpaque() bool {
	if o == nil || o.Opaque == nil {
		var ret bool
		return ret
	}
	return *o.Opaque
}

// GetOpaqueOk returns a tuple with the Opaque field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetOpaqueOk() (*bool, bool) {
	if o == nil || o.Opaque == nil {
		return nil, false
	}
	return o.Opaque, true
}

// HasOpaque returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasOpaque() bool {
	return o != nil && o.Opaque != nil
}

// SetOpaque gets a reference to the given bool and assigns it to the Opaque field.
func (o *ScaRequestDataAttributesDependenciesItems) SetOpaque(v bool) {
	o.Opaque = &v
}

// GetPackageManager returns the PackageManager field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetPackageManager() string {
	if o == nil || o.PackageManager == nil {
		var ret string
		return ret
	}
	return *o.PackageManager
}

// GetPackageManagerOk returns a tuple with the PackageManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetPackageManagerOk() (*string, bool) {
	if o == nil || o.PackageManager == nil {
		return nil, false
	}
	return o.PackageManager, true
}

// HasPackageManager returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasPackageManager() bool {
	return o != nil && o.PackageManager != nil
}

// SetPackageManager gets a reference to the given string and assigns it to the PackageManager field.
func (o *ScaRequestDataAttributesDependenciesItems) SetPackageManager(v string) {
	o.PackageManager = &v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasPurl() bool {
	return o != nil && o.Purl != nil
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *ScaRequestDataAttributesDependenciesItems) SetPurl(v string) {
	o.Purl = &v
}

// GetReachableSymbolProperties returns the ReachableSymbolProperties field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetReachableSymbolProperties() []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems {
	if o == nil || o.ReachableSymbolProperties == nil {
		var ret []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems
		return ret
	}
	return o.ReachableSymbolProperties
}

// GetReachableSymbolPropertiesOk returns a tuple with the ReachableSymbolProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetReachableSymbolPropertiesOk() (*[]ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems, bool) {
	if o == nil || o.ReachableSymbolProperties == nil {
		return nil, false
	}
	return &o.ReachableSymbolProperties, true
}

// HasReachableSymbolProperties returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasReachableSymbolProperties() bool {
	return o != nil && o.ReachableSymbolProperties != nil
}

// SetReachableSymbolProperties gets a reference to the given []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems and assigns it to the ReachableSymbolProperties field.
func (o *ScaRequestDataAttributesDependenciesItems) SetReachableSymbolProperties(v []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems) {
	o.ReachableSymbolProperties = v
}

// GetRequiresTransitiveEnrichment returns the RequiresTransitiveEnrichment field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetRequiresTransitiveEnrichment() bool {
	if o == nil || o.RequiresTransitiveEnrichment == nil {
		var ret bool
		return ret
	}
	return *o.RequiresTransitiveEnrichment
}

// GetRequiresTransitiveEnrichmentOk returns a tuple with the RequiresTransitiveEnrichment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetRequiresTransitiveEnrichmentOk() (*bool, bool) {
	if o == nil || o.RequiresTransitiveEnrichment == nil {
		return nil, false
	}
	return o.RequiresTransitiveEnrichment, true
}

// HasRequiresTransitiveEnrichment returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasRequiresTransitiveEnrichment() bool {
	return o != nil && o.RequiresTransitiveEnrichment != nil
}

// SetRequiresTransitiveEnrichment gets a reference to the given bool and assigns it to the RequiresTransitiveEnrichment field.
func (o *ScaRequestDataAttributesDependenciesItems) SetRequiresTransitiveEnrichment(v bool) {
	o.RequiresTransitiveEnrichment = &v
}

// GetTargetFrameworks returns the TargetFrameworks field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetTargetFrameworks() []string {
	if o == nil || o.TargetFrameworks == nil {
		var ret []string
		return ret
	}
	return o.TargetFrameworks
}

// GetTargetFrameworksOk returns a tuple with the TargetFrameworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetTargetFrameworksOk() (*[]string, bool) {
	if o == nil || o.TargetFrameworks == nil {
		return nil, false
	}
	return &o.TargetFrameworks, true
}

// HasTargetFrameworks returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasTargetFrameworks() bool {
	return o != nil && o.TargetFrameworks != nil
}

// SetTargetFrameworks gets a reference to the given []string and assigns it to the TargetFrameworks field.
func (o *ScaRequestDataAttributesDependenciesItems) SetTargetFrameworks(v []string) {
	o.TargetFrameworks = v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScaRequestDataAttributesDependenciesItems) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasVersion() bool {
	return o != nil && o.Version.IsSet()
}

// SetVersion gets a reference to the given datadog.NullableString and assigns it to the Version field.
func (o *ScaRequestDataAttributesDependenciesItems) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItems) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil.
func (o *ScaRequestDataAttributesDependenciesItems) UnsetVersion() {
	o.Version.Unset()
}

// GetVersionConstraint returns the VersionConstraint field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersionConstraint() bool {
	if o == nil || o.VersionConstraint == nil {
		var ret bool
		return ret
	}
	return *o.VersionConstraint
}

// GetVersionConstraintOk returns a tuple with the VersionConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersionConstraintOk() (*bool, bool) {
	if o == nil || o.VersionConstraint == nil {
		return nil, false
	}
	return o.VersionConstraint, true
}

// HasVersionConstraint returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasVersionConstraint() bool {
	return o != nil && o.VersionConstraint != nil
}

// SetVersionConstraint gets a reference to the given bool and assigns it to the VersionConstraint field.
func (o *ScaRequestDataAttributesDependenciesItems) SetVersionConstraint(v bool) {
	o.VersionConstraint = &v
}

// GetVersionRange returns the VersionRange field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersionRange() string {
	if o == nil || o.VersionRange == nil {
		var ret string
		return ret
	}
	return *o.VersionRange
}

// GetVersionRangeOk returns a tuple with the VersionRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersionRangeOk() (*string, bool) {
	if o == nil || o.VersionRange == nil {
		return nil, false
	}
	return o.VersionRange, true
}

// HasVersionRange returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasVersionRange() bool {
	return o != nil && o.VersionRange != nil
}

// SetVersionRange gets a reference to the given string and assigns it to the VersionRange field.
func (o *ScaRequestDataAttributesDependenciesItems) SetVersionRange(v string) {
	o.VersionRange = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Exclusions != nil {
		toSerialize["exclusions"] = o.Exclusions
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.IsDev != nil {
		toSerialize["is_dev"] = o.IsDev
	}
	if o.IsDirect.IsSet() {
		toSerialize["is_direct"] = o.IsDirect.Get()
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Opaque != nil {
		toSerialize["opaque"] = o.Opaque
	}
	if o.PackageManager != nil {
		toSerialize["package_manager"] = o.PackageManager
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	if o.ReachableSymbolProperties != nil {
		toSerialize["reachable_symbol_properties"] = o.ReachableSymbolProperties
	}
	if o.RequiresTransitiveEnrichment != nil {
		toSerialize["requires_transitive_enrichment"] = o.RequiresTransitiveEnrichment
	}
	if o.TargetFrameworks != nil {
		toSerialize["target_frameworks"] = o.TargetFrameworks
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.VersionConstraint != nil {
		toSerialize["version_constraint"] = o.VersionConstraint
	}
	if o.VersionRange != nil {
		toSerialize["version_range"] = o.VersionRange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Exclusions                   []string                                                                  `json:"exclusions,omitempty"`
		Group                        datadog.NullableString                                                    `json:"group,omitempty"`
		IsDev                        *bool                                                                     `json:"is_dev,omitempty"`
		IsDirect                     datadog.NullableBool                                                      `json:"is_direct,omitempty"`
		Language                     *string                                                                   `json:"language,omitempty"`
		Locations                    []ScaRequestDataAttributesDependenciesItemsLocationsItems                 `json:"locations,omitempty"`
		Name                         *string                                                                   `json:"name,omitempty"`
		Opaque                       *bool                                                                     `json:"opaque,omitempty"`
		PackageManager               *string                                                                   `json:"package_manager,omitempty"`
		Purl                         *string                                                                   `json:"purl,omitempty"`
		ReachableSymbolProperties    []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems `json:"reachable_symbol_properties,omitempty"`
		RequiresTransitiveEnrichment *bool                                                                     `json:"requires_transitive_enrichment,omitempty"`
		TargetFrameworks             []string                                                                  `json:"target_frameworks,omitempty"`
		Version                      datadog.NullableString                                                    `json:"version,omitempty"`
		VersionConstraint            *bool                                                                     `json:"version_constraint,omitempty"`
		VersionRange                 *string                                                                   `json:"version_range,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclusions", "group", "is_dev", "is_direct", "language", "locations", "name", "opaque", "package_manager", "purl", "reachable_symbol_properties", "requires_transitive_enrichment", "target_frameworks", "version", "version_constraint", "version_range"})
	} else {
		return err
	}
	o.Exclusions = all.Exclusions
	o.Group = all.Group
	o.IsDev = all.IsDev
	o.IsDirect = all.IsDirect
	o.Language = all.Language
	o.Locations = all.Locations
	o.Name = all.Name
	o.Opaque = all.Opaque
	o.PackageManager = all.PackageManager
	o.Purl = all.Purl
	o.ReachableSymbolProperties = all.ReachableSymbolProperties
	o.RequiresTransitiveEnrichment = all.RequiresTransitiveEnrichment
	o.TargetFrameworks = all.TargetFrameworks
	o.Version = all.Version
	o.VersionConstraint = all.VersionConstraint
	o.VersionRange = all.VersionRange

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
