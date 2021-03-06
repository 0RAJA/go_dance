package source

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Caiyun struct {
}

type CaiyunDictRequest struct {
	TransType string `json:"trans_type,omitempty"`
	Source    string `json:"source,omitempty"`
}

type CaiyunDictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func (c *Caiyun) Transform(word, homeLanguage, targetLanguage string) (ret []string, err error) {
	client := &http.Client{Timeout: time.Second * 10}
	request := &CaiyunDictRequest{
		TransType: homeLanguage + "2" + targetLanguage,
		Source:    word,
	}
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %v", err)
	}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", reader)
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.4985.0 Safari/537.36 Edg/102.0.1235.1")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
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
	var dictResponse CaiyunDictResponse
	if err := json.Unmarshal(response, &dictResponse); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %v", err)
	}
	ret = append(ret, fmt.Sprintf("UK: %s US: %s", dictResponse.Dictionary.Prons.En, dictResponse.Dictionary.Prons.EnUs))
	for _, item := range dictResponse.Dictionary.Explanations {
		ret = append(ret, item)
	}
	return
}

func (c *Caiyun) Name() string {
	return "????????????"
}
