package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)


//股票基本信息
type stockInfo struct {
	name string  //东风汽车
	code string  //600006
	sinaCode string //sh600006

}

//查询接口
func search(keyword string) []*stockInfo {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://suggest3.sinajs.cn/suggest/name=info&key="+keyword, nil)
	rep, _ := client.Do(req)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	str := ConvertByte2String(body, GB18030)
	var stockList = []*stockInfo{}
	if len(str)>15 {
		list := strings.Split(str[10:], ";")
		for _, v := range list {
			infos := strings.Split(v,",")
			if len(infos)>1 {
				if infos[1]=="11" {
					stockList=append(stockList, &stockInfo{name:infos[0],code:infos[2],sinaCode:infos[3]})
				}
			}


		}
	}

	return stockList


}

