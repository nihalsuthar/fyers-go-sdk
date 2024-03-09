package fyersapi

import (
	"reflect"
	"testing"
)

func TestQuotesQuery_makeString(t *testing.T) {
	type fields struct {
		Symbols []string
	}
	tests := []struct {
		name   string
		fields fields
		want   quotesQuery
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			fields: fields{
				Symbols: []string{"a", "b", "c"},
			},
			want: quotesQuery{"a,b,c"},
		},
		{
			name: "test2",
			fields: fields{
				Symbols: nil,
			},
			want: quotesQuery{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := QuotesQuery{
				Symbols: tt.fields.Symbols,
			}
			got := q.makeString()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuotesQuery.makeString() = %v, want %v", got, tt.want)
				t.Log(got)
			}
			t.Log(got)
		})
	}
}
