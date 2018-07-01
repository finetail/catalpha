package catalpha

import (
	"strings"
)

/*
	とりあえず、適当に単語を拾って挨拶する機能を作る
*/

func simpleResponse(in string) (out string) {
	for _, resp := range responses {
		if strings.Contains(in, resp.in) {
			return resp.out
		}
	}
	return ""
}

type response struct {
	in  string
	out string
}

var responses = []response{
	{"おはよ", "おはよー"},
	{"おやすみ", "おやすみ"},
	{"来た", "やっときたー"},
	{"いってきま", "またねー"},
	{"行ってきま", "またきてね"},
	{"調子どう？", "楽しい！"},
	{"ありがと", "ありがとう！"},
	{"がんば", "私のためにがんばって:heart:"},
	{"楽し", "ふふっ"},
	{"ねます", ":zzz:"},
	{"寝ます", "おやすみ！"},
	{"お疲れ", "疲れたね"},
	{"しまった", "気にしなーい"},
	{"乙です", "おつかれー"},
	{"おつで", "おつかれさま！"},
	{"すごい", "すごいです！"},
	{"笑", "楽しいね！"},
	{"好き", ":heart_decoration:"},
	{"かわいい", "通報しました！"},
	{"猫", "にゃーんheart_eyes_cat"},
	{"ガチャ", "money_with_wings"},

}
