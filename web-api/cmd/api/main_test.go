package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home Handler",
			handler:        homeHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Health Handler",
			handler:        healthHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "About Handler",
			handler:        aboutHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "Roger Zheng - CMPS2242 Lab #3",
		},
		{
			name:           "Time Handler",
			handler:        timeHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "", //Will validate status only since time varies
		},
		{
			name:           "Random Sum Handler",
			handler:        randomSumHandler,
			expectedStatus: http.StatusOK,
			expectedBody:   "", //Will validate status only since output varies due to random numbers
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			tt.handler.ServeHTTP(rr, req)
			if rr.Code != tt.expectedStatus {
				t.Errorf("got %v, expected %v", rr.Code, tt.expectedStatus)
			}
			if tt.expectedBody != "" {
				if rr.Body.String() != tt.expectedBody {
					t.Errorf("got %v, expected %v", rr.Body.String(), tt.expectedBody)
				}
			}
		})
	}
}
