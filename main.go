package main

import (
	"flag"
	"log"

	"github.com/mshogin/tx2ledger/deutschebank"
)

func main() {
	txPath := flag.String(
		"tx",
		"",
		"Path to transaction file")
	category := flag.String(
		"c",
		"",
		"Path to category")
	output := flag.String(
		"o",
		"",
		"Path to ledger file")
	parser := flag.String(
		"parser",
		"db",
		"Transactions parser")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()

	switch *parser {
	case deutschebank.ParserName:
		p := deutschebank.NewParser(*txPath)
		if err := p.Load(); err != nil {
			log.Fatalf("cannot load the transactions: %s", err)
		}
		if err := p.Parse(*category); err != nil {
			log.Fatalf("cannot parse the transactions: %s", err)
		}
		if err := p.Dump(*output); err != nil {
			log.Fatalf("cannot dump write ledger file: %s", err)
		}
	}

}
