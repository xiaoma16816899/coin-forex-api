package types

import "strings"

// Rule "key:action"
type Permision string

func NewPermission(key string, action string) string {
	return strings.Join([]string{key, action}, ":")
}

func (p Permision) GetAction() string {
	return strings.Split(string(p), ":")[1]
}

func (p Permision) GetKey() string {
	return strings.Split(string(p), ":")[0]
}
