package model

// Completion of a securities settlement instruction, wherein securities are delivered/debited from a securities account and received/credited to the designated securities account.
type Transfer14 struct {

	// Date and time at which the transfer was received and processed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt"`

	// Date and time at which a transaction is completed and cleared, ie, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit3 `xml:"UnitsDtls,omitempty"`

	// Total quantity of securities settled.
	PortfolioTransferOutRate *PercentageRate `xml:"PrtflTrfOutRate,omitempty"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Value of a security, as booked in an account. Book value is often different from the current market value of the security.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer14) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer14) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer14) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer14) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *Transfer14) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer14) AddUnitsDetails() *Unit3 {
	newValue := new(Unit3)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
	return newValue
}

func (t *Transfer14) SetPortfolioTransferOutRate(value string) {
	t.PortfolioTransferOutRate = (*PercentageRate)(&value)
}

func (t *Transfer14) SetRounding(value string) {
	t.Rounding = (*RoundingDirection2Code)(&value)
}

func (t *Transfer14) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer14) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer14) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}
