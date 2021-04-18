package article

import (
	"blackspeak/sentence"
	"math/rand"
	"time"
)

// 文章结构体
type article struct{}

// 随机生成文章 参数 文章中含有多少句子
func RandomArticle(sentenceNum int) (string, error) {
	var res string
	var err error

	for i := 0; i < sentenceNum; i++ {
		rand.Seed(time.Now().UnixNano())
		var tmp string
		if tmp, err = sentence.RandomSentence(true, rand.Intn(3)+1); err != nil {
			return "", err
		}
		res = res + tmp
	}

	return res, nil
}
