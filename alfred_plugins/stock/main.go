package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var COMMAND_MAP = make(map[string]Handler)

const patt ="^l[0-9]$"

type Handler interface  {
	Handle()
}


func init(){
	COMMAND_MAP["01"]=&Index{}
	COMMAND_MAP["add"]=&Add{}
	COMMAND_MAP["del"]=&Del{}
	COMMAND_MAP["l"]=&Add{}

}



func main() {
	//解析命令
	query := os.Args[1]
	handler,ok := COMMAND_MAP[query]
	matched, _ := regexp.MatchString(patt, query)
	if ok {
		handler.Handle()
	} else if matched  {
		//查询自选的
		codes := ReadFromFile("./optional_" + query)
		list := strings.Split(codes, ",")
		rows := getPriceBySindCodes(list[0 : len(list)-1])
		result := &Result{rows: rows}
		fmt.Println(result.GetResult())

	} else {
		//查询股票信息，例如code这些
		infos := search(query)
		s := Stock{}
		//获取报价
		s.Price(infos)
	}
	
}

type Stock struct {

}


func (this *Stock) Price( sinfo []*stockInfo){

	var sinaCodes = []string{}
	for _, v := range sinfo {
		sinaCodes = append(sinaCodes,v.sinaCode )
	}
	rows := getPriceBySindCodes(sinaCodes)
	result := &Result{rows: rows}
	fmt.Println(result.GetResult())

}

func getPriceBySindCodes(sinaCodes []string) []*Row {
	rows := []*Row{}
	price4 := getPrice4(sinaCodes)
	for _, p := range price4 {
		row := &Row{title: p.name + " " + p.code, subtitle: "涨跌百分比:" + p.percent + "% 现价:" + p.price +" 涨跌值:" + p.change, arg: p.sinaCode}
		setColor(p.percent, row)
		rows = append(rows, row)
	}
	return rows
}

//根据涨跌设置颜色
func setColor( percent string, row *Row) {
	percent_num, _:= strconv.ParseFloat(percent, 10)
	if percent_num > 0 {
		row.icon = 1
	} else if percent_num < 0 {
		row.icon = -1
	} else {
		row.icon = 0
	}
}




type Index struct {}

func (this *Index) Handle(){
	sh := getSH()
	sz :=  getSZ()
	cy :=  getCY()
	rows:=[]*Row{}

	sh_r := &Row{title: sh.name, subtitle: "涨跌:" + sh.percent + " 指数:" + sh.price}
	setColor(sh.percent, sh_r)
	list := append(rows, sh_r)

	sz_r := &Row{title: sz.name, subtitle: "涨跌:" + sz.percent + " 指数:" + sz.price}
	setColor(sz.percent, sz_r)
	list = append(list, sz_r)

	cy_r := &Row{title: cy.name, subtitle: "涨跌:" + cy.percent + " 指数:" + cy.price}
	setColor(cy.percent, cy_r)
	list = append(list, cy_r)
	result := &Result{rows: list}
	fmt.Println(result.GetResult())

}

type Add struct {}
//添加到自选
func (this *Add) Handle(){
	if len(os.Args)<5 { //1是 add 、2是l1\l2\l3.... 3是股票名称 4是end字符
		StandOutput("参数不够")
		return
	}
	if strings.ToLower(os.Args[4])!="end" {
		StandOutput("没结束符")
		return
	}
	result := ReadFromFile("./optional_"+os.Args[2])
	if strings.Contains(result,os.Args[3]){
		StandOutput("已经存在")
		return
	}else {
		result+=(os.Args[3]+",")
	}
	ExportToFile("./optional_"+os.Args[2],result)
	StandOutput("已经添加:"+os.Args[3]+"到自选队列"+os.Args[2])
}

type Del struct {}
//删除自选
func (this *Del) Handle(){
	if len(os.Args)<5 { //1是 del 、2是l1\l2\l3.... 3是股票名称 4是end字符
		StandOutput("参数不够")
		return
	}
	if strings.ToLower(os.Args[4])!="end" {
		StandOutput("没结束符")
		return
	}
	result := ReadFromFile("./optional_"+os.Args[2])
	if strings.Contains(result,os.Args[3]){
		result=strings.ReplaceAll(result,os.Args[3]+",","")

	}else {
		StandOutput("不存在")
		return
	}
	ExportToFile("./optional_"+os.Args[2],result)
	StandOutput("已经把:"+os.Args[3]+"从自选队列"+os.Args[2]+"删除")
}





type Result struct {
	rows []*Row
}

/**
 * 拼接输出格式
 */
func (this *Result) GetResult() string{
	var buffer = strings.Builder{}

	buffer.WriteString(`<?xml version="1.0"?><items>`)
	for i, r := range this.rows {
		buffer.WriteString(`<item arg="`)
		buffer.WriteString(r.arg)
		buffer.WriteString(`" uid="`)
		buffer.WriteString(strconv.Itoa(i))
		buffer.WriteString(`" valid="yes">`)
		buffer.WriteString("<title>")
		buffer.WriteString(r.title)
		buffer.WriteString("</title>")
		buffer.WriteString("<subtitle>")
		buffer.WriteString(r.subtitle)
		buffer.WriteString(`</subtitle>`)
		//切换颜色
		if r.icon>0 {
			buffer.WriteString(`<icon>icon-red.png</icon>`)
		}else if r.icon <0{
			buffer.WriteString(`<icon>icon-green.png</icon>`)
		}else {
			buffer.WriteString(`<icon>icon.png</icon>`)
		}
		buffer.WriteString(`
		<subtitle mod="alt">添加到自选，不过需要在code 前添加l1、l2等，然后输入end</subtitle>
		<subtitle mod="cmd">打开此股票的详情页面</subtitle>
		<subtitle mod="ctrl">删除自选，不过需要在code 前添加l1、l2等，然后输入end</subtitle>`)
		buffer.WriteString(`</item>`)

	}
	buffer.WriteString(`</items>`)
	return buffer.String()
}


type Row struct {
	title string
	subtitle string
	icon int
	arg string
}


func StandOutput(result string){
	str:=`<?xml version="1.0"?>
<items>
	<item arg="{result}" uid="26083CF2-6C95-480B-9443-67090DFE7696" valid="yes">
		<title>{query}</title>
		<subtitle>{result}</subtitle>
		<icon>icon.png</icon>
		<subtitle mod="alt">拷贝到剪贴板</subtitle>
		<subtitle mod="cmd">插入到当前应用</subtitle>
	</item>`
	all := strings.ReplaceAll(str, "{result}", result)
	fmt.Println(all)
}





