//Package protein does protein translation.
package protein

var pr = map[string]string{"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

//FromCodon returns the polypeptide string
func FromCodon(unit string) string {
	return pr[unit]
}

//FromRNA returns the protein sequence as an []string
func FromRNA(seq string) []string {

	for i, rV := range seq {
		substr := ""

	}
}
