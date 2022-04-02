package models

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrSlotBusy   = errors.New("slot busy")
	ErrValueError = errors.New("value error")
	ErrDataError  = errors.New("data inconsistency error")
)
