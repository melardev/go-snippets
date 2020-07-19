package net

import "github.com/melardev/SnippetsGoSTL/net/auxiliary"

func LaunchNetSnippets() {
	auxiliary.DnsLookup()
	auxiliary.DnsLookupCustomResolver()
	auxiliary.GetFLD()
	auxiliary.GetHostname()
	auxiliary.GetTLD()
	auxiliary.GetIpFromString()
	auxiliary.IpInsideRange()
	auxiliary.IP2Integer()
	auxiliary.IsIpInCIDR()
	auxiliary.LookupAddressSnippet()
	auxiliary.UrlParseSnippet()

	getWithCustomHost()
	GetWithOpaque()
	getRequest()
	GetWithoutSSLCheck()
	parseQuery()
	postFormFile()
	simpleGet()
	simplePostForm()
	simpleJsonPost()
	GetWithEnvProxy()
	GetWithProxy()
	GetWithCustomHost()

}
