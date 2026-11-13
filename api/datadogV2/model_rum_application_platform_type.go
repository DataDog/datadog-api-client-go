// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationPlatformType Type of the RUM application, that is the platform of the RUM SDK. Defaults to `browser` on create.
type RUMApplicationPlatformType string

// List of RUMApplicationPlatformType.
const (
	RUMAPPLICATIONPLATFORMTYPE_BROWSER              RUMApplicationPlatformType = "browser"
	RUMAPPLICATIONPLATFORMTYPE_IOS                  RUMApplicationPlatformType = "ios"
	RUMAPPLICATIONPLATFORMTYPE_ANDROID              RUMApplicationPlatformType = "android"
	RUMAPPLICATIONPLATFORMTYPE_REACT_NATIVE         RUMApplicationPlatformType = "react-native"
	RUMAPPLICATIONPLATFORMTYPE_FLUTTER              RUMApplicationPlatformType = "flutter"
	RUMAPPLICATIONPLATFORMTYPE_ROKU                 RUMApplicationPlatformType = "roku"
	RUMAPPLICATIONPLATFORMTYPE_ELECTRON             RUMApplicationPlatformType = "electron"
	RUMAPPLICATIONPLATFORMTYPE_UNITY                RUMApplicationPlatformType = "unity"
	RUMAPPLICATIONPLATFORMTYPE_KOTLIN_MULTIPLATFORM RUMApplicationPlatformType = "kotlin-multiplatform"
	RUMAPPLICATIONPLATFORMTYPE_REACT                RUMApplicationPlatformType = "react"
	RUMAPPLICATIONPLATFORMTYPE_NEXTJS               RUMApplicationPlatformType = "nextjs"
	RUMAPPLICATIONPLATFORMTYPE_ANGULAR              RUMApplicationPlatformType = "angular"
	RUMAPPLICATIONPLATFORMTYPE_VUEJS                RUMApplicationPlatformType = "vuejs"
	RUMAPPLICATIONPLATFORMTYPE_NUXT                 RUMApplicationPlatformType = "nuxt"
	RUMAPPLICATIONPLATFORMTYPE_TANSTACK             RUMApplicationPlatformType = "tanstack"
	RUMAPPLICATIONPLATFORMTYPE_MAUI                 RUMApplicationPlatformType = "maui"
	RUMAPPLICATIONPLATFORMTYPE_SALESFORCE_LWC       RUMApplicationPlatformType = "salesforce-lwc"
	RUMAPPLICATIONPLATFORMTYPE_SHOPIFY              RUMApplicationPlatformType = "shopify"
	RUMAPPLICATIONPLATFORMTYPE_CPP                  RUMApplicationPlatformType = "cpp"
	RUMAPPLICATIONPLATFORMTYPE_UNREAL               RUMApplicationPlatformType = "unreal"
)

var allowedRUMApplicationPlatformTypeEnumValues = []RUMApplicationPlatformType{
	RUMAPPLICATIONPLATFORMTYPE_BROWSER,
	RUMAPPLICATIONPLATFORMTYPE_IOS,
	RUMAPPLICATIONPLATFORMTYPE_ANDROID,
	RUMAPPLICATIONPLATFORMTYPE_REACT_NATIVE,
	RUMAPPLICATIONPLATFORMTYPE_FLUTTER,
	RUMAPPLICATIONPLATFORMTYPE_ROKU,
	RUMAPPLICATIONPLATFORMTYPE_ELECTRON,
	RUMAPPLICATIONPLATFORMTYPE_UNITY,
	RUMAPPLICATIONPLATFORMTYPE_KOTLIN_MULTIPLATFORM,
	RUMAPPLICATIONPLATFORMTYPE_REACT,
	RUMAPPLICATIONPLATFORMTYPE_NEXTJS,
	RUMAPPLICATIONPLATFORMTYPE_ANGULAR,
	RUMAPPLICATIONPLATFORMTYPE_VUEJS,
	RUMAPPLICATIONPLATFORMTYPE_NUXT,
	RUMAPPLICATIONPLATFORMTYPE_TANSTACK,
	RUMAPPLICATIONPLATFORMTYPE_MAUI,
	RUMAPPLICATIONPLATFORMTYPE_SALESFORCE_LWC,
	RUMAPPLICATIONPLATFORMTYPE_SHOPIFY,
	RUMAPPLICATIONPLATFORMTYPE_CPP,
	RUMAPPLICATIONPLATFORMTYPE_UNREAL,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RUMApplicationPlatformType) GetAllowedValues() []RUMApplicationPlatformType {
	return allowedRUMApplicationPlatformTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RUMApplicationPlatformType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RUMApplicationPlatformType(value)
	return nil
}

// NewRUMApplicationPlatformTypeFromValue returns a pointer to a valid RUMApplicationPlatformType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRUMApplicationPlatformTypeFromValue(v string) (*RUMApplicationPlatformType, error) {
	ev := RUMApplicationPlatformType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RUMApplicationPlatformType: valid values are %v", v, allowedRUMApplicationPlatformTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RUMApplicationPlatformType) IsValid() bool {
	for _, existing := range allowedRUMApplicationPlatformTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RUMApplicationPlatformType value.
func (v RUMApplicationPlatformType) Ptr() *RUMApplicationPlatformType {
	return &v
}
