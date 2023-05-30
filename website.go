package u

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/jmoiron/sqlx"
	"net/http"
	"html/template"
	"log"
)

type Website struct {
	Language       string
	Ctx            *gin.Context
	DB             *sqlx.DB
	Errors         string
	Message        string
	FieldErrors    map[string]string
	Session        sessions.Session
}

func (w *Website) Init(ctx *gin.Context, db *sqlx.DB) {
	w.Language    = "es"
	w.Ctx         = ctx
	w.DB          = db
	w.FieldErrors = map[string]string{}
	w.Session     = sessions.Default(ctx)
	return
}

func (w *Website) T(i string) string {
	return Translate(i, w.Language)
}

func (w *Website) Skip() bool {
	return false
}

func (w *Website) Error(errl ...error) {
	for _, err := range errl {
		msg, field := SplitError(err)
		if len(msg) > 0 {
			w.Errors += "* " + msg + ".\n"
		}
		if len(field) > 0 {
			w.FieldErrors[field] = msg
		}
	}
}

func (w *Website) SetMessage(msg string) {
	w.Message = msg
}

func (w *Website) Fatal(err error) {
	w.Ctx.AbortWithError(http.StatusInternalServerError, err)
	log.Print(err)
	return
}

func (w *Website) FieldError(f string) (r template.HTML) {
	msg, found := w.FieldErrors[f]
	if !found {
		return template.HTML("")
	}
	return template.HTML("<span class=\"error text-danger\">"+ msg +" </span>")
}

func (w *Website) ErrorsFound() bool {
	return len(w.Errors)>0
}

func (w *Website) MessageFound() bool {
	return w.Message != ""
}
