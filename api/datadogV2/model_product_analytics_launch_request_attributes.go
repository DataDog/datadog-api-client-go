// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchRequestAttributes Attributes of a launch to create or replace.
type ProductAnalyticsLaunchRequestAttributes struct {
	// UUID of the RUM application that the launch belongs to.
	ApplicationId datadog.NullableString `json:"application_id,omitempty"`
	// Context of the feature being launched, used to build the tracking plan.
	Context string `json:"context"`
	// Description of the feature being launched.
	Description datadog.NullableString `json:"description,omitempty"`
	// UUID of the Datadog feature flag that gates the launch. Do not set it together with `feature_flag_key`.
	FeatureFlagId datadog.NullableString `json:"feature_flag_id,omitempty"`
	// Key of a feature flag managed outside Datadog, as evaluated in code. At most 256 characters, with no
	// leading or trailing spaces. Do not set it together with `feature_flag_id`. `null` clears it. On replace,
	// omitting it keeps the stored key unless `feature_flag_id` is set.
	FeatureFlagKey datadog.NullableString `json:"feature_flag_key,omitempty"`
	// KPIs that the launch is meant to move.
	Kpis []ProductAnalyticsLaunchKPI `json:"kpis,omitempty"`
	// Name of the launch.
	Name string `json:"name"`
	// Lifecycle phase of a launch. A new launch starts in `plan` or `setup`, and defaults to `plan`.
	// Allowed changes: `plan` to `setup`, `setup` to `ready`, `ready` to `setup` or `launched`, and `launched` to
	// `ready`.
	Phase *ProductAnalyticsLaunchPhase `json:"phase,omitempty"`
	// ID of the code repository that implements the feature.
	RepositoryId *string `json:"repository_id,omitempty"`
	// Steps of the launch. The tracking plan is required.
	Steps ProductAnalyticsLaunchStepsRequest `json:"steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchRequestAttributes instantiates a new ProductAnalyticsLaunchRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchRequestAttributes(context string, name string, steps ProductAnalyticsLaunchStepsRequest) *ProductAnalyticsLaunchRequestAttributes {
	this := ProductAnalyticsLaunchRequestAttributes{}
	this.Context = context
	this.Name = name
	this.Steps = steps
	return &this
}

// NewProductAnalyticsLaunchRequestAttributesWithDefaults instantiates a new ProductAnalyticsLaunchRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchRequestAttributesWithDefaults() *ProductAnalyticsLaunchRequestAttributes {
	this := ProductAnalyticsLaunchRequestAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchRequestAttributes) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchRequestAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId.IsSet()
}

// SetApplicationId gets a reference to the given datadog.NullableString and assigns it to the ApplicationId field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}

// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetContext returns the Context field value.
func (o *ProductAnalyticsLaunchRequestAttributes) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value.
func (o *ProductAnalyticsLaunchRequestAttributes) SetContext(v string) {
	o.Context = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchRequestAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetFeatureFlagId returns the FeatureFlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchRequestAttributes) GetFeatureFlagId() string {
	if o == nil || o.FeatureFlagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagId.Get()
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchRequestAttributes) GetFeatureFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlagId.Get(), o.FeatureFlagId.IsSet()
}

// HasFeatureFlagId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasFeatureFlagId() bool {
	return o != nil && o.FeatureFlagId.IsSet()
}

// SetFeatureFlagId gets a reference to the given datadog.NullableString and assigns it to the FeatureFlagId field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetFeatureFlagId(v string) {
	o.FeatureFlagId.Set(&v)
}

// SetFeatureFlagIdNil sets the value for FeatureFlagId to be an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) SetFeatureFlagIdNil() {
	o.FeatureFlagId.Set(nil)
}

// UnsetFeatureFlagId ensures that no value is present for FeatureFlagId, not even an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) UnsetFeatureFlagId() {
	o.FeatureFlagId.Unset()
}

// GetFeatureFlagKey returns the FeatureFlagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchRequestAttributes) GetFeatureFlagKey() string {
	if o == nil || o.FeatureFlagKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagKey.Get()
}

// GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchRequestAttributes) GetFeatureFlagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlagKey.Get(), o.FeatureFlagKey.IsSet()
}

// HasFeatureFlagKey returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasFeatureFlagKey() bool {
	return o != nil && o.FeatureFlagKey.IsSet()
}

// SetFeatureFlagKey gets a reference to the given datadog.NullableString and assigns it to the FeatureFlagKey field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetFeatureFlagKey(v string) {
	o.FeatureFlagKey.Set(&v)
}

// SetFeatureFlagKeyNil sets the value for FeatureFlagKey to be an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) SetFeatureFlagKeyNil() {
	o.FeatureFlagKey.Set(nil)
}

// UnsetFeatureFlagKey ensures that no value is present for FeatureFlagKey, not even an explicit nil.
func (o *ProductAnalyticsLaunchRequestAttributes) UnsetFeatureFlagKey() {
	o.FeatureFlagKey.Unset()
}

// GetKpis returns the Kpis field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchRequestAttributes) GetKpis() []ProductAnalyticsLaunchKPI {
	if o == nil || o.Kpis == nil {
		var ret []ProductAnalyticsLaunchKPI
		return ret
	}
	return o.Kpis
}

// GetKpisOk returns a tuple with the Kpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) GetKpisOk() (*[]ProductAnalyticsLaunchKPI, bool) {
	if o == nil || o.Kpis == nil {
		return nil, false
	}
	return &o.Kpis, true
}

// HasKpis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasKpis() bool {
	return o != nil && o.Kpis != nil
}

// SetKpis gets a reference to the given []ProductAnalyticsLaunchKPI and assigns it to the Kpis field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetKpis(v []ProductAnalyticsLaunchKPI) {
	o.Kpis = v
}

// GetName returns the Name field value.
func (o *ProductAnalyticsLaunchRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsLaunchRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchRequestAttributes) GetPhase() ProductAnalyticsLaunchPhase {
	if o == nil || o.Phase == nil {
		var ret ProductAnalyticsLaunchPhase
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) GetPhaseOk() (*ProductAnalyticsLaunchPhase, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given ProductAnalyticsLaunchPhase and assigns it to the Phase field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetPhase(v ProductAnalyticsLaunchPhase) {
	o.Phase = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchRequestAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *ProductAnalyticsLaunchRequestAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetSteps returns the Steps field value.
func (o *ProductAnalyticsLaunchRequestAttributes) GetSteps() ProductAnalyticsLaunchStepsRequest {
	if o == nil {
		var ret ProductAnalyticsLaunchStepsRequest
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchRequestAttributes) GetStepsOk() (*ProductAnalyticsLaunchStepsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value.
func (o *ProductAnalyticsLaunchRequestAttributes) SetSteps(v ProductAnalyticsLaunchStepsRequest) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId.IsSet() {
		toSerialize["application_id"] = o.ApplicationId.Get()
	}
	toSerialize["context"] = o.Context
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
	toSerialize["name"] = o.Name
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	toSerialize["steps"] = o.Steps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId  datadog.NullableString              `json:"application_id,omitempty"`
		Context        *string                             `json:"context"`
		Description    datadog.NullableString              `json:"description,omitempty"`
		FeatureFlagId  datadog.NullableString              `json:"feature_flag_id,omitempty"`
		FeatureFlagKey datadog.NullableString              `json:"feature_flag_key,omitempty"`
		Kpis           []ProductAnalyticsLaunchKPI         `json:"kpis,omitempty"`
		Name           *string                             `json:"name"`
		Phase          *ProductAnalyticsLaunchPhase        `json:"phase,omitempty"`
		RepositoryId   *string                             `json:"repository_id,omitempty"`
		Steps          *ProductAnalyticsLaunchStepsRequest `json:"steps"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Context == nil {
		return fmt.Errorf("required field context missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Steps == nil {
		return fmt.Errorf("required field steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "context", "description", "feature_flag_id", "feature_flag_key", "kpis", "name", "phase", "repository_id", "steps"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	o.Context = *all.Context
	o.Description = all.Description
	o.FeatureFlagId = all.FeatureFlagId
	o.FeatureFlagKey = all.FeatureFlagKey
	o.Kpis = all.Kpis
	o.Name = *all.Name
	if all.Phase != nil && !all.Phase.IsValid() {
		hasInvalidField = true
	} else {
		o.Phase = all.Phase
	}
	o.RepositoryId = all.RepositoryId
	if all.Steps.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Steps = *all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
