package format

import (
	"testing"
)

func TestYamlScalars(t *testing.T) {
	cases := []interface{}{
		1,
		1.0,
		1.234,
		"hello",
		true,
		false,
	}

	expecteds := []string{
		"1",
		"1",
		"1.234",
		"hello",
		"true",
		"false",
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
	cases := []interface{}{
		[]interface{}{},
		[]interface{}{1},
		[]interface{}{
			1,
			"foo",
			true,
		},
		[]interface{}{
			[]interface{}{
				"foo",
				"bar",
			},
			"baz",
		},
		[]interface{}{
			[]interface{}{
				[]interface{}{
					"foo",
					"bar",
				},
				"baz",
			},
			"quux",
		},
		[]interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz":  "quux",
				"mooz": "xyzzy",
			},
		},
	}

	expecteds := []string{
		"[]",
		"- 1",
		"- 1\n- foo\n- true",
		"- - foo\n  - bar\n- baz",
		"- - - foo\n    - bar\n  - baz\n- quux",
		"- foo: bar\n- baz: quux\n\n  mooz: xyzzy",
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
	cases := []interface{}{
		map[string]interface{}{},
		map[string]interface{}{
			"foo": "bar",
		},
		map[string]interface{}{
			"foo": "bar",
			"baz": "quux",
		},
		map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
			"quux": "mooz",
		},
		map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "quux",
				},
				"mooz": "xyzzy",
			},
			"alpha": "beta",
		},
		map[string]interface{}{
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
	cases := []interface{}{
		map[string]interface{}{
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
	cases := []interface{}{
		map[string]interface{}{
			"foo": map[string]interface{}{
				"Ref": "bar",
			},
		},
		map[string]interface{}{
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
	cases := []string{
		"foo",
		"*",
		"* bar",
		"2012-05-02",
		"today is 2012-05-02",
		": thing",
		"Yes",
		"No",
		"multi\nline",
	}

	expecteds := []string{
		"foo",
		"\"*\"",
		"\"* bar\"",
		"\"2012-05-02\"",
		"today is 2012-05-02",
		"\": thing\"",
		"\"Yes\"",
		"\"No\"",
		"\"multi\\nline\"",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}

func TestLongStrings(t *testing.T) {
	cases := []string{}

	expecteds := []string{}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}
