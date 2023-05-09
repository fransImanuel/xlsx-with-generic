package xlsx

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
	"xlsx-with-generic/models"

	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
	Create Excel
*/

func GenerateToExcel[V models.GenericReport](data []V, start, end string) (string, error) {
	var TargetRowNumber, TargetColAlp string
	var ColTitleArr []string
	var fileName string
	xlsx := excelize.NewFile()
	SheetName := "Sheet1"

	fmt.Println(reflect.TypeOf(data).String())
	switch reflect.TypeOf(data).String() {
	case "[]models.Penjualan":
		ColTitleArr = []string{
			"Date",
			"Order Number",
			"Email",
			"Income",
			"Additional Income",
			"Cogs",
			"Shipping China",
			"Shipping Idn",
			"Estimated Shipping Cost",
			"Refund",
			"Order Complete",
			"Order Cancel",
			"Payment Type",
			"Payment Date",
			"Link",
			"New Member",
		}
	case "[]models.Register":
		ColTitleArr = []string{
			"Sales",
			"Nama",
			"Email",
			"Telepon",
			"Tanggal",
			"Verified",
		}
	case "[]models.CustLevelAndSpend":
		ColTitleArr = []string{
			"Email",
			"Telepon",
			"Sales",
			"Nama",
			"TanggalTransaksiTerakhir",
			"TransaksiTerakhirBelanja",
			"JumlahTransaksi",
			"TotalBelanja",
			"Level",
			"JarakTransaksiTerakhirBulan",
			"Status",
		}
	case "[]models.Pengajuan":
		ColTitleArr = []string{
			"JumlahBeli",
			"Id",
			"Produk",
			"Gambar",
			"Link",
			"Nama",
			"Email",
			"Telepon",
			"Keterangan",
			"Status",
			"Toko",
			"Tanggal",
			"IdPage",
			"ProdukCh",
			"Updater",
			"TanggalUpdate",
			"Sales",
			"Kategori",
			"Alasan",
		}
	case "[]models.ExportSendToIDN":
		ColTitleArr = []string{
			"Marking Code",
			"ID Order",
			"ID Karton",
			"Volume",
		}

	}

	//Set Excel Header
	//ascii A = 65, Z = 90
	style, err := xlsx.NewStyle(`{"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1}}`)
	if err != nil {
		return "", err
	}

	for i, title := range ColTitleArr {
		TargetRowNumber = strconv.Itoa(1)
		TargetColAlp = string(65 + i)
		xlsx.SetCellValue(SheetName, TargetColAlp+TargetRowNumber, title)
		xlsx.SetCellStyle(SheetName, TargetColAlp+TargetRowNumber, TargetColAlp+TargetRowNumber, style)
	}

	//Set Excel Content
	for i, d := range data {
		SetCellValue(xlsx, d, i+2)
	}
	xlsx.SetActiveSheet(1)

	// fileName := "./generated_file/excel/" + SetFileName(start, end) + ".xlsx"

	switch reflect.TypeOf(data).String() {
	case "[]models.Penjualan":
		fileName = "Penjualan#" + SetFileName(start, end) + ".xlsx"
	case "[]models.Register":
		fileName = "Register#" + SetFileName(start, end) + ".xlsx"
	case "[]models.CustLevelAndSpend":
		fileName = "CustLevelAndSpend#" + SetFileName(start, end) + ".xlsx"
	case "[]models.Pengajuan":
		fileName = "Pengajuan#" + SetFileName(start, end) + ".xlsx"
	case "[]models.ExportSendToIDN":
		fileName = "SendToIDN#" + start + ".xlsx"
	}
	// Save the spreadsheet by the given path.

	if err = xlsx.SaveAs(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}

func SetCellValue[V models.GenericReport](f *excelize.File, d V, rowIndex int) {
	rowIndexStr := strconv.Itoa(rowIndex)
	v := reflect.ValueOf(d)
	for i := 0; i < v.NumField(); i++ {
		TargetColAlph := string(65 + i)
		f.SetCellValue("Sheet1", TargetColAlph+rowIndexStr, v.Field(i).Interface())
	}
}

func SetFileName(start, end string) string {
	//file name = export data order +start + end + tanggal sekarang
	now := time.Now().Format("2006-01-02")
	fileName := start + "-" + end + "#" + now
	return fileName
}

/*
	Helper
*/

func removeFile(filepath string) error {
	if err := os.Remove(filepath); err != nil {
		return err
	}
	return nil
}
