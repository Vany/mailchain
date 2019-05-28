// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mailchain/mailchain/internal/keystore"
	"github.com/mailchain/mailchain/internal/testutil"
	"github.com/mailchain/mailchain/internal/testutil/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetAddresses(t *testing.T) {
	assert := assert.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	type args struct {
		ks keystore.Store
	}
	tests := []struct {
		name       string
		args       args
		wantBody   string
		wantStatus int
	}{
		{
			"err-GetAddresses",
			args{
				func() keystore.Store {
					store := mocks.NewMockStore(mockCtrl)
					store.EXPECT().GetAddresses().Return(
						nil,
						errors.Errorf("error getting address"),
					).Times(1)

					return store
				}(),
			},
			"{\"code\":500,\"message\":\"error getting address\"}\n",
			http.StatusInternalServerError,
		},
		{
			"empty-address",
			args{
				func() keystore.Store {
					store := mocks.NewMockStore(mockCtrl)
					store.EXPECT().GetAddresses().Return(
						[][]byte{},
						nil,
					).Times(1)

					return store
				}(),
			},
			"{\"addresses\":[]}\n",
			http.StatusOK,
		},
		{
			"single-address",
			args{
				func() keystore.Store {
					store := mocks.NewMockStore(mockCtrl)
					store.EXPECT().GetAddresses().Return(
						[][]byte{testutil.MustHexDecodeString("5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761")},
						nil,
					).Times(1)

					return store
				}(),
			},
			"{\"addresses\":[\"5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761\"]}\n",
			http.StatusOK,
		},
		{
			"multi-address",
			args{
				func() keystore.Store {
					store := mocks.NewMockStore(mockCtrl)
					store.EXPECT().GetAddresses().Return(
						[][]byte{
							testutil.MustHexDecodeString("5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761"),
							testutil.MustHexDecodeString("4cb0a77b76667dac586c40cc9523ace73b5d772bd503c63ed0ca596eae1658b2"),
						},
						nil,
					).Times(1)

					return store
				}(),
			},
			"{\"addresses\":[\"5602ea95540bee46d03ba335eed6f49d117eab95c8ab8b71bae2cdd1e564a761\",\"4cb0a77b76667dac586c40cc9523ace73b5d772bd503c63ed0ca596eae1658b2\"]}\n",
			http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetAddresses(tt.args.ks))

			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if !assert.Equal(tt.wantStatus, rr.Code) {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.wantStatus)
			}
			if !assert.Equal(tt.wantBody, rr.Body.String()) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.wantBody)
			}
		})
	}
}
