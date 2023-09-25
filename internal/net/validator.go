package net

type Validator[K ContentKey] interface {
	ValidateContent(key K, value []byte) error
}
