package storage_test

import (
	"context"
	"net"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/tamarakaufler/grpc-char-vs-rune/internal/storage"
)

// TestRedisStoragetests setting and retrieval of cached data.
func TestRedisStorage(t *testing.T) {
	redisClient, redisCleanupF := setup(t)
	defer t.Cleanup(redisCleanupF)
	defer redisClient.Close()

	t.Run("StoreRetrieveRuneToChar", func(t *testing.T) {
		ctx := context.Background()
		s := "Hello 日本"
		rs := []byte(s)
		d := 200 * time.Millisecond

		r := &storage.Redis{
			Client: redisClient,
			TTL:    time.Duration(d),
		}
		err := r.StoreRuneToChar(ctx, rs, s)
		require.NoError(t, err)

		v, err := r.GetRuneToChar(ctx, rs)
		require.NoError(t, err)
		require.Equal(t, s, v)

		tick := time.After(d)
		select {
		case <-tick:
			v, err := r.GetRuneToChar(ctx, rs)
			require.EqualError(t, err, "redis: nil")
			require.Equal(t, "", v)
		}

		err = r.StoreRuneToChar(ctx, rs, s)
		require.NoError(t, err)

		v, err = r.GetRuneToChar(ctx, rs)
		require.NoError(t, err)
		require.Equal(t, s, v)
	})

	t.Run("StoreRetrieveCharToRune", func(t *testing.T) {
		ctx := context.Background()
		s := "Hello 日本"
		expected := runeToUint32(s)

		// TTL set through expire cannot be less than 1 second
		d := 2 * time.Second

		r := &storage.Redis{
			Logger: &logrus.Entry{},
			Client: redisClient,
			TTL:    time.Duration(d),
		}
		err := r.StoreCharToRune(ctx, s, expected)
		require.NoError(t, err)

		v, err := r.GetCharToRune(ctx, s)
		require.NoError(t, err)
		require.EqualValues(t, expected, v)

		select {
		case <-time.After(d):
			v, err := r.GetCharToRune(ctx, s)
			require.EqualError(t, err, "redis: nil")
			require.Equal(t, []uint32(nil), v)
		}

		err = r.StoreCharToRune(ctx, s, expected)
		require.NoError(t, err)

		v, err = r.GetCharToRune(ctx, s)
		require.NoError(t, err)
		require.EqualValues(t, expected, v)
	})
}

// setup brings up redis container
func setup(t *testing.T) (*redis.Client, func()) {
	t.Helper()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Failed to start Dockertest: %+v", err)
	}
	resource, err := pool.Run("redis", "5-alpine", nil)
	if err != nil {
		t.Fatalf("Failed to start redis: %+v", err)
	}
	addr := net.JoinHostPort("localhost", resource.GetPort("6379/tcp"))

	ctx := context.Background()
	// waiting for the container to be ready
	err = pool.Retry(func() error {
		var e error
		client := redis.NewClient(&redis.Options{
			Addr:       addr,
			MaxRetries: 10,
		})
		defer client.Close()

		_, e = client.Ping(ctx).Result()
		return e
	})
	if err != nil {
		t.Fatalf("Failed to ping Redis: %+v", err)
	}

	cleanupRedisF := func() {
		_ = pool.Purge(resource)
	}
	client := redis.NewClient(&redis.Options{Addr: addr})

	return client, cleanupRedisF
}

// runeToUint32 creates a list of uint32/runes corresponding to the provided string.
func runeToUint32(s string) []uint32 {
	uis := make([]uint32, utf8.RuneCountInString(s))
	for i, r := range []rune(s) {
		uis[i] = uint32(r)
	}
	return uis
}
