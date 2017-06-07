package main

import (
	"fmt"
	"os"
	"monitise.com/gopherinthemist/exercise/step5/api"
	"sync"
	"time"
)

type Pair struct {
	Movie api.Movie
	Cast api.Person
}


func main() {
	result, err := api.Search("Wonder woman")
	if err != nil {
		fmt.Println("An error occured", err.Error())
		os.Exit(1)
	}
	var wg sync.WaitGroup
	casting := make(chan Pair)
	done := make(chan bool)

	fmt.Println("There is", result.TotalResults, "movies")

	wg.Add(result.TotalResults)
	result.Iterate(func(v api.Movie) {
		go func() {
			movieCast, err := api.Cast(v.Id)

			if err == nil {
				movieCast.Iterate(func(c api.Person) {
					time.Sleep(10 * time.Millisecond)
					casting <- Pair{v, c}
				})
			}
			wg.Done()
		}()
	})

	go func() {
		wg.Wait()
		close(casting)
		done <- true
	}()

	for pair := range casting {
		fmt.Println("Movie: ", pair.Movie.Title, "Role: ", pair.Cast.Name, pair.Cast.Character)
	}

	<-done
}