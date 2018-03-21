// functions used to convert a HTML file to a "static" version that is independent from outside sources (like the database)
// variables are being converted to strings via JSON.stringify and put as strings into the HTML file
function exportHTML ( fileName ) {
  if (IS_STATIC) {
    handleError('Cannot export a HTML file, that is already "static", to "static" HTML.');
    return;
  }
  const DELIMITER = '<!-- main script end -->';
  var html = document.documentElement.outerHTML;
  var docType = new XMLSerializer().serializeToString(document.doctype);
  var fullHTML = docType + '\n' + html;
  var fullHTMLTmp = fullHTML.split(DELIMITER);
  if (fullHTMLTmp.length <= 1) {
    var msg = 'Error occurred while extracting HTML from page.';
    handleError(msg);
    return;
  }
  fullHTML = fullHTMLTmp[0] + DELIMITER + '\n</body>\n</html>';
  fullHTML = fullHTML.replace(new RegExp('const IS_STATIC = false', ''), 'const IS_STATIC = true');
  STATIC_ARRAY.forEach ( function ( entry ) {
    if (entry == 'CITATIONS') {
      var citationData = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf(entry)];
      if (!!citationData && !!citationData.response) {
        // getting rid of BibTex entries, since they cause JSON.parse errors
        citationData.response.forEach ( function ( citationEntry ) {
          if (!!citationEntry.bib) delete citationEntry.bib;
        });
      }
      fullHTML = fullHTML.replace(new RegExp('var STATIC_' + entry, ''), 'var STATIC_' + entry + ' = `' + unescape(JSON.stringify(citationData)).replace(/\\\"/g, "'") + '`');
    } else fullHTML = fullHTML.replace(new RegExp('var STATIC_' + entry, ''), 'var STATIC_' + entry + ' = `' + unescape(JSON.stringify(DYNAMIC_ARRAY[STATIC_ARRAY.indexOf(entry)])) + '`');
  });

  var scripts = document.scripts;
  var links = document.getElementsByTagName('link');
  var promises = [];

  // fetch JS files
  for ( var i = 0; i < scripts.length; i++ ) {
    var srcURL = $(scripts[i]).attr('src');
    var webURL = $(scripts[i]).attr('data-web');
    if (!!webURL) {
      fullHTML = fullHTML.replace(new RegExp('<script src="' + srcURL + '" data-web="' + webURL + '">', 'g'), '<script src="' + webURL + '" data-web="' + srcURL + '">');
    } else if (!!srcURL && srcURL != '' && srcURL.split('.js').length > 1) {
      var fetchPromise = new Promise ( function ( resolve, reject ) {
        var url = srcURL;
        $.ajax({
          type: "GET",  
          url: url,
          success: function(data){
            if (!data) reject('Cannot fetch file "' + url + '" from server.');
            else resolve({url: url, text: data, type: 'script'}); 
          },
          error: function(XMLHttpRequest, textStatus, errorThrown) { 
            reject(errorThrown);
          }       
        });
      });
      promises.push(fetchPromise);
    }
  }
  // fetch CSS files
  for ( var i = 0; i < links.length; i++ ) {
    var srcURL = $(links[i]).attr('href');
    var webURL = $(links[i]).attr('data-web');
    console.log(links[i])
    if (!!webURL) {
      fullHTML = fullHTML.replace(new RegExp('<link href="' + srcURL + '" type="text/css" rel="stylesheet" data-web="' + webURL + '">', 'g'), '<link href="' + webURL + '" type="text/css" rel="stylesheet" data-web="' + srcURL + '">');
    } else if (!!srcURL && srcURL != '' && srcURL.split('.css').length > 1) { //
      var fetchPromise = new Promise ( function ( resolve, reject ) {
        var url = srcURL;
        $.ajax({
          type: "GET",  
          url: url,
          success: function(data){
            if (!data) reject('Cannot fetch file "' + url + '" from server.');
            else resolve({url: url, text: data, type: 'link'}); 
          },
          error: function(XMLHttpRequest, textStatus, errorThrown) { 
            reject(errorThrown);
          }       
        });
      });
      promises.push(fetchPromise);
    }
  }
  if (promises.length == 0) saveAs(new Blob([fullHTML], {type: "text/html;charset=utf-8"}), fileName);
  else {
    Promise.all(promises)
      .then ( function ( results ) {
        results.forEach ( function ( result ) {
          console.log(new RegExp('<script src="' + result.url + '">', 'g'))
          if (result.type == 'script') fullHTML = fullHTML.replace(new RegExp('<script src="' + result.url + '">', 'g'), '<script>' + result.text);
          else if (result.type == 'link') fullHTML = fullHTML.replace(new RegExp('<link href="' + result.url + '" type="text/css" rel="stylesheet">', 'g'), '<style>' + result.text + '</style>');
        });
        saveAs(new Blob([fullHTML], {type: "text/html;charset=utf-8"}), fileName);
      }).catch ( function ( err ) {
        console.log("Error occurred while fetching script from server: ", err);
        handleError(err);
      });
  }
}