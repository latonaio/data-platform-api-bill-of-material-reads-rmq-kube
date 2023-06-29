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
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Product,
			&pm.Buyer,
			&pm.Seller,
			&pm.DestinationDeliverToParty,
			&pm.DestinationDeliverToPlant,
			&pm.DepartureDeliverFromParty,
			&pm.DepartureDeliverFromPlant,
			&pm.OwnerProductionPlantBusinessPartner,
			&pm.OwnerProductionPlant,
			&pm.ProductBaseUnit,
			&pm.ProductDeliveryUnit,
			&pm.ProductProductionUnit,
			&pm.ProductStandardQuantityInBaseUnit,
			&pm.ProductStandardQuantityInDeliveryUnit,
			&pm.ProductStandardQuantityInProductionUnit,
			&pm.BillOfMaterialHeaderText,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			BillOfMaterial:                           data.BillOfMaterial,
			BillOfMaterialType:                       data.BillOfMaterialType,
			SupplyChainRelationshipID:                data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:        data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:   data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipProductionPlantID: data.SupplyChainRelationshipProductionPlantID,
			Product:                                  data.Product,
			Buyer:                                    data.Buyer,
			Seller:                                   data.Seller,
			DestinationDeliverToParty:                data.DestinationDeliverToParty,
			DestinationDeliverToPlant:                data.DestinationDeliverToPlant,
			DepartureDeliverFromParty:                data.DepartureDeliverFromParty,
			DepartureDeliverFromPlant:                data.DepartureDeliverFromPlant,
			OwnerProductionPlantBusinessPartner:      data.OwnerProductionPlantBusinessPartner,
			OwnerProductionPlant:                     data.OwnerProductionPlant,
			ProductBaseUnit:                          data.ProductBaseUnit,
			ProductDeliveryUnit:                      data.ProductDeliveryUnit,
			ProductProductionUnit:                    data.ProductProductionUnit,
			ProductStandardQuantityInBaseUnit:        data.ProductStandardQuantityInBaseUnit,
			ProductStandardQuantityInDeliveryUnit:    data.ProductStandardQuantityInDeliveryUnit,
			ProductStandardQuantityInProductionUnit:  data.ProductStandardQuantityInProductionUnit,
			BillOfMaterialHeaderText:                 data.BillOfMaterialHeaderText,
			ValidityStartDate:                        data.ValidityStartDate,
			ValidityEndDate:                          data.ValidityEndDate,
			CreationDate:                             data.CreationDate,
			LastChangeDate:                           data.LastChangeDate,
			IsMarkedForDeletion:                      data.IsMarkedForDeletion,
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
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.Product,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.ComponentProductStandardQuantityInBaseUnuit,
			&pm.ComponentProductStandardQuantityInDeliveryUnuit,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.ComponentProductStandardScrapInPercent,
			&pm.IsMarkedForBackflush,
			&pm.BillOfMaterialItemText,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			BillOfMaterial:                                  data.BillOfMaterial,
			BillOfMaterialItem:                              data.BillOfMaterialItem,
			SupplyChainRelationshipID:                       data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:               data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:          data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:         data.SupplyChainRelationshipStockConfPlantID,
			Product:                                         data.Product,
			ProductionPlantBusinessPartner:                  data.ProductionPlantBusinessPartner,
			ProductionPlant:                                 data.ProductionPlant,
			ComponentProduct:                                data.ComponentProduct,
			ComponentProductBuyer:                           data.ComponentProductBuyer,
			ComponentProductSeller:                          data.ComponentProductSeller,
			ComponentProductDeliverToParty:                  data.ComponentProductDeliverToParty,
			ComponentProductDeliverToPlant:                  data.ComponentProductDeliverToPlant,
			ComponentProductDeliverFromParty:                data.ComponentProductDeliverFromParty,
			ComponentProductDeliverFromPlant:                data.ComponentProductDeliverFromPlant,
			StockConfirmationBusinessPartner:                data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                          data.StockConfirmationPlant,
			ComponentProductStandardQuantityInBaseUnuit:     data.ComponentProductStandardQuantityInBaseUnuit,
			ComponentProductStandardQuantityInDeliveryUnuit: data.ComponentProductStandardQuantityInDeliveryUnuit,
			ComponentProductBaseUnit:                        data.ComponentProductBaseUnit,
			ComponentProductDeliveryUnit:                    data.ComponentProductDeliveryUnit,
			ComponentProductStandardScrapInPercent:          data.ComponentProductStandardScrapInPercent,
			IsMarkedForBackflush:                            data.IsMarkedForBackflush,
			BillOfMaterialItemText:                          data.BillOfMaterialItemText,
			ValidityStartDate:                               data.ValidityStartDate,
			ValidityEndDate:                                 data.ValidityEndDate,
			CreationDate:                                    data.CreationDate,
			LastChangeDate:                                  data.LastChangeDate,
			IsMarkedForDeletion:                             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
