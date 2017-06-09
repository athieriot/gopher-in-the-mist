package main

import (
	"fmt"
	"os"
	"monitise.com/gopherinthemist/prep2/api"
	"time"
	"sync"
)

type Pair struct {
	movie api.Movie
	actor api.Cast
}

func main() {
	result, err := api.Search("Wonder woman")
	if err != nil {
		fmt.Println("Oops. Something happened:", err.Error())
		os.Exit(1)
	}

	var wg sync.WaitGroup
	pairs := make(chan Pair)

	fmt.Println("Movies found:", result.TotalResults)
	result.IterateMovies(func (movie api.Movie) {
		wg.Add(1)
		go func() {
			credits, err := api.Credits(movie.Id)
			if err != nil {
				fmt.Println("No Credits:", err.Error())
			}

			credits.IterateCast(func (person api.Cast) {
				time.Sleep(10 * time.Millisecond)
				pairs <- Pair{movie, person}
			})
			wg.Done()
		}()
	})
	go func() {
		wg.Wait()
		close(pairs)
	}()

	for pair := range pairs {
		fmt.Println(pair.movie.Title, pair.actor.Name, "Character:", pair.actor.Character)
	}
}