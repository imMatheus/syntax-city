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
	{
		Id:            "brown-2",
		Name:          "Array Alley",
		Street:        "Brown",
		BuyPrice:      60,
		HousePrice:    50,
		HotelPrice:    50,
		MortgageValue: 30,
		Rents: []PropertyRent{
			{NumberOfHouses: 0, Rent: 4},
			{NumberOfHouses: 1, Rent: 20},
			{NumberOfHouses: 2, Rent: 60},
			{NumberOfHouses: 3, Rent: 180},
			{NumberOfHouses: 4, Rent: 320},
			{NumberOfHouses: 5, Rent: 450},
		},
	},
	{
		Id:            "lightblue-1",
		Name:          "Function Falls",
		Street:        "Light Blue",
		BuyPrice:      100,
		HousePrice:    50,
		HotelPrice:    50,
		MortgageValue: 50,
		Rents: []PropertyRent{
			{NumberOfHouses: 0, Rent: 6},
			{NumberOfHouses: 1, Rent: 30},
			{NumberOfHouses: 2, Rent: 90},
			{NumberOfHouses: 3, Rent: 270},
			{NumberOfHouses: 4, Rent: 400},
			{NumberOfHouses: 5, Rent: 550},
		},
	},
	{
		Id:            "lightblue-2",
		Name:          "Method Meadows",
		Street:        "Light Blue",
		BuyPrice:      100,
		HousePrice:    50,
		HotelPrice:    50,
		MortgageValue: 50,
		Rents: []PropertyRent{
			{NumberOfHouses: 0, Rent: 6},
			{NumberOfHouses: 1, Rent: 30},
			{NumberOfHouses: 2, Rent: 90},
			{NumberOfHouses: 3, Rent: 270},
			{NumberOfHouses: 4, Rent: 400},
			{NumberOfHouses: 5, Rent: 550},
		},
	},
	{
		Id:            "lightblue-3",
		Name:          "Class Cove",
		Street:        "Light Blue",
		BuyPrice:      120,
		HousePrice:    50,
		HotelPrice:    50,
		MortgageValue: 60,
		Rents: []PropertyRent{
			{NumberOfHouses: 0, Rent: 8},
			{NumberOfHouses: 1, Rent: 40},
			{NumberOfHouses: 2, Rent: 100},
			{NumberOfHouses: 3, Rent: 300},
			{NumberOfHouses: 4, Rent: 450},
			{NumberOfHouses: 5, Rent: 600},
		},
	},
}
