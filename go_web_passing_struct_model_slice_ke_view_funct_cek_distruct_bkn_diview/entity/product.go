package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

//cara pengcekean didalaem model atau struct tidak di view lg

//kenapa parameter nya di belakan contoh aja ini ceritanya buat pengecekan
//klo func  StockStatus(string id){}
//nah itu bisa kita gunakan klo mau post atau get atau delete lebih pefer ke situ
func (p Product) StockStatus() string {
	//variabel tampung
	var status string

	if p.Stock < 3 {
		status = "stok mau habis"
	} else if p.Stock > 3 && p.Stock < 4 {
		status = "stok masih banyak"
	} else {
		status = "barang di drop uot"
	}
	return status
}
