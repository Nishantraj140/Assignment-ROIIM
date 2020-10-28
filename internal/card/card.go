package card

type VerifyReq struct {
	Card           Card           `json:"card"`
	MerchantRefNum string         `json:"merchantRefNum"`
	Profile        Profile        `json:"profile"`
	BillingDetails BillingDetails `json:"billingDetails"`
	CustomerIP     string         `json:"customerIp"`
	DupCheck       bool           `json:"dupCheck"`
	Desc           string         `json:"description"`
}

type Card struct {
	CardNum    string `json:"cardNum"`
	CardExpiry Expiry `json:"cardExpiry"`
	Cvv        string `json:"cvv"`
}

type Expiry struct {
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Profile struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
}

type BillingDetails struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
	Phone   string `json:"phone"`
}

type DetailVerifyResp struct {
	Id                        string `json:"id"`
	Merchantrefnum            string `json:"merchantRefNum"`
	Txntime                   string `json:"txnTime"`
	Status                    string `json:"status"`
	Card                      CardRes
	Cardexpiry                Expiry
	Authcode                  string
	Profile                   Profile `json:"profile"`
	BillingDetails            BillingDetails
	CustomerIp                string
	MerchantDescriptor        MerchantDescriptor
	VisaadditionalAuthDetails string
	Description               string
	Currencycode              string
	Avsresponse               string
	Cvvverification           string
	Links                     []Link
}

type CardRes struct {
	Type       string
	Lastdigits string
}

type MerchantDescriptor struct {
	DynamicDescriptor string
	Phone             string
}

type Link struct {
	Rel  string
	Href string
}

type CardDetails struct {
	Id                 string    `json:"id"`
	MerchantRefNum     string    `json:"merchantRefNum"`
	Status             string    `json:"status"`
	Usage              string    `json:"usage"`
	PaymentType        string    `json:"paymentType"`
	Action             string    `json:"action"`
	PaymentHandleToken string    `json:"paymentHandleToken"`
	BillingDetailsId   string    `json:"billingDetailsId"`
	Card               CardBrief `json:"card"`
}

type CreateCard struct {
	MerchantRefNum   string `json:"merchantRefNum"`
	PaymentType      string `json:"paymentType"`
	CurrencyCode     string `json:"currencyCode"`
	CustomerIp       string `json:"customerIp"`
	BillingDetailsId string `json:"billingDetailsId"`
	Card             Card   `json:"card"`
}

type CardBrief struct {
	LastDigits string `json:"lastDigits"`
	CardExpiry Expiry `json:"cardExpiry"`
}
