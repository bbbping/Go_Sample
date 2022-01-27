package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strings"
)

type stock struct {
	name string
	code string
	sinaCode string
	price string
	//变化百分比
	percent string
	//每股涨跌值，单位元
	change string

}

func getSinglePrice(sinaCode string)stock{
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://hq.sinajs.cn/list=s_"+sinaCode, nil)
	rep, _ := client.Do(req)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	str := ConvertByte2String(body, GB18030)
	list := strings.Split(str, ",")
	colon_index:=strings.Index(list[0],`"`)
	code_index:=strings.LastIndex(list[0],"_")+3
	return stock{
		name: list[0][colon_index+1:],
		price: list[1],
		percent: list[3],
		code:list[0][code_index:code_index+6],
	}
}

func getPrice4(sinaCode []string)[]stock{
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://hq.sinajs.cn/list=s_"+strings.Join(sinaCode,",s_"), nil)
	rep, _ := client.Do(req)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	str := ConvertByte2String(body, GB18030)
	fmt.Println("结果：",str)
	ls:=strings.Split(str,";")
	var result = []stock{}

	for _, v := range ls {
		if len(v)>15 {
			list := strings.Split(v, ",")
			colon_index:=strings.Index(list[0],`"`)
			code_index:=strings.LastIndex(list[0],"_")+3
			result = append(result,stock{
				name: list[0][colon_index+1:],
				price: list[1],
				change: list[2],
				percent: list[3],
				code:list[0][code_index:code_index+6],
				sinaCode:list[0][code_index-2:code_index+6],
			} )
		}

	}

	return result

}

func getSZ() stock {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://hq.sinajs.cn/list=s_sz399001", nil)
	rep, _ := client.Do(req)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	str := ConvertByte2String(body, GB18030)
	list := strings.Split(str, ",")
	colon_index:=strings.Index(list[0],`"`)
	return stock{
		name: list[0][colon_index+1:],
		price: list[1],
		percent: list[3],
	}
}

func getCY() stock {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://hq.sinajs.cn/list=s_sz399006", nil)
	rep, _ := client.Do(req)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	str := ConvertByte2String(body, GB18030)
	list := strings.Split(str, ",")
	colon_index:=strings.Index(list[0],`"`)
	return stock{
		name: list[0][colon_index+1:],
		price: list[1],
		percent: list[3],
	}
}

func getSH() stock {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://hq.sinajs.cn/list=s_sh000001", nil)
	rep, _ := client.Do(req)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	str := ConvertByte2String(body, GB18030)
	list := strings.Split(str, ",")
	colon_index:=strings.Index(list[0],`"`)
	return stock{
		name:list[0][colon_index+1:],
		price: list[1],
		percent: list[3],
	}
}



type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}


