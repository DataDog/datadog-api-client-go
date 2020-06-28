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
// IpRanges IP ranges.
type IpRanges struct {
	Agents IpPrefixesAgents `json:"agents,omitempty"`
	Api IpPrefixesApi `json:"api,omitempty"`
	Apm IpPrefixesApm `json:"apm,omitempty"`
	Logs IpPrefixesLogs `json:"logs,omitempty"`
	// Date when last updated, in the form `YYYY-MM-DD-hh-mm-ss`.
	Modified string `json:"modified,omitempty"`
	Process IpPrefixesProcess `json:"process,omitempty"`
	Synthetics IpPrefixesSynthetics `json:"synthetics,omitempty"`
	// Version of the IP list.
	Version int64 `json:"version,omitempty"`
	Webhooks IpPrefixesWebhooks `json:"webhooks,omitempty"`
}
