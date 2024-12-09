package models

type Item struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Discount struct {
	Coupon   Coupon   `json:"coupon"`
	OnTop    OnTop    `json:"onTop"`
	Seasonal Seasonal `json:"seasonal"`
}

type Coupon struct {
	Campaigns  string  `json:"campaigns"`
	Parameters int `json:"parameters"`
}

type OnTop struct {
	Campaigns  string          `json:"campaigns"`
	Parameters OnTopParameters `json:"parameters"`
}

type Seasonal struct {
	Campaigns  string             `json:"campaigns"`
	Parameters SeasonalParameters `json:"parameters"`
}

type OnTopParameters struct {
	Category string  `json:"cate"`
	Discount int `json:"discount"`
	Max      int `json:"Max"`
}

type SeasonalParameters struct {
	Every    int `json:"every"`
	Discount int `json:"discount"`
}

type Request struct {
	Items    []Item   `json:"items"`
	Discount Discount `json:"discount"`
}
