package processData

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type AutoGenerated struct {
	Path    string `json:"path"`
	Ref     string `json:"ref"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
func ProcessKc()  {
	baseUrl:= "https://www.kancloud.cn/liupengjie/go/570005"
	// Create client
	client := &http.Client{}
	req,err:= http.NewRequest("GET",baseUrl,nil);
	if err!= nil{
		panic("xxx")
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	//strings.NewReader(html)
	resp,err:= client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	var  lists []string


	ss:=strings.NewReader(string(respBody))
	re,err:= goquery.NewDocumentFromReader(ss)
	if err!=nil{
		log.Fatalln(err)
	}

	re.Find(".catalog > ul>li>a").Each(func(i int, selection *goquery.Selection) {
		x,_:=selection.Attr("href")
		fmt.Println(x)
		lists=append(lists, x)
	})
	sumFIles,err:=os.Create("./GOlang小书.md")
	defer sumFIles.Close()
	for i,v:= range lists {
		baseUrl = baseUrl[:strings.LastIndex(baseUrl,"/") +1] +v;
		log.Println(baseUrl)
		client := &http.Client{}
		req,err:= http.NewRequest("GET",baseUrl,nil);
		if err!= nil{
			panic("xxx")
		}
		req.Header.Set("Accept","application/json, text/javascript, */*; q=0.01")
		resp,err:= client.Do(req)
		respBody, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		//ss:=strings.NewReader(string(respBody))
		//re,err:= goquery.NewDocumentFromReader(ss)

		if err!=nil{
			log.Fatalln(err)
		}
		x:=AutoGenerated{}

		json.Unmarshal(respBody,&x)
		sumFIles.WriteString("# -----------第" +strconv.Itoa(i+1)+"小节-------------- \n")
		sumFIles.WriteString(x.Content)
	}

}