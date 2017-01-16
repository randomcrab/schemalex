package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strconv"
)

var (
	packageName = flag.String("package", "schemalex", "name of package")
	fileName    = flag.String("file", "token.go", "name of file")
)

func main() {
	flag.Parse()
	if err := _main(); err != nil {
		log.Fatal(err)
	}
}

func _main() error {
	var buf bytes.Buffer

	buf.WriteString(`// generated by internal/cmd/gentokens/main.go`)
	buf.WriteString("\n\npackage " + *packageName)
	buf.WriteString("\n\n// TokenType describes the possible types of tokens that schemalex understands")
	buf.WriteString("\ntype TokenType int")

	buf.WriteString("\n\n// Token represents a token")
	buf.WriteString("\ntype Token struct {")
	buf.WriteString("\nType TokenType")
	buf.WriteString("\nValue string")
	buf.WriteString("\nPos int")
	buf.WriteString("\nLine int")
	buf.WriteString("\nCol int")
	buf.WriteString("\nEOF bool")
	buf.WriteString("\n}")

	buf.WriteString("\n\n// NewToken creates a new token of type `t`, with value `v`")
	buf.WriteString("\nfunc NewToken(t TokenType, v string) *Token {")
	buf.WriteString("\nreturn &Token{Type: t, Value: v}")
	buf.WriteString("\n}")

	buf.WriteString("\n\n// List of possible tokens")
	buf.WriteString("\nconst (")
	buf.WriteString("\nILLEGAL TokenType = iota")

	tokens := []struct {
		Comment string
		Ident   string
	}{
		{Ident: "EOF"},
		{Ident: "SPACE"},
		{Ident: "IDENT"},
		{Ident: "BACKTICK_IDENT"},
		{Ident: "DOUBLE_QUOTE_IDENT"},
		{Ident: "SINGLE_QUOTE_IDENT"},
		{Ident: "NUMBER"},
		{Ident: "LPAREN", Comment: "("},
		{Ident: "RPAREN", Comment: ")"},
		{Ident: "COMMA", Comment: ","},
		{Ident: "SEMICOLON", Comment: ";"},
		{Ident: "DOT", Comment: "."},
		{Ident: "SLASH", Comment: "/"},
		{Ident: "ASTERISK", Comment: "*"},
		{Ident: "DASH", Comment: "-"},
		{Ident: "PLUS", Comment: "+"},
		{Ident: "SINGLE_QUOTE", Comment: "'"},
		{Ident: "DOUBLE_QUOTE", Comment: "\""},
		{Ident: "EQUAL", Comment: "="},
		{Ident: "COMMENT_IDENT", Comment: `// /*   */, --, #`},
		{Ident: "ACTION"},
		{Ident: "AUTO_INCREMENT"},
		{Ident: "AVG_ROW_LENGTH"},
		{Ident: "BIGINT"},
		{Ident: "BINARY"},
		{Ident: "BIT"},
		{Ident: "BLOB"},
		{Ident: "BTREE"},
		{Ident: "CASCADE"},
		{Ident: "CHAR"},
		{Ident: "CHARACTER"},
		{Ident: "CHECK"},
		{Ident: "CHECKSUM"},
		{Ident: "COLLATE"},
		{Ident: "COMMENT"},
		{Ident: "COMPACT"},
		{Ident: "COMPRESSED"},
		{Ident: "CONNECTION"},
		{Ident: "CONSTRAINT"},
		{Ident: "CREATE"},
		{Ident: "CURRENT_TIMESTAMP"},
		{Ident: "DATA"},
		{Ident: "DATABASE"},
		{Ident: "DATE"},
		{Ident: "DATETIME"},
		{Ident: "DECIMAL"},
		{Ident: "DEFAULT"},
		{Ident: "DELAY_KEY_WRITE"},
		{Ident: "DELETE"},
		{Ident: "DIRECTORY"},
		{Ident: "DISK"},
		{Ident: "DOUBLE"},
		{Ident: "DROP"},
		{Ident: "DYNAMIC"},
		{Ident: "ENGINE"},
		{Ident: "ENUM"},
		{Ident: "EXISTS"},
		{Ident: "FIRST"},
		{Ident: "FIXED"},
		{Ident: "FLOAT"},
		{Ident: "FOREIGN"},
		{Ident: "FULL"},
		{Ident: "FULLTEXT"},
		{Ident: "HASH"},
		{Ident: "IF"},
		{Ident: "INDEX"},
		{Ident: "INSERT_METHOD"},
		{Ident: "INT"},
		{Ident: "INTEGER"},
		{Ident: "KEY"},
		{Ident: "KEY_BLOCK_SIZE"},
		{Ident: "LAST"},
		{Ident: "LONGBLOB"},
		{Ident: "LONGTEXT"},
		{Ident: "MATCH"},
		{Ident: "MAX_ROWS"},
		{Ident: "MEDIUMBLOB"},
		{Ident: "MEDIUMINT"},
		{Ident: "MEDIUMTEXT"},
		{Ident: "MEMORY"},
		{Ident: "MIN_ROWS"},
		{Ident: "NO"},
		{Ident: "NOT"},
		{Ident: "NULL"},
		{Ident: "NUMERIC"},
		{Ident: "ON"},
		{Ident: "PACK_KEYS"},
		{Ident: "PARTIAL"},
		{Ident: "PASSWORD"},
		{Ident: "PRIMARY"},
		{Ident: "REAL"},
		{Ident: "REDUNDANT"},
		{Ident: "REFERENCES"},
		{Ident: "RESTRICT"},
		{Ident: "ROW_FORMAT"},
		{Ident: "SET"},
		{Ident: "SIMPLE"},
		{Ident: "SMALLINT"},
		{Ident: "SPATIAL"},
		{Ident: "STATS_AUTO_RECALC"},
		{Ident: "STATS_PERSISTENT"},
		{Ident: "STATS_SAMPLE_PAGES"},
		{Ident: "STORAGE"},
		{Ident: "TABLE"},
		{Ident: "TABLESPACE"},
		{Ident: "TEMPORARY"},
		{Ident: "TEXT"},
		{Ident: "TIME"},
		{Ident: "TIMESTAMP"},
		{Ident: "TINYBLOB"},
		{Ident: "TINYINT"},
		{Ident: "TINYTEXT"},
		{Ident: "UNION"},
		{Ident: "UNIQUE"},
		{Ident: "UNSIGNED"},
		{Ident: "UPDATE"},
		{Ident: "USE"},
		{Ident: "USING"},
		{Ident: "VARBINARY"},
		{Ident: "VARCHAR"},
		{Ident: "YEAR"},
		{Ident: "ZEROFILL"},
	}

	for _, tok := range tokens {
		buf.WriteString("\n" + tok.Ident)
		if c := tok.Comment; c != "" {
			buf.WriteString("// " + c)
		}
	}
	buf.WriteString("\n)") // end const (

	buf.WriteString("\n\nvar keywordIdentMap = map[string]TokenType{")
	for _, tok := range tokens[20:] {
		buf.WriteString("\n" + strconv.Quote(tok.Ident) + ": " + tok.Ident + ",")
	}
	buf.WriteString("\n}")

	buf.WriteString("\n\nfunc (t TokenType) String() string {")
	buf.WriteString("\nswitch t {")
	buf.WriteString("\ncase ILLEGAL:")
	buf.WriteString("\nreturn \"ILLEGAL\"")
	for _, tok := range tokens {
		buf.WriteString("\ncase ")
		buf.WriteString(tok.Ident)
		buf.WriteByte(':')
		buf.WriteString("\nreturn ")
		buf.WriteString(strconv.Quote(tok.Ident))
	}
	buf.WriteString("\n}")
	buf.WriteString("\nreturn \"(invalid)\"")
	buf.WriteString("\n}")

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("%s", buf.Bytes())
		return err
	}

	f, err := os.Create(*fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(formatted)
	return nil
}
