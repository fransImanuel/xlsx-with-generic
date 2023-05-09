package models

import "time"

type GenericReport interface {
	Penjualan |
		Register |
		CustLevelAndSpend |
		Pengajuan |
		ExportSendToIDN
}

type Penjualan struct {
	Date                  time.Time `json:"date" gorm:"column:tanggal"`
	OrderNumber           int64     `json:"order_number" gorm:"column:order_id"`
	Email                 string    `json:"email" gorm:"column:email"`
	Income                float64   `json:"income" gorm:"column:income"`
	AdditionalIncome      float64   `json:"additional_income" gorm:"column:additional_income"`
	Cogs                  float64   `json:"cogs" gorm:"column:cogs"`
	ShippingChina         float64   `json:"shipping_china" gorm:"column:shipping_china"`
	ShippingIdn           float64   `json:"shipping_idn" gorm:"column:shipping_idn"`
	EstimatedShippingCost float64   `json:"estimated_shipping_cost" gorm:"column:estimated_cost_shipping"`
	Refund                float64   `json:"refund" gorm:"column:refund"`
	OrderComplete         string    `json:"order_complete" gorm:"column:order_complete"`
	OrderCancel           string    `json:"order_cancel" gorm:"column:order_cancel"`
	PaymentType           string    `json:"payment_type" gorm:"column:payment_type"`
	PaymentDate           string    `json:"payment_date" gorm:"column:payment_date"`
	Link                  string    `json:"link" gorm:"column:link"`
	NewMember             string    `json:"new_member" gorm:"column:new_member"`
}

type Register struct {
	Sales    string    `json:"sales" gorm:"column:sales"`
	Nama     string    `json:"nama" gorm:"column:nama"`
	Email    string    `json:"email" gorm:"column:email"`
	Telepon  string    `json:"telepon" gorm:"column:telepon"`
	Tanggal  time.Time `json:"tanggal" gorm:"column:tanggal"`
	Verified int64     `json:"verified" gorm:"column:verified"`
}

type CustLevelAndSpend struct {
	Email                       string    `json:"email" gorm:"column:email"`
	Telepon                     string    `json:"telepon" gorm:"column:telepon"`
	Sales                       string    `json:"sales" gorm:"column:sales"`
	Nama                        string    `json:"nama" gorm:"column:nama"`
	TanggalTransaksiTerakhir    time.Time `json:"tanggal_transaksi_terakhir" gorm:"column:transaksi_terakhir_tgl"`
	TransaksiTerakhirBelanja    string    `json:"transaksi_terakhir_belanja" gorm:"column:transaksi_terakhir_belanja"`
	JumlahTransaksi             int64     `json:"jumlah_transaksi" gorm:"column:jml_transaksi"`
	TotalBelanja                float64   `json:"total_belanja" gorm:"column:total_belanja"`
	Level                       string    `json:"level" gorm:"column:level"`
	JarakTransaksiTerakhirBulan int64     `json:"jarak_transaksi_terakhir_bulan" gorm:"column:jarak_transaksi_terakhir_bulan"`
	Status                      string    `json:"status" gorm:"column:status"`
}

type Pengajuan struct {
	JumlahBeli    int64     `json:"jumlah_beli" gorm:"column:jumlah_beli"`
	Id            int64     `json:"id" gorm:"column:id"`
	Produk        string    `json:"produk" gorm:"column:produk"`
	Gambar        string    `json:"gambar" gorm:"column:gambar"`
	Link          string    `json:"link" gorm:"column:link"`
	Nama          string    `json:"nama" gorm:"column:nama"`
	Email         string    `json:"email" gorm:"column:email"`
	Telepon       string    `json:"telepon" gorm:"column:telepon"`
	Keterangan    string    `json:"keterangan" gorm:"column:keterangan"`
	Status        string    `json:"status" gorm:"column:status"`
	Toko          string    `json:"toko" gorm:"column:toko"`
	Tanggal       time.Time `json:"tanggal" gorm:"column:tanggal"`
	IdPage        string    `json:"id_page" gorm:"column:id_page"`
	ProdukCh      string    `json:"produk_ch" gorm:"column:produk_ch"`
	Updater       string    `json:"updater" gorm:"column:updater"`
	TanggalUpdate time.Time `json:"tanggal_update" gorm:"column:date(request.updated)"`
	Sales         string    `json:"sales" gorm:"column:sales"`
	Kategori      int64     `json:"kategori" gorm:"column:kategori"`
	Alasan        string    `json:"alasan" gorm:"column:alasan"`
}

type ExportSendToIDN struct {
	MarkingCode string `json:"marking_code" gorm:"column:marking_code"`
	IDSO        int64  `json:"idso" gorm:"column:id_so"`
	IDKarton    string `json:"id_karton" gorm:"column:id_karton"`
	Volume      int64  `json:"volume" gorm:"column:volume"`
}
