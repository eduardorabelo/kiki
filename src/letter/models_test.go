package letter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/schollz/kiki/src/keypair"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	zack, _ := keypair.New()
	bob, _ := keypair.New()
	jane, _ := keypair.New()
	donald, _ := keypair.New()
	regionKey, _ := keypair.New()

	l := letter.NewText("hello, bob and jane")
	e, err := New(l, zack, []keypair.KeyPair{bob, jane}, regionKey)
	assert.Nil(t, err)

	_, err = json.Marshal(e)
	assert.Nil(t, err)

	err = e.Unseal([]keypair.KeyPair{zack}, regionKey)
	assert.Nil(t, err)
	assert.Equal(t, e.Letter.Text, "hello, bob and jane")

	err = e.Unseal([]keypair.KeyPair{bob}, regionKey)
	assert.Nil(t, err)
	err = e.Unseal([]keypair.KeyPair{jane}, regionKey)
	assert.Nil(t, err)
	err = e.Unseal([]keypair.KeyPair{donald}, regionKey)
	assert.NotNil(t, err)

	bE, _ := json.Marshal(e)
	ioutil.WriteFile("e.json", bE, 0644)

	myPeople := []keypair.KeyPair{donald, bob, jane, zack}
	err = e.Unseal(myPeople, regionKey)
	for _, p := range myPeople {
		fmt.Println(p.Keys.Public)
	}
	for _, p := range e.DeterminedRecipients {
		fmt.Println(p)
	}

	eBytes, _ := json.Marshal(e)
	fmt.Println(string(eBytes))

	err = e.Validate(regionKey)
	assert.Nil(t, err)
	err = e.Validate(donald)
	assert.NotNil(t, err)
}
