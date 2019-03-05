package parser

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func NewFloat64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue float64
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseFloat(defaultStr, 64)
		if err != nil {
			return &Flag{}, fmt.Errorf("could not parse float64 default value %q: %s", defaultStr, err)
		}
	}

	var f Flag
	short, ok := tags.Lookup("short")
	if ok {
		set.Float64Var(field.Addr().Interface().(*float64), short, defaultValue, "")
		f.flags = append(f.flags, set.Lookup(short))
		f.name = fmt.Sprintf("-%s", short)
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.Float64Var(field.Addr().Interface().(*float64), long, defaultValue, "")
		f.flags = append(f.flags, set.Lookup(short))
		f.name = fmt.Sprintf("--%s", long)
	}

	alias, ok := tags.Lookup("alias")
	if ok {
		set.Float64Var(field.Addr().Interface().(*float64), alias, defaultValue, "")
		f.flags = append(f.flags, set.Lookup(alias))
		f.name = fmt.Sprintf("--%s", alias)
	}

	env, ok := tags.Lookup("env")
	if ok {
		envOpts := strings.Split(env, ",")

		for _, envOpt := range envOpts {
			envStr := os.Getenv(envOpt)
			if envStr != "" {
				envValue, err := strconv.ParseFloat(envStr, 64)
				if err != nil {
					return &Flag{}, fmt.Errorf("could not parse float64 environment variable %s value %q: %s", envOpt, envStr, err)
				}

				field.SetFloat(envValue)
				f.set = true
				break
			}
		}
	}

	_, f.required = tags.Lookup("required")

	return &f, nil
}
