package cache

import (
	"context"
	"time"

	fiberRedis "github.com/gofiber/storage/redis"
	configs "server.com/pkg/config"
	"server.com/pkg/types"
	"server.com/pkg/utils/logger"
	redislock "server.com/platform/cache/redisLock"
)

var globalStore *GlobalStore

type GlobalStore struct {
	*fiberRedis.Storage
}

func GetStore() *GlobalStore {
	if globalStore == nil {
		globalStore = &GlobalStore{
			fiberRedis.New(configs.FiberRedisConfig()),
		}
	}

	return globalStore
}

// Pattern here => https://redis.io/commands/keys/
func (store GlobalStore) DeleteByPattern(pattern string) error {
	redis := store.Conn()
	ctx := context.Background()
	cmd := redis.Keys(ctx, pattern)
	err := cmd.Err()

	if targets := cmd.Val(); cmd.Err() == nil && len(targets) > 0 {
		for _, key := range targets {
			err = store.Delete(key)
		}
	}

	return err
}

// Run func once in one redis world, non concurrency and avoid conflicts
func RunOnce(key string, fn func()) error {
	lockKey := PrefixLock + key
	locker := redislock.New(GetStore().Conn())
	ctx := context.Background()

	// Try to obtain lock.
	lock, err := locker.Obtain(ctx, lockKey, 100*time.Millisecond, nil)
	if err == redislock.ErrNotObtained {
		logger.Verbose("Could not obtain lock!, [KEY]:", lockKey)
	} else if err != nil {
		return err
	}

	defer lock.Release(ctx)
	fn()

	return err
}

// Get data from redis or function
func Wrap[T any](key string, ttl time.Duration, fn func() (T, error)) T {
	var (
		value T
		bytes types.JSON
		err   error
	)

	store := GetStore()
	bytes, err = store.Get(key)

	if err != nil || bytes == nil {
		value, err = fn()
		if err == nil {
			bytes.Set(value)
			store.Set(key, bytes, ttl)
		}

		return value
	}

	bytes.Out(&value)
	return value
}

func WrapWithError[T any](key string, ttl time.Duration, fn func() (T, error)) (T, error) {
	var (
		value T
		bytes types.JSON
		err   error
	)

	store := GetStore()
	bytes, err = store.Get(key)

	if err != nil || bytes == nil {
		value, err = fn()
		if err == nil {
			bytes.Set(value)
			store.Set(key, bytes, ttl)
		}

		return value, nil
	}

	bytes.Out(&value)
	return value, err
}

// clear cache
func ClearCache(key string) error {
	store := GetStore()
	err := store.Delete(key)
	return err
}

// clear cache by key**
func ClearCacheByPattern(pattern string) error {
	return GetStore().DeleteByPattern(pattern)
}

// Update cache
func Patch(key string, value interface{}, age time.Duration) error {
	var (
		bytes types.JSON
	)
	store := GetStore()
	err := bytes.Set(value)
	if err != nil {
		return err
	}
	err = store.Set(key, bytes, age)
	if err != nil {
		return err
	}
	return nil
}

func FlushCache() error {
	store := GetStore()
	err := store.Reset()
	return err
}
