package string_diff_matcher

import (
	"fmt"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const defaultStyle = "\x1b[0m"
const redColor = "\x1b[91m"

func EqualWithDiff(expected interface{}) types.GomegaMatcher {
	return &equalWithDiffMatcher{
		expected: expected,
	}
}

type equalWithDiffMatcher struct {
	expected interface{}
}

func (matcher *equalWithDiffMatcher) Match(actual interface{}) (success bool, err error) {
	em := matchers.EqualMatcher{Expected: matcher.expected}
	return em.Match(actual)
}

func (matcher *equalWithDiffMatcher) FailureMessage(actual interface{}) (message string) {
	dmp := diffmatchpatch.New()
	as, _ := actual.(string)
	es, _ := matcher.expected.(string)
	diffs := dmp.DiffMain(es, as, true)
	return fmt.Sprintf("Expected: \n%s\n%s \n%sto match but found diff: %s\n\n\x1b[0m%s",
		defaultStyle, actual, redColor, defaultStyle, dmp.DiffPrettyText(diffs))
}

func (matcher *equalWithDiffMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to equal", matcher.expected)
}
