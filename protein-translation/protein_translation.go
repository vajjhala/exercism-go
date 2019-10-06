//Package protein does protein translation.
package protein

import (
	"errors"
)

//ErrStop when the codon is a STOP codon
var ErrStop = errors.New("stop")

//ErrInvalidBase is an invalid base codon
var ErrInvalidBase = errors.New("Invalid codon")

var pr = map[string]string{"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

//FromCodon returns the polypeptide string
func FromCodon(codon string) (string, error) {
	p, ok := pr[codon]
	if !ok {
		return p, ErrInvalidBase
	}
	if p == "STOP" {
		return "", ErrStop
	}
	return p, nil
}

//FromRNA returns the protein sequence as an []string
func FromRNA(rna string) ([]string, error) {
	var prtns = make([]string, 0, (len(rna)/3)+1)
	var er error
	var cdn string
	for i := 0; i < len(rna); i += 3 {
		cdn = rna[i : i+3]
		p, ok := FromCodon(cdn)
		if ok == ErrInvalidBase {
			er = ok
			break
		}
		if ok == ErrStop {
			er = nil
			break
		}
		er = nil
		prtns = append(prtns, p)
	}
	return prtns, er
}
