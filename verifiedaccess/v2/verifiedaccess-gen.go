// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package verifiedaccess provides access to the Chrome Verified Access API.
//
// For product documentation, see: https://developers.google.com/chrome/verified-access
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/verifiedaccess/v2"
//	...
//	ctx := context.Background()
//	verifiedaccessService, err := verifiedaccess.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	verifiedaccessService, err := verifiedaccess.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	verifiedaccessService, err := verifiedaccess.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package verifiedaccess // import "google.golang.org/api/verifiedaccess/v2"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "verifiedaccess:v2"
const apiName = "verifiedaccess"
const apiVersion = "v2"
const basePath = "https://verifiedaccess.googleapis.com/"
const mtlsBasePath = "https://verifiedaccess.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Verify your enterprise credentials
	VerifiedaccessScope = "https://www.googleapis.com/auth/verifiedaccess"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/verifiedaccess",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Challenge = NewChallengeService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Challenge *ChallengeService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewChallengeService(s *Service) *ChallengeService {
	rs := &ChallengeService{s: s}
	return rs
}

type ChallengeService struct {
	s *Service
}

// Challenge: Result message for VerifiedAccess.GenerateChallenge.
type Challenge struct {
	// Challenge: Generated challenge, the bytes representation of
	// SignedData.
	Challenge string `json:"challenge,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Challenge") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Challenge") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Challenge) MarshalJSON() ([]byte, error) {
	type NoMethod Challenge
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CrowdStrikeAgent: Properties of the CrowdStrike agent installed on a
// device.
type CrowdStrikeAgent struct {
	// AgentId: The Agent ID of the Crowdstrike agent.
	AgentId string `json:"agentId,omitempty"`

	// CustomerId: The Customer ID to which the agent belongs to.
	CustomerId string `json:"customerId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AgentId") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AgentId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CrowdStrikeAgent) MarshalJSON() ([]byte, error) {
	type NoMethod CrowdStrikeAgent
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DeviceSignals: The device signals as reported by Chrome. Unless
// otherwise specified, signals are available on all platforms.
type DeviceSignals struct {
	// AllowScreenLock: Value of the AllowScreenLock policy on the device.
	// See https://chromeenterprise.google/policies/?policy=AllowScreenLock
	// for more details. Available on ChromeOS only.
	AllowScreenLock bool `json:"allowScreenLock,omitempty"`

	// BrowserVersion: Current version of the Chrome browser which generated
	// this set of signals. Example value: "107.0.5286.0".
	BrowserVersion string `json:"browserVersion,omitempty"`

	// BuiltInDnsClientEnabled: Whether Chrome's built-in DNS client is
	// used. The OS DNS client is otherwise used. This value may be
	// controlled by an enterprise policy:
	// https://chromeenterprise.google/policies/#BuiltInDnsClientEnabled.
	BuiltInDnsClientEnabled bool `json:"builtInDnsClientEnabled,omitempty"`

	// ChromeRemoteDesktopAppBlocked: Whether access to the Chrome Remote
	// Desktop application is blocked via a policy.
	ChromeRemoteDesktopAppBlocked bool `json:"chromeRemoteDesktopAppBlocked,omitempty"`

	// CrowdStrikeAgent: Crowdstrike agent properties installed on the
	// device, if any. Available on Windows and MacOS only.
	CrowdStrikeAgent *CrowdStrikeAgent `json:"crowdStrikeAgent,omitempty"`

	// DeviceAffiliationIds: Affiliation IDs of the organizations that are
	// affiliated with the organization that is currently managing the
	// device. When the sets of device and profile affiliation IDs overlap,
	// it means that the organizations managing the device and user are
	// affiliated. To learn more about user affiliation, visit
	// https://support.google.com/chrome/a/answer/12801245?ref_topic=9027936.
	DeviceAffiliationIds []string `json:"deviceAffiliationIds,omitempty"`

	// DeviceEnrollmentDomain: Enrollment domain of the customer which is
	// currently managing the device.
	DeviceEnrollmentDomain string `json:"deviceEnrollmentDomain,omitempty"`

	// DeviceManufacturer: The name of the device's manufacturer.
	DeviceManufacturer string `json:"deviceManufacturer,omitempty"`

	// DeviceModel: The name of the device's model.
	DeviceModel string `json:"deviceModel,omitempty"`

	// DiskEncryption: The encryption state of the disk. On ChromeOS, the
	// main disk is always ENCRYPTED.
	//
	// Possible values:
	//   "DISK_ENCRYPTION_UNSPECIFIED" - Unspecified.
	//   "DISK_ENCRYPTION_UNKNOWN" - Chrome could not evaluate the
	// encryption state.
	//   "DISK_ENCRYPTION_DISABLED" - The main disk is not encrypted.
	//   "DISK_ENCRYPTION_ENCRYPTED" - The main disk is encrypted.
	DiskEncryption string `json:"diskEncryption,omitempty"`

	// DisplayName: The display name of the device, as defined by the user.
	DisplayName string `json:"displayName,omitempty"`

	// Hostname: Hostname of the device.
	Hostname string `json:"hostname,omitempty"`

	// Imei: International Mobile Equipment Identity (IMEI) of the device.
	// Available on ChromeOS only.
	Imei []string `json:"imei,omitempty"`

	// MacAddresses: MAC addresses of the device.
	MacAddresses []string `json:"macAddresses,omitempty"`

	// Meid: Mobile Equipment Identifier (MEID) of the device. Available on
	// ChromeOS only.
	Meid []string `json:"meid,omitempty"`

	// OperatingSystem: The type of the Operating System currently running
	// on the device.
	//
	// Possible values:
	//   "OPERATING_SYSTEM_UNSPECIFIED" - UNSPECIFIED.
	//   "CHROME_OS" - ChromeOS.
	//   "CHROMIUM_OS" - ChromiumOS.
	//   "WINDOWS" - Windows.
	//   "MAC_OS_X" - Mac Os X.
	//   "LINUX" - Linux
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// OsFirewall: The state of the OS level firewall. On ChromeOS, the
	// value will always be ENABLED on regular devices and UNKNOWN on
	// devices in developer mode.
	//
	// Possible values:
	//   "OS_FIREWALL_UNSPECIFIED" - Unspecified.
	//   "OS_FIREWALL_UNKNOWN" - Chrome could not evaluate the OS firewall
	// state.
	//   "OS_FIREWALL_DISABLED" - The OS firewall is disabled.
	//   "OS_FIREWALL_ENABLED" - The OS firewall is enabled.
	OsFirewall string `json:"osFirewall,omitempty"`

	// OsVersion: The current version of the Operating System. On Windows
	// and linux, the value will also include the security patch
	// information.
	OsVersion string `json:"osVersion,omitempty"`

	// PasswordProtectionWarningTrigger: Whether the Password Protection
	// Warning feature is enabled or not. Password protection alerts users
	// when they reuse their protected password on potentially suspicious
	// sites. This setting is controlled by an enterprise policy:
	// https://chromeenterprise.google/policies/#PasswordProtectionWarningTrigger.
	// Note that the policy unset does not have the same effects as having
	// the policy explicitly set to `PASSWORD_PROTECTION_OFF`.
	//
	// Possible values:
	//   "PASSWORD_PROTECTION_WARNING_TRIGGER_UNSPECIFIED" - Unspecified.
	//   "POLICY_UNSET" - The policy is not set.
	//   "PASSWORD_PROTECTION_OFF" - No password protection warning will be
	// shown.
	//   "PASSWORD_REUSE" - Password protection warning is shown if a
	// protected password is re-used.
	//   "PHISHING_REUSE" - Password protection warning is shown if a
	// protected password is re-used on a known phishing website.
	PasswordProtectionWarningTrigger string `json:"passwordProtectionWarningTrigger,omitempty"`

	// ProfileAffiliationIds: Affiliation IDs of the organizations that are
	// affiliated with the organization that is currently managing the
	// Chrome Profile’s user or ChromeOS user.
	ProfileAffiliationIds []string `json:"profileAffiliationIds,omitempty"`

	// RealtimeUrlCheckMode: Whether Enterprise-grade (i.e. custom) unsafe
	// URL scanning is enabled or not. This setting may be controlled by an
	// enterprise policy:
	// https://chromeenterprise.google/policies/#EnterpriseRealTimeUrlCheckMode
	//
	// Possible values:
	//   "REALTIME_URL_CHECK_MODE_UNSPECIFIED" - Unspecified.
	//   "REALTIME_URL_CHECK_MODE_DISABLED" - Disabled. Consumer Safe
	// Browsing checks are applied.
	//   "REALTIME_URL_CHECK_MODE_ENABLED_MAIN_FRAME" - Realtime check for
	// main frame URLs is enabled.
	RealtimeUrlCheckMode string `json:"realtimeUrlCheckMode,omitempty"`

	// SafeBrowsingProtectionLevel: Safe Browsing Protection Level. That
	// setting may be controlled by an enterprise policy:
	// https://chromeenterprise.google/policies/#SafeBrowsingProtectionLevel.
	//
	// Possible values:
	//   "SAFE_BROWSING_PROTECTION_LEVEL_UNSPECIFIED" - Unspecified.
	//   "INACTIVE" - Safe Browsing is disabled.
	//   "STANDARD" - Safe Browsing is active in the standard mode.
	//   "ENHANCED" - Safe Browsing is active in the enhanced mode.
	SafeBrowsingProtectionLevel string `json:"safeBrowsingProtectionLevel,omitempty"`

	// ScreenLockSecured: The state of the Screen Lock password protection.
	// On ChromeOS, this value will always be ENABLED as there is not way to
	// disable requiring a password or pin when unlocking the device.
	//
	// Possible values:
	//   "SCREEN_LOCK_SECURED_UNSPECIFIED" - Unspecified.
	//   "SCREEN_LOCK_SECURED_UNKNOWN" - Chrome could not evaluate the state
	// of the Screen Lock mechanism.
	//   "SCREEN_LOCK_SECURED_DISABLED" - The Screen Lock is not
	// password-protected.
	//   "SCREEN_LOCK_SECURED_ENABLED" - The Screen Lock is
	// password-protected.
	ScreenLockSecured string `json:"screenLockSecured,omitempty"`

	// SecureBootMode: Whether the device's startup software has its Secure
	// Boot feature enabled. Available on Windows only.
	//
	// Possible values:
	//   "SECURE_BOOT_MODE_UNSPECIFIED" - Unspecified.
	//   "SECURE_BOOT_MODE_UNKNOWN" - Chrome was unable to determine the
	// Secure Boot mode.
	//   "SECURE_BOOT_MODE_DISABLED" - Secure Boot was disabled on the
	// startup software.
	//   "SECURE_BOOT_MODE_ENABLED" - Secure Boot was enabled on the startup
	// software.
	SecureBootMode string `json:"secureBootMode,omitempty"`

	// SerialNumber: The serial number of the device. On Windows, this
	// represents the BIOS's serial number. Not available on most Linux
	// distributions.
	SerialNumber string `json:"serialNumber,omitempty"`

	// SiteIsolationEnabled: Whether the Site Isolation (a.k.a Site Per
	// Process) setting is enabled. That setting may be controlled by an
	// enterprise policy:
	// https://chromeenterprise.google/policies/#SitePerProcess
	SiteIsolationEnabled bool `json:"siteIsolationEnabled,omitempty"`

	// SystemDnsServers: List of the addesses of all OS level DNS servers
	// configured in the device's network settings.
	SystemDnsServers []string `json:"systemDnsServers,omitempty"`

	// ThirdPartyBlockingEnabled: Whether Chrome is blocking third-party
	// software injection or not. This setting may be controlled by an
	// enterprise policy:
	// https://chromeenterprise.google/policies/?policy=ThirdPartyBlockingEnabled.
	// Available on Windows only.
	ThirdPartyBlockingEnabled bool `json:"thirdPartyBlockingEnabled,omitempty"`

	// Trigger: The trigger which generated this set of signals.
	//
	// Possible values:
	//   "TRIGGER_UNSPECIFIED" - Unspecified.
	//   "TRIGGER_BROWSER_NAVIGATION" - When navigating to an URL inside a
	// browser.
	//   "TRIGGER_LOGIN_SCREEN" - When signing into an account on the
	// ChromeOS login screen.
	Trigger string `json:"trigger,omitempty"`

	// WindowsMachineDomain: Windows domain that the current machine has
	// joined. Available on Windows only.
	WindowsMachineDomain string `json:"windowsMachineDomain,omitempty"`

	// WindowsUserDomain: Windows domain for the current OS user. Available
	// on Windows only.
	WindowsUserDomain string `json:"windowsUserDomain,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowScreenLock") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllowScreenLock") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *DeviceSignals) MarshalJSON() ([]byte, error) {
	type NoMethod DeviceSignals
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance: service Foo { rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty); }
type Empty struct {
}

// VerifyChallengeResponseRequest: Signed ChallengeResponse.
type VerifyChallengeResponseRequest struct {
	// ChallengeResponse: Required. The generated response to the challenge,
	// the bytes representation of SignedData.
	ChallengeResponse string `json:"challengeResponse,omitempty"`

	// ExpectedIdentity: Optional. Service can optionally provide identity
	// information about the device or user associated with the key. For an
	// EMK, this value is the enrolled domain. For an EUK, this value is the
	// user's email address. If present, this value will be checked against
	// contents of the response, and verification will fail if there is no
	// match.
	ExpectedIdentity string `json:"expectedIdentity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ChallengeResponse")
	// to unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ChallengeResponse") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *VerifyChallengeResponseRequest) MarshalJSON() ([]byte, error) {
	type NoMethod VerifyChallengeResponseRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VerifyChallengeResponseResult: Result message for
// VerifiedAccess.VerifyChallengeResponse.
type VerifyChallengeResponseResult struct {
	// AttestedDeviceId: Attested device ID (ADID).
	AttestedDeviceId string `json:"attestedDeviceId,omitempty"`

	// CustomerId: Unique customer id that this device belongs to, as
	// defined by the Google Admin SDK at
	// https://developers.google.com/admin-sdk/directory/v1/guides/manage-customers
	CustomerId string `json:"customerId,omitempty"`

	// DeviceEnrollmentId: Device enrollment id for ChromeOS devices.
	DeviceEnrollmentId string `json:"deviceEnrollmentId,omitempty"`

	// DevicePermanentId: Device permanent id is returned in this field (for
	// the machine response only).
	DevicePermanentId string `json:"devicePermanentId,omitempty"`

	// DeviceSignal: Deprecated. Device signal in json string
	// representation. Prefer using `device_signals` instead.
	DeviceSignal string `json:"deviceSignal,omitempty"`

	// DeviceSignals: Device signals.
	DeviceSignals *DeviceSignals `json:"deviceSignals,omitempty"`

	// KeyTrustLevel: Device attested key trust level.
	//
	// Possible values:
	//   "KEY_TRUST_LEVEL_UNSPECIFIED" - UNSPECIFIED.
	//   "CHROME_OS_VERIFIED_MODE" - ChromeOS device in verified mode.
	//   "CHROME_OS_DEVELOPER_MODE" - ChromeOS device in developer mode.
	//   "CHROME_BROWSER_HW_KEY" - Chrome Browser with the key stored in the
	// device hardware.
	//   "CHROME_BROWSER_OS_KEY" - Chrome Browser with the key stored at OS
	// level.
	//   "CHROME_BROWSER_NO_KEY" - Chrome Browser without an attestation
	// key.
	KeyTrustLevel string `json:"keyTrustLevel,omitempty"`

	// ProfileCustomerId: Unique customer id that this profile belongs to,
	// as defined by the Google Admin SDK at
	// https://developers.google.com/admin-sdk/directory/v1/guides/manage-customers
	ProfileCustomerId string `json:"profileCustomerId,omitempty"`

	// ProfileKeyTrustLevel: Profile attested key trust level.
	//
	// Possible values:
	//   "KEY_TRUST_LEVEL_UNSPECIFIED" - UNSPECIFIED.
	//   "CHROME_OS_VERIFIED_MODE" - ChromeOS device in verified mode.
	//   "CHROME_OS_DEVELOPER_MODE" - ChromeOS device in developer mode.
	//   "CHROME_BROWSER_HW_KEY" - Chrome Browser with the key stored in the
	// device hardware.
	//   "CHROME_BROWSER_OS_KEY" - Chrome Browser with the key stored at OS
	// level.
	//   "CHROME_BROWSER_NO_KEY" - Chrome Browser without an attestation
	// key.
	ProfileKeyTrustLevel string `json:"profileKeyTrustLevel,omitempty"`

	// SignedPublicKeyAndChallenge: Certificate Signing Request (in the
	// SPKAC format, base64 encoded) is returned in this field. This field
	// will be set only if device has included CSR in its challenge
	// response. (the option to include CSR is now available for both user
	// and machine responses)
	SignedPublicKeyAndChallenge string `json:"signedPublicKeyAndChallenge,omitempty"`

	// VirtualDeviceId: Virtual device id of the device. The definition of
	// virtual device id is platform-specific.
	VirtualDeviceId string `json:"virtualDeviceId,omitempty"`

	// VirtualProfileId: The ID of a profile on the device.
	VirtualProfileId string `json:"virtualProfileId,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AttestedDeviceId") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttestedDeviceId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *VerifyChallengeResponseResult) MarshalJSON() ([]byte, error) {
	type NoMethod VerifyChallengeResponseResult
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "verifiedaccess.challenge.generate":

type ChallengeGenerateCall struct {
	s          *Service
	empty      *Empty
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Generate: Generates a new challenge.
func (r *ChallengeService) Generate(empty *Empty) *ChallengeGenerateCall {
	c := &ChallengeGenerateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.empty = empty
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChallengeGenerateCall) Fields(s ...googleapi.Field) *ChallengeGenerateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChallengeGenerateCall) Context(ctx context.Context) *ChallengeGenerateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChallengeGenerateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChallengeGenerateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/"+internal.Version)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.empty)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/challenge:generate")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "verifiedaccess.challenge.generate" call.
// Exactly one of *Challenge or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Challenge.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ChallengeGenerateCall) Do(opts ...googleapi.CallOption) (*Challenge, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Challenge{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Generates a new challenge.",
	//   "flatPath": "v2/challenge:generate",
	//   "httpMethod": "POST",
	//   "id": "verifiedaccess.challenge.generate",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/challenge:generate",
	//   "request": {
	//     "$ref": "Empty"
	//   },
	//   "response": {
	//     "$ref": "Challenge"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/verifiedaccess"
	//   ]
	// }

}

// method id "verifiedaccess.challenge.verify":

type ChallengeVerifyCall struct {
	s                              *Service
	verifychallengeresponserequest *VerifyChallengeResponseRequest
	urlParams_                     gensupport.URLParams
	ctx_                           context.Context
	header_                        http.Header
}

// Verify: Verifies the challenge response.
func (r *ChallengeService) Verify(verifychallengeresponserequest *VerifyChallengeResponseRequest) *ChallengeVerifyCall {
	c := &ChallengeVerifyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.verifychallengeresponserequest = verifychallengeresponserequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ChallengeVerifyCall) Fields(s ...googleapi.Field) *ChallengeVerifyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ChallengeVerifyCall) Context(ctx context.Context) *ChallengeVerifyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ChallengeVerifyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ChallengeVerifyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/"+internal.Version)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.verifychallengeresponserequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/challenge:verify")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "verifiedaccess.challenge.verify" call.
// Exactly one of *VerifyChallengeResponseResult or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *VerifyChallengeResponseResult.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ChallengeVerifyCall) Do(opts ...googleapi.CallOption) (*VerifyChallengeResponseResult, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &VerifyChallengeResponseResult{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Verifies the challenge response.",
	//   "flatPath": "v2/challenge:verify",
	//   "httpMethod": "POST",
	//   "id": "verifiedaccess.challenge.verify",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/challenge:verify",
	//   "request": {
	//     "$ref": "VerifyChallengeResponseRequest"
	//   },
	//   "response": {
	//     "$ref": "VerifyChallengeResponseResult"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/verifiedaccess"
	//   ]
	// }

}
