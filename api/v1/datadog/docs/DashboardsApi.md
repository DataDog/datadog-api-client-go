# \DashboardsApi

All URIs are relative to *https://api.datadoghq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsApi.md#CreateDashboard) | **Post** /api/v1/dashboard | Create a new dashboard
[**DeleteDashboard**](DashboardsApi.md#DeleteDashboard) | **Delete** /api/v1/dashboard/{dashboard_id} | Delete a dashboard
[**GetDashboard**](DashboardsApi.md#GetDashboard) | **Get** /api/v1/dashboard/{dashboard_id} | Get a dashboard
[**ListDashboards**](DashboardsApi.md#ListDashboards) | **Get** /api/v1/dashboard | Get all dashboards
[**UpdateDashboard**](DashboardsApi.md#UpdateDashboard) | **Put** /api/v1/dashboard/{dashboard_id} | Update a dashboard



## CreateDashboard

> Dashboard CreateDashboard(ctx).Body(body).Execute()

Create a new dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := openapiclient.Dashboard{AuthorHandle: "AuthorHandle_example", CreatedAt: "TODO", Description: "Description_example", Id: "Id_example", IsReadOnly: false, LayoutType: openapiclient.DashboardLayoutType{}, ModifiedAt: "TODO", NotifyList: []string{"NotifyList_example"), TemplateVariablePresets: []DashboardTemplateVariablePreset{openapiclient.DashboardTemplateVariablePreset{Name: "Name_example", TemplateVariables: []DashboardTemplateVariablePresetValue{openapiclient.DashboardTemplateVariablePresetValue{Name: "Name_example", Value: "Value_example"})}), TemplateVariables: []DashboardTemplateVariables{openapiclient.Dashboard_template_variables{Default: "Default_example", Name: "Name_example", Prefix: "Prefix_example"}), Title: "Title_example", Url: "Url_example", Widgets: []Widget{openapiclient.Widget{Definition: openapiclient.WidgetDefinition{AlertId: "AlertId_example", Time: openapiclient.WidgetTime{LiveSpan: openapiclient.WidgetLiveSpan{}}, Title: "Title_example", TitleAlign: openapiclient.WidgetTextAlign{}, TitleSize: "TitleSize_example", Type: "Type_example", VizType: openapiclient.WidgetVizType{}, Precision: int64(123), TextAlign: openapiclient.WidgetTextAlign{}, Unit: "Unit_example", Requests: []ToplistWidgetRequest{openapiclient.ToplistWidgetRequest{ApmQuery: openapiclient.LogQueryDefinition{Compute: openapiclient.LogsQueryCompute{Aggregation: "Aggregation_example", Facet: "Facet_example", Interval: int64(123)}, GroupBy: []LogQueryDefinitionGroupBy{openapiclient.LogQueryDefinition_group_by{Facet: "Facet_example", Limit: int64(123), Sort: openapiclient.LogQueryDefinition_sort{Aggregation: "Aggregation_example", Facet: "Facet_example", Order: openapiclient.WidgetSort{}}}), Index: "Index_example", MultiCompute: []LogsQueryCompute{openapiclient.LogsQueryCompute{Aggregation: "Aggregation_example", Facet: "Facet_example", Interval: int64(123)}), Search: openapiclient.LogQueryDefinition_search{Query: "Query_example"}}, ConditionalFormats: []WidgetConditionalFormat{openapiclient.WidgetConditionalFormat{Comparator: openapiclient.WidgetComparator{}, CustomBgColor: "CustomBgColor_example", CustomFgColor: "CustomFgColor_example", HideValue: false, ImageUrl: "ImageUrl_example", Metric: "Metric_example", Palette: openapiclient.WidgetPalette{}, Timeframe: "Timeframe_example", Value: 123}), EventQuery: openapiclient.EventQueryDefinition{Search: "Search_example", TagsExecution: "TagsExecution_example"}, LogQuery: openapiclient.LogQueryDefinition{Compute: , GroupBy: []LogQueryDefinitionGroupBy{openapiclient.LogQueryDefinition_group_by{Facet: "Facet_example", Limit: int64(123), Sort: openapiclient.LogQueryDefinition_sort{Aggregation: "Aggregation_example", Facet: "Facet_example", Order: openapiclient.WidgetSort{}}}), Index: "Index_example", MultiCompute: []LogsQueryCompute{), Search: openapiclient.LogQueryDefinition_search{Query: "Query_example"}}, NetworkQuery: , ProcessQuery: openapiclient.ProcessQueryDefinition{FilterBy: []string{"FilterBy_example"), Limit: int64(123), Metric: "Metric_example", SearchBy: "SearchBy_example"}, Q: "Q_example", RumQuery: , Style: openapiclient.TimeseriesWidgetRequest_style{LineType: openapiclient.WidgetLineType{}, LineWidth: openapiclient.WidgetLineWidth{}, Palette: "Palette_example"}}), Check: "Check_example", Group: []string{"Group_example"), GroupBy: []string{"GroupBy_example"), Grouping: openapiclient.WidgetGrouping{}, Tags: []string{"Tags_example"), LegendSize: openapiclient.WidgetLegendSize{}, ShowLegend: false, EventSize: openapiclient.WidgetEventSize{}, Query: "Query_example", TagsExecution: "TagsExecution_example", Color: "Color_example", FontSize: "FontSize_example", Text: "Text_example", LayoutType: openapiclient.WidgetLayoutType{}, Widgets: []Widget{openapiclient.Widget{Definition: openapiclient.WidgetDefinition{AlertId: "AlertId_example", Time: openapiclient.WidgetTime{LiveSpan: openapiclient.WidgetLiveSpan{}}, Title: "Title_example", TitleAlign: , TitleSize: "TitleSize_example", Type: "Type_example", VizType: openapiclient.WidgetVizType{}, Precision: int64(123), TextAlign: , Unit: "Unit_example", Requests: []ToplistWidgetRequest{openapiclient.ToplistWidgetRequest{ApmQuery: , ConditionalFormats: []WidgetConditionalFormat{openapiclient.WidgetConditionalFormat{Comparator: openapiclient.WidgetComparator{}, CustomBgColor: "CustomBgColor_example", CustomFgColor: "CustomFgColor_example", HideValue: false, ImageUrl: "ImageUrl_example", Metric: "Metric_example", Palette: openapiclient.WidgetPalette{}, Timeframe: "Timeframe_example", Value: 123}), EventQuery: openapiclient.EventQueryDefinition{Search: "Search_example", TagsExecution: "TagsExecution_example"}, LogQuery: , NetworkQuery: , ProcessQuery: openapiclient.ProcessQueryDefinition{FilterBy: []string{"FilterBy_example"), Limit: int64(123), Metric: "Metric_example", SearchBy: "SearchBy_example"}, Q: "Q_example", RumQuery: , Style: openapiclient.TimeseriesWidgetRequest_style{LineType: openapiclient.WidgetLineType{}, LineWidth: openapiclient.WidgetLineWidth{}, Palette: "Palette_example"}}), Check: "Check_example", Group: []string{"Group_example"), GroupBy: []string{"GroupBy_example"), Grouping: openapiclient.WidgetGrouping{}, Tags: []string{"Tags_example"), LegendSize: openapiclient.WidgetLegendSize{}, ShowLegend: false, EventSize: openapiclient.WidgetEventSize{}, Query: "Query_example", TagsExecution: "TagsExecution_example", Color: "Color_example", FontSize: "FontSize_example", Text: "Text_example", LayoutType: openapiclient.WidgetLayoutType{}, Widgets: []Widget{), Events: []WidgetEvent{openapiclient.WidgetEvent{Q: "Q_example", TagsExecution: "TagsExecution_example"}), Yaxis: openapiclient.WidgetAxis{IncludeZero: false, Label: "Label_example", Max: "Max_example", Min: "Min_example", Scale: "Scale_example"}, NoGroupHosts: false, NoMetricHosts: false, NodeType: openapiclient.WidgetNodeType{}, Notes: "Notes_example", Scope: []string{"Scope_example"), Style: openapiclient.HostMapWidgetDefinition_style{FillMax: "FillMax_example", FillMin: "FillMin_example", Palette: "Palette_example", PaletteFlip: false}, Url: "Url_example", Margin: openapiclient.WidgetMargin{}, Sizing: openapiclient.WidgetImageSizing{}, Columns: []string{"Columns_example"), Indexes: []string{"Indexes_example"), MessageDisplay: openapiclient.WidgetMessageDisplay{}, ShowDateColumn: false, ShowMessageColumn: false, Sort: , ColorPreference: openapiclient.WidgetColorPreference{}, DisplayFormat: openapiclient.WidgetServiceSummaryDisplayFormat{}, HideZeroCounts: false, ShowLastTriggered: false, SummaryType: openapiclient.WidgetSummaryType{}, BackgroundColor: "BackgroundColor_example", Content: "Content_example", ShowTick: false, TickEdge: openapiclient.WidgetTickEdge{}, TickPos: "TickPos_example", Autoscale: false, CustomUnit: "CustomUnit_example", ColorByGroups: []string{"ColorByGroups_example"), Xaxis: openapiclient.WidgetAxis{IncludeZero: false, Label: "Label_example", Max: "Max_example", Min: "Min_example", Scale: "Scale_example"}, ShowErrorBudget: false, SloId: "SloId_example", TimeWindows: []WidgetTimeWindows{openapiclient.WidgetTimeWindows{}), ViewMode: openapiclient.WidgetViewMode{}, ViewType: "ViewType_example", Filters: []string{"Filters_example"), Service: "Service_example", Env: "Env_example", ShowBreakdown: false, ShowDistribution: false, ShowErrors: false, ShowHits: false, ShowLatency: false, ShowResourceList: false, SizeFormat: openapiclient.WidgetSizeFormat{}, SpanName: "SpanName_example", Markers: []WidgetMarker{openapiclient.WidgetMarker{DisplayType: "DisplayType_example", Label: "Label_example", Time: "Time_example", Value: "Value_example"})}, Id: int64(123), Layout: openapiclient.WidgetLayout{Height: int64(123), Width: int64(123), X: int64(123), Y: int64(123)}}), Events: []WidgetEvent{openapiclient.WidgetEvent{Q: "Q_example", TagsExecution: "TagsExecution_example"}), Yaxis: , NoGroupHosts: false, NoMetricHosts: false, NodeType: openapiclient.WidgetNodeType{}, Notes: "Notes_example", Scope: []string{"Scope_example"), Style: openapiclient.HostMapWidgetDefinition_style{FillMax: "FillMax_example", FillMin: "FillMin_example", Palette: "Palette_example", PaletteFlip: false}, Url: "Url_example", Margin: openapiclient.WidgetMargin{}, Sizing: openapiclient.WidgetImageSizing{}, Columns: []string{"Columns_example"), Indexes: []string{"Indexes_example"), MessageDisplay: openapiclient.WidgetMessageDisplay{}, ShowDateColumn: false, ShowMessageColumn: false, Sort: , ColorPreference: openapiclient.WidgetColorPreference{}, DisplayFormat: openapiclient.WidgetServiceSummaryDisplayFormat{}, HideZeroCounts: false, ShowLastTriggered: false, SummaryType: openapiclient.WidgetSummaryType{}, BackgroundColor: "BackgroundColor_example", Content: "Content_example", ShowTick: false, TickEdge: openapiclient.WidgetTickEdge{}, TickPos: "TickPos_example", Autoscale: false, CustomUnit: "CustomUnit_example", ColorByGroups: []string{"ColorByGroups_example"), Xaxis: , ShowErrorBudget: false, SloId: "SloId_example", TimeWindows: []WidgetTimeWindows{openapiclient.WidgetTimeWindows{}), ViewMode: openapiclient.WidgetViewMode{}, ViewType: "ViewType_example", Filters: []string{"Filters_example"), Service: "Service_example", Env: "Env_example", ShowBreakdown: false, ShowDistribution: false, ShowErrors: false, ShowHits: false, ShowLatency: false, ShowResourceList: false, SizeFormat: openapiclient.WidgetSizeFormat{}, SpanName: "SpanName_example", Markers: []WidgetMarker{openapiclient.WidgetMarker{DisplayType: "DisplayType_example", Label: "Label_example", Time: "Time_example", Value: "Value_example"})}, Id: int64(123), Layout: openapiclient.WidgetLayout{Height: int64(123), Width: int64(123), X: int64(123), Y: int64(123)}})} // Dashboard | Create a dashboard request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.CreateDashboard(context.Background(), body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Dashboard**](Dashboard.md) | Create a dashboard request body. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DashboardDeleteResponse DeleteDashboard(ctx, dashboardId).Execute()

Delete a dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dashboardId := "dashboardId_example" // string | The ID of the dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.DeleteDashboard(context.Background(), dashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDashboard`: DashboardDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.DeleteDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardDeleteResponse**](DashboardDeleteResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> Dashboard GetDashboard(ctx, dashboardId).Execute()

Get a dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dashboardId := "dashboardId_example" // string | The ID of the dashboard.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.GetDashboard(context.Background(), dashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDashboards

> DashboardSummary ListDashboards(ctx).Execute()

Get all dashboards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.ListDashboards(context.Background(), ).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.ListDashboards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDashboards`: DashboardSummary
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.ListDashboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDashboardsRequest struct via the builder pattern


### Return type

[**DashboardSummary**](DashboardSummary.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, dashboardId).Body(body).Execute()

Update a dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dashboardId := "dashboardId_example" // string | The ID of the dashboard.
    body := openapiclient.Dashboard{AuthorHandle: "AuthorHandle_example", CreatedAt: "TODO", Description: "Description_example", Id: "Id_example", IsReadOnly: false, LayoutType: openapiclient.DashboardLayoutType{}, ModifiedAt: "TODO", NotifyList: []string{"NotifyList_example"), TemplateVariablePresets: []DashboardTemplateVariablePreset{openapiclient.DashboardTemplateVariablePreset{Name: "Name_example", TemplateVariables: []DashboardTemplateVariablePresetValue{openapiclient.DashboardTemplateVariablePresetValue{Name: "Name_example", Value: "Value_example"})}), TemplateVariables: []DashboardTemplateVariables{openapiclient.Dashboard_template_variables{Default: "Default_example", Name: "Name_example", Prefix: "Prefix_example"}), Title: "Title_example", Url: "Url_example", Widgets: []Widget{)} // Dashboard | Update Dashboard request body.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.UpdateDashboard(context.Background(), dashboardId, body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.UpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Dashboard**](Dashboard.md) | Update Dashboard request body. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [appKeyAuth](../README.md#appKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

