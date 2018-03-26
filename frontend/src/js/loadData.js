// functions used to load data from the database

  function showDimension ( dimension, useCY ) {
    console.log("showing dimension: " + dimension)
    var dimensions = DYNAMIC_ARRAY[STATIC_INDEX_DIMENSIONS];
    var relationTypes = DYNAMIC_ARRAY[STATIC_INDEX_RELATIONTYPES];
    var citationCounts = DYNAMIC_ARRAY[STATIC_INDEX_CITATIONCOUNTS];
    if (dimension == DEFAULT_DIMENSION_NAME) {
      var interDimensional = true;
      var attributes = DYNAMIC_ARRAY[STATIC_INDEX_MAJORATTRIBUTES];
      var relations = DYNAMIC_ARRAY[STATIC_INDEX_INTERDIMENSIONALRELATIONS];
    } else {
      var interDimensional = false;
      var index = -1;
      for ( var i = 0; i < dimensions.length; i++ ) {
        if (dimensions[i].text == dimension) {
          index = i;
          break;
        }
      }
      if (index < 0) return;
      var attributes = DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTESPERDIMENSION][index];
      var relations = DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTERELATIONS][index];
      dimensions = [{text: dimension}];
    }
    if (!!useCY) {
      DYNAMIC_ARRAY[STATIC_INDEX_CY] = JSON.parse(STATIC_CY);
      createJSON(true, !!dimensions ? dimensions : [], interDimensional, !!attributes ? attributes : [], !!relations ? relations : [], !!relationTypes ? relationTypes : [], !!citationCounts ? citationCounts : []);
    } else createJSON(false, !!dimensions ? dimensions : [], interDimensional, !!attributes ? attributes : [], !!relations ? relations : [], !!relationTypes ? relationTypes : [], !!citationCounts ? citationCounts : []);
  }

    function loadTaxonomyData ( taxonomyID, attributeURL, citationCountsURL, useCY ) {
      var dimensionHashText = unescape(window.location.hash.split('_').pop());
      if (!IS_STATIC) {
        var request = {'taxonomy_id': taxonomyID};
        var promises = [];
        var getDimensionsPromise = new Promise ( function ( resolve, reject ) {
          var request = {'taxonomy_id': taxonomyID};
          $.ajax
            ({
              type: "POST",
              url: 'dimension',
              dataType: 'json',
              contentType:'application/json',
              async: true,
              data: JSON.stringify(request),
              success: function ( dimensionResults ) {
                if (!dimensionResults) {
                  reject('Cannot get dimensions from DB.');
                  return;
                }
                if (!dimensionResults.response) dimensionResults.response = [];
                dimensions = dimensionResults.response;
                if (dimensions.length == 0) {
                  resolve({name: 'DIMENSIONDATA', value: []});
                  return;
                }
                var dimensionPromises = [];
                for ( var i = 0; i < dimensions.length; i++ ) {
                  var getAttributesPerDimensionPromise = new Promise ( function ( resolve, reject ) {
                    var index = i;
                    var dimension = dimensions[i].text;
                    $.ajax
                      ({
                        type: "POST",
                        url: 'attributesPerDimension',
                        dataType: 'json',
                        contentType:'application/json',
                        async: true,
                        data: JSON.stringify({'taxonomy_id': taxonomyID, dimension: dimension}),
                        success: function ( attributesPerDimension ) {
                          if (!attributesPerDimension) {
                            reject('Cannot get attributes of dimension "' + dimension + '" from DB.')
                            return;
                          }
                          if (!attributesPerDimension.response) attributesPerDimension.response = [];
                          resolve({name: 'ATTRIBUTESPERDIMENSION', value: {index: index, value: attributesPerDimension.response}});
                        }
                      });
                  });
                  dimensionPromises.push(getAttributesPerDimensionPromise);
                  var getRelationsPerDimensionPromise = new Promise ( function ( resolve, reject ) {
                    var index = i;
                    var dimension = dimensions[i].text;
                    $.ajax
                      ({
                          type: "POST",
                          url: 'attributeRelations',
                          dataType: 'json',
                          contentType:'application/json',
                          async: true,
                          data: JSON.stringify({'taxonomy_id': taxonomyID, dimension: dimension}),
                          success: function ( attributeRelations ) {
                            if (!attributeRelations) {
                              reject('Cannot get attribute relations of dimension "' + dimension + '" from DB.')
                              return;
                            }
                            if (!attributeRelations.response) attributeRelations.response = [];
                            resolve({name: 'ATTRIBUTERELATIONS', value: {index: index, value: attributeRelations.response}});
                          }
                      });
                  });
                  dimensionPromises.push(getRelationsPerDimensionPromise);
                }
                Promise.all(dimensionPromises)
                  .then ( function ( results ) {
                    resolve({name: 'DIMENSIONDATA', value: results});
                  }).catch ( function ( err ) {
                    reject(err);
                  });
              }
          });
        });
        promises.push(getDimensionsPromise);
        var getMajorAttributesPromise = new Promise ( function ( resolve, reject ) {
          var request = {'taxonomy_id': taxonomyID};
          $.ajax
            ({
                type: "POST",
                url: attributeURL,
                dataType: 'json',
                contentType:'application/json',
                async: true,
                data: JSON.stringify(request),
                success: function ( majorAttributes ) {
                  if (!majorAttributes) {
                    reject('Cannot get major attributes from DB.')
                    return;
                  }
                  if (!majorAttributes.response) majorAttributes.response = [];
                  resolve({name: 'MAJORATTRIBUTES', value: majorAttributes.response});
                }
            });
        });
        promises.push(getMajorAttributesPromise);
        var getInterdimensionalRelationsPromise = new Promise ( function ( resolve, reject ) {
          var request = {'taxonomy_id': taxonomyID};
          $.ajax
          ({
              type: "POST",
              url: 'interdimensionalRelations',
              dataType: 'json',
              contentType:'application/json',
              async: true,
              data: JSON.stringify(request),
              success: function ( interdimensionalRelations ) {
                if (!interdimensionalRelations) {
                  reject('Cannot get interdimensional relations from DB.');
                  return;
                }
                if (!interdimensionalRelations.response) interdimensionalRelations.response = [];
                resolve({name: 'INTERDIMENSIONALRELATIONS', value: interdimensionalRelations.response});
              }
          });
        });
        promises.push(getInterdimensionalRelationsPromise);
        var getRelationTypesPromise = new Promise ( function ( resolve, reject ) {
          $.get('relationTypes', function ( relationTypes ) {
            if (!relationTypes) {
              reject('Cannot get relation types from DB.');
              return;
            }
            if (!relationTypes.response) relationTypes.response = [];
            resolve({name: 'RELATIONTYPES', value: relationTypes.response});
          });
        });
        promises.push(getRelationTypesPromise);
        var getCitationCountsPromise = new Promise ( function ( resolve, reject ) {
          var request = {'taxonomy_id': taxonomyID};
          $.ajax
            ({
                type: "POST",
                url: citationCountsURL,
                dataType: 'json',
                contentType:'application/json',
                async: true,
                data: JSON.stringify(request),
                success: function ( citationCounts ) {
                  if (!citationCounts) {
                    reject('Cannot get citation counts from DB.')
                    return;
                  }
                  if (!citationCounts.response) citationCounts.response = [];
                  resolve({name: 'CITATIONCOUNTS', value: citationCounts.response});
                }
            });
        });
        promises.push(getCitationCountsPromise);
        Promise.all(promises)
          .then ( function ( results ) {
            DYNAMIC_ARRAY[STATIC_INDEX_DIMENSIONS] = dimensions;
            var foundDimension = false;
            dimensions.forEach ( function ( dimension ) {
              DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTESPERDIMENSION] = [];
              DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTERELATIONS] = [];
              if (dimension.text == dimensionHashText) foundDimension = true;
            });
            results.forEach ( function ( result ) {
              if (result.name == 'DIMENSIONDATA') {
                console.log('dim data: ', result.value)
                result.value.forEach ( function ( dimensionData ) {
                  if (dimensionData.name == 'ATTRIBUTESPERDIMENSION') DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTESPERDIMENSION][dimensionData.value.index] = dimensionData.value.value;
                  else DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTERELATIONS][dimensionData.value.index] = dimensionData.value.value;
                });
              } else DYNAMIC_ARRAY[STATIC_ARRAY.indexOf(result.name)] = result.value;
            });
            if (foundDimension) displayedDimension = dimensionHashText;
            else displayedDimension = DEFAULT_DIMENSION_NAME;
            showDimension(displayedDimension, useCY);
          }).catch ( function ( err ) {
            console.log('Error loading data from DB: ', err);
            handleError(err);
          });
      } else {
        try {
          var dimensions = JSON.parse(STATIC_DIMENSIONS);
          DYNAMIC_ARRAY[STATIC_INDEX_DIMENSIONS] = dimensions;
          DYNAMIC_ARRAY[STATIC_INDEX_MAJORATTRIBUTES] = JSON.parse(STATIC_MAJORATTRIBUTES);
          DYNAMIC_ARRAY[STATIC_INDEX_INTERDIMENSIONALRELATIONS] = JSON.parse(STATIC_INTERDIMENSIONALRELATIONS);
          DYNAMIC_ARRAY[STATIC_INDEX_RELATIONTYPES] = JSON.parse(STATIC_RELATIONTYPES);
          DYNAMIC_ARRAY[STATIC_INDEX_CITATIONCOUNTS] = JSON.parse(STATIC_CITATIONCOUNTS);

          DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTESPERDIMENSION] = [];
          DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTERELATIONS] = [];

          var attributesPerDimension = JSON.parse(STATIC_ATTRIBUTESPERDIMENSION);
          var attributeRelations = JSON.parse(STATIC_ATTRIBUTERELATIONS);

          var foundDimension = false;
          for ( var i = 0; i < dimensions.length; i++ ) {
            if (attributesPerDimension.length > i) DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTESPERDIMENSION][i] = attributesPerDimension[i];
            if (attributeRelations.length > i) DYNAMIC_ARRAY[STATIC_INDEX_ATTRIBUTERELATIONS][i] = attributeRelations[i];
            if (dimensions[i].text == dimensionHashText) foundDimension = true;
          }
          if (foundDimension) displayedDimension = dimensionHashText;
          else displayedDimension = DEFAULT_DIMENSION_NAME;
          showDimension(displayedDimension, useCY);
        } catch ( err ) {
          console.log('Error parsing STATIC data: ', err);
          handleError('Error parsing STATIC data.');
        }
      }
    }