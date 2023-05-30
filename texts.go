package u

var TranslateFunction func (lang, fr string) string = nil

func T(fr string) string {
	return fr
}

func Translate(t string, lang string) string {
	if TranslateFunction != nil {
		return TranslateFunction(t, lang)
	} else {
		return t
	}
}
