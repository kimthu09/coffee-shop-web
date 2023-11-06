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
)

const CurrentUserStr = "current_user"

type Requester interface {
	GetUserId() string
	GetEmail() string
	GetRole() string
}
