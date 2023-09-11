package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimalDog(t *testing.T) {
	qDog := 10
	alimento := animalDog(qDog)
	assert.Equal(t, 100.0, alimento, "Alimento deberia ser 100 kg")
}

func TestAnimalCat(t *testing.T) {
	qCat := 10
	alimento := animalCat(qCat)
	assert.Equal(t, 50.0, alimento, "Alimento deberia ser 50 kg")
}

func TestAnimalHamster(t *testing.T) {
	qHamster := 10
	alimento := animalHamster(qHamster)
	assert.Equal(t, 2.5, alimento, "Alimento deberia ser 2.5 kg")
}

func TestAnimalSpider(t *testing.T) {
	qSpider := 10
	alimento := animalSpider(qSpider)
	assert.Equal(t, 1.25, alimento, "Alimento deberia ser 1.25 kg")
}
