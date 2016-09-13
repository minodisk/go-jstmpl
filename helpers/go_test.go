package helpers

import (
	"testing"

	"github.com/lestrrat/go-jsschema"
)

type TestCaseConvertTagsForGo struct {
	ColumnName string
	Name       string
	Expect     string
	Title      string
}

type TestCaseExtras struct {
	Extras           map[string]interface{}
	ExpectColumnName string
	ExpectTableName  string
	ExpectDbType     string
	ExpectGoType     string
	Title            string
}

func TestConvertTagsForGo(t *testing.T) {
	tests := []TestCaseConvertTagsForGo{{
		ColumnName: "test_column",
		Name:       "test_name",
		Expect:     "`json:\"test_name, omitempty\" xorm:\"test_column\"`",
		Title:      "pass all column have value test",
	}, {
		ColumnName: "",
		Name:       "test_name",
		Expect:     "`json:\"test_name, omitempty\" xorm:\"-\"`",
		Title:      "pass ColumnName column empty test",
	}, {
		ColumnName: "test_column",
		Name:       "",
		Expect:     "`json:\"-\" xorm:\"test_column\"`",
		Title:      "pass Name column empty test",
	}, {
		ColumnName: "",
		Name:       "",
		Expect:     "`json:\"-\" xorm:\"-\"`",
		Title:      "pass all column empty test",
	}}

	for _, test := range tests {
		s := ConvertTagsForGo(test.Name, test.ColumnName)
		if test.Expect != s {
			t.Errorf("%s: Expect: %s, Result: %s", test.Title, test.Expect, s)
		}
	}
}

func TestGetExtraData(t *testing.T) {

	tests := []TestCaseExtras{{
		Extras: map[string]interface{}{
			"go_type": "go_type_test",
			"column": map[string]interface{}{
				"db_type": "db_type_test",
				"name":    "column_name_test",
			}},
		ExpectGoType:     "go_type_test",
		ExpectDbType:     "db_type_test",
		ExpectColumnName: "column_name_test",
		Title:            "pass all column have value test",
	}, {
		Extras: map[string]interface{}{
			"go_type": "go_type_test",
		},
		ExpectGoType: "go_type_test",
		Title:        "pass go_type column have value test",
	}, {
		Extras: map[string]interface{}{
			"column": map[string]interface{}{
				"db_type": "db_type_test",
				"name":    "column_name_test",
			},
		},
		ExpectDbType:     "db_type_test",
		ExpectColumnName: "column_name_test",
		Title:            "pass column column have value test",
	}, {
		Extras: map[string]interface{}{
			"table": map[string]interface{}{
				"name": "table_name_test",
			},
		},
		ExpectTableName: "table_name_test",
		Title:           "pass table column have value test",
	}}

	for _, test := range tests {
		s := schema.New()
		s.Extras = test.Extras
		gt, err := GetGoTypeData(s)
		if err != nil {
			t.Errorf("%s: in GetGoTypeData, Extras: %s, Error: %s", test.Title, test.Extras, err)
		}

		cn, ct, err := GetColumnData(s)
		if err != nil {
			t.Errorf("%s: in GetColumnData, Extras: %s, Error: %s", test.Title, test.Extras, err)
		}

		tn, err := GetTableData(s)
		if err != nil {
			t.Errorf("%s: in GetTableData, Extras: %s, Error: %s", test.Title, test.Extras, err)
		}

		if gt != test.ExpectGoType || cn != test.ExpectColumnName || ct != test.ExpectDbType || tn != test.ExpectTableName {
			t.Errorf("%s: Expect(%s, %s, %s, %s), Result(%s, %s, %s, %s)", test.Title,
				test.ExpectGoType, test.ExpectColumnName, test.ExpectDbType, test.ExpectTableName,
				gt, cn, ct, tn)
		}
	}
}
