// Package oracle contains the types for schema 'django'.
package oracle

// GENERATED BY XOXO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// DjangoAdminLog represents a row from 'django.django_admin_log'.
type DjangoAdminLog struct {
	ID            float64         // id
	ActionTime    time.Time       // action_time
	ObjectID      sql.NullString  // object_id
	ObjectRepr    sql.NullString  // object_repr
	ActionFlag    float64         // action_flag
	ChangeMessage sql.NullString  // change_message
	ContentTypeID sql.NullFloat64 // content_type_id
	UserID        float64         // user_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoAdminLog exists in the database.
func (dal *DjangoAdminLog) Exists() bool {
	return dal._exists
}

// Deleted provides information if the DjangoAdminLog has been deleted from the database.
func (dal *DjangoAdminLog) Deleted() bool {
	return dal._deleted
}

// Insert inserts the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dal._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO django.django_admin_log (` +
		`action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) VALUES (` +
		`:1, :2, :3, :4, :5, :6, :7` +
		`) RETURNING id /*lastInsertId*/ INTO :pk`

	// run query
	XOLog(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, nil)
	res, err := db.Exec(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, nil)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	dal.ID = float64(id)
	dal._exists = true

	return nil
}

// Update updates the DjangoAdminLog in the database.
func (dal *DjangoAdminLog) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dal._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dal._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django.django_admin_log SET ` +
		`action_time = :1, object_id = :2, object_repr = :3, action_flag = :4, change_message = :5, content_type_id = :6, user_id = :7` +
		` WHERE id = :8`

	// run query
	XOLog(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ID)
	_, err = db.Exec(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ID)
	return err
}

// Save saves the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Save(db XODB) error {
	if dal.Exists() {
		return dal.Update(db)
	}

	return dal.Insert(db)
}

// Delete deletes the DjangoAdminLog from the database.
func (dal *DjangoAdminLog) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dal._exists {
		return nil
	}

	// if deleted, bail
	if dal._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django.django_admin_log WHERE id = :1`

	// run query
	XOLog(sqlstr, dal.ID)
	_, err = db.Exec(sqlstr, dal.ID)
	if err != nil {
		return err
	}

	// set deleted
	dal._deleted = true

	return nil
}

// DjangoContentType returns the DjangoContentType associated with the DjangoAdminLog's ContentTypeID (content_type_id).
//
// Generated from foreign key 'd66d09ff3188e34f6a42ff99b3e0eb'.
func (dal *DjangoAdminLog) DjangoContentType(db XODB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(db, dal.ContentTypeID.Float64)
}

// AuthUser returns the AuthUser associated with the DjangoAdminLog's UserID (user_id).
//
// Generated from foreign key 'e3cca4fe721a3e54da54a29c371971'.
func (dal *DjangoAdminLog) AuthUser(db XODB) (*AuthUser, error) {
	return AuthUserByID(db, dal.UserID)
}

// DjangoAdminLogsByContentTypeID retrieves a row from 'django.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_417f1b1c'.
func DjangoAdminLogsByContentTypeID(db XODB, contentTypeID sql.NullFloat64) ([]*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM django.django_admin_log ` +
		`WHERE content_type_id = :1`

	// run query
	XOLog(sqlstr, contentTypeID)
	q, err := db.Query(sqlstr, contentTypeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoAdminLog{}
	for q.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}

		// scan
		err = q.Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID)
		if err != nil {
			return nil, err
		}

		res = append(res, &dal)
	}

	return res, nil
}

// DjangoAdminLogsByUserID retrieves a row from 'django.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_e8701ad4'.
func DjangoAdminLogsByUserID(db XODB, userID float64) ([]*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM django.django_admin_log ` +
		`WHERE user_id = :1`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoAdminLog{}
	for q.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}

		// scan
		err = q.Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID)
		if err != nil {
			return nil, err
		}

		res = append(res, &dal)
	}

	return res, nil
}

// DjangoAdminLogByID retrieves a row from 'django.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'sys_c005002'.
func DjangoAdminLogByID(db XODB, id float64) (*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM django.django_admin_log ` +
		`WHERE id = :1`

	// run query
	XOLog(sqlstr, id)
	dal := DjangoAdminLog{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID)
	if err != nil {
		return nil, err
	}

	return &dal, nil
}
