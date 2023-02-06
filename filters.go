package goenv

import "strings"

const (
	EnvRequiredRule = "required"
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
