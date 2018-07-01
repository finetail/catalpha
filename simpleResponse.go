package catalpha

import "strings"

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
	{"おはよ", "おはようございます"},
	{"おやすみ", "おやすみなさい"},
	{"きた", "いらっしゃいませ"},
	{"いってきま", "いってらっしゃい"},
	{"行ってきま", "いってらっしゃい"},
	{"調子どう？", "良いです"},
	{"ありがと", "どういたしまして"},
	{"がんば", "がんばって"},
	{"楽し", "楽しいです"},
	{"ねます", "おやすみなさい"},
	{"寝ます", "おやすみなさい"},
	{"お疲れ", "お疲れ様です"},
	{"しまった", "どんまい"},
	{"乙です", "お疲れ様です"},
	{"おつで", "お疲れ様です"},
	{"すごい", "すごいです"},
	{"笑", "wwwwwwwwww"},
}
