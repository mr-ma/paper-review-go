  function scopusLookup ( text, callback ) {
    const API_KEY_LOCAL = 'c4ddcf695994bcc8b4a1c6d4a593a974';
    const API_KEY = '90c6191fcbabf8d9969eb9c750cc898e';
    $.ajax
      ({
        type: "GET",
        // TODO replace other letters like (), : etc
        url: 'https://api.elsevier.com/content/search/scopus?query=TITLE(' + escape(text) + ')&content=all',
        //dataType: 'json',
        contentType:'application/json',
        async: true,
        headers: {'Accept':'application/json', 'X-ELS-APIKey': API_KEY_LOCAL},
        //data: JSON.stringify({'query': 'KEY(test)', 'field': 'All', 'content': 'core'}),
        success: function ( result ) {
          if (!result || !result['search-results'] || !result['search-results'].entry) return;
          var entry = result['search-results'].entry;
          for ( var i = 0; i < entry.length; i++ ) {
            if (!!entry[i]['dc:title']) {
              console.log('obj: ', entry[i])
              var referenceCount = entry[i]['citedby-count'] - 0;
              if (isNaN(referenceCount)) referenceCount = 0;
              callback({success: true, msg: {citation: entry[i]['dc:title'], referenceCount: referenceCount}});
              break;
            }
          }
          callback({success: false, msg: 'No entries found for: ', text});
        },
        error: function ( err ) {
          console.log('Scopus ajax error: ', err);
          callback({success: false, msg: 'Ajax error'});
        }
      });
    }