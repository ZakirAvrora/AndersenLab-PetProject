package models

type User struct {
	UserId   int     `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Username string  `json:"username" db:"username"`
	Email    string  `json:"email" db:"email"`
	Address  Address `json:"address" db:"address"`
	Phone    string  `json:"phone" db:"phone"`
	Website  string  `json:"website" db:"website"`
	Company  Company `json:"company" db:"company"`
}

type Address struct {
	Street  string `json:"street" db:"street"`
	Suite   string `json:"suite" db:"suite"`
	City    string `json:"city" db:"city"`
	Zipcode string `json:"zipcode" db:"zipcode"`
	Geo     Geo    `json:"geo" db:"geo"`
}

type Geo struct {
	Lat string `json:"lat" db:"lat"`
	Lng string `json:"lng" db:"lng"`
}

type Company struct {
	Name        string `json:"name" db:"company_name"`
	CatchPhrase string `json:"catchPhrase" db:"catch_phase"`
	Bs          string `json:"bs" db:"BS"`
}
