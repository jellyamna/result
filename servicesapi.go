package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (m model) RunServices() Apiresponse {
	var apiresponse Apiresponse

	//https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go

	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//url := "https://app07.canalisdev.corp.bankmandiri.co.id/canalis/migration/api/migration/run-migration/v.1"

	//url := "https://app07.canalisdev.corp.bankmandiri.co.id/canalis/bmri-core-rest/api/transaction/v.4/disburse/"

	m.infoLog.Println("mf url   :", *m.FinalData.Urlrest)
	m.infoLog.Println("mf select  :", m.selectedmf)
	url := *m.FinalData.Urlrest + m.selectedmf

	method := "POST"

	client := &http.Client{}

	//req, err := http.NewRequest(method, url, payload)
	//req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(*m.FinalData.Restbody)))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Authorization", *m.FinalData.Authorization)
	req.Header.Add("Authentication", *m.FinalData.Authentication)
	req.Header.Add("mac", *m.FinalData.Mac)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")

	m.infoLog.Println("Url hit :", url)
	res, err := client.Do(req)
	if err != nil {
		m.errorLog.Fatal("Can not hit canalis service  --> ", err.Error())
	}

	code := res.StatusCode

	//ioutil.ReadAll(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		m.errorLog.Fatal("Can not read body response  canalis service  --> ", err.Error())

	}

	bodyString := "Response Services : " + strconv.Itoa(code) + "\n" + string(body)
	m.infoLog.Println(bodyString)
	//apix := "Response Services adalah : "
	apiresponse.Response = &bodyString

	//m.infoLog.Println(bodyString)
	// m.response_service = &bodyString
	//fmt.Println(bodyString)
	defer res.Body.Close()

	return apiresponse

}
