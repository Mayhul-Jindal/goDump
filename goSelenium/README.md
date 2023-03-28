- selenium grid setup 

- chrome aur firefox ke parallely aur independenty hee hone chahiye

- 100 tests in one browser
    1. Extract the urls and store it in a slice
    2. Now just for url := urls {
        go getNet(url)
    }

    getNet(){
        // extracting the network logs 
        netch <- netlogs
    }   

    // dont care about the order in which 