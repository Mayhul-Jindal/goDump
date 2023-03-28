//! make a cli out of this + golang testing benchmarking dekhna hain. Yeh kal taq toh ho jana chahiye

package main

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main(){
    // setting up capabilities
    caps := selenium.Capabilities{}
    chromeCaps := chrome.Capabilities{
        Path: "",
        Args: []string{
            "--headless",
            "--disable-gpu",
            "--no-sandbox",
            "--disable-dev-shm-usage",
        },
    }
    caps.AddChrome(chromeCaps)

    // connect to our docker
    wd, err := selenium.NewRemote(caps, "http://localhost:4444");
    if err != nil{
        fmt.Fprintf(os.Stderr, "Error occured in NewRemote line: %v", err)
    }
    defer wd.Quit()

    // Navigate to a URL
    err = wd.Get("https://www.google.com")
    if err != nil {
        panic(err)
    }
    
    // Grab the title of the results page and print it
    title, err := wd.Title()
    if err != nil {
        panic(err)
    }

    println(title)
}
