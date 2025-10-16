// Create App returns "Created" response

package main


import (
	"context"
	"fmt"
	"os"

    "github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"
)

func main() {
	body := datadogV2.CreateAppRequest{
Data: &datadogV2.CreateAppRequestData{
Type: datadogV2.APPDEFINITIONTYPE_APPDEFINITIONS,
Attributes: &datadogV2.CreateAppRequestDataAttributes{
RootInstanceName: datadog.PtrString("grid0"),
Components: []datadogV2.ComponentGrid{
{
Name: "grid0",
Type: datadogV2.COMPONENTGRIDTYPE_GRID,
Properties: datadogV2.ComponentGridProperties{
Children: []datadogV2.Component{
{
Type: datadogV2.COMPONENTTYPE_GRIDCELL,
Name: "gridCell0",
Properties: datadogV2.ComponentProperties{
Children: []datadogV2.Component{
{
Name: "text0",
Type: datadogV2.COMPONENTTYPE_TEXT,
Properties: datadogV2.ComponentProperties{
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
Bool: datadog.PtrBool(true)},
},
Events: []datadogV2.AppBuilderEvent{
},
},
},
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
String: datadog.PtrString("true")},
},
Events: []datadogV2.AppBuilderEvent{
},
},
{
Type: datadogV2.COMPONENTTYPE_GRIDCELL,
Name: "gridCell2",
Properties: datadogV2.ComponentProperties{
Children: []datadogV2.Component{
{
Name: "table0",
Type: datadogV2.COMPONENTTYPE_TABLE,
Properties: datadogV2.ComponentProperties{
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
Bool: datadog.PtrBool(true)},
},
Events: []datadogV2.AppBuilderEvent{
},
},
},
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
String: datadog.PtrString("true")},
},
Events: []datadogV2.AppBuilderEvent{
},
},
{
Type: datadogV2.COMPONENTTYPE_GRIDCELL,
Name: "gridCell1",
Properties: datadogV2.ComponentProperties{
Children: []datadogV2.Component{
{
Name: "text1",
Type: datadogV2.COMPONENTTYPE_TEXT,
Properties: datadogV2.ComponentProperties{
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
Bool: datadog.PtrBool(true)},
},
Events: []datadogV2.AppBuilderEvent{
},
},
},
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
String: datadog.PtrString("true")},
},
Events: []datadogV2.AppBuilderEvent{
},
},
{
Type: datadogV2.COMPONENTTYPE_GRIDCELL,
Name: "gridCell3",
Properties: datadogV2.ComponentProperties{
Children: []datadogV2.Component{
{
Name: "button0",
Type: datadogV2.COMPONENTTYPE_BUTTON,
Properties: datadogV2.ComponentProperties{
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
Bool: datadog.PtrBool(true)},
},
Events: []datadogV2.AppBuilderEvent{
{
Name: datadogV2.APPBUILDEREVENTNAME_CLICK.Ptr(),
Type: datadogV2.APPBUILDEREVENTTYPE_SETSTATEVARIABLEVALUE.Ptr(),
},
},
},
},
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
String: datadog.PtrString("true")},
},
Events: []datadogV2.AppBuilderEvent{
},
},
{
Type: datadogV2.COMPONENTTYPE_GRIDCELL,
Name: "gridCell4",
Properties: datadogV2.ComponentProperties{
Children: []datadogV2.Component{
{
Name: "button1",
Type: datadogV2.COMPONENTTYPE_BUTTON,
Properties: datadogV2.ComponentProperties{
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
Bool: datadog.PtrBool(true)},
},
Events: []datadogV2.AppBuilderEvent{
{
Name: datadogV2.APPBUILDEREVENTNAME_CLICK.Ptr(),
Type: datadogV2.APPBUILDEREVENTTYPE_SETSTATEVARIABLEVALUE.Ptr(),
},
},
},
},
IsVisible: &datadogV2.ComponentPropertiesIsVisible{
String: datadog.PtrString("true")},
},
Events: []datadogV2.AppBuilderEvent{
},
},
},
BackgroundColor: datadog.PtrString("default"),
},
Events: []datadogV2.AppBuilderEvent{
},
},
},
Queries: []datadogV2.Query{
datadogV2.Query{
ActionQuery: &datadogV2.ActionQuery{
Id: uuid.MustParse("92ff0bb8-553b-4f31-87c7-ef5bd16d47d5"),
Type: datadogV2.ACTIONQUERYTYPE_ACTION,
Name: "fetchFacts",
Events: []datadogV2.AppBuilderEvent{
},
Properties: datadogV2.ActionQueryProperties{
Spec: datadogV2.ActionQuerySpec{
ActionQuerySpecObject: &datadogV2.ActionQuerySpecObject{
Fqn: "com.datadoghq.http.request",
ConnectionId: datadog.PtrString("5e63f4a8-4ce6-47de-ba11-f6617c1d54f3"),
Inputs: &datadogV2.ActionQuerySpecInputs{
ActionQuerySpecInput: map[string]interface{}{
"verb": "GET",
"url": "https://catfact.ninja/facts",
"urlParams": "[{'key': 'limit', 'value': '${pageSize.value.toString()}'}, {'key': 'page', 'value': '${(table0.pageIndex + 1).toString()}'}]",
}},
}},
},
}},
datadogV2.Query{
StateVariable: &datadogV2.StateVariable{
Type: datadogV2.STATEVARIABLETYPE_STATEVARIABLE,
Name: "pageSize",
Properties: datadogV2.StateVariableProperties{
DefaultValue: "${20}",
},
Id: uuid.MustParse("afd03c81-4075-4432-8618-ba09d52d2f2d"),
}},
datadogV2.Query{
DataTransform: &datadogV2.DataTransform{
Type: datadogV2.DATATRANSFORMTYPE_DATATRANSFORM,
Name: "randomFact",
Properties: datadogV2.DataTransformProperties{
Outputs: datadog.PtrString(`${(() => {const facts = fetchFacts.outputs.body.data
return facts[Math.floor(Math.random()*facts.length)]
})()}`),
},
Id: uuid.MustParse("0fb22859-47dc-4137-9e41-7b67d04c525c"),
}},
},
Name: datadog.PtrString("Example Cat Facts Viewer"),
Description: datadog.PtrString("This is a slightly complicated example app that fetches and displays cat facts"),
},
},
}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewAppBuilderApi(apiClient)
	resp, r, err := api.CreateApp(ctx, body, )

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppBuilderApi.CreateApp`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AppBuilderApi.CreateApp`:\n%s\n", responseContent)
}