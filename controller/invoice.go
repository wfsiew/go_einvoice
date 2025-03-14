package controller

import (
	"einvoice/dbmodel"
	"einvoice/service"

	"github.com/gofiber/fiber/v2"
)

// GetInvoiceXml
//
// @Tags Invoice
// @Produce xml
// @Success 200
// @Router /invoice/xml [get]
func GetInvoiceXml(c *fiber.Ctx) error {
	st, _ := service.GetInvoiceSetting()
	o, _ := service.GetInvoiceEvent("01")
	var (
		buyer dbmodel.INTF_INT_EINVOICE_BUYER
		supp  dbmodel.INTF_INT_EINVOICE_SUPPLIER
		items []dbmodel.INTF_INT_EINVOICE_LINE_ITEM
	)

	if o.EVENT_SEQ_NO != "" {
		id := o.EVENT_SEQ_NO
		buyer, _ = service.GetInvoiceBuyer(id)
		supp, _ = service.GetInvoiceSupplier(id)
		items, _ = service.GetInvoiceLineItem(id)
	}

	prepaid_payment_date := ""
	prepaid_payment_time := ""
	if o.PRE_PAYMENT_DATE != "" {
		prepaid_payment_date = o.PRE_PAYMENT_DATE[:10]
		prepaid_payment_time = o.PRE_PAYMENT_DATE[11:19]
	}

	err := c.Render("invoice.xml", fiber.Map{
		"issue_date":           o.EINVOICE_DATE[:10],
		"issue_time":           o.EINVOICE_DATE[11:19],
		"billing_date":         o.BILLING_DATE[:10],
		"prepaid_payment_date": prepaid_payment_date,
		"prepaid_payment_time": prepaid_payment_time,
		"setting":              st,
		"event":                o,
		"buyer":                buyer,
		"supp":                 supp,
		"items":                items,
	})
	c.Set(fiber.HeaderContentType, fiber.MIMETextXML)
	return err
}

// GetInvoiceEvent
//
// @Tags Invoice
// @Produce json
// @Success 200 {object} dbmodel.INTF_INT_EINVOICE_EVENTS
// @Router /invoice/events [get]
func GetInvoiceEvent(c *fiber.Ctx) error {
	st, _ := service.GetInvoiceSetting()
	o, _ := service.GetInvoiceEvent("01")
	if o.EVENT_SEQ_NO != "" {
		id := o.EVENT_SEQ_NO
		buyer, _ := service.GetInvoiceBuyer(id)
		supp, _ := service.GetInvoiceSupplier(id)
		items, _ := service.GetInvoiceLineItem(id)
		return c.JSON(fiber.Map{
			"setting": st,
			"event":   o,
			"buyer":   buyer,
			"supp":    supp,
			"items":   items,
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ "error": "notFound" })
}