package store

import (
    "fmt"
    "github.com/go-redis/redis"
    "time"
)

type StorageService struct {
    redisClient *redis.Client
}

var storService = &StorageService{}

const CacheDuration = 6 * time.Hour

func InitializeStorageService() *StorageService {
	storService.redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := storService.redisClient.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis - Error: %v", err))
	}

	fmt.Println("\nRedis started successfully")
	return storService
}
func SaveUrlMapping(shortUrl string, originalUrl string, userId string){
	err := storService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result

}