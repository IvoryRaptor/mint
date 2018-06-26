package http

import (
	"github.com/IvoryRaptor/dragonfly"
	"net/http"
	"fmt"
	"strings"
	"log"
	"encoding/json"
	"github.com/IvoryRaptor/mint"
	"sync"
)

type Service struct {
	kernel mint.IMint
	srv    *http.Server
}

func (s *Service) convertMatrix(name string, m *sync.Map) map[string]interface{} {
	angler := map[string]interface{}{}
	result := map[string]interface{}{"name": name, "angler": angler}

	m.Range(func(key, value interface{}) bool {
		angler[key.(string)] = s.convertAngler(value.(*sync.Map))
		return true
	})
	return result
}

func (s *Service) convertAngler(a *sync.Map) map[string]interface{} {
	result := map[string]interface{}{}
	a.Range(func(key, value interface{}) bool {
		result[key.(string)] = s.convertCopy(value.(*sync.Map))
		return true
	})
	return result
}

func (s *Service) convertCopy(b *sync.Map) map[string]interface{} {
	result := map[string]interface{}{}
	b.Range(func(key, value interface{}) bool {
		result[fmt.Sprintf("%d", key)] = value
		return true
	})
	return result
}

func (s *Service) Config(k dragonfly.IKernel, config map[interface{}]interface{}) error {
	kernel := k.(mint.IMint)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		matrixMap := kernel.GetMatrix()
		sp := strings.Split(r.URL.Path[1:], "/")
		count := len(sp)
		var result interface{}
		var err error
		switch count {
		case 1:
			if sp[0] == "" {
				t := make([]interface{}, 0)
				matrixMap.Range(func(key, value interface{}) bool {
					t = append(t, s.convertMatrix(key.(string), value.(*sync.Map)))
					return true
				})
				result = t
			} else {
				value, ok := matrixMap.Load(sp[0])
				if ok {
					result = s.convertMatrix(sp[0], value.(*sync.Map))
				} else {
					w.WriteHeader(404)
					return
				}
			}
		case 2:
			value, ok := matrixMap.Load(sp[0])
			if !ok {
				w.WriteHeader(404)
				return
			} else {
				matrix := value.(*sync.Map)
				value, ok := matrix.Load(sp[1])
				if !ok {
					w.WriteHeader(404)
					return
				} else {
					result = s.convertAngler(value.(*sync.Map))
				}
			}
		case 3:
			value, ok := matrixMap.Load(sp[0])
			if !ok {
				w.WriteHeader(404)
				return
			} else {
				matrix := value.(*sync.Map)
				value, ok := matrix.Load(sp[1])
				if !ok {
					w.WriteHeader(404)
					return
				} else {
					angler:=value.(*sync.Map)
					value, ok := angler.Load(sp[2])
					if !ok{
						w.WriteHeader(404)
						return
					}else{
						result = s.convertCopy(value.(*sync.Map))
					}
				}
			}
		}
		if result == nil {
			w.WriteHeader(404)
			return
		}
		var data []byte
		data, err = json.Marshal(result)
		if err != nil {
			w.WriteHeader(500)
		}
		w.Write(data)
	})
	port := fmt.Sprintf(":%d", config["port"])
	s.srv = &http.Server{Addr: port}
	log.Printf("Port: %s", port)
	return nil
}

func (s *Service) Start() error {
	go func() {
		s.srv.ListenAndServe()
	}()
	return nil
}

func (s *Service) Stop() {
	s.srv.Close()

}
