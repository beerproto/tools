package scraper

import (
	"golang.org/x/text/language"
	"tawesoft.co.uk/go/lxstrconv"
)

var llesmaltiers = NewMalting("Les Maltiers", "FRA", lxstrconv.NewDecimalFormat(language.French),
	WithBaseURL("https://www.lesmaltiers.fr/nos-malts/"),
	WithProductSelector(".shop-page-content .product-small .product-title a[href]"),
	WithTitleSelector("h1.product-title"),
	WithProductRowSelector(".woocommerce-tabs tr"),
	WithProductHeaderSelector("th:first-child"),
	WithProductHeaderSelector("td:first-child"),
	WithProductValueSelector("td:nth-child(2)"),
	WithProductValueSelector("td"),
	WithColor("Couleur"),
	WithColor("EBC Colour"),
	WithMoisture("Moisture"),
	WithProtein("Protéines totales"),
	WithSolubleProtein("Protéines solubles"),
	WithFan("EBC FAN in Wort"),
	WithViscosity("EBC Wort Viscosity"),
	WithDiastaticPower("Pouvoir diastasique"),
	WithDiPh("pH:"),
	WithAlphaAmylase("Alpha Amylase dry basis"),
	WithBetaGlucan("B glucans"),
	WithFriability("Friabilité"),
	WithYield("Extrait"),
	WithYield("EBC Extract 0.2mm dry basis"),
)
