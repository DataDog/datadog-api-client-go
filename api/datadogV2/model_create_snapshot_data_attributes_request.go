// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateSnapshotDataAttributesRequest Attributes for snapshot creation.
type CreateSnapshotDataAttributesRequest struct {
	// Additional configuration options for snapshot creation.
	AdditionalConfig *CreateSnapshotAdditionalConfig `json:"additional_config,omitempty"`
	// End of the time window for the snapshot, in milliseconds since Unix epoch.
	End int64 `json:"end"`
	// The height of the rendered snapshot in pixels.
	Height *int64 `json:"height,omitempty"`
	// Whether the snapshot requires authentication to view. Authenticated snapshots are scoped to the creating organization.
	IsAuthenticated *bool `json:"is_authenticated,omitempty"`
	// Start of the time window for the snapshot, in milliseconds since Unix epoch.
	Start int64 `json:"start"`
	// The time-to-live for the snapshot. This value corresponds to storage lifecycle policies that automatically delete the snapshot after the specified period.
	Ttl *CreateSnapshotTTL `json:"ttl,omitempty"`
	// The widget definition to render as a snapshot. Must include a valid `type` field and non-empty `requests` array.
	WidgetDefinition map[string]interface{} `json:"widget_definition"`
	// The width of the rendered snapshot in pixels.
	Width *int64 `json:"width,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateSnapshotDataAttributesRequest instantiates a new CreateSnapshotDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateSnapshotDataAttributesRequest(end int64, start int64, widgetDefinition map[string]interface{}) *CreateSnapshotDataAttributesRequest {
	this := CreateSnapshotDataAttributesRequest{}
	this.End = end
	this.Start = start
	this.WidgetDefinition = widgetDefinition
	return &this
}

// NewCreateSnapshotDataAttributesRequestWithDefaults instantiates a new CreateSnapshotDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateSnapshotDataAttributesRequestWithDefaults() *CreateSnapshotDataAttributesRequest {
	this := CreateSnapshotDataAttributesRequest{}
	return &this
}

// GetAdditionalConfig returns the AdditionalConfig field value if set, zero value otherwise.
func (o *CreateSnapshotDataAttributesRequest) GetAdditionalConfig() CreateSnapshotAdditionalConfig {
	if o == nil || o.AdditionalConfig == nil {
		var ret CreateSnapshotAdditionalConfig
		return ret
	}
	return *o.AdditionalConfig
}

// GetAdditionalConfigOk returns a tuple with the AdditionalConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetAdditionalConfigOk() (*CreateSnapshotAdditionalConfig, bool) {
	if o == nil || o.AdditionalConfig == nil {
		return nil, false
	}
	return o.AdditionalConfig, true
}

// HasAdditionalConfig returns a boolean if a field has been set.
func (o *CreateSnapshotDataAttributesRequest) HasAdditionalConfig() bool {
	return o != nil && o.AdditionalConfig != nil
}

// SetAdditionalConfig gets a reference to the given CreateSnapshotAdditionalConfig and assigns it to the AdditionalConfig field.
func (o *CreateSnapshotDataAttributesRequest) SetAdditionalConfig(v CreateSnapshotAdditionalConfig) {
	o.AdditionalConfig = &v
}

// GetEnd returns the End field value.
func (o *CreateSnapshotDataAttributesRequest) GetEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *CreateSnapshotDataAttributesRequest) SetEnd(v int64) {
	o.End = v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CreateSnapshotDataAttributesRequest) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CreateSnapshotDataAttributesRequest) HasHeight() bool {
	return o != nil && o.Height != nil
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *CreateSnapshotDataAttributesRequest) SetHeight(v int64) {
	o.Height = &v
}

// GetIsAuthenticated returns the IsAuthenticated field value if set, zero value otherwise.
func (o *CreateSnapshotDataAttributesRequest) GetIsAuthenticated() bool {
	if o == nil || o.IsAuthenticated == nil {
		var ret bool
		return ret
	}
	return *o.IsAuthenticated
}

// GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetIsAuthenticatedOk() (*bool, bool) {
	if o == nil || o.IsAuthenticated == nil {
		return nil, false
	}
	return o.IsAuthenticated, true
}

// HasIsAuthenticated returns a boolean if a field has been set.
func (o *CreateSnapshotDataAttributesRequest) HasIsAuthenticated() bool {
	return o != nil && o.IsAuthenticated != nil
}

// SetIsAuthenticated gets a reference to the given bool and assigns it to the IsAuthenticated field.
func (o *CreateSnapshotDataAttributesRequest) SetIsAuthenticated(v bool) {
	o.IsAuthenticated = &v
}

// GetStart returns the Start field value.
func (o *CreateSnapshotDataAttributesRequest) GetStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *CreateSnapshotDataAttributesRequest) SetStart(v int64) {
	o.Start = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CreateSnapshotDataAttributesRequest) GetTtl() CreateSnapshotTTL {
	if o == nil || o.Ttl == nil {
		var ret CreateSnapshotTTL
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetTtlOk() (*CreateSnapshotTTL, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CreateSnapshotDataAttributesRequest) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given CreateSnapshotTTL and assigns it to the Ttl field.
func (o *CreateSnapshotDataAttributesRequest) SetTtl(v CreateSnapshotTTL) {
	o.Ttl = &v
}

// GetWidgetDefinition returns the WidgetDefinition field value.
func (o *CreateSnapshotDataAttributesRequest) GetWidgetDefinition() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.WidgetDefinition
}

// GetWidgetDefinitionOk returns a tuple with the WidgetDefinition field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetWidgetDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetDefinition, true
}

// SetWidgetDefinition sets field value.
func (o *CreateSnapshotDataAttributesRequest) SetWidgetDefinition(v map[string]interface{}) {
	o.WidgetDefinition = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CreateSnapshotDataAttributesRequest) GetWidth() int64 {
	if o == nil || o.Width == nil {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotDataAttributesRequest) GetWidthOk() (*int64, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CreateSnapshotDataAttributesRequest) HasWidth() bool {
	return o != nil && o.Width != nil
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *CreateSnapshotDataAttributesRequest) SetWidth(v int64) {
	o.Width = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateSnapshotDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalConfig != nil {
		toSerialize["additional_config"] = o.AdditionalConfig
	}
	toSerialize["end"] = o.End
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.IsAuthenticated != nil {
		toSerialize["is_authenticated"] = o.IsAuthenticated
	}
	toSerialize["start"] = o.Start
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	toSerialize["widget_definition"] = o.WidgetDefinition
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateSnapshotDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalConfig *CreateSnapshotAdditionalConfig `json:"additional_config,omitempty"`
		End              *int64                          `json:"end"`
		Height           *int64                          `json:"height,omitempty"`
		IsAuthenticated  *bool                           `json:"is_authenticated,omitempty"`
		Start            *int64                          `json:"start"`
		Ttl              *CreateSnapshotTTL              `json:"ttl,omitempty"`
		WidgetDefinition *map[string]interface{}         `json:"widget_definition"`
		Width            *int64                          `json:"width,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.WidgetDefinition == nil {
		return fmt.Errorf("required field widget_definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additional_config", "end", "height", "is_authenticated", "start", "ttl", "widget_definition", "width"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AdditionalConfig != nil && all.AdditionalConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AdditionalConfig = all.AdditionalConfig
	o.End = *all.End
	o.Height = all.Height
	o.IsAuthenticated = all.IsAuthenticated
	o.Start = *all.Start
	if all.Ttl != nil && !all.Ttl.IsValid() {
		hasInvalidField = true
	} else {
		o.Ttl = all.Ttl
	}
	o.WidgetDefinition = *all.WidgetDefinition
	o.Width = all.Width

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
