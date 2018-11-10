package main

import (
	"os"

	"github.com/moisespsena/go-ioutil"
	"github.com/mps-edu/glivros/cmd/cli"
	"github.com/mps-edu/glivros/core"
)

func main() {
	LivrosManager, err := core.NewLivrosManager()
	if err != nil {
		// se o erro for de que o arquivo "livros.json" nao existe, ignora, pois o mesmo será criado no primeiro
		// cadastro. Caso contrario, termina com erro.
		if os.IsNotExist(err) {
			err = nil
		} else {
			panic(err)
		}
	}

	lmCli := cli.LivrosManagerCLI{
		LineReader:    iou.STDMessageLR,
		LivrosManager: LivrosManager,
	}

	if err := lmCli.Menu(); err != nil {
		panic(err)
	}
}
