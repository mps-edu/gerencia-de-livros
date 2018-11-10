package cli

import (
	"fmt"
	"strconv"

	"github.com/moisespsena/go-ioutil"
	"github.com/mps-edu/glivros/core"
)

type LivrosManagerCLI struct {
	LineReader    *iou.MessageLineReader
	LivrosManager *core.LivrosManager
}

func (cli *LivrosManagerCLI) Menu() (err error) {
	l := cli.LineReader
	defer l.WithPrintInput()()
	formatter := &iou.FOptionsPairs{
		Message: "Escolha uma opção",
		Options: iou.StringsToPairs(
			"Cadastrar",
			"Mostrar",
			"Listar",
		).
			AddBlank().
			Add("q", "Sair"),
	}

	var option interface{}
	for option != "q" {
		if option, err = l.RequireF(formatter); err != nil {
			return err
		}

		switch option {
		case 1:
			err = cli.Cadastrar()
		case 2:
			err = cli.Mostrar()
		case 3:
			err = cli.Listar()
		case "q":
			return
		}

		if err != nil {
			return
		}
	}
	return
}

func (cli *LivrosManagerCLI) Cadastrar() (err error) {
	l := cli.LineReader
	var livro core.Livro
	if livro.Titulo, err = l.RequireS("Título"); err != nil {
		return fmt.Errorf("Erro ao ler o título: %v", err)
	}
	if livro.Descricao, err = l.ReadS("Descrição"); err != nil {
		return fmt.Errorf("Erro ao ler a descrição: %v", err)
	}
	if _, err = cli.LivrosManager.Cadastrar(&livro); err != nil {
		return fmt.Errorf("Erro ao cadastrar o livro: %v", err)
	} else {
		fmt.Fprintf(l.Writer, "Livro #%d cadastrado com Sucesso!\n", livro.ID)
	}
	return
}

func (cli *LivrosManagerCLI) Mostrar() (err error) {
	l := cli.LineReader
	var ids string
	if ids, err = l.RequireS("Codigo do livro"); err != nil {
		return fmt.Errorf("Erro ao ler o codigo do livro: %v", err)
	}
	var id int
	if id, err = strconv.Atoi(ids); err != nil {
		return fmt.Errorf("Erro tranforma em inteiro; %v", err)
	}
	//fmt.Printf("O id digitado e: %v\n",id)
	livro, ok := cli.LivrosManager.Get(id)
	if !ok {
		fmt.Fprintf(l.Writer, "o livro com o codigo %v nao esta cadastrado\n", id)
		return
	}

	fmt.Fprintln(l.Writer, livro)
	return
}

func (cli *LivrosManagerCLI) Listar() (err error) {
	livros := cli.LivrosManager.Listar()
	for i, livro := range livros {
		fmt.Fprintf(cli.LineReader.Writer, "%d.\t%v\n", i+1, livro.String())
	}
	return
}
