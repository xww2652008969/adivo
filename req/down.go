package req

import (
	"encoding/base64"
	"fmt"
	"github.com/grafov/m3u8"
	"hdow/utils"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var key = "x3GZk8tbgc6xSPSiBdSPBQ=="

func Redts(url string) {
	s := strings.Split(url, "/")
	var news string
	for k, v := range s {
		if k == 0 {
			news = news + v + "//"
		} else if k > 1 && k < len(s)-1 {
			news = news + v + "/"
		}
	}
	var urllist []string
	res, _ := utils.Httpget(url, nil)
	p, listType, _ := m3u8.DecodeFrom(res.Body, true)
	switch listType {
	case m3u8.MEDIA:
		mediapl := p.(*m3u8.MediaPlaylist)
		for _, v := range mediapl.Segments {
			if v != nil {
				urllist = append(urllist, news+v.URI)
			}
		}
		fmt.Println(len(urllist))
	case m3u8.MASTER:
		masterpl := p.(*m3u8.MasterPlaylist)
		fmt.Printf("%+v\n", masterpl)
	}
	for k, v := range urllist {
		go do(k, v)
		fmt.Println(k, v)
	}
	time.Sleep(10)
	//var hda []byte
	//for k, v := range urllist {
	//	fmt.Println(k)
	//	res, _ = utils.Httpget(v, nil)
	//	utils.PwdKey, _ = base64.StdEncoding.DecodeString(key)
	//	d, _ := ioutil.ReadAll(res.Body)
	//	dde, _ := utils.AesDecrypt(d, utils.PwdKey)
	//	hda = append(hda, dde...)
	//}
	//
	//utils.Writefile(strconv.Itoa(rand.Int())+".mp4", hda)
	//fmt.Println("写入完成")
}
func do(k int, url string) {
	res, _ := utils.Httpget(url, nil)
	utils.PwdKey, _ = base64.StdEncoding.DecodeString(key)
	d, _ := ioutil.ReadAll(res.Body)
	dde, _ := utils.AesDecrypt(d, utils.PwdKey)
	utils.Writefile(strconv.Itoa(k)+".mp4", dde)
}
