package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/abidkiller/hotel_reservation_backend/db"
	"github.com/abidkiller/hotel_reservation_backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	tus db.UserStore
}

func (tdb *testdb) tearDown(t *testing.T) {
	if err := tdb.tus.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		tus: db.NewMongoUserStore(client, db.TEST_DBNAME),
	}

}
func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.tus)

	app.Post("/", userHandler.HandleCreateUser)

	body := &types.CreateUserParams{
		FirstName: "abid",
		LastName:  "ahsan",
		Email:     "abid@gmail.com",
		Password:  "Abimin@29",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Errorf(err.Error())
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(jsonBody))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Test(req)

	respBody, err := io.ReadAll(resp.Body)
	fmt.Printf("restponse body: %+v \n status: %+v", string(respBody), resp.Status)

	// fmt.Println(resp.Body)
	// fmt.Print(resp.Status)
}
