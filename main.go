package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS
	var data = readGoogleTrends()
	
	var err = xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("error:", err)
	}
 
	fmt.Println("Kita akan mengambil data API dari Google Trends ")
	fmt.Println("---------------------------------------------------")

	for i := range r.Channel.ItemList  {
		var rank = i+1
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.ItemList[i].Title)
		fmt.Println("Link to trend:", r.Channel.ItemList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("Link to article:", r.Channel.ItemList[i].NewsItems[0].HeadlineLink)
		fmt.Println("--------------------------")		
	}
}

func getGoogleTrends() *http.Response  {
	var resp, err = http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=ID")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return resp
}

func readGoogleTrends() []byte  {	
	var resp = getGoogleTrends()	
	var data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}