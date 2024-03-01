package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
)

func (cfg *apiConfig) scraper() error {
	limit := int32(10)
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	var wg sync.WaitGroup
	for {
		select {
		case <-ticker.C:
			feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), limit)
			if err != nil {
				return err
			}

			for _, feed := range feeds {
				wg.Add(1)
				go func(feed database.Feed) {
					wg.Done()
					processFeed(feed)
				}(feed)
			}
			wg.Wait()

		case <-quit:
			ticker.Stop()
			return nil
		}
	}
}

func processFeed(feed database.Feed) error {
	rssFeed, err := fetchRSSFeed(feed.Url)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", rssFeed.Channel.Title)
	for _, post := range rssFeed.Channel.Items {
		fmt.Printf("%s\n", post.Title)
	}
	return nil
}
