package functions

import (
	"context"
	"net/http"
	"strings"
	"time"

	// Import the package that contains uuid_generate_v4

	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/services"
	"github.com/piyush7833/Chat-Api/types"
)

func SignUp(user types.SignupType) (interface{}, types.ErrorType) {
	// fmt.Println("User:", user)
	hashedPassword, err := helpers.CreateHashedPassword(*user.Password)
	if err != nil {
		return nil, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	name := helpers.GetNullableValue(user.Name)
	username := helpers.GetNullableValue(user.Username)
	email := helpers.GetNullableValue(user.Email)
	password := hashedPassword
	image := helpers.GetNullableValue(user.Image)
	phone := helpers.GetNullableValue(user.Phone)
	// fmt.Println(name, username, email, password, image, phone)
	result, err := services.Db.Exec(`
    INSERT INTO "users" (name, username, email, password, image, phone)
    VALUES ($1, $2, $3, $4, $5, $6)
`, name, username, email, password, image, phone)
	if err != nil && strings.Contains(err.Error(), " duplicate key") {
		return nil, types.ErrorType{
			Message:    "User with given user name or email or passwrod already exist",
			StatusCode: 403,
		}
	} else if err != nil {
		return nil, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}
	return result, types.ErrorType{}
}

// func Signin(user types.SignInType, w http.ResponseWriter) (interface{}, types.ErrorType) {
// 	fmt.Println(user, "user")
// 	username := helpers.GetNullableValue(user.Username)
// 	// email := helpers.GetNullableValue(user.Email)
// 	// phone := helpers.GetNullableValue(user.Phone)
// 	// if username != nil {
// 	// 	username = username.(string)
// 	// }
// 	// if email != nil {
// 	// 	email = email.(string)
// 	// }
// 	// if phone != nil {
// 	// 	phone = phone.(string)
// 	// }
// 	// query := `SELECT password,id FROM "users" WHERE username = $1 Or email = $2 or phone = $3`
// 	// row, err := services.Db.QueryContext(ctx, query, username, email, phone)
// 	// if err != nil {
// 	// 	// fmt.Println("Error executing query:", err.Error())
// 	// 	return nil, types.ErrorType{
// 	// 		Message:    err.Error(),
// 	// 		StatusCode: 500,
// 	// 	}
// 	// }
// 	// defer row.Close()
// 	whereClause := fmt.Sprintf(`username = "%s"`, username.(string))
// 	res, error := services.GetRows("users", 0, []string{"id", "password"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
// 	fmt.Println(res, error, "res")
// 	resMap := res[0]
// 	fmt.Println(resMap, "resmap")
// 	if error.StatusCode != 0 {
// 		return nil, error
// 	}
// 	// if !row.Next() {
// 	// 	if err := row.Err(); err != nil {
// 	// 		// Handle the error
// 	// 		// fmt.Println("Error iterating rows:", err.Error())
// 	// 		return nil, types.ErrorType{
// 	// 			Message:    err.Error(),
// 	// 			StatusCode: 500,
// 	// 		}
// 	// 	}
// 	// 	// fmt.Println("No rows returned")

// 	// 	return nil, types.ErrorType{
// 	// 		Message:    "No user found",
// 	// 		StatusCode: 404,
// 	// 	}
// 	// }

//		// Scan the row data
//		// var password string
//		// var id string
//		// if err := row.Scan(&password, &id); err != nil {
//		// 	return nil, types.ErrorType{
//		// 		Message:    err.Error(),
//		// 		StatusCode: 500,
//		// 	}
//		// }
//		userPassword := helpers.GetNullableValue(user.Password)
//		if userPassword == nil {
//			return nil, types.ErrorType{
//				Message:    "Password is missing",
//				StatusCode: 500,
//			}
//		}
//		// pswdErr := helpers.VerifyHashedPassword(fmt.Sprintf("%v", resMap["password"]), userPassword.(string))
//		// if pswdErr != nil {
//		// 	return nil, types.ErrorType{
//		// 		Message:    "Incorrect password",
//		// 		StatusCode: 403,
//		// 	}
//		// }
//		token, _ := helpers.GenerateJWT(resMap["id"].(string))
//		setErr := helpers.SetCookie(w, "token", token, time.Now().Add(24*time.Hour))
//		if setErr != nil {
//			return nil, types.ErrorType{
//				Message:    setErr.Error(),
//				StatusCode: 500,
//			}
//		}
//		// User found, return the user dataW
//		return nil, types.ErrorType{}
//	}
func Signin(user types.SignInType, w http.ResponseWriter) (interface{}, types.ErrorType) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout*time.Second)
	defer cancel()
	username := helpers.GetNullableValue(user.Username)
	email := helpers.GetNullableValue(user.Email)
	phone := helpers.GetNullableValue(user.Phone)
	if user.Password == nil {
		return nil, types.ErrorType{
			Message:    "Password is missing",
			StatusCode: 500,
		}
	}
	query := `SELECT password,id FROM "users" WHERE username = $1 Or email = $2 or phone = $3`
	row, err := services.Db.QueryContext(ctx, query, username, email, phone)
	if err != nil {
		// fmt.Println("Error executing query:", err.Error())
		return nil, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}
	defer row.Close()

	if !row.Next() {
		if err := row.Err(); err != nil {
			// Handle the error
			// fmt.Println("Error iterating rows:", err.Error())
			return nil, types.ErrorType{
				Message:    err.Error(),
				StatusCode: 500,
			}
		}
		// fmt.Println("No rows returned")

		return nil, types.ErrorType{
			Message:    "No user found",
			StatusCode: 404,
		}
	}

	// Scan the row data
	var password string
	var id string
	if err := row.Scan(&password, &id); err != nil {
		return nil, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}
	pswdErr := helpers.VerifyHashedPassword(password, *user.Password)
	if pswdErr != nil {
		return nil, types.ErrorType{
			Message:    "Incorrect password",
			StatusCode: 403,
		}
	}
	token, _ := helpers.GenerateJWT(id)
	setErr := helpers.SetCookie(w, "token", token, time.Now().Add(24*time.Hour))
	if setErr != nil {
		return nil, types.ErrorType{
			Message:    setErr.Error(),
			StatusCode: 500,
		}
	}
	// User found, return the user dataW
	return nil, types.ErrorType{}
}
