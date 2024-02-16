package repository

import (
	"fmt"
	"github.com/indrawanagung/loyalty_management_api/db"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/util"
	"testing"
)

var status = domain.Status{
	Name: "active",
}
var statusRepository = NewStatusRepository()
var dbStatus = db.OpenConnection(util.LoadConfig("../").DBSource)

func SaveStatus() int {
	return statusRepository.Save(dbStatus, status)
}
func TestStatusRepositoryImpl_FindAll(t *testing.T) {
	id := SaveStatus()
	fmt.Println("id : ", id)
	statusResponse := statusRepository.FindAll(dbStatus)
	fmt.Println(statusResponse)
}
