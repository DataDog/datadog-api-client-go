// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentEntity Kind of entity that a segment contains.
type SegmentEntity string

// List of SegmentEntity.
const (
	SEGMENTENTITY_USER    SegmentEntity = "user"
	SEGMENTENTITY_ACCOUNT SegmentEntity = "account"
)

var allowedSegmentEntityEnumValues = []SegmentEntity{
	SEGMENTENTITY_USER,
	SEGMENTENTITY_ACCOUNT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SegmentEntity) GetAllowedValues() []SegmentEntity {
	return allowedSegmentEntityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SegmentEntity) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SegmentEntity(value)
	return nil
}

// NewSegmentEntityFromValue returns a pointer to a valid SegmentEntity
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSegmentEntityFromValue(v string) (*SegmentEntity, error) {
	ev := SegmentEntity(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SegmentEntity: valid values are %v", v, allowedSegmentEntityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SegmentEntity) IsValid() bool {
	for _, existing := range allowedSegmentEntityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SegmentEntity value.
func (v SegmentEntity) Ptr() *SegmentEntity {
	return &v
}
