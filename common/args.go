package common

import (
	"os"
	"strings"
)

type Arguments struct {
	keyValues map[string]string
}

func Init() Arguments {
	return Arguments{
		keyValues: make(map[string]string),
	}
}

func (args *Arguments) GetKeyValues() map[string]string {
	return args.keyValues
}

func (args *Arguments) Set(key string, val string) {
	args.keyValues[key] = val
}

func (args *Arguments) Get(key string) string {
	return args.keyValues[key]
}

func (args *Arguments) GetRequired(key string) string {
	if !args.HasKey(key) {
		panic("Missing required key: " + key)
	}
	return args.keyValues[key]
}

func (args *Arguments) GetFirst(keys ...string) string {
	for _, k := range keys {
		if val, ok := args.keyValues[k]; ok {
			return val
		}
	}
	return ""
}

func (args *Arguments) GetFirstDefault(def string, keys ...string) string {
	for _, k := range keys {
		if val, ok := args.keyValues[k]; ok {
			return val
		}
	}
	return def
}

func (args *Arguments) GetFirstRequired(keys ...string) string {
	for _, k := range keys {
		if val, ok := args.keyValues[k]; ok {
			return val
		}
	}
	panic("Missing any of required keys: " + strings.Join(keys, ", "))
}

func (args *Arguments) HasKey(key string) bool {
	_, ok := args.keyValues[key]
	return ok
}

func (args *Arguments) HasAnyKey(keys ...string) bool {
	for _, k := range keys {
		if args.HasKey(k) {
			return true
		}
	}
	return false
}

func GetArgsRaw() []string {
	return os.Args[1:]
}

func GetArgs() Arguments {
	return ParseArgs(GetArgsRaw())
}

func ParseArgs(args []string) Arguments {
	arguments := Init()
	currentKey := ""
	for _, a := range args {
		if strings.HasPrefix(a, "--") {
			s := strings.Split(a, "=")
			k := strings.TrimLeft(s[0], "--")
			if len(s) > 1 {
				arguments.Set(k, s[1])
			} else {
				arguments.Set(k, "")
			}
		} else if strings.HasPrefix(a, "-") {
			k := strings.TrimLeft(a, "--")
			arguments.Set(k, "")
			currentKey = k
		} else if currentKey != "" {
			arguments.Set(currentKey, a)
			currentKey = ""
		}
	}
	return arguments
}
