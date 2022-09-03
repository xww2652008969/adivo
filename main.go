package main

import (
	"fmt"
	"hdow/glo"
	"hdow/req"
	"hdow/utils"
)

func main() {
	var sc string
	//fmt.Println(req.Getlist("30", "6", "2"))
	//fmt.Println(req.Getvedio("32033"))
	////for i := 16125; i < 40000; i++ {
	////	utils.Writefile("date\\"+strconv.Itoa(i), []byte(req.Getvedio(strconv.Itoa(i))))
	//}
	req.Gettoken()                            //初始化
	utils.PwdKey = []byte(glo.GloToken[2:18]) //初始化密钥
	fmt.Println("输入1获取列表，输入2下载东西")
	////fmt.Println(req.Getlist("30", "20", "1"))
	//fmt.Println(req.Getvedio("58647"))
	//req.Redts("https://vedss.ilovyo.cn/lfb/o0/w0/4f/v1/5eeba35c9c7a40aeaa6b1b19f37a3d38.m3u8")
	fmt.Println(glo.GloToken)
	for {
		fmt.Scanln(&sc)
		if sc == "1" {
			fmt.Println("你输入1-30试试")
			fmt.Scanln(&sc)
			fmt.Println(req.Getlist("30", sc, "1")) //获取列表
			sc = ""
		}
		if sc == "2" {
			fmt.Println("你输入刚才获取的id试试")
			fmt.Scanln(&sc)
			req.Redts(req.Getvedio(sc))
			fmt.Println("你看是不是有视频")
		}
	}
}
