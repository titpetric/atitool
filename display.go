package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func romVendorId(field uint16) string {
	switch field {
	case 0x1002:
		return "AMD"
	default:
		hasUnknownIds = true
		return fmt.Sprintf("0x%x", field)
	}
}

func romDeviceId(field uint16) string {
	switch field {
	case 0x67c0:
		return "Ellesmere core (Polaris 10 family)"
	case 0x67df:
		return "Ellesmere core (Polaris 10 family)"
	case 0x67e0:
		return "Baffin core (Polaris 11 family)"
	case 0x67e1:
		return "Baffin core (Polaris 11 family)"
	case 0x67e9:
		return "Baffin core (Polaris 11 family)"
	case 0x67eb:
		return "Baffin core (Polaris 11 family)"
	case 0x67ff:
		return "Baffin core (Polaris 11 family)"
	case 0x6860:
		return "Unknown (Vega 10 family)"
	case 0x6861:
		return "Unknown (Vega 10 family)"
	case 0x6862:
		return "Unknown (Vega 10 family)"
	case 0x6863:
		return "Unknown (Vega 10 family)"
	case 0x6867:
		return "Unknown (Vega 10 family)"
	case 0x686c:
		return "Unknown (Vega 10 family)"
	case 0x687f:
		return "Unknown (Vega 10 family)"
	case 0x6980:
		return "Unknown (Polaris 12 family)"
	case 0x6981:
		return "Unknown (Polaris 12 family)"
	case 0x6985:
		return "Unknown (Polaris 12 family)"
	case 0x6986:
		return "Unknown (Polaris 12 family)"
	case 0x6987:
		return "Unknown (Polaris 12 family)"
	case 0x6995:
		return "Unknown (Polaris 12 family)"
	case 0x699F:
		return "Unknown (Polaris 12 family)"
	default:
		hasUnknownIds = true
		return fmt.Sprintf("0x%x", field)
	}
}

func subVendorId(field uint16) string {
	baseUrl := "http://pcidatabase.com"
	url := fmt.Sprintf("%s/search.php?vendor_search_str=0x%x&vendor_search.x=0&vendor_search.y=0&vendor_search=search+vendors",
		baseUrl,
		field)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		hasUnknownIds = true
		return fmt.Sprintf("0x%x", field)
	}

	var subVendorName string
	//var subVendorNameShort string
	//var subVendorNameShortLink string

	doc.Find("tr.odd").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		subVendorName = s.Find("a").Text()

		// Follow link to get the vendors short name (example: MSI).
		//link, exists := s.Find("a").Attr("href")
		//if exists {
		//	subVendorNameShortLink = fmt.Sprintf("%s\n", link)
		//
		//	url := fmt.Sprintf("%s/vendor_details.php?id=705",
		//		baseUrl)
		//	fmt.Println(url)
		//
		//	doc, err := goquery.NewDocument(url)
		//	if err != nil {
		//		doc.Find("tr.odd").Each(func(i int, s *goquery.Selection) {
		//
		//			//fmt.Println(s.Find("a").Text())
		//		})
		//	}
		//}
	})

	if subVendorName == "" {
		hasUnknownIds = true
		return fmt.Sprintf("0x%x", field)
	}

	return subVendorName
}


func vramVendorId(field byte) string {
	vendorId := uint16(field)




	switch vendorId {
	case 0x3:
		return "Elpida"
	case 0x66:
		return "Hynix"
	default:
		hasUnknownIds = true
		return fmt.Sprintf("0x%x", field)
	}
}