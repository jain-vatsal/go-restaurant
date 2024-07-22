package endpoints

var (
	Users      = "/users"
	UserId     = "/users/:user_id"
	UserSignUp = "/users/sign_up"
	UserLogin  = "/users/login"
)

var (
	Foods  = "/foods"
	FoodId = "/foods/:food_id"
)

var (
	Invoices  = "/invoices"
	InvoiceId = "/invoices/:invoice_id"
)

var (
	Menus  = "/menus"
	MenuId = "/menus/:menu_id"
)

var (
	Orders  = "/orders"
	OrderId = "/orders/:order_id"
)

var (
	OrderItems         = "/orderitems"
	OrderItemId        = "/orderitems/:orderitem_id"
	AllItemsForOrderID = "/orderitems-order/:order_id"
)

var (
	Tables  = "/tables"
	TableId = "/tables/:order_id"
)
