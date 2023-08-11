package models

import (
	"net/http"

	"private_test/db"
)

type DBCheck struct {
	Id string `json:"id"`
}

func HealthCheck() (Response, error) {
	var obj DBCheck
	var arrobj []DBCheck
	var res Response

	con := db.CreatCon()
	sqlStatement := "select 1"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Messages = "Success!"
	res.Data = arrobj

	return res, nil
}
