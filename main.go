package main

import (
	"fmt"
	"github.com/d-kolpakov/ct/cmd"
	"log"
	"runtime"
	"sync"
)

func main()  {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	err := runtime.StartTrace()

	if err != nil {
		log.Println(err)
	}

	c, err := cmd.GoThrow(wg)

	if err != nil {
		trace := runtime.ReadTrace()
		log.Fatalln(err, string(trace))
	}

	wg.Wait()

	wg.Add(1)

	err = c.Commit(wg)

	if err != nil {
		trace := runtime.ReadTrace()
		log.Fatalln(err, string(trace))
	}

	wg.Wait()
	fmt.Println("done")
}