package path

import (
	"encoding/json"
	"testing"
)

var data = []byte(`
{
  "_links": {
    "self": {
      "href": "https://us.api.blizzard.com/data/wow/connected-realm/3726?namespace=dynamic-us"
    }
  },
  "id": 3726,
  "has_queue": false,
  "status": {
    "type": "UP",
    "name": {
      "en_US": "Up",
      "es_MX": "Disponible",
      "pt_BR": "Para Cima",
      "de_DE": "Verfügbar",
      "en_GB": "Up",
      "es_ES": "Activo",
      "fr_FR": "En ligne",
      "it_IT": "Attivo",
      "ru_RU": "Работает",
      "ko_KR": "정상",
      "zh_TW": "正常",
      "zh_CN": "正常"
    }
  },
  "population": {
    "type": "HIGH",
    "name": {
      "en_US": "High",
      "es_MX": "Alto",
      "pt_BR": "Alto",
      "de_DE": "Hoch",
      "en_GB": "High",
      "es_ES": "Alto",
      "fr_FR": "Élevée",
      "it_IT": "Alta",
      "ru_RU": "Высокая",
      "ko_KR": "혼잡",
      "zh_TW": "高",
      "zh_CN": "高"
    }
  },
  "realms": [
    {
      "id": 3726,
      "region": {
        "key": {
          "href": "https://us.api.blizzard.com/data/wow/region/1?namespace=dynamic-us"
        },
        "name": {
          "en_US": "North America",
          "es_MX": "Norteamérica",
          "pt_BR": "América do Norte",
          "de_DE": "Nordamerika",
          "en_GB": "North America",
          "es_ES": "Norteamérica",
          "fr_FR": "Amérique du Nord",
          "it_IT": "Nord America",
          "ru_RU": "Северная Америка",
          "ko_KR": "미국",
          "zh_TW": "北美",
          "zh_CN": "北美"
        },
        "id": 1
      },
      "connected_realm": {
        "href": "https://us.api.blizzard.com/data/wow/connected-realm/3726?namespace=dynamic-us"
      },
      "name": {
        "en_US": "Khaz'goroth",
        "es_MX": "Khaz'goroth",
        "pt_BR": "Khaz'goroth",
        "de_DE": "Khaz'goroth",
        "en_GB": "Khaz'goroth",
        "es_ES": "Khaz'goroth",
        "fr_FR": "Khaz'goroth",
        "it_IT": "Khaz'goroth",
        "ru_RU": "Khaz'goroth US9",
        "ko_KR": "Khaz'goroth",
        "zh_TW": "卡茲格羅斯",
        "zh_CN": "卡兹格罗斯"
      },
      "category": {
        "en_US": "Oceanic",
        "es_MX": "Oceánico",
        "pt_BR": "Oceania",
        "de_DE": "Ozeanisch",
        "en_GB": "Oceanic",
        "es_ES": "Oceánico",
        "fr_FR": "Océanique",
        "it_IT": "Oceania",
        "ru_RU": "Океания",
        "ko_KR": "오세아니아",
        "zh_TW": "大洋洲",
        "zh_CN": "Oceanic"
      },
      "locale": "enUS",
      "timezone": "Australia/Melbourne",
      "type": {
        "type": "NORMAL",
        "name": {
          "en_US": "Normal",
          "es_MX": "Normal",
          "pt_BR": "Normal",
          "de_DE": "Normal",
          "en_GB": "Normal",
          "es_ES": "Normal",
          "fr_FR": "Normal",
          "it_IT": "Normale",
          "ru_RU": "Обычный",
          "ko_KR": "일반",
          "zh_TW": "一般",
          "zh_CN": "普通"
        }
      },
      "is_tournament": false,
      "slug": "khazgoroth"
    },
    {
      "id": 3722,
      "region": {
        "key": {
          "href": "https://us.api.blizzard.com/data/wow/region/1?namespace=dynamic-us"
        },
        "name": {
          "en_US": "North America",
          "es_MX": "Norteamérica",
          "pt_BR": "América do Norte",
          "de_DE": "Nordamerika",
          "en_GB": "North America",
          "es_ES": "Norteamérica",
          "fr_FR": "Amérique du Nord",
          "it_IT": "Nord America",
          "ru_RU": "Северная Америка",
          "ko_KR": "미국",
          "zh_TW": "北美",
          "zh_CN": "北美"
        },
        "id": 1
      },
      "connected_realm": {
        "href": "https://us.api.blizzard.com/data/wow/connected-realm/3726?namespace=dynamic-us"
      },
      "name": {
        "en_US": "Aman'Thul",
        "es_MX": "Aman'Thul",
        "pt_BR": "Aman'Thul",
        "de_DE": "Aman'thul",
        "en_GB": "Aman'Thul",
        "es_ES": "Aman'thul",
        "fr_FR": "Aman'Thul",
        "it_IT": "Aman'Thul",
        "ru_RU": "Aman'Thul US9",
        "ko_KR": "Aman'Thul",
        "zh_TW": "阿曼蘇爾",
        "zh_CN": "阿曼苏尔"
      },
      "category": {
        "en_US": "Oceanic",
        "es_MX": "Oceánico",
        "pt_BR": "Oceania",
        "de_DE": "Ozeanisch",
        "en_GB": "Oceanic",
        "es_ES": "Oceánico",
        "fr_FR": "Océanique",
        "it_IT": "Oceania",
        "ru_RU": "Океания",
        "ko_KR": "오세아니아",
        "zh_TW": "大洋洲",
        "zh_CN": "Oceanic"
      },
      "locale": "enUS",
      "timezone": "Australia/Melbourne",
      "type": {
        "type": "NORMAL",
        "name": {
          "en_US": "Normal",
          "es_MX": "Normal",
          "pt_BR": "Normal",
          "de_DE": "Normal",
          "en_GB": "Normal",
          "es_ES": "Normal",
          "fr_FR": "Normal",
          "it_IT": "Normale",
          "ru_RU": "Обычный",
          "ko_KR": "일반",
          "zh_TW": "一般",
          "zh_CN": "普通"
        }
      },
      "is_tournament": false,
      "slug": "amanthul"
    },
    {
      "id": 3735,
      "region": {
        "key": {
          "href": "https://us.api.blizzard.com/data/wow/region/1?namespace=dynamic-us"
        },
        "name": {
          "en_US": "North America",
          "es_MX": "Norteamérica",
          "pt_BR": "América do Norte",
          "de_DE": "Nordamerika",
          "en_GB": "North America",
          "es_ES": "Norteamérica",
          "fr_FR": "Amérique du Nord",
          "it_IT": "Nord America",
          "ru_RU": "Северная Америка",
          "ko_KR": "미국",
          "zh_TW": "北美",
          "zh_CN": "北美"
        },
        "id": 1
      },
      "connected_realm": {
        "href": "https://us.api.blizzard.com/data/wow/connected-realm/3726?namespace=dynamic-us"
      },
      "name": {
        "en_US": "Dath'Remar",
        "es_MX": "Dath'Remar",
        "pt_BR": "Dath'Remar",
        "de_DE": "Dath'Remar",
        "en_GB": "Dath'Remar",
        "es_ES": "Dath'Remar",
        "fr_FR": "Dath'Remar",
        "it_IT": "Dath'Remar",
        "ru_RU": "Dath'Remar US9",
        "ko_KR": "Dath'Remar",
        "zh_TW": "達斯雷瑪",
        "zh_CN": "达斯雷玛"
      },
      "category": {
        "en_US": "Oceanic",
        "es_MX": "Oceánico",
        "pt_BR": "Oceania",
        "de_DE": "Ozeanisch",
        "en_GB": "Oceanic",
        "es_ES": "Oceánico",
        "fr_FR": "Océanique",
        "it_IT": "Oceania",
        "ru_RU": "Океания",
        "ko_KR": "오세아니아",
        "zh_TW": "大洋洲",
        "zh_CN": "Oceanic"
      },
      "locale": "enUS",
      "timezone": "Australia/Melbourne",
      "type": {
        "type": "NORMAL",
        "name": {
          "en_US": "Normal",
          "es_MX": "Normal",
          "pt_BR": "Normal",
          "de_DE": "Normal",
          "en_GB": "Normal",
          "es_ES": "Normal",
          "fr_FR": "Normal",
          "it_IT": "Normale",
          "ru_RU": "Обычный",
          "ko_KR": "일반",
          "zh_TW": "一般",
          "zh_CN": "普通"
        }
      },
      "is_tournament": false,
      "slug": "dathremar"
    }
  ],
  "mythic_leaderboards": {
    "href": "https://us.api.blizzard.com/data/wow/connected-realm/3726/mythic-leaderboard/?namespace=dynamic-us"
  },
  "auctions": {
    "href": "https://us.api.blizzard.com/data/wow/connected-realm/3726/auctions?namespace=dynamic-us"
  }
}
`)

var (
	res interface{}

	p1 = "/realms/id"
	p2 = "/realms/slug"
	p3 = "/realms/name/en_US"
	p4 = "/groups/items/id"
)

//var results = make(map[string]interface{}, 0)
func BenchmarkJson_Find(b *testing.B) {
	j, _ := JsonParse(data)

	path, _ := ParsePath(p1, "/")
	path2, _ := ParsePath(p2, "/")
	path3, _ := ParsePath(p3, "/")
	path4, _ := ParsePath(p4, "/")

	for n := 0; n < b.N; n++ {
		_ = j.FindValues(path, path2, path3, path4)
		res = path.values
		res = path2.values
		res = path3.values
		res = path4.values
	}
}

var dst = make(map[string]interface{})

func BenchmarkJson_Unmarshal(b *testing.B) {
	_ = json.Unmarshal(data, &dst)

	for n := 0; n < b.N; n++ {
		realms := dst["realms"].([]interface{})
		for _, r := range realms {
			realm := r.(map[string]interface{})

			//res = realm["id"].(float64)
			//res = realm["slug"].(string)
			//res =realm["name"].(map[string]interface{})["en_US"].(string)

			res = realm["id"]
			res = realm["slug"]
			res = realm["name"].(map[string]interface{})["en_US"]
		}
	}
}

func Test_tokenize(t *testing.T) {
	var err error

	j, err := JsonParse(data)
	if err != nil {
		t.Error(err)
	}

	path, _ := ParsePath(p1, "/")
	path2, _ := ParsePath(p2, "/")
	path3, _ := ParsePath(p3, "/")

	err = j.FindValues(path, path2, path3)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < len(path.values); i++ {
		t.Logf("%s:%v", path.raw, path.values[i])
		t.Logf("%s:%v", path2.raw, path2.values[i])
		t.Logf("%s:%v", path3.raw, path3.values[i])
	}
}
