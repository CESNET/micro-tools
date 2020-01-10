package oidc

// The ProviderMetadata describes an idp.
// see https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata
type ProviderMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`
	//claims_parameter_supported
	ClaimsSupported []string `json:"claims_supported,omitempty"`
	//grant_types_supported
	IDTokenSigningAlgValuesSupported []string `json:"id_token_signing_alg_values_supported,omitempty"`
	Issuer                           string   `json:"issuer,omitempty"`
	JwksURI                          string   `json:"jwks_uri,omitempty"`
	//registration_endpoint
	//request_object_signing_alg_values_supported
	//request_parameter_supported
	//request_uri_parameter_supported
	//require_request_uri_registration
	//response_modes_supported
	ResponseTypesSupported []string `json:"response_types_supported,omitempty"`
	ScopesSupported        []string `json:"scopes_supported,omitempty"`
	SubjectTypesSupported  []string `json:"subject_types_supported,omitempty"`
	TokenEndpoint          string   `json:"token_endpoint,omitempty"`
	//token_endpoint_auth_methods_supported
	//token_endpoint_auth_signing_alg_values_supported
	UserinfoEndpoint string `json:"userinfo_endpoint,omitempty"`
	//userinfo_signing_alg_values_supported
	//code_challenge_methods_supported
	IntrospectionEndpoint string `json:"introspection_endpoint,omitempty"`
	//introspection_endpoint_auth_methods_supported
	//introspection_endpoint_auth_signing_alg_values_supported
	RevocationEndpoint string `json:"revocation_endpoint,omitempty"`
	//revocation_endpoint_auth_methods_supported
	//revocation_endpoint_auth_signing_alg_values_supported
	//id_token_encryption_alg_values_supported
	//id_token_encryption_enc_values_supported
	//userinfo_encryption_alg_values_supported
	//userinfo_encryption_enc_values_supported
	//request_object_encryption_alg_values_supported
	//request_object_encryption_enc_values_supported
	CheckSessionIframe string `json:"check_session_iframe,omitempty"`
	EndSessionEndpoint string `json:"end_session_endpoint,omitempty"`
	//claim_types_supported
}

// StandardClaims will be stored in the context to be consumed by the oidc user manager
// They can be requested to be returned either in the UserInfo Response, per
// Section 5.3.2, or in the ID Token, per Section 2.
// see https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
type StandardClaims struct {
	Iss string `json:"iss"`
	// Subject - Identifier for the End-User at the Issuer.
	Sub string `json:"sub,omitempty"`

	// End-User's full name in displayable form including all name parts, possibly
	// including titles and suffixes, ordered according to the End-User's locale
	// and preferences.
	Name string `json:"name,omitempty"`

	// Given name(s) or first name(s) of the End-User. Note that in some cultures,
	// people can have multiple given names; all can be present, with the names
	// being separated by space characters.
	GivenName string `json:"given_name,omitempty"`

	// Surname(s) or last name(s) of the End-User. Note that in some cultures,
	// people can have multiple family names or no family name; all can be present,
	// with the names being separated by space characters.
	FamilyName string `json:"family_name,omitempty"`

	// Middle name(s) of the End-User. Note that in some cultures, people can have
	// multiple middle names; all can be present, with the names being separated by
	// space characters. Also note that in some cultures, middle names are not used.
	MiddleName string `json:"middle_name,omitempty"`

	// Shorthand name by which the End-User wishes to be referred to at the RP, such
	// as janedoe or j.doe. This value MAY be any valid JSON string including special
	// characters such as @, /, or whitespace. The RP MUST NOT rely upon this value
	// being unique, as discussed in Section 5.7.
	PreferredUsername string `json:"preferred_username,omitempty"`

	// End-User's preferred e-mail address. Its value MUST conform to the RFC 5322
	// addr-spec syntax. The RP MUST NOT rely upon this value being unique, as
	// discussed in Section 5.7.
	Email string `json:"email,omitempty"`

	// String from zoneinfo time zone database representing the End-User's time
	// zone. For example, Europe/Paris or America/Los_Angeles.
	Zoneinfo string `json:"zoneinfo,omitempty"`

	// End-User's locale, represented as a BCP47 [RFC5646] language tag.
	// This is typically an ISO 639-1 Alpha-2 [ISO639‑1] language code in
	// lowercase and an ISO 3166-1 Alpha-2 [ISO3166‑1] country code in
	// uppercase, separated by a dash. For example, en-US or fr-CA. As a
	// compatibility note, some implementations have used an underscore as
	// the separator rather than a dash, for example, en_US; Relying Parties
	// MAY choose to accept this locale syntax as well.
	Locale string `json:"locale,omitempty"`

	// End-User's preferred telephone number. E.164 [E.164] is RECOMMENDED
	// as the format of this Claim, for example, +1 (425) 555-1212 or
	// +56 (2) 687 2400. If the phone number contains an extension, it is
	// RECOMMENDED that the extension be represented using the RFC 3966
	// extension syntax, for example, +1 (604) 555-1234;ext=5678.
	PhoneNumber string `json:"phone_number,omitempty"`

	Groups       []string `json:"eduPersonEntitlement,omitempty"`
	LastEligible string   `json:"isCesnetEligibleLastSeen,omitempty"`
	Organization string   `json:"organization,omitempty"`
}
