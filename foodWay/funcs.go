package foodWay

import (
	"database/sql"
)

type cl []cluster.ClusterPoint

func (c cl) Len() int           { return len(c) }
func (c cl) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c cl) Less(i, j int) bool { return c[i].NumPoints > c[j].NumPoints }

type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Name      string  `json:"name"`
}

func (p Point) GetCoordinates() cluster.GeoCoordinates {
	return cluster.GeoCoordinates{p.Latitude, p.Longitude}
}

func GetMalls(db accessToDB) ([]Point, error) {
	var (
		err    error
		rows   *sql.Rows
		points []Point
		p      Point
	)
	if rows, err = db.Query("SELECT name, latitude, longitude FROM mall ORDER BY id DESC"); err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&p.Name, &p.Latitude, &p.Longitude); err != nil {
			return nil, err
		}
		points = append(points, p)
	}
	return points, nil
}
