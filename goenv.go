package goenv

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

func Load[T any](s *T) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	v := reflect.ValueOf(*s)
	ptr := reflect.ValueOf(s)

	errs := []error{}

	for i := 0; i < v.Type().NumField(); i++ {
		s := ptr.Elem()

		if s.Kind() == reflect.Struct {
			env := v.Type().Field(i).Tag.Get("env")
			sl := strings.Split(env, ",")

			envName := sl[0]
			val := os.Getenv(envName)

			if isEnvRequired(env) && val == "" {
				errs = append(errs, fmt.Errorf("Env %s is required", envName))
			}

			defaultValue := getEnvDefaultValue(env)
			if val == "" {
				val = defaultValue
			}

			f := s.Field(i)
			if f.IsValid() && f.CanSet() {
				f.SetString(val)
			}
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func MustLoad[T any](s *T) {
	err := Load[T](s)
	if err != nil {
		panic(err)
	}
}
