package design

import (
	. "goa.design/goa/v3/dsl"
)

var MailAddressPaylaod = Type("MailPayload", func() {
	Description("Payload of MailType")

	Attribute("address", String, "Address of user mail", func() {
		Format(FormatEmail)
		MinLength(1)
	})

	Attribute("active", Boolean, "Status of email address", func() {
		Default(true)
	})

	Required("address")
})

var MailInterface = Type("Mail", func() {
	Description("Mail type")
	Extend(MailAddressPaylaod)

	Attribute("id", String, "Unique mail ID")
	Attribute("createdAt", String, "Created at date of mail")
	Attribute("updateAt", String, "Updated at date of mail")

	Required("id", "address", "active", "createdAt", "updateAt")
})

var MailList = ResultType("MailList", func() {
	Description("List of mails")
	Attribute("data", ArrayOf(MailInterface))
	Attribute("page", Int, "Page number")
	Attribute("limit", Int, "Items per page")
	Attribute("total", Int, "Total number of mails")

	Required("data", "page", "limit", "total")
})

var _ = Service("MailAddresses", func() {
	Description("Servicio de recoleccion de datos sobre el talon de pago")

	Method("listMailAddresses", func() {
		Payload(func() {
			Attribute("page", Int, "Page number", func() {
				Minimum(0)
				Default(1)
			})
			Attribute("limit", Int, "Items per page", func() {
				Minimum(0)
				Maximum(100)
				Default(10)
			})
		})

		Result(MailList)

		HTTP(func() {
			GET("/mail-address")

			Param("page", Int, "Page number", func() {
				Minimum(0)
			})
			Param("limit", Int, "Number of items per page", func() {
				Minimum(0)
				Maximum(100)
			})

			Response(StatusOK)
		})
	})

	Method("getMailAddresses", func() {
		Description("Get mail address by ID")

		Payload(func() {
			Attribute("id", String, "Unique mail ID")
			Required("id")
		})

		Result(MailInterface)
		Error("not_found", ErrorResult, "Mail address not found")

		HTTP(func() {
			GET("/mail-address/{id}")

			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("createMailAddress", func() {
		Payload(MailAddressPaylaod)

		Result(MailInterface)

		HTTP(func() {
			POST("/mail-address")

			Response(StatusCreated)
		})
	})

	Method("updateMailAddress", func() {
		Payload(func() {
			Extend(MailAddressPaylaod)
			Attribute("id", String, "Unique mail ID")
			Required("id", "address", "active")
		})

		Error("not_found", ErrorResult, "Mail address not found")

		HTTP(func() {
			PUT("/mail-address/{id}")

			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("deleteMailAddress", func() {
		Payload(func() {
			Attribute("id", String, "Unique mail ID")
			Required("id")
		})

		Error("not_found", ErrorResult, "Mail address not found")

		HTTP(func() {
			DELETE("/mail-address/{id}")

			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})
})
