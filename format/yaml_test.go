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
		"- foo: bar\n- baz: quux\n  mooz: xyzzy",
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
		"baz: quux\nfoo: bar",
		"foo:\n  bar: baz\nquux: mooz",
		"alpha: beta\nfoo:\n  bar:\n    baz: quux\n\n  mooz: xyzzy",
		"foo:\n  - bar\n  - baz\nquux:\n  - mooz",
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
		"Parameters: baz\nResources: xyzzy\nFoo: bar\nQuux: mooz",
	}

	for i, testCase := range cases {
		expected := expecteds[i]

		actual := Yaml(testCase)

		if actual != expected {
			t.Errorf("from %T %v:\n%#v != %#v\n", testCase, testCase, actual, expected)
		}
	}
}
