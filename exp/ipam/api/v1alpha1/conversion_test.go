/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"

	ipamv1 "sigs.k8s.io/cluster-api/exp/ipam/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
)

func TestFuzzyConversion(t *testing.T) {
	t.Run("for IPAddress", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:         &ipamv1.IPAddress{},
		Spoke:       &IPAddress{},
		FuzzerFuncs: []fuzzer.FuzzerFuncs{},
	}))
	t.Run("for IPAddressClaim", utilconversion.FuzzTestFunc(utilconversion.FuzzTestFuncInput{
		Hub:         &ipamv1.IPAddressClaim{},
		Spoke:       &IPAddressClaim{},
		FuzzerFuncs: []fuzzer.FuzzerFuncs{},
	}))
}
