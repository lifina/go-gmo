package payment

// RegisterBankAccountRequest ... input parameter for registering bank account
type RegisterBankAccountRequest struct {
	MemberID   string `json:"MemberID" validate:"required"`
	MemberName string `json:"MemberName"`
	// 0 ... do not create member, 1 ... create member
	CreateMember string `json:"CreateMember" validate:"required,oneof=0 1"`
	RetURL       string `json:"RetURL" validate:"required,url"`
	BankCode     string `json:"BankCode" validate:"required,len=4"`
	BranchCode   string `json:"BranchCode"`
	// 1 ... normal type, 2 ... touza type
	AccountType      string `json:"AccountType"`
	AccountNumber    string `json:"AccountNumber"`
	AccountName      string `json:"AccountName"`
	AccountNameKanji string `json:"AccountNameKanji"`
	ConsumerDevice   string `json:"ConsumerDevice" validate:"required,oneof=i ez sb pc"`
}

// RegisterBankAccountResponse ... response parameter
type RegisterBankAccountResponse struct {
	TrainID  string `json:"TrainID"`
	Token    string `json:"token"`
	StartURL string `json:"StartUrl"`
	ErrCode  string `json:"ErrCode"`
	ErrInfo  string `json:"ErrInfo"`
}

// RegisterBankAccount ... register bank account
func (cli *Client) RegisterBankAccount(req *RegisterBankAccountRequest) (res *RegisterBankAccountResponse, err error) {

	return nil, nil
}
