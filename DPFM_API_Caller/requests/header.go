package requests

type Header struct {
	BillOfMaterial              int      `json:"BillOfMaterial"`
	BillOfMaterialType          *string  `json:"BillOfMaterialType"`
	Product                     *string  `json:"Product"`
	OwnerBusinessPartner        *int     `json:"OwnerBusinessPartner"`
	OwnerPlant                  *string  `json:"OwnerPlant"`
	BOMAlternativeText          *string  `json:"BOMAlternativeText"`
	BOMHeaderBaseUnit           *string  `json:"BOMHeaderBaseUnit"`
	BOMHeaderQuantityInBaseUnit *float32 `json:"BOMHeaderQuantityInBaseUnit"`
	ValidityStartDate           *string  `json:"ValidityStartDate"`
	ValidityEndDate             *string  `json:"ValidityEndDate"`
	CreationDate                string   `json:"CreationDate"`
	LastChangeDate              string   `json:"LastChangeDate"`
	BillOfMaterialHeaderText    *string  `json:"BillOfMaterialHeaderText"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion"`
}
