package model

type Taxonomy struct {
	ID             int       `json:"id"`
	Text      string    `json:"text"`
  Dimensions []Dimension `json:"dimensions"`
}

type Dimension struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  X string `json:"x"`
  Y string `json:"y"`
  XMajor string `json:"xMajor"`
  YMajor string `json:"yMajor"`
  Attributes []Attribute `json:"attributes"`
}

type Attribute struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  ParentID int `json:"parentID"`
  ParentText string `json:"parentText"`
  X string `json:"x"`
  Y string `json:"y"`
  XMajor string `json:"xMajor"`
  YMajor string `json:"xMajor"`
  X3D string `json:"x3D"`
  Y3D string `json:"y3D"`
  Z3D string `json:"z3D"`
  XMajor3D string `json:"xMajor3D"`
  YMajor3D string `json:"yMajor3D"`
  ZMajor3D string `json:"zMajor3D"`
  Major int8 `json:"major"`
  Dimension string `json:"dimension"`
  MappedPapers []Paper `json:"papers"`
  Relations []Relation `json:"relations"`
}
type Paper struct {
  ID int `json:"id"`
  Citation string `json:"citation"`
  Text string `json:"text"`
  Attributes []Attribute `json:"attributes"`
	StrAttributes string `json:"str_attributes"`
	Bib string `json:"bib"`
  ReferenceCount int64 `json:"referenceCount"`
}

type Position struct {
  ID string `json:"id"`
  Table string `json:"table"`
  X string `json:"x"`
  Y string `json:"y"`
  Z string `json:"z"`
}

type AttributeRequest struct{
  Text string  `json:"text"`
  X string `json:"x"`
  Y string `json:"y"`
  XMajor string `json:"xMajor"`
  YMajor string `json:"xMajor"`
  Major int8 `json:"major"`
  Dimension string `json:"dimension"`
}

type CorrelationRequest struct {
	TaxonomyID int64 `json:"taxonomy_id"`
	Attributes []Attribute `json:"attributes"`
}

type RenameAttributeRequest struct {
  PreviousName string `json:"previousName"`
  NewName string `json:"newName"`
}

type AttributeRelationsRequest struct {
  TaxonomyID int64 `json:"taxonomy_id"`
  AttributeSrc string `json:"attributeSrc"`
  AttributeDest string `json:"attributeDest"`
  Text string `json:"text"`
  Dimension string `json:"dimension"`
  MinValue int64 `json:"minValue"`
  MaxValue int64 `json:"maxValue"`
}

type MajorAttributesRequest struct {
  TaxonomyID int64 `json:"taxonomy_id"`
}

type AllChildrenAttributesRequest struct {
  TaxonomyID int64 `json:"taxonomy_id"`
  Parent string `json:"parent"`
}

type SavePositionsRequest struct {
  Positions []Position `json:"positions"`
}

type SaveEdgeBendPointsRequest struct {
  AttributeSrc string `json:"attributeSrc"`
  AttributeDest string `json:"attributeDest"`
  EdgeBendPoints string `json:"edgeBendPoints"`
}

type CitationsPerAttributeRequest struct {
  Attribute string `json:"attribute"`
}

type SharedPapersRequest struct {
  Text1 string `json:"text1"`
  Text2 string `json:"text2"`
  Text3 string `json:"text3"`
}

type AttributeDetailsRequest struct {
  Text string `json:"text"`
}

type MergeAttributesRequest struct {
  Text1 string  `json:"text1"`
  Text2 string  `json:"text2"`
  Dimension1 string `json:"dimension1"`
  Dimension2 string `json:"dimension2"`
}

type ForkAttributeRequest struct {
  Text string `json:"text"`
  Dimension string `json:"dimension"`
  Parents1 []AttributeRelation `json:"parents1"`
  Parents2 []AttributeRelation `json:"parents2"`
  Children1 []AttributeRelation `json:"children1"`
  Children2 []AttributeRelation `json:"children2"`
  Citations1 []Paper `json:"citations1"`
  Citations2 []Paper `json:"citations2"`
}

type Relation struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  Comment string  `json:"comment"`
  SourceAttribute Attribute `json:"src_attribute"`
  DestinationAttribute Attribute `json:"dest_attribute"`
}

type RelationType struct{
  ID int `json:"id"`
  Text string  `json:"text"`
  Comment string `json:"comment"`
}

type ConceptCorrelation struct {
  Attribute1 string `json:"attribute1"`
  Attribute2 string `json:"attribute2"`
  Attribute3 string `json:"attribute3"`
  Text1 string `json:"text1"`
  Text2 string `json:"text2"`
  Text3 string `json:"text3"`
  ID1 int `json:"id1"`
  ID2 int `json:"id2"`
  ID3 int `json:"id3"`
  Value int `json:"value"`
}

type AttributeCoverage struct {
  AttributeName string `json:"attributeName"`
  PaperName string `json:"paperName"`
  Text1 string `json:"text1"`
  Text2 string `json:"text2"`
  AttributeID int `json:"attributeID"`
  PaperID int `json:"paperID"`
  Value int `json:"value"`
}

type AttributeRelation struct {
  TaxonomyID int64 `json:"taxonomy_id"`
  AttributeSrc string `json:"attributeSrc"`
  AttributeDest string `json:"attributeDest"`
  Text string `json:"text"`
  DimensionSrc string `json:"dimensionSrc"`
  DimensionDest string `json:"dimensionDest"`
  Dimension string `json:"dimension"`
  Relation string `json:"relation"`
  EdgeBendPoints string `json:"edgeBendPoints"`
  Annotation string `json:"annotation"`
}

type CitationCount struct {
  Attribute string `json:"attribute"`
  CitationCount int64 `json:"citationCount"`
  MaxReferenceCount int64 `json:"maxReferenceCount"`
}

type Result struct {
  Success bool `json:"success"`
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
