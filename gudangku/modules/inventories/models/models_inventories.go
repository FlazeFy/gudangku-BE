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
)
