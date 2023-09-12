// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationAttributes The custom destination attributes.
type CustomDestinationAttributes struct {
	// The max amount of bytes to buffer for the custom destination.
	BufferMaxBytes *int64 `json:"bufferMaxBytes,omitempty"`
	// The timeout amount in seconds to buffer for the custom destination.
	BufferTimeoutSeconds *int64 `json:"bufferTimeoutSeconds,omitempty"`
	// The compression method used for the custom destination.
	Compression *CustomDestinationCompressionType `json:"compression,omitempty"`
	// The enabled status of the custom destination.
	Enabled *bool `json:"enabled,omitempty"`
	// The archiving destination to fall back to.
	FallbackDestination *CustomDestinationFallbackDestination `json:"fallbackDestination,omitempty"`
	// The destination to forward events to.
	ForwarderDestination CustomDestinationForwarderDestination `json:"forwarderDestination"`
	// The retry duration in seconds for the custom destination.
	MaxRetryDurationSeconds *int64 `json:"maxRetryDurationSeconds,omitempty"`
	// The name of the custom destination.
	Name string `json:"name"`
	// The search query of the custom destination.
	Query *string `json:"query,omitempty"`
	// The version of the custom destination to prevent concurrent changes.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCustomDestinationAttributes instantiates a new CustomDestinationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationAttributes(forwarderDestination CustomDestinationForwarderDestination, name string, version int64) *CustomDestinationAttributes {
	this := CustomDestinationAttributes{}
	this.ForwarderDestination = forwarderDestination
	this.Name = name
	this.Version = version
	return &this
}

// NewCustomDestinationAttributesWithDefaults instantiates a new CustomDestinationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationAttributesWithDefaults() *CustomDestinationAttributes {
	this := CustomDestinationAttributes{}
	return &this
}

// GetBufferMaxBytes returns the BufferMaxBytes field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetBufferMaxBytes() int64 {
	if o == nil || o.BufferMaxBytes == nil {
		var ret int64
		return ret
	}
	return *o.BufferMaxBytes
}

// GetBufferMaxBytesOk returns a tuple with the BufferMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetBufferMaxBytesOk() (*int64, bool) {
	if o == nil || o.BufferMaxBytes == nil {
		return nil, false
	}
	return o.BufferMaxBytes, true
}

// HasBufferMaxBytes returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasBufferMaxBytes() bool {
	return o != nil && o.BufferMaxBytes != nil
}

// SetBufferMaxBytes gets a reference to the given int64 and assigns it to the BufferMaxBytes field.
func (o *CustomDestinationAttributes) SetBufferMaxBytes(v int64) {
	o.BufferMaxBytes = &v
}

// GetBufferTimeoutSeconds returns the BufferTimeoutSeconds field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetBufferTimeoutSeconds() int64 {
	if o == nil || o.BufferTimeoutSeconds == nil {
		var ret int64
		return ret
	}
	return *o.BufferTimeoutSeconds
}

// GetBufferTimeoutSecondsOk returns a tuple with the BufferTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetBufferTimeoutSecondsOk() (*int64, bool) {
	if o == nil || o.BufferTimeoutSeconds == nil {
		return nil, false
	}
	return o.BufferTimeoutSeconds, true
}

// HasBufferTimeoutSeconds returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasBufferTimeoutSeconds() bool {
	return o != nil && o.BufferTimeoutSeconds != nil
}

// SetBufferTimeoutSeconds gets a reference to the given int64 and assigns it to the BufferTimeoutSeconds field.
func (o *CustomDestinationAttributes) SetBufferTimeoutSeconds(v int64) {
	o.BufferTimeoutSeconds = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetCompression() CustomDestinationCompressionType {
	if o == nil || o.Compression == nil {
		var ret CustomDestinationCompressionType
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetCompressionOk() (*CustomDestinationCompressionType, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given CustomDestinationCompressionType and assigns it to the Compression field.
func (o *CustomDestinationAttributes) SetCompression(v CustomDestinationCompressionType) {
	o.Compression = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomDestinationAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFallbackDestination returns the FallbackDestination field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetFallbackDestination() CustomDestinationFallbackDestination {
	if o == nil || o.FallbackDestination == nil {
		var ret CustomDestinationFallbackDestination
		return ret
	}
	return *o.FallbackDestination
}

// GetFallbackDestinationOk returns a tuple with the FallbackDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetFallbackDestinationOk() (*CustomDestinationFallbackDestination, bool) {
	if o == nil || o.FallbackDestination == nil {
		return nil, false
	}
	return o.FallbackDestination, true
}

// HasFallbackDestination returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasFallbackDestination() bool {
	return o != nil && o.FallbackDestination != nil
}

// SetFallbackDestination gets a reference to the given CustomDestinationFallbackDestination and assigns it to the FallbackDestination field.
func (o *CustomDestinationAttributes) SetFallbackDestination(v CustomDestinationFallbackDestination) {
	o.FallbackDestination = &v
}

// GetForwarderDestination returns the ForwarderDestination field value.
func (o *CustomDestinationAttributes) GetForwarderDestination() CustomDestinationForwarderDestination {
	if o == nil {
		var ret CustomDestinationForwarderDestination
		return ret
	}
	return o.ForwarderDestination
}

// GetForwarderDestinationOk returns a tuple with the ForwarderDestination field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetForwarderDestinationOk() (*CustomDestinationForwarderDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForwarderDestination, true
}

// SetForwarderDestination sets field value.
func (o *CustomDestinationAttributes) SetForwarderDestination(v CustomDestinationForwarderDestination) {
	o.ForwarderDestination = v
}

// GetMaxRetryDurationSeconds returns the MaxRetryDurationSeconds field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetMaxRetryDurationSeconds() int64 {
	if o == nil || o.MaxRetryDurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.MaxRetryDurationSeconds
}

// GetMaxRetryDurationSecondsOk returns a tuple with the MaxRetryDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetMaxRetryDurationSecondsOk() (*int64, bool) {
	if o == nil || o.MaxRetryDurationSeconds == nil {
		return nil, false
	}
	return o.MaxRetryDurationSeconds, true
}

// HasMaxRetryDurationSeconds returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasMaxRetryDurationSeconds() bool {
	return o != nil && o.MaxRetryDurationSeconds != nil
}

// SetMaxRetryDurationSeconds gets a reference to the given int64 and assigns it to the MaxRetryDurationSeconds field.
func (o *CustomDestinationAttributes) SetMaxRetryDurationSeconds(v int64) {
	o.MaxRetryDurationSeconds = &v
}

// GetName returns the Name field value.
func (o *CustomDestinationAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomDestinationAttributes) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CustomDestinationAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CustomDestinationAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CustomDestinationAttributes) SetQuery(v string) {
	o.Query = &v
}

// GetVersion returns the Version field value.
func (o *CustomDestinationAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *CustomDestinationAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.BufferMaxBytes != nil {
		toSerialize["bufferMaxBytes"] = o.BufferMaxBytes
	}
	if o.BufferTimeoutSeconds != nil {
		toSerialize["bufferTimeoutSeconds"] = o.BufferTimeoutSeconds
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FallbackDestination != nil {
		toSerialize["fallbackDestination"] = o.FallbackDestination
	}
	toSerialize["forwarderDestination"] = o.ForwarderDestination
	if o.MaxRetryDurationSeconds != nil {
		toSerialize["maxRetryDurationSeconds"] = o.MaxRetryDurationSeconds
	}
	toSerialize["name"] = o.Name
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BufferMaxBytes          *int64                                 `json:"bufferMaxBytes,omitempty"`
		BufferTimeoutSeconds    *int64                                 `json:"bufferTimeoutSeconds,omitempty"`
		Compression             *CustomDestinationCompressionType      `json:"compression,omitempty"`
		Enabled                 *bool                                  `json:"enabled,omitempty"`
		FallbackDestination     *CustomDestinationFallbackDestination  `json:"fallbackDestination,omitempty"`
		ForwarderDestination    *CustomDestinationForwarderDestination `json:"forwarderDestination"`
		MaxRetryDurationSeconds *int64                                 `json:"maxRetryDurationSeconds,omitempty"`
		Name                    *string                                `json:"name"`
		Query                   *string                                `json:"query,omitempty"`
		Version                 *int64                                 `json:"version"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ForwarderDestination == nil {
		return fmt.Errorf("required field forwarderDestination missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bufferMaxBytes", "bufferTimeoutSeconds", "compression", "enabled", "fallbackDestination", "forwarderDestination", "maxRetryDurationSeconds", "name", "query", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BufferMaxBytes = all.BufferMaxBytes
	o.BufferTimeoutSeconds = all.BufferTimeoutSeconds
	if all.Compression != nil && !all.Compression.IsValid() {
		hasInvalidField = true
	} else {
		o.Compression = all.Compression
	}
	o.Enabled = all.Enabled
	o.FallbackDestination = all.FallbackDestination
	o.ForwarderDestination = *all.ForwarderDestination
	o.MaxRetryDurationSeconds = all.MaxRetryDurationSeconds
	o.Name = *all.Name
	o.Query = all.Query
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
