package services

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"

	"github.com/ramb0111/tax-calc/db"
	"github.com/ramb0111/tax-calc/src"
)

var dB *sql.DB

func init() {
	dB = db.PsqlConn()
}

// Function to get the Tax Id from db
func getTaxId(tx *sql.Tx, pl *src.Payload) (int64, error) {

	query := `
	SELECT
		id 
	FROM 
		tax
	where
		code = $1;
	`
	var taxId int64
	if err := tx.QueryRow(query, pl.TaxCode).Scan(&taxId); err != nil {
		return -1, err
	}
	return taxId, nil
}

// Function to insert the Item details and amount into db
func insertItem(tx *sql.Tx, pl *src.Payload, taxId int64, taxAmount float64) (int64, error) {
	query := `
	INSERT INTO
		item(name, amount, tax_amount, total_amount, tax_id)
	VALUES(
		$1, $2, $3, $4, $5
	) RETURNING
		id;
	`
	var itemId int64
	err := tx.QueryRow(query, pl.Name, pl.Amount, taxAmount, pl.Amount+taxAmount, taxId).Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	return itemId, nil
}

// Function to insert the user-item mapping into the db
func insertUserItemMapping(tx *sql.Tx, userId, itemId int64) error {
	query := `
	INSERT INTO
		user_item_mapping(user_id, item_id)
	VALUES(
		$1, $2
	);
	`
	if _, err := tx.Exec(query, userId, itemId); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// Wrapper function for all the insert functions
func Insert(userId int64, pl *src.Payload) (bool, error) {
	tx, err := dB.Begin()
	if err != nil {
		return false, err
	}

	taxAmount := pl.GetTaxAmount()
	if taxAmount == -1 {
		return false, errors.New("Invalid Tax Code")
	}

	taxId, err := getTaxId(tx, pl)
	if err != nil {
		return false, err
	}

	itemId, err := insertItem(tx, pl, taxId, taxAmount)
	if err != nil {
		return false, err
	}

	if err = insertUserItemMapping(tx, userId, itemId); err != nil {
		return false, err
	}

	tx.Commit()
	return true, nil
}
