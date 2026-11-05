// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureDeleteRequestDataAttributes Revisions and confirmation used to delete the override.
type AiGatewayCaptureDeleteRequestDataAttributes struct {
	// Opaque revision of the organization default and all overrides. Writes require the revision returned by a current read or preview.
	ConfigurationRevision string `json:"configuration_revision"`
	// Set to `true` to confirm a change that can enable full capture, including a deletion that causes a `full_capture` override or the organization default to apply.
	ConfirmFullCapture *bool `json:"confirm_full_capture,omitempty"`
	// Capture mode for matching traffic.
	Mode *AiGatewayCaptureDeleteRequestDataAttributesMode `json:"mode,omitempty"`
	// Display name of the override. Names need not be unique; identify overrides by `id`.
	Name *string `json:"name,omitempty"`
	// Proposed operation for preview. Write endpoints infer it from the HTTP method.
	Operation *AiGatewayCaptureDeleteRequestDataAttributesOperation `json:"operation,omitempty"`
	// Current override revision, required for replacement or deletion.
	Revision int64 `json:"revision"`
	// All selected filters combine with AND, with one value per filter. The `model` and `provider` tags apply to each actual direct or routed attempt, including fallback destinations. Routes are not selectable filters. At least one filter is required.
	// If any matching override uses `metadata_only`, that mode applies. Otherwise, if any matching override uses `full_capture`, that mode applies. The organization default applies only when no override matches. Overrides have no priority or specificity ordering.
	Scope *AiGatewayCaptureScope `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureDeleteRequestDataAttributes instantiates a new AiGatewayCaptureDeleteRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureDeleteRequestDataAttributes(configurationRevision string, revision int64) *AiGatewayCaptureDeleteRequestDataAttributes {
	this := AiGatewayCaptureDeleteRequestDataAttributes{}
	this.ConfigurationRevision = configurationRevision
	this.Revision = revision
	return &this
}

// NewAiGatewayCaptureDeleteRequestDataAttributesWithDefaults instantiates a new AiGatewayCaptureDeleteRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureDeleteRequestDataAttributesWithDefaults() *AiGatewayCaptureDeleteRequestDataAttributes {
	this := AiGatewayCaptureDeleteRequestDataAttributes{}
	return &this
}

// GetConfigurationRevision returns the ConfigurationRevision field value.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetConfigurationRevision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigurationRevision
}

// GetConfigurationRevisionOk returns a tuple with the ConfigurationRevision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetConfigurationRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationRevision, true
}

// SetConfigurationRevision sets field value.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetConfigurationRevision(v string) {
	o.ConfigurationRevision = v
}

// GetConfirmFullCapture returns the ConfirmFullCapture field value if set, zero value otherwise.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetConfirmFullCapture() bool {
	if o == nil || o.ConfirmFullCapture == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmFullCapture
}

// GetConfirmFullCaptureOk returns a tuple with the ConfirmFullCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetConfirmFullCaptureOk() (*bool, bool) {
	if o == nil || o.ConfirmFullCapture == nil {
		return nil, false
	}
	return o.ConfirmFullCapture, true
}

// HasConfirmFullCapture returns a boolean if a field has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) HasConfirmFullCapture() bool {
	return o != nil && o.ConfirmFullCapture != nil
}

// SetConfirmFullCapture gets a reference to the given bool and assigns it to the ConfirmFullCapture field.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetConfirmFullCapture(v bool) {
	o.ConfirmFullCapture = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetMode() AiGatewayCaptureDeleteRequestDataAttributesMode {
	if o == nil || o.Mode == nil {
		var ret AiGatewayCaptureDeleteRequestDataAttributesMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetModeOk() (*AiGatewayCaptureDeleteRequestDataAttributesMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given AiGatewayCaptureDeleteRequestDataAttributesMode and assigns it to the Mode field.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetMode(v AiGatewayCaptureDeleteRequestDataAttributesMode) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetOperation() AiGatewayCaptureDeleteRequestDataAttributesOperation {
	if o == nil || o.Operation == nil {
		var ret AiGatewayCaptureDeleteRequestDataAttributesOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetOperationOk() (*AiGatewayCaptureDeleteRequestDataAttributesOperation, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given AiGatewayCaptureDeleteRequestDataAttributesOperation and assigns it to the Operation field.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetOperation(v AiGatewayCaptureDeleteRequestDataAttributesOperation) {
	o.Operation = &v
}

// GetRevision returns the Revision field value.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetRevision() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetRevisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetRevision(v int64) {
	o.Revision = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetScope() AiGatewayCaptureScope {
	if o == nil || o.Scope == nil {
		var ret AiGatewayCaptureScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) GetScopeOk() (*AiGatewayCaptureScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiGatewayCaptureScope and assigns it to the Scope field.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) SetScope(v AiGatewayCaptureScope) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureDeleteRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configuration_revision"] = o.ConfigurationRevision
	if o.ConfirmFullCapture != nil {
		toSerialize["confirm_full_capture"] = o.ConfirmFullCapture
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	toSerialize["revision"] = o.Revision
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayCaptureDeleteRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigurationRevision *string                                               `json:"configuration_revision"`
		ConfirmFullCapture    *bool                                                 `json:"confirm_full_capture,omitempty"`
		Mode                  *AiGatewayCaptureDeleteRequestDataAttributesMode      `json:"mode,omitempty"`
		Name                  *string                                               `json:"name,omitempty"`
		Operation             *AiGatewayCaptureDeleteRequestDataAttributesOperation `json:"operation,omitempty"`
		Revision              *int64                                                `json:"revision"`
		Scope                 *AiGatewayCaptureScope                                `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigurationRevision == nil {
		return fmt.Errorf("required field configuration_revision missing")
	}
	if all.Revision == nil {
		return fmt.Errorf("required field revision missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configuration_revision", "confirm_full_capture", "mode", "name", "operation", "revision", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConfigurationRevision = *all.ConfigurationRevision
	o.ConfirmFullCapture = all.ConfirmFullCapture
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.Name = all.Name
	if all.Operation != nil && !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = all.Operation
	}
	o.Revision = *all.Revision
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
