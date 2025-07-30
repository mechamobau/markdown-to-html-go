package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <arquivo_de_entrada.md> <arquivo_de_saida.html>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	markdown, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo de entrada: %v\n", err)
		os.Exit(1)
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)

	var htmlBuffer bytes.Buffer
	if err := md.Convert(markdown, &htmlBuffer); err != nil {
		fmt.Printf("Erro ao converter o Markdown: %v\n", err)
		os.Exit(1)
	}

	p := bluemonday.UGCPolicy()
	sanitizedHTML := p.SanitizeBytes(htmlBuffer.Bytes())

	err = os.WriteFile(outputFile, sanitizedHTML, 0644)
	if err != nil {
		fmt.Printf("Erro ao escrever o arquivo de saída: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Arquivo '%s' convertido com sucesso para '%s'\n", inputFile, outputFile)
}
