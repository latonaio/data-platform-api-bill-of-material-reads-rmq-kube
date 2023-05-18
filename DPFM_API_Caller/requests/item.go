package requests

type Item struct {
	BillOfMaterial                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem              int      `json:"BillOfMaterialItem"`
	ComponentProduct                string   `json:"ComponentProduct"`
	ComponentProductBusinessPartner *int     `json:"ComponentProductBusinessPartner"`
	StockConfirmationPlant          *string  `json:"StockConfirmationPlant"`
	BOMAlternativeText              *string  `json:"BOMAlternativeText"`
	BOMItemBaseUnit                 *string  `json:"BOMItemBaseUnit"`
	BOMItemQuantityInBaseUnit       *float32 `json:"BOMItemQuantityInBaseUnit"`
	ComponentScrapInPercent         *int     `json:"ComponentScrapInPercent"`
	ValidityStartDate               *string  `json:"ValidityStartDate"`
	ValidityEndDate                 *string  `json:"ValidityEndDate"`
	BillOfMaterialItemText          *string  `json:"BillOfMaterialItemText"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion"`
}
