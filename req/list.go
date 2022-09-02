package req

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hdow/glo"
	"hdow/utils"
	"io/ioutil"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

var geturl = "https://mhapp.vxnbbrs.xyz/api/video/classify/getClassifyVideos"

func Getlist(pagesize string, classid string, sort string) ([]map[string]string, error) {
	new := time.Now().UnixNano()
	str := strconv.FormatInt(new, 10)
	md5 := utils.Gedmd5([]byte(str[3:8]))
	u := url.Values{}
	u.Add("pagesize", pagesize)
	u.Add("classifyId", classid)
	u.Add("sort", sort)
	a := u.Encode()
	header := make(map[string]string)
	header["User-Agent"] = "Mozilla/5.0 (Linux; Android 11; Subsystem for Android(TM) Build/RD2A.211001.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.120 Mobile Safari/537.36;SuiRui/mh/ver=1.0.2"
	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	header["Host"] = "mhapp.vxnbbrs.xyz"
	header["Connection"] = "Keep-Alive"
	header["Accept-Encoding"] = "gzip"
	header["s"] = md5
	header["t"] = str
	header["Authorization"] = glo.GloToken
	res, err := utils.Httpget(geturl+"?"+a, header)
	if err != nil {
		return nil, err
	}
	date, _ := ioutil.ReadAll(res.Body)
	w := make(map[string]string)
	err = json.Unmarshal(date, &w)
	if err != nil {

	}
	wout, _ := base64.StdEncoding.DecodeString(w["encData"])
	aaa, _ := utils.AesDecrypt(wout, utils.PwdKey)
	aaas := string(aaa)
	tire, err := regexp.Compile("\"title.*?\":\".*?\"")
	idre, err := regexp.Compile("\"videoId\":[0-9]*,")
	if err != nil {
	}
	var mapvdid []map[string]string
	var mapvdid1 []map[string]string
	vdid := make(map[string]string)
	var tiok []string
	tiok = tire.FindAllString(aaas, -1)
	fmt.Println(len(tiok))
	var idok []string
	idok = idre.FindAllString(aaas, -1)
	for k, _ := range tiok {
		vdid["title"] = tiok[k][9 : len(tiok[k])-1]
		vdid["id"] = idok[k][10 : len(idok[k])-1]
		fmt.Println(vdid)
		mapvdid = append(mapvdid1, vdid)

	}
	fmt.Println(mapvdid)
	return mapvdid, nil
}
