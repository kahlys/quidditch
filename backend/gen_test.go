package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generatePlayers(t *testing.T) {
	for _, ply := range generatePlayers() {
		assert.NotEmpty(t, ply.FirstName, "firstname")
		assert.NotEmpty(t, ply.LastName, "lastname")
		assert.NotEmpty(t, ply.Country, "nationality")
		assert.NotEmpty(t, ply.Power, "power")
		assert.NotEmpty(t, ply.Stamina, "power")
	}
}
