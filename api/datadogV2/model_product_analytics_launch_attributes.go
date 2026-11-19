// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchAttributes Attributes of a launch.
type ProductAnalyticsLaunchAttributes struct {
	// UUID of the RUM application that the launch belongs to.
	ApplicationId datadog.NullableString `json:"application_id,omitempty"`
	// Context of the feature being launched.
	Context datadog.NullableString `json:"context,omitempty"`
	// Time the launch was created.
	CreatedAt time.Time `json:"created_at"`
	// UUID of the user who created the launch.
	CreatedBy string `json:"created_by"`
	// Description of the feature being launched.
	Description datadog.NullableString `json:"description,omitempty"`
	// UUID of the Datadog feature flag that gates the launch.
	FeatureFlagId datadog.NullableString `json:"feature_flag_id,omitempty"`
	// Key of a feature flag managed outside Datadog.
	FeatureFlagKey datadog.NullableString `json:"feature_flag_key,omitempty"`
	// KPIs that the launch is meant to move.
	Kpis []ProductAnalyticsLaunchKPI `json:"kpis,omitempty"`
	// Name of the launch.
	Name string `json:"name"`
	// ID of the organization that owns the launch.
	OrgId int64 `json:"org_id"`
	// Lifecycle phase of a launch. A new launch starts in `plan` or `setup`, and defaults to `plan`.
	// Allowed changes: `plan` to `setup`, `setup` to `ready`, `ready` to `setup` or `launched`, and `launched` to
	// `ready`.
	Phase ProductAnalyticsLaunchPhase `json:"phase"`
	// ID of the code repository that implements the feature.
	RepositoryId *string `json:"repository_id,omitempty"`
	// Steps of the launch.
	Steps ProductAnalyticsLaunchSteps `json:"steps"`
	// Time the launch was last updated.
	UpdatedAt datadog.NullableTime `json:"updated_at,omitempty"`
	// UUID of the user who last updated the launch.
	UpdatedBy datadog.NullableString `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchAttributes instantiates a new ProductAnalyticsLaunchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchAttributes(createdAt time.Time, createdBy string, name string, orgId int64, phase ProductAnalyticsLaunchPhase, steps ProductAnalyticsLaunchSteps) *ProductAnalyticsLaunchAttributes {
	this := ProductAnalyticsLaunchAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Name = name
	this.OrgId = orgId
	this.Phase = phase
	this.Steps = steps
	return &this
}

// NewProductAnalyticsLaunchAttributesWithDefaults instantiates a new ProductAnalyticsLaunchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchAttributesWithDefaults() *ProductAnalyticsLaunchAttributes {
	this := ProductAnalyticsLaunchAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetApplicationId() string {
	if o == nil || o.ApplicationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId.Get()
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationId.Get(), o.ApplicationId.IsSet()
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId.IsSet()
}

// SetApplicationId gets a reference to the given datadog.NullableString and assigns it to the ApplicationId field.
func (o *ProductAnalyticsLaunchAttributes) SetApplicationId(v string) {
	o.ApplicationId.Set(&v)
}

// SetApplicationIdNil sets the value for ApplicationId to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetApplicationIdNil() {
	o.ApplicationId.Set(nil)
}

// UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetApplicationId() {
	o.ApplicationId.Unset()
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetContext() string {
	if o == nil || o.Context.Get() == nil {
		var ret string
		return ret
	}
	return *o.Context.Get()
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context.Get(), o.Context.IsSet()
}

// HasContext returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasContext() bool {
	return o != nil && o.Context.IsSet()
}

// SetContext gets a reference to the given datadog.NullableString and assigns it to the Context field.
func (o *ProductAnalyticsLaunchAttributes) SetContext(v string) {
	o.Context.Set(&v)
}

// SetContextNil sets the value for Context to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetContextNil() {
	o.Context.Set(nil)
}

// UnsetContext ensures that no value is present for Context, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetContext() {
	o.Context.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ProductAnalyticsLaunchAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ProductAnalyticsLaunchAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *ProductAnalyticsLaunchAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *ProductAnalyticsLaunchAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ProductAnalyticsLaunchAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetFeatureFlagId returns the FeatureFlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetFeatureFlagId() string {
	if o == nil || o.FeatureFlagId.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagId.Get()
}

// GetFeatureFlagIdOk returns a tuple with the FeatureFlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetFeatureFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlagId.Get(), o.FeatureFlagId.IsSet()
}

// HasFeatureFlagId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasFeatureFlagId() bool {
	return o != nil && o.FeatureFlagId.IsSet()
}

// SetFeatureFlagId gets a reference to the given datadog.NullableString and assigns it to the FeatureFlagId field.
func (o *ProductAnalyticsLaunchAttributes) SetFeatureFlagId(v string) {
	o.FeatureFlagId.Set(&v)
}

// SetFeatureFlagIdNil sets the value for FeatureFlagId to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetFeatureFlagIdNil() {
	o.FeatureFlagId.Set(nil)
}

// UnsetFeatureFlagId ensures that no value is present for FeatureFlagId, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetFeatureFlagId() {
	o.FeatureFlagId.Unset()
}

// GetFeatureFlagKey returns the FeatureFlagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetFeatureFlagKey() string {
	if o == nil || o.FeatureFlagKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.FeatureFlagKey.Get()
}

// GetFeatureFlagKeyOk returns a tuple with the FeatureFlagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetFeatureFlagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureFlagKey.Get(), o.FeatureFlagKey.IsSet()
}

// HasFeatureFlagKey returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasFeatureFlagKey() bool {
	return o != nil && o.FeatureFlagKey.IsSet()
}

// SetFeatureFlagKey gets a reference to the given datadog.NullableString and assigns it to the FeatureFlagKey field.
func (o *ProductAnalyticsLaunchAttributes) SetFeatureFlagKey(v string) {
	o.FeatureFlagKey.Set(&v)
}

// SetFeatureFlagKeyNil sets the value for FeatureFlagKey to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetFeatureFlagKeyNil() {
	o.FeatureFlagKey.Set(nil)
}

// UnsetFeatureFlagKey ensures that no value is present for FeatureFlagKey, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetFeatureFlagKey() {
	o.FeatureFlagKey.Unset()
}

// GetKpis returns the Kpis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetKpis() []ProductAnalyticsLaunchKPI {
	if o == nil {
		var ret []ProductAnalyticsLaunchKPI
		return ret
	}
	return o.Kpis
}

// GetKpisOk returns a tuple with the Kpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetKpisOk() (*[]ProductAnalyticsLaunchKPI, bool) {
	if o == nil || o.Kpis == nil {
		return nil, false
	}
	return &o.Kpis, true
}

// HasKpis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasKpis() bool {
	return o != nil && o.Kpis != nil
}

// SetKpis gets a reference to the given []ProductAnalyticsLaunchKPI and assigns it to the Kpis field.
func (o *ProductAnalyticsLaunchAttributes) SetKpis(v []ProductAnalyticsLaunchKPI) {
	o.Kpis = v
}

// GetName returns the Name field value.
func (o *ProductAnalyticsLaunchAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsLaunchAttributes) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value.
func (o *ProductAnalyticsLaunchAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *ProductAnalyticsLaunchAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetPhase returns the Phase field value.
func (o *ProductAnalyticsLaunchAttributes) GetPhase() ProductAnalyticsLaunchPhase {
	if o == nil {
		var ret ProductAnalyticsLaunchPhase
		return ret
	}
	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetPhaseOk() (*ProductAnalyticsLaunchPhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value.
func (o *ProductAnalyticsLaunchAttributes) SetPhase(v ProductAnalyticsLaunchPhase) {
	o.Phase = v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchAttributes) GetRepositoryId() string {
	if o == nil || o.RepositoryId == nil {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetRepositoryIdOk() (*string, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasRepositoryId() bool {
	return o != nil && o.RepositoryId != nil
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *ProductAnalyticsLaunchAttributes) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetSteps returns the Steps field value.
func (o *ProductAnalyticsLaunchAttributes) GetSteps() ProductAnalyticsLaunchSteps {
	if o == nil {
		var ret ProductAnalyticsLaunchSteps
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchAttributes) GetStepsOk() (*ProductAnalyticsLaunchSteps, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value.
func (o *ProductAnalyticsLaunchAttributes) SetSteps(v ProductAnalyticsLaunchSteps) {
	o.Steps = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given datadog.NullableTime and assigns it to the UpdatedAt field.
func (o *ProductAnalyticsLaunchAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAnalyticsLaunchAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ProductAnalyticsLaunchAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy.IsSet()
}

// SetUpdatedBy gets a reference to the given datadog.NullableString and assigns it to the UpdatedBy field.
func (o *ProductAnalyticsLaunchAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil.
func (o *ProductAnalyticsLaunchAttributes) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId.IsSet() {
		toSerialize["application_id"] = o.ApplicationId.Get()
	}
	if o.Context.IsSet() {
		toSerialize["context"] = o.Context.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
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
	toSerialize["org_id"] = o.OrgId
	toSerialize["phase"] = o.Phase
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	toSerialize["steps"] = o.Steps
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updated_by"] = o.UpdatedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId  datadog.NullableString       `json:"application_id,omitempty"`
		Context        datadog.NullableString       `json:"context,omitempty"`
		CreatedAt      *time.Time                   `json:"created_at"`
		CreatedBy      *string                      `json:"created_by"`
		Description    datadog.NullableString       `json:"description,omitempty"`
		FeatureFlagId  datadog.NullableString       `json:"feature_flag_id,omitempty"`
		FeatureFlagKey datadog.NullableString       `json:"feature_flag_key,omitempty"`
		Kpis           []ProductAnalyticsLaunchKPI  `json:"kpis,omitempty"`
		Name           *string                      `json:"name"`
		OrgId          *int64                       `json:"org_id"`
		Phase          *ProductAnalyticsLaunchPhase `json:"phase"`
		RepositoryId   *string                      `json:"repository_id,omitempty"`
		Steps          *ProductAnalyticsLaunchSteps `json:"steps"`
		UpdatedAt      datadog.NullableTime         `json:"updated_at,omitempty"`
		UpdatedBy      datadog.NullableString       `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Phase == nil {
		return fmt.Errorf("required field phase missing")
	}
	if all.Steps == nil {
		return fmt.Errorf("required field steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "context", "created_at", "created_by", "description", "feature_flag_id", "feature_flag_key", "kpis", "name", "org_id", "phase", "repository_id", "steps", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	o.Context = all.Context
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Description = all.Description
	o.FeatureFlagId = all.FeatureFlagId
	o.FeatureFlagKey = all.FeatureFlagKey
	o.Kpis = all.Kpis
	o.Name = *all.Name
	o.OrgId = *all.OrgId
	if !all.Phase.IsValid() {
		hasInvalidField = true
	} else {
		o.Phase = *all.Phase
	}
	o.RepositoryId = all.RepositoryId
	if all.Steps.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Steps = *all.Steps
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
