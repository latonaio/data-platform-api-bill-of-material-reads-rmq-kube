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
			&pm.ComponentProductStandardQuantityInBaseUnit,
			&pm.ComponentProductStandardQuantityInDeliveryUnit,
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
			BillOfMaterial:                                 data.BillOfMaterial,
			BillOfMaterialItem:                             data.BillOfMaterialItem,
			SupplyChainRelationshipID:                      data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:              data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:         data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:        data.SupplyChainRelationshipStockConfPlantID,
			Product:                                        data.Product,
			ProductionPlantBusinessPartner:                 data.ProductionPlantBusinessPartner,
			ProductionPlant:                                data.ProductionPlant,
			ComponentProduct:                               data.ComponentProduct,
			ComponentProductBuyer:                          data.ComponentProductBuyer,
			ComponentProductSeller:                         data.ComponentProductSeller,
			ComponentProductDeliverToParty:                 data.ComponentProductDeliverToParty,
			ComponentProductDeliverToPlant:                 data.ComponentProductDeliverToPlant,
			ComponentProductDeliverFromParty:               data.ComponentProductDeliverFromParty,
			ComponentProductDeliverFromPlant:               data.ComponentProductDeliverFromPlant,
			StockConfirmationBusinessPartner:               data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:                         data.StockConfirmationPlant,
			ComponentProductStandardQuantityInBaseUnit:     data.ComponentProductStandardQuantityInBaseUnit,
			ComponentProductStandardQuantityInDeliveryUnit: data.ComponentProductStandardQuantityInDeliveryUnit,
			ComponentProductBaseUnit:                       data.ComponentProductBaseUnit,
			ComponentProductDeliveryUnit:                   data.ComponentProductDeliveryUnit,
			ComponentProductStandardScrapInPercent:         data.ComponentProductStandardScrapInPercent,
			IsMarkedForBackflush:                           data.IsMarkedForBackflush,
			BillOfMaterialItemText:                         data.BillOfMaterialItemText,
			ValidityStartDate:                              data.ValidityStartDate,
			ValidityEndDate:                                data.ValidityEndDate,
			CreationDate:                                   data.CreationDate,
			LastChangeDate:                                 data.LastChangeDate,
			IsMarkedForDeletion:                            data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}

func ConvertToItemPricingElement(rows *sql.Rows) (*[]ItemPricingElement, error) {
	defer rows.Close()
	itemPricingElement := make([]ItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPricingElement{}

		err := rows.Scan(
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.PricingProcedureCounter,
			&pm.SupplyChainRelationshipID,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.TaxCode,
			&pm.ConditionAmount,
			&pm.TransactionCurrency,
			&pm.ConditionIsManuallyChanged,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemPricingElement, err
		}

		data := pm
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			BillOfMaterial:             data.BillOfMaterial,
			BillOfMaterialItem:         data.BillOfMaterialItem,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			ComponentProductBuyer:      data.ComponentProductBuyer,
			ComponentProductSeller:     data.ComponentProductSeller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			TransactionCurrency:        data.TransactionCurrency,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
			CreationDate:               data.CreationDate,
			CreationTime:               data.CreationTime,
			LastChangeDate:             data.LastChangeDate,
			LastChangeTime:             data.LastChangeTime,
			IsCancelled:                data.IsCancelled,
			IsMarkedForDeletion:        data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemPricingElement, nil
	}

	return &itemPricingElement, nil
}
