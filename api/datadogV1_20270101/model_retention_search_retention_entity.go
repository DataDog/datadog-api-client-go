// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearchRetentionEntity Test-only generated schema for the dashboard versioning proof of concept.
type RetentionSearchRetentionEntity string

// List of RetentionSearchRetentionEntity.
const (
	RETENTIONSEARCHRETENTIONENTITY_USR_ID     RetentionSearchRetentionEntity = "@usr.id"
	RETENTIONSEARCHRETENTIONENTITY_ACCOUNT_ID RetentionSearchRetentionEntity = "@account.id"
)

var allowedRetentionSearchRetentionEntityEnumValues = []RetentionSearchRetentionEntity{
	RETENTIONSEARCHRETENTIONENTITY_USR_ID,
	RETENTIONSEARCHRETENTIONENTITY_ACCOUNT_ID,
}

// GetAllowedValues reeturns the list of possible values.
func (v *RetentionSearchRetentionEntity) GetAllowedValues() []RetentionSearchRetentionEntity {
	return allowedRetentionSearchRetentionEntityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RetentionSearchRetentionEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RetentionSearchRetentionEntity(value)
	return nil
}

// NewRetentionSearchRetentionEntityFromValue returns a pointer to a valid RetentionSearchRetentionEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRetentionSearchRetentionEntityFromValue(v string) (*RetentionSearchRetentionEntity, error) {
	ev := RetentionSearchRetentionEntity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RetentionSearchRetentionEntity: valid values are %v", v, allowedRetentionSearchRetentionEntityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RetentionSearchRetentionEntity) IsValid() bool {
	for _, existing := range allowedRetentionSearchRetentionEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetentionSearchRetentionEntity value.
func (v RetentionSearchRetentionEntity) Ptr() *RetentionSearchRetentionEntity {
	return &v
}
