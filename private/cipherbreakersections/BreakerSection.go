package cipherbreakersections

import (
	"github.com/titosilva/cipherbreaker-go/internal/cipherbreakershared/cbsharedviews"
	"github.com/titosilva/cipherbreaker-go/pkg/tui/section"
)

// BreakerSection struct
type BreakerSection struct {
	CipherType int
	InputMode  int
	FilePath   string
	CipherText string
	Mode       int
}

// Run method of BreakerSection
func (bs BreakerSection) Run() section.Section {
	if bs.Mode == 0 {
		cancelled, selection := cbsharedviews.MultipleSelectionView("Select a cipher algorithm: ", []string{"Caesar", "Vigenere", "Exit"})

		if selection == 2 || cancelled {
			return MainSection{}
		}

		return BreakerSection{CipherType: selection, Mode: 1}
	} else if bs.Mode == 1 {
		cancelled, selection := cbsharedviews.MultipleSelectionView("Select a input mode: ", []string{"From file", "Type", "Exit"})

		if selection == 2 || cancelled {
			return MainSection{}
		}

		return BreakerSection{CipherType: bs.CipherType, InputMode: selection, Mode: 2}
	} else if bs.Mode == 2 && bs.InputMode == 0 {
		cancelled, cancelButton, doneButton, filePath := cbsharedviews.OneInputView("Type the file path: ", "Path: ", "Cancel", "Done")

		if cancelled || cancelButton {
			return MainSection{}
		}

		if doneButton {
			return BreakerSection{CipherType: bs.CipherType, InputMode: bs.InputMode, Mode: 3, FilePath: filePath}
		}

		return MainSection{}
	} else if bs.Mode == 2 && bs.InputMode == 1 {
		cancelled, cancelButton, doneButton, cipherText := cbsharedviews.OneInputView("Type the ciphertext: ", "Ciphertext: ", "Cancel", "Done")

		if cancelled || cancelButton {
			return MainSection{}
		}

		if doneButton {
			return BreakerSection{CipherType: bs.CipherType, InputMode: bs.InputMode, Mode: 3, CipherText: cipherText}
		}

		return MainSection{}
	} else if bs.Mode == 3 && bs.InputMode == 0 {

	} else if bs.Mode == 3 && bs.InputMode == 1 {

	}

	return MainSection{}
}

// IsEnd method of BreakerSection
func (bs BreakerSection) IsEnd() bool {
	return false
}
