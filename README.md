# Markdown to HTML Converter

Um conversor simples de Markdown para HTML escrito em Go, que utiliza bibliotecas robustas para garantir conversão segura e eficiente.

## 📋 Descrição

Este projeto converte arquivos Markdown em HTML usando as seguintes bibliotecas:
- **Goldmark**: Parser e renderizador Markdown rápido e extensível
- **Bluemonday**: Sanitizador HTML para segurança
- **GFM**: Suporte a GitHub Flavored Markdown
- **Footnotes**: Suporte a notas de rodapé

## 🚀 Funcionalidades

- ✅ Conversão de Markdown para HTML
- ✅ Suporte a GitHub Flavored Markdown (GFM)
- ✅ Suporte a notas de rodapé
- ✅ Sanitização automática de HTML
- ✅ IDs automáticos para cabeçalhos
- ✅ Quebras de linha hard-wrapped
- ✅ Interface de linha de comando simples

## 📦 Pré-requisitos

- Go 1.24.5 ou superior

## 🛠️ Instalação

1. Clone o repositório:
```bash
git clone <url-do-repositorio>
cd markdown-to-html-go
```

2. Instale as dependências:
```bash
go mod tidy
```

## 💻 Como usar

### Execução direta
```bash
go run main.go <arquivo_de_entrada.md> <arquivo_de_saida.html>
```

### Exemplo
```bash
go run main.go exemplo.md saida.html
```

### Compilação
```bash
go build -o markdown-converter main.go
./markdown-converter exemplo.md saida.html
```

## 📝 Exemplo de uso

O projeto inclui um arquivo de exemplo (`exemplo.md`) que demonstra as funcionalidades:

```markdown
# Título Principal

Este é um exemplo de conversão de **Markdown** para *HTML* em Go.

## Lista de Itens

* Item 1
* Item 2
* Item 3

## Código

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

## 🔧 Dependências

- `github.com/yuin/goldmark` - Parser e renderizador Markdown
- `github.com/microcosm-cc/bluemonday` - Sanitizador HTML
- `github.com/yuin/goldmark/extension` - Extensões do Goldmark
- `github.com/yuin/goldmark/parser` - Parser do Goldmark
- `github.com/yuin/goldmark/renderer/html` - Renderizador HTML

## 🛡️ Segurança

O projeto utiliza o Bluemonday para sanitizar o HTML gerado, removendo conteúdo potencialmente perigoso como:
- Scripts maliciosos
- Eventos JavaScript
- Tags HTML inseguras

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para:
1. Fazer um fork do projeto
2. Criar uma branch para sua feature
3. Fazer commit das suas mudanças
4. Fazer push para a branch
5. Abrir um Pull Request

## 📞 Suporte

Se você encontrar algum problema ou tiver dúvidas, abra uma issue no repositório. 