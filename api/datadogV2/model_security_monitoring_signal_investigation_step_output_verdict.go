// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalInvestigationStepOutputVerdict The verdict from the investigation step.
type SecurityMonitoringSignalInvestigationStepOutputVerdict string

// List of SecurityMonitoringSignalInvestigationStepOutputVerdict.
const (
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_UNSPECIFIED  SecurityMonitoringSignalInvestigationStepOutputVerdict = "unspecified"
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_BENIGN       SecurityMonitoringSignalInvestigationStepOutputVerdict = "benign"
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_SUSPICIOUS   SecurityMonitoringSignalInvestigationStepOutputVerdict = "suspicious"
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_INCONCLUSIVE SecurityMonitoringSignalInvestigationStepOutputVerdict = "inconclusive"
)

var allowedSecurityMonitoringSignalInvestigationStepOutputVerdictEnumValues = []SecurityMonitoringSignalInvestigationStepOutputVerdict{
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_UNSPECIFIED,
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_BENIGN,
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_SUSPICIOUS,
	SECURITYMONITORINGSIGNALINVESTIGATIONSTEPOUTPUTVERDICT_INCONCLUSIVE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SecurityMonitoringSignalInvestigationStepOutputVerdict) GetAllowedValues() []SecurityMonitoringSignalInvestigationStepOutputVerdict {
	return allowedSecurityMonitoringSignalInvestigationStepOutputVerdictEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SecurityMonitoringSignalInvestigationStepOutputVerdict) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SecurityMonitoringSignalInvestigationStepOutputVerdict(value)
	return nil
}

// NewSecurityMonitoringSignalInvestigationStepOutputVerdictFromValue returns a pointer to a valid SecurityMonitoringSignalInvestigationStepOutputVerdict
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSecurityMonitoringSignalInvestigationStepOutputVerdictFromValue(v string) (*SecurityMonitoringSignalInvestigationStepOutputVerdict, error) {
	ev := SecurityMonitoringSignalInvestigationStepOutputVerdict(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SecurityMonitoringSignalInvestigationStepOutputVerdict: valid values are %v", v, allowedSecurityMonitoringSignalInvestigationStepOutputVerdictEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SecurityMonitoringSignalInvestigationStepOutputVerdict) IsValid() bool {
	for _, existing := range allowedSecurityMonitoringSignalInvestigationStepOutputVerdictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMonitoringSignalInvestigationStepOutputVerdict value.
func (v SecurityMonitoringSignalInvestigationStepOutputVerdict) Ptr() *SecurityMonitoringSignalInvestigationStepOutputVerdict {
	return &v
}
