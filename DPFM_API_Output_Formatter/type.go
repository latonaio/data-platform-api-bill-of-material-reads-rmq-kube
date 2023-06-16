package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header *[]Header `json:"Header"`
	Item   *[]Item   `json:"Item"`
}

type Header struct {
	BillOfMaterial                           int      `json:"BillOfMaterial"`
	BillOfMaterialType                       *string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  *string  `json:"Product"`
	Buyer                                    *int     `json:"Buyer"`
	Seller                                   *int     `json:"Seller"`
	DepartureDeliverFromParty                *int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                *string  `json:"DepartureDeliverFromPlant"`
	DestinationDeliverToParty                *int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                *string  `json:"DestinationDeliverToPlant"`
	OwnerProductionBusinessPartner           *int     `json:"OwnerProductionBusinessPartner"`
	OwnerProductionPlant                     *string  `json:"OwnerProductionPlant"`
	ProductBaseUnit                          *string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      *string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    *string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit        *float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    *float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  *float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                 *string  `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string  `json:"ValidityStartDate"`
	ValidityEndDate                          *string  `json:"ValidityEndDate"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}

type Item struct {
	BillOfMaterial                                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem                              int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                       *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID         *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         *string  `json:"Product"`
	ProductionPlantBusinessPartner                  *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 *string  `json:"ProductionPlant"`
	ComponentProduct                                *string  `json:"ComponentProduct"`
	ComponentProductBuyer                           *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                          *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverFromParty                *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductDeliverToParty                  *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                  *string  `json:"ComponentProductDeliverToPlant"`
	StockConfirmationBusinessPartner                *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                          *string  `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnuit     *float32 `json:"ComponentProductStandardQuantityInBaseUnuit"`
	ComponentProductStandardQuantityInDeliveryUnuit *float32 `json:"ComponentProductStandardQuantityInDeliveryUnuit"`
	ComponentProductBaseUnit                        *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                    *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
	CreationDate                                    *string  `json:"CreationDate"`
	LastChangeDate                                  *string  `json:"LastChangeDate"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}
