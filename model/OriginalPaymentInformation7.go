package model

// Provides information on the original payment transaction, to which the remittance message applies.
type OriginalPaymentInformation7 struct {

	// Identifies the underlying transaction.
	References *TransactionReferences4 `xml:"Refs"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt,omitempty"`

	// Provides details of the currency exchange rate and contract.
	ExchangeRateInformation *ExchangeRate1 `xml:"XchgRateInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`
}

func (o *OriginalPaymentInformation7) AddReferences() *TransactionReferences4 {
	o.References = new(TransactionReferences4)
	return o.References
}

func (o *OriginalPaymentInformation7) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	o.PaymentTypeInformation = new(PaymentTypeInformation19)
	return o.PaymentTypeInformation
}

func (o *OriginalPaymentInformation7) AddAmount() *AmountType3Choice {
	o.Amount = new(AmountType3Choice)
	return o.Amount
}

func (o *OriginalPaymentInformation7) AddExchangeRateInformation() *ExchangeRate1 {
	o.ExchangeRateInformation = new(ExchangeRate1)
	return o.ExchangeRateInformation
}

func (o *OriginalPaymentInformation7) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	o.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return o.RequestedExecutionDate
}

func (o *OriginalPaymentInformation7) SetRequestedCollectionDate(value string) {
	o.RequestedCollectionDate = (*ISODate)(&value)
}

func (o *OriginalPaymentInformation7) AddDebtor() *PartyIdentification43 {
	o.Debtor = new(PartyIdentification43)
	return o.Debtor
}

func (o *OriginalPaymentInformation7) AddDebtorAccount() *CashAccount24 {
	o.DebtorAccount = new(CashAccount24)
	return o.DebtorAccount
}

func (o *OriginalPaymentInformation7) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalPaymentInformation7) AddCreditor() *PartyIdentification43 {
	o.Creditor = new(PartyIdentification43)
	return o.Creditor
}

func (o *OriginalPaymentInformation7) AddCreditorAccount() *CashAccount24 {
	o.CreditorAccount = new(CashAccount24)
	return o.CreditorAccount
}

func (o *OriginalPaymentInformation7) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.CreditorAgent
}
