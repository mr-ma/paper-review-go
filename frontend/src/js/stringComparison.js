  // exact matching using regular expressions
  function occurrencesRegex ( string, subStrings ) {
    var counter = 0;
    var occurrences = [];
    subStrings.forEach ( function ( subString ) {
      var value = (string.match(new RegExp(subString, 'gi')) || []).length;
      if (value > 0) {
        counter += value;
        occurrences.push({name: subString, value: value});
      }
    });
    return {value: counter, occurrences: occurrences};
  }

  // string comparison based on "Dice's Coefficient"
  function occurrencesDice ( string, subStrings ) {
    var words = (!!string && string != '') ? string.replace(/([^\w\d\s]+)/g, '').replace(/\s(\s+)/g, ' ').split(/\s/) : [];
    var counter = 0;
    var occurrences = [];
    words.forEach ( function ( entry ) {
      if (entry.length > 2) {
        subStrings.forEach ( function ( subString ) {
          var value = compareTwoStrings(entry, subString);
          if (value >= minSimilarity) {
            counter++;
            var index = -1;
            for ( var i = 0; i < occurrences.length; i++ ) {
              if (occurrences[i].name == entry) {
                index = i;
                break;
              }
            }
            if (index < 0) occurrences.push({name: entry, value: 1});
            else occurrences[index].value++;
          }
        });
      }
    });
    return {value: counter, occurrences: occurrences};
  }

  // string comparison based on the "Levenshtein Distance"
  function occurrencesLevenshtein ( string, subStrings ) {
    var levenshtein = window.Levenshtein;
    var words = (!!string && string != '') ? string.replace(/([^\w\d\s]+)/g, '').replace(/\s(\s+)/g, ' ').split(/\s/) : [];
    var counter = 0;
    var occurrences = [];
    words.forEach ( function ( entry ) {
      if (entry.length > 2) {
        subStrings.forEach ( function ( subString ) {
          var value = levenshtein.get(entry, subString);
          if (value <= minSimilarity) {
            counter++;
            var index = -1;
            for ( var i = 0; i < occurrences.length; i++ ) {
              if (occurrences[i].name == entry) {
                index = i;
                break;
              }
            }
            if (index < 0) occurrences.push({name: entry, value: 1});
            else occurrences[index].value++;
          }
        });
      }
    });
    return {value: counter, occurrences: occurrences};
  }

  // string comparison used in source code editors (e.g. Sublime Text)
  function occurrencesFuzzysort ( string, subStrings ) {
    var words = (!!string && string != '') ? string.replace(/([^\w\d\s]+)/g, '').replace(/\s(\s+)/g, ' ').split(/\s/) : [];
    var counter = 0;
    var occurrences = [];
    words.forEach ( function ( entry ) {
      if (entry.length > 2) {
        var entryPrepared = fuzzysort.prepare(entry)
        subStrings.forEach ( function ( subString ) {
          var value = fuzzysort.single(subString, entryPrepared);
          if (!!value && value.score >= FUZZY_SORT_MIN_SIMILARITY) {
            counter++;
            var index = -1;
            for ( var i = 0; i < occurrences.length; i++ ) {
              if (occurrences[i].name == entry) {
                index = i;
                break;
              }
            }
            if (index < 0) occurrences.push({name: entry, value: 1});
            else occurrences[index].value++;
          }
        });
      }
    });
    return {value: counter, occurrences: occurrences};
  }