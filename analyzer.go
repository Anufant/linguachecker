package linguachecker

import (
	"flag"
	"log"
	"strings"

	"github.com/pemistahl/lingua-go"
	"golang.org/x/tools/go/analysis"
)

var flags = flag.NewFlagSet("linguachecker", flag.ExitOnError)

type arrayFlags []string

func (i *arrayFlags) String() string {
	if i == nil {
		return "<nil>"
	}

	return strings.Join(*i, ",")
}

func (i *arrayFlags) Append(value string) error {
	*i = append(*i, value)
	return nil
}

func (i *arrayFlags) Set(value string) error {
	*i = strings.Split(value, ",")
	return nil
}

func init() {
	include := arrayFlags{}
	exclude := arrayFlags{}
	flags.Var(&include, "include", "List of allowed languages. If emty - allow all languages. If not empty - allow only listed languages.")
	flags.Var(&exclude, "exclude", "List of prohibited languages")

	Analyzer.Flags = *flags
}

var Analyzer = &analysis.Analyzer{
	Name: "Linguachecker",
	Doc:  "Linguachecker checks for language of comments in code",
	Run:  run,
}

// run какой-то комментарий на русском языке.
func run(pass *analysis.Pass) (interface{}, error) {
	include := []lingua.Language{}
	if incl := pass.Analyzer.Flags.Lookup("include"); incl.Value != nil {
		if val, ok := incl.Value.(*arrayFlags); ok && val != nil {
			for _, lang := range *val {
				code := lingua.GetIsoCode639_1FromValue(lang)
				if code != lingua.UnknownIsoCode639_1 {
					include = append(
						include,
						lingua.GetLanguageFromIsoCode639_1(code),
					)
				}
			}
		}
	}

	exclude := []lingua.Language{}
	if excl := pass.Analyzer.Flags.Lookup("exclude"); excl.Value != nil {
		if val, ok := excl.Value.(*arrayFlags); ok && val != nil {
			for _, lang := range *val {
				code := lingua.GetIsoCode639_1FromValue(lang)
				if code != lingua.UnknownIsoCode639_1 {
					exclude = append(
						exclude,
						lingua.GetLanguageFromIsoCode639_1(code),
					)
				}
			}
		}
	}

	var detector lingua.LanguageDetector
	if len(include) > 1 {
		detector = lingua.NewLanguageDetectorBuilder().FromLanguages(include...).Build()
	} else {
		detector = lingua.NewLanguageDetectorBuilder().FromAllLanguages().Build()
	}

	for _, f := range pass.Files {
		for _, group := range f.Comments {
			for _, cmt := range group.List {
				if detectedLanguage, ok := detector.DetectLanguageOf(sanitizeComment(cmt.Text)); ok {
					log.Printf("Detected language: %v", detectedLanguage)
					excluded := false
					if len(exclude) > 0 {
						for _, excludedLanguage := range exclude {
							if detectedLanguage == excludedLanguage {
								pass.Reportf(cmt.Pos(), "comment is not written in acceptable language")
							}
						}
					}

					log.Printf("Excluded: %v", excluded)
					if !excluded && len(include) > 0 {
						included := false
						for _, includedLanguage := range include {
							log.Printf("Included language: %v", includedLanguage)
							if detectedLanguage == includedLanguage {
								included = true
							}
						}
						log.Printf("Included: %v", included)
						if !included {
							pass.Reportf(cmt.Pos(), "comment is not written in acceptable language")
						}
					}
				} else {
					pass.Reportf(cmt.Pos(), "comment is not written in acceptable language")
				}
			}
		}
	}

	return nil, nil
}

// TODO: add support for multiple languages in comments
func sanitizeComment(comment string) string {
	if strings.LastIndex(comment, "// want") != -1 {
		comment = comment[:strings.LastIndex(comment, "// want")]
	}

	return comment
}
