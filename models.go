package main

import (
	"time"

	"github.com/SamiZeinsAI/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

type Feed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

func databaseUsersFeedsToUsersFeeds(usersFeeds database.FeedsUser) FeedsUser {
	return FeedsUser{
		ID:        usersFeeds.ID,
		FeedID:    usersFeeds.FeedID,
		UserID:    usersFeeds.UserID,
		CreatedAt: usersFeeds.CreatedAt,
		UpdatedAt: usersFeeds.UpdatedAt,
	}
}

type FeedsUser struct {
	ID        uuid.UUID
	FeedID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
