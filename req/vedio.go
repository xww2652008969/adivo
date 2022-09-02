package req

import (
	"encoding/base64"
	"encoding/json"
	"hdow/glo"
	"hdow/utils"
	"io/ioutil"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

var vidourl = "https://mhapp.vxerfxs.shop/api/video/getVideoById?"

func Getvedio(id string) string {
	new := time.Now().UnixNano()
	str := strconv.FormatInt(new, 10)
	md5 := utils.Gedmd5([]byte(str[3:8]))
	u := url.Values{}
	u.Add("videoId", id)
	a := u.Encode()
	header := make(map[string]string)
	header["User-Agent"] = "Mozilla/5.0 (Linux; Android 11; Subsystem for Android(TM) Build/RD2A.211001.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.120 Mobile Safari/537.36;SuiRui/mh/ver=1.0.2"
	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	header["deviceId"] = "a99a87c3bf3307317dc78278a3161375"
	header["Host"] = "mhapp.vxnbbrs.xyz"
	header["Connection"] = "Keep-Alive"
	header["Accept-Encoding"] = "gzip"
	header["s"] = md5
	header["t"] = str
	header["Authorization"] = glo.GloToken
	res, _ := utils.Httpget(vidourl+a, header)
	date, _ := ioutil.ReadAll(res.Body)
	w := make(map[string]string)
	err := json.Unmarshal(date, &w)
	if err != nil {

	}
	bade, _ := base64.StdEncoding.DecodeString(w["encData"])
	aaa, _ := utils.AesDecrypt(bade, utils.PwdKey)
	m3u8re, err := regexp.Compile("l\":\".*pre")
	out := m3u8re.FindString(string(aaa))
	return "https://vedss.ilovyo.cn/" + out[4:len(out)-6]
	//return string(aaa)
}
