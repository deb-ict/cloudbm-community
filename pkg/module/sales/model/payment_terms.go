package model

type PaymentTerms struct {
	Id           string
	Translations []*PaymentTermsTranslation
	Days         int32
	IsEnabled    bool
}

type PaymentTermsTranslation struct {
	Language    string
	Name        string
	Description string
}
