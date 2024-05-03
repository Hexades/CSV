package csv

import (
	"log"
	"testing"
)

func TestCSVNBus(t *testing.T) {
	//_ = os.Remove("test_sqlite.db")
	NewCSVParser()
	evt := NewEvent("test.csv")
	log.Println("Sending csv  file event")
	SendEvent(evt)
	SendEvent(NewEvent(""))
	//	assert.Nil(t, err)
	//	assert.NotNil(t, resp)
	//	assert.Equal(t, http.StatusOK, resp.StatusCode)
	//}
	//time.Sleep(10 * time.Second)

}
