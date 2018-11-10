package core

import (
	"fmt"
	"time"
)

//======== LIVRO ========

type Livro struct {
	ID               int
	Titulo           string
	Descricao        string
	DataDePublicacao *time.Time
}

func (l *Livro) String() (s string) {
	if l.ID != 0 {
		s = "#" + fmt.Sprint(l.ID) + " "
	}
	s += l.Titulo
	if l.DataDePublicacao != nil {
		s += ", publicado em " + l.DataDePublicacao.Format("Mon Jan _2 15:04:05 2006")
	}
	if l.Descricao != "" {
		s += " -> " + l.Descricao
	}
	return s
}

func (l *Livro) IsPublicado() bool {
	return l.DataDePublicacao != nil
}
