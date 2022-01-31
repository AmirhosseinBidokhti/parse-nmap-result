package main

import (
	"encoding/xml"
	"fmt"
	"parse_nmap_result/structs"
	utils "parse_nmap_result/utilities"
)

// func getDomain(fullUrl string) string {
// 	u, err := url.Parse(fullUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(u.Hostname())
// 	return u.Hostname()
// }

func parser(XMLdata []byte) []string {

	var nmapRes structs.Nmaprun
	xml.Unmarshal(XMLdata, &nmapRes)

	var hosts []string

	for _, host := range nmapRes.Host {

		hosts = append(hosts, host.Hostnames.Hostname.Name)

	}

	return utils.RemoveDuplicateStr(hosts)
	// for _, v := range hosts {
	// 	fmt.Println(v)
	// }

}

func main() {
	var XMLdata = utils.ReadXML("./sample.xml")
	d := parser(XMLdata)
	for _, v := range d {
		fmt.Println(v)
	}
}

// func main() {
// 	urls := []string{
// 		"google.com",
// 		"http://blog.google",
// 		"medi-cal.ca.gov/",
// 		"ato.gov.au",
// 		"a.very.complex-domain.co.uk:8080/foo/bar",
// 		"a.domain.that.is.unmanaged",
// 	}
// 	for _, url := range urls {
// 		u, _ := tld.Parse(url)
// 		fmt.Printf("%50s = [ %s ] [ %s ] [ %s ] [ %s ] [ %s ] [ %t ]\n",
// 			u, u.Subdomain, u.Domain, u.TLD, u.Port, u.Path, u.ICANN)
// 	}
// }

// tld package only works with an scheme specified for the url
// adding a fake one.
// url := "https://" + host.Hostnames.Hostname.Name

// parsedUrl, _ := tld.Parse(url)
// parsedUrl, _ := tld.Parse(url)

// hosts = append(hosts, parsedUrl.Domain+"."+parsedUrl.TLD)
// 	hosts = append(hosts, parsedUrl.Domain+"."+parsedUrl.TLD)
