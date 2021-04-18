package sentence

import (
	"bytes"
	"math/rand"
	"text/template"
	"time"

	dict "blackspeak/dictionary"
)

// 句式
type sentence struct {
	V string // 动词
	N string // 名词
	P string // 短语

	N1 string // 名词
	N2 string // 名词
	N3 string // 名词

	V1 string // 动词
}

// 句式模版
var (
	// 1参数
	ST1 = []string{"亮点是{{.N}}", "达到{{.N}}标准", "优势是{{.N}}"}

	// 2参数
	ST2 = []string{"{{.N}}是{{.P}}", "{{.V}}行业{{.N}}", "{{.V}}整个{{.N}}", "{{.P}}作为{{.N}}的评判标准"}

	// 3参数
	ST3 = []string{"{{.V}}{{.N}}作为{{.P}}为产品赋能", "{{.N}}是{{.V}}{{.P}}", "通过{{.N}}和{{.P}}达到{{.V}}"}

	// 4参数
	ST4 = []string{"{{.N}}是{{.V}}{{.N1}}{{.N2}}", "{{.V}}{{.N1}}{{.V1}}{{.N2}}", "{{.N}}是在{{.N1}}采用{{.N2}}打法达成{{.N3}}"}
)

// 标点符号
var (
	q = []string{"。", "，"}
)

// 随机生成句式 参数 是否采用模版 几参数的句式
func RandomSentence(isTemplate bool, params int) (string, error) {
	if isTemplate {
		var st []string
		switch params {
		case 1:
			st = ST1
		case 2:
			st = ST2
		case 3:
			st = ST3
		}
		rand.Seed(time.Now().UnixNano())
		tmpl, err := template.New("sentence").Parse(st[rand.Intn(len(st))])
		if err != nil {
			return "", err
		}
		buf := new(bytes.Buffer)
		s := sentence{
			V:  dict.GetRandomV(),
			N:  dict.GetRandomN(),
			P:  dict.GetRandomP(),
			N1: dict.GetRandomN(),
			N2: dict.GetRandomN(),
			N3: dict.GetRandomN(),
			V1: dict.GetRandomV(),
		}
		if err := tmpl.Execute(buf, s); err != nil {
			return "", err
		}

		return buf.String() + q[rand.Intn(len(q))], nil
	}
	return "", nil
}
