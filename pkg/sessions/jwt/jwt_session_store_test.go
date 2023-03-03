package jwt

import (
	"testing"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
	sessionsapi "github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/sessions"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/logger"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/sessions/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	JWTKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAtZ4zRs2J8EAWLxWR+b0C3saGx5j8hnvjFmQp+PigqpA/Eue/
agk21/y9idPFMWcNqtX8NvkYL+fEsk/1Na17do6zbvZTBgTIPCy75IF8oc3OR/zt
aIic3YeJoKH4UcLxifOqnR+B1CaQMTjuBp2xZC4QK7FyXRJEN3taW3DmyCfvDfIA
HjDp3n5hQha0Vtp5J+nQmjDdiQFYpS69/WtJ1WhFJlnOcCAXEFA2+UhwNkhshgtA
+VnzE1AzZwRuCsqM97kMU9Eja+opkhn3BZnAaAE7XFwm645W95Ac3X67uNgMD4bh
auTjWy27IBhcN/CwxAhWo5z68GApRd7L38IJHQIDAQABAoIBAQCIj4HC1T3I1odX
tAJlJEgKNoCViGUowfKInZwpxtkYJwomLvdwxajlUvc5sXBuqyxNrkTNGROkwcLW
yOR6Dg3toXMuFi1rMyFUjdZiBTMvfs6Ctp3UohRBRm6nx+ItqEVyEzPQnSZD3RNC
z6m6c8w0paYnFHAHp3p/tVLFuujsykWYoAauos42LySFwfh3EdEaTido7lCdETAG
Vu+T8Mga9hbs4n+0jO0BdR/XITDTNLpge50tsoI4G6TJT5gHPVnuxR9K8xNjIHK2
3eC6fPSePXfNLh6isf4eJaRzI8ZWPHFooxbOWRMG4CEQ1hEO8U7QybQQWj98O3Pe
2piBEi/JAoGBAOD7eP9hAOflSxjp4Bv6uuOLQu+LQVmSuk7vtjSn/I6iUcm1/oWf
u6qxsKtUFdK0sfAmsUmIwHXthGV1qEjC19SyE+/dycH7fpBg071KFqO30H1/8cg/
76AVm84VvWSO1Db0otVKrCvtIZjF20BKixKTCLH9Q4v/R38yveojhJRXAoGBAM6o
OeW+BXoWY6eodptjamxE9BbE6HaH6mUM00M87lS8AbPkrQ0I49atrihTXvB6m6mI
GedEIcErAJIfWCDMce3k9dIUFIBhLdHyRsc+KHazhnA709xauE0B6lHfWY6i7hvV
s3TBpEcCjneRB6R9GX3OH9Zi9+cpj7EovL741MWrAoGBAMnOiJB60LcyJBSq5M30
L+OfrWD1xp60UM4xk3zUGmVPEJIg37e4uju4u8JS4GhqkRnbezd8pTai4RmpWlQ6
AiPVwLBuf2WzU6nqUMQASyJ75VZNh/GZ+DXebC2Frqcevxi0g8NTAfE8+d/xymN2
+hylKy2NAiP3zog4WcZGKcxtAoGBAMq7NbEv7OeMN08ucMyXhruYGWyM1xAQ3d0r
68S2bYgqt/DmkO2MnxbnY0akIyr+3N4/akn6CLMboH+4yBfE+K9MQetJT6NxsiWX
699iFwf7rhNEXd56EPtauah/17eaFsSvrFEJ9kLDO0gIutqe7vb/0zPZ+yCHITPG
pwMh0HnpAoGBAI+bNyeV4V2mcSrOimZEBI/A2C8+/e9iuDj5XnTd2TTCg4HRiHXi
MwT0luCoc/P4mz7Zvr6bnI5Jy1D6sS0MY8CTk6PLly0UFxB1Tw9d1RJNMJVUdT6H
tx0A+NS7+eKOZzPwm6oJjnQKFykzdoWlJoI0iXE/bpgx/HqGy7cumWs8
-----END RSA PRIVATE KEY-----`
)

func TestSessionStore(t *testing.T) {
	logger.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	RunSpecs(t, "JWT SessionStore")
}

var _ = Describe("JWT SessionStore Tests", func() {
	tests.RunSessionStoreTests(
		func(opts *options.SessionOptions, cookieOpts *options.Cookie) (sessionsapi.SessionStore, error) {
			// Set the connection URL
			opts.Type = options.JWTSessionStoreType
			opts.JWT.JWTKey = JWTKey
			return NewJWTSessionStore(opts, cookieOpts)
		}, nil)
})
