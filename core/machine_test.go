package core_test

import (
	. "github.com/zhyuri/afanty/core"

	"encoding/json"
)

var _ = Describe("Core", func() {

	var (
		exampleStateMachineJSON = []byte(`{
		  "Comment": "An example that adds two numbers.",
		  "StartAt": "Add",
		  "Version": "1.0",
		  "TimeoutSeconds": 10,
		  "States":
			{
				"Add": {
				  "Comment": "Test Task",
				  "Type": "Task",
				  "Resource": "arn:aws:lambda:us-east-1:123456789012:function:Add",
				  "End": true
				}
			}
		}`)
		sm  StateMachine
		err error
	)
	BeforeEach(func() {
		sm, err = NewStateMachineFromJSON(exampleStateMachineJSON)
	})
	Describe("Init StateMachine", func() {
		Context("When json parse successfully", func() {
			It("Should be no error", func() {
				Ω(err).ShouldNot(HaveOccurred())
			})
			It("Should parse json correctly", func() {
				Ω(sm.StartAt).Should(Equal("Add"))
				Ω(sm.States[sm.StartAt]).ShouldNot(BeNil())
			})
		})

		Context("When json is wrong", func() {
			BeforeEach(func() {
				sm, err = NewStateMachineFromJSON(exampleStateMachineJSON[:2])
			})
			It("Should have an error", func() {
				Ω(err).Should(HaveOccurred())
			})
			It("Should have default values", func() {
				Ω(sm.StartAt).Should(Equal("InitState"))
			})
			It("Should have default timeout", func() {
				Ω(sm.TimeoutSeconds).Should(Equal(int32(10)))
			})
		})

	})

	var (
		exampleDataJSON = json.RawMessage{}
		data            json.RawMessage
	)
	Describe("Execute StateMachine", func() {
		BeforeEach(func() {
			data = exampleDataJSON
			err = sm.Execute(&data)
		})
		Context("", func() {
			It("should be no error", func() {
				Skip("No lib supported now")
				Ω(err).Should(BeNil())
				Ω(sm.StartAt).Should(Equal("Add"))
				Ω(sm.States[sm.StartAt]).ShouldNot(BeNil())
			})
		})
	})

})
