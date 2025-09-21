type Price = int

func pack(shop, movie, price int) Price {
	return (price << 49) | (shop << 17) | (movie << 1)
}

func unpack(x Price) (shop, movie, price, rented int) {
	rented = x & 1
	movie = (x >> 1) & 0xFFFF
	shop = (x >> 17) & 0xFFFFFFFF
	price = x >> 49
	return
}

type MovieRentingSystem struct {
	priceMap                 map[[2]int]Price
	cheapestUnrentedByMovies map[int][]Price
	cheapestRented           []Price
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	prices := make(map[[2]int]Price, len(entries))
	movies := make(map[int][]Price, 32)
	for _, entry := range entries {
		price := pack(entry[0], entry[1], entry[2])
		movies[entry[1]] = append(movies[entry[1]], price)
		prices[[2]int{entry[0], entry[1]}] = price
	}
	for movie := range movies {
		slices.Sort(movies[movie])
	}
	return MovieRentingSystem{
		priceMap:                 prices,
		cheapestUnrentedByMovies: movies,
		cheapestRented:           make([]Price, 0, 16),
	}
}

func (mvr *MovieRentingSystem) Search(movie int) []int {
	cheapestUnrented := make([]int, 0, 5)
	unrented := mvr.cheapestUnrentedByMovies[movie]
	for i := 0; i < len(unrented) && len(cheapestUnrented) < 5; i++ {
		shop, _, _, _ := unpack(unrented[i])
		if mvr.priceMap[[2]int{shop, movie}]&1 == 0 {
			cheapestUnrented = append(cheapestUnrented, shop)
		}
	}
	return cheapestUnrented
}

func (mvr *MovieRentingSystem) Rent(shop int, movie int) {
	mvr.priceMap[[2]int{shop, movie}] |= 1
	mvr.cheapestRented = append(mvr.cheapestRented, mvr.priceMap[[2]int{shop, movie}])
}

func (mvr *MovieRentingSystem) Drop(shop int, movie int) {
	price := mvr.priceMap[[2]int{shop, movie}]
	i := slices.Index(mvr.cheapestRented, price)
	n := len(mvr.cheapestRented) - 1
	mvr.cheapestRented[i], mvr.cheapestRented[n] = mvr.cheapestRented[n], mvr.cheapestRented[i]
	mvr.cheapestRented = mvr.cheapestRented[0:n]
	mvr.priceMap[[2]int{shop, movie}] = price &^ 1
}

func (mvr *MovieRentingSystem) Report() [][]int {
	slices.Sort(mvr.cheapestRented)
	n := min(5, len(mvr.cheapestRented))
	report := make([][]int, n)
	for i := range n {
		shop, movie, _, _ := unpack(mvr.cheapestRented[i])
		report[i] = []int{shop, movie}
	}
	return report
}
