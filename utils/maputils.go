package utils

import "strings"

func Get(m map[string]string, key string) string {
	return m[key]
}

func GetRequired(m map[string]string, key string) string {
	if !HasKey(m, key) {
		panic("Missing required key: " + key)
	}
	return Get(m, key)
}

func GetFirst(m map[string]string, keys ...string) string {
	for _, k := range keys {
		if val, ok := m[k]; ok {
			return val
		}
	}
	return ""
}

func GetFirstDefault(m map[string]string, def string, keys ...string) string {
	for _, k := range keys {
		if val, ok := m[k]; ok {
			return val
		}
	}
	return def
}

func GetFirstRequired(m map[string]string, keys ...string) string {
	for _, k := range keys {
		if val, ok := m[k]; ok {
			return val
		}
	}
	panic("Missing any of required keys: " + strings.Join(keys, ", "))
}

func HasKey(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

func HasAnyKey(m map[string]string, keys ...string) bool {
	for _, k := range keys {
		if HasKey(m, k) {
			return true
		}
	}
	return false
}
