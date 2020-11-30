package loglevel

import (
	"fmt"
	golog "log"
)

func main() {
	fmt.Println("Start")
	golog.Println("normal log")
	log := logger.GetLogger()

}
