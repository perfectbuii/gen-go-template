package utils

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"github.com/cucumber/messages-go/v16"
	"github.com/perfectbuii/gen-go-template/cmd/gen-layer/utils/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const FeatureExt = ".feature"

func GetStepsContent(filePath string, stepMap map[string]bool) []string {
	var stepTexts []string
	features, err := parser.ParseFeatures([]string{filePath})
	if err != nil {
		panic(err)
	}
	for _, f := range features {
		if f.GherkinDocument.Feature == nil {
			return stepTexts
		}
		for _, children := range f.GherkinDocument.Feature.Children {
			steps := []*messages.Step{}
			if children.Scenario != nil {
				steps = children.Scenario.Steps
			}
			if children.Background != nil {
				steps = append(steps, children.Background.Steps...)
			}
			for _, step := range steps {
				expr, name := GetExprAndFuncName(step.Text)
				if _, ok := stepMap[expr]; !ok {
					stepTexts = append(stepTexts, "`"+expr+"`: s."+name+",\n")
					stepMap[expr] = true
				}
			}
		}
	}
	return stepTexts
}

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

// some snippet formatting regexps
var (
	snippetExprCleanup = regexp.MustCompile(`([\/\[\]\(\)\\^\\$\.\|\?\*\+\'])`)
	snippetExprQuoted  = regexp.MustCompile(`(\W|^)\"(?:[^\"]*)\"(\W|$)`)
	snippetMethodName  = regexp.MustCompile(`[^a-zA-Z\_\ ]`)
	snippetNumbers     = regexp.MustCompile(`(\d+)`)
)

func GetArgs(expr string) (ret string) {
	var (
		args      []string
		pos       int
		breakLoop bool
	)

	for !breakLoop {
		part := expr[pos:]
		ipos := strings.Index(part, "(\\d+)")
		spos := strings.Index(part, "\"([^\"]*)\"")

		switch {
		case spos == -1 && ipos == -1:
			breakLoop = true
		case spos == -1:
			pos += ipos + len("(\\d+)")
			args = append(args, reflect.Int.String())
		case ipos == -1:
			pos += spos + len("\"([^\"]*)\"")
			args = append(args, reflect.String.String())
		case ipos < spos:
			pos += ipos + len("(\\d+)")
			args = append(args, reflect.Int.String())
		case spos < ipos:
			pos += spos + len("\"([^\"]*)\"")
			args = append(args, reflect.String.String())
		}
	}

	var last string

	for i, arg := range args {
		if last == "" || last == arg {
			ret += fmt.Sprintf("arg%d, ", i+1)
		} else {
			ret = strings.TrimRight(ret, ", ") + fmt.Sprintf(" %s, arg%d, ", last, i+1)
		}

		last = arg
	}

	return strings.TrimSpace(strings.TrimRight(ret, ", ") + " " + last)
}

func GetExprAndFuncName(step string) (string, string) {
	expr := snippetExprCleanup.ReplaceAllString(step, "\\$1")
	expr = snippetNumbers.ReplaceAllString(expr, "(\\d+)")
	expr = snippetExprQuoted.ReplaceAllString(expr, "$1\"([^\"]*)\"$2")
	expr = "^" + strings.TrimSpace(expr) + "$"

	name := snippetNumbers.ReplaceAllString(step, " ")
	name = snippetExprQuoted.ReplaceAllString(name, " ")
	name = strings.TrimSpace(snippetMethodName.ReplaceAllString(name, ""))
	nameSplitArr := strings.Split(name, " ")
	words := make([]string, 0, len(nameSplitArr))
	for i, w := range nameSplitArr {
		switch {
		case i != 0:
			caser := cases.Title(language.English)
			w = caser.String(w)

		case len(w) > 0:
			w = string(unicode.ToLower(rune(w[0]))) + w[1:]
		}
		words = append(words, w)
	}
	name = strings.Join(words, "")

	return expr, name
}
