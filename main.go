package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"parse_nmap_result/structs"
	utils "parse_nmap_result/utilities"
	"regexp"

	tld "github.com/jpillora/go-tld"
)

func main() {

	var domains bool
	var subdomains bool
	var stdinRead bool

	flag.BoolVar(&domains, "d", false, "print domains")
	flag.BoolVar(&subdomains, "s", false, "print subdomains")
	flag.BoolVar(&stdinRead, "std", false, "read file from stdin")
	filePath := flag.String("fp", "test.txt", "file path to read from")
	flag.Parse()

	var XMLdata = utils.ReadXML(*filePath)

	var hosts []string
	var uniqueHosts []string

	d := parser(XMLdata)
	for _, v := range d {
		// tld package only works with an scheme specified for the url
		// adding a fake one so we can work with it.
		url := "https://" + v
		hosts = append(hosts, url)
	}

	// TODO: Ability to read file from stdin so we can use pipes
	// if stdinRead {
	// }

	if domains {
		for _, v := range hosts {
			parsedUrl, _ := tld.Parse(v)
			uniqueHosts = append(uniqueHosts, parsedUrl.Domain+"."+parsedUrl.TLD)
		}

		for _, v := range utils.RemoveDuplicateStr(uniqueHosts) {
			fmt.Println(v)
		}
	}

	if subdomains {
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
	}

}

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
			}
		}
	}

	for _, host := range nmapRes.Host {
		hosts = append(hosts, host.Hostnames.Hostname.Name)
	}

	return utils.RemoveDuplicateStr(hosts)
}
