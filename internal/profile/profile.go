package profile

type CreateProfile struct {
	MerchantCustomerId string      `json:"merchantCustomerId"`
	Locale             string      `json:"locale"`
	FirstName          string      `json:"firstName"`
	MiddleName         string      `json:"middleName"`
	LastName           string      `json:"lastName"`
	DateOfBirth        DateOfBirth `json:"dateOfBirth"`
	Email              string      `json:"email"`
	Phone              string      `json:"phone"`
	Ip                 string      `json:"ip"`
	Gender             string      `json:"gender"`
	Nationality        string      `json:"nationality"`
	CellPhone          string      `json:"cellPhone"`
}

type DateOfBirth struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type CreateProfileRes struct {
	Id                 string      `json:"id"`
	MerchantCustomerId string      `json:"merchantCustomerId"`
	Locale             string      `json:"locale"`
	FirstName          string      `json:"firstName"`
	MiddleName         string      `json:"middleName"`
	LastName           string      `json:"lastName"`
	DateOfBirth        DateOfBirth `json:"dateOfBirth"`
	Email              string      `json:"email"`
	Phone              string      `json:"phone"`
	Ip                 string      `json:"ip"`
	Gender             string      `json:"gender"`
	Nationality        string      `json:"nationality"`
	CellPhone          string      `json:"cellPhone"`
	Status             string      `json:"status"`
}
