package models

type (
	GetListInventoryModel struct {
		ID            string `json:"id"`
		InventoryName string `json:"inventory_name"`
		InventoryVol  int    `json:"inventory_vol"`
		InventoryUnit string `json:"inventory_unit"`
	}
	GetListCalendarModel struct {
		ID             string `json:"id"`
		InventoryName  string `json:"inventory_name"`
		InventoryPrice int    `json:"inventory_price"`
		CreatedAt      string `json:"created_at"`
	}
	GetListContextModel struct {
		Context string `json:"context"`
	}
	GetInventoryByStorageModel struct {
		ID                string `json:"id"`
		InventoryName     string `json:"inventory_name"`
		InventoryVol      int    `json:"inventory_vol"`
		InventoryUnit     string `json:"inventory_unit"`
		InventoryCategory string `json:"inventory_category"`
		InventoryPrice    int    `json:"inventory_price"`
	}
	GetInventoryDetailModel struct {
		ID                    string `json:"id"`
		InventoryName         string `json:"inventory_name"`
		InventoryCategory     string `json:"inventory_category"`
		InventoryDesc         string `json:"inventory_desc"`
		InventoryMerk         string `json:"inventory_merk"`
		InventoryRoom         string `json:"inventory_room"`
		InventoryStorage      string `json:"inventory_storage"`
		InventoryRack         string `json:"inventory_rack"`
		InventoryPrice        int    `json:"inventory_price"`
		InventoryImage        string `json:"inventory_image"`
		InventoryUnit         string `json:"inventory_unit"`
		InventoryVol          int    `json:"inventory_vol"`
		InventoryCapacityUnit string `json:"inventory_capacity_unit"`
		InventoryCapacityVol  int    `json:"inventory_capacity_vol"`
		InventoryColor        string `json:"inventory_color"`
		IsFavorite            int    `json:"is_favorite"`
		IsReminder            int    `json:"is_reminder"`
		CreatedAt             string `json:"created_at"`
		CreatedBy             string `json:"created_by"`
		UpdatedAt             string `json:"updated_at"`
		DeletedAt             string `json:"deleted_at"`
	}
)
