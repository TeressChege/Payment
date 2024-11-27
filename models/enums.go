package models

type CommandId string

const (
	CustomerToMerchantMobileMoneyPayment     CommandId = "CustomerToMerchantMobileMoneyPayment"
	MerchantToCustomerMobileMoneyPayment     CommandId = "MerchantToCustomerMobileMoneyPayment"
	MerchantTo3rdPartyMerchantPayment        CommandId = "MerchantTo3rdPartyMerchantPayment"
)

type ServiceProviderId string

const (
	MPESA       ServiceProviderId = "MPESA"
	AIRTELMONEY ServiceProviderId = "AIRTELMONEY"
)
