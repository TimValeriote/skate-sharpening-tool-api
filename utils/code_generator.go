package utils

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	coreModels "phl-skate-sharpening-api/core/models"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateString(length int) string {
	return StringWithCharset(length, charset)
}

func CreateNewStoreCodes(db *sql.DB) error {
	sql := `SELECT id, name, address, city, country, phone_number, store_number FROM store`
	query, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("Failed to prepare")
		return err
	}
	defer query.Close()

	stores, err := getAllStoresFromQuery(query)
	if err != nil {
		fmt.Println("Failed to get stores")
		return err
	}

	for _, store := range stores {
		newStoreCode := GenerateString(5)

		err := updateStoreCode(store.ID, newStoreCode, db)
		if err != nil {
			return err
		} else {
			fmt.Println(fmt.Sprintf("%s's code was update with: %s", store.Name, newStoreCode))
		}
	}

	return nil
}

func updateStoreCode(storeId int, code string, db *sql.DB) error {
	sql := `UPDATE daily_sharpening_code SET code = ? WHERE store_id = ?`

	_, err := db.Exec(sql, code, storeId)
	if err != nil {
		return err
	}

	return nil
}

func getAllStoresFromQuery(query *sql.Stmt) ([]coreModels.StoreStruct, error) {
	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stores := make([]coreModels.StoreStruct, 0)
	for rows.Next() {
		var store coreModels.StoreStruct
		err = rows.Scan(
			&store.ID,
			&store.Name,
			&store.Address,
			&store.City,
			&store.Country,
			&store.PhoneNumber,
			&store.StoreNumber,
		)
		if err != nil {
			return nil, err
		}

		stores = append(stores, store)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return stores, nil
}
