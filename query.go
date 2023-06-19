package u

import (
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
	"strings"
)

// QueryToMap gets the parameters from URL parameters and the
// POST form and makes a map.
func QueryToMap(ctx *gin.Context) (kv map[string]string) {
	kv = map[string]string{}
	for k, v := range ctx.Request.URL.Query() {
		kv[k] = v[0]
	}
	for k, v := range ctx.Request.PostForm {
		kv[k] = v[0]
	}
	return kv
}

// ArrayToMap converts a "KEY=VALUE" array into a map.
func ArrayToMap(args []string) (ret map[string]string) {
	ret = map[string]string {}
	for _, a := range args {
		varval := strings.SplitN(a, "=", 2)
		if len(varval) == 2 {
			ret[varval[0]] = varval[1]
		} else {
			ret[varval[0]] = ""
		}
	}
	return
}

// From a route defined as "/something/:p" and gotten "/something/v"
// fetch "v".
func GetIDFromPath(ctx *gin.Context, p string) (val int64, err error) {
	var valS string
	valS = ctx.Param(p)
	if len(valS) <= 0 {
		err = errors.New("Invalid PATH")
		return
	}
	val, err = strconv.ParseInt(valS, 10, 64)
	if err != nil {
		return
	}
	return
}

// From a route as "/something?p=v" fetch the "v" part.
func GetIDFromQuery(ctx *gin.Context, p string) (val int64, found bool) {
	var err error
	val, err = strconv.ParseInt(ctx.Query(p), 10, 64)
	return val, err == nil
}

