package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason26Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason24Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason26Choice) SetCode(value string) {
	r.Code = (*RejectionReason24Code)(&value)
}

func (r *RejectionAndRepairReason26Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
