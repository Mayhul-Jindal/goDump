package main

import "time"

//! using context wala example, yeh part voh stream se dekhulunga
func fetchThirdPartyStuffWhichCanBeSlow()(int, error){
	time.Sleep(time.Millisecond*500)
	return 666, nil
}
func main(){
	// Pattern1()
	// Pattern2()
	fetchThirdPartyStuffWhichCanBeSlow()
}