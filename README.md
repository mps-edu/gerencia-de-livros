# glivros

Projeto exemplo de Gestão de Livros.


## Instalação

Execute no terminal de comandos:

```bash
go get -u github.com/mps-edu/glivros
```

## Diretório do código fonte

O código fonte foi instalado no diretório `$GOPATH/src/github.com/mps-edu/glivros`. Para acessá-lo, execete o comando
abaixo no terminal:

```bash
cd $GOPATH/src/github.com/mps-edu/glivros
```

## Divisões

O projeto está dividido em dois módulos, veja a seguir.

### 1. core

O programa principal, isto é, a "Gestão de Livros" de fato.

### 2. cmd

Contém os micro programas que fazem a interface com o usuário. Este, também está subdividido em outros.

#### 2.1 cmd/cli

Contém o programa que funciona e interage através da linha de comandos.

Para acessar, considerando o diretório atual como sendo a raiz do projeto, isto é
`$GOPATH/src/github.com/mps-edu/glivros`, acesse o diretório `cmd/cli` (`cd cmd/cli`).

O diretório `glivros-cli`, contém o `main.go` desse programa. Para executar ou compilar, acesse-o `cd glivros-cli`.

##### Compilando e Executando

Para executar somente, execute o comando `go run main.go` no terminal.

Para compilar, execute no terminal o comando `go build`. Será criado um arquivo executáve chamado `glivros-cli`.

Para executar o executável gerado, execute `./glivros-cli` no terminal.

# Obrigado!

Moisés P. Sena - https://github.com/moisespsena