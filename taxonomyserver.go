package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"./data"
	"./model"
	"./go-tigertonic"
)

//MyRequest standard request
type MyRequest struct {
	ID    string      `json:"id"`
	Stuff interface{} `json:"stuff"`
}

//MyResponse standard response
type MyResponse struct {
	ID       string      `json:"id"`
	Count    int         `json:"count"`
	Response interface{} `json:"response"`
}
type Page struct {
	Title string
	Body  []byte
}

var (
	mysqlUser     = flag.String("mysqluser", "foo", "a mysql user")
	mysqlPassword = flag.String("mysqlpass", "bar", "the mysql password")
	cert          = flag.String("cert", "", "certificate pathname")
	key           = flag.String("key", "", "private key pathname")
	config        = flag.String("config", "", "pathname of JSON configuration file")
	listen        = flag.String("listen", "127.0.0.1:8002", "listen address")
)

func main() {

	flag.Parse()

	cors := tigertonic.NewCORSBuilder().AddAllowedOrigins(*listen) //.AddAllowedHeaders("Origin, X-Requested-With, Content-Type, Accept")

	mux := tigertonic.NewTrieServeMux()
	mux.Handle("POST", "/correlation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCorrelationHandler), "getCorrelationHandler", nil)))
	mux.Handle("POST", "/attributesPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesPerDimensionHandler), "getAttributesPerDimensionHandler", nil)))
	mux.Handle("POST", "/leafAttributesPerDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getLeafAttributesPerDimensionHandler), "getLeafAttributesPerDimensionHandler", nil)))
	mux.Handle("POST", "/allChildrenAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllChildrenAttributesHandler), "getAllChildrenAttributesHandler", nil)))
	mux.Handle("POST", "/allChildrenLeafAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllChildrenLeafAttributesHandler), "getAllChildrenLeafAttributesHandler", nil)))
	mux.Handle("POST", "/attributeRelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeRelationsHandler), "getAttributeRelationsHandler", nil)))
	mux.Handle("POST", "/interdimensionalRelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getInterdimensionalRelationsHandler), "getInterdimensionalRelationsHandler", nil)))
	mux.Handle("POST", "/intermediateAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIntermediateAttributesHandler), "getIntermediateAttributesHandler", nil)))
	mux.Handle("POST", "/majorAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMajorAttributesHandler), "getMajorAttributesHandler", nil)))
	mux.Handle("POST", "/savePositions", cors.Build(tigertonic.Timed(tigertonic.Marshaled(savePositionsHandler), "savePositionsHandler", nil)))
	mux.Handle("POST", "/saveMajorPositions", cors.Build(tigertonic.Timed(tigertonic.Marshaled(saveMajorPositionsHandler), "saveMajorPositionsHandler", nil)))
	mux.Handle("POST", "/save3DPositions", cors.Build(tigertonic.Timed(tigertonic.Marshaled(save3DPositionsHandler), "save3DPositionsHandler", nil)))
	mux.Handle("POST", "/saveMajor3DPositions", cors.Build(tigertonic.Timed(tigertonic.Marshaled(saveMajor3DPositionsHandler), "saveMajor3DPositionsHandler", nil)))
	mux.Handle("POST", "/saveEdgeBendPoints", cors.Build(tigertonic.Timed(tigertonic.Marshaled(saveEdgeBendPointsHandler), "saveEdgeBendPointsHandler", nil)))
	mux.Handle("POST", "/citationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsPerAttributeHandler), "getCitationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/citationsPerAttributeIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsPerAttributeIncludingChildrenHandler), "getCitationsPerAttributeIncludingChildrenHandler", nil)))
	mux.Handle("POST", "/addAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(addAttributeHandler), "addAttributeHandler", nil)))
	mux.Handle("POST", "/addDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(addDimensionHandler), "addDimensionHandler", nil)))
	mux.Handle("POST", "/removeAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(removeAttributeHandler), "removeAttributeHandler", nil)))
	mux.Handle("POST", "/removeDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(removeDimensionHandler), "removeDimensionHandler", nil)))
	mux.Handle("POST", "/renameAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(renameAttributeHandler), "renameAttributeHandler", nil)))
	mux.Handle("POST", "/renameDimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(renameDimensionHandler), "renameDimensionHandler", nil)))
	mux.Handle("POST", "/addTaxonomyRelation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(addTaxonomyRelationHandler), "addTaxonomyRelationHandler", nil)))
	mux.Handle("POST", "/removeTaxonomyRelation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(removeTaxonomyRelationHandler), "removeTaxonomyRelationHandler", nil)))
	mux.Handle("POST", "/updateTaxonomyRelationType", cors.Build(tigertonic.Timed(tigertonic.Marshaled(updateTaxonomyRelationTypeHandler), "updateTaxonomyRelationTypeHandler", nil)))
	mux.Handle("POST", "/updateTaxonomyRelationAnnotation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(updateTaxonomyRelationAnnotationHandler), "updateTaxonomyRelationAnnotationHandler", nil)))
	mux.Handle("GET", "/attribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributesHandler), "getAttributesHandler", nil)))
	mux.Handle("GET", "/leafAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getLeafAttributesHandler), "getLeafAttributesHandler", nil)))
	mux.Handle("GET", "/dimension", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getDimensionsHandler), "getDimensionsHandler", nil)))
	mux.Handle("GET", "/citation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationsHandler), "getCitationsHandler", nil)))
	mux.Handle("GET", "/citationCount", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountHandler), "getCitationCountHandler", nil)))
	mux.Handle("GET", "/citationCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountsHandler), "getCitationCountsHandler", nil)))
	mux.Handle("POST", "/updateCitationReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getUpdateCitationReferenceCountsHandler), "getUpdateCitationReferenceCountsHandler", nil)))
	mux.Handle("GET", "/citationCountsIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationCountsIncludingChildrenHandler), "getCitationCountsIncludingChildrenHandler", nil)))
	mux.Handle("GET", "/relationTypes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getRelationTypesHandler), "getRelationTypesHandler", nil)))
	mux.Handle("GET", "/conceptCorrelation", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsHandler), "getConceptCorrelationsHandler", nil)))
	mux.Handle("GET", "/conceptCorrelation3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelations3DHandler), "getConceptCorrelations3DHandler", nil)))
	mux.Handle("GET", "/conceptCorrelationWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsWithReferenceCountsHandler), "getConceptCorrelationsWithReferenceCountsHandler", nil)))
	mux.Handle("GET", "/conceptCorrelationWithReferenceCounts3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getConceptCorrelationsWithReferenceCounts3DHandler), "getConceptCorrelationsWithReferenceCounts3DHandler", nil)))
	mux.Handle("GET", "/allConceptCorrelations", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsHandler), "getAllConceptCorrelationsHandler", nil)))
	mux.Handle("GET", "/allConceptCorrelations3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelations3DHandler), "getAllConceptCorrelations3DHandler", nil)))
	mux.Handle("GET", "/allConceptCorrelationsWithReferenceCounts", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsWithReferenceCountsHandler), "getAllConceptCorrelationsWithReferenceCountsHandler", nil)))
	mux.Handle("GET", "/allConceptCorrelationsWithReferenceCounts3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAllConceptCorrelationsWithReferenceCounts3DHandler), "getAllConceptCorrelationsWithReferenceCounts3DHandler", nil)))
	mux.Handle("POST", "/parentRelationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(GetParentRelationsPerAttributeHandler), "GetParentRelationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/childRelationsPerAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(GetChildRelationsPerAttributeHandler), "GetChildRelationsPerAttributeHandler", nil)))
	mux.Handle("POST", "/sharedPapers", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersHandler), "getSharedPapersHandler", nil)))
	mux.Handle("POST", "/sharedPapers3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapers3DHandler), "getSharedPapers3DHandler", nil)))
	mux.Handle("POST", "/sharedPapersIncludingChildren", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersIncludingChildrenHandler), "getSharedPapersIncludingChildrenHandler", nil)))
	mux.Handle("POST", "/sharedPapersIncludingChildren3D", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getSharedPapersIncludingChildren3DHandler), "getSharedPapersIncludingChildren3DHandler", nil)))
	mux.Handle("POST", "/attributeDetails", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeDetailsHandler), "getAttributeDetailsHandler", nil)))
	mux.Handle("POST", "/citationDetails", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getCitationDetailsHandler), "getCitationDetailsHandler", nil)))
	mux.Handle("POST", "/mergeAttributes", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getMergeAttributesHandler), "getMergeAttributesHandler", nil)))
	mux.Handle("POST", "/forkAttribute", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getForkAttributeHandler), "getForkAttributeHandler", nil)))
	mux.Handle("GET", "/attributeCoverage", cors.Build(tigertonic.Timed(tigertonic.Marshaled(getAttributeCoverageHandler), "getAttributeCoverageHandler", nil)))
	mux.HandleFunc("GET", "/error.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/error.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/zoomInIcon.png", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/icons/zoom_in_128.png")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/zoomOutIcon.png", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/icons/zoom_out_128.png")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bluebird.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bluebird.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/multiselect.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/multiselect.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/three.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/three.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/threex.dynamictexture.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/threex.dynamictexture.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/threex.dynamictext2dobject.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/threex.dynamictext2dobject.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/DragControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/DragControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/TrackballControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/TrackballControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/PointerLockControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/PointerLockControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/FlyControls.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/three/FlyControls.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/stats.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/stats.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/cytoscape.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/jquery.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-context-menu.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3-context-menu.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-context-menu.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/d3-context-menu.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bootstrap.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/bootstrap.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/selectize.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/selectize.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/loginForm.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/loginForm.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/selectize.bootstrap3.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/selectize.bootstrap3.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/style.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/style.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fonts/glyphicons-halflings-regular.woff", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.woff")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fonts/glyphicons-halflings-regular.woff2", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.woff2")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/fonts/glyphicons-halflings-regular.ttf", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/fonts/glyphicons-halflings-regular.ttf")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-dialog.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/bootstrap-dialog.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/bootstrap-dialog.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/bootstrap-dialog.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.v2.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.v2.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3.v4.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3.v4.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/viz.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/viz.v1.1.2.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vis.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/vis.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vis.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/vis.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-undo-redo.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-undo-redo.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-clipboard.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-clipboard.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-graphml.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-graphml.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-view-utilities.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-view-utilities.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-context-menus.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-context-menus.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-context-menus.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-context-menus.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-cxtmenu.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-cxtmenu.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-autopan-on-drag.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-autopan-on-drag.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-edge-bend-editing.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-edge-bend-editing.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/CytoscapeEdgeEditation.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/CytoscapeEdgeEditation.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-edgehandles.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-edgehandles.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-expand-collapse.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-expand-collapse.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.js-navigator.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape.js-navigator.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-navigator.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-navigator.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-node-resize.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-node-resize.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/konva.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/konva.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-noderesize.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-noderesize.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-panzoom.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-panzoom.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape.js-panzoom.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape.js-panzoom.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.qtip.min.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/jquery.qtip.min.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/jquery.qtip.min.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/jquery.qtip.min.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-qtip.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-qtip.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-tip.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/js/d3-tip.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/d3-tip.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/d3-tip.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/cytoscape-snap-to-grid.js", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/extensions/cytoscape-snap-to-grid.js")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/treemap.css", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/src/css/treemap.css")
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, "%s", p)
	})
	// mux.Handle("GET","/",cors.Build(tigertonic.Timed(tigertonic.Marshaled(getIndexHandler), "getIndexHandler", nil)))
	mux.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/index.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix2D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_two_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix3D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_three_dimensional.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/conceptCorrelationMatrix3D2", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/conceptCorrelations_three_dimensional2.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/attributeCoverageMatrix", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/interactive/attributeCoverage.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/3D", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/3D/hierarchy.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/treemap", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/treemap.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/chord", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/chordMap.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/horizontal", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/BPHorizontal.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/vertical", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/BPVertical.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/correlationMap", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/index.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/coverage", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/correlationMap/attributeCoverage.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/taxonomyRelations", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/cytoscape/taxonomyRelations.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/scopus", func(w http.ResponseWriter, r *http.Request) {
		p := loadPage("frontend/taxonomy/scopus/scopusAPI.html")
		fmt.Fprintf(w, "%s", p)
	})
	mux.HandleFunc("GET", "/pdf/{file}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling pdf")
		fp := filepath.Clean(r.URL.Path)
		var p []byte
		if strings.Contains(fp, "system") {
			p = loadPage("frontend/pdfs/system.pdf")
		} else if strings.Contains(fp, "attack") {
			p = loadPage("frontend/pdfs/attack.pdf")
		} else if strings.Contains(fp, "defense") {
			p = loadPage("frontend/pdfs/defense.pdf")
		} else if strings.Contains(fp, "relations") {
			p = loadPage("frontend/pdfs/view-relations.pdf")
		}
		fmt.Printf("Printing pdf %d", len(p))
		b := bytes.NewBuffer(p)
		// stream straight to client(browser)
		w.Header().Set("Content-type", "application/pdf")

		if _, err := b.WriteTo(w); err != nil { // <----- here!
			fmt.Fprintf(w, "%s", err)
		}
		fmt.Println("handling pdf end")
		w.Write([]byte("PDF Generated"))
	})
	mux.HandleFunc("GET", "/png/{file}", func(w http.ResponseWriter, r *http.Request) {
		fp := filepath.Clean(r.URL.Path)
		var p []byte
		if strings.Contains(fp, "system") {
			p = loadPage("frontend/pngs/system.png")
		} else if strings.Contains(fp, "attack") {
			p = loadPage("frontend/pngs/attack.png")
		} else if strings.Contains(fp, "defense") {
			p = loadPage("frontend/pngs/defense.png")
		} else if strings.Contains(fp, "relations") {
			p = loadPage("frontend/pngs/relations.png")
		}

		b := bytes.NewBuffer(p)
		// stream straight to client(browser)
		w.Header().Set("Content-type", "application/png")

		if _, err := b.WriteTo(w); err != nil { // <----- here!
			fmt.Fprintf(w, "%s", err)
		}
		w.Write([]byte("PNG Generated"))
	})

	// c := &Config{}
	// if err := tigertonic.Configure(*config, c); nil != err {
	// 	checkErr(err)
	// }

	server := tigertonic.NewServer(*listen, tigertonic.Logged(mux, nil))
	go func() {
		var err error
		if "" != *cert && "" != *key {
			err = server.ListenAndServeTLS(*cert, *key)
		} else {
			err = server.ListenAndServe()
		}
		if nil != err {
			checkErr(err)
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	fmt.Println(<-ch)
	server.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCorrelationHandler(u *url.URL, h http.Header, correlationRequest *model.CorrelationRequest) (int, http.Header, *MyResponse, error) {
	if len(correlationRequest.Attributes) == 0 {
		return http.StatusNotAcceptable, nil,
			&MyResponse{"0", 0, "I need some attributes to produce correlations"}, nil
	}
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.ExportCorrelations(
		correlationRequest.Attributes, correlationRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(papers), papers}, nil
}
func getAttributesPerDimensionHandler(u *url.URL, h http.Header, attributesPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAttributesPerDimension(
		attributesPerDimensionRequest.TaxonomyID, attributesPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getLeafAttributesPerDimensionHandler(u *url.URL, h http.Header, attributesPerDimensionRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetLeafAttributesPerDimension(
		attributesPerDimensionRequest.TaxonomyID, attributesPerDimensionRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAllChildrenAttributesHandler(u *url.URL, h http.Header, allChildrenAttributesRequest *model.AllChildrenAttributesRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAllChildrenAttributes(
		allChildrenAttributesRequest.TaxonomyID, allChildrenAttributesRequest.Parent)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getAllChildrenLeafAttributesHandler(u *url.URL, h http.Header, allChildrenAttributesRequest *model.AllChildrenAttributesRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAllChildrenLeafAttributes(
		allChildrenAttributesRequest.TaxonomyID, allChildrenAttributesRequest.Parent)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributes), attributes}, nil
}
func getIntermediateAttributesHandler(u *url.URL, h http.Header, intermediateAttributesRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	intermediateAttributes, err := driver.GetIntermediateAttributes(
		intermediateAttributesRequest.TaxonomyID, intermediateAttributesRequest.MinValue, intermediateAttributesRequest.MaxValue)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(intermediateAttributes), intermediateAttributes}, nil
}
func getMajorAttributesHandler(u *url.URL, h http.Header, majorAttributesRequest *model.MajorAttributesRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	majorAttributes, err := driver.GetMajorAttributes(
		majorAttributesRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(majorAttributes), majorAttributes}, nil
}
func getAttributeRelationsHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeRelations, err := driver.GetAttributeRelationsPerDimension(
		attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeRelations), attributeRelations}, nil
}
func getInterdimensionalRelationsHandler(u *url.URL, h http.Header, interdimensionalRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeRelations, err := driver.GetInterdimensionalRelations(
		interdimensionalRelationsRequest.TaxonomyID)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", len(attributeRelations), attributeRelations}, nil
}
func savePositionsHandler(u *url.URL, h http.Header, savePositionsRequest *model.SavePositionsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	err := driver.SavePositions(
		savePositionsRequest.Positions)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", 1, err}, nil // TODO
}
func saveMajorPositionsHandler(u *url.URL, h http.Header, saveMajorPositionsRequest *model.SavePositionsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	err := driver.SaveMajorPositions(
		saveMajorPositionsRequest.Positions)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", 1, err}, nil // TODO
}
func save3DPositionsHandler(u *url.URL, h http.Header, save3DPositionsRequest *model.SavePositionsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	err := driver.Save3DPositions(
		save3DPositionsRequest.Positions)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", 1, err}, nil // TODO
}
func saveMajor3DPositionsHandler(u *url.URL, h http.Header, saveMajor3DPositionsRequest *model.SavePositionsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	err := driver.SaveMajor3DPositions(
		saveMajor3DPositionsRequest.Positions)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", 1, err}, nil // TODO
}
func saveEdgeBendPointsHandler(u *url.URL, h http.Header, saveEdgeBendPointsRequest *model.SaveEdgeBendPointsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.SaveEdgeBendPoints(
		saveEdgeBendPointsRequest.TaxonomyID, saveEdgeBendPointsRequest.AttributeSrc, saveEdgeBendPointsRequest.AttributeDest, saveEdgeBendPointsRequest.EdgeBendPoints, saveEdgeBendPointsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil,
		&MyResponse{"0", 1, result}, nil // TODO
}
func getCitationsPerAttributeHandler(u *url.URL, h http.Header, citationsPerAttributeRequest *model.CitationsPerAttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := driver.GetCitationsPerAttribute(citationsPerAttributeRequest.Attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func getCitationsPerAttributeIncludingChildrenHandler(u *url.URL, h http.Header, citationsPerAttributeRequest *model.CitationsPerAttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := driver.GetCitationsPerAttributeIncludingChildren(citationsPerAttributeRequest.Attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func addAttributeHandler(u *url.URL, h http.Header, attributeRequest *model.AttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute := model.Attribute{Text: attributeRequest.Text, X: attributeRequest.X, Y: attributeRequest.Y, XMajor: attributeRequest.XMajor, YMajor: attributeRequest.YMajor, Major: attributeRequest.Major, Dimension: attributeRequest.Dimension}
	result, err := driver.AddAttribute(attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func addDimensionHandler(u *url.URL, h http.Header, dimensionRequest *model.DimensionRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.AddDimension(dimensionRequest.Text)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func removeAttributeHandler(u *url.URL, h http.Header, attributeRequest *model.AttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute := model.Attribute{Text: attributeRequest.Text}
	result, err := driver.RemoveAttribute(attribute)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func removeDimensionHandler(u *url.URL, h http.Header, dimensionRequest *model.AttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	dimension := model.Dimension{Text: dimensionRequest.Text}
	result, err := driver.RemoveDimension(dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func renameAttributeHandler(u *url.URL, h http.Header, renameAttributeRequest *model.RenameAttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.RenameAttribute(renameAttributeRequest.PreviousName, renameAttributeRequest.NewName)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func renameDimensionHandler(u *url.URL, h http.Header, renameDimensionRequest *model.RenameAttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.RenameDimension(renameDimensionRequest.PreviousName, renameDimensionRequest.NewName)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func addTaxonomyRelationHandler(u *url.URL, h http.Header, taxonomyRelationRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := driver.AddTaxonomyRelation(relation)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func removeTaxonomyRelationHandler(u *url.URL, h http.Header, taxonomyRelationRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := driver.RemoveTaxonomyRelation(relation)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func updateTaxonomyRelationTypeHandler(u *url.URL, h http.Header, taxonomyRelationRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Relation: taxonomyRelationRequest.Text}
	result, err := driver.UpdateTaxonomyRelationType(relation)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func updateTaxonomyRelationAnnotationHandler(u *url.URL, h http.Header, taxonomyRelationRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relation := model.AttributeRelation{TaxonomyID: taxonomyRelationRequest.TaxonomyID, AttributeSrc: taxonomyRelationRequest.AttributeSrc, AttributeDest: taxonomyRelationRequest.AttributeDest, Dimension: taxonomyRelationRequest.Dimension, Annotation: taxonomyRelationRequest.Text}
	result, err := driver.UpdateTaxonomyRelationAnnotation(relation)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func getAttributesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetAllAttributes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(attributes), attributes}, nil
}
func getLeafAttributesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributes, err := driver.GetLeafAttributes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(attributes), attributes}, nil
}
func getDimensionsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	dimensions, err := driver.GetAllDimensions()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(dimensions), dimensions}, nil
}
func getCitationsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citations, err := driver.GetAllCitations()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citations), citations}, nil
}
func getCitationCountHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := driver.GetCitationCount()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citationCounts), citationCounts}, nil
}
func getCitationCountsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := driver.GetCitationCounts()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citationCounts), citationCounts}, nil
}
func getUpdateCitationReferenceCountsHandler(u *url.URL, h http.Header, updateReferenceCountsRequest *model.UpdateReferenceCountsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.UpdateCitationReferenceCounts(updateReferenceCountsRequest.ReferenceCounts)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func getCitationCountsIncludingChildrenHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationCounts, err := driver.GetCitationCountsIncludingChildren()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citationCounts), citationCounts}, nil
}
func getRelationTypesHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relationTypes, err := driver.GetRelationTypes()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relationTypes), relationTypes}, nil
}
func getConceptCorrelationsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelations()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelations3DHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelations3D()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelationsWithReferenceCountsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelationsWithReferenceCounts()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getConceptCorrelationsWithReferenceCounts3DHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetConceptRelationsWithReferenceCounts3D()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelations()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelations3DHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelations3D()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsWithReferenceCountsHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelationsWithReferenceCounts()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func getAllConceptCorrelationsWithReferenceCounts3DHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	conceptRelations, err := driver.GetAllConceptRelationsWithReferenceCounts3D()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(conceptRelations), conceptRelations}, nil
}
func GetParentRelationsPerAttributeHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relations, err := driver.GetParentRelationsPerAttribute(attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Text, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relations), relations}, nil
}
func GetChildRelationsPerAttributeHandler(u *url.URL, h http.Header, attributeRelationsRequest *model.AttributeRelationsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	relations, err := driver.GetChildRelationsPerAttribute(attributeRelationsRequest.TaxonomyID, attributeRelationsRequest.Text, attributeRelationsRequest.Dimension)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(relations), relations}, nil
}
func getSharedPapersHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapers(sharedPapersRequest.Text1, sharedPapersRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapers3DHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapers3D(sharedPapersRequest.Text1, sharedPapersRequest.Text2, sharedPapersRequest.Text3)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapersIncludingChildrenHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapersIncludingChildren(sharedPapersRequest.Text1, sharedPapersRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getSharedPapersIncludingChildren3DHandler(u *url.URL, h http.Header, sharedPapersRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	papers, err := driver.GetSharedPapersIncludingChildren3D(sharedPapersRequest.Text1, sharedPapersRequest.Text2, sharedPapersRequest.Text3)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(papers), papers}, nil
}
func getAttributeDetailsHandler(u *url.URL, h http.Header, attributeDetailsRequest *model.AttributeDetailsRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeDetails, err := driver.GetAttributeDetails(attributeDetailsRequest.Text)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(attributeDetails), attributeDetails}, nil
}
func getCitationDetailsHandler(u *url.URL, h http.Header, citationDetailsRequest *model.SharedPapersRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	citationDetails, err := driver.GetCitationDetails(citationDetailsRequest.Text1, citationDetailsRequest.Text2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(citationDetails), citationDetails}, nil
}
func getMergeAttributesHandler(u *url.URL, h http.Header, mergeAttributesRequest *model.MergeAttributesRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attribute1 := model.Attribute{Text: mergeAttributesRequest.Text1, Dimension: mergeAttributesRequest.Dimension1}
	attribute2 := model.Attribute{Text: mergeAttributesRequest.Text2, Dimension: mergeAttributesRequest.Dimension2}
	result, err := driver.MergeAttributes(attribute1, attribute2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func getForkAttributeHandler(u *url.URL, h http.Header, forkAttributeRequest *model.ForkAttributeRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	result, err := driver.ForkAttribute(forkAttributeRequest.Text, forkAttributeRequest.Dimension, forkAttributeRequest.Parents1, forkAttributeRequest.Parents2, forkAttributeRequest.Children1, forkAttributeRequest.Children2, forkAttributeRequest.Citations1, forkAttributeRequest.Citations2)
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", 1, result}, nil
}
func getAttributeCoverageHandler(u *url.URL, h http.Header, r *MyRequest) (int, http.Header, *MyResponse, error) {
	driver := data.InitClassificationDriver(*mysqlUser, *mysqlPassword)
	attributeCoverage, err := driver.GetAttributeCoverage()
	checkErr(err)
	return http.StatusOK, nil, &MyResponse{"0", len(attributeCoverage), attributeCoverage}, nil
}

func loadPage(filename string) (body []byte) {
	// fmt.Println("handling loadpage")
	body, err := ioutil.ReadFile(filename)
	// fmt.Printf("%d", len(body))
	checkErr(err)
	return body
}
