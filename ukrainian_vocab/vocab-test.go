package ukrainian_vocab_test

import (
	"testing"
)

// Test a large set of words for which we know
// the correct stemmed form.
//
func Test_LargeVocabulary(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{

		{"адміністратор", "адміністра"},
		{"абсолютний", "абсолют"},
		{"абетка", "абетк"},
		{"асоціація", "асоція"},
		{"автобіографію", "автобіограф"},
		{"автографом", "автограф"},
		{"автомобілі", "автомоб"},
		{"автомобіль", "автомобіл"},
		{"автор", "автор"},
		{"авторам", "автор"},
		{"авторитет", "авторитет"},
		{"авторитету", "авторитет"},
		{"авторитетні", "авторитет"},
		{"агроном", "агрон"},
		{"агрономии", "агроном"},
		{"ад", "ад"},
		{"ада", "ад"},
		{"адам", "ад"},
		{"адамант", "адамант"},
		{"адвокат", "адвокат"},
		{"адвокатом", "адвокат"},
		{"адвокатську", "адвокатськ"},
		{"адміністративний", "адміністративн"},
		{"адміністративні", "адміністративн"},
		{"адміністратор", "адміністратор"},
		{"адмірали", "адмірал"},
		{"ад", "ад"},
		{"адреса", "адрес"},
		{"адресовано", "адресова"},
		{"адресою", "адрес"},
		{"адресу", "адрес"},
		{"адська", "адськ"},
		{"азарт", "азарт"},
		{"азартний", "азарт"},
		{"абетка", "абетк"},
		{"території", "територі"},
		{"реклами", "реклам"},
		{"рекламою", "реклам"},
		{"реклама", "реклам"},
		{"рекламу", "реклам"},
		{"рекламувати", "реклам"},
		{"рекламний", "реклам"},
		{"рекламні", "реклам"},
		{"рекламами", "реклам"},
		{"рекламі", "реклам"},
	}
	for _, testCase := range testCases {
		result := ukrainian.Stem(testCase.in, true)
		if result != testCase.out {
			t.Errorf("Expected %v -> %v, but got %v", testCase.in, testCase.out, result)
		}
	}
}
