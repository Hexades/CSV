package csv

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestCSVNBus(t *testing.T) {
	//_ = os.Remove("test_sqlite.db")
	NewCSVParser()
	readAll(t)
	//	assert.Nil(t, err)
	//	assert.NotNil(t, resp)
	//	assert.Equal(t, http.StatusOK, resp.StatusCode)
	//}
	//time.Sleep(10 * time.Second)

}
func readAll(t *testing.T) {
	evt := NewReadFileEvent("test.csv", TestModel{})
	SendEvent(evt)
	log.Println("Sent open event:", evt)
	response := <-evt.ResponseChannel()

	log.Println(reflect.TypeOf(response))
	for _, row := range response {
		fmt.Println(row)
	}
	//log.Println(response)
}

type TestModel struct {
	Index, CustomerId, FirstName, LastName, Company, City, Country, Phone1, Phone2, Email, SubscriptionDate, Website string
}
