package net

import (
	"github.com/optimism-java/jiraiya/internal/api"
)

//go:generate mockgen -destination mock/storage_mock.go -package mock -source storage.go ContentKey,ContentStore

type ContentKey interface {
	ContentId() [32]byte

	Bytes() []byte

	Hex() string
}

type ContentStore[K ContentKey, V []byte] interface {
	Get(key K) (V, error)

	Put(key K, value V) error

	IsKeyWithinRadiusAndUnavailable(key K) (bool, error)

	Radius() *api.Distance
}
