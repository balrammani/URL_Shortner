package randomgenerator

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type suite struct {
	ctx  context.Context
	ctrl *gomock.Controller
	t    *testing.T

	random RandomService
	err    error
}

func (s *suite) setupTest() {
	s.random, s.err = New()
	assert.NoError(s.t, s.err)

}

func TestGenerateRandomString(t *testing.T) {
	var s suite

	scenarios := []struct {
		name  string
		setup func()
	}{
		{
			name: "success",
			setup: func() {

			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			s.ctrl = gomock.NewController(t)
			s.setupTest()
			scenario.setup()

			response := s.random.GenerateRandomString()
			assert.Equal(t, len(response), shortURLLength)

		})
	}

}
