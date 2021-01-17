package port

type Port struct {
	ID             int64    `db:"id"`
	IDStr          string   `db:"id_str"`
	Name           string   `db:"name"`
	City           string   `db:"city"`
	Country        string   `db:"country"`
	CoordinatesLat *float64 `db:"coord_long"`
	CoordinatesLon *float64 `db:"coord_lat"`
	Provice        string   `db:"province"`
	Timezone       string   `db:"timezone"`
	Code           string   `db:"code"`
	Regions        string   `db:"regions"`
	Unlocs         string   `db:"unlocs"`
	Alias          string   `db:"alias"`
}
