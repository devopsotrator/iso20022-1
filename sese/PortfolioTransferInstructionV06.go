package sese

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200106 struct {
	XMLName xml.Name                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.012.001.06 Document"`
	Message *PortfolioTransferInstructionV06 `xml:"PrtflTrfInstr"`
}

func (d *Document01200106) AddMessage() *PortfolioTransferInstructionV06 {
	d.Message = new(PortfolioTransferInstructionV06)
	return d.Message
}

// Scope
// An instructing party, for example, a (new) plan manager (Transferee), sends the PortfolioTransferInstruction message to the executing party, for example, a (old) plan manager (Transferor), on behalf of the initiating party, for example, an investor (client), to instruct the transfer of financial instruments from the clients account at the old plan manager (Transferor) to the clients account at the new plan manager (Transferee) through a nominee account.
// Usage
// The PortfolioTransferInstruction message is used to instruct the withdrawal of one or more ISA or portfolio products from one account and deliver them to another account.
// The PortfolioTransferInstruction message is used to instruct one or more transfers for one client. Each transfer is for delivery to the same account. The account may be owned by one or more individual investors or one or more corporate investors. Each transfer is identified in TransferIdentification.
// If the instructing party does not have enough information to instruct the transfer, then it must first send a AccountHoldingInformationRequest message to the executing party in order to receive a AccountHoldingInformation message.
type PortfolioTransferInstructionV06 struct {

	// Identifies the message.
	MessageReference *model.MessageIdentification1 `xml:"MsgRef"`

	// Collective reference identifying a set of messages.
	PoolReference *model.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *model.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *model.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Information identifying the primary individual investor, for example, name, address, social security number and date of birth.
	PrimaryIndividualInvestor *model.IndividualPerson8 `xml:"PmryIndvInvstr,omitempty"`

	// Information identifying the secondary individual investor, for example, name, address, social security number and date of birth.
	SecondaryIndividualInvestor *model.IndividualPerson8 `xml:"ScndryIndvInvstr,omitempty"`

	// Information identifying the other individual investors, for example, name, address, social security number and date of birth.
	OtherIndividualInvestor []*model.IndividualPerson8 `xml:"OthrIndvInvstr,omitempty"`

	// Information identifying the primary corporate investor, for example, name and address.
	PrimaryCorporateInvestor *model.Organisation4 `xml:"PmryCorpInvstr,omitempty"`

	// Information identifying the secondary corporate investor, for example, name and address.
	SecondaryCorporateInvestor *model.Organisation4 `xml:"ScndryCorpInvstr,omitempty"`

	// Information identifying the other corporate investors, for example, name and address.
	OtherCorporateInvestor []*model.Organisation4 `xml:"OthrCorpInvstr,omitempty"`

	// Identification of an account owned by the investor at the old plan manager (account servicer).
	TransferorAccount *model.Account15 `xml:"TrfrAcct"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	NomineeAccount *model.Account16 `xml:"NmneeAcct,omitempty"`

	// Information related to the institution to which the financial instrument is to be transferred.
	Transferee *model.PartyIdentification2Choice `xml:"Trfee"`

	// Identification of an account owned by the investor to which a cash entry is made based on the transfer of asset(s).
	CashAccount *model.CashAccount29 `xml:"CshAcct,omitempty"`

	// Provides information related to the asset(s) transferred.
	ProductTransfer []*model.ISATransfer18 `xml:"PdctTrf"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *model.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*model.Extension1 `xml:"Xtnsn,omitempty"`
}

func (p *PortfolioTransferInstructionV06) AddMessageReference() *model.MessageIdentification1 {
	p.MessageReference = new(model.MessageIdentification1)
	return p.MessageReference
}

func (p *PortfolioTransferInstructionV06) AddPoolReference() *model.AdditionalReference3 {
	p.PoolReference = new(model.AdditionalReference3)
	return p.PoolReference
}

func (p *PortfolioTransferInstructionV06) AddPreviousReference() *model.AdditionalReference3 {
	p.PreviousReference = new(model.AdditionalReference3)
	return p.PreviousReference
}

func (p *PortfolioTransferInstructionV06) AddRelatedReference() *model.AdditionalReference3 {
	p.RelatedReference = new(model.AdditionalReference3)
	return p.RelatedReference
}

func (p *PortfolioTransferInstructionV06) AddPrimaryIndividualInvestor() *model.IndividualPerson8 {
	p.PrimaryIndividualInvestor = new(model.IndividualPerson8)
	return p.PrimaryIndividualInvestor
}

func (p *PortfolioTransferInstructionV06) AddSecondaryIndividualInvestor() *model.IndividualPerson8 {
	p.SecondaryIndividualInvestor = new(model.IndividualPerson8)
	return p.SecondaryIndividualInvestor
}

func (p *PortfolioTransferInstructionV06) AddOtherIndividualInvestor() *model.IndividualPerson8 {
	newValue := new(model.IndividualPerson8)
	p.OtherIndividualInvestor = append(p.OtherIndividualInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV06) AddPrimaryCorporateInvestor() *model.Organisation4 {
	p.PrimaryCorporateInvestor = new(model.Organisation4)
	return p.PrimaryCorporateInvestor
}

func (p *PortfolioTransferInstructionV06) AddSecondaryCorporateInvestor() *model.Organisation4 {
	p.SecondaryCorporateInvestor = new(model.Organisation4)
	return p.SecondaryCorporateInvestor
}

func (p *PortfolioTransferInstructionV06) AddOtherCorporateInvestor() *model.Organisation4 {
	newValue := new(model.Organisation4)
	p.OtherCorporateInvestor = append(p.OtherCorporateInvestor, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV06) AddTransferorAccount() *model.Account15 {
	p.TransferorAccount = new(model.Account15)
	return p.TransferorAccount
}

func (p *PortfolioTransferInstructionV06) AddNomineeAccount() *model.Account16 {
	p.NomineeAccount = new(model.Account16)
	return p.NomineeAccount
}

func (p *PortfolioTransferInstructionV06) AddTransferee() *model.PartyIdentification2Choice {
	p.Transferee = new(model.PartyIdentification2Choice)
	return p.Transferee
}

func (p *PortfolioTransferInstructionV06) AddCashAccount() *model.CashAccount29 {
	p.CashAccount = new(model.CashAccount29)
	return p.CashAccount
}

func (p *PortfolioTransferInstructionV06) AddProductTransfer() *model.ISATransfer18 {
	newValue := new(model.ISATransfer18)
	p.ProductTransfer = append(p.ProductTransfer, newValue)
	return newValue
}

func (p *PortfolioTransferInstructionV06) AddMarketPracticeVersion() *model.MarketPracticeVersion1 {
	p.MarketPracticeVersion = new(model.MarketPracticeVersion1)
	return p.MarketPracticeVersion
}

func (p *PortfolioTransferInstructionV06) AddExtension() *model.Extension1 {
	newValue := new(model.Extension1)
	p.Extension = append(p.Extension, newValue)
	return newValue
}
