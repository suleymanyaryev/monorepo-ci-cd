package datastore

import (
	"context"

	"example.com/monorepo-backend/responses"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

const (
	sqlGetTodoList   = `SELECT name, status FROM tbl_todo`
	sqlGetTodoByName = `SELECT name, status FROM tbl_todo WHERE name = $1`
	sqlCreateTodo    = `INSERT INTO tbl_todo(name, status) VALUES($1, $2)`
	sqlUpdateToDo    = `UPDATE tbl_todo SET name = $1 WHERE name = $2`
	sqlDeleteToDo    = `DELETE FROM tbl_todo WHERE name = $1`
	sqlChangeStatus  = `UPDATE tbl_todo SET status = $1 WHERE name = $2`
)

func (d *PgAccess) GetToDoList(ctx context.Context) (item *[]responses.ToDO, err error) {
	clog := log.WithFields(log.Fields{
		"method": "PgAccess.GetList",
	})
	err = d.runQuery(ctx, clog, func(conn *pgxpool.Conn) (err error) {
		defer func() {
			if err != nil {
				item = nil
			}
		}()

		todoes := make([]responses.ToDO, 0)
		rows, err := conn.Query(ctx, sqlGetTodoList)
		if err != nil {
			return
		}
		for rows.Next() {
			todo := responses.ToDO{}
			err = rows.Scan(
				&todo.Name,
				&todo.Status,
			)
			if err != nil {
				return
			}
			todoes = append(todoes, todo)

		}
		item = &todoes
		return
	})
	if err != nil {
		eMsg := "error in GetList()"
		clog.WithError(err).Error(eMsg)
	}
	return
}

func (d *PgAccess) GetToDoByName(ctx context.Context, name string) (item *responses.ToDO, err error) {
	clog := log.WithFields(log.Fields{
		"method": "PgAccess.GetByName",
	})

	err = d.runQuery(ctx, clog, func(conn *pgxpool.Conn) (err error) {
		defer func() {
			if err != nil {
				item = nil
			}
		}()
		item = &responses.ToDO{}

		row := conn.QueryRow(ctx, sqlGetTodoByName, name)
		err = row.Scan(
			&item.Name,
			&item.Status,
		)
		if err != nil {
			if err == pgx.ErrNoRows {
				err = nil
				item = nil
				return
			}
			return
		}
		return
	})
	if err != nil {
		eMsg := "error in GetByName()"
		clog.WithError(err).Error(eMsg)
	}
	return
}

func (d *PgAccess) CreateToDo(ctx context.Context, todo *responses.ToDO) (err error) {
	clog := log.WithFields(log.Fields{
		"method": "PgAccess.CreateToDo",
	})

	err = d.runInTx(ctx, nil, clog, func(tx pgx.Tx) (rollback bool, err error) {
		rollback = true

		_, err = tx.Exec(ctx, sqlCreateTodo, todo.Name, todo.Status)
		if err != nil {
			return
		}
		return false, nil
	})
	if err != nil {
		eMsg := "error in CreateToDo()"
		clog.WithError(err).Error(eMsg)
	}
	return
}

func (d *PgAccess) UpdateToDo(ctx context.Context, oldName, newName string) (err error) {
	clog := log.WithFields(log.Fields{
		"method": "PgAccess.UpdateToDo",
	})

	err = d.runInTx(ctx, nil, clog, func(tx pgx.Tx) (rollback bool, err error) {
		rollback = true

		_, err = tx.Exec(ctx, sqlUpdateToDo, newName, oldName)
		if err != nil {
			return
		}
		return false, nil
	})
	if err != nil {
		eMsg := "error in UpdateToDo()"
		clog.WithError(err).Error(eMsg)
	}
	return
}

func (d *PgAccess) DeleteToDo(ctx context.Context, name string) (err error) {
	clog := log.WithFields(log.Fields{
		"method": "PgAccess.DeleteToDo",
	})

	err = d.runInTx(ctx, nil, clog, func(tx pgx.Tx) (rollback bool, err error) {
		rollback = true

		_, err = tx.Exec(ctx, sqlDeleteToDo, name)
		if err != nil {
			return
		}
		return false, nil
	})
	if err != nil {
		eMsg := "error in DeleteToDo()"
		clog.WithError(err).Error(eMsg)
	}
	return
}

func (d *PgAccess) ChangeToDoStatus(ctx context.Context, name, status string) (err error) {
	clog := log.WithFields(log.Fields{
		"method": "PgAccess.ChangeToDoStatus",
	})

	err = d.runInTx(ctx, nil, clog, func(tx pgx.Tx) (rollback bool, err error) {
		rollback = true

		_, err = tx.Exec(ctx, sqlChangeStatus, status, name)
		if err != nil {
			return
		}
		return false, nil
	})
	if err != nil {
		eMsg := "error in ChangeToDoStatus()"
		clog.WithError(err).Error(eMsg)
	}
	return
}
