// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayCaptureOverrideResponseDataAttributes Saved name, capture mode, request filters, and revision of the override.
type AiGatewayCaptureOverrideResponseDataAttributes struct {
	// Items Metadata-only capture adds, each listed once. An item requires the items whose information it already reveals: `tool_output_size`, `tool_names` and `conversation_shape` require `tool_usage`, and `input_size_by_role` requires `tool_output_size`. A set that leaves out a required item is rejected. When several Metadata-only overrides match a request, it gets only the items all of them select.
	MetadataCapture []AiGatewayMetadataCaptureItem `json:"metadata_capture,omitempty"`
	// Capture mode for matching traffic.
	Mode AiGatewayCaptureOverrideResponseDataAttributesMode `json:"mode"`
	// Display name of the override. Names need not be unique; identify overrides by `id`.
	Name string `json:"name"`
	// Current override revision, required for replacement or deletion.
	Revision int64 `json:"revision"`
	// Visible request filters. Provider and model tags are withheld when the caller cannot view them. An empty object or empty tags can represent a fully restricted scope; scope_restricted identifies redaction.
	Scope AiGatewayCaptureResponseScope `json:"scope"`
	// Whether provider or model scope metadata was withheld. Omit unchanged scope on replacement to preserve it.
	ScopeRestricted bool `json:"scope_restricted"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayCaptureOverrideResponseDataAttributes instantiates a new AiGatewayCaptureOverrideResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayCaptureOverrideResponseDataAttributes(mode AiGatewayCaptureOverrideResponseDataAttributesMode, name string, revision int64, scope AiGatewayCaptureResponseScope, scopeRestricted bool) *AiGatewayCaptureOverrideResponseDataAttributes {
	this := AiGatewayCaptureOverrideResponseDataAttributes{}
	this.Mode = mode
	this.Name = name
	this.Revision = revision
	this.Scope = scope
	this.ScopeRestricted = scopeRestricted
	return &this
}

// NewAiGatewayCaptureOverrideResponseDataAttributesWithDefaults instantiates a new AiGatewayCaptureOverrideResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayCaptureOverrideResponseDataAttributesWithDefaults() *AiGatewayCaptureOverrideResponseDataAttributes {
	this := AiGatewayCaptureOverrideResponseDataAttributes{}
	return &this
}

// GetMetadataCapture returns the MetadataCapture field value if set, zero value otherwise.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetMetadataCapture() []AiGatewayMetadataCaptureItem {
	if o == nil || o.MetadataCapture == nil {
		var ret []AiGatewayMetadataCaptureItem
		return ret
	}
	return o.MetadataCapture
}

// GetMetadataCaptureOk returns a tuple with the MetadataCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetMetadataCaptureOk() (*[]AiGatewayMetadataCaptureItem, bool) {
	if o == nil || o.MetadataCapture == nil {
		return nil, false
	}
	return &o.MetadataCapture, true
}

// HasMetadataCapture returns a boolean if a field has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) HasMetadataCapture() bool {
	return o != nil && o.MetadataCapture != nil
}

// SetMetadataCapture gets a reference to the given []AiGatewayMetadataCaptureItem and assigns it to the MetadataCapture field.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) SetMetadataCapture(v []AiGatewayMetadataCaptureItem) {
	o.MetadataCapture = v
}

// GetMode returns the Mode field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetMode() AiGatewayCaptureOverrideResponseDataAttributesMode {
	if o == nil {
		var ret AiGatewayCaptureOverrideResponseDataAttributesMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetModeOk() (*AiGatewayCaptureOverrideResponseDataAttributesMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) SetMode(v AiGatewayCaptureOverrideResponseDataAttributesMode) {
	o.Mode = v
}

// GetName returns the Name field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) SetName(v string) {
	o.Name = v
}

// GetRevision returns the Revision field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetRevision() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetRevisionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) SetRevision(v int64) {
	o.Revision = v
}

// GetScope returns the Scope field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetScope() AiGatewayCaptureResponseScope {
	if o == nil {
		var ret AiGatewayCaptureResponseScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetScopeOk() (*AiGatewayCaptureResponseScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) SetScope(v AiGatewayCaptureResponseScope) {
	o.Scope = v
}

// GetScopeRestricted returns the ScopeRestricted field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetScopeRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ScopeRestricted
}

// GetScopeRestrictedOk returns a tuple with the ScopeRestricted field value
// and a boolean to check if the value has been set.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) GetScopeRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeRestricted, true
}

// SetScopeRestricted sets field value.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) SetScopeRestricted(v bool) {
	o.ScopeRestricted = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayCaptureOverrideResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MetadataCapture != nil {
		toSerialize["metadata_capture"] = o.MetadataCapture
	}
	toSerialize["mode"] = o.Mode
	toSerialize["name"] = o.Name
	toSerialize["revision"] = o.Revision
	toSerialize["scope"] = o.Scope
	toSerialize["scope_restricted"] = o.ScopeRestricted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayCaptureOverrideResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetadataCapture []AiGatewayMetadataCaptureItem                      `json:"metadata_capture,omitempty"`
		Mode            *AiGatewayCaptureOverrideResponseDataAttributesMode `json:"mode"`
		Name            *string                                             `json:"name"`
		Revision        *int64                                              `json:"revision"`
		Scope           *AiGatewayCaptureResponseScope                      `json:"scope"`
		ScopeRestricted *bool                                               `json:"scope_restricted"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Revision == nil {
		return fmt.Errorf("required field revision missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.ScopeRestricted == nil {
		return fmt.Errorf("required field scope_restricted missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metadata_capture", "mode", "name", "revision", "scope", "scope_restricted"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MetadataCapture = all.MetadataCapture
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.Name = *all.Name
	o.Revision = *all.Revision
	if all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = *all.Scope
	o.ScopeRestricted = *all.ScopeRestricted

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
