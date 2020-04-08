package models

// CoronaStates <model>
// is used to describe article model.
type CoronaStates struct {
	ID                  string
	TotalCases          string
	TotalDeaths         string
	TotalRecovered      string
	CurrentlyCases      string
	MildCases           string
	CriticalCases       string
	OutcomeCases        string
	RecoveredDischarged string
}
