package utils

import (
    
)

type Response struct {
    Code    int
    Data    interface{}
    Msg     string
}

// func ResponseBody(code int, data interface{}, msg string) Response {
//     var response Response
//     response.code = code
//     response.data = data
//     response.msg  = msg

//     return response
// }


