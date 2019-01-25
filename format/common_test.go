package format

import (
	"testing"
)

func TestIndent(t *testing.T) {
	cases := []string{
		"foo",
		"  bar  ",
		"foo\nbar",
		"foo\nbar\n",
		"\nfoo\n  bar  \nquux",
	}

	expecteds := []string{
		"foo",
		"bar  ",
		"foo\n  bar",
		"foo\n  bar\n",
		"\n  foo\n    bar  \n  quux",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := indent(testCase)

		if actual != expected {
			t.Errorf("%q != %q\n", actual, expected)
		}
	}
}

func TestIntrinsicKeys(t *testing.T) {
	cases := []map[string]interface{}{
		{"foo": "bar"},
		{"Ref": "cake"},
		{"Fn::IncludeIf": "banana"},
		{"Fn::IncludeIf": "banana", "Ref": "cake"},
	}

	expecteds := []string{
		"",
		"Ref",
		"Fn::IncludeIf",
		"",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual, _ := intrinsicKey(testCase)

		if actual != expected {
			t.Errorf("%q != %q\n", actual, expected)
		}
	}
}

func TestFormatString(t *testing.T) {
	cases := []string{
		"foo",
		"\"quoted\"",
	}

	expecteds := []string{
		"foo",
		"\"\\\"quoted\\\"\"",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := formatString(testCase)

		if actual != expected {
			t.Errorf("%q != %q\n", actual, expected)
		}
	}
}
