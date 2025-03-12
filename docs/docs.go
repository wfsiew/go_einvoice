// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/invoice/events": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.INTF_INT_EINVOICE_EVENTS"
                        }
                    }
                }
            }
        },
        "/invoice/setting": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.INTF_INT_EINVOICE_SETTING"
                        }
                    }
                }
            }
        },
        "/invoice/xml": {
            "get": {
                "produces": [
                    "text/xml"
                ],
                "tags": [
                    "Invoice"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "dbmodel.INTF_INT_EINVOICE_EVENTS": {
            "type": "object",
            "properties": {
                "bill_REF_NO": {
                    "type": "string"
                },
                "billing_DATE": {
                    "type": "string"
                },
                "einvoice_CODE": {
                    "type": "string"
                },
                "einvoice_DATE": {
                    "type": "string"
                },
                "einvoice_TYPE_CODE": {
                    "type": "string"
                },
                "event_SEQ_NO": {
                    "type": "string"
                },
                "payment_MODE": {
                    "type": "string"
                },
                "payment_TERMS": {
                    "type": "string"
                },
                "pre_PAYMENT_AMT": {
                    "type": "number"
                },
                "pre_PAYMENT_DATE": {
                    "type": "string"
                },
                "pre_PAYMENT_REF_NO": {
                    "type": "string"
                },
                "rounding_AMT": {
                    "type": "number"
                },
                "supplier_BANK_ACC": {
                    "type": "string"
                },
                "tax_EXEMPTION_AMT": {
                    "type": "number"
                },
                "tax_EXEMPTION_DETAILS": {
                    "type": "string"
                },
                "total_EXCLUDING_TAX": {
                    "type": "number"
                },
                "total_INCLUDING_TAX": {
                    "type": "number"
                },
                "total_NET_AMT": {
                    "type": "number"
                },
                "total_PAYABLE_AMT": {
                    "type": "number"
                },
                "total_TAXABLE_AMT_PER_TAX_TYPE": {
                    "type": "number"
                },
                "total_TAX_AMT": {
                    "type": "number"
                },
                "total_TAX_AMT_PER_TAX_TYPE": {
                    "type": "number"
                }
            }
        },
        "dbmodel.INTF_INT_EINVOICE_SETTING": {
            "type": "object",
            "properties": {
                "billing_FREQUENCY": {
                    "type": "string"
                },
                "currency_CODE": {
                    "type": "string"
                },
                "currency_EXCHANGE_RATE": {
                    "type": "string"
                },
                "einvoice_VERSION": {
                    "type": "string"
                },
                "tax_TYPE": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/einvoice",
	Schemes:          []string{},
	Title:            "Swagger E-Invoice API",
	Description:      "This is a e-invoice server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
