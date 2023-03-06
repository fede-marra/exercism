package lasagna

// PreparationTime calc time for cook of lasagna,
// receives layers and time like parameters
// for calculing cock time
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
		return (len(layers) * time)
	} else {
		return (len(layers) * time)
	}
}
// Quantities Calculate amounts of noodles and sauce,
// receives layers like parameter and returned noodles and
// sauce that need
func Quantities(layers []string) (int, float64) {
	var (
		noodles int
		sauce float64
	)
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles +=1
		} else if layers[i] == "sauce" {
			sauce += 1
		}
	}
	return noodles * 50, sauce * 0.2
}
// AddSecretIngredient Add the secret ingredient to our list using with
// reference the list of our friends
func AddSecretIngredient(friendsList, mylist []string) {
	for i :=0; i < len(friendsList); i++ {
		for j := 0; j < len(mylist); j++ {
			if friendsList[i] != mylist[j] {
				mylist[len(mylist)-1] = friendsList[i]
			}
		}
	}
}
// ScaleRecipe Calculate ingredients that need for cocking of the lasagna
// according to of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var (
		scaleQuantities []float64
		value float64
	)
	scaleQuantities = []float64{}
	if portions > 2 {
		value = float64(portions) / float64(2)
		
		for i := 0; i < len(quantities) ; i++ {
			scaleQuantities = append(scaleQuantities, (quantities[i] * value))
		}
	} else if portions < 2 {
		value = float64(portions) / float64(2)
		for i := 0; i < len(quantities) ; i++ {
			scaleQuantities = append(scaleQuantities,  quantities[i] * value)
		}
	} else {
		for i := 0; i < len(quantities) ; i++ {
			scaleQuantities = append(scaleQuantities,  quantities[i] * float64(portions))
		}
	}
	return scaleQuantities
}