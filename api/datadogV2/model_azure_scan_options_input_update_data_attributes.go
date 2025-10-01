// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureScanOptionsInputUpdateDataAttributes The definition of `AzureScanOptionsInputUpdateDataAttributes` object.
type AzureScanOptionsInputUpdateDataAttributes struct {
	// The `attributes` `vuln_containers_os`.
	VulnContainersOs *bool `json:"vuln_containers_os,omitempty"`
	// The `attributes` `vuln_host_os`.
	VulnHostOs *bool `json:"vuln_host_os,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureScanOptionsInputUpdateDataAttributes instantiates a new AzureScanOptionsInputUpdateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureScanOptionsInputUpdateDataAttributes() *AzureScanOptionsInputUpdateDataAttributes {
	this := AzureScanOptionsInputUpdateDataAttributes{}
	return &this
}

// NewAzureScanOptionsInputUpdateDataAttributesWithDefaults instantiates a new AzureScanOptionsInputUpdateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureScanOptionsInputUpdateDataAttributesWithDefaults() *AzureScanOptionsInputUpdateDataAttributes {
	this := AzureScanOptionsInputUpdateDataAttributes{}
	return &this
}

// GetVulnContainersOs returns the VulnContainersOs field value if set, zero value otherwise.
func (o *AzureScanOptionsInputUpdateDataAttributes) GetVulnContainersOs() bool {
	if o == nil || o.VulnContainersOs == nil {
		var ret bool
		return ret
	}
	return *o.VulnContainersOs
}

// GetVulnContainersOsOk returns a tuple with the VulnContainersOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureScanOptionsInputUpdateDataAttributes) GetVulnContainersOsOk() (*bool, bool) {
	if o == nil || o.VulnContainersOs == nil {
		return nil, false
	}
	return o.VulnContainersOs, true
}

// HasVulnContainersOs returns a boolean if a field has been set.
func (o *AzureScanOptionsInputUpdateDataAttributes) HasVulnContainersOs() bool {
	return o != nil && o.VulnContainersOs != nil
}

// SetVulnContainersOs gets a reference to the given bool and assigns it to the VulnContainersOs field.
func (o *AzureScanOptionsInputUpdateDataAttributes) SetVulnContainersOs(v bool) {
	o.VulnContainersOs = &v
}

// GetVulnHostOs returns the VulnHostOs field value if set, zero value otherwise.
func (o *AzureScanOptionsInputUpdateDataAttributes) GetVulnHostOs() bool {
	if o == nil || o.VulnHostOs == nil {
		var ret bool
		return ret
	}
	return *o.VulnHostOs
}

// GetVulnHostOsOk returns a tuple with the VulnHostOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureScanOptionsInputUpdateDataAttributes) GetVulnHostOsOk() (*bool, bool) {
	if o == nil || o.VulnHostOs == nil {
		return nil, false
	}
	return o.VulnHostOs, true
}

// HasVulnHostOs returns a boolean if a field has been set.
func (o *AzureScanOptionsInputUpdateDataAttributes) HasVulnHostOs() bool {
	return o != nil && o.VulnHostOs != nil
}

// SetVulnHostOs gets a reference to the given bool and assigns it to the VulnHostOs field.
func (o *AzureScanOptionsInputUpdateDataAttributes) SetVulnHostOs(v bool) {
	o.VulnHostOs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureScanOptionsInputUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.VulnContainersOs != nil {
		toSerialize["vuln_containers_os"] = o.VulnContainersOs
	}
	if o.VulnHostOs != nil {
		toSerialize["vuln_host_os"] = o.VulnHostOs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureScanOptionsInputUpdateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		VulnContainersOs *bool `json:"vuln_containers_os,omitempty"`
		VulnHostOs       *bool `json:"vuln_host_os,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"vuln_containers_os", "vuln_host_os"})
	} else {
		return err
	}
	o.VulnContainersOs = all.VulnContainersOs
	o.VulnHostOs = all.VulnHostOs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
