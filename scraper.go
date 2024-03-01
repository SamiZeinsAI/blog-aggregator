package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) scraper() error {
	limit := int32(10)
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	var wg sync.WaitGroup
	for {
		select {
		case <-ticker.C:
			fmt.Println("start of scraper loop")
			feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), limit)
			if err != nil {
				return err
			}

			for _, feed := range feeds {
				wg.Add(1)
				go cfg.processFeed(&wg, feed)
			}
			wg.Wait()

		case <-quit:
			ticker.Stop()
			return nil
		}
	}
}

func (cfg *apiConfig) processFeed(wg *sync.WaitGroup, feed database.Feed) error {
	defer wg.Done()
	rssFeed, err := fetchRSSFeed(feed.Url)
	if err != nil {
		return err
	}
	for _, item := range rssFeed.Channel.Items {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, _ = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
	}
	return nil
}
