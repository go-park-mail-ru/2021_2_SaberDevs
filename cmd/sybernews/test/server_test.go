package testing

import (
	"testing"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		db, err := server.DbConfig()
		assert.NoError(t, err)
		assert.NotEmpty(t, db)
	})
}

func TestDbConnect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		db, err := server.DbConnect()
		assert.NoError(t, err)
		assert.NotEmpty(t, db)
	})
}

func TestTarantoolConnect(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		db, err := server.TarantoolConnect()
		assert.NoError(t, err)
		assert.NotEmpty(t, db)
	})
}

func TestNewConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("success", func(t *testing.T) {
		log := wrapper.NewLogger()
		assert.NotEmpty(t, log)
	})
}
