package main

import (
	"fmt"
	"os"

	"github.com/TeressChege/payment"
	"github.com/TeressChege/payment/models"
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
	TestWalletBalances()
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
			{ID: "shortCode", Value: "xxxx", Label: "Short code"},
			{ID: "accountNumber", Value: "xxxx", Label: "Phone number"},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", resp)
	respp, err := tandaPay.StatusQuery("xxx", "xxxx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", respp)
}

// Testing for paybills
func TestB2BPAYBILL() {
	rep, err := tandaPay.MerchantToCustomer(models.PaymentRequest{
		CommandId:         models.MerchantTo3rdPartyBusinessPayment,
		ServiceProviderId: string(models.MPESA),
		Reference:         "MY2549921",
		CallBackUrl:       "https://webhook.site/9666c087-579a-48cc-b3ae-0256bd52e0d1",
		AccountNumber:     "xxx",
		AccountNumberRef:  "xxx",
		AccountName:       "test",
		Narration:         "Payment for testing",
		ShortCode:         "xxxx",
		Amount:            "100",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", rep)
}

func TestB2BTILL() {
	rep, err := tandaPay.MerchantToCustomer(models.PaymentRequest{
		CommandId:         models.MerchantTo3rdPartyMerchantPayment,
		ServiceProviderId: string(models.MPESA),
		Reference:         "MY25499921",
		CallBackUrl:       "https://webhook.site/9666c087-579a-48cc-b3ae-0256bd52e0d1",
		AccountNumber:     "xxxxx",
		AccountNumberRef:  "xxxx",
		AccountName:       "test",
		Narration:         "Payment for testing",
		ShortCode:         "xxxx",
		Amount:            "100",
		BankCode:          "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", rep)
}

// Send funds to mpesa,airtel customer
func TestB2C() {
	rep, err := tandaPay.MerchantToCustomer(models.PaymentRequest{
		CommandId:         models.MerchantToCustomerMobileMoneyPayment,
		ServiceProviderId: string(models.MPESA),
		Reference:         "MY240d3U4992136",
		CallBackUrl:       "https://webhook.site/9666c087-579a-48cc-b3ae-0256bd52e0d1",
		AccountNumber:     "xxxxx",
		AccountName:       "Tanda",
		Narration:         "Payment for testing",
		ShortCode:         "xxxxx",
		Amount:            "150",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", rep)
}
func TestB2BANK() {
	rep, err := tandaPay.MerchantToCustomer(models.PaymentRequest{
		CommandId:         models.MerchantToCustomerBankPayment,
		ServiceProviderId: string(models.PESALINK),
		Reference:         "MY240699273136",
		CallBackUrl:       "https://webhook.site/9666c087-579a-48cc-b3ae-0256bd52e0d1",
		AccountNumber:     "xxxx",
		AccountNumberRef:  "xxx",
		AccountName:       "test",
		Narration:         "Payment for testing",
		ShortCode:         "xxxxx",
		Amount:            "150",
		BankCode:          "xxxx",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", rep)
}

func TestWalletBalances() {
	rep, err := tandaPay.GetWalletBalances("xxxxxx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %v", rep)
}
