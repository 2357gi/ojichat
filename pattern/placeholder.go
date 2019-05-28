package pattern

import (
	"math/rand"
	"strings"
	"time"
)

// 文章中一種類に統一されるタグ
var uniqTags = map[string][]string{
	// 対象の名前
	"{TARGET_NAME}": []string{
		"優子",
	},
	// おじさんの一人称
	"{FIRST_PERSON}": []string{
		"僕",
		"ボク",
		"俺",
		"オレ",
		"小生",
	},
	// 曜日
	"{DAY_OF_WEEK}": []string{
		"月", "火", "水", "木", "金", "土", "日",
	},
	// 食べ物
	"{FOOD}": []string{
		"お寿司🍣",
		"イタリアン🍕🍝",
		"パスタ🍝",
	},
}

// 文章中複数回変更&繰り返されるタグ
var flexTags = map[string][]string{
	// ポジティブな表現の絵文字/顔文字
	"{EMOJI_POS}": []string{
		"❗",
		"(^_^)",
	},
	// ネガティヴな表現の絵文字/顔文字
	"{EMOJI_NEG}": []string{
		"(T_T)",
		"💦",
	},
	// ニュートラルな感情を表す絵文字/顔文字
	"{EMOJI_NEUT}": []string{
		"(^^;;",
		"💤",
		"😴",
	},
	// 疑問を投げかけるときに利用される絵文字/顔文字
	"{EMOJI_ASK}": []string{
		"❓",
	},
}

// ConvertTags ; message内にあるタグを置換する
func ConvertTags(message string) string {
	rand.Seed(time.Now().UnixNano())
	for tag, pat := range uniqTags {
		content := pat[rand.Intn(len(pat))]
		message = strings.Replace(message, tag, content, -1)
	}
	return message
}
