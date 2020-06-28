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
import (
	"time"
)
// Dashboard A dashboard is Datadog’s tool for visually tracking, analyzing, and displaying key performance metrics, which enable you to monitor the health of your infrastructure.
type Dashboard struct {
	// Identifier of the dashboard author.
	AuthorHandle string `json:"author_handle,omitempty"`
	// Creation date of the dashboard.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Description of the dashboard.
	Description *string `json:"description,omitempty"`
	// ID of the dashboard.
	Id string `json:"id,omitempty"`
	// Whether this dashboard is read-only. If True, only the author and admins can make changes to it.
	IsReadOnly bool `json:"is_read_only,omitempty"`
	LayoutType DashboardLayoutType `json:"layout_type"`
	// Modification date of the dashboard.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// List of handles of users to notify when changes are made to this dashboard.
	NotifyList *[]string `json:"notify_list,omitempty"`
	// Array of template variables saved views.
	TemplateVariablePresets *[]DashboardTemplateVariablePreset `json:"template_variable_presets,omitempty"`
	// List of template variables for this dashboard.
	TemplateVariables *[]DashboardTemplateVariables `json:"template_variables,omitempty"`
	// Title of the dashboard.
	Title string `json:"title"`
	// The URL of the dashboard.
	Url string `json:"url,omitempty"`
	// List of widgets to display on the dashboard.
	Widgets []Widget `json:"widgets"`
}
