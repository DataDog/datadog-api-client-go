// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureCreateRequestDataAttributes Settings and configuration revision for creating the override.
type AiGatewayCaptureCreateRequestDataAttributes struct {
	// Opaque revision of the organization default and all overrides. Writes require the revision returned by a current read or preview.
	ConfigurationRevision string `json:"configuration_revision"`
	// Set to `true` to confirm a change that can enable full capture, including a deletion that causes a `full_capture` override or the organization default to apply.
	ConfirmFullCapture *bool `json:"confirm_full_capture,omitempty"`
	// Items Metadata-only capture adds, each listed once. An item requires the items whose information it already reveals: `tool_output_size`, `tool_names` and `conversation_shape` require `tool_usage`, and `input_size_by_role` requires `tool_output_size`. A set that leaves out a required item is rejected. When several Metadata-only overrides match a request, it gets only the items all of them select.
	MetadataCapture []AiGatewayMetadataCaptureItem `json:"metadata_capture,omitempty"`
	// Capture mode for matching traffic.
	Mode AiGatewayCaptureCreateRequestDataAttributesMode `json:"mode"`
	// Display name of the override. Names need not be unique; identify overrides by `id`.
	Name string `json:"name"`
	// Proposed operation for preview. Write endpoints infer it from the HTTP method.
	Operation *AiGatewayCaptureCreateRequestDataAttributesOperation `json:"operation,omitempty"`
	// Current override revision, required for replacement or deletion.
	Revision *int64 `json:"revision,omitempty"`
	// All selected filters combine with AND, with one value per filter. The `model` and `provider` tags apply to each actual direct or routed attempt, including fallback destinations. Routes are not selectable filters. At least one filter is required.
	// If any matching override uses `metadata_only`, that mode applies. Otherwise, if any matching override uses `full_capture`, that mode applies. The organization default applies only when no override matches. Overrides have no priority or specificity ordering.
	Scope AiGatewayCaptureScope `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureCreateRequestDataAttributes instantiates a new AiGatewayCaptureCreateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureCreateRequestDataAttributes(configurationRevision string, mode AiGatewayCaptureCreateRequestDataAttributesMode, name string, scope AiGatewayCaptureScope) *AiGatewayCaptureCreateRequestDataAttributes {
	this := AiGatewayCaptureCreateRequestDataAttributes{}
	this.ConfigurationRevision = configurationRevision
	this.Mode = mode
	this.Name = name
	this.Scope = scope
	return &this
}

// NewAiGatewayCaptureCreateRequestDataAttributesWithDefaults instantiates a new AiGatewayCaptureCreateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureCreateRequestDataAttributesWithDefaults() *AiGatewayCaptureCreateRequestDataAttributes {
	this := AiGatewayCaptureCreateRequestDataAttributes{}
	return &this
}

// GetConfigurationRevision returns the ConfigurationRevision field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetConfigurationRevision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigurationRevision
}

// GetConfigurationRevisionOk returns a tuple with the ConfigurationRevision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetConfigurationRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationRevision, true
}

// SetConfigurationRevision sets field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetConfigurationRevision(v string) {
	o.ConfigurationRevision = v
}

// GetConfirmFullCapture returns the ConfirmFullCapture field value if set, zero value otherwise.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetConfirmFullCapture() bool {
	if o == nil || o.ConfirmFullCapture == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmFullCapture
}

// GetConfirmFullCaptureOk returns a tuple with the ConfirmFullCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetConfirmFullCaptureOk() (*bool, bool) {
	if o == nil || o.ConfirmFullCapture == nil {
		return nil, false
	}
	return o.ConfirmFullCapture, true
}

// HasConfirmFullCapture returns a boolean if a field has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) HasConfirmFullCapture() bool {
	return o != nil && o.ConfirmFullCapture != nil
}

// SetConfirmFullCapture gets a reference to the given bool and assigns it to the ConfirmFullCapture field.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetConfirmFullCapture(v bool) {
	o.ConfirmFullCapture = &v
}

// GetMetadataCapture returns the MetadataCapture field value if set, zero value otherwise.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetMetadataCapture() []AiGatewayMetadataCaptureItem {
	if o == nil || o.MetadataCapture == nil {
		var ret []AiGatewayMetadataCaptureItem
		return ret
	}
	return o.MetadataCapture
}

// GetMetadataCaptureOk returns a tuple with the MetadataCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetMetadataCaptureOk() (*[]AiGatewayMetadataCaptureItem, bool) {
	if o == nil || o.MetadataCapture == nil {
		return nil, false
	}
	return &o.MetadataCapture, true
}

// HasMetadataCapture returns a boolean if a field has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) HasMetadataCapture() bool {
	return o != nil && o.MetadataCapture != nil
}

// SetMetadataCapture gets a reference to the given []AiGatewayMetadataCaptureItem and assigns it to the MetadataCapture field.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetMetadataCapture(v []AiGatewayMetadataCaptureItem) {
	o.MetadataCapture = v
}

// GetMode returns the Mode field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetMode() AiGatewayCaptureCreateRequestDataAttributesMode {
	if o == nil {
		var ret AiGatewayCaptureCreateRequestDataAttributesMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetModeOk() (*AiGatewayCaptureCreateRequestDataAttributesMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetMode(v AiGatewayCaptureCreateRequestDataAttributesMode) {
	o.Mode = v
}

// GetName returns the Name field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetOperation() AiGatewayCaptureCreateRequestDataAttributesOperation {
	if o == nil || o.Operation == nil {
		var ret AiGatewayCaptureCreateRequestDataAttributesOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetOperationOk() (*AiGatewayCaptureCreateRequestDataAttributesOperation, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given AiGatewayCaptureCreateRequestDataAttributesOperation and assigns it to the Operation field.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetOperation(v AiGatewayCaptureCreateRequestDataAttributesOperation) {
	o.Operation = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetRevision() int64 {
	if o == nil || o.Revision == nil {
		var ret int64
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetRevisionOk() (*int64, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) HasRevision() bool {
	return o != nil && o.Revision != nil
}

// SetRevision gets a reference to the given int64 and assigns it to the Revision field.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetRevision(v int64) {
	o.Revision = &v
}

// GetScope returns the Scope field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetScope() AiGatewayCaptureScope {
	if o == nil {
		var ret AiGatewayCaptureScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureCreateRequestDataAttributes) GetScopeOk() (*AiGatewayCaptureScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiGatewayCaptureCreateRequestDataAttributes) SetScope(v AiGatewayCaptureScope) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configuration_revision"] = o.ConfigurationRevision
	if o.ConfirmFullCapture != nil {
		toSerialize["confirm_full_capture"] = o.ConfirmFullCapture
	}
	if o.MetadataCapture != nil {
		toSerialize["metadata_capture"] = o.MetadataCapture
	}
	toSerialize["mode"] = o.Mode
	toSerialize["name"] = o.Name
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayCaptureCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigurationRevision *string                                               `json:"configuration_revision"`
		ConfirmFullCapture    *bool                                                 `json:"confirm_full_capture,omitempty"`
		MetadataCapture       []AiGatewayMetadataCaptureItem                        `json:"metadata_capture,omitempty"`
		Mode                  *AiGatewayCaptureCreateRequestDataAttributesMode      `json:"mode"`
		Name                  *string                                               `json:"name"`
		Operation             *AiGatewayCaptureCreateRequestDataAttributesOperation `json:"operation,omitempty"`
		Revision              *int64                                                `json:"revision,omitempty"`
		Scope                 *AiGatewayCaptureScope                                `json:"scope"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigurationRevision == nil {
		return fmt.Errorf("required field configuration_revision missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration_revision", "confirm_full_capture", "metadata_capture", "mode", "name", "operation", "revision", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfigurationRevision = *all.ConfigurationRevision
	o.ConfirmFullCapture = all.ConfirmFullCapture
	o.MetadataCapture = all.MetadataCapture
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.Name = *all.Name
	if all.Operation != nil && !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = all.Operation
	}
	o.Revision = all.Revision
	if all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = *all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
