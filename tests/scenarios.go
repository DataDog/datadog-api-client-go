/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/mcuadros/go-lookup"
)

type ctxKey struct{}
type apiKey struct{}
type requestKey struct{}
type requestArgsKey struct{}
type requestParamsKey struct{}
type responseKey struct{}
type dataKey struct{}
type bodyKey struct{}
type cleanupKey struct{}

func Decorate(impl interface{}) interface{} {
	fn := reflect.ValueOf(impl)

	inner := func(in []reflect.Value) []reflect.Value {
		f := reflect.ValueOf(impl)
		ret := f.Call(in)
		return ret
	}

	v := reflect.MakeFunc(fn.Type(), inner)
	return v.Interface()
}

// Templated replaces {{ path }} in source with value from data[path].
func Templated(data interface{}, source string) string {
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

func GetCtx(ctx gobdd.Context) context.Context {
	c, err := ctx.Get(ctxKey{})
	if err != nil {
		panic(err)
	}
	return c.(context.Context)
}

func SetCtx(ctx gobdd.Context, value context.Context) {
	ctx.Set(ctxKey{}, value)
}

func SetAPI(ctx gobdd.Context, value interface{}) {
	ctx.Set(apiKey{}, value)
}

func GetData(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(dataKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]interface{})
}

func SetData(ctx gobdd.Context, value map[string]interface{}) {
	ctx.Set(dataKey{}, value)
}

func GetRequestParameters(ctx gobdd.Context) map[string]interface{} {
	c, err := ctx.Get(requestParamsKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]interface{})
}

func GetRequestArguments(ctx gobdd.Context) []interface{} {
	c, err := ctx.Get(requestArgsKey{})
	if err != nil {
		panic(err)
	}
	return c.([]interface{})
}

func SetCleanup(ctx gobdd.Context, value map[string]func()) {
	ctx.Set(cleanupKey{}, value)
}

func GetCleanup(ctx gobdd.Context) map[string]func() {
	c, err := ctx.Get(cleanupKey{})
	if err != nil {
		panic(err)
	}
	return c.(map[string]func())
}

func RunCleanup(ctx gobdd.Context) gobdd.Context {
	c, _ := ctx.Get(cleanupKey{})
	cc := c.(map[string]func())
	fmt.Println("run", cc, ctx)
	for name, cleanup := range cc {
		defer fmt.Println(name)
		defer cleanup() // in reverse order
	}
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
	value, err := lookup.LookupStringI(GetData(ctx), path)
	if err != nil {
		panic(err)
	}
	GetRequestParameters(ctx)[name] = value
	ctx.Set(requestArgsKey{}, append(GetRequestArguments(ctx), value))
	return ctx
}

func requestIsSent(t gobdd.TestingT, ctx gobdd.Context) gobdd.Context {
	c, err := ctx.Get(requestKey{})
	if err != nil {
		t.Error("Missing requestKey{}")
	}
	f := c.(reflect.Value)

	requestParameters := GetRequestParameters(ctx)
	in := make([]reflect.Value, f.Type().NumIn())

	// first argument is always context.Context
	in[0] = reflect.ValueOf(GetCtx(ctx))

	for i := 1; i < f.Type().NumIn(); i++ {
		object := GetRequestArguments(ctx)[i-1]
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
	data := GetData(ctx)
	name := strings.Join(strings.Split(t.(*testing.T).Name(), "/")[:3], "/")
	data["unique"] = WithUniqueSurrounding(GetCtx(ctx), name)
	GetRequestParameters(ctx)["body"] = Templated(data, body)
	return ctx
}

// Steps implementation
var Steps = map[string]interface{}{
	`new "([^"]+)" request`:       newRequest,
	`parameter ([^ ]+) from (.*)`: addParameterFrom,
	`the request is sent`:         requestIsSent,
	`the status is (\d+) (.*)`:    statusIs,
	`body (.*)`:                   body,
}
