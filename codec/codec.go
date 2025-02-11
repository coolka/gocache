package codec

import (
	"context"
	"time"

	"github.com/eko/gocache/v2/store"
)

// Stats allows to returns some statistics of codec usage
type Stats struct {
	Hits              int
	Miss              int
	SetSuccess        int
	SetError          int
	DeleteSuccess     int
	DeleteError       int
	InvalidateSuccess int
	InvalidateError   int
	ClearSuccess      int
	ClearError        int
}

// Codec represents an instance of a cache store
type Codec struct {
	store store.StoreInterface
	stats *Stats
}

// New return a new codec instance
func New(store store.StoreInterface) *Codec {
	return &Codec{
		store: store,
		stats: &Stats{},
	}
}

// Get allows to retrieve the value from a given key identifier
func (c *Codec) Get(ctx context.Context, key interface{}) (interface{}, error) {
	val, err := c.store.Get(ctx, key)

	if err == nil {
		c.stats.Hits++
	} else {
		c.stats.Miss++
	}

	return val, err
}

// GetWithTTL allows to retrieve the value from a given key identifier and its corresponding TTL
func (c *Codec) GetWithTTL(ctx context.Context, key interface{}) (interface{}, time.Duration, error) {
	val, ttl, err := c.store.GetWithTTL(ctx, key)

	if err == nil {
		c.stats.Hits++
	} else {
		c.stats.Miss++
	}

	return val, ttl, err
}

// Set allows to set a value for a given key identifier and also allows to specify
// an expiration time
func (c *Codec) Set(ctx context.Context, key interface{}, value interface{}, options *store.Options) error {
	err := c.store.Set(ctx, key, value, options)

	if err == nil {
		c.stats.SetSuccess++
	} else {
		c.stats.SetError++
	}

	return err
}

// Delete allows to remove a value for a given key identifier
func (c *Codec) Delete(ctx context.Context, key interface{}) error {
	err := c.store.Delete(ctx, key)

	if err == nil {
		c.stats.DeleteSuccess++
	} else {
		c.stats.DeleteError++
	}

	return err
}

// Invalidate invalidates some cach items from given options
func (c *Codec) Invalidate(ctx context.Context, options store.InvalidateOptions) error {
	err := c.store.Invalidate(ctx, options)

	if err == nil {
		c.stats.InvalidateSuccess++
	} else {
		c.stats.InvalidateError++
	}

	return err
}

// Clear resets all codec store data
func (c *Codec) Clear(ctx context.Context) error {
	err := c.store.Clear(ctx)

	if err == nil {
		c.stats.ClearSuccess++
	} else {
		c.stats.ClearError++
	}

	return err
}

// GetStore returns the store associated to this codec
func (c *Codec) GetStore() store.StoreInterface {
	return c.store
}

// GetStats returns some statistics about the current codec
func (c *Codec) GetStats() *Stats {
	return c.stats
}
