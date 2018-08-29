package docker

import (
	"regexp"

	"github.com/go-courier/envconf"
)

var reEnvVar = regexp.MustCompile("(\\$?\\$)\\{?([A-Z0-9_]+)\\}?")

func ParseEnvVars(s string, envVars *envconf.EnvVars) string {
	result := reEnvVar.ReplaceAllStringFunc(s, func(str string) string {
		matched := reEnvVar.FindAllStringSubmatch(str, -1)[0]

		// skip $${ }
		if matched[1] == "$$" {
			return "${" + matched[2] + "}"
		}

		if value := envVars.Get(matched[2]); value != "" {
			return value
		}

		return "${" + matched[2] + "}"
	})

	return result
}

func stringSome(list []string, checker func(item string, i int) bool) bool {
	for i, item := range list {
		if checker(item, i) {
			return true
		}
	}
	return false
}

func stringIncludes(list []string, target string) bool {
	return stringSome(list, func(item string, i int) bool {
		return item == target
	})
}
