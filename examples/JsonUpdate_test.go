package examples

import (
	"encoding/json"
	"testing"

	"github.com/DataWiseHQ/grule-rule-engine/ast"
	"github.com/DataWiseHQ/grule-rule-engine/builder"
	"github.com/DataWiseHQ/grule-rule-engine/engine"
	"github.com/DataWiseHQ/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
)

func TestUpdatingJsonField(t *testing.T) {
	dataContext := ast.NewDataContext()

	byteSlice := []byte(`
	{
	"testField": 100, 
	"returnField": false
	}`)

	err := dataContext.AddJSON("json", byteSlice)
	assert.NoError(t, err)

	rule := `
	rule CheckReturnValue {
	   	when
	   		json.testField == 100
	   	then
	   		json.returnField = true;
			Retract("CheckReturnValue");
	}`

	// Prepare knowledgebase library and load it with our rule.
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)

	err = rb.BuildRuleFromResource("TestJsonFact", "0.0.1", pkg.NewBytesResource([]byte(rule)))

	eng1 := &engine.GruleEngine{MaxCycle: 5}

	kb, err := lib.NewKnowledgeBaseInstance("TestJsonFact", "0.0.1")

	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)

	jsonFact := make(map[string]any)

	err = json.Unmarshal(byteSlice, &jsonFact)

	assert.Equal(t, true, jsonFact["returnField"].(bool),
		"Returned Value not Updated from Rule Engine Execution")

}
