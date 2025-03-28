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

const (
	rule2 = `
rule AgeNameCheck "test" {
    when
      Pogo.GetStringLength("9999") > 0  &&
      Pogo.Result == ""
    then
      Pogo.Result = "String len above 0";
}
`

	rule3 = `
rule AgeNameCheck "test"  salience 10{
    when
      Pogo.Compare(User.Name, "Calo")  
    then
      User.Name = "Success";
      Log(User.Name);
      Retract("AgeNameCheck");
}
`
)

// MyPoGo serve as example plain Plai Old Go Object.
type MyPoGo struct {
	Result string
}

// GetStringLength will return the length of provided string argument
func (p *MyPoGo) GetStringLength(sarg string) int {
	return len(sarg)
}

// Compare will compare the equality between the two string.
func (p *MyPoGo) Compare(t1, t2 string) bool {
	fmt.Println(t1, t2)
	return t1 == t2
}

// User is an example user struct.
type User struct {
	Name string
	Age  int
	Male bool
}

func TestMyPoGo_GetStringLength(t *testing.T) {
	dataContext := ast.NewDataContext()
	pogo := &MyPoGo{}
	err := dataContext.Add("Pogo", pogo)
	if err != nil {
		t.Fatal(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rule2)))
	assert.NoError(t, err)
	kb, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 1}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	assert.Equal(t, "String len above 0", pogo.Result)
}

func TestMyPoGo_Compare(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("User", user)
	if err != nil {
		t.Fatal(err)
	}
	err = dataContext.Add("Pogo", &MyPoGo{})
	if err != nil {
		t.Fatal(err)
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rule3)))
	assert.NoError(t, err)
	kb, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	assert.NoError(t, err)
	eng1 := &engine.GruleEngine{MaxCycle: 100}
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)
	t.Log(user)
	assert.Equal(t, "Success", user.Name, "User should have changed name by rule to Success, but %s", user.Name)
}
