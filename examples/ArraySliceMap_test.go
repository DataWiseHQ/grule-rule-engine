//  Copyright DataWiseHQ/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package examples

import (
	"fmt"
	"github.com/DataWiseHQ/grule-rule-engine/ast"
	"github.com/DataWiseHQ/grule-rule-engine/builder"
	"github.com/DataWiseHQ/grule-rule-engine/engine"
	"github.com/DataWiseHQ/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArrayNode struct {
	Name        string
	StringArray []string
	NumberArray []int
	ChildArray  []*ArrayNode
}

func (node *ArrayNode) CallMyName() string {
	fmt.Println("You have call my name", node.Name)
	return node.Name
}

func (node *ArrayNode) GetChild(idx int64) *ArrayNode {
	return node.ChildArray[idx]
}

func TestArraySlice(t *testing.T) {
	//logrus.SetLevel(logrus.TraceLevel)
	Tree := &ArrayNode{
		Name:        "Node",
		StringArray: []string{"NodeString1", "NodeString2"},
		NumberArray: []int{235, 633},
		ChildArray: []*ArrayNode{
			&ArrayNode{
				Name:        "NodeChild1",
				StringArray: []string{"NodeChildString11", "NodeChildString12"},
				NumberArray: []int{578, 296},
				ChildArray: []*ArrayNode{
					&ArrayNode{
						Name:        "NodeChild11",
						StringArray: []string{"NodeChildString111", "NodeChildString112"},
						NumberArray: []int{578, 296},
						ChildArray:  nil,
					}, &ArrayNode{
						Name:        "NodeChild12",
						StringArray: []string{"NodeChildString121", "NodeChildString122"},
						NumberArray: []int{744, 895},
						ChildArray:  nil,
					},
				},
			}, &ArrayNode{
				Name:        "NodeChild2",
				StringArray: []string{"NodeChildString21", "NodeChildString22"},
				NumberArray: []int{744, 895},
				ChildArray:  nil,
			},
		},
	}

	rule := `
rule SetTreeName "Set the top most tree name" {
	when
		Tree.Name.ToUpper() == "NODE" &&
		Tree.StringArray[0].ToUpper() == "NODESTRING1" &&
		Tree.StringArray[1].ToLower() == "nodestring2" &&
		Tree.NumberArray[0] == 235 &&
		Tree.NumberArray[1] == 633 &&
		Tree.ChildArray[0].Name == "NodeChild1" &&
		Tree.ChildArray[0].CallMyName() == "NodeChild1" &&
		Tree.GetChild(0).ChildArray[0].Name == "NodeChild11" &&
		Tree.GetChild(0).ChildArray[0].CallMyName() == "NodeChild11" &&
		Tree.ChildArray[0].StringArray[1] == "NodeChildString12"
	then
		Tree.Name = "VERIFIED".ToLower();
		Tree.ChildArray[0].StringArray[0] = "SetSuccessful";
		Tree.NumberArray[1] = 1000;
		Tree.ChildArray[0].CallMyName();
		Retract("SetTreeName");
}
`

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Tree", Tree)
	assert.NoError(t, err)

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err = ruleBuilder.BuildRuleFromResource("TestFuncChaining", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	kb, err := lib.NewKnowledgeBaseInstance("TestFuncChaining", "0.0.1")
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "verified", Tree.Name)
	assert.Equal(t, "SetSuccessful", Tree.ChildArray[0].StringArray[0])
	assert.Equal(t, 1000, Tree.NumberArray[1])
}
