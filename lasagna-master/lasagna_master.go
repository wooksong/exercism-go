package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layerDescs []string, requiredMinPerLayer int) int {
	if requiredMinPerLayer <= 0 {
		requiredMinPerLayer = 2
	}

	return len(layerDescs) * requiredMinPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layerDescs []string) (int, float64) {
	const gramsPerNoodles = 50
	const litersPerSauce = 0.2

	numNoodlesLayers := 0
	numSauceLayers := 0
	for _, v := range layerDescs {
		switch v {
		case "noodles":
			numNoodlesLayers++
		case "sauce":
			numSauceLayers++
		}
	}
	return numNoodlesLayers * gramsPerNoodles, float64(numSauceLayers) * litersPerSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, mine []string) {
	secretingredient := friendsList[len(friendsList)-1]
	mine[len(mine)-1] = secretingredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsForTwoPortions []float64, portionsToCook int) []float64 {
	ret := make([]float64, len(amountsForTwoPortions))
	scale := float64(portionsToCook) / float64(2)

	for i, v := range amountsForTwoPortions {
		ret[i] = v * scale
	}

	return ret
}
