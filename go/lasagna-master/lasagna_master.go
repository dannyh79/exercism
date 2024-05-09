package lasagna

func PreparationTime(l []string, t int) int {
	const defaultTime = 2
	if t == 0 {
		t = defaultTime
	}
	return len(l) * t
}

func Quantities(l []string) (int, float64) {
	const (
		noodles         = "noodles"
		noodlesPerLayer = 50
		sauce           = "sauce"
		saucePerLayer   = 0.2
	)
	var qn, qs int
	for _, i := range l {
		switch i {
		case noodles:
			qn++
		case sauce:
			qs++
		}
	}

	return qn * noodlesPerLayer, float64(qs) * saucePerLayer
}

func AddSecretIngredient(fl, ml []string) {
	ml[len(ml)-1] = fl[len(fl)-1]
}

func ScaleRecipe(i []float64, p int) []float64 {
	const portionPerSlice = 2

	var scale = float64(p) / 2
	var r = make([]float64, len(i))
	for idx := range i {
		r[idx] = i[idx] * scale
	}
	return r
}
