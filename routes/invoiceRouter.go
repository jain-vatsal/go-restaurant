package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jain-vatsal/go-restaurant/controllers"
	"github.com/jain-vatsal/go-restaurant/endpoints"
)

func InvoiceRoutes(invoiceRoutes *gin.Engine) {

	invoiceRoutes.GET(endpoints.Invoices, controllers.GetInvoices())
	invoiceRoutes.GET(endpoints.InvoiceId, controllers.GetInvoice())
	invoiceRoutes.POST(endpoints.Invoices, controllers.CreateInvoices())
	invoiceRoutes.PATCH(endpoints.InvoiceId, controllers.UpdateInvoices())

}
