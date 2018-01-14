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

function uploadFiles ( files, callback ) {
	console.log("files: ", files);
	var formData = new FormData();
	for ( var i = 0; i < files.length; i++ ) {
		formData.append('uploads[file]', files[i], files[i].name);
	}
	$.ajax({
	  url: '/upload',
	  type: 'POST',
	  data: formData,
	  processData: false,
	  contentType: false,
	  success: function ( result ) {
        if (!result || !result.response || !result.response.success) {
          handleErrorHelper('Cannot upload files to server.');
          return;
        }
        console.log('File upload successful.');
        if (!!callback) callback();
	  }
	});
}

	function gettext(pdfUrl){
	var pdf = PDFJS.getDocument(pdfUrl);
	return pdf.then(function(pdf) { // get all pages text
	     var maxPages = pdf.pdfInfo.numPages;
	     console.log("page amount: " + maxPages);
	     var countPromises = []; // collecting all page promises
	     for (var j = 1; j <= maxPages; j++) {
	        var page = pdf.getPage(j);
	        var txt = "";
	        countPromises.push(page.then(function(page) { // add page promise
	            var textContent = page.getTextContent();
	            return textContent.then(function(text){ // return content promise
	                return text.items.map(function (s) { return s.str; }).join(' '); // value page text 

	            });
	        }));
	     }
	     return Promise.all(countPromises).then(function (texts) {
	       return texts.join(' ');
	     });
	});
}

function importPDF ( file, callback ) {
    var fileReader = new FileReader();
    fileReader.onload = function() {
        var typedarray = new Uint8Array(this.result);
		gettext(typedarray).then(function (text) {
		  var tmp = text.replace(/([^\w\d\s]+)/g, '').replace(/\s(\s+)/g, ' ').split(/\s/);
		  console.log("words: " + tmp.length);
		  console.log("last 2: " + tmp[tmp.length-2] + ', ' + tmp[tmp.length-1]);
		  if (!!callback) callback({success: true, msg: text});
		}, function (err) {
			callback({success: false, msg: err});
		});
    };
    fileReader.readAsArrayBuffer(file);
}