package main

import (
    "fmt"
    "bufio"
	"log"
	"os"
    "net/http"
)

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	var sitelist string
    fmt.Print("Input your site list :")
	fmt.Scanln(&sitelist)

	file, err := os.Open(sitelist)

	if err != nil{
		log.Fatalf("Failed to open")
	}
	scanner := bufio.NewScanner(file) 
	scanner.Split(bufio.ScanLines) 
    var text []string 
  
    for scanner.Scan() { 
        text = append(text, scanner.Text()) 
    } 
	file.Close() 
  
    // and then a loop iterates through  
    // and prints each of the slice values. 
    for _, each_ln := range text { 
		resp, err := http.Get("https://"+each_ln+"/.git")
		if err != nil {
			log.Fatal(err)
		}
	

    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println(string(colorGreen), each_ln, ">>>" ,"HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
    } else {
        fmt.Println(string(colorRed), each_ln,">>> Not Found")
    }
    } 

	
}