package main

import (
	"context"
	"database/sql"
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
		description := sql.NullString{}
		if item.Description != "" {
			description.Valid = true
			description.String = item.Description
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %s: %s", item.PubDate, err)
			continue
		}

		_, errCreatePost := database.CreatePost(context.Background(), db.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: pubDate,
			FeedID:      int32(feed.ID),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if errCreatePost != nil {
			log.Printf("Error creating post: %s", errCreatePost)
			continue
		}
	}
	log.Printf("Done fetching %s", feed.Url)
}
