package kafka

type ProfileForNormalization struct {
	Name        string
	Age         string
	City        string
	Height      string
	Weight      string
	BreastSize  string
	HairColor   string
	Nationality string
	EyeColor    string
	Piercings   bool
	Prices      map[string]string
	Phone       string
	Images      []string
}
