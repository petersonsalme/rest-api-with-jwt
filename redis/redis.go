package redis

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/petersonsalme/rest-api-with-jwt/model"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

// Connect should connect to Redis
func Connect() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")

	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatal(err.Error())
	}

	return client
}

// CreateAuth createAuth
func CreateAuth(userid uint64, token *model.Token) error {
	at := time.Unix(token.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(ctx, token.AccessUUID, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(ctx, token.RefreshUUID, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// FetchAuth FetchAuth
func FetchAuth(authD *model.AccessDetails) (uint64, error) {
	userid, err := client.Get(ctx, authD.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

// DeleteAuth DeleteAuth
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := client.Del(ctx, givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
