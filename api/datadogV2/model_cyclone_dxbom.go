// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXBOM CycloneDX 1.5 Bill of Materials (BOM) for importing vulnerabilities.
type CycloneDXBOM struct {
	// The format of the BOM. Must be "CycloneDX".
	BomFormat string `json:"bomFormat"`
	// List of components (libraries, applications, or operating systems) that are affected by vulnerabilities.
	Components []CycloneDXComponent `json:"components"`
	// Metadata for the CycloneDX BOM.
	Metadata CycloneDXMetadata `json:"metadata"`
	// The version of the CycloneDX specification. Must be "1.5".
	SpecVersion string `json:"specVersion"`
	// The version of the BOM.
	Version int64 `json:"version"`
	// List of vulnerabilities to be imported.
	Vulnerabilities []CycloneDXVulnerability `json:"vulnerabilities"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXBOM instantiates a new CycloneDXBOM object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXBOM(bomFormat string, components []CycloneDXComponent, metadata CycloneDXMetadata, specVersion string, version int64, vulnerabilities []CycloneDXVulnerability) *CycloneDXBOM {
	this := CycloneDXBOM{}
	this.BomFormat = bomFormat
	this.Components = components
	this.Metadata = metadata
	this.SpecVersion = specVersion
	this.Version = version
	this.Vulnerabilities = vulnerabilities
	return &this
}

// NewCycloneDXBOMWithDefaults instantiates a new CycloneDXBOM object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXBOMWithDefaults() *CycloneDXBOM {
	this := CycloneDXBOM{}
	return &this
}

// GetBomFormat returns the BomFormat field value.
func (o *CycloneDXBOM) GetBomFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BomFormat
}

// GetBomFormatOk returns a tuple with the BomFormat field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBOM) GetBomFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BomFormat, true
}

// SetBomFormat sets field value.
func (o *CycloneDXBOM) SetBomFormat(v string) {
	o.BomFormat = v
}

// GetComponents returns the Components field value.
func (o *CycloneDXBOM) GetComponents() []CycloneDXComponent {
	if o == nil {
		var ret []CycloneDXComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBOM) GetComponentsOk() (*[]CycloneDXComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *CycloneDXBOM) SetComponents(v []CycloneDXComponent) {
	o.Components = v
}

// GetMetadata returns the Metadata field value.
func (o *CycloneDXBOM) GetMetadata() CycloneDXMetadata {
	if o == nil {
		var ret CycloneDXMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBOM) GetMetadataOk() (*CycloneDXMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *CycloneDXBOM) SetMetadata(v CycloneDXMetadata) {
	o.Metadata = v
}

// GetSpecVersion returns the SpecVersion field value.
func (o *CycloneDXBOM) GetSpecVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBOM) GetSpecVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value.
func (o *CycloneDXBOM) SetSpecVersion(v string) {
	o.SpecVersion = v
}

// GetVersion returns the Version field value.
func (o *CycloneDXBOM) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBOM) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *CycloneDXBOM) SetVersion(v int64) {
	o.Version = v
}

// GetVulnerabilities returns the Vulnerabilities field value.
func (o *CycloneDXBOM) GetVulnerabilities() []CycloneDXVulnerability {
	if o == nil {
		var ret []CycloneDXVulnerability
		return ret
	}
	return o.Vulnerabilities
}

// GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBOM) GetVulnerabilitiesOk() (*[]CycloneDXVulnerability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vulnerabilities, true
}

// SetVulnerabilities sets field value.
func (o *CycloneDXBOM) SetVulnerabilities(v []CycloneDXVulnerability) {
	o.Vulnerabilities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXBOM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bomFormat"] = o.BomFormat
	toSerialize["components"] = o.Components
	toSerialize["metadata"] = o.Metadata
	toSerialize["specVersion"] = o.SpecVersion
	toSerialize["version"] = o.Version
	toSerialize["vulnerabilities"] = o.Vulnerabilities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXBOM) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BomFormat       *string                   `json:"bomFormat"`
		Components      *[]CycloneDXComponent     `json:"components"`
		Metadata        *CycloneDXMetadata        `json:"metadata"`
		SpecVersion     *string                   `json:"specVersion"`
		Version         *int64                    `json:"version"`
		Vulnerabilities *[]CycloneDXVulnerability `json:"vulnerabilities"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BomFormat == nil {
		return fmt.Errorf("required field bomFormat missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.SpecVersion == nil {
		return fmt.Errorf("required field specVersion missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	if all.Vulnerabilities == nil {
		return fmt.Errorf("required field vulnerabilities missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bomFormat", "components", "metadata", "specVersion", "version", "vulnerabilities"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BomFormat = *all.BomFormat
	o.Components = *all.Components
	if all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = *all.Metadata
	o.SpecVersion = *all.SpecVersion
	o.Version = *all.Version
	o.Vulnerabilities = *all.Vulnerabilities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
