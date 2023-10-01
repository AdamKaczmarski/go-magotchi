package src 

type food struct {
	happiness float32
	hunger    float32
	weight    int16
}

type feedFoods struct {
	burger food
	cake   food
}

var feedingFoods feedFoods = feedFoods{
	burger: food{
		happiness: 0.1,
		hunger:    0.75,
		weight:    3,
	},
	cake: food{
		happiness: 0.5,
		hunger:    0.5,
		weight:    4,
	},
}
