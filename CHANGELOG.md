# Changelog

### 0.5.0 (2026-04-23)

* [PLT-2583] Bump oauth2-proxy upstream to v7.15.2 (security fixes: authentication bypasses CVE-2026-34986, CVE-2026-32281 and others)

## Previous development

### Branched to branch-0.4 (2026-02-25)

* [PLT-3428] Bump oauth2 proxy upstream version to 7.13.0
* [PLT-2291] Fix: Handle missing JWT cookie on oauth2-proxy logout
* [PLT-2674] Bump oauth2 proxy upstream version to 7.12.0
* [PLT-3777] Feature: Adaptate to Semantic Versioning

* Oauth2 proxy upstream version: 7.13.0

#

<!--
## Changes since v7.15.2

# V7.15.2

## Release Highlights

- 🔵 Golang version upgrade to v1.25.9
    - Upgrade of all dependencies to their latest versions
    - [CVE-2026-34986](https://nvd.nist.gov/vuln/detail/CVE-2026-34986)
    - [CVE-2026-32281](https://nvd.nist.gov/vuln/detail/CVE-2026-32281)
    - [CVE-2026-32289](https://nvd.nist.gov/vuln/detail/CVE-2026-32289)
    - [CVE-2026-32288](https://nvd.nist.gov/vuln/detail/CVE-2026-32288)
    - [CVE-2026-32280](https://nvd.nist.gov/vuln/detail/CVE-2026-32280)
    - [CVE-2026-32282](https://nvd.nist.gov/vuln/detail/CVE-2026-32282)
    - [CVE-2026-32283](https://nvd.nist.gov/vuln/detail/CVE-2026-32283)
-  🕵️‍♀️ Vulnerabilities have been addressed

## Important Notes

We have had security audits performed on OAuth2 Proxy in the past couple of weeks and as a result we have fixed
several CRITICAL vulnerabilities.

The security vulnerabilities include multiple authentication bypasses and a potential session fixation attack.
For more details and to identify if you are effects, we urge all users of OAuth2 Proxy to read the security 
disclosures.

- (Critical) [GHSA-5hvv-m4w4-gf6v](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-5hvv-m4w4-gf6v) fix: health check user-agent authentication bypass 
- (Critical) [GHSA-7x63-xv5r-3p2x](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-7x63-xv5r-3p2x) fix: authentication bypass via X-Forwarded-Uri header spoofing
- (High) [GHSA-pxq7-h93f-9jrg](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-pxq7-h93f-9jrg) fix: fragment evaluation as part of the allowed routes
- (Moderate) [GHSA-c5c4-8r6x-56w3](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-c5c4-8r6x-56w3) fix: email validation bypass via malformed multi-@ email claims

Furthermore, for improving the security of OAuth2 Proxy we introduced a new flag `--trusted-proxy-ip` that allows users
to explicitly specify trusted reverse proxy IPs for the `X-Forwarded-*` headers. This is an important step to prevent 
potential header spoofing attacks and to ensure that OAuth2 Proxy only trusts headers from known and trusted sources.
We highly recommend users to review their deployment architecture and consider using this flag to enhance the security 
of their OAuth2 Proxy instances. Check the docs for more details: https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview#proxy-options

Furthermore, we want to thank everyone who contributed to the audits and reported potential issues to make open source 
software like OAuth2 Proxy more secure for everyone. 

## Breaking Changes

## Changes since v7.15.1

- [#3411](https://github.com/oauth2-proxy/oauth2-proxy/pull/3411) chore(deps): update gomod dependencies (@tuunit)
- [#3333](https://github.com/oauth2-proxy/oauth2-proxy/pull/3333) fix: invalidate session on fatal OAuth2 refresh errors (@frhack)
- [GHSA-f24x-5g9q-753f](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-f24x-5g9q-753f) fix: clear session cookie at beginning of signinpage handler (@fnoehWM / @bella-WI / @tuunit)
- [GHSA-5hvv-m4w4-gf6v](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-5hvv-m4w4-gf6v) fix: health check user-agent authentication bypass (@tuunit)
- [GHSA-7x63-xv5r-3p2x](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-7x63-xv5r-3p2x) fix: authentication bypass via X-Forwarded-Uri header spoofing (@tuunit)
- [GHSA-pxq7-h93f-9jrg](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-pxq7-h93f-9jrg) fix: fragment evaluation as part of the allowed routes (@tuunit)
- [GHSA-c5c4-8r6x-56w3](https://github.com/oauth2-proxy/oauth2-proxy/security/advisories/GHSA-c5c4-8r6x-56w3) fix: email validation bypass via malformed multi-@ email claims (@tuunit)

# V7.15.1

## Release Highlights

- 🐛 Squashed some bugs
-  🕵️‍♀️ Vulnerabilities have been addressed
    - [CVE-2026-33186](https://nvd.nist.gov/vuln/detail/CVE-2026-33186)
      OAuth2 Proxy was not impacted by this vulnerability as it isn't in the path of execution

## Important Notes

## Breaking Changes

## Changes since v7.15.0

- [#3382](https://github.com/oauth2-proxy/oauth2-proxy/pull/3382) chore(deps): update gomod and golangci/golangci-lint to v2.11.4 (@tuunit)
- [#3374](https://github.com/oauth2-proxy/oauth2-proxy/pull/3374) fix: handle Unix socket RemoteAddr in IP resolution (@H1net)
- [#3381](https://github.com/oauth2-proxy/oauth2-proxy/pull/3381) fix: do not log error for backend logout 204 (@artificiosus)
- [#3327](https://github.com/oauth2-proxy/oauth2-proxy/pull/3327) fix: improve logging when session refresh token is missing (@yosri-brh)
- [#2767](https://github.com/oauth2-proxy/oauth2-proxy/pull/2767) fix: propagate errors during route building (@sybereal)

# V7.15.0

## Release Highlights

- 🔒 OIDC JWT signing algorithms can now be configured
- 🍪 CSRF cookie improvements (SameSite option, proper expiration validation)
- 🧪 Configuration validation flag: --config-test
- 🔌 Unix socket file mode support
- 👤 Session state can now be extend with arbitrary claims from ID Token and upstream IDP user profiles endpoint
    - This opens the door for multiple features like:
    - Additional arbitrary header values for any claims your IDP provides
    - Extended OAuth2 Proxy UserInfo endpoint with all additional claims
    - Read the docs [here](https://oauth2-proxy.github.io/oauth2-proxy/configuration/alpha-config#how-to-utilize-arbitrary-claims-provided-by-your-identity-provider)

## Important Notes

CSRF cookie validation now correctly uses `CSRFExpire` instead of `Expire`. If you relied on the previous behavior, review your session timeout configuration.
Check the [documentation(https://oauth2-proxy.github.io/oauth2-proxy/configuration/overview#cookie-options) for `cookie-csrf-expire`.

## Breaking Changes

## Changes since v7.14.3

- [#3352](https://github.com/oauth2-proxy/oauth2-proxy/pull/3352) fix: backend logout URL call on sign out (#3172)(@vsejpal)
- [#3332](https://github.com/oauth2-proxy/oauth2-proxy/pull/3332) ci: distribute windows binary with .exe extension (@igitur)
- [#2685](https://github.com/oauth2-proxy/oauth2-proxy/pull/2685) feat: allow arbitrary claims from the IDToken and IdentityProvider UserInfo endpoint to be added to the session state (@vegetablest)
- [#3278](https://github.com/oauth2-proxy/oauth2-proxy/pull/3278) feat: possibility to inject id_token in redirect url during sign out (@albanf)
- [#2851](https://github.com/oauth2-proxy/oauth2-proxy/pull/2851) feat: add support for specifying allowed OIDC JWT signing algorithms (#2753) (@andoks / @tuunit)
- [#3369](https://github.com/oauth2-proxy/oauth2-proxy/pull/3369) fix: use CSRFExpire instead of Expire for CSRF cookie validation (@Br1an67)
- [#3365](https://github.com/oauth2-proxy/oauth2-proxy/pull/3365) fix: filter empty strings from allowed groups (@Br1an67)
- [#3338](https://github.com/oauth2-proxy/oauth2-proxy/pull/3338) feat: add --config-test flag for validating configuration (@MayorFaj)
- [#3347](https://github.com/oauth2-proxy/oauth2-proxy/pull/3347) feat: add same site option for csrf cookies (@jvnoije)
- [#3376](https://github.com/oauth2-proxy/oauth2-proxy/pull/3376) feat: allow setting unix socket file mode when declaring listener (@Tristan971 / @tuunit)

# V7.14.3

## Release Highlights

- 🔵 Go1.25.7 and upgrade of dependencies to latest versions
  - Fixes [CVE-2025-68121](https://nvd.nist.gov/vuln/detail/cve-2025-68121)
- 🐛 Bug fixes
  - Allow Redis URL parameters to configure username, password and max idle connection timeout if the matching configuration is empty.

## Important Notes

We improved our supply chain security by added additional checks to prevent potential command injection in the publish release workflow and to ensure that it can only be triggered from branches originating in the local repository. This potential issue was reported by automated systems as well as a couple of security researchers, and we want to thank everyone for their diligence in looking out for the security of the project. Especially Aastha Aggarwal for her detailed report and follow-up. @Aastha2602


## Breaking Changes

## Changes since v7.14.2

- [#3183](https://github.com/oauth2-proxy/oauth2-proxy/pull/3183) fix: allow URL parameters to configure username, password and max idle connection timeout if the matching configuration is empty.

# V7.14.2

## Release Highlights

- Revert AuthOnly endpoint change from v7.14.1 that caused issues when using `skip-provider-button` enabled

## Important Notes

- This release reverts the change made in v7.14.1 that caused issues when using the `skip-provider-button` enabled. Now, when a session does not exist, the AuthOnly endpoint will send a 401 status code as expected instead of a 302 redirect. And instead we extended the documentation to clarify the behavior when using `nginx` with `auth_request` and `skip-provider-button` and how to properly configure redirects for browser and API routes.

## Breaking Changes

## Changes since v7.14.1

- [#3314](https://github.com/oauth2-proxy/oauth2-proxy/pull/3314) revert: fix: skip provider button auth only redirect (#3309) (@StefanMarkmann / @tuunit)
- [#3315](https://github.com/oauth2-proxy/oauth2-proxy/pull/3315) docs: clarify browser vs API routes for nginx auth_request redirects (@StefanMarkmann)

# V7.14.1

## Release Highlights

- 🔵 Go1.25.6 and upgrade of dependencies to latest versions
- 🐛 Bug fixes
  - AuthOnly now starts the auth flow and send status code 302 if no session exists and skip-provider-button is true
  - Fixed static upstream validation issue due to incorrect defaults

## Important Notes

## Breaking Changes

## Changes since v7.14.0

- [#3309](https://github.com/oauth2-proxy/oauth2-proxy/pull/3309) fix: Return 302 redirect from AuthOnly endpoint when skip-provider-button is true (@StefanMarkmann)
- [#3302](https://github.com/oauth2-proxy/oauth2-proxy/pull/3302) fix: static upstreams failing validation due to `passHostHeader` and `proxyWebSockets` defaults being set incorrectly (@sourava01 / @tuunit)
- [#3312](https://github.com/oauth2-proxy/oauth2-proxy/pull/3312) chore(deps): upgrade to go1.25.6 and latest dependencies (@tuunit)

# V7.14.0

## Release Highlights

- 🕵️‍♀️ Vulnerabilities have been addressed
  - [CVE-2025-61729](https://access.redhat.com/security/cve/cve-2025-61729)
  - [CVE-2025-61727](https://access.redhat.com/security/cve/cve-2025-61727)
  - [CVE-2025-47914](https://access.redhat.com/security/cve/cve-2025-47914)
  - [CVE-2025-58181](https://access.redhat.com/security/cve/cve-2025-58181)
- 🗂️ AMajor Alpha Config YAML parsing revamped for better extensibility and preparing v8
- 🐛 Squashed some bugs

## Important Notes

This release introduces a breaking change for Alpha Config users and moves us significantly 
closer to removing legacy configuration parameters, making the codebase of OAuth2 Proxy more
future proof and extensible.

From v7.14.0 onward, header injection sources must be explicitly nested. If you
previously relied on squashed fields, update to the new structure before upgrading:

```yaml
# before v7.14.0
injectRequestHeaders:
- name: X-Forwarded-User
  values:
  - claim: user
- name: X-Custom-Secret-header
  values:
  - value: my-super-secret

# v7.14.0 and later
injectRequestHeaders:
- name: X-Forwarded-User
  values:
  - claimSource:
      claim: user
- name: X-Custom-Secret-header
  values:
  - secretSource:
      value: my-super-secret
```

Furthermore, Alpha Config now fully supports configuring the `Server` struct using YAML.

```yaml
// Server represents the configuration for the Proxy HTTP(S) configuration.
type Server struct {
	// BindAddress is the address on which to serve traffic.
	BindAddress string `yaml:"bindAddress,omitempty"`

	// SecureBindAddress is the address on which to serve secure traffic.
	SecureBindAddress string `yaml:"secureBindAddress,omitempty"`

	// TLS contains the information for loading the certificate and key for the
	// secure traffic and further configuration for the TLS server.
	TLS *TLS `yaml:"tls,omitempty"`
}

// TLS contains the information for loading a TLS certificate and key
// as well as an optional minimal TLS version that is acceptable.
type TLS struct {
    // Key is the TLS key data to use.
	Key *SecretSource `yaml:"key,omitempty"`
    // Cert is the TLS certificate data to use.
	Cert *SecretSource `yaml:"cert,omitempty"`
    // MinVersion is the minimal TLS version that is acceptable.
	MinVersion string `yaml:"minVersion,omitempty"`
    // CipherSuites is a list of TLS cipher suites that are allowed.
	CipherSuites []string `yaml:"cipherSuites,omitempty"`
}
```

More about how to use Alpha Config can be found in the [documentation](https://oauth2-proxy.github.io/oauth2-proxy/configuration/alpha-config#server).

We are committed to Semantic Versioning and usually avoid breaking changes without a major version release.
Advancing Alpha Config toward its Beta stage required this exception, and even for the Alpha Config we try
to keep breaking changes in v7 to a minium. Thank you for understanding the need for this step to prepare
the project for future maintainability and future improvements like structured logging.

## Breaking Changes

- Alpha Config: header injection no longer supports squashed claim/secret sources; they must now be set explicitly (see example above).

## Changes since v7.13.0

- [#2628](https://github.com/oauth2-proxy/oauth2-proxy/pull/2628) feat(structured config): revamp of yaml parsing using mapstructure decoder and custom decoders (@tuunit)
- [#3197](https://github.com/oauth2-proxy/oauth2-proxy/pull/3197) fix: NewRemoteKeySet is not using DefaultHTTPClient (@rsrdesarrollo / @tuunit)
- [#3292](https://github.com/oauth2-proxy/oauth2-proxy/pull/3292) chore(deps): upgrade gomod and bump to golang v1.25.5 (@tuunit)
- [#3304](https://github.com/oauth2-proxy/oauth2-proxy/pull/3304) fix: added conditional so default is not always set and env vars are honored fixes 3303 (@pixeldrew)
- [#3264](https://github.com/oauth2-proxy/oauth2-proxy/pull/3264) fix: more aggressively truncate logged access_token (@MartinNowak / @tuunit)
- [#3267](https://github.com/oauth2-proxy/oauth2-proxy/pull/3267) fix: Session refresh handling in OIDC provider (@gysel) 
- [#3290](https://github.com/oauth2-proxy/oauth2-proxy/pull/3290) fix: WebSocket proxy to respect PassHostHeader setting (@UnsignedLong)

# V7.13.0
-->

### 7.4.0-0.2.0 (2023-07-12)

### 7.1.2-0.1.1 (2023-02-01)

* [EOS-10808] Clear extra cookies with same domain as session cookie

### 7.1.2-0.1.0 (2022-07-21)

* Use new versioning schema
* Adapt repo to new CICD
* Bump alpine version to fix vulnerabilities
* [EOS-5416] Make sis path configurable
* [EOS-5112] Clear extra cookies whenever session cookie is removed
* [EOS-5112] Use extra cookies info from request
* Clear extra cookies on sign out
* Redirect to provider specific URL on sign out
* Add jwt session store
* Add sis provider
* Adapt repo to Stratio CICD flow
* Use https://github.com/oauth2-proxy/oauth2-proxy/releases/tag/v7.1.2 as base

### Branched to branch-6.1 (2020-12-09)

* Add tenant and groups to userinfo
* Add SIS provider and JWT session support
* Adapt repo to Stratio CICD flow
