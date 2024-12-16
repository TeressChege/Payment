package models

type CommandId string

const (
	CustomerToMerchantMobileMoneyPayment CommandId = "CustomerToMerchantMobileMoneyPayment"
	MerchantToCustomerMobileMoneyPayment CommandId = "MerchantToCustomerMobileMoneyPayment"
	MerchantTo3rdPartyMerchantPayment    CommandId = "MerchantTo3rdPartyMerchantPayment"
	MerchantTo3rdPartyBusinessPayment    CommandId = "MerchantTo3rdPartyBusinessPayment"
	MerchantToCustomerBankPayment        CommandId = "MerchantToCustomerBankPayment"
	
)

type ServiceProviderId string

const (
	MPESA       ServiceProviderId = "MPESA"
	AIRTELMONEY ServiceProviderId = "AIRTELMONEY"
	PESALINK    ServiceProviderId = "PESALINK"
)
