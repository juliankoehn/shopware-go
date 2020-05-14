package shopware

// CustomerRecovery see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Customer/Aggregate/CustomerRecovery/CustomerRecoveryEntity.php
// sometimes shopware devs make weirdo things
// EntityIdTrait mutates ID to customerRecovery
// but they still reference it?
type CustomerRecovery struct {
	Entity
	CustomerID string    `json:"customerId"`
	Hash       string    `json:"hash"`
	Customer   *Customer `json:"customer"`
}
