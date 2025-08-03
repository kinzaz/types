package contracts

import "time"

type Profile struct {
	ID            string    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Description   string    `json:"description" db:"description"`
	CityID        int       `json:"city_id" db:"city_id"`
	Age           int       `json:"age" db:"age"`
	Height        int       `json:"height" db:"height"`
	Weight        int       `json:"weight" db:"weight"`
	BreastSize    string    `json:"breast_size" db:"breast_size"`
	HairColorID   string    `json:"hair_color_id" db:"hair_color_id"`
	NationalityID string    `json:"nationality_id" db:"nationality_id"`
	Prices        Prices    `json:"prices" db:"prices"`
	Phone         string    `json:"phone" db:"phone"`
	Images        []string  `json:"images" db:"images"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type Prices struct {
	OneHour int `json:"1_hour"`
	TwoHour int `json:"2_hour"`
	Night   int `json:"night"`
}
