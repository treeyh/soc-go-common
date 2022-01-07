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

	assert.Equal(t, Get("zh", "aaa.asdfasdf"), "7")
	assert.Equal(t, GetByDefault("zh", "aaa.asdfaasdf", "default"), "default")
}
