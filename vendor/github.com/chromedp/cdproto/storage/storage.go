// Package storage provides the Chrome DevTools Protocol
// commands, types, and events for the Storage domain.
//
// Generated by the cdproto-gen command.
package storage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
)

// GetStorageKeyForFrameParams returns a storage key given a frame id.
type GetStorageKeyForFrameParams struct {
	FrameID cdp.FrameID `json:"frameId"`
}

// GetStorageKeyForFrame returns a storage key given a frame id.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getStorageKeyForFrame
//
// parameters:
//
//	frameID
func GetStorageKeyForFrame(frameID cdp.FrameID) *GetStorageKeyForFrameParams {
	return &GetStorageKeyForFrameParams{
		FrameID: frameID,
	}
}

// GetStorageKeyForFrameReturns return values.
type GetStorageKeyForFrameReturns struct {
	StorageKey SerializedStorageKey `json:"storageKey,omitempty"`
}

// Do executes Storage.getStorageKeyForFrame against the provided context.
//
// returns:
//
//	storageKey
func (p *GetStorageKeyForFrameParams) Do(ctx context.Context) (storageKey SerializedStorageKey, err error) {
	// execute
	var res GetStorageKeyForFrameReturns
	err = cdp.Execute(ctx, CommandGetStorageKeyForFrame, p, &res)
	if err != nil {
		return "", err
	}

	return res.StorageKey, nil
}

// ClearDataForOriginParams clears storage for origin.
type ClearDataForOriginParams struct {
	Origin       string `json:"origin"`       // Security origin.
	StorageTypes string `json:"storageTypes"` // Comma separated list of StorageType to clear.
}

// ClearDataForOrigin clears storage for origin.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearDataForOrigin
//
// parameters:
//
//	origin - Security origin.
//	storageTypes - Comma separated list of StorageType to clear.
func ClearDataForOrigin(origin string, storageTypes string) *ClearDataForOriginParams {
	return &ClearDataForOriginParams{
		Origin:       origin,
		StorageTypes: storageTypes,
	}
}

// Do executes Storage.clearDataForOrigin against the provided context.
func (p *ClearDataForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearDataForOrigin, p, nil)
}

// ClearDataForStorageKeyParams clears storage for storage key.
type ClearDataForStorageKeyParams struct {
	StorageKey   string `json:"storageKey"`   // Storage key.
	StorageTypes string `json:"storageTypes"` // Comma separated list of StorageType to clear.
}

// ClearDataForStorageKey clears storage for storage key.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearDataForStorageKey
//
// parameters:
//
//	storageKey - Storage key.
//	storageTypes - Comma separated list of StorageType to clear.
func ClearDataForStorageKey(storageKey string, storageTypes string) *ClearDataForStorageKeyParams {
	return &ClearDataForStorageKeyParams{
		StorageKey:   storageKey,
		StorageTypes: storageTypes,
	}
}

// Do executes Storage.clearDataForStorageKey against the provided context.
func (p *ClearDataForStorageKeyParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearDataForStorageKey, p, nil)
}

// GetCookiesParams returns all browser cookies.
type GetCookiesParams struct {
	BrowserContextID cdp.BrowserContextID `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// GetCookies returns all browser cookies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getCookies
//
// parameters:
func GetCookies() *GetCookiesParams {
	return &GetCookiesParams{}
}

// WithBrowserContextID browser context to use when called on the browser
// endpoint.
func (p GetCookiesParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *GetCookiesParams {
	p.BrowserContextID = browserContextID
	return &p
}

// GetCookiesReturns return values.
type GetCookiesReturns struct {
	Cookies []*network.Cookie `json:"cookies,omitempty"` // Array of cookie objects.
}

// Do executes Storage.getCookies against the provided context.
//
// returns:
//
//	cookies - Array of cookie objects.
func (p *GetCookiesParams) Do(ctx context.Context) (cookies []*network.Cookie, err error) {
	// execute
	var res GetCookiesReturns
	err = cdp.Execute(ctx, CommandGetCookies, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Cookies, nil
}

// SetCookiesParams sets given cookies.
type SetCookiesParams struct {
	Cookies          []*network.CookieParam `json:"cookies"`                    // Cookies to be set.
	BrowserContextID cdp.BrowserContextID   `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// SetCookies sets given cookies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-setCookies
//
// parameters:
//
//	cookies - Cookies to be set.
func SetCookies(cookies []*network.CookieParam) *SetCookiesParams {
	return &SetCookiesParams{
		Cookies: cookies,
	}
}

// WithBrowserContextID browser context to use when called on the browser
// endpoint.
func (p SetCookiesParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *SetCookiesParams {
	p.BrowserContextID = browserContextID
	return &p
}

// Do executes Storage.setCookies against the provided context.
func (p *SetCookiesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetCookies, p, nil)
}

// ClearCookiesParams clears cookies.
type ClearCookiesParams struct {
	BrowserContextID cdp.BrowserContextID `json:"browserContextId,omitempty"` // Browser context to use when called on the browser endpoint.
}

// ClearCookies clears cookies.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearCookies
//
// parameters:
func ClearCookies() *ClearCookiesParams {
	return &ClearCookiesParams{}
}

// WithBrowserContextID browser context to use when called on the browser
// endpoint.
func (p ClearCookiesParams) WithBrowserContextID(browserContextID cdp.BrowserContextID) *ClearCookiesParams {
	p.BrowserContextID = browserContextID
	return &p
}

// Do executes Storage.clearCookies against the provided context.
func (p *ClearCookiesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearCookies, p, nil)
}

// GetUsageAndQuotaParams returns usage and quota in bytes.
type GetUsageAndQuotaParams struct {
	Origin string `json:"origin"` // Security origin.
}

// GetUsageAndQuota returns usage and quota in bytes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getUsageAndQuota
//
// parameters:
//
//	origin - Security origin.
func GetUsageAndQuota(origin string) *GetUsageAndQuotaParams {
	return &GetUsageAndQuotaParams{
		Origin: origin,
	}
}

// GetUsageAndQuotaReturns return values.
type GetUsageAndQuotaReturns struct {
	Usage          float64         `json:"usage,omitempty"`          // Storage usage (bytes).
	Quota          float64         `json:"quota,omitempty"`          // Storage quota (bytes).
	OverrideActive bool            `json:"overrideActive,omitempty"` // Whether or not the origin has an active storage quota override
	UsageBreakdown []*UsageForType `json:"usageBreakdown,omitempty"` // Storage usage per type (bytes).
}

// Do executes Storage.getUsageAndQuota against the provided context.
//
// returns:
//
//	usage - Storage usage (bytes).
//	quota - Storage quota (bytes).
//	overrideActive - Whether or not the origin has an active storage quota override
//	usageBreakdown - Storage usage per type (bytes).
func (p *GetUsageAndQuotaParams) Do(ctx context.Context) (usage float64, quota float64, overrideActive bool, usageBreakdown []*UsageForType, err error) {
	// execute
	var res GetUsageAndQuotaReturns
	err = cdp.Execute(ctx, CommandGetUsageAndQuota, p, &res)
	if err != nil {
		return 0, 0, false, nil, err
	}

	return res.Usage, res.Quota, res.OverrideActive, res.UsageBreakdown, nil
}

// OverrideQuotaForOriginParams override quota for the specified origin.
type OverrideQuotaForOriginParams struct {
	Origin    string  `json:"origin"`              // Security origin.
	QuotaSize float64 `json:"quotaSize,omitempty"` // The quota size (in bytes) to override the original quota with. If this is called multiple times, the overridden quota will be equal to the quotaSize provided in the final call. If this is called without specifying a quotaSize, the quota will be reset to the default value for the specified origin. If this is called multiple times with different origins, the override will be maintained for each origin until it is disabled (called without a quotaSize).
}

// OverrideQuotaForOrigin override quota for the specified origin.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-overrideQuotaForOrigin
//
// parameters:
//
//	origin - Security origin.
func OverrideQuotaForOrigin(origin string) *OverrideQuotaForOriginParams {
	return &OverrideQuotaForOriginParams{
		Origin: origin,
	}
}

// WithQuotaSize the quota size (in bytes) to override the original quota
// with. If this is called multiple times, the overridden quota will be equal to
// the quotaSize provided in the final call. If this is called without
// specifying a quotaSize, the quota will be reset to the default value for the
// specified origin. If this is called multiple times with different origins,
// the override will be maintained for each origin until it is disabled (called
// without a quotaSize).
func (p OverrideQuotaForOriginParams) WithQuotaSize(quotaSize float64) *OverrideQuotaForOriginParams {
	p.QuotaSize = quotaSize
	return &p
}

// Do executes Storage.overrideQuotaForOrigin against the provided context.
func (p *OverrideQuotaForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandOverrideQuotaForOrigin, p, nil)
}

// TrackCacheStorageForOriginParams registers origin to be notified when an
// update occurs to its cache storage list.
type TrackCacheStorageForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// TrackCacheStorageForOrigin registers origin to be notified when an update
// occurs to its cache storage list.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackCacheStorageForOrigin
//
// parameters:
//
//	origin - Security origin.
func TrackCacheStorageForOrigin(origin string) *TrackCacheStorageForOriginParams {
	return &TrackCacheStorageForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.trackCacheStorageForOrigin against the provided context.
func (p *TrackCacheStorageForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackCacheStorageForOrigin, p, nil)
}

// TrackCacheStorageForStorageKeyParams registers storage key to be notified
// when an update occurs to its cache storage list.
type TrackCacheStorageForStorageKeyParams struct {
	StorageKey string `json:"storageKey"` // Storage key.
}

// TrackCacheStorageForStorageKey registers storage key to be notified when
// an update occurs to its cache storage list.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackCacheStorageForStorageKey
//
// parameters:
//
//	storageKey - Storage key.
func TrackCacheStorageForStorageKey(storageKey string) *TrackCacheStorageForStorageKeyParams {
	return &TrackCacheStorageForStorageKeyParams{
		StorageKey: storageKey,
	}
}

// Do executes Storage.trackCacheStorageForStorageKey against the provided context.
func (p *TrackCacheStorageForStorageKeyParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackCacheStorageForStorageKey, p, nil)
}

// TrackIndexedDBForOriginParams registers origin to be notified when an
// update occurs to its IndexedDB.
type TrackIndexedDBForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// TrackIndexedDBForOrigin registers origin to be notified when an update
// occurs to its IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackIndexedDBForOrigin
//
// parameters:
//
//	origin - Security origin.
func TrackIndexedDBForOrigin(origin string) *TrackIndexedDBForOriginParams {
	return &TrackIndexedDBForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.trackIndexedDBForOrigin against the provided context.
func (p *TrackIndexedDBForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackIndexedDBForOrigin, p, nil)
}

// TrackIndexedDBForStorageKeyParams registers storage key to be notified
// when an update occurs to its IndexedDB.
type TrackIndexedDBForStorageKeyParams struct {
	StorageKey string `json:"storageKey"` // Storage key.
}

// TrackIndexedDBForStorageKey registers storage key to be notified when an
// update occurs to its IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-trackIndexedDBForStorageKey
//
// parameters:
//
//	storageKey - Storage key.
func TrackIndexedDBForStorageKey(storageKey string) *TrackIndexedDBForStorageKeyParams {
	return &TrackIndexedDBForStorageKeyParams{
		StorageKey: storageKey,
	}
}

// Do executes Storage.trackIndexedDBForStorageKey against the provided context.
func (p *TrackIndexedDBForStorageKeyParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandTrackIndexedDBForStorageKey, p, nil)
}

// UntrackCacheStorageForOriginParams unregisters origin from receiving
// notifications for cache storage.
type UntrackCacheStorageForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// UntrackCacheStorageForOrigin unregisters origin from receiving
// notifications for cache storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackCacheStorageForOrigin
//
// parameters:
//
//	origin - Security origin.
func UntrackCacheStorageForOrigin(origin string) *UntrackCacheStorageForOriginParams {
	return &UntrackCacheStorageForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.untrackCacheStorageForOrigin against the provided context.
func (p *UntrackCacheStorageForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUntrackCacheStorageForOrigin, p, nil)
}

// UntrackCacheStorageForStorageKeyParams unregisters storage key from
// receiving notifications for cache storage.
type UntrackCacheStorageForStorageKeyParams struct {
	StorageKey string `json:"storageKey"` // Storage key.
}

// UntrackCacheStorageForStorageKey unregisters storage key from receiving
// notifications for cache storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackCacheStorageForStorageKey
//
// parameters:
//
//	storageKey - Storage key.
func UntrackCacheStorageForStorageKey(storageKey string) *UntrackCacheStorageForStorageKeyParams {
	return &UntrackCacheStorageForStorageKeyParams{
		StorageKey: storageKey,
	}
}

// Do executes Storage.untrackCacheStorageForStorageKey against the provided context.
func (p *UntrackCacheStorageForStorageKeyParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUntrackCacheStorageForStorageKey, p, nil)
}

// UntrackIndexedDBForOriginParams unregisters origin from receiving
// notifications for IndexedDB.
type UntrackIndexedDBForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// UntrackIndexedDBForOrigin unregisters origin from receiving notifications
// for IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackIndexedDBForOrigin
//
// parameters:
//
//	origin - Security origin.
func UntrackIndexedDBForOrigin(origin string) *UntrackIndexedDBForOriginParams {
	return &UntrackIndexedDBForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.untrackIndexedDBForOrigin against the provided context.
func (p *UntrackIndexedDBForOriginParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUntrackIndexedDBForOrigin, p, nil)
}

// UntrackIndexedDBForStorageKeyParams unregisters storage key from receiving
// notifications for IndexedDB.
type UntrackIndexedDBForStorageKeyParams struct {
	StorageKey string `json:"storageKey"` // Storage key.
}

// UntrackIndexedDBForStorageKey unregisters storage key from receiving
// notifications for IndexedDB.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-untrackIndexedDBForStorageKey
//
// parameters:
//
//	storageKey - Storage key.
func UntrackIndexedDBForStorageKey(storageKey string) *UntrackIndexedDBForStorageKeyParams {
	return &UntrackIndexedDBForStorageKeyParams{
		StorageKey: storageKey,
	}
}

// Do executes Storage.untrackIndexedDBForStorageKey against the provided context.
func (p *UntrackIndexedDBForStorageKeyParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandUntrackIndexedDBForStorageKey, p, nil)
}

// GetTrustTokensParams returns the number of stored Trust Tokens per issuer
// for the current browsing context.
type GetTrustTokensParams struct{}

// GetTrustTokens returns the number of stored Trust Tokens per issuer for
// the current browsing context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getTrustTokens
func GetTrustTokens() *GetTrustTokensParams {
	return &GetTrustTokensParams{}
}

// GetTrustTokensReturns return values.
type GetTrustTokensReturns struct {
	Tokens []*TrustTokens `json:"tokens,omitempty"`
}

// Do executes Storage.getTrustTokens against the provided context.
//
// returns:
//
//	tokens
func (p *GetTrustTokensParams) Do(ctx context.Context) (tokens []*TrustTokens, err error) {
	// execute
	var res GetTrustTokensReturns
	err = cdp.Execute(ctx, CommandGetTrustTokens, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Tokens, nil
}

// ClearTrustTokensParams removes all Trust Tokens issued by the provided
// issuerOrigin. Leaves other stored data, including the issuer's Redemption
// Records, intact.
type ClearTrustTokensParams struct {
	IssuerOrigin string `json:"issuerOrigin"`
}

// ClearTrustTokens removes all Trust Tokens issued by the provided
// issuerOrigin. Leaves other stored data, including the issuer's Redemption
// Records, intact.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearTrustTokens
//
// parameters:
//
//	issuerOrigin
func ClearTrustTokens(issuerOrigin string) *ClearTrustTokensParams {
	return &ClearTrustTokensParams{
		IssuerOrigin: issuerOrigin,
	}
}

// ClearTrustTokensReturns return values.
type ClearTrustTokensReturns struct {
	DidDeleteTokens bool `json:"didDeleteTokens,omitempty"` // True if any tokens were deleted, false otherwise.
}

// Do executes Storage.clearTrustTokens against the provided context.
//
// returns:
//
//	didDeleteTokens - True if any tokens were deleted, false otherwise.
func (p *ClearTrustTokensParams) Do(ctx context.Context) (didDeleteTokens bool, err error) {
	// execute
	var res ClearTrustTokensReturns
	err = cdp.Execute(ctx, CommandClearTrustTokens, p, &res)
	if err != nil {
		return false, err
	}

	return res.DidDeleteTokens, nil
}

// GetInterestGroupDetailsParams gets details for a named interest group.
type GetInterestGroupDetailsParams struct {
	OwnerOrigin string `json:"ownerOrigin"`
	Name        string `json:"name"`
}

// GetInterestGroupDetails gets details for a named interest group.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getInterestGroupDetails
//
// parameters:
//
//	ownerOrigin
//	name
func GetInterestGroupDetails(ownerOrigin string, name string) *GetInterestGroupDetailsParams {
	return &GetInterestGroupDetailsParams{
		OwnerOrigin: ownerOrigin,
		Name:        name,
	}
}

// GetInterestGroupDetailsReturns return values.
type GetInterestGroupDetailsReturns struct {
	Details *InterestGroupDetails `json:"details,omitempty"`
}

// Do executes Storage.getInterestGroupDetails against the provided context.
//
// returns:
//
//	details
func (p *GetInterestGroupDetailsParams) Do(ctx context.Context) (details *InterestGroupDetails, err error) {
	// execute
	var res GetInterestGroupDetailsReturns
	err = cdp.Execute(ctx, CommandGetInterestGroupDetails, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Details, nil
}

// SetInterestGroupTrackingParams enables/Disables issuing of
// interestGroupAccessed events.
type SetInterestGroupTrackingParams struct {
	Enable bool `json:"enable"`
}

// SetInterestGroupTracking enables/Disables issuing of interestGroupAccessed
// events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-setInterestGroupTracking
//
// parameters:
//
//	enable
func SetInterestGroupTracking(enable bool) *SetInterestGroupTrackingParams {
	return &SetInterestGroupTrackingParams{
		Enable: enable,
	}
}

// Do executes Storage.setInterestGroupTracking against the provided context.
func (p *SetInterestGroupTrackingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetInterestGroupTracking, p, nil)
}

// GetSharedStorageMetadataParams gets metadata for an origin's shared
// storage.
type GetSharedStorageMetadataParams struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

// GetSharedStorageMetadata gets metadata for an origin's shared storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getSharedStorageMetadata
//
// parameters:
//
//	ownerOrigin
func GetSharedStorageMetadata(ownerOrigin string) *GetSharedStorageMetadataParams {
	return &GetSharedStorageMetadataParams{
		OwnerOrigin: ownerOrigin,
	}
}

// GetSharedStorageMetadataReturns return values.
type GetSharedStorageMetadataReturns struct {
	Metadata *SharedStorageMetadata `json:"metadata,omitempty"`
}

// Do executes Storage.getSharedStorageMetadata against the provided context.
//
// returns:
//
//	metadata
func (p *GetSharedStorageMetadataParams) Do(ctx context.Context) (metadata *SharedStorageMetadata, err error) {
	// execute
	var res GetSharedStorageMetadataReturns
	err = cdp.Execute(ctx, CommandGetSharedStorageMetadata, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Metadata, nil
}

// GetSharedStorageEntriesParams gets the entries in an given origin's shared
// storage.
type GetSharedStorageEntriesParams struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

// GetSharedStorageEntries gets the entries in an given origin's shared
// storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-getSharedStorageEntries
//
// parameters:
//
//	ownerOrigin
func GetSharedStorageEntries(ownerOrigin string) *GetSharedStorageEntriesParams {
	return &GetSharedStorageEntriesParams{
		OwnerOrigin: ownerOrigin,
	}
}

// GetSharedStorageEntriesReturns return values.
type GetSharedStorageEntriesReturns struct {
	Entries []*SharedStorageEntry `json:"entries,omitempty"`
}

// Do executes Storage.getSharedStorageEntries against the provided context.
//
// returns:
//
//	entries
func (p *GetSharedStorageEntriesParams) Do(ctx context.Context) (entries []*SharedStorageEntry, err error) {
	// execute
	var res GetSharedStorageEntriesReturns
	err = cdp.Execute(ctx, CommandGetSharedStorageEntries, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Entries, nil
}

// SetSharedStorageEntryParams sets entry with key and value for a given
// origin's shared storage.
type SetSharedStorageEntryParams struct {
	OwnerOrigin     string `json:"ownerOrigin"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	IgnoreIfPresent bool   `json:"ignoreIfPresent,omitempty"` // If ignoreIfPresent is included and true, then only sets the entry if key doesn't already exist.
}

// SetSharedStorageEntry sets entry with key and value for a given origin's
// shared storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-setSharedStorageEntry
//
// parameters:
//
//	ownerOrigin
//	key
//	value
func SetSharedStorageEntry(ownerOrigin string, key string, value string) *SetSharedStorageEntryParams {
	return &SetSharedStorageEntryParams{
		OwnerOrigin: ownerOrigin,
		Key:         key,
		Value:       value,
	}
}

// WithIgnoreIfPresent if ignoreIfPresent is included and true, then only
// sets the entry if key doesn't already exist.
func (p SetSharedStorageEntryParams) WithIgnoreIfPresent(ignoreIfPresent bool) *SetSharedStorageEntryParams {
	p.IgnoreIfPresent = ignoreIfPresent
	return &p
}

// Do executes Storage.setSharedStorageEntry against the provided context.
func (p *SetSharedStorageEntryParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetSharedStorageEntry, p, nil)
}

// DeleteSharedStorageEntryParams deletes entry for key (if it exists) for a
// given origin's shared storage.
type DeleteSharedStorageEntryParams struct {
	OwnerOrigin string `json:"ownerOrigin"`
	Key         string `json:"key"`
}

// DeleteSharedStorageEntry deletes entry for key (if it exists) for a given
// origin's shared storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-deleteSharedStorageEntry
//
// parameters:
//
//	ownerOrigin
//	key
func DeleteSharedStorageEntry(ownerOrigin string, key string) *DeleteSharedStorageEntryParams {
	return &DeleteSharedStorageEntryParams{
		OwnerOrigin: ownerOrigin,
		Key:         key,
	}
}

// Do executes Storage.deleteSharedStorageEntry against the provided context.
func (p *DeleteSharedStorageEntryParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDeleteSharedStorageEntry, p, nil)
}

// ClearSharedStorageEntriesParams clears all entries for a given origin's
// shared storage.
type ClearSharedStorageEntriesParams struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

// ClearSharedStorageEntries clears all entries for a given origin's shared
// storage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-clearSharedStorageEntries
//
// parameters:
//
//	ownerOrigin
func ClearSharedStorageEntries(ownerOrigin string) *ClearSharedStorageEntriesParams {
	return &ClearSharedStorageEntriesParams{
		OwnerOrigin: ownerOrigin,
	}
}

// Do executes Storage.clearSharedStorageEntries against the provided context.
func (p *ClearSharedStorageEntriesParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearSharedStorageEntries, p, nil)
}

// ResetSharedStorageBudgetParams resets the budget for ownerOrigin by
// clearing all budget withdrawals.
type ResetSharedStorageBudgetParams struct {
	OwnerOrigin string `json:"ownerOrigin"`
}

// ResetSharedStorageBudget resets the budget for ownerOrigin by clearing all
// budget withdrawals.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-resetSharedStorageBudget
//
// parameters:
//
//	ownerOrigin
func ResetSharedStorageBudget(ownerOrigin string) *ResetSharedStorageBudgetParams {
	return &ResetSharedStorageBudgetParams{
		OwnerOrigin: ownerOrigin,
	}
}

// Do executes Storage.resetSharedStorageBudget against the provided context.
func (p *ResetSharedStorageBudgetParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandResetSharedStorageBudget, p, nil)
}

// SetSharedStorageTrackingParams enables/disables issuing of
// sharedStorageAccessed events.
type SetSharedStorageTrackingParams struct {
	Enable bool `json:"enable"`
}

// SetSharedStorageTracking enables/disables issuing of sharedStorageAccessed
// events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-setSharedStorageTracking
//
// parameters:
//
//	enable
func SetSharedStorageTracking(enable bool) *SetSharedStorageTrackingParams {
	return &SetSharedStorageTrackingParams{
		Enable: enable,
	}
}

// Do executes Storage.setSharedStorageTracking against the provided context.
func (p *SetSharedStorageTrackingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetSharedStorageTracking, p, nil)
}

// SetStorageBucketTrackingParams set tracking for a storage key's buckets.
type SetStorageBucketTrackingParams struct {
	StorageKey string `json:"storageKey"`
	Enable     bool   `json:"enable"`
}

// SetStorageBucketTracking set tracking for a storage key's buckets.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-setStorageBucketTracking
//
// parameters:
//
//	storageKey
//	enable
func SetStorageBucketTracking(storageKey string, enable bool) *SetStorageBucketTrackingParams {
	return &SetStorageBucketTrackingParams{
		StorageKey: storageKey,
		Enable:     enable,
	}
}

// Do executes Storage.setStorageBucketTracking against the provided context.
func (p *SetStorageBucketTrackingParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetStorageBucketTracking, p, nil)
}

// DeleteStorageBucketParams deletes the Storage Bucket with the given
// storage key and bucket name.
type DeleteStorageBucketParams struct {
	StorageKey string `json:"storageKey"`
	BucketName string `json:"bucketName"`
}

// DeleteStorageBucket deletes the Storage Bucket with the given storage key
// and bucket name.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#method-deleteStorageBucket
//
// parameters:
//
//	storageKey
//	bucketName
func DeleteStorageBucket(storageKey string, bucketName string) *DeleteStorageBucketParams {
	return &DeleteStorageBucketParams{
		StorageKey: storageKey,
		BucketName: bucketName,
	}
}

// Do executes Storage.deleteStorageBucket against the provided context.
func (p *DeleteStorageBucketParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandDeleteStorageBucket, p, nil)
}

// Command names.
const (
	CommandGetStorageKeyForFrame            = "Storage.getStorageKeyForFrame"
	CommandClearDataForOrigin               = "Storage.clearDataForOrigin"
	CommandClearDataForStorageKey           = "Storage.clearDataForStorageKey"
	CommandGetCookies                       = "Storage.getCookies"
	CommandSetCookies                       = "Storage.setCookies"
	CommandClearCookies                     = "Storage.clearCookies"
	CommandGetUsageAndQuota                 = "Storage.getUsageAndQuota"
	CommandOverrideQuotaForOrigin           = "Storage.overrideQuotaForOrigin"
	CommandTrackCacheStorageForOrigin       = "Storage.trackCacheStorageForOrigin"
	CommandTrackCacheStorageForStorageKey   = "Storage.trackCacheStorageForStorageKey"
	CommandTrackIndexedDBForOrigin          = "Storage.trackIndexedDBForOrigin"
	CommandTrackIndexedDBForStorageKey      = "Storage.trackIndexedDBForStorageKey"
	CommandUntrackCacheStorageForOrigin     = "Storage.untrackCacheStorageForOrigin"
	CommandUntrackCacheStorageForStorageKey = "Storage.untrackCacheStorageForStorageKey"
	CommandUntrackIndexedDBForOrigin        = "Storage.untrackIndexedDBForOrigin"
	CommandUntrackIndexedDBForStorageKey    = "Storage.untrackIndexedDBForStorageKey"
	CommandGetTrustTokens                   = "Storage.getTrustTokens"
	CommandClearTrustTokens                 = "Storage.clearTrustTokens"
	CommandGetInterestGroupDetails          = "Storage.getInterestGroupDetails"
	CommandSetInterestGroupTracking         = "Storage.setInterestGroupTracking"
	CommandGetSharedStorageMetadata         = "Storage.getSharedStorageMetadata"
	CommandGetSharedStorageEntries          = "Storage.getSharedStorageEntries"
	CommandSetSharedStorageEntry            = "Storage.setSharedStorageEntry"
	CommandDeleteSharedStorageEntry         = "Storage.deleteSharedStorageEntry"
	CommandClearSharedStorageEntries        = "Storage.clearSharedStorageEntries"
	CommandResetSharedStorageBudget         = "Storage.resetSharedStorageBudget"
	CommandSetSharedStorageTracking         = "Storage.setSharedStorageTracking"
	CommandSetStorageBucketTracking         = "Storage.setStorageBucketTracking"
	CommandDeleteStorageBucket              = "Storage.deleteStorageBucket"
)
