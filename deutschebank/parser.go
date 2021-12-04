package deutschebank

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	category "github.com/mshogin/tx2ledger/caterory"
)

const ParserName = "db"

type dbParser struct {
	path string

	txs    []*tx
	txsRaw []string
}

type tx struct {
	Date     time.Time
	Name     string
	Category string
	Amount   float64
}

func NewParser(path string) *dbParser {
	return &dbParser{path: path}
}

// Load transactions from file
func (m *dbParser) Load() error {
	fh, err := os.Open(m.path)
	if err != nil {
		return fmt.Errorf("cannot open file %q: %w", m.path, err)
	}
	defer fh.Close()
	return m.LoadStream(fh)
}

// Load transactions from reader
func (m *dbParser) LoadStream(r io.Reader) error {
	transactionsBlock := false
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "Booking date") {
			transactionsBlock = true
			continue
		}
		if strings.HasPrefix(l, "Account balance") {
			transactionsBlock = false
			continue
		}
		if transactionsBlock {
			m.txsRaw = append(m.txsRaw, l)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("cannot scan transactions file: %w", err)
	}

	return nil
}

// Parse transactions
func (m *dbParser) Parse(categoryConfigPath string) error {
	fh, err := os.Open(categoryConfigPath)
	if err != nil {
		return fmt.Errorf("cannot open category config: %w", err)
	}
	defer fh.Close()
	categoryParser, err := category.CreateParser(fh)
	if err != nil {
		return fmt.Errorf("cannot create category parser: %w", err)
	}
	for _, l := range m.txsRaw {
		tx := parseTransaction(categoryParser, l)
		m.txs = append(m.txs, tx)
		fmt.Printf("%+v\n", m.txs[len(m.txs)-1]) // output for debug
	}
	return nil
}

func parseTransaction(categoryParser category.Categories, l string) *tx {
	parts := strings.Split(l, ";")
	date := parseDate(parts[0])
	name, category, err := categoryParser.Parse(parts[4], parts[3], parts[2])
	if err != nil {
		panic(fmt.Sprintf("cannot parse transaction: %s", strings.Join(parts, "\n")))
	}
	amount := parseFloat(parts[15], parts[16])
	return &tx{
		Date:     date,
		Name:     name,
		Category: category,
		Amount:   amount,
	}
}

func parseDate(l string) time.Time {
	p := strings.Split(l, "/")
	m, errM := strconv.Atoi(p[0])
	d, errD := strconv.Atoi(p[1])
	y, errY := strconv.Atoi(p[2])
	if errM != nil || errD != nil || errY != nil {
		panic("wrong date time")
	}
	t := time.Time{}
	t = t.AddDate(y-1, m-1, d-1)
	return t
}

func parseFloat(credit, debit string) float64 {
	credit = strings.ReplaceAll(credit, ",", "")
	debit = strings.ReplaceAll(debit, ",", "")
	if s, err := strconv.ParseFloat(credit, 64); err == nil {
		return math.Abs(s)
	}
	if s, err := strconv.ParseFloat(debit, 64); err == nil {
		return math.Abs(s)
	}
	panic("could not parse float from credit: " + credit + " debit:" + debit)
}
