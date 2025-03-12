package models

type DeviceCatalogue struct {
	Slno              int    `json:"slno" gorm:"column:slno"`
	SkuCode           string `json:"skucode" gorm:"column:skucode" binding:"required"`
	ProductName       string `json:"productname" gorm:"column:productname"`
	Description       string `json:"description" gorm:"column:description"`
	InternalBattery   string `json:"internalbattery" gorm:"column:internalbattery"`
	Specification     string `json:"specification" gorm:"column:specification"`
	CompatibleNations string `json:"compatiblenations" gorm:"column:compatiblenations"`
	NumberofInputs    string `json:"numberofinputs" gorm:"column:numberofinputs"`
	NumberofOutputs   string `json:"numberofoutputs" gorm:"column:numberofoutputs"`
	DrawingLink2d     string `json:"drawinglink2d" gorm:"column:drawinglink2d"`
	StpFileLink       string `json:"stpfilelink" gorm:"column:stpfilelink"`
	DevicePictureLink string `json:"devicepicturelink" gorm:"column:devicepicturelink"`
	SpecSheetLink     string `json:"specsheetlink" gorm:"column:specsheetlink"`
	Dimensions        string `json:"dimensions" gorm:"column:dimensions"`
	Stage             string `json:"stage" gorm:"column:stage"`
}

func (DeviceCatalogue) TableName() string {
	return "intellicar.devicecatalogue"
}
