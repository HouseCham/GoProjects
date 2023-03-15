package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//? MX record, or mail exchange record, is a DNS record that routes emails to specified mail servers.
	//? MX records essentially point to the IP addresses of a mail server's domain.
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	//? An SPF record added to Domain Name Service (DNS) servers tells recipient email servers that a message 
	//? came from an authorized sender IP address or could be from a phishing campaign.
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}


	//? DMARC, which stands for Domain-based Message Authentication, Reporting, and Conformance, 
	//? is a DNS TXT Record that can be published for a domain to control what happens if a message 
	//? fails authentication (i.e. the recipient server can't verify that the message's sender is who they say they are).
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}