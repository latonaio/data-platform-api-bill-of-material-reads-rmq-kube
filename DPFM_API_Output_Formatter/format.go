package dpfm_api_output_formatter

import (
	"data-platform-api-bill-of-material-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.BillOfMaterial,
			&pm.BillOfMaterialType,
			&pm.Product,
			&pm.OwnerBusinessPartner,
			&pm.OwnerPlant,
			&pm.BOMAlternativeText,
			&pm.BOMHeaderBaseUnit,
			&pm.BOMHeaderQuantityInBaseUnit,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.BillOfMaterialHeaderText,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			BillOfMaterial:              data.BillOfMaterial,
			BillOfMaterialType:          data.BillOfMaterialType,
			Product:                     data.Product,
			OwnerBusinessPartner:        data.OwnerBusinessPartner,
			OwnerPlant:                  data.OwnerPlant,
			BOMAlternativeText:          data.BOMAlternativeText,
			BOMHeaderBaseUnit:           data.BOMHeaderBaseUnit,
			BOMHeaderQuantityInBaseUnit: data.BOMHeaderQuantityInBaseUnit,
			ValidityStartDate:           data.ValidityStartDate,
			ValidityEndDate:             data.ValidityEndDate,
			CreationDate:                data.CreationDate,
			LastChangeDate:              data.LastChangeDate,
			BillOfMaterialHeaderText:    data.BillOfMaterialHeaderText,
			IsMarkedForDeletion:         data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ComponentProduct,
			&pm.ComponentProductBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.BOMAlternativeText,
			&pm.BOMItemBaseUnit,
			&pm.BOMItemQuantityInBaseUnit,
			&pm.ComponentScrapInPercent,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.BillOfMaterialItemText,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			BillOfMaterial:                  data.BillOfMaterial,
			BillOfMaterialItem:              data.BillOfMaterialItem,
			ComponentProduct:                data.ComponentProduct,
			ComponentProductBusinessPartner: data.ComponentProductBusinessPartner,
			StockConfirmationPlant:          data.StockConfirmationPlant,
			BOMAlternativeText:              data.BOMAlternativeText,
			BOMItemBaseUnit:                 data.BOMItemBaseUnit,
			BOMItemQuantityInBaseUnit:       data.BOMItemQuantityInBaseUnit,
			ComponentScrapInPercent:         data.ComponentScrapInPercent,
			ValidityStartDate:               data.ValidityStartDate,
			ValidityEndDate:                 data.ValidityEndDate,
			BillOfMaterialItemText:          data.BillOfMaterialItemText,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
