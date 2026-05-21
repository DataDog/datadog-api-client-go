// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchIncidentsExportRequest Request to export incidents as CSV.
type IncidentSearchIncidentsExportRequest struct {
	// The list of fields to include in the export.
	Fields []string `json:"fields"`
	// The search query to filter incidents for export.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentSearchIncidentsExportRequest instantiates a new IncidentSearchIncidentsExportRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSearchIncidentsExportRequest(fields []string, query string) *IncidentSearchIncidentsExportRequest {
	this := IncidentSearchIncidentsExportRequest{}
	this.Fields = fields
	this.Query = query
	return &this
}

// NewIncidentSearchIncidentsExportRequestWithDefaults instantiates a new IncidentSearchIncidentsExportRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSearchIncidentsExportRequestWithDefaults() *IncidentSearchIncidentsExportRequest {
	this := IncidentSearchIncidentsExportRequest{}
	return &this
}

// GetFields returns the Fields field value.
func (o *IncidentSearchIncidentsExportRequest) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchIncidentsExportRequest) GetFieldsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *IncidentSearchIncidentsExportRequest) SetFields(v []string) {
	o.Fields = v
}

// GetQuery returns the Query field value.
func (o *IncidentSearchIncidentsExportRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchIncidentsExportRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *IncidentSearchIncidentsExportRequest) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSearchIncidentsExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentSearchIncidentsExportRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields *[]string `json:"fields"`
		Query  *string   `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "query"})
	} else {
		return err
	}
	o.Fields = *all.Fields
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
