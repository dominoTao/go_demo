package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
)

func TestHttp(t *testing.T) {
	//get, err := http.Get("https://www.iana.org/domains/reserved")
	//if err != nil {
	//	return
	//}
	//defer get.Body.Close()
	//io.Copy(os.Stdout, get.Body)

	//post, err := http.Post("http://example.com/upload", "image/jpeg", nil)
	//if err != nil {
	//	return
	//}
	//if post.StatusCode != http.StatusOK {
	//	return
	//}
	//io.Copy(os.Stdout, post.Body)

	//http.PostForm("http://examle.com/posts", url.Values{"title":{"article title"}, "content":{"article body"}})

	//http.Head("http://example.com/")

	//req, _ := http.NewRequest("GET", "http://example.com", nil)
	//req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	//client := &http.Client{}
	//do, _ := client.Do(req)
	//fmt.Println(do.Header)

	//client := &http.Client{
	//	CheckRedirect: redirectPolicyFunc,
	//}
	//
	////resp, err := client.Get("http://example.com")
	//req, _ := http.NewRequest("GET", "http://example.com", nil)
	//req.Header.Add("User-Agent", "Our Custom User-Agent")
	//req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
	//resp, _ := client.Do(req)
	//fmt.Println(resp.Body)

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://example.com")
	fmt.Println(resp.Body)
}


func TestNet(t *testing.T){
	//ip := net.ParseIP("39.156.69.79")
	//for index, value := range ip {
	//	fmt.Println(index, value)
	//}

	// 根据域名找ip
	//net.ResolveIPAddr()
	host, err := net.LookupHost("www.google.com")
	checkError2(err)
	for index,value := range host {
		fmt.Println(index, value)
	}
}
func checkError2(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
