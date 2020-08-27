package test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
	"github.com/go-bdd/gobdd"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func undoIgnore(ctx gobdd.Context) {}

func undoCreateUser(ctx gobdd.Context) {
	user := tests.GetResponse(ctx)[0].Interface().(datadog.UserResponse)
	cctx := tests.GetCtx(ctx)
	Client(cctx).UsersApi.DisableUser(cctx, user.Data.GetId()).Execute()
}

func undoCreateRole(ctx gobdd.Context) {
	role := tests.GetResponse(ctx)[0].Interface().(datadog.RoleCreateResponse)
	cctx := tests.GetCtx(ctx)
	Client(cctx).RolesApi.DeleteRole(cctx, role.Data.GetId()).Execute()
}

var requestsUndo = map[string]func(ctx gobdd.Context){
	"AddPermissionToRole":      undoIgnore,
	"AddUserToRole":            undoIgnore,
	"CreateRole":               undoCreateRole,
	"CreateUser":               undoCreateUser,
	"DisableUser":              undoIgnore,
	"DeleteRole":               undoIgnore,
	"GetInvitation":            undoIgnore,
	"GetRole":                  undoIgnore,
	"GetUser":                  undoIgnore,
	"ListPermissions":          undoIgnore,
	"ListRolePermissions":      undoIgnore,
	"ListRoles":                undoIgnore,
	"ListRoleUsers":            undoIgnore,
	"ListUsers":                undoIgnore,
	"ListUserPermissions":      undoIgnore,
	"RemovePermissionFromRole": undoIgnore,
	"RemoveUserFromRole":       undoIgnore,
	"SendInvitations":          undoIgnore,
	"UpdateRole":               undoIgnore,
	"UpdateUser":               undoIgnore,
}

func TestScenarios(t *testing.T) {
	s := gobdd.NewSuite(
		t,
		gobdd.WithIgnoredTags([]string{"@skip"}),
		gobdd.WithBeforeScenario(func(ctx gobdd.Context) {
			ct, _ := ctx.Get(gobdd.TestingTKey{})
			cctx, finish := WithRecorder(
				context.WithValue(
					context.Background(),
					datadog.ContextAPIKeys,
					map[string]datadog.APIKey{},
				),
				ct.(*testing.T),
			)
			tests.SetCtx(ctx, cctx)
			tests.SetRequestsUndo(ctx, requestsUndo)
			tests.SetData(ctx, make(map[string]interface{}))
			tests.SetCleanup(ctx, map[string]func(){"99-finish": finish})
		}), gobdd.WithAfterScenario(func(ctx gobdd.Context) {
			tests.RunCleanup(ctx)
		}),
		gobdd.WithBeforeStep(func(ctx gobdd.Context) {
			ct, _ := ctx.Get(gobdd.TestingTKey{})
			cctx := tests.GetCtx(ctx)
			parts := strings.Split(ct.(*testing.T).Name(), "/")
			if parent, ok := tracer.SpanFromContext(cctx); ok {
				ctx.Set("parentSpan", parent)
			} else {
				ctx.Set("parentSpan", nil)
			}
			_, cctx = tracer.StartSpanFromContext(
				cctx,
				"step",
				tracer.SpanType("step"),
				tracer.ResourceName(parts[len(parts)-1]),
			)
			tests.SetCtx(ctx, cctx)
		}),
		gobdd.WithAfterStep(func(ctx gobdd.Context) {
			cctx := tests.GetCtx(ctx)
			if span, ok := tracer.SpanFromContext(cctx); ok {
				ct, _ := ctx.Get(gobdd.TestingTKey{})
				failed := ct.(*testing.T).Failed()
				if failed {
					span.SetTag(ext.Error, failed)
				}
				span.Finish()

				if parent, err := ctx.Get("parentSpan"); err == nil && parent != nil {
					tests.SetCtx(ctx, tracer.ContextWithSpan(cctx, parent.(ddtrace.Span)))
				}
			}
		}),
	)
	tests.ConfigureSteps(s)

	s.AddStep(`a valid "apiKeyAuth" key in the system`, aValidAPIKeyAuth)
	s.AddStep(`a valid "appKeyAuth" key in the system`, aValidAppKeyAuth)
	s.AddStep(`an instance of "([^"]+)" API`, anInstanceOf)
	s.AddStep(`there is a valid "user" in the system`, user)
	s.AddStep(`there is a valid "role" in the system`, role)
	s.AddStep(`the "user" has the "role"`, userHasRole)
	s.AddStep(`the "user" has a "user_invitation"`, userHasInvitation)
	s.AddStep(`there is a valid "permission" in the system`, permission)
	s.AddStep(`the "permission" is granted to the "role"`, permissionIsGrantedRole)
	s.Run()
}

func aValidAPIKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	key, ok := os.LookupEnv("DD_TEST_CLIENT_API_KEY")
	if !ok && tests.GetRecording() != tests.ModeReplaying {
		t.Fatal("DD_TEST_CLIENT_API_KEY must be set")
	}
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["apiKeyAuth"] = datadog.APIKey{
			Key: key,
		}
	} else {
		t.Fatal("could not set API key")
	}
}

func aValidAppKeyAuth(t gobdd.StepTest, ctx gobdd.Context) {
	key, ok := os.LookupEnv("DD_TEST_CLIENT_APP_KEY")
	if !ok && tests.GetRecording() != tests.ModeReplaying {
		t.Fatal("DD_TEST_CLIENT_APP_KEY must be set")
	}
	if keys, ok := tests.GetCtx(ctx).Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
		keys["appKeyAuth"] = datadog.APIKey{
			Key: key,
		}
	} else {
		t.Fatal("could not set App key")
	}
}

// anInstanceOf sets API callable to apiKey{}
func anInstanceOf(t gobdd.StepTest, ctx gobdd.Context, name string) {
	client := Client(tests.GetCtx(ctx))
	ct := reflect.ValueOf(client)
	f := reflect.Indirect(ct).FieldByName(name + "Api")
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid API name %s", name))
	}
	tests.SetAPI(ctx, f)
}

func user(t gobdd.StepTest, ctx gobdd.Context) {
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
		t.Fatalf("error creating user: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	urData := ur.GetData()
	uid := urData.GetId()
	cctx := tests.GetCtx(ctx)
	tests.GetCleanup(ctx)["20-user"] = func() {
		disableUser(cctx, uid)
	}

	tests.GetData(ctx)["user"] = ur
}

func role(t gobdd.StepTest, ctx gobdd.Context) {
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
	tests.GetCleanup(ctx)["10-role"] = func() {
		deleteRole(cctx, rid)
	}

	tests.GetData(ctx)["role"] = rr
}

func userHasRole(t gobdd.StepTest, ctx gobdd.Context) {
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
}

func userHasInvitation(t gobdd.StepTest, ctx gobdd.Context) {
	client := Client(tests.GetCtx(ctx))
	data := tests.GetData(ctx)
	ur := data["user"].(datadog.UserResponse)
	urData := ur.GetData()
	id := urData.GetId()

	rtud := datadog.NewRelationshipToUserData()
	rtud.SetId(id)
	rtu := datadog.NewRelationshipToUser()
	rtu.SetData(*rtud)
	uir := datadog.NewUserInvitationRelationships()
	uir.SetUser(*rtu)
	uid := datadog.NewUserInvitationData()
	uid.SetRelationships(*uir)
	uireq := datadog.NewUserInvitationsRequest()
	uireq.SetData([]datadog.UserInvitationData{*uid})

	ui, _, err := client.UsersApi.SendInvitations(tests.GetCtx(ctx)).Body(*uireq).Execute()
	if err != nil {
		t.Fatalf("%v", err)
	}

	tests.GetData(ctx)["user_invitation"] = ui.GetData()[0]
}

func permission(t gobdd.StepTest, ctx gobdd.Context) {
	client := Client(tests.GetCtx(ctx))

	psr, _, err := client.RolesApi.ListPermissions(tests.GetCtx(ctx)).Execute()
	if err != nil {
		t.Fatalf("error listing permissions: response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}

	tests.GetData(ctx)["permission"] = psr.GetData()[0]
}

func permissionIsGrantedRole(t gobdd.StepTest, ctx gobdd.Context) {
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
}
