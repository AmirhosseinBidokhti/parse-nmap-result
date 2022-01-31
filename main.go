package main

import (
	"encoding/xml"
	"fmt"
	"parse_nmap_result/structs"
	utils "parse_nmap_result/utilities"
	"regexp"

	tld "github.com/jpillora/go-tld"
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

	Host := nmapRes.Host

	for _, v := range Host {

		for _, v := range v.Ports.Port {

			output := v.Script.Output

			if len(output) > 0 {

				re := regexp.MustCompile(`(?:[\w-]+\.)+[\w-]+`)

				submatchall := re.FindAllString(output, -1)

				hosts = append(hosts, submatchall...)

				// for _, element := range submatchall {
				// 	//fmt.Println(element)
				// 	hosts = append(hosts, element)
				// }

			}

		}

	}

	for _, host := range nmapRes.Host {

		hosts = append(hosts, host.Hostnames.Hostname.Name)

	}

	return utils.RemoveDuplicateStr(hosts)

}

func main() {
	var XMLdata = utils.ReadXML("./walmart.sample.xml")

	var hosts []string
	var uniqueHosts []string

	d := parser(XMLdata)
	for _, v := range d {
		//fmt.Println(v)
		url := "https://" + v
		//fmt.Println(url) // so far all unique values in the d
		hosts = append(hosts, url)
	}

	// Domains
	// for _, v := range hosts {
	// 	parsedUrl, _ := tld.Parse(v)
	// 	//fmt.Println()
	// 	uniqueHosts = append(uniqueHosts, parsedUrl.Domain+"."+parsedUrl.TLD)
	// }

	// for _, v := range utils.RemoveDuplicateStr(uniqueHosts) {
	// 	fmt.Println(v)
	// }

	// Subdomains
	for _, v := range hosts {
		parsedUrl, _ := tld.Parse(v)

		// If there is a subdomain except the "www" concat it
		if len(parsedUrl.Subdomain) > 0 && parsedUrl.Subdomain != "www" {
			uniqueHosts = append(uniqueHosts, parsedUrl.Subdomain+"."+parsedUrl.Domain+"."+parsedUrl.TLD)
		} else {
			uniqueHosts = append(uniqueHosts, parsedUrl.Domain+"."+parsedUrl.TLD)
		}

	}

	for _, v := range utils.RemoveDuplicateStr(uniqueHosts) {
		fmt.Println(v)
	}

	// host = utils.RemoveDuplicateStr(host)
	// // Domains
	// for _, v := range host {

	// 	parsedUrl, _ := tld.Parse(v)

	// 	fmt.Println(parsedUrl.Domain + "." + parsedUrl.TLD)

	// }

	// Sub-domains
	// for _, v := range utils.RemoveDuplicateStr(host) {

	// 	parsedUrl, _ := tld.Parse(v)

	// 	fmt.Println(parsedUrl.Subdomain + "." + parsedUrl.Domain + "." + parsedUrl.TLD)

	// }

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
