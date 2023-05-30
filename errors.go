package u

import (
	"errors"
	"strings"
	"fmt"
	"os"
)

func SplitError(err error) (s, field string) {
	s = fmt.Sprint(err)
	start := strings.Index(s, "|f=")
	if start >= 0 {
		start += 3
		end := strings.Index(s[start:], "|")
		if end == -1 {
			end = len(s)
		} else {
			end = start + end
		}
		field = s[start:end]
		s = strings.Replace(s, "|f=" + field, "", 1)
	}
	return
}

func AddError2(el *[]error, s string, extra ...string) {
	AddError1(el, errors.New(s), extra...)
}

func AddError1(el *[]error, e error, extra ...string) {
	if *el == nil {
		*el = []error{}
	}
	if len(extra) > 0 {
		s := e.Error()
		for _, ex := range extra {
			s += "|" + ex
		}
		e = errors.New(s)
	}
	
	*el = append(*el, e)
}

func JoinErrors(el []error) (e error) {
	var msg string = ""
	for _, e := range el {
		s, _ := SplitError(e)
		msg += s + ".\n"
	}
	return errors.New(msg)
}

func PrintError1(err error, lang string) {
	msg, _ := SplitError(err)
	PrintError2(msg, lang)
}

func PrintError2(msg string, lang string) {
	if TranslateFunction != nil {
		fmt.Fprintf(os.Stderr, "error: %s.\n", TranslateFunction(msg, lang));
	} else {
		fmt.Fprintf(os.Stderr, "error: %s.\n", msg);
	}
}
