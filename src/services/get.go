package services

import (
	"log"
	"strings"

	"github.com/ramb0111/tax-calc/src"
	"github.com/ramb0111/tax-calc/src/models"
)

func getItemIds(userId int64) []int64 {

	query := `
	SELECT
		item_id 
	FROM 
		user_item_mapping
	where
		user_id = $1;
	`
	var (
		itemIds []int64
		itemId  int64
	)

	rows, err := dB.Query(query, userId)
	if err != nil {
		log.Println(err)
		return itemIds
	}

	for rows.Next() {
		if err = rows.Scan(&itemId); err == nil {
			itemIds = append(itemIds, itemId)
		} else {
			log.Println(err)
		}
	}
	return itemIds
}

func getResult(itemIds []int64) models.Result {
	query := `
	SELECT
		name, amount, tax_amount, total_amount, tax.code AS tax_code, tax.type as tax_type 
	FROM 
		item, tax 
	where
		item.id = ANY($1::int[])
	AND
		item.tax_id = tax.id;
	`

	var (
		items  = []models.Item{}
		item   models.Item
		result = models.Result{Items: items}
	)
	params := "{" + strings.Join(src.ArrayIntToArrayStr(itemIds), ",") + "}"
	log.Println(params)

	rows, err := dB.Query(query, params)
	if err != nil {
		log.Println(err)
		return result
	}

	for rows.Next() {
		err = rows.Scan(
			&item.Name,
			&item.Amount,
			&item.TaxAmount,
			&item.TotalAmount,
			&item.TaxCode,
			&item.Type,
		)

		if err == nil {
			result.TotalAmount += item.Amount
			result.TotalTaxAmount += item.TaxAmount
			items = append(items, item)
		} else {
			log.Println(err)
		}
	}
	result.GrandTotal = result.TotalAmount + result.TotalTaxAmount
	result.Items = items
	return result
}

// Function to get all the items and their details for a userId
func GetAll(userId int64) models.Result {
	itemIds := getItemIds(userId)
	log.Println(itemIds, "items")
	return getResult(itemIds)
}
