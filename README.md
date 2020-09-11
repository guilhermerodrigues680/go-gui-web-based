## Configurações de ambiente

> https://golang.org/doc/code.html

```sh
# Imprimir variaveis de ambiente do go
#go env
#go env | grep "GOPATH\|GOBIN"

# Definir varivaies de ambiente go
#go env -w GOPATH=/Users/guilherme/Projetos/estudos-lista-encadeada-go
#go env -w GOBIN=/Users/guilherme/Projetos/estudos-lista-encadeada-go/bin

# Zerar varivaies de ambiente go
go env -u GOPATH
go env -u GOBIN
```