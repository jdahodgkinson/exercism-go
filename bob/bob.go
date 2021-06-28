// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Determines if a given remark is a silence.
func IsSilence(remark string) bool {
	for i := 0; i < len(remark); i++ {
		if string(remark[i]) != " " {
			return false
		}
	}
	return true
}

// Determines if a given remark is a question.
func IsQuestion(remark string) bool {
	n := len(remark)
	if n > 0 {
		return string(remark[n-1]) == "?"
	} else {
		return false
	}
}

// Determines if a given remark is a yell.
func IsYelling(remark string) bool {
	isSameAsLower := remark == strings.ToLower(remark)
	isSameAsUpper := remark == strings.ToUpper(remark)
	isYelling := !isSameAsLower && isSameAsUpper
	return isYelling
}

// Returns Bob's appropriate response to a given remark.
func Hey(remark string) string {
	var response string
    isSilence := IsSilence(remark)
	isQuestion := IsQuestion(remark)
	isYelling := IsYelling(remark)
	if isSilence {
		response = "Fine. Be that way!"
	} else if isQuestion && isYelling {
		response = "Calm down, I know what I'm doing!"
	} else if isQuestion {
		response = "Sure."
	} else if isYelling {
		response = "Whoa, chill out!"
	} else {
		response = "Whatever."
	}
	return response
}
