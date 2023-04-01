package env

import (
	"os"
	"strconv"
)

func Get[T any](k string, d T) T {
	v, ok := os.LookupEnv(k)
	if !ok {
		return d
	}

	var y any
	switch x := any(d).(type) {
	case string:
		y = convString(v, x)
	case int:
		y = convInt(v, x)
	case bool:
		y = convBool(v, x)
	default:
		return d
	}

	return y.(T)
}

func convString(v, d string) string {
	return v
}

func convInt(v string, d int) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return d
	}
	return i
}

func convBool(v string, d bool) bool {
	i, err := strconv.ParseBool(v)
	if err != nil {
		return d
	}
	return i
}
