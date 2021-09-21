package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// TODO: Burada ki mapi thread safe hale getirebilirsiniz.
// 102-concurrency egitimindeki mutex orneklerine bakabilirsiniz.
// Ref: https://pmihaylov.com/thread-safety-concerns-go/
// Ref: https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
var cache = map[string]Cache{}

type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}

func EvictCacheHandler(c *fiber.Ctx) error {
	// TODO: [DELETE] /cache/:key/* pathine istek atildiginda memorydeki cachei temizleyen handleri implement edebilirsiniz.
	// TODO: implement me!
	return nil
}

func NewCacheProxy(key string, ttl time.Duration) CacheProxy {
	return CacheProxy{
		key: key,
		ttl: ttl,
	}
}

func (p CacheProxy) Accept(key string) bool {
	return p.key == key
}

func (p CacheProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()
	key := c.Params("key")

	if r, ok := cache[path]; ok && r.ttl.After(time.Now()) {
		c.Response().SetBody(r.body)
		c.Response().Header.Add("cache-control", fmt.Sprintf("max-age:%d", p.ttl/time.Second))

		return nil
	}

	// https://mocki.io/v1/d4d63bce-1809-4250-91ac-0470aa392ca5
	url := "https://mocki.io/" + strings.TrimPrefix(path, "/"+key+"/")

	fmt.Printf("http request redirecting to %s \n", url)

	if err := proxy.Do(c, url); err != nil {
		return err
	}

	ch := Cache{
		ttl:  time.Now().Add(p.ttl),
		body: c.Response().Body(),
	}

	cache[path] = ch

	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}
