// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GenerateCampaignReportRequestDataType
type GenerateCampaignReportRequestDataType string

// List of GenerateCampaignReportRequestDataType.
const (
	GENERATECAMPAIGNREPORTREQUESTDATATYPE_CAMPAIGN_REPORT GenerateCampaignReportRequestDataType = "campaign-report"
)

var allowedGenerateCampaignReportRequestDataTypeEnumValues = []GenerateCampaignReportRequestDataType{
	GENERATECAMPAIGNREPORTREQUESTDATATYPE_CAMPAIGN_REPORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GenerateCampaignReportRequestDataType) GetAllowedValues() []GenerateCampaignReportRequestDataType {
	return allowedGenerateCampaignReportRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GenerateCampaignReportRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GenerateCampaignReportRequestDataType(value)
	return nil
}

// NewGenerateCampaignReportRequestDataTypeFromValue returns a pointer to a valid GenerateCampaignReportRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGenerateCampaignReportRequestDataTypeFromValue(v string) (*GenerateCampaignReportRequestDataType, error) {
	ev := GenerateCampaignReportRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GenerateCampaignReportRequestDataType: valid values are %v", v, allowedGenerateCampaignReportRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GenerateCampaignReportRequestDataType) IsValid() bool {
	for _, existing := range allowedGenerateCampaignReportRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GenerateCampaignReportRequestDataType value.
func (v GenerateCampaignReportRequestDataType) Ptr() *GenerateCampaignReportRequestDataType {
	return &v
}
