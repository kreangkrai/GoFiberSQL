package Controller

import (
	"GoFiberSQL/Models"
	"context"
	"database/sql"
	"log"
)

func GetsData() ([]Models.Data, error) {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT Device,Date,Value FROM Data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []Models.Data

	for rows.Next() {
		var device string
		var date string
		var value string
		err := rows.Scan(&device, &date, &value)
		if err != nil {
			return nil, err
		}
		data := Models.Data{
			Device: device,
			Date:   date,
			Value:  value,
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func GetData(device string) ([]Models.Data, error) {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT Device,Date,Value FROM Data WHERE Device='"+device+"'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []Models.Data

	for rows.Next() {
		var device string
		var date string
		var value string
		err := rows.Scan(&device, &date, &value)
		if err != nil {
			return nil, err
		}
		data := Models.Data{
			Device: device,
			Date:   date,
			Value:  value,
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func Insert(datas Models.Data) (string, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", err
	}
	_, execErr := tx.ExecContext(ctx, "Insert Into Data(Device,Date,Value) VALUES (?,?,?)", datas.Device, datas.Date, datas.Value)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr.Error(), execErr
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		return "", err
	}

	return "Insert Success", nil
}
func Update(data Models.Data) (string, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", err
	}
	_, execErr := tx.ExecContext(ctx, "UPDATE Data SET Value = ? WHERE Device = ?", data.Value, data.Device)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr.Error(), execErr
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		return "", err
	}

	return "Update Success", nil
}

func Delete(id string) (string, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", err
	}
	_, execErr := tx.ExecContext(ctx, "Delete From Data WHERE ID = ? ", id)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr.Error(), execErr
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		return "", err
	}

	return "Delete Success", nil
}
