package main 

import (
	"fmt"
	"utils/logger"
)

var log = logger.New()	

var a = 1


func main() {
	fmt.Println(a)
	fmt.Println("Hello World")
}

func init(){
	log.Formatter = new(logger.TextFormatter) 
	log.Formatter.(*logger.TextFormatter).DisableColors = true
	log.Formatter.(*logger.TextFormatter).FullTimestamp = false
	log.Level = logger.TraceLevel
	log.Trace( "Value:", a )
	log.Errorln( "Value:", a )
	a := a +1
	fmt.Println(a)
}