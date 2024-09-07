package domain

type OntologyClass struct {
	Name    string
	Objects []OntologyObject
}

type OntologyObject struct {
	Name       string
	Attributes []string
}
