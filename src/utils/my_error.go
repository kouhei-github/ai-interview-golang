package utils

//type error interface {
//	Error() string
//}

// MyError はカスタム・エラー・タイプを表す。
//
// エラーメッセージを保持する Message フィールドを含みます。
type MyError struct {
	Message string
}

// Error represents a method for the MyError type that implements the error interface.
// It returns a string representation of the error message by adding a code tag around the message.
// The error message is contained within the MyError type, which has a Message string field.
// The code tag format is "[code]\n" + myErr.Message + "[/code]\n".
// This method is used to format the error message and provide a string representation of the error.
// Example usage:
//
//	err := MyError{"Something went wrong"}
//	errMsg := err.Error()
//	fmt.Println(errMsg)
//
// Output:
//
//	[code]
//	Something went wrong
//	[/code]
func (myErr MyError) Error() string {
	msg := "[code]\n" + myErr.Message + "[/code]\n"
	return msg
}
