package main

import "fmt"
import "net/http"
import "io/ioutil"

func main(){
	
	resp, err := http.Get("https://www.googleapis.com/language/translate/v1?source=en&target=fr&q=Hello")
	if(err!=nil){
		fmt.Println(err)
		return
	}
	
	body, err :=ioutil.ReadAll(resp.Body)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
