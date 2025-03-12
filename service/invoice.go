package service

import (
	"database/sql"
	"einvoice/database"
	"einvoice/dbmodel"
	"einvoice/utils"
)

func GetInvoiceSetting() (dbmodel.INTF_INT_EINVOICE_SETTING, error) {
	o := dbmodel.INTF_INT_EINVOICE_SETTING{}
	db := database.GetDb()
    if db == nil {
        utils.ILogger.Info().Msg("db is nil")
        return o, nil
    }

	q := `SELECT EINVOICE_VERSION, CURRENCY_CODE, CURRENCY_EXCHANGE_RATE, BILLING_FREQUENCY, TAX_TYPE 
	    FROM INTF_INT_EINVOICE_SETTING 
		FETCH FIRST 1 ROWS ONLY`
	rows, err := db.Query(q)
    if err != nil {
        if rows != nil {
            defer rows.Close()
        }
        
        utils.Logger.Err(err).Msg(err.Error())
        return o, err
    }

    defer rows.Close()

	var (
		EINVOICE_VERSION sql.NullString
		CURRENCY_CODE sql.NullString
		CURRENCY_EXCHANGE_RATE sql.NullString
		BILLING_FREQUENCY sql.NullString
		TAX_TYPE sql.NullString
	)

	if rows.Next() {
		err := rows.Scan(&EINVOICE_VERSION, &CURRENCY_CODE, &CURRENCY_EXCHANGE_RATE, &BILLING_FREQUENCY, &TAX_TYPE)

        if err != nil {
            utils.Logger.Err(err).Msg(err.Error())
			return o, err
        }

		o.EINVOICE_VERSION = EINVOICE_VERSION.String
		o.CURRENCY_CODE = CURRENCY_CODE.String
		o.CURRENCY_EXCHANGE_RATE = CURRENCY_EXCHANGE_RATE.String
		o.BILLING_FREQUENCY = BILLING_FREQUENCY.String
		o.TAX_TYPE = TAX_TYPE.String
	}

	return o, nil
}

func GetInvoiceEvent() (dbmodel.INTF_INT_EINVOICE_EVENTS, error) {
	o := dbmodel.INTF_INT_EINVOICE_EVENTS{}
	db := database.GetDb()
    if db == nil {
        utils.ILogger.Info().Msg("db is nil")
        return o, nil
    }

	q := `SELECT EVENT_SEQ_NO, EINVOICE_TYPE_CODE, EINVOICE_CODE, EINVOICE_DATE,
	    BILLING_DATE, PAYMENT_MODE, SUPPLIER_BANK_ACC, PAYMENT_TERMS, PRE_PAYMENT_AMT,
	    PRE_PAYMENT_DATE, PRE_PAYMENT_REF_NO, BILL_REF_NO, TOTAL_EXCLUDING_TAX, TOTAL_INCLUDING_TAX,
	    TOTAL_PAYABLE_AMT, TOTAL_NET_AMT, TOTAL_TAX_AMT, ROUNDING_AMT, TOTAL_TAX_AMT_PER_TAX_TYPE,
	    TOTAL_TAXABLE_AMT_PER_TAX_TYPE, TAX_EXEMPTION_DETAILS, TAX_EXEMPTION_AMT
		FROM INTF_INT_EINVOICE_EVENTS
		ORDER BY CREATION_DATE_TIME FETCH FIRST 1 ROWS ONLY`
	rows, err := db.Query(q)
	if err != nil {
		if rows != nil {
			defer rows.Close()
		}
		
		utils.Logger.Err(err).Msg(err.Error())
		return o, err
	}

	defer rows.Close()

	var (
		EVENT_SEQ_NO sql.NullString
		EINVOICE_TYPE_CODE sql.NullString
		EINVOICE_CODE sql.NullString
		EINVOICE_DATE sql.NullString
		BILLING_DATE sql.NullString
		PAYMENT_MODE sql.NullString
		SUPPLIER_BANK_ACC sql.NullString
		PAYMENT_TERMS sql.NullString
		PRE_PAYMENT_AMT sql.NullFloat64
		PRE_PAYMENT_DATE sql.NullString
		PRE_PAYMENT_REF_NO sql.NullString
		BILL_REF_NO sql.NullString
		TOTAL_EXCLUDING_TAX sql.NullFloat64
		TOTAL_INCLUDING_TAX sql.NullFloat64
		TOTAL_PAYABLE_AMT sql.NullFloat64
		TOTAL_NET_AMT sql.NullFloat64
		TOTAL_TAX_AMT sql.NullFloat64
		ROUNDING_AMT sql.NullFloat64
		TOTAL_TAX_AMT_PER_TAX_TYPE sql.NullFloat64
		TOTAL_TAXABLE_AMT_PER_TAX_TYPE sql.NullFloat64
		TAX_EXEMPTION_DETAILS sql.NullString
		TAX_EXEMPTION_AMT sql.NullFloat64
	)

	if rows.Next() {
		err := rows.Scan(&EVENT_SEQ_NO, &EINVOICE_TYPE_CODE, &EINVOICE_CODE, &EINVOICE_DATE, &BILLING_DATE,
		    &PAYMENT_MODE, &SUPPLIER_BANK_ACC, &PAYMENT_TERMS, &PRE_PAYMENT_AMT, &PRE_PAYMENT_DATE,
		    &PRE_PAYMENT_REF_NO, &BILL_REF_NO, &TOTAL_EXCLUDING_TAX, &TOTAL_INCLUDING_TAX, &TOTAL_PAYABLE_AMT,
	        &TOTAL_NET_AMT, &TOTAL_TAX_AMT, &ROUNDING_AMT, &TOTAL_TAX_AMT_PER_TAX_TYPE, &TOTAL_TAXABLE_AMT_PER_TAX_TYPE,
			&TAX_EXEMPTION_DETAILS, &TAX_EXEMPTION_AMT)

        if err != nil {
            utils.Logger.Err(err).Msg(err.Error())
			return o, err
        }

		o.EVENT_SEQ_NO = EVENT_SEQ_NO.String
		o.EINVOICE_TYPE_CODE = EINVOICE_TYPE_CODE.String
		o.EINVOICE_CODE = EINVOICE_CODE.String
		o.EINVOICE_DATE = EINVOICE_DATE.String
		o.BILLING_DATE = BILLING_DATE.String

		o.PAYMENT_MODE = PAYMENT_MODE.String
		o.SUPPLIER_BANK_ACC = SUPPLIER_BANK_ACC.String
		o.PAYMENT_TERMS = PAYMENT_TERMS.String
		o.PRE_PAYMENT_AMT = PRE_PAYMENT_AMT.Float64
		o.PRE_PAYMENT_DATE = PRE_PAYMENT_DATE.String

		o.PRE_PAYMENT_REF_NO = PRE_PAYMENT_REF_NO.String
		o.BILL_REF_NO = BILL_REF_NO.String
		o.TOTAL_EXCLUDING_TAX = TOTAL_EXCLUDING_TAX.Float64
		o.TOTAL_INCLUDING_TAX = TOTAL_INCLUDING_TAX.Float64
		o.TOTAL_PAYABLE_AMT = TOTAL_PAYABLE_AMT.Float64

		o.TOTAL_NET_AMT = TOTAL_NET_AMT.Float64
		o.TOTAL_TAX_AMT = TOTAL_TAX_AMT.Float64
		o.ROUNDING_AMT = ROUNDING_AMT.Float64
		o.TOTAL_TAX_AMT_PER_TAX_TYPE = TOTAL_TAX_AMT_PER_TAX_TYPE.Float64
		o.TOTAL_TAXABLE_AMT_PER_TAX_TYPE = TOTAL_TAXABLE_AMT_PER_TAX_TYPE.Float64

		o.TAX_EXEMPTION_DETAILS = TAX_EXEMPTION_DETAILS.String
		o.TAX_EXEMPTION_AMT = TAX_EXEMPTION_AMT.Float64
	}

	return o, nil
}

func GetInvoiceSupplier(eventSeqNo string) (dbmodel.INTF_INT_EINVOICE_SUPPLIER, error) {
	o := dbmodel.INTF_INT_EINVOICE_SUPPLIER{}
	db := database.GetDb()
    if db == nil {
        utils.ILogger.Info().Msg("db is nil")
        return o, nil
    }

	q := `SELECT EVENT_SEQ_NO, SUPPLIER_NAME, SUPPLIER_TIN, SUPPLIER_REG_ID, SUPPLIER_SST_REG_NO,
	    SUPPLIER_TOURISM_TAX_REG_NO, SUPPLIER_EMAIL, SUPPLIER_MSIC_CODE, SUPPLIER_BIZ_ACTY_DESC, ADDRESS_0,
		ADDRESS_1, ADDRESS_2, ADDRESS_POSTAL, ADDRESS_CITY, ADDRESS_STATE,
		ADDRESS_COUNTRY, CONTACT_NO
	    FROM INTF_INT_EINVOICE_SUPPLIER
	    WHERE EVENT_SEQ_NO = :eventSeqNo
		FETCH FIRST 1 ROWS ONLY`
	rows, err := db.Query(q, eventSeqNo)
	if err != nil {
		if rows != nil {
			defer rows.Close()
		}
		
		utils.Logger.Err(err).Msg(err.Error())
		return o, err
	}

	defer rows.Close()

	var (
		EVENT_SEQ_NO sql.NullString
		SUPPLIER_NAME sql.NullString
		SUPPLIER_TIN sql.NullString
		SUPPLIER_REG_ID sql.NullString
		SUPPLIER_SST_REG_NO sql.NullString
		SUPPLIER_TOURISM_TAX_REG_NO sql.NullString
		SUPPLIER_EMAIL sql.NullString
		SUPPLIER_MSIC_CODE sql.NullString
		SUPPLIER_BIZ_ACTY_DESC sql.NullString
		ADDRESS_0 sql.NullString
		ADDRESS_1 sql.NullString
		ADDRESS_2 sql.NullString
		ADDRESS_POSTAL sql.NullString
		ADDRESS_CITY sql.NullString
		ADDRESS_STATE sql.NullString
		ADDRESS_COUNTRY sql.NullString
		CONTACT_NO sql.NullString
	)

	if rows.Next() {
		err := rows.Scan(&EVENT_SEQ_NO, &SUPPLIER_NAME, &SUPPLIER_TIN, &SUPPLIER_REG_ID, &SUPPLIER_SST_REG_NO,
		    &SUPPLIER_TOURISM_TAX_REG_NO, &SUPPLIER_EMAIL, &SUPPLIER_MSIC_CODE, &SUPPLIER_BIZ_ACTY_DESC, &ADDRESS_0,
		    &ADDRESS_1, &ADDRESS_2, &ADDRESS_POSTAL, &ADDRESS_CITY, &ADDRESS_STATE, &ADDRESS_COUNTRY, &CONTACT_NO)

        if err != nil {
            utils.Logger.Err(err).Msg(err.Error())
			return o, err
        }

		o.EVENT_SEQ_NO = EVENT_SEQ_NO.String
		o.SUPPLIER_NAME = SUPPLIER_NAME.String
		o.SUPPLIER_TIN = SUPPLIER_TIN.String
		o.SUPPLIER_REG_ID = SUPPLIER_REG_ID.String
		o.SUPPLIER_SST_REG_NO = SUPPLIER_SST_REG_NO.String

		o.SUPPLIER_TOURISM_TAX_REG_NO = SUPPLIER_TOURISM_TAX_REG_NO.String
		o.SUPPLIER_EMAIL = SUPPLIER_EMAIL.String
		o.SUPPLIER_MSIC_CODE = SUPPLIER_MSIC_CODE.String
		o.SUPPLIER_BIZ_ACTY_DESC = SUPPLIER_BIZ_ACTY_DESC.String
		o.ADDRESS_0 = ADDRESS_0.String

		o.ADDRESS_1 = ADDRESS_1.String
		o.ADDRESS_2 = ADDRESS_2.String
		o.ADDRESS_POSTAL = ADDRESS_POSTAL.String
		o.ADDRESS_CITY = ADDRESS_CITY.String
		o.ADDRESS_STATE = ADDRESS_STATE.String
		o.ADDRESS_COUNTRY = ADDRESS_COUNTRY.String
		o.CONTACT_NO = CONTACT_NO.String
	}

	return o, nil
}

func GetInvoiceBuyer(eventSeqNo string) (dbmodel.INTF_INT_EINVOICE_BUYER, error) {
	o := dbmodel.INTF_INT_EINVOICE_BUYER{}
	db := database.GetDb()
    if db == nil {
        utils.ILogger.Info().Msg("db is nil")
        return o, nil
    }

	q := `SELECT EVENT_SEQ_NO, BUYER_NAME, BUYER_TIN, BUYER_REG_ID, BUYER_SST_REG_NO,
	    BUYER_EMAIL, ADDRESS_0, ADDRESS_1, ADDRESS_2, ADDRESS_POSTAL,
		ADDRESS_CITY, ADDRESS_STATE, ADDRESS_COUNTRY, CONTACT_NO
	    FROM INTF_INT_EINVOICE_BUYER
		WHERE EVENT_SEQ_NO = :eventSeqNo
		FETCH FIRST 1 ROWS ONLY`
	rows, err := db.Query(q, eventSeqNo)
	if err != nil {
		if rows != nil {
			defer rows.Close()
		}
		
		utils.Logger.Err(err).Msg(err.Error())
		return o, err
	}

	defer rows.Close()

	var (
		EVENT_SEQ_NO sql.NullString
		BUYER_NAME sql.NullString
		BUYER_TIN sql.NullString
		BUYER_REG_ID sql.NullString
		BUYER_SST_REG_NO sql.NullString
		BUYER_EMAIL sql.NullString
		ADDRESS_0 sql.NullString
		ADDRESS_1 sql.NullString
		ADDRESS_2 sql.NullString
		ADDRESS_POSTAL sql.NullString
		ADDRESS_CITY sql.NullString
		ADDRESS_STATE sql.NullString
		ADDRESS_COUNTRY sql.NullString
		CONTACT_NO sql.NullString
	)

	if rows.Next() {
		err := rows.Scan(&EVENT_SEQ_NO, &BUYER_NAME, &BUYER_TIN, &BUYER_REG_ID, &BUYER_SST_REG_NO,
		    &BUYER_EMAIL, &ADDRESS_0, &ADDRESS_1, &ADDRESS_2, &ADDRESS_POSTAL,
		    &ADDRESS_CITY, &ADDRESS_STATE, &ADDRESS_COUNTRY, &CONTACT_NO)

        if err != nil {
            utils.Logger.Err(err).Msg(err.Error())
			return o, err
        }

		o.EVENT_SEQ_NO = EVENT_SEQ_NO.String
		o.BUYER_NAME = BUYER_NAME.String
		o.BUYER_TIN = BUYER_TIN.String
		o.BUYER_REG_ID = BUYER_REG_ID.String
		o.BUYER_SST_REG_NO = BUYER_SST_REG_NO.String

		o.BUYER_EMAIL = BUYER_EMAIL.String
		o.ADDRESS_0 = ADDRESS_0.String
		o.ADDRESS_1 = ADDRESS_1.String
		o.ADDRESS_2 = ADDRESS_2.String
		o.ADDRESS_POSTAL = ADDRESS_POSTAL.String
		
		o.ADDRESS_CITY = ADDRESS_CITY.String
		o.ADDRESS_STATE = ADDRESS_STATE.String
		o.ADDRESS_COUNTRY = ADDRESS_COUNTRY.String
		o.CONTACT_NO = CONTACT_NO.String
	}

	return o, nil
}

func GetInvoiceLineItem(eventSeqNo string) ([]dbmodel.INTF_INT_EINVOICE_LINE_ITEM, error) {
	ls := make([]dbmodel.INTF_INT_EINVOICE_LINE_ITEM, 0)
	db := database.GetDb()
    if db == nil {
        utils.ILogger.Info().Msg("db is nil")
        return ls, nil
    }

	q := `SELECT EVENT_SEQ_NO, CLASSIFICATION_CODE, ITEM_DESC, UNIT_PRICE, TAX_TYPE, 
	    TAX_RATE, TAX_AMT, TAX_EXEMPTION_DESC, TAX_EXMEMPTED_AMT, SUBTOTAL, 
		TOTAL_EXCLUDING_TAX, QUANTITY, MEASUREMENT, DISCOUNT_RATE, DISCOUNT_AMT, 
		CHARGE_RATE, CHARGE_AMT
        FROM INTF_INT_EINVOICE_LINE_ITEM WHERE EVENT_SEQ_NO = :eventSeqNo`
	rows, err := db.Query(q, eventSeqNo)
	if err != nil {
		if rows != nil {
			defer rows.Close()
		}
		
		utils.Logger.Err(err).Msg(err.Error())
		return ls, err
	}

	defer rows.Close()

	var (
		EVENT_SEQ_NO sql.NullString
		CLASSIFICATION_CODE sql.NullString
		ITEM_DESC sql.NullString
		UNIT_PRICE sql.NullFloat64
		TAX_TYPE sql.NullString
		TAX_RATE sql.NullFloat64
		TAX_AMT sql.NullFloat64
		TAX_EXEMPTION_DESC sql.NullString
		TAX_EXMEMPTED_AMT sql.NullFloat64
		SUBTOTAL sql.NullFloat64
		TOTAL_EXCLUDING_TAX sql.NullFloat64
		QUANTITY sql.NullFloat64
		MEASUREMENT sql.NullString
		DISCOUNT_RATE sql.NullFloat64
		DISCOUNT_AMT sql.NullFloat64
		CHARGE_RATE sql.NullFloat64
		CHARGE_AMT sql.NullFloat64
	)

	for rows.Next() {
        err := rows.Scan(&EVENT_SEQ_NO, &CLASSIFICATION_CODE, &ITEM_DESC, &UNIT_PRICE, &TAX_TYPE,
		    &TAX_RATE, &TAX_AMT, &TAX_EXEMPTION_DESC, &TAX_EXMEMPTED_AMT, &SUBTOTAL,
		    &TOTAL_EXCLUDING_TAX, &QUANTITY, &MEASUREMENT, &DISCOUNT_RATE, &DISCOUNT_AMT,
		    &CHARGE_RATE, &CHARGE_AMT)

        if err != nil {
            utils.Logger.Err(err).Msg(err.Error())
            continue
        }

        o := dbmodel.INTF_INT_EINVOICE_LINE_ITEM{
            EVENT_SEQ_NO: EVENT_SEQ_NO.String,
            CLASSIFICATION_CODE: CLASSIFICATION_CODE.String,
            ITEM_DESC: ITEM_DESC.String,
			UNIT_PRICE: UNIT_PRICE.Float64,
			TAX_TYPE: TAX_TYPE.String,

			TAX_RATE: TAX_RATE.Float64,
			TAX_AMT: TAX_AMT.Float64,
			TAX_EXEMPTION_DESC: TAX_EXEMPTION_DESC.String,
			TAX_EXMEMPTED_AMT: TAX_EXMEMPTED_AMT.Float64,
			SUBTOTAL: SUBTOTAL.Float64,

			TOTAL_EXCLUDING_TAX: TOTAL_EXCLUDING_TAX.Float64,
			QUANTITY: QUANTITY.Float64,
			MEASUREMENT: MEASUREMENT.String,
			DISCOUNT_RATE: DISCOUNT_RATE.Float64,
			DISCOUNT_AMT: DISCOUNT_AMT.Float64,

			CHARGE_RATE: CHARGE_RATE.Float64,
			CHARGE_AMT: CHARGE_AMT.Float64,
        }
        ls = append(ls, o)
    }

	return ls, nil
}