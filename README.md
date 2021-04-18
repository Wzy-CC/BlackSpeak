# BlackSpeak
中文互联网黑话生成器/不完全指南

sample:调用RandomArticle方法生成任意句子数量的文章

```golang
	a, err := article.RandomArticle(20)
	if err != nil {
		panic("error")
	}
	fmt.Println(a)
```

更新日志
v0.1功能:支持自定义模版文章，自定义模版句式，自定义短语
后期支持功能:支持任意的生成文章，非模版
