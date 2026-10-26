// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueSampleAttributes Object containing the information of a sample.
type IssueSampleAttributes struct {
	// RUM application identifier of the error.
	ApplicationId *string `json:"application_id,omitempty"`
	// Biases the sample satisfies. A single error can satisfy several biases.
	Biases []IssueSampleBias `json:"biases,omitempty"`
	// Build identifier of the mobile application.
	BuildId *string `json:"build_id,omitempty"`
	// Bundle version of the mobile application.
	BundleVersion *string `json:"bundle_version,omitempty"`
	// Product that reported the error. The value is `NONE` when the error has no known origin.
	DataSource *IssueSampleDataSource `json:"data_source,omitempty"`
	// Debug identifier used to resolve the stack trace of the error to source code.
	DebugId *string `json:"debug_id,omitempty"`
	// Whether an Exception Replay snapshot was captured for the error.
	DebugInfoCaptured *bool `json:"debug_info_captured,omitempty"`
	// Source type set by the SDK or tracer that reported the error, such as `browser` for RUM browser errors
	// or `Crashtracking` for backend crash reports. Usually not set for APM and logs errors.
	ErrorSourceType *string `json:"error_source_type,omitempty"`
	// Parsed stack trace of the error. Only the sample picked by the `git_metadata` bias has frames.
	ErrorStack IssueSampleErrorStack `json:"error_stack"`
	// Raw error event, as stored by its data source. Its fields differ between APM, logs, and RUM.
	Event map[string]interface{} `json:"event,omitempty"`
	// Git commit SHA of the service version.
	GitCommitSha *string `json:"git_commit_sha,omitempty"`
	// Git repository identifier of the service.
	GitRepositoryId *string `json:"git_repository_id,omitempty"`
	// Git repository URL of the service.
	GitRepositoryUrl *string `json:"git_repository_url,omitempty"`
	// Whether a RUM session replay is available for the error.
	HasSessionReplay *bool `json:"has_session_replay,omitempty"`
	// Programming language of the service.
	Language *string `json:"language,omitempty"`
	// Source code locations of the stack frames that resolve to a repository file.
	Links []IssueSampleFrameLink `json:"links,omitempty"`
	// Operating system of the device that reported the error.
	OsName *string `json:"os_name,omitempty"`
	// Service name.
	Service *string `json:"service,omitempty"`
	// RUM session identifier of the error.
	SessionId *string `json:"session_id,omitempty"`
	// URL of the RUM session replay of the error.
	SessionReplayUrl *string `json:"session_replay_url,omitempty"`
	// Identifier of the Exception Replay snapshot of the error.
	SnapshotId *string `json:"snapshot_id,omitempty"`
	// APM span identifier of the error.
	SpanId *string `json:"span_id,omitempty"`
	// Timestamp of the error.
	Timestamp time.Time `json:"timestamp"`
	// APM trace identifier of the error.
	TraceId *string `json:"trace_id,omitempty"`
	// Build variant of the mobile application.
	Variant *string `json:"variant,omitempty"`
	// Version of the service.
	Version *string `json:"version,omitempty"`
	// RUM view identifier of the error.
	ViewId *string `json:"view_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueSampleAttributes instantiates a new IssueSampleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueSampleAttributes(errorStack IssueSampleErrorStack, timestamp time.Time) *IssueSampleAttributes {
	this := IssueSampleAttributes{}
	this.ErrorStack = errorStack
	this.Timestamp = timestamp
	return &this
}

// NewIssueSampleAttributesWithDefaults instantiates a new IssueSampleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueSampleAttributesWithDefaults() *IssueSampleAttributes {
	this := IssueSampleAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasApplicationId() bool {
	return o != nil && o.ApplicationId != nil
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *IssueSampleAttributes) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetBiases returns the Biases field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetBiases() []IssueSampleBias {
	if o == nil || o.Biases == nil {
		var ret []IssueSampleBias
		return ret
	}
	return o.Biases
}

// GetBiasesOk returns a tuple with the Biases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetBiasesOk() (*[]IssueSampleBias, bool) {
	if o == nil || o.Biases == nil {
		return nil, false
	}
	return &o.Biases, true
}

// HasBiases returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasBiases() bool {
	return o != nil && o.Biases != nil
}

// SetBiases gets a reference to the given []IssueSampleBias and assigns it to the Biases field.
func (o *IssueSampleAttributes) SetBiases(v []IssueSampleBias) {
	o.Biases = v
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasBuildId() bool {
	return o != nil && o.BuildId != nil
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *IssueSampleAttributes) SetBuildId(v string) {
	o.BuildId = &v
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetBundleVersion() string {
	if o == nil || o.BundleVersion == nil {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetBundleVersionOk() (*string, bool) {
	if o == nil || o.BundleVersion == nil {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasBundleVersion() bool {
	return o != nil && o.BundleVersion != nil
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *IssueSampleAttributes) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetDataSource() IssueSampleDataSource {
	if o == nil || o.DataSource == nil {
		var ret IssueSampleDataSource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetDataSourceOk() (*IssueSampleDataSource, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given IssueSampleDataSource and assigns it to the DataSource field.
func (o *IssueSampleAttributes) SetDataSource(v IssueSampleDataSource) {
	o.DataSource = &v
}

// GetDebugId returns the DebugId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetDebugId() string {
	if o == nil || o.DebugId == nil {
		var ret string
		return ret
	}
	return *o.DebugId
}

// GetDebugIdOk returns a tuple with the DebugId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetDebugIdOk() (*string, bool) {
	if o == nil || o.DebugId == nil {
		return nil, false
	}
	return o.DebugId, true
}

// HasDebugId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasDebugId() bool {
	return o != nil && o.DebugId != nil
}

// SetDebugId gets a reference to the given string and assigns it to the DebugId field.
func (o *IssueSampleAttributes) SetDebugId(v string) {
	o.DebugId = &v
}

// GetDebugInfoCaptured returns the DebugInfoCaptured field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetDebugInfoCaptured() bool {
	if o == nil || o.DebugInfoCaptured == nil {
		var ret bool
		return ret
	}
	return *o.DebugInfoCaptured
}

// GetDebugInfoCapturedOk returns a tuple with the DebugInfoCaptured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetDebugInfoCapturedOk() (*bool, bool) {
	if o == nil || o.DebugInfoCaptured == nil {
		return nil, false
	}
	return o.DebugInfoCaptured, true
}

// HasDebugInfoCaptured returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasDebugInfoCaptured() bool {
	return o != nil && o.DebugInfoCaptured != nil
}

// SetDebugInfoCaptured gets a reference to the given bool and assigns it to the DebugInfoCaptured field.
func (o *IssueSampleAttributes) SetDebugInfoCaptured(v bool) {
	o.DebugInfoCaptured = &v
}

// GetErrorSourceType returns the ErrorSourceType field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetErrorSourceType() string {
	if o == nil || o.ErrorSourceType == nil {
		var ret string
		return ret
	}
	return *o.ErrorSourceType
}

// GetErrorSourceTypeOk returns a tuple with the ErrorSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetErrorSourceTypeOk() (*string, bool) {
	if o == nil || o.ErrorSourceType == nil {
		return nil, false
	}
	return o.ErrorSourceType, true
}

// HasErrorSourceType returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasErrorSourceType() bool {
	return o != nil && o.ErrorSourceType != nil
}

// SetErrorSourceType gets a reference to the given string and assigns it to the ErrorSourceType field.
func (o *IssueSampleAttributes) SetErrorSourceType(v string) {
	o.ErrorSourceType = &v
}

// GetErrorStack returns the ErrorStack field value.
func (o *IssueSampleAttributes) GetErrorStack() IssueSampleErrorStack {
	if o == nil {
		var ret IssueSampleErrorStack
		return ret
	}
	return o.ErrorStack
}

// GetErrorStackOk returns a tuple with the ErrorStack field value
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetErrorStackOk() (*IssueSampleErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorStack, true
}

// SetErrorStack sets field value.
func (o *IssueSampleAttributes) SetErrorStack(v IssueSampleErrorStack) {
	o.ErrorStack = v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetEvent() map[string]interface{} {
	if o == nil || o.Event == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetEventOk() (*map[string]interface{}, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return &o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasEvent() bool {
	return o != nil && o.Event != nil
}

// SetEvent gets a reference to the given map[string]interface{} and assigns it to the Event field.
func (o *IssueSampleAttributes) SetEvent(v map[string]interface{}) {
	o.Event = v
}

// GetGitCommitSha returns the GitCommitSha field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetGitCommitSha() string {
	if o == nil || o.GitCommitSha == nil {
		var ret string
		return ret
	}
	return *o.GitCommitSha
}

// GetGitCommitShaOk returns a tuple with the GitCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetGitCommitShaOk() (*string, bool) {
	if o == nil || o.GitCommitSha == nil {
		return nil, false
	}
	return o.GitCommitSha, true
}

// HasGitCommitSha returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasGitCommitSha() bool {
	return o != nil && o.GitCommitSha != nil
}

// SetGitCommitSha gets a reference to the given string and assigns it to the GitCommitSha field.
func (o *IssueSampleAttributes) SetGitCommitSha(v string) {
	o.GitCommitSha = &v
}

// GetGitRepositoryId returns the GitRepositoryId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetGitRepositoryId() string {
	if o == nil || o.GitRepositoryId == nil {
		var ret string
		return ret
	}
	return *o.GitRepositoryId
}

// GetGitRepositoryIdOk returns a tuple with the GitRepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetGitRepositoryIdOk() (*string, bool) {
	if o == nil || o.GitRepositoryId == nil {
		return nil, false
	}
	return o.GitRepositoryId, true
}

// HasGitRepositoryId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasGitRepositoryId() bool {
	return o != nil && o.GitRepositoryId != nil
}

// SetGitRepositoryId gets a reference to the given string and assigns it to the GitRepositoryId field.
func (o *IssueSampleAttributes) SetGitRepositoryId(v string) {
	o.GitRepositoryId = &v
}

// GetGitRepositoryUrl returns the GitRepositoryUrl field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetGitRepositoryUrl() string {
	if o == nil || o.GitRepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.GitRepositoryUrl
}

// GetGitRepositoryUrlOk returns a tuple with the GitRepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetGitRepositoryUrlOk() (*string, bool) {
	if o == nil || o.GitRepositoryUrl == nil {
		return nil, false
	}
	return o.GitRepositoryUrl, true
}

// HasGitRepositoryUrl returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasGitRepositoryUrl() bool {
	return o != nil && o.GitRepositoryUrl != nil
}

// SetGitRepositoryUrl gets a reference to the given string and assigns it to the GitRepositoryUrl field.
func (o *IssueSampleAttributes) SetGitRepositoryUrl(v string) {
	o.GitRepositoryUrl = &v
}

// GetHasSessionReplay returns the HasSessionReplay field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetHasSessionReplay() bool {
	if o == nil || o.HasSessionReplay == nil {
		var ret bool
		return ret
	}
	return *o.HasSessionReplay
}

// GetHasSessionReplayOk returns a tuple with the HasSessionReplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetHasSessionReplayOk() (*bool, bool) {
	if o == nil || o.HasSessionReplay == nil {
		return nil, false
	}
	return o.HasSessionReplay, true
}

// HasHasSessionReplay returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasHasSessionReplay() bool {
	return o != nil && o.HasSessionReplay != nil
}

// SetHasSessionReplay gets a reference to the given bool and assigns it to the HasSessionReplay field.
func (o *IssueSampleAttributes) SetHasSessionReplay(v bool) {
	o.HasSessionReplay = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *IssueSampleAttributes) SetLanguage(v string) {
	o.Language = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetLinks() []IssueSampleFrameLink {
	if o == nil || o.Links == nil {
		var ret []IssueSampleFrameLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetLinksOk() (*[]IssueSampleFrameLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given []IssueSampleFrameLink and assigns it to the Links field.
func (o *IssueSampleAttributes) SetLinks(v []IssueSampleFrameLink) {
	o.Links = v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasOsName() bool {
	return o != nil && o.OsName != nil
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *IssueSampleAttributes) SetOsName(v string) {
	o.OsName = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *IssueSampleAttributes) SetService(v string) {
	o.Service = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasSessionId() bool {
	return o != nil && o.SessionId != nil
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *IssueSampleAttributes) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSessionReplayUrl returns the SessionReplayUrl field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetSessionReplayUrl() string {
	if o == nil || o.SessionReplayUrl == nil {
		var ret string
		return ret
	}
	return *o.SessionReplayUrl
}

// GetSessionReplayUrlOk returns a tuple with the SessionReplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetSessionReplayUrlOk() (*string, bool) {
	if o == nil || o.SessionReplayUrl == nil {
		return nil, false
	}
	return o.SessionReplayUrl, true
}

// HasSessionReplayUrl returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasSessionReplayUrl() bool {
	return o != nil && o.SessionReplayUrl != nil
}

// SetSessionReplayUrl gets a reference to the given string and assigns it to the SessionReplayUrl field.
func (o *IssueSampleAttributes) SetSessionReplayUrl(v string) {
	o.SessionReplayUrl = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasSnapshotId() bool {
	return o != nil && o.SnapshotId != nil
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *IssueSampleAttributes) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetSpanId() string {
	if o == nil || o.SpanId == nil {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetSpanIdOk() (*string, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasSpanId() bool {
	return o != nil && o.SpanId != nil
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *IssueSampleAttributes) SetSpanId(v string) {
	o.SpanId = &v
}

// GetTimestamp returns the Timestamp field value.
func (o *IssueSampleAttributes) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *IssueSampleAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasTraceId() bool {
	return o != nil && o.TraceId != nil
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *IssueSampleAttributes) SetTraceId(v string) {
	o.TraceId = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetVariant() string {
	if o == nil || o.Variant == nil {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetVariantOk() (*string, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasVariant() bool {
	return o != nil && o.Variant != nil
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *IssueSampleAttributes) SetVariant(v string) {
	o.Variant = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IssueSampleAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *IssueSampleAttributes) GetViewId() string {
	if o == nil || o.ViewId == nil {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSampleAttributes) GetViewIdOk() (*string, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *IssueSampleAttributes) HasViewId() bool {
	return o != nil && o.ViewId != nil
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *IssueSampleAttributes) SetViewId(v string) {
	o.ViewId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueSampleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.Biases != nil {
		toSerialize["biases"] = o.Biases
	}
	if o.BuildId != nil {
		toSerialize["build_id"] = o.BuildId
	}
	if o.BundleVersion != nil {
		toSerialize["bundle_version"] = o.BundleVersion
	}
	if o.DataSource != nil {
		toSerialize["data_source"] = o.DataSource
	}
	if o.DebugId != nil {
		toSerialize["debug_id"] = o.DebugId
	}
	if o.DebugInfoCaptured != nil {
		toSerialize["debug_info_captured"] = o.DebugInfoCaptured
	}
	if o.ErrorSourceType != nil {
		toSerialize["error_source_type"] = o.ErrorSourceType
	}
	toSerialize["error_stack"] = o.ErrorStack
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.GitCommitSha != nil {
		toSerialize["git_commit_sha"] = o.GitCommitSha
	}
	if o.GitRepositoryId != nil {
		toSerialize["git_repository_id"] = o.GitRepositoryId
	}
	if o.GitRepositoryUrl != nil {
		toSerialize["git_repository_url"] = o.GitRepositoryUrl
	}
	if o.HasSessionReplay != nil {
		toSerialize["has_session_replay"] = o.HasSessionReplay
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.OsName != nil {
		toSerialize["os_name"] = o.OsName
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	if o.SessionReplayUrl != nil {
		toSerialize["session_replay_url"] = o.SessionReplayUrl
	}
	if o.SnapshotId != nil {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	if o.SpanId != nil {
		toSerialize["span_id"] = o.SpanId
	}
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.TraceId != nil {
		toSerialize["trace_id"] = o.TraceId
	}
	if o.Variant != nil {
		toSerialize["variant"] = o.Variant
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueSampleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId     *string                `json:"application_id,omitempty"`
		Biases            []IssueSampleBias      `json:"biases,omitempty"`
		BuildId           *string                `json:"build_id,omitempty"`
		BundleVersion     *string                `json:"bundle_version,omitempty"`
		DataSource        *IssueSampleDataSource `json:"data_source,omitempty"`
		DebugId           *string                `json:"debug_id,omitempty"`
		DebugInfoCaptured *bool                  `json:"debug_info_captured,omitempty"`
		ErrorSourceType   *string                `json:"error_source_type,omitempty"`
		ErrorStack        *IssueSampleErrorStack `json:"error_stack"`
		Event             map[string]interface{} `json:"event,omitempty"`
		GitCommitSha      *string                `json:"git_commit_sha,omitempty"`
		GitRepositoryId   *string                `json:"git_repository_id,omitempty"`
		GitRepositoryUrl  *string                `json:"git_repository_url,omitempty"`
		HasSessionReplay  *bool                  `json:"has_session_replay,omitempty"`
		Language          *string                `json:"language,omitempty"`
		Links             []IssueSampleFrameLink `json:"links,omitempty"`
		OsName            *string                `json:"os_name,omitempty"`
		Service           *string                `json:"service,omitempty"`
		SessionId         *string                `json:"session_id,omitempty"`
		SessionReplayUrl  *string                `json:"session_replay_url,omitempty"`
		SnapshotId        *string                `json:"snapshot_id,omitempty"`
		SpanId            *string                `json:"span_id,omitempty"`
		Timestamp         *time.Time             `json:"timestamp"`
		TraceId           *string                `json:"trace_id,omitempty"`
		Variant           *string                `json:"variant,omitempty"`
		Version           *string                `json:"version,omitempty"`
		ViewId            *string                `json:"view_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ErrorStack == nil {
		return fmt.Errorf("required field error_stack missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "biases", "build_id", "bundle_version", "data_source", "debug_id", "debug_info_captured", "error_source_type", "error_stack", "event", "git_commit_sha", "git_repository_id", "git_repository_url", "has_session_replay", "language", "links", "os_name", "service", "session_id", "session_replay_url", "snapshot_id", "span_id", "timestamp", "trace_id", "variant", "version", "view_id"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = all.ApplicationId
	o.Biases = all.Biases
	o.BuildId = all.BuildId
	o.BundleVersion = all.BundleVersion
	if all.DataSource != nil && !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = all.DataSource
	}
	o.DebugId = all.DebugId
	o.DebugInfoCaptured = all.DebugInfoCaptured
	o.ErrorSourceType = all.ErrorSourceType
	if all.ErrorStack.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorStack = *all.ErrorStack
	o.Event = all.Event
	o.GitCommitSha = all.GitCommitSha
	o.GitRepositoryId = all.GitRepositoryId
	o.GitRepositoryUrl = all.GitRepositoryUrl
	o.HasSessionReplay = all.HasSessionReplay
	o.Language = all.Language
	o.Links = all.Links
	o.OsName = all.OsName
	o.Service = all.Service
	o.SessionId = all.SessionId
	o.SessionReplayUrl = all.SessionReplayUrl
	o.SnapshotId = all.SnapshotId
	o.SpanId = all.SpanId
	o.Timestamp = *all.Timestamp
	o.TraceId = all.TraceId
	o.Variant = all.Variant
	o.Version = all.Version
	o.ViewId = all.ViewId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
