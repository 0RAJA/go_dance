package source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Sougou struct{}

func (v *Sougou) Name() string {
	return "搜狗翻译"
}

type SougouDictRequest struct {
	From string
	To   string
	Text string
}

type SougouDictResponse struct {
	Zly     string `json:"zly"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	UUID    string `json:"uuid"`
	Sugg    []struct {
		K string `json:"k"`
		V string `json:"v"`
	} `json:"sugg"`
	Direction string `json:"direction"`
}

func (v *Sougou) Transform(word, homeLanguage, targetLanguage string) (ret []string, err error) {
	client := &http.Client{Timeout: 10 * time.Second}
	request := &SougouDictRequest{
		From: homeLanguage,
		To:   "zh-CHS",
		Text: word,
	}
	var data = fmt.Sprintf(`from=%s&to=%s&client=web&text=%s&uuid=368d7543-0577-402a-a2a2-256ec9e4d039&pid=sogou-dict-vr&addSugg=on`, request.From, request.To, request.Text)
	reader := strings.NewReader(data)
	req, err := http.NewRequest("POST", "https://fanyi.sogou.com/reventondc/suggV3", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "ABTEST=0|1651826306|v17; SNUID=2FF0085C282CF77BE068A1312970164B; IPLOC=CN6101; SUID=07D82075D352A00A000000006274DE82; wuid=1651826306938; FUV=3d06cd19908956588c14ec6780f29c8a; translate.sess=2e2b9d40-26ed-43dc-9f72-369388ee7065; SUV=1651826308137; SGINPUT_UPSCREEN=1651826308163")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://fanyi.sogou.com")
	req.Header.Set("Referer", "https://fanyi.sogou.com/text")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.4985.0 Safari/537.36 Edg/102.0.1235.1")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="102", "Microsoft Edge";v="102"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client send request: %v", err)
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response body: %v", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status: %d ,body:%v", resp.StatusCode, response)
	}
	var dictResponse SougouDictResponse
	if err := json.Unmarshal(response, &dictResponse); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	for _, v := range dictResponse.Sugg {
		ret = append(ret, fmt.Sprintf("%s:%s", v.K, v.V))
	}
	return
}
