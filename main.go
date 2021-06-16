/**
	Site: https://tryhackme.com/room/webfundamentals
	TODO:
		- Write the code to get the GET request flag
		- Write the code to get the POST request flag
		- Write the code to get the flag from cookie
		- Write the code to set a cookie and get the flag

 **/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	var ctf *int
	ctf = flag.Int("ctf", 0, "1-4 which task you want to run")
	flag.Parse()
	if *ctf == 0 {
		flag.PrintDefaults()
	}

	switch(*ctf){
		case 1:
			getCTF()
			break
		case 2:
			postCTF()
			break
		case 3:
			getCookieCTF()
			break
		case 4:
			sendCookieCTF()
			break

	}
}

// Task 1 \ getting flag by just sending GET request to /ctf/get
func getCTF(){
	var apiUrl string = "http://10.10.78.223:8081/ctf/get"
	res, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	}
}

// Task 2 \ getting flag by sending POST request to /ctf/post
func postCTF(){
	var apiUrl string = "http://10.10.78.223:8081/ctf/post"
	// creating client
	var client http.Client
	// setting data
	var data = []byte(`flag_please`)
	// creating request
	res, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	// making request
	resp, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// getting the flag
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
}

// Task 3 \ getting cookie
func getCookieCTF(){
	// setting target
	var apiUrl string = "http://10.10.78.223:8081/ctf/getcookie"
	// creating request
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	// creating client
	var client http.Client
	// making request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println(resp.Cookies())
	}

}

// Task 4 \ setting cookie to get the flag
func sendCookieCTF(){
	// setting target
	var apiUrl string = "http://10.10.78.223:8081/ctf/sendcookie"
	// creating new request
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	// creating client
	var client http.Client
	// setting headers
	req.Header.Set("Cookie", "flagpls=flagpls")
	// making request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// response
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", body)
	}
}