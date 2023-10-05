package parser

import (
	"encoding/json"
	"good-spider/engine"
	"good-spider/model"
	"regexp"
)

const dirctorRegex = `<span ><span class='pl'>导演</span>: <span class='attrs'><a href="/celebrity/[^"]*?" rel="v:directedBy">([^<]+?)</a></span></span><br/>`
const writerRegex = `<a href="/celebrity/[^"]*?">([^<]+?)</a>`
const starringRegex = `<a href="/celebrity/[^"]*?" rel="v:starring">([^<]+?)</a>`
const movieTypeRegex = `<span property="v:genre">([^<]+?)</span>`
const regionRegex = `<span class="pl">制片国家/地区:</span>([^<]+?)<br/>`
const languageRegex = `<span class="pl">语言:</span>([^<]+?)<br/>`
const releaseDateRegex = `<span property="v:initialReleaseDate" content="[^"]*?">([^<]+?)</span>`
const longRegex = `<span property="v:runtime" content="[^"]*?">([^<]+?)</span>`
const snameRegex = `<span class="pl">又名:</span>([^<]+?)<br/>`

var reDirctor = regexp.MustCompile(dirctorRegex)
var reWriter = regexp.MustCompile(writerRegex)
var reStarring = regexp.MustCompile(starringRegex)
var reMovieType = regexp.MustCompile(movieTypeRegex)
var reRegion = regexp.MustCompile(regionRegex)
var reLanguage = regexp.MustCompile(languageRegex)
var reReleaseDate = regexp.MustCompile(releaseDateRegex)
var reLong = regexp.MustCompile(longRegex)
var reSname = regexp.MustCompile(snameRegex)

func ParseMovieInfo(contents []byte, title string) engine.ParseResult {
	var movieInfo model.MovieInfo

	matches := reDirctor.FindAllSubmatch(contents, -1)
	movieInfo.Director = combina(matches)

	matches = reWriter.FindAllSubmatch(contents, -1)
	movieInfo.Screenwriter = combina(matches)

	matches = reStarring.FindAllSubmatch(contents, -1)
	movieInfo.Starring = combina(matches)

	matches = reMovieType.FindAllSubmatch(contents, -1)
	movieInfo.Type = combina(matches)

	matches = reRegion.FindAllSubmatch(contents, -1)
	movieInfo.Region = combina(matches)

	matches = reLanguage.FindAllSubmatch(contents, -1)
	movieInfo.Language = combina(matches)

	matches = reReleaseDate.FindAllSubmatch(contents, -1)
	movieInfo.ReleaseDate = combina(matches)

	matches = reLong.FindAllSubmatch(contents, -1)
	movieInfo.Long = combina(matches)

	matches = reSname.FindAllSubmatch(contents, -1)
	movieInfo.Sname = combina(matches)

	movieInfo.Title = title

	result := engine.ParseResult{}
	jsonStr, err := json.Marshal(movieInfo)
	if err != nil {
		return result
	}
	result.Items = append(result.Items, string(jsonStr))
	println(string(jsonStr))

	return result
}

func combina(res [][][]byte) string {
	var ret string
	for _, m := range res {
		ret += string(m[1]) + "/"
	}
	if len(ret) > 0 {
		ret = ret[:len(ret)-1]
	}

	return ret
}
