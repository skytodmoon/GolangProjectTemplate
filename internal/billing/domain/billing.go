//go:generate protoc -I. --go_out=Minternal/billing/domain/billing.proto=.:. --go_opt=paths=source_relative billing.proto

package domain
