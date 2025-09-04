module github.com/libdns/cloudflare/e2e

go 1.18

require (
	github.com/libdns/cloudflare v0.0.0
	github.com/libdns/libdns v1.1.0
)

replace (
	github.com/libdns/cloudflare => ../
	github.com/libdns/libdns => github.com/AndrianBdn/libdns v0.0.0-20250903124819-40d75e3218d5
)
