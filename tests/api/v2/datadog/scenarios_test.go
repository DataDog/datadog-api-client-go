package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	lookup "github.com/mcuadros/go-lookup"
)

type ctxKey struct{}
type apiClientKey struct{}
type apiKey struct{}
type requestKey struct{}
type requestArgsKey struct{}
type requestParamsKey struct{}
type responseKey struct{}
type dataKey struct{}
type bodyKey struct{}
type cleanupKey struct{}

var steps = make(map[string]interface{})

func AddStep(expr string, step interface{}) {
	steps[expr] = step
}

func TestScenarios(t *testing.T) {
	s := gobdd.NewSuite(
		t,
		gobdd.WithIgnoredTags([]string{"@todo"}),
		gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
			ctx.Set(
				ctxKey{},
				context.WithValue(
					context.Background(),
					datadog.ContextAPIKeys,
					map[string]datadog.APIKey{},
				),
			)
			ctx.Set(
				apiClientKey{},
				datadog.NewAPIClient(NewConfiguration()),
			)
			ctx.Set(
				dataKey{},
				make(map[string]interface{}),
			)
		}), gobdd.WithAfterScenario(func(ctx gobdd.Context) {
			c, err := ctx.Get(cleanupKey{})
			if err == nil {
				c.(func())()
			}
		}),
	)

	for expr, step := range steps {
		s.AddStep(expr, step)
	}

	s.AddStep(`a valid "apiKeyAuth" key`, aValidAPIKeyAuth)
	s.AddStep(`a valid "appKeyAuth" key`, aValidAppKeyAuth)
	s.AddStep(`an instance of "([^"]+)" API`, anInstanceOf)
	s.AddStep(`new "([^"]+)" request`, newRequest)
	s.AddStep(`parameter ([^ ]+) from (.*)`, addParameterFrom)
	s.AddStep(`the request is sent`, requestIsSent)
	s.AddStep(`the status is (\d+) (.*)`, statusIs)
	s.AddStep(`body (.*)`, body)
	s.AddStep(`there is a valid user in the system`, user)
	s.AddStep(`there is a valid role in the system`, role)
	s.AddStep(`the user has the role`, userHasRole)
	s.AddStep(`there is a valid permission in the system`, permission)
	s.AddStep(`the permission is granted to the role`, permissionIsGrantedRole)
	s.Run()
}

func aValidAPIKeyAuth(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	if keys, ok := getCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["apiKeyAuth"] = datadog.APIKey{
			Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
		}
	} else {
		panic("could not set API key")
	}
	return ctx
}

func aValidAppKeyAuth(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	if keys, ok := getCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["appKeyAuth"] = datadog.APIKey{
			Key: os.Getenv("DD_TEST_CLIENT_APP_KEY"),
		}
	} else {
		panic("could not set App key")
	}
	return ctx
}

// anInstanceOf sets API callable to apiKey{}
func anInstanceOf(t gobdd.TestingT, ctx gobdd.Context, name string) gobdd.Context {
	client := getClient(ctx)

	ctx.Set(ctxKey{}, tests.WithClock(getCtx(ctx), t.(*testing.T)))

	r, err := tests.Recorder(getCtx(ctx), t.(*testing.T))
	if err != nil {
		log.Fatal(err)
	}
	ctx.Set(cleanupKey{}, func() { r.Stop() })
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	ct := reflect.ValueOf(client)
	f := reflect.Indirect(ct).FieldByName(name + "Api")
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid API name %s", name))
	}
	ctx.Set(apiKey{}, f)
	return ctx
}

// newRequest sets callable operation to requestKey{}
func newRequest(t gobdd.TestingT, ctx gobdd.Context, name string) gobdd.Context {
	c, err := ctx.Get(apiKey{})
	if err != nil {
		panic(err)
	}
	r := c.(reflect.Value)
	f := r.MethodByName(name)
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid method name %s on API", name))
	}

	ctx.Set(requestKey{}, f)
	ctx.Set(requestParamsKey{}, make(map[string]interface{}))
	ctx.Set(requestArgsKey{}, make([]interface{}, 0))
	return ctx
}

func statusIs(t gobdd.TestingT, ctx gobdd.Context, expected int, text string) gobdd.Context {
	r, err := ctx.Get(responseKey{})
	if err != nil {
		panic(err)
	}
	// Execute() returns tripples -> 2nd value is *http.Response -> get StatusCode
	resp := r.([]reflect.Value)
	code := resp[len(resp)-2].Interface().(*http.Response).StatusCode
	if expected != code {
		t.Fatalf("Excepted %d got %d", expected, code)
	}
	return ctx
}

func addParameterFrom(t gobdd.TestingT, ctx gobdd.Context, name string, path string) gobdd.Context {
	value, err := lookup.LookupStringI(getData(ctx), path)
	if err != nil {
		panic(err)
	}
	getRequestParameters(ctx)[name] = value
	ctx.Set(requestArgsKey{}, append(getRequestArguments(ctx), value))
	return ctx
}

func requestIsSent(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	c, err := ctx.Get(requestKey{})
	if err != nil {
		t.Error("Missing requestKey{}")
	}
	f := c.(reflect.Value)

	requestParameters := getRequestParameters(ctx)
	in := make([]reflect.Value, f.Type().NumIn())

	// first argument is always context.Context
	in[0] = reflect.ValueOf(getCtx(ctx))

	for i := 1; i < f.Type().NumIn(); i++ {
		object := getRequestArguments(ctx)[i-1]
		in[i] = object.(reflect.Value)
	}

	request := f.Call(in)[0]

	if body, ok := requestParameters["body"]; ok {
		a := request.MethodByName("Body").Type().In(0)
		at := reflect.New(a)
		json.Unmarshal([]byte(body.(string)), at.Interface())
		request = request.MethodByName("Body").Call([]reflect.Value{
			at.Elem(),
		})[0]
	}

	result := request.MethodByName("Execute").Call(nil)
	ctx.Set(responseKey{}, result)
	return ctx
}

func body(t gobdd.TestingT, ctx gobdd.Context, body string) gobdd.Context {
	data := getData(ctx)
	name := strings.Join(strings.Split(t.(*testing.T).Name(), "/")[:3], "/")
	data["unique"] = *UniqueEntityName(getCtx(ctx), name)
	getRequestParameters(ctx)["body"] = templated(data, body)
	return ctx
}

func user(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := getClient(ctx)

	uca := datadog.NewUserCreateAttributes()
	name := strings.ToLower(*tests.UniqueEntityName(getCtx(ctx), t.(*testing.T)))
	uca.SetEmail(fmt.Sprintf("%s@datadoghq.com", name))
	uca.SetName(name[:44])
	uca.SetTitle("Big boss")
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreateRequest()
	ucp.SetData(*ucd)
	ur, _, err := client.UsersApi.CreateUser(getCtx(ctx)).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("error creating role: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	/*
		urData := ur.GetData()
		uid := urData.GetId()
		defer disableUser(getCtx(ctx), uid)
	*/
	data := getData(ctx)
	data["user"] = ur

	return ctx
}

func role(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := getClient(ctx)

	rca := datadog.NewRoleCreateAttributes()
	rca.SetName(*tests.UniqueEntityName(getCtx(ctx), t.(*testing.T)))
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequest()
	rcp.SetData(*rcd)
	rr, _, err := client.RolesApi.CreateRole(getCtx(ctx)).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("error creating role: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	/*
		rrData := rr.GetData()
		rid := rrData.GetId()
		defer deleteRole(getCtx(ctx), rid)
	*/
	data := getData(ctx)
	data["role"] = rr
	return ctx
}

func userHasRole(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := getClient(ctx)
	data := getData(ctx)
	ur := data["user"].(datadog.UserResponse)
	rr := data["role"].(datadog.RoleCreateResponse)
	u := ur.GetData()
	r := rr.GetData()

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(u.GetId())
	rtu.SetData(*rtud)

	_, _, err := client.RolesApi.AddUserToRole(getCtx(ctx), r.GetId()).Body(*rtu).Execute()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return ctx
}

func permission(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := getClient(ctx)

	psr, _, err := client.RolesApi.ListPermissions(getCtx(ctx)).Execute()
	if err != nil {
		t.Fatalf("error listing permissions: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	data := getData(ctx)
	data["permission"] = psr.GetData()[0]

	return ctx
}

func permissionIsGrantedRole(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := getClient(ctx)
	data := getData(ctx)
	p := data["permission"].(datadog.Permission)
	rr := data["role"].(datadog.RoleCreateResponse)
	r := rr.GetData()

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(p.GetId())
	rtp.SetData(*rtpd)

	_, _, err := client.RolesApi.AddPermissionToRole(getCtx(ctx), r.GetId()).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return ctx
}

func getCtx(ctx gobdd.Context) context.Context {
	c, err := ctx.Get(ctxKey{})
	if err != nil {
		panic(err)
	}
	return c.(context.Context)
}

func getClient(ctx gobdd.Context) *datadog.APIClient {
	c, err := ctx.Get(apiClientKey{})
	if err != nil {
		panic(err)
	}
	return c.(*datadog.APIClient)
}

func getData(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(dataKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]interface{})
}

func getRequestParameters(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(requestParamsKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]interface{})
}

func getRequestArguments(ctx gobdd.Context) []interface{} {
	c, err := ctx.Get(requestArgsKey{})
	if err != nil {
		panic(err)
	}
	return c.([]interface{})
}

func templated(data interface{}, source string) string {
	re := regexp.MustCompile(`{{ ?([^}])+ ?}}`)
	replace := func(source string) string {
		path := strings.Trim(source, "{ }")
		v, err := lookup.LookupStringI(data, path)
		if err != nil {
			panic(fmt.Sprintf("problem with replacement of %s: %v", source, err))
		}
		return v.String()
	}
	return re.ReplaceAllStringFunc(source, replace)
}

// UniqueEntityName will return a unique string that can be used as a title/description/summary/...
// of an API entity. When used in Azure Pipelines and RECORD=true or RECORD=none, it will include
// BuildId to enable mapping resources that weren't deleted to builds.
func UniqueEntityName(ctx context.Context, name string) *string {
	buildID, present := os.LookupEnv("BUILD_BUILDID")
	if !present || !tests.IsCIRun() || tests.GetRecording() == tests.ModeReplaying {
		buildID = "local"
	}

	// NOTE: some endpoints have limits on certain fields (e.g. Roles V2 names can only be 55 chars long),
	// so we need to keep this short
	result := fmt.Sprintf("go-%s-%s-%d", name, buildID, tests.ClockFromContext(ctx).Now().Unix())
	// In case this is used in URL, make sure we replace the slash that is added by subtests
	result = strings.ReplaceAll(result, "/", "-")
	return &result
}
