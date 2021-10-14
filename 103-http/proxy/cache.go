package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// DONE: Buradaki mapi thread safe hale getirebilirsiniz.
// 102-concurrency egitimindeki mutex orneklerine bakabilirsiniz.
// Ref: https://pmihaylov.com/thread-safety-concerns-go/
// Ref: https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c

var cache = &CacheStore{v: map[string]Cache{}}

type CacheStore struct {
	sync.Mutex
	v map[string]Cache
}

type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}

func NewCacheProxy(key string, ttl time.Duration) CacheProxy {
	return CacheProxy{
		key: key,
		ttl: ttl,
	}
}

func EvictCacheHandler(c *fiber.Ctx) error {
	// DONE: [DELETE] /cache/:key/* pathine istek atildiginda memorydeki cachei temizleyen handleri implement edebilirsiniz.
	// DONE: implement me!

	key := strings.TrimPrefix(c.Path(), "/cache")

	if _, ok := cache.v[key]; !ok {
		return fiber.ErrNotFound
	}

	cache.Delete(key)

	c.Response().SetStatusCode(fiber.StatusNoContent)
	return nil
}

func (p CacheProxy) Accept(key string) bool {
	return p.key == key
}

func (p CacheProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()
	key := c.Params("key")

	if r, ok := cache.v[path]; ok && r.ttl.After(time.Now()) {
		c.Response().SetBody(r.body)
		c.Response().Header.Add(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		c.Response().Header.Add(fiber.HeaderCacheControl, fmt.Sprintf("max-age:%d", p.ttl/time.Second))

		return nil
	}

	// https://mocki.io/v1/d4d63bce-1809-4250-91ac-0470aa392ca5
	url := "https://mocki.io" + strings.TrimPrefix(path, "/"+key)

	fmt.Printf("http request redirecting to %s \n", url)

	if err := proxy.Do(c, url); err != nil {
		return err
	}

	respStatusCode := c.Response().StatusCode()

	if respStatusCode != fiber.StatusOK {
		return fiber.NewError(respStatusCode, "Check your request")
	}

	ch := Cache{
		ttl:  time.Now().Add(p.ttl),
		body: c.Response().Body(),
	}

	//cache[path] = ch

	// thread safe version
	cache.Set(path, ch)

	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}

func (c *CacheStore) Set(key string, cache Cache) {
	c.Lock()
	c.v[key] = cache
	c.Unlock()
}

func (c *CacheStore) Delete(key string) {
	c.Lock()
	delete(c.v, key)
	c.Unlock()
}
