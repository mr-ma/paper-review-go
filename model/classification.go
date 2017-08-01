package model

type Taxonomy struct {
	ID             int       `json:"id"`
	Text      string    `json:"text"`
  Dimensions []Dimension `json:"dimensions"`
}

type Dimension struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  Attributes []Attribute `json:"attributes"`
}

type Attribute struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  MappedPapers []Paper `json:"papers"`
  Relations []Relation
}
type Paper struct {
  ID int `json:"id"`
  Citation string `json:"citation"`
  Attributes []Attribute `json:"attributes"`
	StrAttributes string `json:"str_attributes"`
	Bib string `json:"bib"`
}

type CorrelationRequest struct {
	TaxonomyID int64 `json:"taxonomy_id"`
	Attributes []Attribute `json:"attributes"`
}

type Relation struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  Comment string  `json:"comment"`
  SourceAttribute Attribute `json:"src_attribute"`
  DestinationAttribute Attribute `json:"dest_attribute"`
}

//
// type VoteStatus int
//
// const (
// 	//UNSURE in vote
// 	UNSURE VoteStatus = -1
// 	//NO in vote
// 	No VoteStatus = 0
// 	//YES in vote
// 	YES VoteStatus = 1
// )
//
// type Research struct {
// 	ID             int       `json:"id"`
// 	Questions      string    `json:"questions"`
// 	ReviewTemplate string    `json:"review_template"`
// 	Articles       []Article `json:"articles"`
// 	Title          string    `json:"title"`
// }
// type Article struct {
// 	ID                   int    `json:"id"`
// 	Title                string `json:"title"`
// 	Authors              string `json:"authors"`
// 	Year                 string `json:"year"`
// 	CitedBy              string `json:"cited_by"`
// 	Keywords             string `json:"keywords"`
// 	Abstract             string `json:"abstract"`
// 	Journal              string `json:"journal"`
// 	File                 string `json:"file"`
// 	Source               string `json:"source"`
// 	AssociatedResearchId int64  `json:"associated_research_id"`
// }
// type Mitarbeiter struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	PassHash string `json:"pass_hash"`
// }
// type Tag struct {
// 	ID         int64  `json:"id"`
// 	Text       string `json:"text"`
// 	ResearchID int    `json:"research_id"`
// }
// type Vote struct {
// 	ID                  int         `json:"id"`
// 	State               VoteStatus  `json:"state"`
// 	Voter               Mitarbeiter `json:"voter"`
// 	Tags                []Tag       `json:"tags"`
// 	AssociatedArticleID int         `json:"associated_article_id"`
// 	Review              string      `json:"review"`
// }
//
// //Stats enholds statistic info about reviews
// type Stats struct {
// 	ReviewedArticles  int `json:"reviewed_articles"`
// 	RemainingArticles int `json:"remaining_articles"`
// 	MitarbeiterID     int `json:"mitarbeiter_id"`
// }
