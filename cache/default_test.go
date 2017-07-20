package cache

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tests default cache", func() {

	var cache Cache
	cache = NewDefaultCache()

	It("should add a new item to the cache.", func() {
		cache.Set("test-key", "test-value", time.Second*10)

		found, err := cache.Exists("test-key")

		Expect(found).To(Equal(true))
		Expect(err).ToNot(HaveOccurred())
	})

	It("should get an item from the cache.", func() {
		cache.Set("test-key", "test-value", time.Second*10)

		item, _ := cache.Get("test-key")

		Expect(item.Value).To(Equal("test-value"))
	})

	It("should remove an item from the cache.", func() {
		key := "test-removal-key"
		cache.Set(key, "data", time.Second*10)
		found, err := cache.Exists(key)
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(Equal(true))

		cache.Delete(key)
		found, err = cache.Exists(key)
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(Equal(false))
	})

	It("should remove all keys", func() {
		key := "test-removal-key"
		cache.Set(key, "data", time.Second*10)
		found, err := cache.Exists(key)
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(Equal(true))

		cache.FlushAll()
		found, err = cache.Exists(key)
		Expect(err).ToNot(HaveOccurred())
		Expect(found).To(Equal(false))
	})
})
