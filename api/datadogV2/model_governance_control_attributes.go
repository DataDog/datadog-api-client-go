// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlAttributes The attributes of a governance control.
type GovernanceControlAttributes struct {
	// The number of active detections for the control.
	ActiveDetectionsCount int64 `json:"active_detections_count"`
	// The value driver the control is grouped under, such as `security` or `cost`.
	Category string `json:"category"`
	// The time the control configuration was created.
	CreatedAt time.Time `json:"created_at"`
	// The UUID of the user who created the control configuration.
	CreatedBy string `json:"created_by"`
	// A human-readable description of what the control detects.
	Description string `json:"description"`
	// A free-form map of parameter names to their configured values.
	DetectionParameters map[string]interface{} `json:"detection_parameters"`
	// The insight slugs associated with the control.
	Insights []string `json:"insights"`
	// The time of the most recent detection for the control. `null` when there are no detections.
	LastDetectionAt datadog.NullableTime `json:"last_detection_at"`
	// The number of mitigated detections for the control.
	MitigatedDetectionsCount int64 `json:"mitigated_detections_count"`
	// A free-form map of parameter names to their configured values.
	MitigationParameters map[string]interface{} `json:"mitigation_parameters"`
	// The configured mitigation type for the control. Empty when not configured.
	MitigationType string `json:"mitigation_type"`
	// The mitigations available for a control.
	Mitigations []GovernanceControlMitigationDefinition `json:"mitigations"`
	// Human-readable name of the control.
	Name string `json:"name"`
	// The priority of the control, such as `High`.
	Priority string `json:"priority"`
	// The product the control belongs to.
	Product string `json:"product"`
	// The type of resource the control evaluates.
	ResourceType string `json:"resource_type"`
	// The human-readable name of the resource type.
	ResourceTypeDisplayName string `json:"resource_type_display_name"`
	// An array of parameter definitions.
	SupportedDetectionParameters []GovernanceControlParameterDefinition `json:"supported_detection_parameters"`
	// The control type, such as `Proactive` or `Detection`.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlAttributes instantiates a new GovernanceControlAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlAttributes(activeDetectionsCount int64, category string, createdAt time.Time, createdBy string, description string, detectionParameters map[string]interface{}, insights []string, lastDetectionAt datadog.NullableTime, mitigatedDetectionsCount int64, mitigationParameters map[string]interface{}, mitigationType string, mitigations []GovernanceControlMitigationDefinition, name string, priority string, product string, resourceType string, resourceTypeDisplayName string, supportedDetectionParameters []GovernanceControlParameterDefinition, typeVar string) *GovernanceControlAttributes {
	this := GovernanceControlAttributes{}
	this.ActiveDetectionsCount = activeDetectionsCount
	this.Category = category
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.DetectionParameters = detectionParameters
	this.Insights = insights
	this.LastDetectionAt = lastDetectionAt
	this.MitigatedDetectionsCount = mitigatedDetectionsCount
	this.MitigationParameters = mitigationParameters
	this.MitigationType = mitigationType
	this.Mitigations = mitigations
	this.Name = name
	this.Priority = priority
	this.Product = product
	this.ResourceType = resourceType
	this.ResourceTypeDisplayName = resourceTypeDisplayName
	this.SupportedDetectionParameters = supportedDetectionParameters
	this.Type = typeVar
	return &this
}

// NewGovernanceControlAttributesWithDefaults instantiates a new GovernanceControlAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlAttributesWithDefaults() *GovernanceControlAttributes {
	this := GovernanceControlAttributes{}
	return &this
}

// GetActiveDetectionsCount returns the ActiveDetectionsCount field value.
func (o *GovernanceControlAttributes) GetActiveDetectionsCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ActiveDetectionsCount
}

// GetActiveDetectionsCountOk returns a tuple with the ActiveDetectionsCount field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetActiveDetectionsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveDetectionsCount, true
}

// SetActiveDetectionsCount sets field value.
func (o *GovernanceControlAttributes) SetActiveDetectionsCount(v int64) {
	o.ActiveDetectionsCount = v
}

// GetCategory returns the Category field value.
func (o *GovernanceControlAttributes) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *GovernanceControlAttributes) SetCategory(v string) {
	o.Category = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *GovernanceControlAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *GovernanceControlAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *GovernanceControlAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *GovernanceControlAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value.
func (o *GovernanceControlAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GovernanceControlAttributes) SetDescription(v string) {
	o.Description = v
}

// GetDetectionParameters returns the DetectionParameters field value.
func (o *GovernanceControlAttributes) GetDetectionParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DetectionParameters
}

// GetDetectionParametersOk returns a tuple with the DetectionParameters field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetDetectionParametersOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionParameters, true
}

// SetDetectionParameters sets field value.
func (o *GovernanceControlAttributes) SetDetectionParameters(v map[string]interface{}) {
	o.DetectionParameters = v
}

// GetInsights returns the Insights field value.
func (o *GovernanceControlAttributes) GetInsights() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Insights
}

// GetInsightsOk returns a tuple with the Insights field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetInsightsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Insights, true
}

// SetInsights sets field value.
func (o *GovernanceControlAttributes) SetInsights(v []string) {
	o.Insights = v
}

// GetLastDetectionAt returns the LastDetectionAt field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *GovernanceControlAttributes) GetLastDetectionAt() time.Time {
	if o == nil || o.LastDetectionAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDetectionAt.Get()
}

// GetLastDetectionAtOk returns a tuple with the LastDetectionAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *GovernanceControlAttributes) GetLastDetectionAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastDetectionAt.Get(), o.LastDetectionAt.IsSet()
}

// SetLastDetectionAt sets field value.
func (o *GovernanceControlAttributes) SetLastDetectionAt(v time.Time) {
	o.LastDetectionAt.Set(&v)
}

// GetMitigatedDetectionsCount returns the MitigatedDetectionsCount field value.
func (o *GovernanceControlAttributes) GetMitigatedDetectionsCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MitigatedDetectionsCount
}

// GetMitigatedDetectionsCountOk returns a tuple with the MitigatedDetectionsCount field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetMitigatedDetectionsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MitigatedDetectionsCount, true
}

// SetMitigatedDetectionsCount sets field value.
func (o *GovernanceControlAttributes) SetMitigatedDetectionsCount(v int64) {
	o.MitigatedDetectionsCount = v
}

// GetMitigationParameters returns the MitigationParameters field value.
func (o *GovernanceControlAttributes) GetMitigationParameters() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MitigationParameters
}

// GetMitigationParametersOk returns a tuple with the MitigationParameters field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetMitigationParametersOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MitigationParameters, true
}

// SetMitigationParameters sets field value.
func (o *GovernanceControlAttributes) SetMitigationParameters(v map[string]interface{}) {
	o.MitigationParameters = v
}

// GetMitigationType returns the MitigationType field value.
func (o *GovernanceControlAttributes) GetMitigationType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MitigationType
}

// GetMitigationTypeOk returns a tuple with the MitigationType field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetMitigationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MitigationType, true
}

// SetMitigationType sets field value.
func (o *GovernanceControlAttributes) SetMitigationType(v string) {
	o.MitigationType = v
}

// GetMitigations returns the Mitigations field value.
func (o *GovernanceControlAttributes) GetMitigations() []GovernanceControlMitigationDefinition {
	if o == nil {
		var ret []GovernanceControlMitigationDefinition
		return ret
	}
	return o.Mitigations
}

// GetMitigationsOk returns a tuple with the Mitigations field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetMitigationsOk() (*[]GovernanceControlMitigationDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mitigations, true
}

// SetMitigations sets field value.
func (o *GovernanceControlAttributes) SetMitigations(v []GovernanceControlMitigationDefinition) {
	o.Mitigations = v
}

// GetName returns the Name field value.
func (o *GovernanceControlAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GovernanceControlAttributes) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value.
func (o *GovernanceControlAttributes) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *GovernanceControlAttributes) SetPriority(v string) {
	o.Priority = v
}

// GetProduct returns the Product field value.
func (o *GovernanceControlAttributes) GetProduct() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value.
func (o *GovernanceControlAttributes) SetProduct(v string) {
	o.Product = v
}

// GetResourceType returns the ResourceType field value.
func (o *GovernanceControlAttributes) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *GovernanceControlAttributes) SetResourceType(v string) {
	o.ResourceType = v
}

// GetResourceTypeDisplayName returns the ResourceTypeDisplayName field value.
func (o *GovernanceControlAttributes) GetResourceTypeDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceTypeDisplayName
}

// GetResourceTypeDisplayNameOk returns a tuple with the ResourceTypeDisplayName field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetResourceTypeDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceTypeDisplayName, true
}

// SetResourceTypeDisplayName sets field value.
func (o *GovernanceControlAttributes) SetResourceTypeDisplayName(v string) {
	o.ResourceTypeDisplayName = v
}

// GetSupportedDetectionParameters returns the SupportedDetectionParameters field value.
func (o *GovernanceControlAttributes) GetSupportedDetectionParameters() []GovernanceControlParameterDefinition {
	if o == nil {
		var ret []GovernanceControlParameterDefinition
		return ret
	}
	return o.SupportedDetectionParameters
}

// GetSupportedDetectionParametersOk returns a tuple with the SupportedDetectionParameters field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetSupportedDetectionParametersOk() (*[]GovernanceControlParameterDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedDetectionParameters, true
}

// SetSupportedDetectionParameters sets field value.
func (o *GovernanceControlAttributes) SetSupportedDetectionParameters(v []GovernanceControlParameterDefinition) {
	o.SupportedDetectionParameters = v
}

// GetType returns the Type field value.
func (o *GovernanceControlAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GovernanceControlAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GovernanceControlAttributes) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["active_detections_count"] = o.ActiveDetectionsCount
	toSerialize["category"] = o.Category
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["description"] = o.Description
	toSerialize["detection_parameters"] = o.DetectionParameters
	toSerialize["insights"] = o.Insights
	toSerialize["last_detection_at"] = o.LastDetectionAt.Get()
	toSerialize["mitigated_detections_count"] = o.MitigatedDetectionsCount
	toSerialize["mitigation_parameters"] = o.MitigationParameters
	toSerialize["mitigation_type"] = o.MitigationType
	toSerialize["mitigations"] = o.Mitigations
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	toSerialize["product"] = o.Product
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["resource_type_display_name"] = o.ResourceTypeDisplayName
	toSerialize["supported_detection_parameters"] = o.SupportedDetectionParameters
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActiveDetectionsCount        *int64                                   `json:"active_detections_count"`
		Category                     *string                                  `json:"category"`
		CreatedAt                    *time.Time                               `json:"created_at"`
		CreatedBy                    *string                                  `json:"created_by"`
		Description                  *string                                  `json:"description"`
		DetectionParameters          *map[string]interface{}                  `json:"detection_parameters"`
		Insights                     *[]string                                `json:"insights"`
		LastDetectionAt              datadog.NullableTime                     `json:"last_detection_at"`
		MitigatedDetectionsCount     *int64                                   `json:"mitigated_detections_count"`
		MitigationParameters         *map[string]interface{}                  `json:"mitigation_parameters"`
		MitigationType               *string                                  `json:"mitigation_type"`
		Mitigations                  *[]GovernanceControlMitigationDefinition `json:"mitigations"`
		Name                         *string                                  `json:"name"`
		Priority                     *string                                  `json:"priority"`
		Product                      *string                                  `json:"product"`
		ResourceType                 *string                                  `json:"resource_type"`
		ResourceTypeDisplayName      *string                                  `json:"resource_type_display_name"`
		SupportedDetectionParameters *[]GovernanceControlParameterDefinition  `json:"supported_detection_parameters"`
		Type                         *string                                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActiveDetectionsCount == nil {
		return fmt.Errorf("required field active_detections_count missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DetectionParameters == nil {
		return fmt.Errorf("required field detection_parameters missing")
	}
	if all.Insights == nil {
		return fmt.Errorf("required field insights missing")
	}
	if !all.LastDetectionAt.IsSet() {
		return fmt.Errorf("required field last_detection_at missing")
	}
	if all.MitigatedDetectionsCount == nil {
		return fmt.Errorf("required field mitigated_detections_count missing")
	}
	if all.MitigationParameters == nil {
		return fmt.Errorf("required field mitigation_parameters missing")
	}
	if all.MitigationType == nil {
		return fmt.Errorf("required field mitigation_type missing")
	}
	if all.Mitigations == nil {
		return fmt.Errorf("required field mitigations missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.Product == nil {
		return fmt.Errorf("required field product missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.ResourceTypeDisplayName == nil {
		return fmt.Errorf("required field resource_type_display_name missing")
	}
	if all.SupportedDetectionParameters == nil {
		return fmt.Errorf("required field supported_detection_parameters missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"active_detections_count", "category", "created_at", "created_by", "description", "detection_parameters", "insights", "last_detection_at", "mitigated_detections_count", "mitigation_parameters", "mitigation_type", "mitigations", "name", "priority", "product", "resource_type", "resource_type_display_name", "supported_detection_parameters", "type"})
	} else {
		return err
	}
	o.ActiveDetectionsCount = *all.ActiveDetectionsCount
	o.Category = *all.Category
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Description = *all.Description
	o.DetectionParameters = *all.DetectionParameters
	o.Insights = *all.Insights
	o.LastDetectionAt = all.LastDetectionAt
	o.MitigatedDetectionsCount = *all.MitigatedDetectionsCount
	o.MitigationParameters = *all.MitigationParameters
	o.MitigationType = *all.MitigationType
	o.Mitigations = *all.Mitigations
	o.Name = *all.Name
	o.Priority = *all.Priority
	o.Product = *all.Product
	o.ResourceType = *all.ResourceType
	o.ResourceTypeDisplayName = *all.ResourceTypeDisplayName
	o.SupportedDetectionParameters = *all.SupportedDetectionParameters
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
