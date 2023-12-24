package contextlearning

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// Context is used to store some values to pass to go routines but the most important feature is to add a timeout to a specific function
func ExecuteContext() {
	contextTest()
}

func contextTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	userId := 42069

	_, err := externalSlowApiCall(ctx, userId)
	if err != nil {
		log.Println("The external api call exeeded the timeout err: ", err)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond) // The timeout start from when i declare the context
	defer cancel2()

	present, err := externalApiCall(ctx2, userId)
	if err != nil {
		log.Println("The external api call exeeded the timeout err: ", err)
	}
	fmt.Println("Does the user exist? ", present)

}

func externalSlowApiCall(cxt context.Context, userID int) (bool, error) {
	time.Sleep(time.Millisecond * 400)

	if cxt.Err() == context.DeadlineExceeded {
		return false, errors.New("time exeeded")
	}

	return true, nil
}

func externalApiCall(cxt context.Context, userID int) (bool, error) {

	if cxt.Err() == context.DeadlineExceeded {
		return false, errors.New("time exeeded")
	}

	return true, nil
}
