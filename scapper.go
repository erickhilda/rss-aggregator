package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/erickhilda/rssagg/internal/db"
)

func startScapper(database *db.Queries, concurency int, timeBetWeenRequest time.Duration) {
	log.Printf("Starting scapper with %d workers", concurency)
	ticker := time.NewTicker(timeBetWeenRequest)
	for ; ; <-ticker.C {
		feeds, err := database.GetNextFeedsToFetch(context.Background(), int32(concurency))
		if err != nil {
			log.Printf("Error getting feeds to fetch: %s", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(database, feed, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(database *db.Queries, feed db.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := database.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %s", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s: %s", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		log.Printf("Creating item %s", item.Title)
	}
	log.Printf("Done fetching %s", feed.Url)
	// log.Printf("Fetching %s", feed.Url)
	// rssFeed, err := urlToFeed(feed.Url)
	// if err != nil {
	// 	log.Printf("Error fetching feed %s: %s", feed.Url, err)
	// 	return
	// }
	// for _, item := range rssFeed.Channel.Items {
	// 	_, err := database.CreateItem(context.Background(), db.CreateItemParams{
	// 		Title:       item.Title,
	// 		Link:        item.Link,
	// 		Description: item.Description,
	// 		PubDate:     item.PubDate,
	// 		FeedID:      feed.ID,
	// 	})
	// 	if err != nil {
	// 		log.Printf("Error creating item: %s", err)
	// 	}
	// }
	// _, err = database.UpdateFeed(context.Background(), db.UpdateFeedParams{
	// 	ID: feed.ID,
	// 	LastScrapedAt: time.Now(),
	// })
	// if err != nil {
	// 	log.Printf("Error updating feed: %s", err)
	// }
}