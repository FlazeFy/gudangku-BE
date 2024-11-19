package models

type (
	GetListInventoryModel struct {
		ID            string `json:"id"`
		InventoryName string `json:"inventory_name"`
		InventoryVol  int    `json:"inventory_vol"`
		InventoryUnit string `json:"inventory_unit"`
	}
)
