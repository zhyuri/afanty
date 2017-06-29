package core_test

import (
	. "github.com/zhyuri/afanty/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var (
		passStateJSON = []byte(`{
		  "Type": "Pass",
		  "ResultPath": "$.coords",
		  "Next": "End"
		}`)
		stateInter interface{}
		passState  *PassState
		err        error
		ok         bool
	)

	BeforeEach(func() {
		stateInter, err = BuildState(passStateJSON)
	})
	Describe("Parse the json of state", func() {
		Context("When json is correct", func() {
			It("Should not be any error", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
		Context("When assert type", func() {
			BeforeEach(func() {
				passState, ok = stateInter.(*PassState)
			})
			It("Should have correct type", func() {
				Ω(ok).Should(Equal(true))
			})
			It("Should have correct value", func() {
				Ω(passState.Type).Should(Equal(Name_PassState))
				Ω(passState.ResultPath).Should(Equal("$.coords"))
				Ω(passState.Next).Should(Equal("End"))
			})
		})

	})
})
