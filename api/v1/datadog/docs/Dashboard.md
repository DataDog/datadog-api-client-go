# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorHandle** | **string** | Identifier of the dashboard author. | [optional] [readonly] 
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date of the dashboard. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the dashboard. | [optional] 
**Id** | **string** | ID of the dashboard. | [optional] [readonly] 
**IsReadOnly** | **bool** | Whether this dashboard is read-only. If True, only the author and admins can make changes to it. | [optional] [default to false]
**LayoutType** | [**DashboardLayoutType**](DashboardLayoutType.md) |  | 
**ModifiedAt** | [**time.Time**](time.Time.md) | Modification date of the dashboard. | [optional] [readonly] 
**NotifyList** | Pointer to **[]string** | List of handles of users to notify when changes are made to this dashboard. | [optional] 
**TemplateVariablePresets** | Pointer to [**[]DashboardTemplateVariablePreset**](DashboardTemplateVariablePreset.md) | Array of template variables saved views. | [optional] 
**TemplateVariables** | Pointer to [**[]DashboardTemplateVariables**](Dashboard_template_variables.md) | List of template variables for this dashboard. | [optional] 
**Title** | **string** | Title of the dashboard. | 
**Url** | **string** | The URL of the dashboard. | [optional] [readonly] 
**Widgets** | [**[]Widget**](Widget.md) | List of widgets to display on the dashboard. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


