package paymetgateway

import (
	"github.com/midtrans/midtrans-go"
	_ "github.com/midtrans/midtrans-go"
	_ "github.com/midtrans/midtrans-go/coreapi"
	_ "github.com/midtrans/midtrans-go/iris"
	"github.com/midtrans/midtrans-go/snap"
)

func CreatePayment() (*snap.Response, error) {
	var s = snap.Client{}

	s.New("SB-Mid-server-SncqFmlAhQF1Or6Wd8rsoLuw", midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-1234522+12222",
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, err := s.CreateTransaction(req)

	if err != nil {
		return nil, err
	}

	return snapResp, nil

}
