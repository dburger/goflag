/*
Package goflag provides additional flag utilities.
*/
package goflag

import "strings"

type ArrayFlags []string

func (af *ArrayFlags) String() string {
	return strings.Join(*af, ",")
}

func (af *ArrayFlags) Set(value string) error {
	*af = append(*af, value)
	return nil
}
