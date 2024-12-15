package files

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/anatolygg/tg_bot/internal/storage"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = fmt.Errorf("error: %w", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err := os.Mkdir(filePath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath := filepath.Join(filePath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err = gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
