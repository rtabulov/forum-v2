package helpers

import (
	"strings"
)

// MatchAndParams func
func MatchAndParams(pattern, path string) (bool, map[string]string) {
	params := map[string]string{}

	patternFields := strings.Split(pattern, "/")
	pathFields := strings.Split(path, "/")

	if len(patternFields) != len(pathFields) {
		return false, nil
	}

	for i := range patternFields {
		if patternFields[i] == pathFields[i] {
			continue
		}

		if param, ok := getParam(patternFields[i]); ok {
			params[param] = pathFields[i]
			continue
		}

		return false, nil
	}

	return true, params
}

func getParam(segment string) (string, bool) {
	if strings.HasPrefix(segment, ":") {
		return strings.TrimPrefix(segment, ":"), true
	}

	return "", false
}

// MergeStringMaps func
func MergeStringMaps(maps ...map[string]string) map[string]string {
	res := map[string]string{}

	for _, mp := range maps {
		for key, value := range mp {
			res[key] = value
		}
	}

	return res
}
