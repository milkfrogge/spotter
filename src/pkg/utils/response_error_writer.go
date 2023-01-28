/*
 * Copyright (c)  Nikita Cherkasov, 2023.
 * Spotter Project
 */

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResponseError(w http.ResponseWriter, e error) {
	a, _ := json.Marshal(map[string]string{"err": fmt.Sprintf("%v", e)})
	w.WriteHeader(500)
	w.Write(a)
}
