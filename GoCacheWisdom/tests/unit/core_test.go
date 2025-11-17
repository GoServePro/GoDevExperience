package unit

import (
	"github.com/GoServePro/GoDevExperience/GoCacheWisdom/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicCache(t *testing.T) {
	c := cache.New()

	// 测试 Set + Get
	c.Set("key1", "value1")
	val, ok := c.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	// 测试覆盖 key
	c.Set("key1", "value2")
	val, ok = c.Get("key1")
	assert.Equal(t, "value2", val)

	// 测试 Delete
	c.Delete("key1")
	_, ok = c.Get("key1")
	assert.False(t, ok)
}
