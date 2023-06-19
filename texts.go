package u

var TranslateFunction func (lang, fr string) string = nil

// T marks a string as translatable. Can be used to extract
// internationalizable texts from the program.
func T(fr string) string {
	return fr
}

// Translate uses the TranslateFunction set by the user to
// query translations.
func Translate(t string, lang string) string {
	if TranslateFunction != nil {
		return TranslateFunction(t, lang)
	} else {
		return t
	}
}
