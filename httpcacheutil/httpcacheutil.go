package httpcacheutil

import (
	"fmt"
	"net/url"

	"github.com/gomodule/redigo/redis"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"github.com/gregjones/httpcache/leveldbcache"
	"github.com/gregjones/httpcache/memcache"
	httpcacheredis "github.com/gregjones/httpcache/redis"
)

// NewCacheFromURL returns a new httpcache.Cache according the the scheme of
// urlStr.
func NewCacheFromURL(urlStr string) (httpcache.Cache, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	switch url.Scheme {
	case "file":
		return diskcache.New(url.Path), nil
	case "leveldb":
		return leveldbcache.New(url.Path)
	case "memcache":
		return memcache.New(url.Host), nil
	case "memory":
		return httpcache.NewMemoryCache(), nil
	case "redis":
		conn, err := redis.DialURL(urlStr)
		if err != nil {
			return nil, err
		}
		return httpcacheredis.NewWithClient(conn), nil
	default:
		return nil, fmt.Errorf("unknown scheme: %s", url.Scheme)
	}
}
