package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mvitta/GoRoutineMiniProject/fetch"
)

// https://app.quicktype.io/
type Book struct {
	NumFound      int64       `json:"numFound"`
	Start         int64       `json:"start"`
	NumFoundExact bool        `json:"numFoundExact"`
	Docs          []Doc       `json:"docs"`
	Book1NumFound int64       `json:"num_found"`
	Q             string      `json:"q"`
	Offset        interface{} `json:"offset"`
}

type Doc struct {
	AuthorAlternativeName            []string    `json:"author_alternative_name,omitempty"`
	AuthorKey                        []string    `json:"author_key,omitempty"`
	AuthorName                       []string    `json:"author_name,omitempty"`
	Contributor                      []string    `json:"contributor,omitempty"`
	CoverEditionKey                  *string     `json:"cover_edition_key,omitempty"`
	CoverI                           *int64      `json:"cover_i,omitempty"`
	Ddc                              []string    `json:"ddc,omitempty"`
	EbookAccess                      EbookAccess `json:"ebook_access"`
	EbookCountI                      int64       `json:"ebook_count_i"`
	EditionCount                     int64       `json:"edition_count"`
	EditionKey                       []string    `json:"edition_key,omitempty"`
	FirstPublishYear                 *int64      `json:"first_publish_year,omitempty"`
	FirstSentence                    []string    `json:"first_sentence,omitempty"`
	HasFulltext                      bool        `json:"has_fulltext"`
	Ia                               []string    `json:"ia,omitempty"`
	IaCollection                     []string    `json:"ia_collection,omitempty"`
	IaCollectionS                    *string     `json:"ia_collection_s,omitempty"`
	Isbn                             []string    `json:"isbn,omitempty"`
	Key                              string      `json:"key"`
	Language                         []string    `json:"language,omitempty"`
	LastModifiedI                    int64       `json:"last_modified_i"`
	Lcc                              []string    `json:"lcc,omitempty"`
	Lccn                             []string    `json:"lccn,omitempty"`
	LendingEditionS                  *string     `json:"lending_edition_s,omitempty"`
	LendingIdentifierS               *string     `json:"lending_identifier_s,omitempty"`
	NumberOfPagesMedian              *int64      `json:"number_of_pages_median,omitempty"`
	Oclc                             []string    `json:"oclc,omitempty"`
	PrintdisabledS                   *string     `json:"printdisabled_s,omitempty"`
	PublicScanB                      bool        `json:"public_scan_b"`
	PublishDate                      []string    `json:"publish_date,omitempty"`
	PublishPlace                     []string    `json:"publish_place,omitempty"`
	PublishYear                      []int64     `json:"publish_year,omitempty"`
	Publisher                        []string    `json:"publisher,omitempty"`
	Seed                             []string    `json:"seed"`
	Title                            string      `json:"title"`
	TitleSort                        string      `json:"title_sort"`
	TitleSuggest                     string      `json:"title_suggest"`
	Type                             Type        `json:"type"`
	IDLibrarything                   []string    `json:"id_librarything,omitempty"`
	IDGoodreads                      []string    `json:"id_goodreads,omitempty"`
	IDAmazon                         []string    `json:"id_amazon,omitempty"`
	IDDepósitoLegal                  []string    `json:"id_depósito_legal,omitempty"`
	IDAlibrisID                      []string    `json:"id_alibris_id,omitempty"`
	IDGoogle                         []string    `json:"id_google,omitempty"`
	IDPaperbackSwap                  []string    `json:"id_paperback_swap,omitempty"`
	IDWikidata                       []string    `json:"id_wikidata,omitempty"`
	IDOverdrive                      []string    `json:"id_overdrive,omitempty"`
	IDCanadianNationalLibraryArchive []string    `json:"id_canadian_national_library_archive,omitempty"`
	Subject                          []string    `json:"subject,omitempty"`
	Place                            []string    `json:"place,omitempty"`
	Time                             []string    `json:"time,omitempty"`
	Person                           []string    `json:"person,omitempty"`
	IaLoadedID                       []string    `json:"ia_loaded_id,omitempty"`
	IaBoxID                          []string    `json:"ia_box_id,omitempty"`
	RatingsAverage                   *float64    `json:"ratings_average,omitempty"`
	RatingsSortable                  *float64    `json:"ratings_sortable,omitempty"`
	RatingsCount                     *int64      `json:"ratings_count,omitempty"`
	RatingsCount1                    *int64      `json:"ratings_count_1,omitempty"`
	RatingsCount2                    *int64      `json:"ratings_count_2,omitempty"`
	RatingsCount3                    *int64      `json:"ratings_count_3,omitempty"`
	RatingsCount4                    *int64      `json:"ratings_count_4,omitempty"`
	RatingsCount5                    *int64      `json:"ratings_count_5,omitempty"`
	ReadinglogCount                  *int64      `json:"readinglog_count,omitempty"`
	WantToReadCount                  *int64      `json:"want_to_read_count,omitempty"`
	CurrentlyReadingCount            *int64      `json:"currently_reading_count,omitempty"`
	AlreadyReadCount                 *int64      `json:"already_read_count,omitempty"`
	PublisherFacet                   []string    `json:"publisher_facet,omitempty"`
	PersonKey                        []string    `json:"person_key,omitempty"`
	PlaceKey                         []string    `json:"place_key,omitempty"`
	TimeFacet                        []string    `json:"time_facet,omitempty"`
	PersonFacet                      []string    `json:"person_facet,omitempty"`
	SubjectFacet                     []string    `json:"subject_facet,omitempty"`
	Version                          float64     `json:"_version_"`
	PlaceFacet                       []string    `json:"place_facet,omitempty"`
	LccSort                          *string     `json:"lcc_sort,omitempty"`
	AuthorFacet                      []string    `json:"author_facet,omitempty"`
	SubjectKey                       []string    `json:"subject_key,omitempty"`
	TimeKey                          []string    `json:"time_key,omitempty"`
	DdcSort                          *string     `json:"ddc_sort,omitempty"`
	IDNla                            []string    `json:"id_nla,omitempty"`
	IDAmazonCoUkAsin                 []string    `json:"id_amazon_co_uk_asin,omitempty"`
	IDAmazonCAAsin                   []string    `json:"id_amazon_ca_asin,omitempty"`
	IDAmazonDeAsin                   []string    `json:"id_amazon_de_asin,omitempty"`
	IDBetterWorldBooks               []string    `json:"id_better_world_books,omitempty"`
	IDBritishNationalBibliography    []string    `json:"id_british_national_bibliography,omitempty"`
	IDAmazonItAsin                   []string    `json:"id_amazon_it_asin,omitempty"`
	IDBcid                           []string    `json:"id_bcid,omitempty"`
	IDScribd                         []string    `json:"id_scribd,omitempty"`
	IDHathiTrust                     []string    `json:"id_hathi_trust,omitempty"`
	IDBritishLibrary                 []string    `json:"id_british_library,omitempty"`
	IDBibliothèqueNationaleDeFrance  []string    `json:"id_bibliothèque_nationale_de_france,omitempty"`
	IDLibris                         []string    `json:"id_libris,omitempty"`
	IDDnb                            []string    `json:"id_dnb,omitempty"`
	Subtitle                         *string     `json:"subtitle,omitempty"`
}

type EbookAccess string

const (
	Borrowable    EbookAccess = "borrowable"
	NoEbook       EbookAccess = "no_ebook"
	Printdisabled EbookAccess = "printdisabled"
)

type Type string

const (
	Work Type = "work"
)

var HostClient string = "http://localhost:3000"

type Books struct {
	Book1 Book `json:"book1"`
	Book2 Book `json:"book2"`
}
type Response struct {
	Res Books `json:"res"`
}

func HandlerMain(w http.ResponseWriter, r *http.Request) {

	req1 := fetch.Fetch(os.Getenv("API_REQ_1"))
	req2 := fetch.Fetch(os.Getenv("API_REQ_1"))
	var book1 Book
	_ = json.Unmarshal(req1, &book1)
	var book2 Book
	_ = json.Unmarshal(req2, &book2)

	res, _ := json.Marshal(Response{Res: Books{Book1: book1, Book2: book2}})
	w.Header().Add("Access-Control-Allow-Origin", HostClient)
	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
