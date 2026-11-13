// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateEnvironmentAttributes Attributes for updating an environment.
type UpdateEnvironmentAttributes struct {
	// Indicates whether this is a production environment. If omitted, resets to false. Include the existing value to preserve it.
	IsProduction *bool `json:"is_production,omitempty"`
	// The name of the environment. Required on every request and unique within the organization. Must contain 1 to 200 ASCII characters, start with a letter or digit, and contain only letters, digits, spaces, underscores, hyphens, and periods.
	Name *string `json:"name,omitempty"`
	// Controls whether feature flag evaluation events may include original targeting identifiers and full evaluation context. When true, SDKs may include this data. When false, SDKs hash targeting identifiers and omit evaluation context. If omitted, the existing value is preserved.
	ObserveFullEvaluationData *bool `json:"observe_full_evaluation_data,omitempty"`
	// Required list of at least one exact DD_ENV value associated with this environment. Values cannot be used by another environment in the same organization. Wildcards and search expressions are not supported.
	Queries []string `json:"queries,omitempty"`
	// Indicates whether feature flag changes require approval in this environment. If omitted, resets to false. Include the existing value to preserve it.
	RequireFeatureFlagApproval *bool `json:"require_feature_flag_approval,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateEnvironmentAttributes instantiates a new UpdateEnvironmentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateEnvironmentAttributes() *UpdateEnvironmentAttributes {
	this := UpdateEnvironmentAttributes{}
	return &this
}

// NewUpdateEnvironmentAttributesWithDefaults instantiates a new UpdateEnvironmentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateEnvironmentAttributesWithDefaults() *UpdateEnvironmentAttributes {
	this := UpdateEnvironmentAttributes{}
	return &this
}

// GetIsProduction returns the IsProduction field value if set, zero value otherwise.
func (o *UpdateEnvironmentAttributes) GetIsProduction() bool {
	if o == nil || o.IsProduction == nil {
		var ret bool
		return ret
	}
	return *o.IsProduction
}

// GetIsProductionOk returns a tuple with the IsProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentAttributes) GetIsProductionOk() (*bool, bool) {
	if o == nil || o.IsProduction == nil {
		return nil, false
	}
	return o.IsProduction, true
}

// HasIsProduction returns a boolean if a field has been set.
func (o *UpdateEnvironmentAttributes) HasIsProduction() bool {
	return o != nil && o.IsProduction != nil
}

// SetIsProduction gets a reference to the given bool and assigns it to the IsProduction field.
func (o *UpdateEnvironmentAttributes) SetIsProduction(v bool) {
	o.IsProduction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateEnvironmentAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateEnvironmentAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateEnvironmentAttributes) SetName(v string) {
	o.Name = &v
}

// GetObserveFullEvaluationData returns the ObserveFullEvaluationData field value if set, zero value otherwise.
func (o *UpdateEnvironmentAttributes) GetObserveFullEvaluationData() bool {
	if o == nil || o.ObserveFullEvaluationData == nil {
		var ret bool
		return ret
	}
	return *o.ObserveFullEvaluationData
}

// GetObserveFullEvaluationDataOk returns a tuple with the ObserveFullEvaluationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentAttributes) GetObserveFullEvaluationDataOk() (*bool, bool) {
	if o == nil || o.ObserveFullEvaluationData == nil {
		return nil, false
	}
	return o.ObserveFullEvaluationData, true
}

// HasObserveFullEvaluationData returns a boolean if a field has been set.
func (o *UpdateEnvironmentAttributes) HasObserveFullEvaluationData() bool {
	return o != nil && o.ObserveFullEvaluationData != nil
}

// SetObserveFullEvaluationData gets a reference to the given bool and assigns it to the ObserveFullEvaluationData field.
func (o *UpdateEnvironmentAttributes) SetObserveFullEvaluationData(v bool) {
	o.ObserveFullEvaluationData = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *UpdateEnvironmentAttributes) GetQueries() []string {
	if o == nil || o.Queries == nil {
		var ret []string
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentAttributes) GetQueriesOk() (*[]string, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *UpdateEnvironmentAttributes) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []string and assigns it to the Queries field.
func (o *UpdateEnvironmentAttributes) SetQueries(v []string) {
	o.Queries = v
}

// GetRequireFeatureFlagApproval returns the RequireFeatureFlagApproval field value if set, zero value otherwise.
func (o *UpdateEnvironmentAttributes) GetRequireFeatureFlagApproval() bool {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireFeatureFlagApproval
}

// GetRequireFeatureFlagApprovalOk returns a tuple with the RequireFeatureFlagApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentAttributes) GetRequireFeatureFlagApprovalOk() (*bool, bool) {
	if o == nil || o.RequireFeatureFlagApproval == nil {
		return nil, false
	}
	return o.RequireFeatureFlagApproval, true
}

// HasRequireFeatureFlagApproval returns a boolean if a field has been set.
func (o *UpdateEnvironmentAttributes) HasRequireFeatureFlagApproval() bool {
	return o != nil && o.RequireFeatureFlagApproval != nil
}

// SetRequireFeatureFlagApproval gets a reference to the given bool and assigns it to the RequireFeatureFlagApproval field.
func (o *UpdateEnvironmentAttributes) SetRequireFeatureFlagApproval(v bool) {
	o.RequireFeatureFlagApproval = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateEnvironmentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsProduction != nil {
		toSerialize["is_production"] = o.IsProduction
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObserveFullEvaluationData != nil {
		toSerialize["observe_full_evaluation_data"] = o.ObserveFullEvaluationData
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.RequireFeatureFlagApproval != nil {
		toSerialize["require_feature_flag_approval"] = o.RequireFeatureFlagApproval
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateEnvironmentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsProduction               *bool    `json:"is_production,omitempty"`
		Name                       *string  `json:"name,omitempty"`
		ObserveFullEvaluationData  *bool    `json:"observe_full_evaluation_data,omitempty"`
		Queries                    []string `json:"queries,omitempty"`
		RequireFeatureFlagApproval *bool    `json:"require_feature_flag_approval,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_production", "name", "observe_full_evaluation_data", "queries", "require_feature_flag_approval"})
	} else {
		return err
	}
	o.IsProduction = all.IsProduction
	o.Name = all.Name
	o.ObserveFullEvaluationData = all.ObserveFullEvaluationData
	o.Queries = all.Queries
	o.RequireFeatureFlagApproval = all.RequireFeatureFlagApproval

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
