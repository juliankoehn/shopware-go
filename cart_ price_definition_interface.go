package shopware

// PriceDefinitionInterface see:
// https://github.com/shopware/platform/blob/6.2/src/Core/Checkout/Cart/Price/Struct/PriceDefinitionInterface.php
type PriceDefinitionInterface interface {
	// GetPrecision returns the decimal precision for the price.
	// Necessary for \Shopware\Core\Checkout\Cart\Price\PriceRounding::round
	GetPrecision() int

	// GetType returns the type of this PriceDefinition.
	// At the moment three definitions exist: AbsolutePriceDefinition,
	// QuantityPriceDefinition and PercentagePriceDefinition.
	// The type of the definition changes how a price is
	// calculated, but they all share the same datastructure.
	// See the corresponding classes for exactly how each
	// of them works.
	GetType() string

	// GetPriority returns the priority of this price definitions,
	// which determines in which order prices are calculated.
	// Some PriceDefinitions change the final price based on the amount of
	// the prices already calculated, and thus can only be calculated
	// after all others have finished. This applies for example to
	// percental discounts.
	GetPriority() string

	// GetConstraints Returns the constraints of this PriceDefinitions.
	// These are used by PriceDefinitions which calculate their final
	// price from the results of other PriceDefinitions to filter the
	// items they want to apply to. This is used for example to create
	// discounts, which only apply to items of a certain type.
	GetConstraints() *Constraint
}
