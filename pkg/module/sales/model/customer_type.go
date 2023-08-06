package model

type CustomerType uint8

const (
	CustomerType_Undefined CustomerType = iota
	CustomerType_Business
	CustomerType_Private
)
