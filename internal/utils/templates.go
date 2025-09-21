package utils

import (
	twmerge "github.com/Oudwins/tailwind-merge-go"
	"github.com/a-h/templ"
)

// Merge & resolve all the classes into a single string.
func MergeClasses(classes ...string) string {
	return twmerge.Merge(classes...)
}

// Return value if condition is true, else retruns an empty value of type T
func If[T comparable](cond bool, value T) T {
	var initialValue T
	if cond {
		return value
	}
	return initialValue
}

// return trueValue if condition is true else return falseValue.
func IfElse[T any](cond bool, trueValue T, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}

// Merge multiple instances of templ.Attributes into one.
func MergeAttrs(attrs ...templ.Attributes) templ.Attributes {
	mergedAttrs := templ.Attributes{}
	for _, attr := range attrs {
		for k, v := range attr {
			mergedAttrs[k] = v
		}
	}
	return mergedAttrs
}
