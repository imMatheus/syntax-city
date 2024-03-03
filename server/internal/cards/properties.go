package cards

// each step of rent based on the number of houses/hotels
type PropertyRent struct {
	NumberOfHouses int `json:"numberOfHouses"`
	Rent           int `json:"rent"`
}

type PropertyCard struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Street        string         `json:"street"`
	BuyPrice      int32          `json:"buyPrice"`
	HousePrice    int            `json:"housePrice"`
	HotelPrice    int            `json:"hotelPrice"`
	MortgageValue int            `json:"mortgageValue"`
	Rents         []PropertyRent `json:"rents"`
}

var Properties = []PropertyCard{
	{
		Id:         "brown-1",
		Name:       "Dungeon of despair",
		Street:     "Brown",
		BuyPrice:   60,
		HousePrice: 50,
		HotelPrice: 50,
		Rents: []PropertyRent{
			{
				NumberOfHouses: 0,
				Rent:           2,
			},
			{
				NumberOfHouses: 1,
				Rent:           10,
			},
			{
				NumberOfHouses: 2,
				Rent:           30,
			},
			{
				NumberOfHouses: 3,
				Rent:           90,
			},
			{
				NumberOfHouses: 4,
				Rent:           160,
			},
			{
				NumberOfHouses: 5,
				Rent:           250,
			},
		},
	},
}
