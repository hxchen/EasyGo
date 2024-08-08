package main

import "fmt"

func main() {
	var siteMap map[string]string
	siteMap = make(map[string]string)
	siteMap["Google"] = "http://www.google.com"
	siteMap["Runoob"] = "http://www.runoob.com"
	siteMap["Taobao"] = "http://www.taobao.com"
	siteMap["Zhihu"] = "http://www.zhihu.com"
	siteMap["Weibo"] = "http://www.weibo.com"

	for key, value := range siteMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// 判断某个键是否存在
	site := siteMap["Google"]
	fmt.Println("site:", site)

	site, ok := siteMap["Facebook"]
	if ok {
		fmt.Println("site:", site)
	} else {
		fmt.Println("site not found")
	}
}
