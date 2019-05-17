package matrixCellsInDistanceOrder

import (
	"log"
	"testing"
)

func TestSolutions(t *testing.T) {
	rlt := allCellsDistOrder(2,2,0,1)
	log.Println(rlt)
}
