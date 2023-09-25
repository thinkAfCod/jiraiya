package net

type ProtocolId interface {
	From(input string) ProtocolId
}
