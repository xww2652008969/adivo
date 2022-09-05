package req

import (
	"encoding/base64"
	"fmt"
	"github.com/grafov/m3u8"
	"hdow/utils"
	"io/ioutil"
	"strings"
)

var key = "x3GZk8tbgc6xSPSiBdSPBQ=="

type vido struct {
	id   int
	date []byte
}

func Redts(url string) {
	utils.Createfolder("wwt")
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
	var vida = make([]vido, len(urllist))
	ch := make(chan int)
	for k, v := range urllist {
		go do(k, v, ch, vida)
	}
	var a int
	for true {
		if a < len(urllist) {
			a = a + <-ch
		} else {
			var datee []byte
			for _, vv := range vida {
				datee = append(datee, vv.date...)
			}
			utils.Writefile("wwt/"+"xww.mp4", datee)
			break
		}
	}
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
func do(k int, url string, ch chan int, vi []vido) {
	res, _ := utils.Httpget(url, nil)
	utils.PwdKey, _ = base64.StdEncoding.DecodeString(key)
	d, _ := ioutil.ReadAll(res.Body)
	dde, _ := utils.AesDecrypt(d, utils.PwdKey)
	vi[k].id = k
	vi[k].date = dde
	ch <- 1
}
