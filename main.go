package main

import (
	"errors"
	"fmt"
)


	var bingDomains = map[string]string{
	"com": "",
	"uk": "&cc=GB",
	"us": "&cc=US",
	"tr": "&cc=TR",
	"tw": "&cc=TW",
	"ch": "&cc=CH",
	"se": "&cc=SE",
	"es": "&cc=ES",
	"za": "&cc=ZA",
	"sa": "&cc=SA",
	"ru": "&cc=RU",
	"ph": "&cc=PH",
	"pt": "&cc=PT",
	"pl": "&cc=PL",
	"cn": "&cc=CN",
	"no": "&cc=NO",
	"nz": "&cc=NZ",
	"nl": "&cc=NL",
	"mx": "&cc=MX",
	"my": "&cc=MY",
	"kr": "&cc=KR",
	"jp": "&cc=JP",
	"it": "&cc=IT",
	"id": "&cc=ID",
	"in": "&cc=IN",
	"hk": "&cc=HK",
	"de": "&cc=DE",
	"fr": "&cc=FR",
	"fi": "&cc=FI",
	"dk": "&cc=DK",
	"cl": "&cc=CL",
	"ca": "&cc=CA",
	"br": "&cc=BR",
	"be": "&cc=BE",
	"at": "&cc=AT",
	"au": "&cc=AU",
	"ar": "&cc=AR",
}

type searchResult struct{
	ResultRank int;
	ResultURL string;
	ResultTitle string;
	ResultDesc string;
}
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent()string{

}
func buildUrl()string{

}

func scrapeClientRequest(){

}
func bingResultParser(){

}

func bingScrape(searchTerm string,country string,pageCount int,count int,backoff int)([]searchResult,error){
results:=[]searchResult{
	bingpages,err:=buildUrl(searchTerm,country,pageCount,count)
	if err!=nil{
		return nil,err
	}

	for _,page:= range bingpages{
		rank:=len(results)
		res,err:=scrapeClientRequest(page)
		if err!=nil{
			return nil,err
		}
		data,err:=bingResultParser(res,rank)
		if err!=nil{
			return nil,err
		}
		for _,result:= range data{
			results=append(results,result)
		}
		// Give delay between requests
		time.Sleep(time.Duration(backoff)*time.second)
	 }
}
}

func main(){
res,err:=bingScrape("Mayank Kumar", "com", 2, 30,30)
if err==nil{
	for _,res := range res{
		fmt.Println(res)
	}
}else{
	fmt.Println(err)
}
}
 