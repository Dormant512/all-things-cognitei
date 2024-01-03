package util

import "fmt"

type DocWithNameExistsError struct {
	ItemName string
}

func (e DocWithNameExistsError) Error() string {
	return fmt.Sprintf("document with itemName %s already exists", e.ItemName)
}

type DocWithWithIdNotFoundError struct {
	Id string
}

func (e DocWithWithIdNotFoundError) Error() string {
	return fmt.Sprintf("document with id %s not found", e.Id)
}

type MegaItemTypeDoesNotExistError struct {
	InvalidType string
}

func (e MegaItemTypeDoesNotExistError) Error() string {
	return fmt.Sprintf("MegaItem object contains unknown itemType %s", e.InvalidType)
}

type MegaItemIsNotOfSaidTypeError struct {
	SaidType string
}

func (e MegaItemIsNotOfSaidTypeError) Error() string {
	return fmt.Sprintf("MegaItem object does not fit type %s criteria", e.SaidType)
}
