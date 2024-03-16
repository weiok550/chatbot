package config

type AppSettings struct {
	ServiceName  string `yaml:"serviceName" json:"service_name"`
	ServiceURL   string `yaml:"serviceUrl" json:"service_url"`
	CookieDomain string `yaml:"cookieDomain" json:"cookie_domain"`
	AllowOrigin string `yaml:"allowOrigin" json:"allow_origin"`
	LoginCookieDuration int32 `yaml:"loginCookieDuration" json:"login_cookie_duration"`
}