package main

import (
	"time"

	"github.com/erickhilda/rssagg/internal/db"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser db.User) User {
	return User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedToFeed(dbFeed db.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func databaseFeedsToFeeds(dbFeeds []db.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uint64    `json:"id"`
	FeedID    int32     `json:"feed_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow db.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		FeedID:    dbFeedFollow.FeedID,
		UserID:    dbFeedFollow.UserID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollow []db.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, feedFollow := range dbFeedFollow {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}

type Post struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      int32     `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databasePostToPost(dbPost db.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func databasePostsToPosts(dbPosts []db.Post) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
