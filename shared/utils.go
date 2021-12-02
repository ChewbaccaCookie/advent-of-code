package shared

import "regexp"

func RegexMatch(regex string, val string) []string {
	re := regexp.MustCompile(regex)
	return re.FindStringSubmatch(val)
}
