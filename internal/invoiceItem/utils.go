package invoiceItem

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"
)

func BatchUpdateItems(curItems []entity.Item, modifiedItems []invoiceDto.PutItemDto) {
	// !consider refactoring the search coz both slices are sorted
	for modIdx := range modifiedItems {
		for curIdx := range curItems {
			if curItems[curIdx].ID == modifiedItems[modIdx].ID {
				curItems[curIdx].Name = modifiedItems[modIdx].Name
				curItems[curIdx].Quantity = modifiedItems[modIdx].Quantity
				curItems[curIdx].Price = modifiedItems[modIdx].Price
				curItems[curIdx].Total = modifiedItems[modIdx].Total
			}
		}
	}

}

func BatchDeleteItems(curItems *[]entity.Item, deleteItemIds []uint64) []entity.Item {
	var deleteItems []entity.Item

	// !consider refactoring the search coz both slices are sorted?
	for _, deleteId := range deleteItemIds {
		newLength := 0
		for _, curItem := range *curItems {
			if curItem.ID != deleteId {
				(*curItems)[newLength] = curItem
				newLength++
			} else {
				deleteItems = append(deleteItems, curItem)
			}
		}
		(*curItems) = (*curItems)[:newLength]
	}

	return deleteItems
}
