package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Header            Header   `json:"BillOfMaterial"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

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
	CreationDate                *string  `json:"CreationDate"`
	LastChangeDate              *string  `json:"LastChangeDate"`
	BillOfMaterialHeaderText    *string  `json:"BillOfMaterialHeaderText"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion"`
	Item                        []Item   `json:"Item"`
}

type Item struct {
	BillOfMaterial                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem              int      `json:"BillOfMaterialItem"`
	ComponentProduct                *string  `json:"ComponentProduct"`
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
