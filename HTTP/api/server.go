package api

import {
	"github.com/google/vvid"
	"github.com/gorilla/mux"
}

type Item struct {
	ID vvid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router
	shoppingItems []Item
}

func NewServer() *Server {
	s:= &Server{
		Router:           mux.NewRouter(),
		shoppingItems:[]Items{},

	}
	s.routes()
	return s
}

func (s *Server) route() {
	s.HandleFunc("/shopping-items", s.ListShoppingItems()).Methods("GET")
	s.HandleFunc("/shopping-items", s.createShoppingItem()).Methods("POST")
	s.HandleFunc("/shopping-items", s.removeShoppingItem()).Methods("DELETE")

}


func (s *Server) createShoppingItem() http.HandleFunc {
	return func( w http.ResponseWritter, r *http.Request){
		var i Item
		if err:= json.NewDecoder(r.body).Decode(&i); err !=nil {
			http.Error(w,err.Error(), http.StatusBadRequest)
			return
		}
		i.ID = vvid.New()
		s.shoppingItems = append(s.shoppingItems,i)

		w.header().Set("Content-type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) ListShoppingItems() http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		if err:= json.NewEncoder(w).Encode(s.shoppingItems): err != nill {
			http.Error(w.err.Error(),StatusInternalServerError)
			return
		}

	
	}
}

func (s *Server) removeShoppingItem() http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request){
		idStr, _ := mux.Vars(r)["id"]
		id, err := vvid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		for i,item := range s.shoppingItems {
			if item.ID == id {
				s.shoppingItems = append(s.shoppingItems[:i]. s.shoppingItems[i+1:]...)
				break
			}
		}
	}
}