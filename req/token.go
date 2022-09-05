package req

import (
	"bytes"
	"encoding/json"
	"hdow/glo"
	"hdow/utils"
	"io/ioutil"
	"strconv"
	"time"
)

type Date struct {
	DeviceId string `json:"deviceId"`
	Dev      string `json:"dev"`
	ChCode   string `json:"chCode"`
}

func Gettoken() {
	url := "https://mhapp.vxxsred.xyz/api/user/traveler"
	new := time.Now().UnixNano()
	str := strconv.FormatInt(new, 10)
	md5 := utils.Gedmd5([]byte(str[3:8]))
	date := Date{
		DeviceId: "a99a87c3bf3407317dc78278a31613l6",
		Dev:      "S9S",
		ChCode:   "mlstg",
	}
	body, _ := json.Marshal(date)
	header := make(map[string]string)
	header["User-Agent"] = "Mozilla/5.0 (Linux; Android 11; Subsystem for Android(TM) Build/RD2A.211001.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.120 Mobile Safari/537.36;SuiRui/mh/ver=1.0.2"
	header["Accept"] = "application/json"
	header["Content-Type"] = "application/json"
	header["deviceId"] = "a99a87c3bf3407317dc78278a31613l6"
	header["Connection"] = "Keep-Alive"
	header["Accept-Encoding"] = "gzip"
	header["s"] = md5
	header["t"] = str
	res, _ := utils.Httppost(url, header, bytes.NewReader(body))
	a, _ := ioutil.ReadAll(res.Body)
	w := make(map[string]interface{})
	err := json.Unmarshal(a, &w)
	if err != nil {

	}
	token := w["data"].(map[string]interface{})["token"]
	glo.GloToken = token.(string)
}
