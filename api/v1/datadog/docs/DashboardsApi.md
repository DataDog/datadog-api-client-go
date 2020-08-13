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
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    body := datadog.Dashboard{AuthorHandle: "AuthorHandle_example", CreatedAt: "TODO", Description: "Description_example", Id: "Id_example", IsReadOnly: false, LayoutType: datadog.DashboardLayoutType{}, ModifiedAt: "TODO", NotifyList: []string{"NotifyList_example"), TemplateVariablePresets: []DashboardTemplateVariablePreset{datadog.DashboardTemplateVariablePreset{Name: "Name_example", TemplateVariables: []DashboardTemplateVariablePresetValue{datadog.DashboardTemplateVariablePresetValue{Name: "Name_example", Value: "Value_example"})}), TemplateVariables: []DashboardTemplateVariables{datadog.Dashboard_template_variables{Default: "Default_example", Name: "Name_example", Prefix: "Prefix_example"}), Title: "Title_example", Url: "Url_example", Widgets: []Widget{datadog.Widget{Definition: datadog.WidgetDefinition{AlertId: "AlertId_example", Time: datadog.WidgetTime{LiveSpan: datadog.WidgetLiveSpan{}}, Title: "Title_example", TitleAlign: datadog.WidgetTextAlign{}, TitleSize: "TitleSize_example", Type: datadog.ToplistWidgetDefinitionType{}, VizType: datadog.WidgetVizType{}, Precision: int64(123), TextAlign: datadog.WidgetTextAlign{}, Unit: "Unit_example", Requests: []ToplistWidgetRequest{datadog.ToplistWidgetRequest{ApmQuery: datadog.LogQueryDefinition{Compute: datadog.LogsQueryCompute{Aggregation: "Aggregation_example", Facet: "Facet_example", Interval: int64(123)}, GroupBy: []LogQueryDefinitionGroupBy{datadog.LogQueryDefinition_group_by{Facet: "Facet_example", Limit: int64(123), Sort: datadog.LogQueryDefinition_sort{Aggregation: "Aggregation_example", Facet: "Facet_example", Order: datadog.WidgetSort{}}}), Index: "Index_example", MultiCompute: []LogsQueryCompute{datadog.LogsQueryCompute{Aggregation: "Aggregation_example", Facet: "Facet_example", Interval: int64(123)}), Search: datadog.LogQueryDefinition_search{Query: "Query_example"}}, ConditionalFormats: []WidgetConditionalFormat{datadog.WidgetConditionalFormat{Comparator: datadog.WidgetComparator{}, CustomBgColor: "CustomBgColor_example", CustomFgColor: "CustomFgColor_example", HideValue: false, ImageUrl: "ImageUrl_example", Metric: "Metric_example", Palette: datadog.WidgetPalette{}, Timeframe: "Timeframe_example", Value: 123}), EventQuery: datadog.EventQueryDefinition{Search: "Search_example", TagsExecution: "TagsExecution_example"}, LogQuery: datadog.LogQueryDefinition{Compute: , GroupBy: []LogQueryDefinitionGroupBy{datadog.LogQueryDefinition_group_by{Facet: "Facet_example", Limit: int64(123), Sort: datadog.LogQueryDefinition_sort{Aggregation: "Aggregation_example", Facet: "Facet_example", Order: datadog.WidgetSort{}}}), Index: "Index_example", MultiCompute: []LogsQueryCompute{), Search: datadog.LogQueryDefinition_search{Query: "Query_example"}}, NetworkQuery: , ProcessQuery: datadog.ProcessQueryDefinition{FilterBy: []string{"FilterBy_example"), Limit: int64(123), Metric: "Metric_example", SearchBy: "SearchBy_example"}, Q: "Q_example", RumQuery: , SecurityQuery: , Style: datadog.WidgetRequestStyle{LineType: datadog.WidgetLineType{}, LineWidth: datadog.WidgetLineWidth{}, Palette: "Palette_example"}}), Check: "Check_example", Group: []string{"Group_example"), GroupBy: []string{"GroupBy_example"), Grouping: datadog.WidgetGrouping{}, Tags: []string{"Tags_example"), LegendSize: "LegendSize_example", ShowLegend: false, EventSize: datadog.WidgetEventSize{}, Query: "Query_example", TagsExecution: "TagsExecution_example", Color: "Color_example", FontSize: "FontSize_example", Text: "Text_example", LayoutType: datadog.WidgetLayoutType{}, Widgets: []Widget{datadog.Widget{Definition: datadog.WidgetDefinition{AlertId: "AlertId_example", Time: datadog.WidgetTime{LiveSpan: datadog.WidgetLiveSpan{}}, Title: "Title_example", TitleAlign: , TitleSize: "TitleSize_example", Type: datadog.ToplistWidgetDefinitionType{}, VizType: datadog.WidgetVizType{}, Precision: int64(123), TextAlign: , Unit: "Unit_example", Requests: []ToplistWidgetRequest{datadog.ToplistWidgetRequest{ApmQuery: , ConditionalFormats: []WidgetConditionalFormat{datadog.WidgetConditionalFormat{Comparator: datadog.WidgetComparator{}, CustomBgColor: "CustomBgColor_example", CustomFgColor: "CustomFgColor_example", HideValue: false, ImageUrl: "ImageUrl_example", Metric: "Metric_example", Palette: datadog.WidgetPalette{}, Timeframe: "Timeframe_example", Value: 123}), EventQuery: datadog.EventQueryDefinition{Search: "Search_example", TagsExecution: "TagsExecution_example"}, LogQuery: , NetworkQuery: , ProcessQuery: datadog.ProcessQueryDefinition{FilterBy: []string{"FilterBy_example"), Limit: int64(123), Metric: "Metric_example", SearchBy: "SearchBy_example"}, Q: "Q_example", RumQuery: , SecurityQuery: , Style: datadog.WidgetRequestStyle{LineType: datadog.WidgetLineType{}, LineWidth: datadog.WidgetLineWidth{}, Palette: "Palette_example"}}), Check: "Check_example", Group: []string{"Group_example"), GroupBy: []string{"GroupBy_example"), Grouping: datadog.WidgetGrouping{}, Tags: []string{"Tags_example"), LegendSize: "LegendSize_example", ShowLegend: false, EventSize: datadog.WidgetEventSize{}, Query: "Query_example", TagsExecution: "TagsExecution_example", Color: "Color_example", FontSize: "FontSize_example", Text: "Text_example", LayoutType: datadog.WidgetLayoutType{}, Widgets: []Widget{), Events: []WidgetEvent{datadog.WidgetEvent{Q: "Q_example", TagsExecution: "TagsExecution_example"}), Yaxis: datadog.WidgetAxis{IncludeZero: false, Label: "Label_example", Max: "Max_example", Min: "Min_example", Scale: "Scale_example"}, NoGroupHosts: false, NoMetricHosts: false, NodeType: datadog.WidgetNodeType{}, Notes: "Notes_example", Scope: []string{"Scope_example"), Style: datadog.HostMapWidgetDefinition_style{FillMax: "FillMax_example", FillMin: "FillMin_example", Palette: "Palette_example", PaletteFlip: false}, Url: "Url_example", Margin: datadog.WidgetMargin{}, Sizing: datadog.WidgetImageSizing{}, Columns: []string{"Columns_example"), Indexes: []string{"Indexes_example"), Logset: "Logset_example", MessageDisplay: datadog.WidgetMessageDisplay{}, ShowDateColumn: false, ShowMessageColumn: false, Sort: datadog.WidgetMonitorSummarySort{}, ColorPreference: datadog.WidgetColorPreference{}, Count: int64(123), DisplayFormat: datadog.WidgetServiceSummaryDisplayFormat{}, HideZeroCounts: false, ShowLastTriggered: false, Start: int64(123), SummaryType: datadog.WidgetSummaryType{}, BackgroundColor: "BackgroundColor_example", Content: "Content_example", ShowTick: false, TickEdge: datadog.WidgetTickEdge{}, TickPos: "TickPos_example", Autoscale: false, CustomUnit: "CustomUnit_example", ColorByGroups: []string{"ColorByGroups_example"), Xaxis: datadog.WidgetAxis{IncludeZero: false, Label: "Label_example", Max: "Max_example", Min: "Min_example", Scale: "Scale_example"}, ShowErrorBudget: false, SloId: "SloId_example", TimeWindows: []WidgetTimeWindows{datadog.WidgetTimeWindows{}), ViewMode: datadog.WidgetViewMode{}, ViewType: "ViewType_example", Filters: []string{"Filters_example"), Service: "Service_example", Env: "Env_example", ShowBreakdown: false, ShowDistribution: false, ShowErrors: false, ShowHits: false, ShowLatency: false, ShowResourceList: false, SizeFormat: datadog.WidgetSizeFormat{}, SpanName: "SpanName_example", Markers: []WidgetMarker{datadog.WidgetMarker{DisplayType: "DisplayType_example", Label: "Label_example", Time: "Time_example", Value: "Value_example"})}, Id: int64(123), Layout: datadog.WidgetLayout{Height: int64(123), Width: int64(123), X: int64(123), Y: int64(123)}}), Events: []WidgetEvent{datadog.WidgetEvent{Q: "Q_example", TagsExecution: "TagsExecution_example"}), Yaxis: , NoGroupHosts: false, NoMetricHosts: false, NodeType: datadog.WidgetNodeType{}, Notes: "Notes_example", Scope: []string{"Scope_example"), Style: datadog.HostMapWidgetDefinition_style{FillMax: "FillMax_example", FillMin: "FillMin_example", Palette: "Palette_example", PaletteFlip: false}, Url: "Url_example", Margin: datadog.WidgetMargin{}, Sizing: datadog.WidgetImageSizing{}, Columns: []string{"Columns_example"), Indexes: []string{"Indexes_example"), Logset: "Logset_example", MessageDisplay: datadog.WidgetMessageDisplay{}, ShowDateColumn: false, ShowMessageColumn: false, Sort: datadog.WidgetMonitorSummarySort{}, ColorPreference: datadog.WidgetColorPreference{}, Count: int64(123), DisplayFormat: datadog.WidgetServiceSummaryDisplayFormat{}, HideZeroCounts: false, ShowLastTriggered: false, Start: int64(123), SummaryType: datadog.WidgetSummaryType{}, BackgroundColor: "BackgroundColor_example", Content: "Content_example", ShowTick: false, TickEdge: datadog.WidgetTickEdge{}, TickPos: "TickPos_example", Autoscale: false, CustomUnit: "CustomUnit_example", ColorByGroups: []string{"ColorByGroups_example"), Xaxis: , ShowErrorBudget: false, SloId: "SloId_example", TimeWindows: []WidgetTimeWindows{datadog.WidgetTimeWindows{}), ViewMode: datadog.WidgetViewMode{}, ViewType: "ViewType_example", Filters: []string{"Filters_example"), Service: "Service_example", Env: "Env_example", ShowBreakdown: false, ShowDistribution: false, ShowErrors: false, ShowHits: false, ShowLatency: false, ShowResourceList: false, SizeFormat: datadog.WidgetSizeFormat{}, SpanName: "SpanName_example", Markers: []WidgetMarker{datadog.WidgetMarker{DisplayType: "DisplayType_example", Label: "Label_example", Time: "Time_example", Value: "Value_example"})}, Id: int64(123), Layout: datadog.WidgetLayout{Height: int64(123), Width: int64(123), X: int64(123), Y: int64(123)}})} // Dashboard | Create a dashboard request body.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.CreateDashboard(ctx, body).Execute()
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
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    dashboardId := "dashboardId_example" // string | The ID of the dashboard.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.DeleteDashboard(ctx, dashboardId).Execute()
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
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    dashboardId := "dashboardId_example" // string | The ID of the dashboard.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.GetDashboard(ctx, dashboardId).Execute()
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
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )


    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.ListDashboards(ctx).Execute()
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
    datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        datadog.ContextAPIKeys,
        map[string]datadog.APIKey{
            "apiKeyAuth": {
                Key: os.Getenv("DD_CLIENT_API_KEY"),
            },
            "appKeyAuth": {
                Key: os.Getenv("DD_CLIENT_APP_KEY"),
            },
        },
    )

    dashboardId := "dashboardId_example" // string | The ID of the dashboard.
    body := datadog.Dashboard{AuthorHandle: "AuthorHandle_example", CreatedAt: "TODO", Description: "Description_example", Id: "Id_example", IsReadOnly: false, LayoutType: datadog.DashboardLayoutType{}, ModifiedAt: "TODO", NotifyList: []string{"NotifyList_example"), TemplateVariablePresets: []DashboardTemplateVariablePreset{datadog.DashboardTemplateVariablePreset{Name: "Name_example", TemplateVariables: []DashboardTemplateVariablePresetValue{datadog.DashboardTemplateVariablePresetValue{Name: "Name_example", Value: "Value_example"})}), TemplateVariables: []DashboardTemplateVariables{datadog.Dashboard_template_variables{Default: "Default_example", Name: "Name_example", Prefix: "Prefix_example"}), Title: "Title_example", Url: "Url_example", Widgets: []Widget{)} // Dashboard | Update Dashboard request body.

    configuration := datadog.NewConfiguration()
    api_client := datadog.NewAPIClient(configuration)
    resp, r, err := api_client.DashboardsApi.UpdateDashboard(ctx, dashboardId, body).Execute()
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

