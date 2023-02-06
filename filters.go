package goenv

import "strings"

const (
	EnvRequiredRule = "required"
	EnvDefaultRule  = "default"
)

func isEnvRequired(env string) bool {
	s := strings.Split(env, ",")

	for _, item := range s {
		if item == EnvRequiredRule {
			return true
		}
	}

	return false
}

func getEnvDefaultValue(env string) string {
	s := strings.Split(env, ",")

	for _, item := range s {
		if strings.HasPrefix(item, EnvDefaultRule) {
			rule := strings.Split(item, "=")
			if len(rule) < 2 {
				return ""
			}

			v := rule[1]
			return v
		}
	}

	return ""
}
