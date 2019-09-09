package calc

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/caarlos0/httperr"
	"github.com/felipeweb/distributedcalc/parser"
)

func Handler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return httperr.Wrap(errors.New(http.StatusText(http.StatusNotFound)),
			http.StatusNotFound)
	}
	input := parser.Input{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return httperr.Wrap(err, http.StatusBadRequest)
	}
	defer r.Body.Close()
	result, err := parser.Eval(r.Context(), input)
	if err != nil {
		return httperr.Wrap(err, http.StatusBadRequest)
	}
	err = json.NewEncoder(w).Encode(Result{
		Result: result,
	})
	if err != nil {
		return httperr.Wrap(err, http.StatusInternalServerError)
	}
	return nil
}

type Result struct {
	Result float64 `json:"result,omitempty"`
}
