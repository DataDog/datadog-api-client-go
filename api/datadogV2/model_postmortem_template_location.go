// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PostmortemTemplateLocation The location where the postmortem is created and stored.
type PostmortemTemplateLocation string

// List of PostmortemTemplateLocation.
const (
	POSTMORTEMTEMPLATELOCATION_DATADOG_NOTEBOOKS PostmortemTemplateLocation = "datadog_notebooks"
	POSTMORTEMTEMPLATELOCATION_CONFLUENCE        PostmortemTemplateLocation = "confluence"
	POSTMORTEMTEMPLATELOCATION_GOOGLE_DOCS       PostmortemTemplateLocation = "google_docs"
)

var allowedPostmortemTemplateLocationEnumValues = []PostmortemTemplateLocation{
	POSTMORTEMTEMPLATELOCATION_DATADOG_NOTEBOOKS,
	POSTMORTEMTEMPLATELOCATION_CONFLUENCE,
	POSTMORTEMTEMPLATELOCATION_GOOGLE_DOCS,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PostmortemTemplateLocation) GetAllowedValues() []PostmortemTemplateLocation {
	return allowedPostmortemTemplateLocationEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PostmortemTemplateLocation) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PostmortemTemplateLocation(value)
	return nil
}

// NewPostmortemTemplateLocationFromValue returns a pointer to a valid PostmortemTemplateLocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPostmortemTemplateLocationFromValue(v string) (*PostmortemTemplateLocation, error) {
	ev := PostmortemTemplateLocation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PostmortemTemplateLocation: valid values are %v", v, allowedPostmortemTemplateLocationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PostmortemTemplateLocation) IsValid() bool {
	for _, existing := range allowedPostmortemTemplateLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostmortemTemplateLocation value.
func (v PostmortemTemplateLocation) Ptr() *PostmortemTemplateLocation {
	return &v
}
