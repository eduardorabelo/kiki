package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenClose(t *testing.T) {
	os.Remove("kiki.db")
	db, err := Open()
	assert.Nil(t, err)
	err = db.Close()
	assert.Nil(t, err)
}

func TestKeyStore(t *testing.T) {
	type A struct {
		B int
		C string
	}
	a := A{
		B: 3,
		C: "hi",
	}
	os.Remove("kiki.db")
	db, err := Open()
	assert.Nil(t, err)
	defer db.Close()
	err = db.Set("Astuff", "a", a)
	assert.Nil(t, err)
	var a2 A
	err = db.Get("Astuff", "a", &a2)
	assert.Nil(t, err)
	assert.Equal(t, a, a2)
}
func TestAddLetter(t *testing.T) {

}
