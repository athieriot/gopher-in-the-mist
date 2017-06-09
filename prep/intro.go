package main

import (
	"fmt"
	"os"
	"time"
	"monitise.com/gopherinthemist/prep/api"
	"sync"
)

func ToTime(date string) time.Time {
	time, _ := time.Parse("2006-01-02", date)
	return time
}


func main() {
	result, err := api.Search("Wonder Woman")
	if err != nil {
		fmt.Println("An error happened while fetching movies", err.Error())
		os.Exit(1)
	}

	var wg sync.WaitGroup
	actors := make(chan api.Cast)

	result.Iterate(func (movie api.Movie) {
		date := ToTime(movie.ReleaseDate)
		fmt.Println(movie.Title, "Note:", movie.VoteAverage, "Released:", date.Year())
		wg.Add(1)

		go func() {
			credits, err := api.Credits(movie.Id)
			if err != nil {
				fmt.Println("Dohh !", err.Error())
			}
			credits.Iterate(func (person api.Cast) {
				actors <- person
			})
			wg.Done()
		}()
	})
	go func() {
		wg.Wait()
		close(actors)
	}()

	for actor := range actors {
		fmt.Println(actor.Character, "Actor:", actor.Name)
	}
}