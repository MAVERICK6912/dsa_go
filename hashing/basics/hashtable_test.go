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

var testDataString = []struct {
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

var testDataInt = []struct {
	HeroNumber  int
	HeroDetails SuperHero
}{
	{
		HeroNumber:  45613548,
		HeroDetails: SuperHero{OriginalIdentity: "Bruce Wayne", City: "Gotham"},
	},
	{
		HeroNumber:  78123545,
		HeroDetails: SuperHero{OriginalIdentity: "Clark Kent", City: "Metropolis"},
	},
	{
		HeroNumber:  32156974,
		HeroDetails: SuperHero{OriginalIdentity: "Arthur Curry", City: "Atlantis"},
	},
}

func TestInsertAndGet(t *testing.T) {
	var ht *Hashtable[string, SuperHero]
	ht = ht.New(utils.StringHash[string], 3)

	for _, v := range testDataString {
		ht.Insert(v.HeroName, v.HeroDetails)
	}

	actual, err := ht.Search(testDataString[2].HeroName)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testDataString[2].HeroDetails, actual)

	actual, err = ht.Search(testDataString[0].HeroName)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testDataString[0].HeroDetails, actual)

	actual, err = ht.Search(testDataString[1].HeroName)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testDataString[1].HeroDetails, actual)

	_, err = ht.Search("hawk girl")
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}

func TestRemove(t *testing.T) {
	var ht *Hashtable[string, SuperHero]
	ht = ht.New(utils.StringHash[string], 3)

	for _, v := range testDataString {
		ht.Insert(v.HeroName, v.HeroDetails)
	}

	err := ht.Remove(testDataString[1].HeroName)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testDataString[1].HeroName)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove(testDataString[2].HeroName)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testDataString[2].HeroName)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove(testDataString[0].HeroName)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testDataString[0].HeroName)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
	// removing a key which does not exist in hashtable
	err = ht.Remove("flash")
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}

func TestInsertAndGetInt(t *testing.T) {
	var ht *Hashtable[int, SuperHero]
	ht = ht.New(utils.IntHash[int], 3)

	for _, v := range testDataInt {
		ht.Insert(v.HeroNumber, v.HeroDetails)
	}

	actual, err := ht.Search(testDataInt[2].HeroNumber)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testDataInt[2].HeroDetails, actual)

	actual, err = ht.Search(testDataInt[0].HeroNumber)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testDataInt[0].HeroDetails, actual)

	actual, err = ht.Search(testDataInt[1].HeroNumber)
	if err != nil {
		t.Error("error while searching in hashtable, error was: ", err.Error())
	}
	assert.Equal(t, testDataInt[1].HeroDetails, actual)

	_, err = ht.Search(752365)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}

func TestRemoveInt(t *testing.T) {
	var ht *Hashtable[int, SuperHero]
	ht = ht.New(utils.IntHash[int], 3)

	for _, v := range testDataInt {
		ht.Insert(v.HeroNumber, v.HeroDetails)
	}

	err := ht.Remove(testDataInt[1].HeroNumber)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testDataInt[1].HeroNumber)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove(testDataInt[2].HeroNumber)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testDataInt[2].HeroNumber)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	err = ht.Remove(testDataInt[0].HeroNumber)
	if err != nil {
		t.Error("error while removing from hashtable, error was: ", err.Error())
	}
	assert.Nil(t, err)
	_, err = ht.Search(testDataInt[0].HeroNumber)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}

	// removing a key which does not exist in hashtable
	err = ht.Remove(123648)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NotFound)
	}
}
