package models

import "time"

type Resolucion struct {
	Id                     int
	NumeroResolucion       string
	FechaExpedicion        time.Time
	Vigencia               int
	DependenciaId          int
	TipoResolucionId       int
	PreambuloResolucion    string
	ConsiderandoResolucion string
	NumeroSemanas          int
	Periodo                int
	Titulo                 string
	DependenciaFirmaId     int
	VigenciaCarga          int
	PeriodoCarga           int
	NuxeoUid               string
	Activo                 bool
	FechaCreacion          string
	FechaModificacion      string
	FechaInicio            time.Time
	FechaFin               time.Time
}
