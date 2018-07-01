package catalpha

import (
	"strings"
	"testing"
)

func TestSimpleResponse(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"おはよう！", "おはようございます"},
		{"今日もおはよう", "おはようございます"},
		{"博多にいってきました", "いってらっしゃい"},
		{"行ってきましょう", "いってらっしゃい"},
		{"おはよおやすみきた", "おはようございます"},
		{"お、はよう", ""},
		{"", ""},
		{strings.Repeat("--------", 2048), ""},
	}
	for i, tt := range testCases {
		if actual := simpleResponse(tt.in); actual != tt.out {
			t.Errorf("case: %d, expected: %s, actual: %s", i, tt.out, actual)
		}
	}
}
