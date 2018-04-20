// functions used to import PDF files into the browser and parse their text

	function getFileType ( url ) {
	  var result = url;
	  var index = result.indexOf('.');
	  while (index > 0) {
	      if (index + 1 >= result.length)
	        return result;
	      result = result.substr(index + 1);
	      index = result.indexOf('.');
	  }
	  return result;
	}

	function readTextHelper ( pdf, maxPages ){
	     var countPromises = [];
	     for (var j = 1; j <= maxPages; j++) {
	        var page = pdf.getPage(j);
	        var txt = "";
	        countPromises.push(page.then(function(page) {
	            var textContent = page.getTextContent();
	            return textContent.then(function(text){
	                return text.items.map(function (s) { return s.str; }).join(' ');
	            });
	        }));
	     }
	     return Promise.all(countPromises).then(function (texts) {
	       return texts.join(' ');
	     });
	}

	function readText ( pdfUrl ){
		var pdf = PDFJS.getDocument(pdfUrl);
		return  pdf.then(function(pdf) {
		     		var maxPages = pdf.pdfInfo.numPages;
		     		return pdf.getMetadata()
				     	.then ( function ( metaDataResult ) {
				     		if (!!metaDataResult && !!metaDataResult.info) {
				     			return {promise: readTextHelper(pdf, maxPages), metaData: metaDataResult.info};
				     		} else return {promise: readTextHelper(pdf, maxPages), metaData: {}};
					     }).catch ( function ( err ) {
					     	console.log('PDF meta data parse error: ', err);
					     	return {promise: readTextHelper(pdf, maxPages), metaData: {}};
					     });
				}).catch ( function ( err ) {
					console.log('PDF parse error: ', err);
					return {metaData: {}};
				});
	}

	function importPDF ( file, callback ) {
	    var fileReader = new FileReader();
	    fileReader.onload = function() {
	        var typedarray = new Uint8Array(this.result);
			readText(typedarray).then(function ( resultObj ) {
				if (!resultObj.promise) {
					callback({success: false, msg: 'Error parsing text from PDF file.'});
					return;
				}
				var metaData = resultObj.metaData;
				resultObj.promise.then ( function ( text ) {
					if (!!text) callback({success: true, msg: {text: text, metaData: metaData}});
					else callback({success: false, msg: 'Error parsing text from PDF file.'});
			  	}).catch ( function ( err ) {
			  		callback({success: false, msg: err});
			  	});
			}).catch ( function ( err ) {
				callback({success: false, msg: err});
			});
	    };
	    fileReader.readAsArrayBuffer(file);
	}