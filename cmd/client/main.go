package main

import (
	"context"
	"log"
	"net/http"

	todov1 "grpc_todo/gen/todo/v1"
	todoconnect "grpc_todo/gen/todo/v1/todov1connect"

	"connectrpc.com/connect"
)

func main() {
    client := todoconnect.NewToDoServiceClient(
        http.DefaultClient,
        "http://localhost:8080",
    )

	// ====== CreateToDo ======
    res, err := client.CreateToDo(
        context.Background(),
        connect.NewRequest(&todov1.CreateToDoRequest{Title: "Mike's ToDo"}),
    )
	// ====== DeleteToDo ======
    // res, err := client.DeleteToDo(
    //     context.Background(),
    //     connect.NewRequest(&todov1.DeleteToDoRequest{Id: "2dd35e76-d6b6-4719-8a39-e00c09f4ed3f"}),
    // )
	// ====== UpdateToDo ======
	// res, err := client.UpdateToDo(
    //     context.Background(),
    //     connect.NewRequest(&todov1.UpdateToDoRequest{
	// 		// idを文字列で指定しても現状エラーになる
	// 		Id: "0a5f83ca-9fa3-4d20-9b1b-ebfb7f7a67b0",
	// 		Status: todov1.ToDoStatus_TO_DO_STATUS_CLOSE,
	// 	}),
    // )

    if err != nil {
        log.Println(err)
        return
    }
	// ====== CreateToDo ======
    log.Println(res.Msg.Todo)
	// ====== DeleteToDo ======
    // log.Println(res.Msg.Id)
	// ====== UpdateToDo ======
    // log.Println(res.Msg.Todo)
}