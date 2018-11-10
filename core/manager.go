package core

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

//======== GERENCIA ========

type LivrosManager struct {
	arquivo string
	items   map[int]*Livro
	lastID  int
}

func NewLivrosManager(arquivo ...string) (lm *LivrosManager, err error) {
	if arquivo == nil || arquivo[0] == "" {
		arquivo = []string{"livros.json"}
	}

	lm = &LivrosManager{arquivo: arquivo[0], items: map[int]*Livro{}, lastID: 0}
	if err = lm.carregar(); err != nil {
		if os.IsNotExist(err) {
			return
		}
		return nil, err
	}
	return
}

func (l *LivrosManager) NewID() int {
	l.lastID++
	return l.lastID
}

func (lm *LivrosManager) Cadastrar(l *Livro) (ID int, err error) {
	l.ID = lm.NewID()
	lm.items[l.ID] = l
	if err = lm.salvar(); err != nil {
		l.ID = 0
	} else {
		ID = l.ID
	}
	return
}

func (lm *LivrosManager) Get(idDoLivro int) (livro *Livro, ok bool) {
	livro, ok = lm.items[idDoLivro]
	return
}

func (lm *LivrosManager) Atualizar(l *Livro) (err error) {
	lm.items[l.ID] = l
	return lm.salvar()
}

func (lm *LivrosManager) Excluir(idDoLivro int) (err error) {
	if _, exists := lm.items[idDoLivro]; exists {
		delete(lm.items, idDoLivro)
		return lm.salvar()
	}
	return fmt.Errorf("O livro #%d não está cadastrado.", idDoLivro)
}

func (lm *LivrosManager) Publicar(idDoLivro int) (err error) {
	if livro, exists := lm.items[idDoLivro]; exists {
		if livro.IsPublicado() {
			return fmt.Errorf("O livro #%d já foi publicado.", idDoLivro)
		}

		now := time.Now()
		livro.DataDePublicacao = &now
		return lm.salvar()
	}
	return fmt.Errorf("O livro #%q não está cadastrado.", idDoLivro)
}

func (lm *LivrosManager) Despublicar(idDoLivro int) (err error) {
	if livro, exists := lm.items[idDoLivro]; exists {
		if !livro.IsPublicado() {
			return fmt.Errorf("O livro #%d não foi publicado.", idDoLivro)
		}

		livro.DataDePublicacao = nil
		return lm.salvar()
	}
	return fmt.Errorf("O livro #%q não está cadastrado.", idDoLivro)
}

func (lm *LivrosManager) Listar() (livros []*Livro) {
	for _, livro := range lm.items {
		livros = append(livros, livro)
	}

	lm.sortByTitulo(livros)
	return
}

func (lm *LivrosManager) Pesquisar(termo string) (livros []*Livro) {
	// coloca tudo em MAIUSCULO
	termo = strings.ToUpper(termo)
	for _, livro := range lm.items {
		// coloca o titulo em MAIUSCULO
		tituloMaiusculo := strings.ToUpper(livro.Titulo)
		// compara tudo em MAIUSCULO para nao fazer distincao entre MAIUSCULO e MINUSCULO
		if strings.Contains(tituloMaiusculo, termo) {
			livros = append(livros, livro)
		}
	}

	lm.sortByTitulo(livros)
	return
}

func (lm *LivrosManager) sortByTitulo(livros []*Livro) {
	// Ordena pelo Titulo do livro
	sort.Slice(livros, func(i, j int) bool {
		return strings.ToUpper(livros[i].Titulo) < strings.ToUpper(livros[j].Titulo)
	})
}
