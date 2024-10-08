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

	// The bufio.Scanner is useful for reading input line by line.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord,hasDMARC,dmarcRecord\n")

	// for scanner.Scan(): This loop continues as long as there is input to read. The Scan() method reads the next line and returns true if there are more lines to read.
	//checkDomain(scanner.Text()): Inside the loop, scanner.Text() retrieves the current line of input (as a string), and this string is passed to the checkDomain function for processing.

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: Could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// net.LookupMX(domain):

	// This function is part of the net package in Go and is used to look up the MX records for a given domain.

	// MX records are like an address book for email. They tell your computer (or the email system) which mailboxes (or mail servers) to send the letter (or email) to.

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	//This line uses a for loop to iterate over each record in the txtRecords slice.

	// txtRecords is expected to contain all the TXT records retrieved for a given domain using net.LookupTXT(domain).

	// if strings.HasPrefix(record, "v=spf1") {

	//This line uses the strings.HasPrefix function to check if the current record starts with the string "v=spf1".

	//SPF records typically start with "v=spf1" to indicate that the record is using SPF version 1.

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
