package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Waire214/goninja/reuse"
)

var validResponse reuse.ResponseType

func GetSingleProduct(paramID string) reuse.ResponseType {
	var productsInfo []ProductsDetails
	var productInfo ProductsDetails
	productDetailsFile, err := os.Open("myFile0.csv")
	if err != nil {
		panic(err)
	}
	defer productDetailsFile.Close()
	rawProductsDetails, err := csv.NewReader(productDetailsFile).ReadAll()
	fmt.Println("file opened successfully")
	if err != nil {
		panic(err)
	}
	for _, rawProductDetails := range rawProductsDetails {
		productInfo.ProductName = rawProductDetails[0]
		productInfo.ProductDescription = rawProductDetails[1]
		productInfo.StockQuantity, _ = strconv.Atoi(rawProductDetails[2])
		productInfo.DisplayImage = rawProductDetails[3]
		productInfo.Price = rawProductDetails[4]
		productInfo.Industry = rawProductDetails[5]
		productInfo.ProductID = rawProductDetails[6]
		productInfo.Tags = rawProductDetails[7]
		productInfo.BusinessID = rawProductDetails[8]
		productsInfo = append(productsInfo, productInfo)
	}
	var perProductInfo ProductsDetails
	for _, perProductInfo = range productsInfo {
		if perProductInfo.ProductID == paramID {
			fmt.Println(validResponse.Response(200, &perProductInfo, "Successful"))
			return validResponse.Response(200, &perProductInfo, "Successful")

		} else {
			return validResponse.Response(400, &perProductInfo, "Unable to read file")
		}
	}
	return validResponse.Response(200, &perProductInfo, "Successful")

	// return validResponse.Response(200, &productInfo, "successful")
}
