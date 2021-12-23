package testing

import (
	"os"
	"os/exec"
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

func TestRun(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		server.Run("crasher")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestRun")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
