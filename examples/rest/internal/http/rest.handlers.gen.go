// Code generated by Rest annotation processor. DO NOT EDIT.
// versions:
//		go-annotation: 0.0.4-alpha
//		Rest: 0.0.1
package handlers

import (
	"net/http"
)

func (s *Container) Handlers() map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		http.MethodDelete: s.deleteContainer,
		http.MethodGet:    s.getContainer,
		http.MethodPatch:  s.patchContainer,
	}
}

func (s *Container) Path() string {
	return "/api/v1/containers/{container_id}"
}
