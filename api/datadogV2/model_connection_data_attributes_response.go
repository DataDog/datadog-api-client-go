// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionDataAttributesResponse Attributes of an action connection.
type ConnectionDataAttributesResponse struct {
	// Indicates if the acting user can resolve the connection.
	ActingUserCanResolve bool `json:"acting_user_can_resolve"`
	// The creation timestamp of the connection.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Information about the user who created the resource.
	CreatedBy *CreatedBy `json:"created_by,omitempty"`
	// The description of the connection.
	Description *string `json:"description,omitempty"`
	// External secrets manager configuration.
	ExternalSecretsManager ExternalSecretsManager `json:"external_secrets_manager"`
	// Integration configuration details.
	Integration interface{} `json:"integration"`
	// Indicates if the connection is marked as favorite.
	IsFavorite bool `json:"is_favorite"`
	// The last updated timestamp of the connection.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// The name of the connection.
	Name string `json:"name"`
	// Information about the private actions runner.
	PrivateActionsRunner *PrivateActionsRunner `json:"private_actions_runner,omitempty"`
	// The ID of the runner associated with the connection.
	RunnerId *string `json:"runner_id,omitempty"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConnectionDataAttributesResponse instantiates a new ConnectionDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConnectionDataAttributesResponse(actingUserCanResolve bool, externalSecretsManager ExternalSecretsManager, integration interface{}, isFavorite bool, name string) *ConnectionDataAttributesResponse {
	this := ConnectionDataAttributesResponse{}
	this.ActingUserCanResolve = actingUserCanResolve
	this.ExternalSecretsManager = externalSecretsManager
	this.Integration = integration
	this.IsFavorite = isFavorite
	this.Name = name
	return &this
}

// NewConnectionDataAttributesResponseWithDefaults instantiates a new ConnectionDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConnectionDataAttributesResponseWithDefaults() *ConnectionDataAttributesResponse {
	this := ConnectionDataAttributesResponse{}
	return &this
}

// GetActingUserCanResolve returns the ActingUserCanResolve field value.
func (o *ConnectionDataAttributesResponse) GetActingUserCanResolve() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ActingUserCanResolve
}

// GetActingUserCanResolveOk returns a tuple with the ActingUserCanResolve field value
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetActingUserCanResolveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActingUserCanResolve, true
}

// SetActingUserCanResolve sets field value.
func (o *ConnectionDataAttributesResponse) SetActingUserCanResolve(v bool) {
	o.ActingUserCanResolve = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ConnectionDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetCreatedBy() CreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret CreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetCreatedByOk() (*CreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given CreatedBy and assigns it to the CreatedBy field.
func (o *ConnectionDataAttributesResponse) SetCreatedBy(v CreatedBy) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectionDataAttributesResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExternalSecretsManager returns the ExternalSecretsManager field value.
func (o *ConnectionDataAttributesResponse) GetExternalSecretsManager() ExternalSecretsManager {
	if o == nil {
		var ret ExternalSecretsManager
		return ret
	}
	return o.ExternalSecretsManager
}

// GetExternalSecretsManagerOk returns a tuple with the ExternalSecretsManager field value
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetExternalSecretsManagerOk() (*ExternalSecretsManager, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalSecretsManager, true
}

// SetExternalSecretsManager sets field value.
func (o *ConnectionDataAttributesResponse) SetExternalSecretsManager(v ExternalSecretsManager) {
	o.ExternalSecretsManager = v
}

// GetIntegration returns the Integration field value.
func (o *ConnectionDataAttributesResponse) GetIntegration() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetIntegrationOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *ConnectionDataAttributesResponse) SetIntegration(v interface{}) {
	o.Integration = v
}

// GetIsFavorite returns the IsFavorite field value.
func (o *ConnectionDataAttributesResponse) GetIsFavorite() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorite, true
}

// SetIsFavorite sets field value.
func (o *ConnectionDataAttributesResponse) SetIsFavorite(v bool) {
	o.IsFavorite = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ConnectionDataAttributesResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value.
func (o *ConnectionDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ConnectionDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetPrivateActionsRunner returns the PrivateActionsRunner field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetPrivateActionsRunner() PrivateActionsRunner {
	if o == nil || o.PrivateActionsRunner == nil {
		var ret PrivateActionsRunner
		return ret
	}
	return *o.PrivateActionsRunner
}

// GetPrivateActionsRunnerOk returns a tuple with the PrivateActionsRunner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetPrivateActionsRunnerOk() (*PrivateActionsRunner, bool) {
	if o == nil || o.PrivateActionsRunner == nil {
		return nil, false
	}
	return o.PrivateActionsRunner, true
}

// HasPrivateActionsRunner returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasPrivateActionsRunner() bool {
	return o != nil && o.PrivateActionsRunner != nil
}

// SetPrivateActionsRunner gets a reference to the given PrivateActionsRunner and assigns it to the PrivateActionsRunner field.
func (o *ConnectionDataAttributesResponse) SetPrivateActionsRunner(v PrivateActionsRunner) {
	o.PrivateActionsRunner = &v
}

// GetRunnerId returns the RunnerId field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetRunnerId() string {
	if o == nil || o.RunnerId == nil {
		var ret string
		return ret
	}
	return *o.RunnerId
}

// GetRunnerIdOk returns a tuple with the RunnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetRunnerIdOk() (*string, bool) {
	if o == nil || o.RunnerId == nil {
		return nil, false
	}
	return o.RunnerId, true
}

// HasRunnerId returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasRunnerId() bool {
	return o != nil && o.RunnerId != nil
}

// SetRunnerId gets a reference to the given string and assigns it to the RunnerId field.
func (o *ConnectionDataAttributesResponse) SetRunnerId(v string) {
	o.RunnerId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConnectionDataAttributesResponse) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionDataAttributesResponse) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConnectionDataAttributesResponse) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConnectionDataAttributesResponse) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConnectionDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["acting_user_can_resolve"] = o.ActingUserCanResolve
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["external_secrets_manager"] = o.ExternalSecretsManager
	toSerialize["integration"] = o.Integration
	toSerialize["is_favorite"] = o.IsFavorite
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	if o.PrivateActionsRunner != nil {
		toSerialize["private_actions_runner"] = o.PrivateActionsRunner
	}
	if o.RunnerId != nil {
		toSerialize["runner_id"] = o.RunnerId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConnectionDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActingUserCanResolve   *bool                   `json:"acting_user_can_resolve"`
		CreatedAt              *time.Time              `json:"created_at,omitempty"`
		CreatedBy              *CreatedBy              `json:"created_by,omitempty"`
		Description            *string                 `json:"description,omitempty"`
		ExternalSecretsManager *ExternalSecretsManager `json:"external_secrets_manager"`
		Integration            *interface{}            `json:"integration"`
		IsFavorite             *bool                   `json:"is_favorite"`
		LastUpdated            *time.Time              `json:"last_updated,omitempty"`
		Name                   *string                 `json:"name"`
		PrivateActionsRunner   *PrivateActionsRunner   `json:"private_actions_runner,omitempty"`
		RunnerId               *string                 `json:"runner_id,omitempty"`
		Tags                   []string                `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActingUserCanResolve == nil {
		return fmt.Errorf("required field acting_user_can_resolve missing")
	}
	if all.ExternalSecretsManager == nil {
		return fmt.Errorf("required field external_secrets_manager missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.IsFavorite == nil {
		return fmt.Errorf("required field is_favorite missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"acting_user_can_resolve", "created_at", "created_by", "description", "external_secrets_manager", "integration", "is_favorite", "last_updated", "name", "private_actions_runner", "runner_id", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActingUserCanResolve = *all.ActingUserCanResolve
	o.CreatedAt = all.CreatedAt
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	o.ExternalSecretsManager = *all.ExternalSecretsManager
	o.Integration = *all.Integration
	o.IsFavorite = *all.IsFavorite
	o.LastUpdated = all.LastUpdated
	o.Name = *all.Name
	if all.PrivateActionsRunner != nil && all.PrivateActionsRunner.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PrivateActionsRunner = all.PrivateActionsRunner
	o.RunnerId = all.RunnerId
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
