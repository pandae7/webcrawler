package crawler

import (
	"context"
	"log"
	"sync"

	"github.com/go-redis/redis"
)

// Crawler Structure will be defined here

type Crawler struct {
	config      CrawlerConfig
	redisClient *redis.Client
	wg          sync.WaitGroup
	logger      *log.Logger
}

type CrawlerConfig struct {
}

func NewCrawler() (*Crawler, error) {
	crawler := &Crawler{}
	return crawler, nil
}

// Based on Crawler Config and Object
// Handles Initialisation of Crawl processes
// encapsulates logic of handling HTML, XML or reading robots.txt
// Calls functions further down to push URLs to queue
func (c *Crawler) Crawl(url string) {
	// Take context
	// initiate workers based on the URL queue

	// Where are you storing URLs?
	// use sync.Map for now,
	// Can migrate to redis later
	// Use redis for storing product URLs though

	// Print out all the URLs in the working example first-

}

func (c *Crawler) worker(ctx context.Context) {
	defer c.wg.Done()

	for url := range c.urlQueue {
		if _, visited := c.visited.LoadOrStore(url, true); visited {
			continue
		}

		if err := c.processURL(ctx, url); err != nil {
			c.logger.Printf("Error processing %s: %v", url, err)
		}
	}
}

func (c *Crawler) processURL(ctx context.Context, url string) {
	// Process URL in the following ways
	// call parse function,
	// handle content variations
	// find if the page has more URLs and queue them
}
