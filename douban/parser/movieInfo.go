package parser

import (
	"encoding/json"
	"good-spider/engine"
	"good-spider/model"
	"regexp"
)

const directorRegex = `<span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/[^"]*?" rel="v:directedBy">([^<]+?)</a></span></span><br/>`
const writerRegex = `<a href="/celebrity/[^"]*?">([^<]+?)</a>`
const starringRegex = `<a href="/celebrity/[^"]*?" rel="v:starring">([^<]+?)</a>`
const movieTypeRegex = `<span property="v:genre">([^<]+?)</span>`
const regionRegex = `<span class="pl">制片国家/地区:</span>([^<]+?)<br/>`
const languageRegex = `<span class="pl">语言:</span>([^<]+?)<br/>`
const releaseDateRegex = `<span property="v:initialReleaseDate" content="[^"]*?">([^<]+?)</span>`
const longRegex = `<span property="v:runtime" content="[^"]*?">([^<]+?)</span>`
const snameRegex = `<span class="pl">又名:</span>([^<]+?)<br/>`

var reDirctor = regexp.MustCompile(directorRegex)
var reWriter = regexp.MustCompile(writerRegex)
var reStarring = regexp.MustCompile(starringRegex)
var reMovieType = regexp.MustCompile(movieTypeRegex)
var reRegion = regexp.MustCompile(regionRegex)
var reLanguage = regexp.MustCompile(languageRegex)
var reReleaseDate = regexp.MustCompile(releaseDateRegex)
var reLong = regexp.MustCompile(longRegex)
var reSname = regexp.MustCompile(snameRegex)

func ParseMovieInfo(contents []byte, rankType model.RankType) engine.ParseResult {
	var movieInfo model.MovieInfo

	matches := reDirctor.FindAllSubmatch(contents, -1)
	movieInfo.Director = combine(matches)

	matches = reWriter.FindAllSubmatch(contents, -1)
	movieInfo.Screenwriter = combine(matches)

	matches = reStarring.FindAllSubmatch(contents, -1)
	movieInfo.Starring = combine(matches)

	matches = reMovieType.FindAllSubmatch(contents, -1)
	movieInfo.Type = combine(matches)

	matches = reRegion.FindAllSubmatch(contents, -1)
	movieInfo.Region = combine(matches)

	matches = reLanguage.FindAllSubmatch(contents, -1)
	movieInfo.Language = combine(matches)

	matches = reReleaseDate.FindAllSubmatch(contents, -1)
	movieInfo.ReleaseDate = combine(matches)

	matches = reLong.FindAllSubmatch(contents, -1)
	movieInfo.Long = combine(matches)

	matches = reSname.FindAllSubmatch(contents, -1)
	movieInfo.Sname = combine(matches)

	movieInfo.Title = rankType.Title
	movieInfo.ID = rankType.ID
	movieInfo.URL = rankType.URL

	result := engine.ParseResult{}
	jsonStr, err := json.Marshal(movieInfo)
	if err != nil {
		return result
	}
	result.Items = append(result.Items, movieInfo)
	println(string(jsonStr))

	return result
}

func combine(res [][][]byte) string {
	var ret string
	itemLimit := 0
	for _, m := range res {
		if itemLimit >= 5 {
			break
		}
		ret += string(m[1]) + "/"
		itemLimit++
	}
	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}

	return ret
}
