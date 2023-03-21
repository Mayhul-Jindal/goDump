package main




// yeh pattern 1 hain but meine vese thodha toh chudap likha hain yaha but samajh toh aa gaya tha jab likha tha
// yeh wait group wali cheez tabh kaam ari jab tum for loop mein chalare verna toh buffered use karke hack kar sakte
// import (
// 	"fmt"
// 	"time"
// )

// type testObject struct{
// 	Id int
// 	Name string
// 	Recommendations []string
// 	Likes int
// }

// var database = []testObject{
// 	{
// 		Id: 1,
// 		Name: "mayhul jindal",
// 		Recommendations: []string{"anthingg", "kunal"},
// 		Likes: 20,
// 	},
// 	{
// 		Id: 2,
// 		Name: "rohan verma",
// 		Recommendations: []string{"coco melon", "code with harry"},
// 		Likes: 1,
// 	},
// 	{
// 		Id: 3,
// 		Name: "prerit rana",
// 		Recommendations: []string{"andrej katharpy", "ml ka chodha hun mein"},
// 		Likes: 200,
// 	},
// }

// func main(){
// 	now := time.Now()

// 	chanP := make(chan string, 2)
// 	chanR := make(chan []string, 3)
// 	chanL := make(chan int, 4)

// 	go fetchUserName(1, chanP)
// 	go fetchUserRecommendationData(1, chanR)
// 	go fetchUserLikesData(1, chanL)

//     var p string
//     var r []string
//     var l int

// 	p = <-chanP
// 	r = <-chanR
// 	l = <-chanL

// 	fmt.Println(p)
// 	fmt.Println(r)
// 	fmt.Println(l)

// 	fmt.Println(time.Since(now))

// }

// func fetchUserName(Id int, chanP chan string){
// 	time.Sleep(1000 * time.Millisecond)

// 	for _, person := range database{
// 		if person.Id == Id{
// 			chanP <- person.Name
// 			return
// 		}
// 	}

// 	chanP <- ""
// }

// func fetchUserRecommendationData(Id int, chanR chan []string){ // weakest link
// 	time.Sleep(100 * time.Millisecond)

// 	for _, person := range database{
// 		if person.Id == Id{
// 			chanR <- person.Recommendations
// 			return
// 		}
// 	}

// 	chanR <- []string{}
// }

// func fetchUserLikesData(Id int, chanL chan int){
// 	time.Sleep(100 * time.Millisecond)

// 	for _, person := range database{
// 		if person.Id == Id{
// 			chanL <- person.Likes
// 			return
// 		}
// 	}

// 	chanL <- 0
// }

