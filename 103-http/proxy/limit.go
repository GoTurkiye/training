package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// DONE: Burada ki mapi thread safe hale getirebilirsiniz.
// 102-concurrency egitimindeki mutex orneklerine bakabilirsiniz.
// Ref: https://pmihaylov.com/thread-safety-concerns-go/
// Ref: https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
var counter = map[string]*Limit{}
var mutex sync.Mutex

type Limit struct {
	count int
	ttl   time.Time
}

type LimitProxy struct {
	key   string
	limit int
	ttl   time.Duration
}

func ResetLimitHandler(c *fiber.Ctx) error {
	// TODO: [DELETE] /limit/:key/* pathine istek atildiginda limiti sifirlayan handleri implement edebilirsiniz.
	// TODO: implement me!
	return nil
}

func NewLimitProxy(key string, limit int, ttl time.Duration) LimitProxy {
	return LimitProxy{
		key:   key,
		limit: limit,
		ttl:   ttl,
	}
}

func (p LimitProxy) Accept(key string) bool {
	return p.key == key
}

func (p LimitProxy) Proxy(c *fiber.Ctx) error {
	path := c.Path()

	if r, ok := counter[path]; ok && r.count >= p.limit {
		if r.ttl.After(time.Now()) {
			c.Response().SetStatusCode(429)

			fmt.Printf("request limit reached for %s \n", path)

			return nil
		} else {
			fmt.Printf("resetting counter values \n")

			incrementCounter(path, p.ttl)
		}
	} else if !ok {
		incrementCounter(path, p.ttl)
	}

	c.SendString("Go Turkiye - 103 Http Package")

	getCounter(path).count++

	return nil
}

func incrementCounter(path string, ttl time.Duration) {
	mutex.Lock()
	counter[path] = &Limit{
		count: 0,
		ttl:   time.Now().Add(ttl),
	}
	mutex.Unlock()
}

func getCounter(path string) *Limit {
	defer mutex.Unlock()
	mutex.Lock()
	return counter[path]
}
