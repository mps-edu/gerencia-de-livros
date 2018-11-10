package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*------------ Métodos para carregar e armazenar os dados ------------*/

func (l *LivrosManager) carregar() (err error) {
	if _, err = os.Stat(l.arquivo); os.IsNotExist(err) {
		return err
	}

	var data []byte
	if data, err = ioutil.ReadFile(l.arquivo); err != nil {
		return fmt.Errorf("Erro ao ler o arquivo %q: %v", l.arquivo, err)
	}

	items := map[int]*Livro{}
	if err = json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("Erro ao importar: %v", err)
	}

	l.items = items
	l.lastID = 0

	for key := range l.items {
		if key > l.lastID {
			l.lastID = key
		}
	}
	return
}

func (l *LivrosManager) salvar() (err error) {
	var data []byte
	if data, err = json.MarshalIndent(l.items, "", " "); err != nil {
		return fmt.Errorf("Erro ao exportar: %v", err)
	}

	var f *os.File
	if f, err = os.Create(l.arquivo); err != nil {
		return fmt.Errorf("Erro ao criar o arquivo %q: %v", l.arquivo, err)
	}

	if err = f.Sync(); err != nil {
		return fmt.Errorf("Erro ao sincronizar o arquivo %q: %v", l.arquivo, err)
	}

	if _, err = f.Write(data); err != nil {
		return fmt.Errorf("Erro escrever os dados no arquivo %q: %v", l.arquivo, err)
	}
	return
}
