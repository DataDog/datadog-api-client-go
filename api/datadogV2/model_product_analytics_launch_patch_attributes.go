// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchPatchAttributes Attributes to update on a launch. Fields that are not set are left unchanged, and `null` clears a nullable
// field.
// Setting a non-empty `feature_flag_id` clears the stored `feature_flag_key` when `feature_flag_key` is not set.
type ProductAnalyticsLaunchPatchAttributes struct {
	// UUID of the RUM application that the launch belongs to.
	ApplicationId datadog.NullableString `json:"application_id,omitempty"`
	// Latest inspection of the repository for the tracking plan instrumentation. A new analysis replaces the stored
	// one.
	// Events and properties that it does not list are not verified.
	CodeAnalysis *ProductAnalyticsLaunchCodeAnalysis `json:"code_analysis,omitempty"`
	// Context of the feature being launched.
	Context *string `json:"context,omitempty"`
	// Description of the feature being launched.
	Description datadog.NullableString `json:"description,omitempty"`
	// UUID of the Datadog feature flag that gates the launch.
	FeatureFlagId datadog.NullableString `json:"feature_flag_id,omitempty"`
	// Key of a feature flag managed outside Datadog. At most 256 characters, with no leading or trailing spaces.
	FeatureFlagKey datadog.NullableString `json:"feature_flag_key,omitempty"`
	// KPIs that the launch is meant to move. Replaces the stored list, and `null` clears it.
	Kpis []ProductAnalyticsLaunchKPI `json:"kpis,omitempty"`
	// Name of the launch.
	Name *string `json:"name,omitempty"`
	// Lifecycle phase of a launch. A new launch starts in `plan` or `setup`, and defaults to `plan`.
	// Allowed changes: `plan` to `setup`, `setup` to `ready`, `ready` to `setup` or `launched`, and `launched` to
	// `ready`.
	Phase *ProductAnalyticsLaunchPhase `json:"phase,omitempty"`
	// ID of the code repository that implements the feature.
	RepositoryId datadog.NullableString `json:"repository_id,omitempty"`
	// Steps to update. Steps that are not set are left unchanged.
	// A `tracking_plan` replaces the whole stored tracking plan, so include every event and property to keep.
	Steps *ProductAnalyticsLaunchPatchSteps `json:"steps,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchPatchAttributes instantiates a new ProductAnalyticsLaunchPatchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchPatchAttributes() *ProductAnalyticsLaunchPatchAttributes {
	this := ProductAnalyticsLaunchPatchAttributes{}
	return &this
}

// NewProductAnalyticsLaunchPatchAttributesWithDefaults instantiates a new ProductAnalyticsLaunchPatchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchPatchAttributesWithDefaults() *ProductAnalyticsLaunchPatchAttributes {
	this := ProductAnalyticsLaunchPatchAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchPatchAttributes) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchPatchAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId.IsSet()
}

// SetApplicationId gets a reference to the given datadog.NullableString and assigns it to the ApplicationId field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}

// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetCodeAnalysis returns the CodeAnalysis field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchAttributes) GetCodeAnalysis() ProductAnalyticsLaunchCodeAnalysis {
	if o == nil || o.CodeAnalysis == nil {
		var ret ProductAnalyticsLaunchCodeAnalysis
		return ret
	}
	return *o.CodeAnalysis
}

// GetCodeAnalysisOk returns a tuple with the CodeAnalysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) GetCodeAnalysisOk() (*ProductAnalyticsLaunchCodeAnalysis, bool) {
	if o == nil || o.CodeAnalysis == nil {
		return nil, false
	}
	return o.CodeAnalysis, true
}

// HasCodeAnalysis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasCodeAnalysis() bool {
	return o != nil && o.CodeAnalysis != nil
}

// SetCodeAnalysis gets a reference to the given ProductAnalyticsLaunchCodeAnalysis and assigns it to the CodeAnalysis field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetCodeAnalysis(v ProductAnalyticsLaunchCodeAnalysis) {
	o.CodeAnalysis = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchAttributes) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetContext(v string) {
	o.Context = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchPatchAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchPatchAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetFeatureFlagId returns the FeatureFlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchPatchAttributes) GetFeatureFlagId() string {
	if o == nil || o.FeatureFlagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagId.Get()
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchPatchAttributes) GetFeatureFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlagId.Get(), o.FeatureFlagId.IsSet()
}

// HasFeatureFlagId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasFeatureFlagId() bool {
	return o != nil && o.FeatureFlagId.IsSet()
}

// SetFeatureFlagId gets a reference to the given datadog.NullableString and assigns it to the FeatureFlagId field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetFeatureFlagId(v string) {
	o.FeatureFlagId.Set(&v)
}

// SetFeatureFlagIdNil sets the value for FeatureFlagId to be an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) SetFeatureFlagIdNil() {
	o.FeatureFlagId.Set(nil)
}

// UnsetFeatureFlagId ensures that no value is present for FeatureFlagId, not even an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) UnsetFeatureFlagId() {
	o.FeatureFlagId.Unset()
}

// GetFeatureFlagKey returns the FeatureFlagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchPatchAttributes) GetFeatureFlagKey() string {
	if o == nil || o.FeatureFlagKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagKey.Get()
}

// GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchPatchAttributes) GetFeatureFlagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlagKey.Get(), o.FeatureFlagKey.IsSet()
}

// HasFeatureFlagKey returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasFeatureFlagKey() bool {
	return o != nil && o.FeatureFlagKey.IsSet()
}

// SetFeatureFlagKey gets a reference to the given datadog.NullableString and assigns it to the FeatureFlagKey field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetFeatureFlagKey(v string) {
	o.FeatureFlagKey.Set(&v)
}

// SetFeatureFlagKeyNil sets the value for FeatureFlagKey to be an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) SetFeatureFlagKeyNil() {
	o.FeatureFlagKey.Set(nil)
}

// UnsetFeatureFlagKey ensures that no value is present for FeatureFlagKey, not even an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) UnsetFeatureFlagKey() {
	o.FeatureFlagKey.Unset()
}

// GetKpis returns the Kpis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchPatchAttributes) GetKpis() []ProductAnalyticsLaunchKPI {
	if o == nil {
		var ret []ProductAnalyticsLaunchKPI
		return ret
	}
	return o.Kpis
}

// GetKpisOk returns a tuple with the Kpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchPatchAttributes) GetKpisOk() (*[]ProductAnalyticsLaunchKPI, bool) {
	if o == nil || o.Kpis == nil {
		return nil, false
	}
	return &o.Kpis, true
}

// HasKpis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasKpis() bool {
	return o != nil && o.Kpis != nil
}

// SetKpis gets a reference to the given []ProductAnalyticsLaunchKPI and assigns it to the Kpis field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetKpis(v []ProductAnalyticsLaunchKPI) {
	o.Kpis = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetName(v string) {
	o.Name = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchAttributes) GetPhase() ProductAnalyticsLaunchPhase {
	if o == nil || o.Phase == nil {
		var ret ProductAnalyticsLaunchPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) GetPhaseOk() (*ProductAnalyticsLaunchPhase, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given ProductAnalyticsLaunchPhase and assigns it to the Phase field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetPhase(v ProductAnalyticsLaunchPhase) {
	o.Phase = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchPatchAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId.Get()
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchPatchAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryId.Get(), o.RepositoryId.IsSet()
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId.IsSet()
}

// SetRepositoryId gets a reference to the given datadog.NullableString and assigns it to the RepositoryId field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetRepositoryId(v string) {
	o.RepositoryId.Set(&v)
}

// SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) SetRepositoryIdNil() {
	o.RepositoryId.Set(nil)
}

// UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil.
func (o *ProductAnalyticsLaunchPatchAttributes) UnsetRepositoryId() {
	o.RepositoryId.Unset()
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchAttributes) GetSteps() ProductAnalyticsLaunchPatchSteps {
	if o == nil || o.Steps == nil {
		var ret ProductAnalyticsLaunchPatchSteps
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) GetStepsOk() (*ProductAnalyticsLaunchPatchSteps, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchAttributes) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given ProductAnalyticsLaunchPatchSteps and assigns it to the Steps field.
func (o *ProductAnalyticsLaunchPatchAttributes) SetSteps(v ProductAnalyticsLaunchPatchSteps) {
	o.Steps = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchPatchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId.IsSet() {
		toSerialize["application_id"] = o.ApplicationId.Get()
	}
	if o.CodeAnalysis != nil {
		toSerialize["code_analysis"] = o.CodeAnalysis
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.FeatureFlagId.IsSet() {
		toSerialize["feature_flag_id"] = o.FeatureFlagId.Get()
	}
	if o.FeatureFlagKey.IsSet() {
		toSerialize["feature_flag_key"] = o.FeatureFlagKey.Get()
	}
	if o.Kpis != nil {
		toSerialize["kpis"] = o.Kpis
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.RepositoryId.IsSet() {
		toSerialize["repository_id"] = o.RepositoryId.Get()
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchPatchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId  datadog.NullableString              `json:"application_id,omitempty"`
		CodeAnalysis   *ProductAnalyticsLaunchCodeAnalysis `json:"code_analysis,omitempty"`
		Context        *string                             `json:"context,omitempty"`
		Description    datadog.NullableString              `json:"description,omitempty"`
		FeatureFlagId  datadog.NullableString              `json:"feature_flag_id,omitempty"`
		FeatureFlagKey datadog.NullableString              `json:"feature_flag_key,omitempty"`
		Kpis           []ProductAnalyticsLaunchKPI         `json:"kpis,omitempty"`
		Name           *string                             `json:"name,omitempty"`
		Phase          *ProductAnalyticsLaunchPhase        `json:"phase,omitempty"`
		RepositoryId   datadog.NullableString              `json:"repository_id,omitempty"`
		Steps          *ProductAnalyticsLaunchPatchSteps   `json:"steps,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "code_analysis", "context", "description", "feature_flag_id", "feature_flag_key", "kpis", "name", "phase", "repository_id", "steps"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	if all.CodeAnalysis != nil && all.CodeAnalysis.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CodeAnalysis = all.CodeAnalysis
	o.Context = all.Context
	o.Description = all.Description
	o.FeatureFlagId = all.FeatureFlagId
	o.FeatureFlagKey = all.FeatureFlagKey
	o.Kpis = all.Kpis
	o.Name = all.Name
	if all.Phase != nil && !all.Phase.IsValid() {
		hasInvalidField = true
	} else {
		o.Phase = all.Phase
	}
	o.RepositoryId = all.RepositoryId
	if all.Steps != nil && all.Steps.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Steps = all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
