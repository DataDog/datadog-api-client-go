// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SunburstWidgetRequest Request definition of sunburst widget.
type SunburstWidgetRequest struct {
	// The log query.
	ApmQuery *LogQueryDefinition `json:"apm_query,omitempty"`
	// The log query.
	AuditQuery *LogQueryDefinition `json:"audit_query,omitempty"`
	// The log query.
	EventQuery *LogQueryDefinition `json:"event_query,omitempty"`
	// List of formulas that operate on queries.
	Formulas []WidgetFormula `json:"formulas,omitempty"`
	// The log query.
	LogQuery *LogQueryDefinition `json:"log_query,omitempty"`
	// The log query.
	NetworkQuery *LogQueryDefinition `json:"network_query,omitempty"`
	// The process query to use in the widget.
	ProcessQuery *ProcessQueryDefinition `json:"process_query,omitempty"`
	// The log query.
	ProfileMetricsQuery *LogQueryDefinition `json:"profile_metrics_query,omitempty"`
	// Widget query.
	Q *string `json:"q,omitempty"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
	// Timeseries, scalar, or event list response. Event list response formats are supported by Geomap widgets.
	ResponseFormat *FormulaAndFunctionResponseFormat `json:"response_format,omitempty"`
	// The log query.
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
	// The log query.
	SecurityQuery *LogQueryDefinition `json:"security_query,omitempty"`
	// Widget style definition.
	Style *WidgetStyle `json:"style,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSunburstWidgetRequest instantiates a new SunburstWidgetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSunburstWidgetRequest() *SunburstWidgetRequest {
	this := SunburstWidgetRequest{}
	return &this
}

// NewSunburstWidgetRequestWithDefaults instantiates a new SunburstWidgetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSunburstWidgetRequestWithDefaults() *SunburstWidgetRequest {
	this := SunburstWidgetRequest{}
	return &this
}

// GetApmQuery returns the ApmQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetApmQuery() LogQueryDefinition {
	if o == nil || o.ApmQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ApmQuery == nil {
		return nil, false
	}
	return o.ApmQuery, true
}

// HasApmQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasApmQuery() bool {
	return o != nil && o.ApmQuery != nil
}

// SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.
func (o *SunburstWidgetRequest) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = &v
}

// GetAuditQuery returns the AuditQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetAuditQuery() LogQueryDefinition {
	if o == nil || o.AuditQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.AuditQuery
}

// GetAuditQueryOk returns a tuple with the AuditQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetAuditQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.AuditQuery == nil {
		return nil, false
	}
	return o.AuditQuery, true
}

// HasAuditQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasAuditQuery() bool {
	return o != nil && o.AuditQuery != nil
}

// SetAuditQuery gets a reference to the given LogQueryDefinition and assigns it to the AuditQuery field.
func (o *SunburstWidgetRequest) SetAuditQuery(v LogQueryDefinition) {
	o.AuditQuery = &v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetEventQuery() LogQueryDefinition {
	if o == nil || o.EventQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetEventQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasEventQuery() bool {
	return o != nil && o.EventQuery != nil
}

// SetEventQuery gets a reference to the given LogQueryDefinition and assigns it to the EventQuery field.
func (o *SunburstWidgetRequest) SetEventQuery(v LogQueryDefinition) {
	o.EventQuery = &v
}

// GetFormulas returns the Formulas field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetFormulas() []WidgetFormula {
	if o == nil || o.Formulas == nil {
		var ret []WidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetFormulasOk() (*[]WidgetFormula, bool) {
	if o == nil || o.Formulas == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// HasFormulas returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasFormulas() bool {
	return o != nil && o.Formulas != nil
}

// SetFormulas gets a reference to the given []WidgetFormula and assigns it to the Formulas field.
func (o *SunburstWidgetRequest) SetFormulas(v []WidgetFormula) {
	o.Formulas = v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasLogQuery() bool {
	return o != nil && o.LogQuery != nil
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *SunburstWidgetRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetNetworkQuery returns the NetworkQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetNetworkQuery() LogQueryDefinition {
	if o == nil || o.NetworkQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.NetworkQuery
}

// GetNetworkQueryOk returns a tuple with the NetworkQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.NetworkQuery == nil {
		return nil, false
	}
	return o.NetworkQuery, true
}

// HasNetworkQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasNetworkQuery() bool {
	return o != nil && o.NetworkQuery != nil
}

// SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.
func (o *SunburstWidgetRequest) SetNetworkQuery(v LogQueryDefinition) {
	o.NetworkQuery = &v
}

// GetProcessQuery returns the ProcessQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetProcessQuery() ProcessQueryDefinition {
	if o == nil || o.ProcessQuery == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return *o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil || o.ProcessQuery == nil {
		return nil, false
	}
	return o.ProcessQuery, true
}

// HasProcessQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasProcessQuery() bool {
	return o != nil && o.ProcessQuery != nil
}

// SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.
func (o *SunburstWidgetRequest) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = &v
}

// GetProfileMetricsQuery returns the ProfileMetricsQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetProfileMetricsQuery() LogQueryDefinition {
	if o == nil || o.ProfileMetricsQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ProfileMetricsQuery
}

// GetProfileMetricsQueryOk returns a tuple with the ProfileMetricsQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetProfileMetricsQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ProfileMetricsQuery == nil {
		return nil, false
	}
	return o.ProfileMetricsQuery, true
}

// HasProfileMetricsQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasProfileMetricsQuery() bool {
	return o != nil && o.ProfileMetricsQuery != nil
}

// SetProfileMetricsQuery gets a reference to the given LogQueryDefinition and assigns it to the ProfileMetricsQuery field.
func (o *SunburstWidgetRequest) SetProfileMetricsQuery(v LogQueryDefinition) {
	o.ProfileMetricsQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasQ() bool {
	return o != nil && o.Q != nil
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *SunburstWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil || o.Queries == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return &o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasQueries() bool {
	return o != nil && o.Queries != nil
}

// SetQueries gets a reference to the given []FormulaAndFunctionQueryDefinition and assigns it to the Queries field.
func (o *SunburstWidgetRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetResponseFormat() FormulaAndFunctionResponseFormat {
	if o == nil || o.ResponseFormat == nil {
		var ret FormulaAndFunctionResponseFormat
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetResponseFormatOk() (*FormulaAndFunctionResponseFormat, bool) {
	if o == nil || o.ResponseFormat == nil {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasResponseFormat() bool {
	return o != nil && o.ResponseFormat != nil
}

// SetResponseFormat gets a reference to the given FormulaAndFunctionResponseFormat and assigns it to the ResponseFormat field.
func (o *SunburstWidgetRequest) SetResponseFormat(v FormulaAndFunctionResponseFormat) {
	o.ResponseFormat = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasRumQuery() bool {
	return o != nil && o.RumQuery != nil
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *SunburstWidgetRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

// GetSecurityQuery returns the SecurityQuery field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetSecurityQuery() LogQueryDefinition {
	if o == nil || o.SecurityQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.SecurityQuery
}

// GetSecurityQueryOk returns a tuple with the SecurityQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetSecurityQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.SecurityQuery == nil {
		return nil, false
	}
	return o.SecurityQuery, true
}

// HasSecurityQuery returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasSecurityQuery() bool {
	return o != nil && o.SecurityQuery != nil
}

// SetSecurityQuery gets a reference to the given LogQueryDefinition and assigns it to the SecurityQuery field.
func (o *SunburstWidgetRequest) SetSecurityQuery(v LogQueryDefinition) {
	o.SecurityQuery = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *SunburstWidgetRequest) GetStyle() WidgetStyle {
	if o == nil || o.Style == nil {
		var ret WidgetStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SunburstWidgetRequest) GetStyleOk() (*WidgetStyle, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *SunburstWidgetRequest) HasStyle() bool {
	return o != nil && o.Style != nil
}

// SetStyle gets a reference to the given WidgetStyle and assigns it to the Style field.
func (o *SunburstWidgetRequest) SetStyle(v WidgetStyle) {
	o.Style = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SunburstWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApmQuery != nil {
		toSerialize["apm_query"] = o.ApmQuery
	}
	if o.AuditQuery != nil {
		toSerialize["audit_query"] = o.AuditQuery
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.Formulas != nil {
		toSerialize["formulas"] = o.Formulas
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.NetworkQuery != nil {
		toSerialize["network_query"] = o.NetworkQuery
	}
	if o.ProcessQuery != nil {
		toSerialize["process_query"] = o.ProcessQuery
	}
	if o.ProfileMetricsQuery != nil {
		toSerialize["profile_metrics_query"] = o.ProfileMetricsQuery
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.ResponseFormat != nil {
		toSerialize["response_format"] = o.ResponseFormat
	}
	if o.RumQuery != nil {
		toSerialize["rum_query"] = o.RumQuery
	}
	if o.SecurityQuery != nil {
		toSerialize["security_query"] = o.SecurityQuery
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SunburstWidgetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApmQuery            *LogQueryDefinition                 `json:"apm_query,omitempty"`
		AuditQuery          *LogQueryDefinition                 `json:"audit_query,omitempty"`
		EventQuery          *LogQueryDefinition                 `json:"event_query,omitempty"`
		Formulas            []WidgetFormula                     `json:"formulas,omitempty"`
		LogQuery            *LogQueryDefinition                 `json:"log_query,omitempty"`
		NetworkQuery        *LogQueryDefinition                 `json:"network_query,omitempty"`
		ProcessQuery        *ProcessQueryDefinition             `json:"process_query,omitempty"`
		ProfileMetricsQuery *LogQueryDefinition                 `json:"profile_metrics_query,omitempty"`
		Q                   *string                             `json:"q,omitempty"`
		Queries             []FormulaAndFunctionQueryDefinition `json:"queries,omitempty"`
		ResponseFormat      *FormulaAndFunctionResponseFormat   `json:"response_format,omitempty"`
		RumQuery            *LogQueryDefinition                 `json:"rum_query,omitempty"`
		SecurityQuery       *LogQueryDefinition                 `json:"security_query,omitempty"`
		Style               *WidgetStyle                        `json:"style,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm_query", "audit_query", "event_query", "formulas", "log_query", "network_query", "process_query", "profile_metrics_query", "q", "queries", "response_format", "rum_query", "security_query", "style"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApmQuery != nil && all.ApmQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApmQuery = all.ApmQuery
	if all.AuditQuery != nil && all.AuditQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuditQuery = all.AuditQuery
	if all.EventQuery != nil && all.EventQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EventQuery = all.EventQuery
	o.Formulas = all.Formulas
	if all.LogQuery != nil && all.LogQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogQuery = all.LogQuery
	if all.NetworkQuery != nil && all.NetworkQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NetworkQuery = all.NetworkQuery
	if all.ProcessQuery != nil && all.ProcessQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProcessQuery = all.ProcessQuery
	if all.ProfileMetricsQuery != nil && all.ProfileMetricsQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProfileMetricsQuery = all.ProfileMetricsQuery
	o.Q = all.Q
	o.Queries = all.Queries
	if all.ResponseFormat != nil && !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = all.ResponseFormat
	}
	if all.RumQuery != nil && all.RumQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RumQuery = all.RumQuery
	if all.SecurityQuery != nil && all.SecurityQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SecurityQuery = all.SecurityQuery
	if all.Style != nil && all.Style.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Style = all.Style

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
