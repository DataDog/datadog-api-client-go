// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScannedAssetMetadataType The JSON:API type.
type ScannedAssetMetadataType string

// List of ScannedAssetMetadataType.
const (
	SCANNEDASSETMETADATATYPE_SCANNED_ASSETS_METADATA ScannedAssetMetadataType = "scanned-assets-metadata"
)

var allowedScannedAssetMetadataTypeEnumValues = []ScannedAssetMetadataType{
	SCANNEDASSETMETADATATYPE_SCANNED_ASSETS_METADATA,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ScannedAssetMetadataType) GetAllowedValues() []ScannedAssetMetadataType {
	return allowedScannedAssetMetadataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ScannedAssetMetadataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ScannedAssetMetadataType(value)
	return nil
}

// NewScannedAssetMetadataTypeFromValue returns a pointer to a valid ScannedAssetMetadataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewScannedAssetMetadataTypeFromValue(v string) (*ScannedAssetMetadataType, error) {
	ev := ScannedAssetMetadataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ScannedAssetMetadataType: valid values are %v", v, allowedScannedAssetMetadataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ScannedAssetMetadataType) IsValid() bool {
	for _, existing := range allowedScannedAssetMetadataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScannedAssetMetadataType value.
func (v ScannedAssetMetadataType) Ptr() *ScannedAssetMetadataType {
	return &v
}
