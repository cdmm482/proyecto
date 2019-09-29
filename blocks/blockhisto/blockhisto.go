package blockhisto

import (
	"proyecto/controller"
	"proyecto/estructura"
	"time"
)

// GenerateBlockHisto genera un bloque del tipo Historial
func GenerateBlockHisto(histoLast estructura.History, codeMed string, codePat string, Medicos []estructura.Medic, Pacientes []estructura.Patient) estructura.History {
	var nwBl estructura.History
	t := time.Now()
	nwBl.Index = histoLast.Index + 1
	findP, findM := controller.FindUser(codePat, codeMed, Medicos, Pacientes)
	nwBl.Medi = findM
	nwBl.Patien = findP
	nwBl.QueryAt = t
	nwBl.PrevHash = histoLast.Hash
	nwBl.Hash = controller.CalculateHashHis(nwBl)

	return nwBl
}
