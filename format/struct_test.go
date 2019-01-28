package format

import (
	"reflect"
	"testing"
)

var s Struct = Struct{
	Data: map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			"quux": []interface{}{
				"mooz",
			},
			"xyzzy": "lorem",
			"ipsum": "dolor", // Uncommented
		},
	},
	Comments: map[interface{}]interface{}{
		"":    "Root comment",
		"foo": "This is foo",
		"baz": map[interface{}]interface{}{
			"": "This is baz",
			"quux": map[interface{}]interface{}{
				"": "This is quux",
				0:  "This is quux[0]",
			},
			"xyzzy": "This is xyzzy",
		},
	},
}
var paths = [][]interface{}{
	{},
	{"foo"},
	{"baz"},
	{"baz", "quux"},
	{"baz", "quux", 0},
	{"baz", "xyzzy"},
	{"baz", "ipsum"},
}

func TestGetComment(t *testing.T) {
	expecteds := []string{
		"Root comment",
		"This is foo",
		"This is baz",
		"This is quux",
		"This is quux[0]",
		"This is xyzzy",
		"",
	}

	for i, path := range paths {
		expected := expecteds[i]

		actual := s.GetComment(path)

		if actual != expected {
			t.Errorf("%v != %v\n", actual, expected)
		}
	}
}

func TestGet(t *testing.T) {
	expecteds := []interface{}{
		s.Data,
		s.Data.(map[string]interface{})["foo"],
		s.Data.(map[string]interface{})["baz"],
		s.Data.(map[string]interface{})["baz"].(map[string]interface{})["quux"],
		s.Data.(map[string]interface{})["baz"].(map[string]interface{})["quux"].([]interface{})[0],
		s.Data.(map[string]interface{})["baz"].(map[string]interface{})["xyzzy"],
		s.Data.(map[string]interface{})["baz"].(map[string]interface{})["ipsum"],
	}

	for i, path := range paths {
		expected := expecteds[i]

		actual := s.Get(path)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v != %v\n", actual, expected)
		}
	}
}
