package main

import (
    "testing"
	"maps"
	"database/sql"
    "github.com/traPtitech/naro-template-backend/handler"
)

func Test_sumPopulationByCountryCode(t *testing.T) {
	// ここにテストケースを書いていく
	cases := []struct {
		name   string           // テストケースの名前
		cities []handler.City   // テストケースの入力
		want   map[string]int64  // 期待される結果
	}{
		{
			name:   "empty input",
			cities: []handler.City{},
			want:   map[string]int64{},
		},
		{
			name: "single country",
			cities: []handler.City{
				{
					CountryCode: sql.NullString{
						String: "JPN",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
			},
			want: map[string]int64{"JPN": 100},
		},
		{
			name: "multiple countries",
			cities: []handler.City{
				{
					CountryCode: sql.NullString{
						String: "JPN",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
				{
					CountryCode: sql.NullString{
						String: "USA",
						Valid:  true,
					},
					Population: sql.NullInt64{
						Int64: 200,
						Valid: true,
					},
				},
			},
			want: map[string]int64{"JPN": 100, "USA": 200},
		},
		{
			name: "empty country code",
			cities: []handler.City{
				{
					CountryCode: sql.NullString{
						String: "",
						Valid:  false,
					},
					Population: sql.NullInt64{
						Int64: 100,
						Valid: true,
					},
				},
			},
			want: map[string]int64{},
		},
	}
	for _, tt := range cases {
		// サブテストの実行
		t.Run(tt.name, func(t *testing.T) {
			got := handler.SumPopulationByCountryCode(tt.cities)
			if !maps.Equal(got, tt.want) {
				t.Errorf("SumPopulationByCountryCode(%v) = %v, want %v", tt.cities, got, tt.want)
			}
		})
	}
}