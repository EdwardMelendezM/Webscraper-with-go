package main

import (
	"fmt"
	"strings"
)

type Signal struct {
	Name string
}

type RiskFactor struct {
	Name string
}

type Influence struct {
	Name string
}

type HarassmentType struct {
	Name string
}

type Synonym struct {
	Name string
}

type Method struct {
	Name string
}

type Prevention struct {
	Name string
}

type Ontology struct {
	Signals         map[string]*Signal
	RiskFactors     map[string]*RiskFactor
	Influences      map[string]*Influence
	HarassmentTypes map[string]*HarassmentType
	Synonyms        map[string]*Synonym
	Methods         map[string]*Method
	Preventions     map[string]*Prevention
	Associations    map[string][]*RiskFactor
}

func NewOntology() *Ontology {
	return &Ontology{
		Signals:         make(map[string]*Signal),
		RiskFactors:     make(map[string]*RiskFactor),
		Influences:      make(map[string]*Influence),
		HarassmentTypes: make(map[string]*HarassmentType),
		Synonyms:        make(map[string]*Synonym),
		Methods:         make(map[string]*Method),
		Preventions:     make(map[string]*Prevention),
		Associations:    make(map[string][]*RiskFactor),
	}
}

func (o *Ontology) AddSignal(name string) {
	o.Signals[name] = &Signal{Name: name}
}

func (o *Ontology) AddRiskFactor(name string) {
	o.RiskFactors[name] = &RiskFactor{Name: name}
}

func (o *Ontology) AddInfluence(name string) {
	o.Influences[name] = &Influence{Name: name}
}

func (o *Ontology) AddHarassmentType(name string) {
	o.HarassmentTypes[name] = &HarassmentType{Name: name}
}

func (o *Ontology) AddSynonym(name string) {
	o.Synonyms[name] = &Synonym{Name: name}
}

func (o *Ontology) AddMethod(name string) {
	o.Methods[name] = &Method{Name: name}
}

func (o *Ontology) AddPrevention(name string) {
	o.Preventions[name] = &Prevention{Name: name}
}

func (o *Ontology) AddAssociation(signalName, riskFactorName string) {
	_, ok := o.Signals[signalName]
	if !ok {
		return
	}
	riskFactor, ok := o.RiskFactors[riskFactorName]
	if !ok {
		return
	}
	o.Associations[signalName] = append(o.Associations[signalName], riskFactor)
}

func (o *Ontology) Classify(title, keyword string) []string {
	var result []string

	for signalName, riskFactors := range o.Associations {
		if strings.Contains(title, signalName) || strings.Contains(title, keyword) {
			for _, rf := range riskFactors {
				result = append(result, fmt.Sprintf("%s está asociado con %s", signalName, rf.Name))
			}
		}
	}
	return result
}

func main() {
	ontology := NewOntology()

	// Agregar señales de acoso
	ontology.AddSignal("Euforia")
	ontology.AddSignal("Deseo de no haber nacido")
	ontology.AddSignal("Me voy a suicidar")
	ontology.AddSignal("Suenio")
	ontology.AddSignal("Angustia")
	ontology.AddSignal("Hablar del acoso")
	ontology.AddSignal("Busqueda de afecto")
	ontology.AddSignal("Agresion a otros")
	ontology.AddSignal("Sin esperanza")
	ontology.AddSignal("Dificultad de aprendizaje")
	ontology.AddSignal("Notas de acoso")
	ontology.AddSignal("Acosar")
	ontology.AddSignal("Regalar pertenencias")
	ontology.AddSignal("Aislarse")
	ontology.AddSignal("Canbios de humor")
	ontology.AddSignal("Cambios corporales")
	ontology.AddSignal("Confusion")
	ontology.AddSignal("Miedo")
	ontology.AddSignal("Insertidumbre")

	// Agregar influencias
	ontology.AddInfluence("Amigos")
	ontology.AddInfluence("Familia")
	ontology.AddInfluence("Honor")
	ontology.AddInfluence("Religion")
	ontology.AddInfluence("Compañeros")

	// Agregar tipos de acoso
	ontology.AddHarassmentType("Acoso verbal")
	ontology.AddHarassmentType("Acoso físico")
	ontology.AddHarassmentType("Acoso sexual")
	ontology.AddHarassmentType("Acoso psicológico")
	ontology.AddHarassmentType("Acoso laboral")
	ontology.AddHarassmentType("Acoso escolar")
	ontology.AddHarassmentType("Ciberacoso")

	// Agregar sinónimos de acoso
	ontology.AddSynonym("Tocamientos indebidos")
	ontology.AddSynonym("Abuso")
	ontology.AddSynonym("Hostigamiento")
	ontology.AddSynonym("Acoso")

	// Agregar formas de llevar a cabo el acoso
	ontology.AddMethod("En el transporte")
	ontology.AddMethod("Redes sociales")
	ontology.AddMethod("Facebook")
	ontology.AddMethod("WhatsApp")
	ontology.AddMethod("Colegio")
	ontology.AddMethod("Escuela")
	ontology.AddMethod("Lugar de trabajo")

	// Agregar prevención
	ontology.AddPrevention("Ayuda telefónica")
	ontology.AddPrevention("Orientación psicológica")
	ontology.AddPrevention("Educación en el hogar")
	ontology.AddPrevention("Campañas de concienciación")

	// Agregar factores de riesgo
	ontology.AddRiskFactor("Pobreza")
	ontology.AddRiskFactor("Discriminación")
	ontology.AddRiskFactor("Violencia")
	ontology.AddRiskFactor("Desempleo")
	ontology.AddRiskFactor("Migración")
	ontology.AddRiskFactor("Indigencia")
	ontology.AddRiskFactor("Intento de acoso")
	ontology.AddRiskFactor("Abuso de sustancias")
	ontology.AddRiskFactor("Drogas, cocaína, cannabis")
	ontology.AddRiskFactor("Menosprecio")
	ontology.AddRiskFactor("Humillación")
	ontology.AddRiskFactor("Alcoholismo")
	ontology.AddRiskFactor("Autoagresión")
	ontology.AddRiskFactor("Decepción amorosa")
	ontology.AddRiskFactor("Esquizofrenia")
	ontology.AddRiskFactor("Trastorno bipolar")
	ontology.AddRiskFactor("Trastorno mental")
	ontology.AddRiskFactor("Estrés postraumático")
	ontology.AddRiskFactor("Depresión")
	ontology.AddRiskFactor("Soledad")
	ontology.AddRiskFactor("Acoso psicológico")
	ontology.AddRiskFactor("Sufrimiento psicológico")
	ontology.AddRiskFactor("Problemas de relaciones interpersonales")
	ontology.AddRiskFactor("Abuso sexual")
	ontology.AddRiskFactor("Acoso escolar")
	ontology.AddRiskFactor("Maltrato familiar")
	ontology.AddRiskFactor("Dificultad financiera")

	// Agregar asociaciones entre señales y factores de riesgo
	ontology.AddAssociation("Euforia", "Depresión")
	ontology.AddAssociation("Deseo de no haber nacido", "Autoagresión")
	ontology.AddAssociation("Me voy a suicidar", "Decepción amorosa")
	ontology.AddAssociation("Angustia", "Acoso escolar")
	ontology.AddAssociation("Aislarse", "Soledad")

	// Clasificar
	title := "Ayuda como metodo de huida de acoso"
	keyword := "Acoso"
	results := ontology.Classify(title, keyword)

	for _, result := range results {
		fmt.Println(result)
	}
}
