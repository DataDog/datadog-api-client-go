// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudcraftWidgetDefinitionProvider Cloud provider for the Cloudcraft widget.
type CloudcraftWidgetDefinitionProvider string

// List of CloudcraftWidgetDefinitionProvider.
const (
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_AWS   CloudcraftWidgetDefinitionProvider = "aws"
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_AZURE CloudcraftWidgetDefinitionProvider = "azure"
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_GCP   CloudcraftWidgetDefinitionProvider = "gcp"
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_NDM   CloudcraftWidgetDefinitionProvider = "ndm"
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_OCI   CloudcraftWidgetDefinitionProvider = "oci"
)

var allowedCloudcraftWidgetDefinitionProviderEnumValues = []CloudcraftWidgetDefinitionProvider{
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_AWS,
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_AZURE,
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_GCP,
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_NDM,
	CLOUDCRAFTWIDGETDEFINITIONPROVIDER_OCI,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CloudcraftWidgetDefinitionProvider) GetAllowedValues() []CloudcraftWidgetDefinitionProvider {
	return allowedCloudcraftWidgetDefinitionProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudcraftWidgetDefinitionProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudcraftWidgetDefinitionProvider(value)
	return nil
}

// NewCloudcraftWidgetDefinitionProviderFromValue returns a pointer to a valid CloudcraftWidgetDefinitionProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudcraftWidgetDefinitionProviderFromValue(v string) (*CloudcraftWidgetDefinitionProvider, error) {
	ev := CloudcraftWidgetDefinitionProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudcraftWidgetDefinitionProvider: valid values are %v", v, allowedCloudcraftWidgetDefinitionProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudcraftWidgetDefinitionProvider) IsValid() bool {
	for _, existing := range allowedCloudcraftWidgetDefinitionProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudcraftWidgetDefinitionProvider value.
func (v CloudcraftWidgetDefinitionProvider) Ptr() *CloudcraftWidgetDefinitionProvider {
	return &v
}
