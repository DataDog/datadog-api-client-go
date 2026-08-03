// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetIntegrationDetailsV2 Detailed information about a single integration.
type FleetIntegrationDetailsV2 struct {
	// Type of data collected, such as metrics or logs.
	DataType *string `json:"data_type,omitempty"`
	// Error messages if the integration has issues.
	ErrorMessages []string `json:"error_messages,omitempty"`
	// Initialization configuration (YAML format).
	InitConfig *string `json:"init_config,omitempty"`
	// Instance-specific configuration (YAML format).
	InstanceConfig *string `json:"instance_config,omitempty"`
	// Whether this is a custom integration.
	IsCustomCheck *bool `json:"is_custom_check,omitempty"`
	// Whether this is a default integration instance.
	IsDefault *bool `json:"is_default,omitempty"`
	// Whether this integration configuration is an init config.
	IsInit *bool `json:"is_init,omitempty"`
	// Log collection configuration (YAML format).
	LogConfig *string `json:"log_config,omitempty"`
	// Name of the integration instance.
	Name *string `json:"name,omitempty"`
	// Number of pods running this integration. Absent from the response when the count is zero.
	PodCount *int64 `json:"pod_count,omitempty"`
	// Index in the configuration file.
	SourceIndex *int64 `json:"source_index,omitempty"`
	// Path to the configuration file.
	SourcePath *string `json:"source_path,omitempty"`
	// Integration type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetIntegrationDetailsV2 instantiates a new FleetIntegrationDetailsV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetIntegrationDetailsV2() *FleetIntegrationDetailsV2 {
	this := FleetIntegrationDetailsV2{}
	return &this
}

// NewFleetIntegrationDetailsV2WithDefaults instantiates a new FleetIntegrationDetailsV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetIntegrationDetailsV2WithDefaults() *FleetIntegrationDetailsV2 {
	this := FleetIntegrationDetailsV2{}
	return &this
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasDataType() bool {
	return o != nil && o.DataType != nil
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *FleetIntegrationDetailsV2) SetDataType(v string) {
	o.DataType = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages == nil {
		var ret []string
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil || o.ErrorMessages == nil {
		return nil, false
	}
	return &o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages != nil
}

// SetErrorMessages gets a reference to the given []string and assigns it to the ErrorMessages field.
func (o *FleetIntegrationDetailsV2) SetErrorMessages(v []string) {
	o.ErrorMessages = v
}

// GetInitConfig returns the InitConfig field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetInitConfig() string {
	if o == nil || o.InitConfig == nil {
		var ret string
		return ret
	}
	return *o.InitConfig
}

// GetInitConfigOk returns a tuple with the InitConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetInitConfigOk() (*string, bool) {
	if o == nil || o.InitConfig == nil {
		return nil, false
	}
	return o.InitConfig, true
}

// HasInitConfig returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasInitConfig() bool {
	return o != nil && o.InitConfig != nil
}

// SetInitConfig gets a reference to the given string and assigns it to the InitConfig field.
func (o *FleetIntegrationDetailsV2) SetInitConfig(v string) {
	o.InitConfig = &v
}

// GetInstanceConfig returns the InstanceConfig field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetInstanceConfig() string {
	if o == nil || o.InstanceConfig == nil {
		var ret string
		return ret
	}
	return *o.InstanceConfig
}

// GetInstanceConfigOk returns a tuple with the InstanceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetInstanceConfigOk() (*string, bool) {
	if o == nil || o.InstanceConfig == nil {
		return nil, false
	}
	return o.InstanceConfig, true
}

// HasInstanceConfig returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasInstanceConfig() bool {
	return o != nil && o.InstanceConfig != nil
}

// SetInstanceConfig gets a reference to the given string and assigns it to the InstanceConfig field.
func (o *FleetIntegrationDetailsV2) SetInstanceConfig(v string) {
	o.InstanceConfig = &v
}

// GetIsCustomCheck returns the IsCustomCheck field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetIsCustomCheck() bool {
	if o == nil || o.IsCustomCheck == nil {
		var ret bool
		return ret
	}
	return *o.IsCustomCheck
}

// GetIsCustomCheckOk returns a tuple with the IsCustomCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetIsCustomCheckOk() (*bool, bool) {
	if o == nil || o.IsCustomCheck == nil {
		return nil, false
	}
	return o.IsCustomCheck, true
}

// HasIsCustomCheck returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasIsCustomCheck() bool {
	return o != nil && o.IsCustomCheck != nil
}

// SetIsCustomCheck gets a reference to the given bool and assigns it to the IsCustomCheck field.
func (o *FleetIntegrationDetailsV2) SetIsCustomCheck(v bool) {
	o.IsCustomCheck = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *FleetIntegrationDetailsV2) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsInit returns the IsInit field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetIsInit() bool {
	if o == nil || o.IsInit == nil {
		var ret bool
		return ret
	}
	return *o.IsInit
}

// GetIsInitOk returns a tuple with the IsInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetIsInitOk() (*bool, bool) {
	if o == nil || o.IsInit == nil {
		return nil, false
	}
	return o.IsInit, true
}

// HasIsInit returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasIsInit() bool {
	return o != nil && o.IsInit != nil
}

// SetIsInit gets a reference to the given bool and assigns it to the IsInit field.
func (o *FleetIntegrationDetailsV2) SetIsInit(v bool) {
	o.IsInit = &v
}

// GetLogConfig returns the LogConfig field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetLogConfig() string {
	if o == nil || o.LogConfig == nil {
		var ret string
		return ret
	}
	return *o.LogConfig
}

// GetLogConfigOk returns a tuple with the LogConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetLogConfigOk() (*string, bool) {
	if o == nil || o.LogConfig == nil {
		return nil, false
	}
	return o.LogConfig, true
}

// HasLogConfig returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasLogConfig() bool {
	return o != nil && o.LogConfig != nil
}

// SetLogConfig gets a reference to the given string and assigns it to the LogConfig field.
func (o *FleetIntegrationDetailsV2) SetLogConfig(v string) {
	o.LogConfig = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FleetIntegrationDetailsV2) SetName(v string) {
	o.Name = &v
}

// GetPodCount returns the PodCount field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetPodCount() int64 {
	if o == nil || o.PodCount == nil {
		var ret int64
		return ret
	}
	return *o.PodCount
}

// GetPodCountOk returns a tuple with the PodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetPodCountOk() (*int64, bool) {
	if o == nil || o.PodCount == nil {
		return nil, false
	}
	return o.PodCount, true
}

// HasPodCount returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasPodCount() bool {
	return o != nil && o.PodCount != nil
}

// SetPodCount gets a reference to the given int64 and assigns it to the PodCount field.
func (o *FleetIntegrationDetailsV2) SetPodCount(v int64) {
	o.PodCount = &v
}

// GetSourceIndex returns the SourceIndex field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetSourceIndex() int64 {
	if o == nil || o.SourceIndex == nil {
		var ret int64
		return ret
	}
	return *o.SourceIndex
}

// GetSourceIndexOk returns a tuple with the SourceIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetSourceIndexOk() (*int64, bool) {
	if o == nil || o.SourceIndex == nil {
		return nil, false
	}
	return o.SourceIndex, true
}

// HasSourceIndex returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasSourceIndex() bool {
	return o != nil && o.SourceIndex != nil
}

// SetSourceIndex gets a reference to the given int64 and assigns it to the SourceIndex field.
func (o *FleetIntegrationDetailsV2) SetSourceIndex(v int64) {
	o.SourceIndex = &v
}

// GetSourcePath returns the SourcePath field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetSourcePath() string {
	if o == nil || o.SourcePath == nil {
		var ret string
		return ret
	}
	return *o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetSourcePathOk() (*string, bool) {
	if o == nil || o.SourcePath == nil {
		return nil, false
	}
	return o.SourcePath, true
}

// HasSourcePath returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasSourcePath() bool {
	return o != nil && o.SourcePath != nil
}

// SetSourcePath gets a reference to the given string and assigns it to the SourcePath field.
func (o *FleetIntegrationDetailsV2) SetSourcePath(v string) {
	o.SourcePath = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FleetIntegrationDetailsV2) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetIntegrationDetailsV2) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FleetIntegrationDetailsV2) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FleetIntegrationDetailsV2) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetIntegrationDetailsV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DataType != nil {
		toSerialize["data_type"] = o.DataType
	}
	if o.ErrorMessages != nil {
		toSerialize["error_messages"] = o.ErrorMessages
	}
	if o.InitConfig != nil {
		toSerialize["init_config"] = o.InitConfig
	}
	if o.InstanceConfig != nil {
		toSerialize["instance_config"] = o.InstanceConfig
	}
	if o.IsCustomCheck != nil {
		toSerialize["is_custom_check"] = o.IsCustomCheck
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.IsInit != nil {
		toSerialize["is_init"] = o.IsInit
	}
	if o.LogConfig != nil {
		toSerialize["log_config"] = o.LogConfig
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PodCount != nil {
		toSerialize["pod_count"] = o.PodCount
	}
	if o.SourceIndex != nil {
		toSerialize["source_index"] = o.SourceIndex
	}
	if o.SourcePath != nil {
		toSerialize["source_path"] = o.SourcePath
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetIntegrationDetailsV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataType       *string  `json:"data_type,omitempty"`
		ErrorMessages  []string `json:"error_messages,omitempty"`
		InitConfig     *string  `json:"init_config,omitempty"`
		InstanceConfig *string  `json:"instance_config,omitempty"`
		IsCustomCheck  *bool    `json:"is_custom_check,omitempty"`
		IsDefault      *bool    `json:"is_default,omitempty"`
		IsInit         *bool    `json:"is_init,omitempty"`
		LogConfig      *string  `json:"log_config,omitempty"`
		Name           *string  `json:"name,omitempty"`
		PodCount       *int64   `json:"pod_count,omitempty"`
		SourceIndex    *int64   `json:"source_index,omitempty"`
		SourcePath     *string  `json:"source_path,omitempty"`
		Type           *string  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_type", "error_messages", "init_config", "instance_config", "is_custom_check", "is_default", "is_init", "log_config", "name", "pod_count", "source_index", "source_path", "type"})
	} else {
		return err
	}
	o.DataType = all.DataType
	o.ErrorMessages = all.ErrorMessages
	o.InitConfig = all.InitConfig
	o.InstanceConfig = all.InstanceConfig
	o.IsCustomCheck = all.IsCustomCheck
	o.IsDefault = all.IsDefault
	o.IsInit = all.IsInit
	o.LogConfig = all.LogConfig
	o.Name = all.Name
	o.PodCount = all.PodCount
	o.SourceIndex = all.SourceIndex
	o.SourcePath = all.SourcePath
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
