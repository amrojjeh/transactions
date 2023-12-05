package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (app *application) indexGet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		data, err := app.newTemplate()
		if err != nil {
			app.serverError(w, err)
			return
		}
		app.render(w, "transactions.tmpl", data, http.StatusOK)
	})

}

type newForm struct {
	TDate       time.Time
	PDate       time.Time
	ToAccount   string
	FromAccount string
	Amount      float64
	Context     string
	Tags        string
	Valid       bool
}

func (app *application) new() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.render(w, "new.tmpl", templateData{Form: newForm{}}, http.StatusOK)
			return
		}
		// Method is post
		err := r.ParseForm()
		if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
		}
		form := newForm{}
		tdate := r.Form.Get("tdate")
		if tdate == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		fmt.Printf("%v\n", tdate)
		form.TDate, err = time.Parse("006", tdate)
		if err != nil {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}

		pdate := r.Form.Get("pdate")
		if pdate == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		form.PDate, err = time.Parse("006", pdate)
		if err != nil {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}

		from_acc := r.Form.Get("from")
		if from_acc == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		form.FromAccount = from_acc

		to_acc := r.Form.Get("to")
		if to_acc == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		form.ToAccount = to_acc

		amount := r.Form.Get("amount")
		if to_acc == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		form.Amount, err = strconv.ParseFloat(amount, 64)
		if err != nil {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}

		context := r.Form.Get("context")
		if context == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		form.Context = context

		tags := r.Form.Get("tags")
		if tags == "" {
			form.Valid = false
			app.render(w, "new.tmpl", templateData{Form: form},
				http.StatusUnprocessableEntity)
			return
		}
		form.Tags = tags

		err = app.tmodel.NewTransaction(form.TDate, form.PDate, form.FromAccount, form.ToAccount,
			form.Amount, form.Context, form.Tags)
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
