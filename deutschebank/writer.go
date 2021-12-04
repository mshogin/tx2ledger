package deutschebank

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

var txTemplate = `
{{.LedgerDate}} * {{.Name}}
        {{.Category}}        {{.Amount}} EUR
        Assets:Checking
`
var txIncomeTemplate = `
{{.LedgerDate}} * {{.Name}}
        Assets:Checking        {{.Amount}} EUR
        {{.Category}}
`

func (m *tx) LedgerDate() string {
	return fmt.Sprintf("%4d/%02d/%02d", m.Date.Year(), int(m.Date.Month()), m.Date.Day())
}

func (m *tx) IsIncome() bool {
	return strings.Contains(m.Category, "Income")
}

func (m *tx) ToByte() []byte {
	t := txTemplate
	if m.IsIncome() {
		t = txIncomeTemplate
	}
	tmpl, err := template.New("render").Parse(t)
	if err != nil {
		panic(fmt.Errorf("cannot parse template: %s", err))
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, m); err != nil {
		panic(fmt.Errorf("cannot execute template: %s", err))
	}

	return buf.Bytes()
}

func (m *dbParser) Dump(fpath string) error {
	fh, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer fh.Close()
	if _, err := fh.Write([]byte("P 2021/09/01 EUR 82.00 RUR\n\n")); err != nil {
		return fmt.Errorf("cannot write the exchange rate: %w", err)
	}

	for _, tx := range m.txs {
		_, err := fh.Write(tx.ToByte())
		if err != nil {
			return fmt.Errorf("cannot write to file: %w", err)
		}
	}

	return nil
}
