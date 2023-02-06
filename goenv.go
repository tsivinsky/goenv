package goenv

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

func Load[T any](s *T) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	v := reflect.ValueOf(*s)
	ptr := reflect.ValueOf(s)

	for i := 0; i < v.Type().NumField(); i++ {
		s := ptr.Elem()

		if s.Kind() == reflect.Struct {
			env := v.Type().Field(i).Tag.Get("env")
			val := os.Getenv(env)

			f := s.Field(i)
			if f.IsValid() && f.CanSet() {
				f.SetString(val)
			}
		}
	}

	return nil
}
