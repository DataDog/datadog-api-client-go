// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GenerateCampaignTeamReportsRequestDataType
type GenerateCampaignTeamReportsRequestDataType string

// List of GenerateCampaignTeamReportsRequestDataType.
const (
	GENERATECAMPAIGNTEAMREPORTSREQUESTDATATYPE_CAMPAIGN_TEAM_REPORT GenerateCampaignTeamReportsRequestDataType = "campaign-team-report"
)

var allowedGenerateCampaignTeamReportsRequestDataTypeEnumValues = []GenerateCampaignTeamReportsRequestDataType{
	GENERATECAMPAIGNTEAMREPORTSREQUESTDATATYPE_CAMPAIGN_TEAM_REPORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *GenerateCampaignTeamReportsRequestDataType) GetAllowedValues() []GenerateCampaignTeamReportsRequestDataType {
	return allowedGenerateCampaignTeamReportsRequestDataTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *GenerateCampaignTeamReportsRequestDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = GenerateCampaignTeamReportsRequestDataType(value)
	return nil
}

// NewGenerateCampaignTeamReportsRequestDataTypeFromValue returns a pointer to a valid GenerateCampaignTeamReportsRequestDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewGenerateCampaignTeamReportsRequestDataTypeFromValue(v string) (*GenerateCampaignTeamReportsRequestDataType, error) {
	ev := GenerateCampaignTeamReportsRequestDataType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for GenerateCampaignTeamReportsRequestDataType: valid values are %v", v, allowedGenerateCampaignTeamReportsRequestDataTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v GenerateCampaignTeamReportsRequestDataType) IsValid() bool {
	for _, existing := range allowedGenerateCampaignTeamReportsRequestDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GenerateCampaignTeamReportsRequestDataType value.
func (v GenerateCampaignTeamReportsRequestDataType) Ptr() *GenerateCampaignTeamReportsRequestDataType {
	return &v
}
