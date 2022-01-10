package i18n

import (
	"github.com/stretchr/testify/assert"
	"github.com/treeyh/soc-go-common/core/config"
	"testing"
)

func TestInitI18n(t *testing.T) {
	InitI18n(&config.I18nConfig{
		Enable: true,
		Path:   "D:\\i18n",
	})

	assert.Equal(t, Get("zh", "aaa.aaa"), "%s 4 3")
	assert.Equal(t, Get("zh", "aaa.bbb"), "5 %s 41")
	assert.Equal(t, Get("zh", "aaa.ccc"), "61 13")
	assert.Equal(t, GetByDefault("zh", "aaa.ddd", "default"), "default")
}
