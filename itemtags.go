package slug

import (
	"strings"
	"unicode"
)

const DELIMITER = ","

// check for duplicate entries in a slice
func removeDuplicates(slice []string) []string {
	var newSlice []string
	encounteredItems := make(map[string]bool)
	for _, val := range slice {

		// new entry, add it to the found list
		if _, found := encounteredItems[val]; !found {
			encounteredItems[val] = true
			newSlice = append(newSlice, val)
		} else {
			// found a duplicate
			newSlice = append(newSlice, "-")
		}
	}

	return newSlice
}

func getTagSlugListWithBlanks(tags string) []string {
	var tagList []string

	for _, tag := range strings.Split(tags, DELIMITER) {
		tag := GetAsciiSlug(tag)

		if len(tag) > 0 {
			tagList = append(tagList, tag)
		} else {
			tagList = append(tagList, "-")
		}
	}

	return tagList
}

func getPlainTagListWithBlanks(tags string) []string {
	var tagList []string
	for _, tag := range strings.Split(tags, DELIMITER) {
		if len(tag) > 0 {
			tagList = append(tagList, tag)
		} else {
			tagList = append(tagList, "-")
		}
	}

	return tagList
}

func GetTagsAndTagSlugs(tags string) ([]string, []string) {
	var tagList, slugList []string

	sl := getTagSlugListWithBlanks(tags)
	ta := getPlainTagListWithBlanks(tags)

	// remove duplicates from slugs before loop
	sl = removeDuplicates(sl)

	for i := 0; i < len(sl); i++ {
		if sl[i] != "-" && ta[i] != "-" {
			tagList = append(tagList, ta[i])
			slugList = append(slugList, sl[i])
		}
	}

	return tagList, slugList
}

func IsItemTag(tag string) bool {
	return isTagValid.MatchString(tag)
}

func IsUTF8ItemTag(sl string) bool {

	var temp string

	// check all the characters before deciding
	for _, c := range sl {
		if !unicode.IsLetter(c) && !isAsciiNumber.MatchString(string(c)) {
			// not a letter or number, is it a hypthen?
			if c == ' ' {
				temp += " "
			} else {
				return false
			}
		} else {
			// replace it with any letter
			temp += "a"
		}
	}

	return isTagValid.MatchString(temp)
}

func IsItemTagListRegex(tags string) bool {
	return isTagListValid.MatchString(tags)
}

func IsItemTagList(tags string) bool {
	for _, tag := range strings.Split(tags, DELIMITER) {
		if !IsItemTag(tag) {
			return false
		}
	}
	return true
}

func IsUTF8ItemTagList(tags string) bool {
	for _, tag := range strings.Split(tags, DELIMITER) {
		if !IsUTF8ItemTag(tag) {
			return false
		}
	}
	return true
}
