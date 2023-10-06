package persist

import (
	"fmt"
	"good-spider/model"
	"good-spider/persist/sqlclient"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Got item #%d : %v", itemCount, item)
			err := save(item)
			if err != nil {
				log.Print(err)
			}
			itemCount++
		}
	}()
	return out
}

func save(items ...interface{}) (err error) {
	var movieInfo []model.MovieInfo
	for _, item := range items {
		if val, ok := item.(model.MovieInfo); ok {
			movieInfo = append(movieInfo, val)
		} else {
			err = fmt.Errorf("interface is not MovieInfo")
		}
	}
	err = sqlclient.MultiInsert(movieInfo...)
	return err
}
