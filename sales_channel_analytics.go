package shopware

// SalesChannelAnalytics see:
// https://github.com/shopware/platform/blob/6.2/src/Core/System/SalesChannel/Aggregate/SalesChannelAnalytics/SalesChannelAnalyticsEntity.php
type SalesChannelAnalytics struct {
	Entity
	TrackingId   string        `json:"trackingId"`
	Active       bool          `json:"active"`
	TrackOrders  bool          `json:"trackOrders"`
	SalesChannel *SalesChannel `json:"salesChannel"`
}
