package slug

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/rainycape/unidecode" // external dependency
)

var allowedSpecialChars, disallowedChars, multipleDashes *regexp.Regexp
var isAsciiNumber, isSlugValid *regexp.Regexp
var isTagValid, isTagListValid *regexp.Regexp

func init() {
	// characters that are replaced with a hypthen (-), (whitespaces, commas, dots, forward slashes, back slashes, hypthens, underscores, equal signs, and pluses)
	// multiple hypthens are replaced with one (Ex: hello------world -> hello-world)
	allowedSpecialChars = regexp.MustCompile(`[\s,./\\-_=+]+`)
	disallowedChars = regexp.MustCompile(`[^A-Za-z0-9-]`)
	multipleDashes = regexp.MustCompile("-+")

	isAsciiNumber = regexp.MustCompile(`^[0-9]+$`)
	isSlugValid = regexp.MustCompile(`^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`)

	isTagValid = regexp.MustCompile(`^[A-Za-z0-9]+(?: [A-Za-z0-9]+)*$`)
	isTagListValid = regexp.MustCompile("^[a-z0-9]+(\\|*[a-z0-9]+(?:-[a-z0-9]+)*)*$")

	//	uses commas for separator character instead of pipe (|)
	//	isTagListValidAlternative = regexp.MustCompile(`^[a-z0-9]+(,*[a-z0-9]+(?:-[a-z0-9]+)*)*$`)
}

func GetAsciiSlug(title string) string {

	// replace unicode characters with ascii
	title = unidecode.Unidecode(title)
	title = strings.ToLower(title)

	title = allowedSpecialChars.ReplaceAllString(title, "-")
	title = disallowedChars.ReplaceAllString(title, "")
	title = multipleDashes.ReplaceAllString(title, "-")
	title = strings.Trim(title, "-")

	return title
}

func IsSlug(sl string) bool {
	return isSlugValid.MatchString(sl)
}

func IsUTF8Slug(sl string) bool {

	var temp string

	// check all the characters before deciding
	for _, c := range sl {
		if !unicode.IsLetter(c) && !isAsciiNumber.MatchString(string(c)) {
			// not a letter or number, is it a hypthen?
			if c == '-' {
				temp += "-"
			} else {
				return false
			}
		} else {
			// replace it with any letter
			temp += "a"
		}
	}

	return isSlugValid.MatchString(temp)
}

// for validating an HTML form field, implements an interface from another package
type isSlugField struct {
}

func IsSlugField() isSlugField {
	return isSlugField{}
}

/*
	A slug must contain at least one ASCII letter or number after being parsed, it cannot be blank.
	This uses another package with lots of examples in the "*_test.go" files
*/
func (i isSlugField) Validate(fields []string, errorMessages map[string]string) (error, []interface{}) {

	var errMsg = errors.New("This field must contain at least one letter or number.")
	var field = ""

	if len(fields) > 0 {
		field = fields[0]
	}

	field = GetAsciiSlug(field)

	if IsSlug(field) {
		return nil, nil
	}

	return errMsg, nil
}
