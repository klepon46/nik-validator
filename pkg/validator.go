package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/klepon46/nik-validator/internal/models"
	"golang.org/x/exp/slices"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

var (
	errInvalidLength = errors.New("NIK must be 16 characters")
	errInvalidNik    = errors.New("invalid NIK")
)

func (v *Validator) Validate(nik string) error {
	var kodeWilayah models.KodeWilayah

	if len(nik) != 16 {
		return errInvalidLength
	}

	filePath := filepath.Join("internal", "assets", "kode-wilayah.json")

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	err = json.Unmarshal(file, &kodeWilayah)
	if err != nil {
		return fmt.Errorf("error unmarshalling file: %w", err)
	}

	if !slices.Contains(kodeWilayah.Kecamatan, nik[:6]) {
		return errInvalidNik
	}

	return nil
}
