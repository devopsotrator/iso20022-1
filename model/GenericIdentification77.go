package model

// Identification of an entity.
type GenericIdentification77 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType12Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType12Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3)
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification77) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification77) SetType(value string) {
	g.Type = (*PartyType12Code)(&value)
}

func (g *GenericIdentification77) SetIssuer(value string) {
	g.Issuer = (*PartyType12Code)(&value)
}

func (g *GenericIdentification77) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification77) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
