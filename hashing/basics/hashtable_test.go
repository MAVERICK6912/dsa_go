package hashing

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

type SuperHero struct {
	OriginalIdentity string
	City             string
}

var testData = []struct {
	HeroName    string
	HeroDetails SuperHero
}{
	{
		HeroName:    "batman",
		HeroDetails: SuperHero{OriginalIdentity: "Bruce Wayne", City: "Gotham"},
	},
	{
		HeroName:    "superman",
		HeroDetails: SuperHero{OriginalIdentity: "Clark Kent", City: "Metropolis"},
	},
	{
		HeroName:    "aquaman",
		HeroDetails: SuperHero{OriginalIdentity: "Arthur Curry", City: "Atlantis"},
	},
}

func TestInsertAndGet(t *testing.T) {
	var ht *Hashtable[string, SuperHero]
	ht = ht.New(utils.StringHash[string], 3)

	for _, v := range testData {
		ht.Insert(v.HeroName, v.HeroDetails)
	}

	actual, err := ht.Search(testData[2].HeroName)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testData[2].HeroDetails, actual)

	actual, err = ht.Search(testData[0].HeroName)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testData[0].HeroDetails, actual)

	actual, err = ht.Search(testData[1].HeroName)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testData[1].HeroDetails, actual)

	_, err = ht.Search("hawk girl")
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}

func TestRemove(t *testing.T) {
	var ht *Hashtable[string, SuperHero]
	ht = ht.New(utils.StringHash[string], 3)

	for _, v := range testData {
		ht.Insert(v.HeroName, v.HeroDetails)
	}

	err := ht.Remove(testData[1].HeroName)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testData[1].HeroName)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove(testData[2].HeroName)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testData[2].HeroName)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove(testData[0].HeroName)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testData[0].HeroName)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove("flash")
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}
