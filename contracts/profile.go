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
	HairColorID   int       `json:"hair_color_id" db:"hair_color_id"`
	NationalityID int       `json:"nationality_id" db:"nationality_id"`
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

type CreateProfile struct {
	Name          string   `json:"name" `
	Description   string   `json:"description"`
	Age           int      `json:"age"`
	Height        int      `json:"height"`
	Weight        int      `json:"weight"`
	BreastSize    string   `json:"breast_size"`
	HairColorID   int      `json:"hair_color_id"`
	NationalityID int      `json:"nationality_id"`
	CityID        int      `json:"city_id"`
	DistrictIDs   []int    `json:"district_ids"`
	MetroIDs      []int    `json:"metro_ids"`
	Prices        Prices   `json:"prices"`
	Phone         string   `json:"phone"`
	Images        []string `json:"images"`
}
