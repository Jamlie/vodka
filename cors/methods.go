package cors

type AllowedHeader = string

const (
	Accept              AllowedHeader = "Accept"
	Authorization       AllowedHeader = "Authorization"
	AcceptLanguage      AllowedHeader = "Accept-Language"
	ContentRange        AllowedHeader = "Content-Range"
	ContentType         AllowedHeader = "Content-Type"
	ContentLanguage     AllowedHeader = "Content-Language"
	CacheControl        AllowedHeader = "Cache-Control"
	IfMatch             AllowedHeader = "If-Match"
	IfNoneMatch         AllowedHeader = "If-None-Match"
	IfModifiedSince     AllowedHeader = "If-Modified-Since"
	IfUnmodifiedSince   AllowedHeader = "If-Unmodified-Since"
	IfRange             AllowedHeader = "If-Range"
	Range               AllowedHeader = "Range"
	Origin              AllowedHeader = "Origin"
	XRequestedWith      AllowedHeader = "X-Requested-With"
	XHTTPHeaderOverride AllowedHeader = "X-HTTP-Method-Override"
	XForwardedFor       AllowedHeader = "X-Forwarded-For"
	XForwardedProto     AllowedHeader = "X-Forwarded-Proto"
	XForwardedHost      AllowedHeader = "X-Forwarded-Host"
	XForwardedPort      AllowedHeader = "X-Forwarded-Port"
	XAccessToken        AllowedHeader = "X-Access-Token"
	XCSRFToken          AllowedHeader = "X-CSRF-Token"
)
