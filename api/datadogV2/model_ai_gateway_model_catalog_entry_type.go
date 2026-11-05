// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelCatalogEntryType Model catalog entry resource type.
type AiGatewayModelCatalogEntryType string

// List of AiGatewayModelCatalogEntryType.
const (
	AIGATEWAYMODELCATALOGENTRYTYPE_MODEL_CATALOG_ENTRIES AiGatewayModelCatalogEntryType = "model-catalog-entries"
)

var allowedAiGatewayModelCatalogEntryTypeEnumValues = []AiGatewayModelCatalogEntryType{
	AIGATEWAYMODELCATALOGENTRYTYPE_MODEL_CATALOG_ENTRIES,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AiGatewayModelCatalogEntryType) GetAllowedValues() []AiGatewayModelCatalogEntryType {
	return allowedAiGatewayModelCatalogEntryTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AiGatewayModelCatalogEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AiGatewayModelCatalogEntryType(value)
	return nil
}

// NewAiGatewayModelCatalogEntryTypeFromValue returns a pointer to a valid AiGatewayModelCatalogEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAiGatewayModelCatalogEntryTypeFromValue(v string) (*AiGatewayModelCatalogEntryType, error) {
	ev := AiGatewayModelCatalogEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AiGatewayModelCatalogEntryType: valid values are %v", v, allowedAiGatewayModelCatalogEntryTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AiGatewayModelCatalogEntryType) IsValid() bool {
	for _, existing := range allowedAiGatewayModelCatalogEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiGatewayModelCatalogEntryType value.
func (v AiGatewayModelCatalogEntryType) Ptr() *AiGatewayModelCatalogEntryType {
	return &v
}
