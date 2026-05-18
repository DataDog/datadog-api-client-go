// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScatterplotTableRequest Scatterplot table request. Supports two modes:
// - **Formulas and functions** (default): `request_type` is absent or `"table"`. Uses `queries` and `formulas`.
// - **Data projection**: `request_type` is `"data_projection"`. Uses `query`, `projection`, and optionally `limit`.
type ScatterplotTableRequest struct {
	// List of Scatterplot formulas that operate on queries.
	Formulas []ScatterplotWidgetFormula `json:"formulas,omitempty"`
	// Maximum number of rows to return. Used when `request_type` is `"data_projection"`.
	Limit *int64 `json:"limit,omitempty"`
	// The projection configuration for a scatterplot data projection request.
	Projection *ScatterplotDataProjectionProjection `json:"projection,omitempty"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
	// The query for a scatterplot data projection request.
	Query *ScatterplotDataProjectionQuery `json:"query,omitempty"`
	// The type of the scatterplot table request.
	RequestType *ScatterplotTableRequestType `json:"request_type,omitempty"`
	// Timeseries, scalar, or event list response. Event list response formats are supported by Geomap widgets.
	ResponseFormat *FormulaAndFunctionResponseFormat `json:"response_format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScatterplotTableRequest instantiates a new ScatterplotTableRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScatterplotTableRequest() *ScatterplotTableRequest {
	this := ScatterplotTableRequest{}
	return &this
}

// NewScatterplotTableRequestWithDefaults instantiates a new ScatterplotTableRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScatterplotTableRequestWithDefaults() *ScatterplotTableRequest {
	this := ScatterplotTableRequest{}
	return &this
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetFormulas() []ScatterplotWidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []ScatterplotWidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetFormulasOk() (*[]ScatterplotWidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []ScatterplotWidgetFormula and assigns it to the Formulas field.
func (o *ScatterplotTableRequest) SetFormulas(v []ScatterplotWidgetFormula) {
	o.Formulas = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ScatterplotTableRequest) SetLimit(v int64) {
	o.Limit = &v
}

// GetProjection returns the Projection field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetProjection() ScatterplotDataProjectionProjection {
	if o == nil || o.Projection == nil {
		var ret ScatterplotDataProjectionProjection
		return ret
	}
	return *o.Projection
}

// GetProjectionOk returns a tuple with the Projection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetProjectionOk() (*ScatterplotDataProjectionProjection, bool) {
	if o == nil || o.Projection == nil {
		return nil, false
	}
	return o.Projection, true
}

// HasProjection returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasProjection() bool {
	return o != nil && o.Projection != nil
}

// SetProjection gets a reference to the given ScatterplotDataProjectionProjection and assigns it to the Projection field.
func (o *ScatterplotTableRequest) SetProjection(v ScatterplotDataProjectionProjection) {
	o.Projection = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil || o.Queries == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []FormulaAndFunctionQueryDefinition and assigns it to the Queries field.
func (o *ScatterplotTableRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetQuery() ScatterplotDataProjectionQuery {
	if o == nil || o.Query == nil {
		var ret ScatterplotDataProjectionQuery
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetQueryOk() (*ScatterplotDataProjectionQuery, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given ScatterplotDataProjectionQuery and assigns it to the Query field.
func (o *ScatterplotTableRequest) SetQuery(v ScatterplotDataProjectionQuery) {
	o.Query = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetRequestType() ScatterplotTableRequestType {
	if o == nil || o.RequestType == nil {
		var ret ScatterplotTableRequestType
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetRequestTypeOk() (*ScatterplotTableRequestType, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasRequestType() bool {
	return o != nil && o.RequestType != nil
}

// SetRequestType gets a reference to the given ScatterplotTableRequestType and assigns it to the RequestType field.
func (o *ScatterplotTableRequest) SetRequestType(v ScatterplotTableRequestType) {
	o.RequestType = &v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *ScatterplotTableRequest) GetResponseFormat() FormulaAndFunctionResponseFormat {
	if o == nil || o.ResponseFormat == nil {
		var ret FormulaAndFunctionResponseFormat
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScatterplotTableRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool) {
	if o == nil || o.ResponseFormat == nil {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *ScatterplotTableRequest) HasResponseFormat() bool {
	return o != nil && o.ResponseFormat != nil
}

// SetResponseFormat gets a reference to the given FormulaAndFunctionResponseFormat and assigns it to the ResponseFormat field.
func (o *ScatterplotTableRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat) {
	o.ResponseFormat = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScatterplotTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Projection != nil {
		toSerialize["projection"] = o.Projection
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.RequestType != nil {
		toSerialize["request_type"] = o.RequestType
	}
	if o.ResponseFormat != nil {
		toSerialize["response_format"] = o.ResponseFormat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScatterplotTableRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       []ScatterplotWidgetFormula           `json:"formulas,omitempty"`
		Limit          *int64                               `json:"limit,omitempty"`
		Projection     *ScatterplotDataProjectionProjection `json:"projection,omitempty"`
		Queries        []FormulaAndFunctionQueryDefinition  `json:"queries,omitempty"`
		Query          *ScatterplotDataProjectionQuery      `json:"query,omitempty"`
		RequestType    *ScatterplotTableRequestType         `json:"request_type,omitempty"`
		ResponseFormat *FormulaAndFunctionResponseFormat    `json:"response_format,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formulas", "limit", "projection", "queries", "query", "request_type", "response_format"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Formulas = all.Formulas
	o.Limit = all.Limit
	if all.Projection != nil && all.Projection.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Projection = all.Projection
	o.Queries = all.Queries
	if all.Query != nil && all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = all.Query
	if all.RequestType != nil && !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = all.RequestType
	}
	if all.ResponseFormat != nil && !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = all.ResponseFormat
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
