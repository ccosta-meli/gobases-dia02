/*
05 - Calcular cantidad de alimento.

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:

1. Perro 10 kg de alimento.
2. Gato 5 kg de alimento.
3. Hamster 250 g de alimento.
4. Tarántula 150 g de alimento.

Se solicita:

1. Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
2. Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

package internal

func Animal(ani string) (func(cant int) float64, string) {
	switch ani {
	case "Perro":
		return animalDog, ""
	case "Gato":
		return animalCat, ""
	case "Hamster":
		return animalHamster, ""
	case "Araña":
		return animalSpider, ""
	}
	return nil, "No existe el animal"
}

func animalDog(cant int) float64 {
	return float64(cant * 10)
}

func animalCat(cant int) float64 {
	return float64(cant * 5)
}

func animalHamster(cant int) float64 {
	return float64(cant) * 0.25
}

func animalSpider(cant int) float64 {
	return float64(cant) * 0.125
}
