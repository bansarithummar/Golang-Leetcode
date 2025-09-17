type Item struct {
	name   string
	rating int
}

type FoodHeap []Item

func (h FoodHeap) Len() int { return len(h) }
func (h FoodHeap) Less(i, j int) bool {
	if h[i].rating == h[j].rating {
		return h[i].name < h[j].name 
	}
	return h[i].rating > h[j].rating 
}
func (h FoodHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *FoodHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}
func (h *FoodHeap) Pop() interface{} {
	old := *h
	n := len(old)
	it := old[n-1]
	*h = old[:n-1]
	return it
}

type FoodRatings struct {
	foodToCuisine map[string]string
	foodToRating  map[string]int
	cuisineHeap   map[string]*FoodHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		foodToCuisine: make(map[string]string),
		foodToRating:  make(map[string]int),
		cuisineHeap:   make(map[string]*FoodHeap),
	}
	for i, f := range foods {
		c := cuisines[i]
		r := ratings[i]
		fr.foodToCuisine[f] = c
		fr.foodToRating[f] = r
		if _, ok := fr.cuisineHeap[c]; !ok {
			h := &FoodHeap{}
			fr.cuisineHeap[c] = h
			heap.Init(h)
		}
		heap.Push(fr.cuisineHeap[c], Item{name: f, rating: r})
	}
	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	this.foodToRating[food] = newRating
	c := this.foodToCuisine[food]
	if _, ok := this.cuisineHeap[c]; !ok {
		h := &FoodHeap{}
		this.cuisineHeap[c] = h
		heap.Init(h)
	}
	heap.Push(this.cuisineHeap[c], Item{name: food, rating: newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h, ok := this.cuisineHeap[cuisine]
	if !ok || h.Len() == 0 {
		return ""
	}
	for h.Len() > 0 {
		top := (*h)[0]
		curr := this.foodToRating[top.name]
		if top.rating == curr {
			return top.name
		}
		heap.Pop(h) 
	}
	return ""
}
