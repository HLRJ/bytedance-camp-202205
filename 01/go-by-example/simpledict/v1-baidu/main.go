package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`kw=good`)
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/sug", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIDUPSID=9EF36AC727158B4F288254D6E4767C6C; PSTM=1647944009; BAIDUID=9EF36AC727158B4FEE4C343ECC25069A:FG=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BA_HECTOR=0g0g8h87ag0gag0lk41h7e5330q; BAIDUID_BFESS=0BA7A3E2C28AB3BB5DFD76C374703593:FG=1; delPer=0; PSINO=1; H_PS_PSSID=36309_31254_34813_35914_36165_34584_35979_36073_36055_36383_35801_36336_26350_36301_22159_36061; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1651977225; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1651977225; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; ab_sr=1.0.1_YTY4Njg0MWMyOGE3ODRmNjE1NzA0ZDY3MTczNDUzODVjMmY4NTNiZDVlNWVkNjc3MWMzZjE4MmI2MGU5N2NlOWYxMjVmZTIyYzk1MDk5OGUzYTlkNTI0MjVkNDE0NmI1MjIzYjcxZmFkNGM0ZGM3NWVjMDE0NDJhZTlmYmE0ZDMyMzMzMzA5YzlmMzMwZTAyZjZmMmEzZmM5ODNmYzk5Yw==")
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Referer", "https://fanyi.baidu.com/?aldtype=16047")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
