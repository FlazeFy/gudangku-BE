package models

type (
	PostReportModel struct {
		ReportTitle    string  `json:"report_title"`
		ReportCategory string  `json:"report_category"`
		ReportDesc     *string `json:"report_desc"`
		ReportItem     *string `json:"report_item"`
		ReportImage    *string `json:"report_image"`
		IsReminder     int     `json:"is_reminder"`
		RemindAt       *string `json:"remind_at"`
	}
	ReportItemModel struct {
		InventoryID *string `json:"inventory_id"`
		ItemName    string  `json:"item_name"`
		ItemDesc    string  `json:"item_desc"`
		ItemQty     int     `json:"item_qty"`
		ItemPrice   int     `json:"item_price"`
	}
)
