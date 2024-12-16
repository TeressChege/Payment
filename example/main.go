package main

import (
	"fmt"
	"os"

	"github.com/danchengash/tanda-plugin"
	"github.com/danchengash/tanda-plugin/models"
	"github.com/joho/godotenv"
)

// defaults
func init() {
	godotenv.Load()
	organizationId = os.Getenv("ORGID")
	clientid = os.Getenv("CLIENTID")
	clientSecret = os.Getenv("CLIENTSECRET")
	shortCode = os.Getenv("SHORTCODE")
}

var organizationId = ""
var clientid = ""
var clientSecret = ""
var shortCode = ""
var tandaPay = tanda.Tanda{}

func main() {
	tandaPay = *tanda.NewTanda(clientid, clientSecret, organizationId, shortCode, tanda.Production, true)
	rep, er := tandaPay.MerchantToCustomer(models.PaymentRequest{
		CommandId:         models.MerchantTo3rdPartyBusinessPayment,
		ServiceProviderId: string(models.MPESA),
		Reference:         "MY2WA90d3U4",
		CallBackUrl:       "https://webhook.site/d03b51de-36a8-464f-ba6f-0c18144105b4",
		AccountNumber:     "756756",
		AccountNumberRef:  "5055",
		AccountName:       "Tanda",
		Narration:         "Payment for testing",
		ShortCode:         "220429",
		Amount:            "50",
		BankCode:          "",
	})
	if er != nil {
		panic(er)
	}
	fmt.Printf("Response: %v", rep)
	// res, er := tandaPay.GetWalletBalances()
	// if er != nil {
	// 	fmt.Println(er)
	// }
	// fmt.Println(res)
}
func TestC2b() {

	resp, err := tandaPay.ReceivePayment(models.PaymentRequest{
		CommandId:         models.CustomerToMerchantMobileMoneyPayment,
		ServiceProviderId: string(models.MPESA),
		Reference:         "MY2WA9034",
		Request: []models.PaymentField{
			{ID: "amount", Value: "50", Label: "Amount"},
			{ID: "narration", Value: "Payment for testing", Label: "Narration"},
			{ID: "ipnUrl", Value: "https://webhook.site/ee627e14-3b4e-46b7-aec6-4625b5bdd629", Label: "Notification URL"},
			{ID: "shortCode", Value: "220429", Label: "Short code"},
			{ID: "accountNumber", Value: "254703545191", Label: "Phone number"},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", resp)
	respp, err := tandaPay.StatusQuery("909c9353-a003-4834-944c-7e274c10ed3d", "220429")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", respp)
}
