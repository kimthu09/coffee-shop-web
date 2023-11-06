package common

const (
	TableUser             string = "MUser"
	TableCategory         string = "Category"
	TableCategoryFood     string = "CategoryFood"
	TableTopping          string = "Topping"
	TableFood             string = "Food"
	TableIngredient       string = "Ingredient"
	TableIngredientDetail string = "IngredientDetail"
	TableRecipe           string = "Recipe"
	TableRecipeDetail     string = "RecipeDetail"
	TableSizeFood         string = "SizeFood"
	TableCancelNote       string = "CancelNote"
	TableCancelNoteDetail string = "CancelNoteDetail"
	TableSupplier         string = "Supplier"
	TableSupplierDebt     string = "SupplierDebt"
	TableCustomer         string = "Customer"
	TableCustomerDebt     string = "CustomerDebt"
	TableImportNote       string = "ImportNote"
	TableImportNoteDetail string = "ImportNoteDetail"
	TableExportNote       string = "ExportNote"
	TableExportNoteDetail string = "ExportNoteDetail"
	TableInvoice          string = "Invoice"
	TableInvoiceDetail    string = "InvoiceDetail"
	TableRole             string = "Role"
	TableFeature          string = "Feature"
	TableRoleFeature      string = "RoleFeature"
)

const MaxLengthIdCanGenerate = 12

const (
	CancelNoteCreateFeatureCode       = "CAN_CREATE"
	CategoryCreateFeatureCode         = "CAT_CREATE"
	CategoryUpdateInfoFeatureCode     = "CAT_UP_INFO"
	CustomerCreateFeatureCode         = "CUS_CREATE"
	CustomerPayFeatureCode            = "CUS_PAY"
	CustomerUpdateInfoFeatureCode     = "CUS_UP_INFO"
	ExportNoteCreateFeatureCode       = "EXP_CREATE"
	FoodCreateFeatureCode             = "FOD_CREATE"
	FoodUpdateInfoFeatureCode         = "FOD_UP_INFO"
	FoodUpdateStatusFeatureCode       = "FOD_UP_STATE"
	ImportNoteCreateFeatureCode       = "IMP_CREATE"
	ImportNoteChangeStatusFeatureCode = "IMP_UP_STATE"
	IngredientCreateFeatureCode       = "ING_CREATE"
	InvoiceCreateFeatureCode          = "INV_CREATE"
	SupplierCreateFeatureCode         = "SUP_CREATE"
	SupplierPayFeatureCode            = "SUP_PAY"
	SupplierUpdateInfoFeatureCode     = "SUP_UP_INFO"
	ToppingCreateFeatureCode          = "TOP_CREATE"
	ToppingUpdateInfoFeatureCode      = "TOP_UP_INFO"
	ToppingUpdateStatusFeatureCode    = "TOP_UP_STATE"
	UserUpdateInfoFeatureCode         = "USE_UP_INFO"
	UserUpdateStatusFeatureCode       = "USE_UP_STATE"
)

const RoleAdminId = "admin"

const DefaultPass = "app123"

const CurrentUserStr = "current_user"
