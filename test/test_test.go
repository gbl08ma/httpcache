package test_test

import (
	"testing"

	"github.com/twpayne/httpcache"
	"github.com/twpayne/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
