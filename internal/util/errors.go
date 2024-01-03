package util

import "fmt"

type DocWithNameExistsError struct {
	ItemName string
}

func (e DocWithNameExistsError) Error() string {
	return fmt.Sprintf("document with itemName %s already exists", e.ItemName)
}

type InvalidMegaItemTypeError struct {
	InvalidType string
}

func (e InvalidMegaItemTypeError) Error() string {
	return fmt.Sprintf("MegaItem object contains unknown itemType %s", e.InvalidType)
}

type InvalidMegaItemError struct {
	InvalidType string
}

func (e InvalidMegaItemError) Error() string {
	return "MegaItem object does not fit item criteria"
}
