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
	}

	expecteds := []string{
		"[]",
		"- 1",
		"- 1\n- foo\n- true",
		"- - foo\n  - bar\n- baz",
		"- - - foo\n    - bar\n  - baz\n- quux",
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
	}

	expecteds := []string{
		"{}",
		"foo: bar",
		"baz: quux\nfoo: bar",
		"foo:\n  bar: baz\nquux: mooz",
		"alpha: beta\nfoo:\n  bar:\n    baz: quux\n  mooz: xyzzy",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}
