package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
)

func TestScenarios(t *testing.T) {
	s := gobdd.NewSuite(
		t,
		gobdd.WithIgnoredTags([]string{"@todo"}),
		gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
			client := datadog.NewAPIClient(NewConfiguration())
			tests.SetCtx(ctx, context.WithValue(context.WithValue(
				context.Background(),
				datadog.ContextAPIKeys,
				map[string]datadog.APIKey{},
			), clientKey, client))
			tests.SetData(ctx, make(map[string]interface{}))
			tests.SetCleanup(ctx, make(map[string]func()))
		}), gobdd.WithAfterScenario(func(ctx gobdd.Context) {
			fmt.Println(ctx)
			ctx = tests.RunCleanup(ctx)
		}),
	)

	for expr, step := range tests.Steps {
		s.AddStep(expr, step)
	}

	s.AddStep(`a valid "apiKeyAuth" key`, aValidAPIKeyAuth)
	s.AddStep(`a valid "appKeyAuth" key`, aValidAppKeyAuth)
	s.AddStep(`an instance of "([^"]+)" API`, anInstanceOf)
	s.AddStep(`there is a valid user in the system`, user)
	s.AddStep(`there is a valid role in the system`, role)
	s.AddStep(`the user has the role`, userHasRole)
	s.AddStep(`there is a valid permission in the system`, permission)
	s.AddStep(`the permission is granted to the role`, permissionIsGrantedRole)
	s.Run()
}

func aValidAPIKeyAuth(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["apiKeyAuth"] = datadog.APIKey{
			Key: os.Getenv("DD_TEST_CLIENT_API_KEY"),
		}
	} else {
		panic("could not set API key")
	}
	return ctx
}

func aValidAppKeyAuth(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
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
	client := Client(tests.GetCtx(ctx))

	tests.SetCtx(ctx, tests.WithClock(tests.GetCtx(ctx), t.(*testing.T)))

	r, err := tests.Recorder(tests.GetCtx(ctx), t.(*testing.T))
	if err != nil {
		log.Fatal(err)
	}
	tests.GetCleanup(ctx)["recorder"] = func() { r.Stop() }
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	ct := reflect.ValueOf(client)
	f := reflect.Indirect(ct).FieldByName(name + "Api")
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid API name %s", name))
	}
	tests.SetAPI(ctx, f)
	return ctx
}

func user(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := Client(tests.GetCtx(ctx))

	uca := datadog.NewUserCreateAttributes()
	name := strings.ToLower(*tests.UniqueEntityName(tests.GetCtx(ctx), t.(*testing.T)))
	uca.SetEmail(fmt.Sprintf("%s@datadoghq.com", name))
	uca.SetName(name[:44])
	uca.SetTitle("Big boss")
	ucd := datadog.NewUserCreateData()
	ucd.SetAttributes(*uca)
	ucp := datadog.NewUserCreateRequest()
	ucp.SetData(*ucd)
	ur, _, err := client.UsersApi.CreateUser(tests.GetCtx(ctx)).Body(*ucp).Execute()
	if err != nil {
		t.Fatalf("error creating role: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	urData := ur.GetData()
	uid := urData.GetId()
	cctx := tests.GetCtx(ctx)
	tests.GetCleanup(ctx)["user"] = func() {
		disableUser(cctx, uid)
	}

	tests.GetData(ctx)["user"] = ur
	return ctx
}

func role(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := Client(tests.GetCtx(ctx))

	rca := datadog.NewRoleCreateAttributes()
	rca.SetName(*tests.UniqueEntityName(tests.GetCtx(ctx), t.(*testing.T)))
	rcd := datadog.NewRoleCreateData()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequest()
	rcp.SetData(*rcd)
	rr, _, err := client.RolesApi.CreateRole(tests.GetCtx(ctx)).Body(*rcp).Execute()
	if err != nil {
		t.Fatalf("error creating role: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	rrData := rr.GetData()
	rid := rrData.GetId()
	cctx := tests.GetCtx(ctx)
	tests.GetCleanup(ctx)["role"] = func() {
		fmt.Println("removing role", rid)
		deleteRole(cctx, rid)
	}

	tests.GetData(ctx)["role"] = rr
	return ctx
}

func userHasRole(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := Client(tests.GetCtx(ctx))
	data := tests.GetData(ctx)
	ur := data["user"].(datadog.UserResponse)
	rr := data["role"].(datadog.RoleCreateResponse)
	u := ur.GetData()
	r := rr.GetData()

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(u.GetId())
	rtu.SetData(*rtud)

	_, _, err := client.RolesApi.AddUserToRole(tests.GetCtx(ctx), r.GetId()).Body(*rtu).Execute()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return ctx
}

func permission(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := Client(tests.GetCtx(ctx))

	psr, _, err := client.RolesApi.ListPermissions(tests.GetCtx(ctx)).Execute()
	if err != nil {
		t.Fatalf("error listing permissions: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	tests.GetData(ctx)["permission"] = psr.GetData()[0]
	return ctx
}

func permissionIsGrantedRole(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	client := Client(tests.GetCtx(ctx))
	data := tests.GetData(ctx)
	p := data["permission"].(datadog.Permission)
	rr := data["role"].(datadog.RoleCreateResponse)
	r := rr.GetData()

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(p.GetId())
	rtp.SetData(*rtpd)

	_, _, err := client.RolesApi.AddPermissionToRole(tests.GetCtx(ctx), r.GetId()).Body(*rtp).Execute()
	if err != nil {
		t.Fatalf("%v", err)
	}
	return ctx
}
