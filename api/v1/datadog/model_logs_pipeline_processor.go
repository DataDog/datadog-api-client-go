/*
 * Datadog API Collection
 *
 * Collection of all Datadog Public endpoints.
 *
 * API version: 1.0
 * Contact: support@datadoghq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog
// LogsPipelineProcessor Nested Pipelines are pipelines within a pipeline. Use Nested Pipelines to split the processing into two steps. For example, first use a high-level filtering such as team and then a second level of filtering based on the integration, service, or any other tag or attribute.  A pipeline can contain Nested Pipelines and Processors whereas a Nested Pipeline can only contain Processors.
type LogsPipelineProcessor struct {
	Filter LogsFilter `json:"filter,omitempty"`
	// Whether or not the processor is enabled.
	IsEnabled bool `json:"is_enabled,omitempty"`
	// Name of the processor.
	Name string `json:"name,omitempty"`
	// Ordered list of processors in this pipeline.
	Processors []LogsProcessor `json:"processors,omitempty"`
	Type LogsPipelineProcessorType `json:"type"`
}
