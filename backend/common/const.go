package common

const (
	TableUser                     string = "MUser"
	TableCategory                 string = "Category"
	TableCategoryFood             string = "CategoryFood"
	TableTopping                  string = "Topping"
	TableFood                     string = "Food"
	TableIngredient               string = "Ingredient"
	TableRecipe                   string = "Recipe"
	TableRecipeDetail             string = "RecipeDetail"
	TableSizeFood                 string = "SizeFood"
	TableSupplier                 string = "Supplier"
	TableSupplierDebt             string = "SupplierDebt"
	TableCustomer                 string = "Customer"
	TableImportNote               string = "ImportNote"
	TableImportNoteDetail         string = "ImportNoteDetail"
	TableExportNote               string = "ExportNote"
	TableExportNoteDetail         string = "ExportNoteDetail"
	TableInventoryCheckNote       string = "InventoryCheckNote"
	TableInventoryCheckNoteDetail string = "InventoryCheckNoteDetail"
	TableInvoice                  string = "Invoice"
	TableInvoiceDetail            string = "InvoiceDetail"
	TableRole                     string = "Role"
	TableFeature                  string = "Feature"
	TableRoleFeature              string = "RoleFeature"
	TableShopGeneral              string = "ShopGeneral"
)

const MaxLengthIdCanGenerate = 12

const (
	InventoryCheckNoteViewFeatureCode   = "ICN_VIEW"
	InventoryCheckNoteCreateFeatureCode = "ICN_CREATE"
	CategoryViewFeatureCode             = "CAT_VIEW"
	CategoryCreateFeatureCode           = "CAT_CREATE"
	CategoryUpdateInfoFeatureCode       = "CAT_UP_INFO"
	CustomerViewFeatureCode             = "CUS_VIEW"
	CustomerCreateFeatureCode           = "CUS_CREATE"
	CustomerUpdateInfoFeatureCode       = "CUS_UP_INFO"
	ExportNoteViewFeatureCode           = "EXP_VIEW"
	ExportNoteCreateFeatureCode         = "EXP_CREATE"
	FoodViewFeatureCode                 = "FOD_VIEW"
	FoodCreateFeatureCode               = "FOD_CREATE"
	FoodUpdateInfoFeatureCode           = "FOD_UP_INFO"
	FoodUpdateStatusFeatureCode         = "FOD_UP_STATE"
	ImportNoteViewFeatureCode           = "IMP_VIEW"
	ImportNoteCreateFeatureCode         = "IMP_CREATE"
	ImportNoteChangeStatusFeatureCode   = "IMP_UP_STATE"
	IngredientViewFeatureCode           = "ING_VIEW"
	IngredientCreateFeatureCode         = "ING_CREATE"
	InvoiceViewFeatureCode              = "INV_VIEW"
	InvoiceCreateFeatureCode            = "INV_CREATE"
	SupplierViewFeatureCode             = "SUP_VIEW"
	SupplierCreateFeatureCode           = "SUP_CREATE"
	SupplierPayFeatureCode              = "SUP_PAY"
	SupplierUpdateInfoFeatureCode       = "SUP_UP_INFO"
	ToppingViewFeatureCode              = "TOP_VIEW"
	ToppingCreateFeatureCode            = "TOP_CREATE"
	ToppingUpdateInfoFeatureCode        = "TOP_UP_INFO"
	ToppingUpdateStatusFeatureCode      = "TOP_UP_STATE"
	UserViewFeatureCode                 = "USE_VIEW"
	UserUpdateInfoFeatureCode           = "USE_UP_INFO"
	UserUpdateStatusFeatureCode         = "USE_UP_STATE"
)

const RoleAdminId = "admin"

const DefaultPass = "app123"

const CurrentUserStr = "current_user"