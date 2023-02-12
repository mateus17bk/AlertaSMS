package main

import(
	"alertasms/sms"
)

func main() {
	sms.SendSMS("Alerta de servidor down", "<55 DD_SEU_NUMERO>")
}
