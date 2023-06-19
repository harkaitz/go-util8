# UTIL8

My utility library for developing websites.

## Go documentation

    package u // import "gitlab.com/harkaitz/go-util8"
    
    var FakedTime time.Time
    var FakedTimeEnabled bool = false
    var TranslateFunction func(lang, fr string) string = nil
    func AddError1(el *[]error, e error, extra ...string)
    func AddError2(el *[]error, s string, extra ...string)
    func ArrayToMap(args []string) (ret map[string]string)
    func GetIDFromPath(ctx *gin.Context, p string) (val int64, err error)
    func GetIDFromQuery(ctx *gin.Context, p string) (val int64, found bool)
    func JoinErrors(el []error) (e error)
    func Money(i, j money.Amount) (o money.Amount)
    func MoneyStr(o money.Amount, c string, l string) (s string)
    func Now() time.Time
    func PrintError1(err error, lang string)
    func PrintError2(msg string, lang string)
    func QueryToMap(ctx *gin.Context) (kv map[string]string)
    func SplitError(err error) (s, field string)
    func T(fr string) string
    func Translate(t string, lang string) string
    type Website struct{ ... }

## Go struct go-util8.Website

    package u // import "gitlab.com/harkaitz/go-util8"
    
    type Website struct {
        Language    string
        Ctx         *gin.Context
        DB          *sqlx.DB
        Errors      string
        Message     string
        FieldErrors map[string]string
        Session     sessions.Session
    }
    
    func (w *Website) Error(errl ...error)
    func (w *Website) ErrorS(err string)
    func (w *Website) ErrorsFound() bool
    func (w *Website) Fatal(err error)
    func (w *Website) FieldError(f string) (r template.HTML)
    func (w *Website) Init(ctx *gin.Context, db *sqlx.DB)
    func (w *Website) IsPhoneUserAgent() bool
    func (w *Website) MessageFound() bool
    func (w *Website) Query(key string) string
    func (w *Website) SetMessage(msg string)
    func (w *Website) Skip() bool
    func (w *Website) T(i string) string

