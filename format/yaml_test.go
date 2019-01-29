package format

import (
	"testing"
)

func TestYamlScalars(t *testing.T) {
	cases := []map[string]interface{}{
		{"foo": 1},
		{"foo": 1.0},
		{"foo": 1.234},
		{"foo": "hello"},
		{"foo": true},
		{"foo": false},
	}

	expecteds := []string{
		"foo: 1",
		"foo: 1",
		"foo: 1.234",
		"foo: hello",
		"foo: true",
		"foo: false",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestYamlList(t *testing.T) {
	cases := []map[string]interface{}{
		{"foo": []interface{}{}},
		{"foo": []interface{}{1}},
		{"foo": []interface{}{
			1,
			"foo",
			true,
		}},
		{"foo": []interface{}{
			[]interface{}{
				"foo",
				"bar",
			},
			"baz",
		}},
		{"foo": []interface{}{
			[]interface{}{
				[]interface{}{
					"foo",
					"bar",
				},
				"baz",
			},
			"quux",
		}},
		{"foo": []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz":  "quux",
				"mooz": "xyzzy",
			},
		}},
	}

	expecteds := []string{
		"foo: []",
		"foo:\n  - 1",
		"foo:\n  - 1\n  - foo\n  - true",
		"foo:\n  - - foo\n    - bar\n  - baz",
		"foo:\n  - - - foo\n      - bar\n    - baz\n  - quux",
		"foo:\n  - foo: bar\n  - baz: quux\n    mooz: xyzzy",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestYamlMap(t *testing.T) {
	cases := []map[string]interface{}{
		{},
		{
			"foo": "bar",
		},
		{
			"foo": "bar",
			"baz": "quux",
		},
		{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
			"quux": "mooz",
		},
		{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "quux",
				},
				"mooz": "xyzzy",
			},
			"alpha": "beta",
		},
		{
			"foo": []interface{}{
				"bar",
				"baz",
			},
			"quux": []interface{}{
				"mooz",
			},
		},
	}

	expecteds := []string{
		"{}",
		"foo: bar",
		"baz: quux\n\nfoo: bar",
		"foo:\n  bar: baz\n\nquux: mooz",
		"alpha: beta\n\nfoo:\n  bar:\n    baz: quux\n\n  mooz: xyzzy",
		"foo:\n  - bar\n  - baz\n\nquux:\n  - mooz",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestCfnYaml(t *testing.T) {
	cases := []map[string]interface{}{
		{
			"Quux":       "mooz",
			"Parameters": "baz",
			"Foo":        "bar",
			"Resources":  "xyzzy",
		},
	}

	expecteds := []string{
		"Parameters: baz\n\nResources: xyzzy\n\nFoo: bar\n\nQuux: mooz",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestIntrinsicKey(t *testing.T) {
	cases := []map[string]interface{}{
		map[string]interface{}{
			"Ref": "foo",
		},
		map[string]interface{}{
			"Fn::Sub": "The cake is a lie",
		},
		map[string]interface{}{
			"Fn::NotARealFn": "This is not real but we'll take it",
		},
		map[string]interface{}{
			"Func::Join": "joined",
		},
		map[string]interface{}{
			"NoFunc": "Not a func for sure!",
		},
	}

	expecteds := []string{
		"Ref",
		"Fn::Sub",
		"Fn::NotARealFn",
		"",
		"",
	}

	for i, testCase := range cases {
		expected := expecteds[i]
		expectedOk := true

		if expected == "" {
			expectedOk = false
		}

		actual, actualOk := intrinsicKey(testCase)

		if actual != expected || actualOk != expectedOk {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestIntrinsics(t *testing.T) {
	cases := []map[string]interface{}{
		{
			"foo": map[string]interface{}{
				"Ref": "bar",
			},
		},
		{
			"foo": map[string]interface{}{
				"Fn::Sub": []interface{}{
					"The ${key} is a ${value}",
					map[string]interface{}{
						"key":   "cake",
						"value": "lie",
					},
				},
			},
		},
	}

	expecteds := []string{
		"foo: !Ref bar",
		"foo: !Sub\n  - The ${key} is a ${value}\n  - key: cake\n    value: lie",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestStrings(t *testing.T) {
	cases := []map[string]interface{}{
		{"foo": "foo"},
		{"foo": "*"},
		{"foo": "* bar"},
		{"foo": "2012-05-02"},
		{"foo": "today is 2012-05-02"},
		{"foo": ": thing"},
		{"foo": "Yes"},
		{"foo": "No"},
		{"foo": "multi\nline"},
	}

	expecteds := []string{
		"foo: foo",
		"foo: \"*\"",
		"foo: \"* bar\"",
		"foo: \"2012-05-02\"",
		"foo: today is 2012-05-02",
		"foo: \": thing\"",
		"foo: \"Yes\"",
		"foo: \"No\"",
		"foo: \"multi\\nline\"",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestYamlComments(t *testing.T) {
	data := map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"quux": "mooz",
		},
		"xyzzy": []interface{}{
			"lorem",
		},
	}

	commentCases := []map[interface{}]interface{}{
		{},
		{"foo": "This is bar"},
		{"baz": "This is baz"},
		{"baz": map[string]interface{}{"quux": "This is quux"}},
		// FIXME: Add a case for commenting on an element in a list
	}

	expecteds := []string{
		"baz:\n  quux: mooz\n\nfoo: bar\n\nxyzzy:\n  - lorem",
		"baz:\n  quux: mooz\n\nfoo: bar  # This is bar\n\nxyzzy:\n  - lorem",
		"baz:  # This is baz\n  quux: mooz\n\nfoo: bar\n\nxyzzy:\n  - lorem",
		"baz:\n  quux: mooz  # This is quux\n\nfoo: bar\n\nxyzzy:\n  - lorem",
	}

	for i, comments := range commentCases {
		expected := expecteds[i]

		actual := YamlWithComments(data, comments)

		if actual != expected {
			t.Errorf("from %q != %q\n", actual, expected)
		}
	}
}
