  function showDimension ( dimension ) {
    var dimensions = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('DIMENSIONS')];
    var relationTypes = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('RELATIONTYPES')];
    var citationCounts = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('CITATIONCOUNTS')];
    if (dimension == 'Interdimensional view') {
      var attributes = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('MAJORATTRIBUTES')];
      var relations = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('INTERDIMENSIONALRELATIONS')];
    } else {
      var index = -1;
      for ( var i = 0; i < dimensions.length; i++ ) {
        if (dimensions[i].text == dimension) {
          index = i;
          break;
        }
      }
      if (index < 0) return;
      var attributes = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTESPERDIMENSION')][index];
      var relations = DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTERELATIONS')][index];
      dimensions = [{text: dimension}];
    }
    createJSON(!!dimensions ? dimensions : [], !!attributes ? attributes : [], !!relations ? relations : [], !!relationTypes ? relationTypes : [], !!citationCounts ? citationCounts : []);
  }

    function loadTaxonomyData ( taxonomyID, attributeURL, citationCountsURL ) {
      if (!IS_STATIC) {
        var request = {'taxonomy_id': taxonomyID};
        var promises = [];
        var getDimensionsPromise = new Promise ( function ( resolve, reject ) {
          $.get('dimension', function ( dimensionResults ) {
            if (!dimensionResults || !dimensionResults.response) {
              reject('Cannot get dimensions from DB.');
              return;
            }
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
                      if (!attributesPerDimension || !attributesPerDimension.response) {
                        reject('Cannot get attributes of dimension "' + dimension + '" from DB.')
                        return;
                      }
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
                        if (!attributeRelations || !attributeRelations.response) {
                          reject('Cannot get attribute relations of dimension "' + dimension + '" from DB.')
                          return;
                        }
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
          });
        });
        promises.push(getDimensionsPromise);
        var getMajorAttributesPromise = new Promise ( function ( resolve, reject ) {
          $.ajax
            ({
                type: "POST",
                url: attributeURL,
                dataType: 'json',
                contentType:'application/json',
                async: true,
                data: JSON.stringify(request),
                success: function ( majorAttributes ) {
                  if (!majorAttributes || !majorAttributes.response) {
                    reject('Cannot get major attributes from DB.')
                    return;
                  }
                  resolve({name: 'MAJORATTRIBUTES', value: majorAttributes.response});
                }
            });
        });
        promises.push(getMajorAttributesPromise);
        var getInterdimensionalRelationsPromise = new Promise ( function ( resolve, reject ) {
          $.ajax
          ({
              type: "POST",
              url: 'interdimensionalRelations',
              dataType: 'json',
              contentType:'application/json',
              async: true,
              data: JSON.stringify(request),
              success: function ( interdimensionalRelations ) {
                if (!interdimensionalRelations || !interdimensionalRelations.response) {
                  reject('Cannot get interdimensional relations from DB.');
                  return;
                }
                resolve({name: 'INTERDIMENSIONALRELATIONS', value: interdimensionalRelations.response});
              }
          });
        });
        promises.push(getInterdimensionalRelationsPromise);
        var getRelationTypesPromise = new Promise ( function ( resolve, reject ) {
          $.get('relationTypes', function ( relationTypes ) {
            if (!relationTypes || !relationTypes.response) {
              reject('Cannot get relation types from DB.');
              return;
            }
            resolve({name: 'RELATIONTYPES', value: relationTypes.response});
          });
        });
        promises.push(getRelationTypesPromise);
        var getCitationCountsPromise = new Promise ( function ( resolve, reject ) {
          $.get(citationCountsURL, function ( citationCounts ) {
            if (!citationCounts || !citationCounts.response) {
              reject('Cannot get citation counts from DB.');
              return;
            }
            console.log("cit: ", citationCounts.response)
            resolve({name: 'CITATIONCOUNTS', value: citationCounts.response});
          });
        });
        promises.push(getCitationCountsPromise);
        Promise.all(promises)
          .then ( function ( results ) {
            DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('DIMENSIONS')] = dimensions;
            dimensions.forEach ( function ( dimension ) {
              DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTESPERDIMENSION')] = [];
              DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTERELATIONS')] = [];
            });
            results.forEach ( function ( result ) {
              if (result.name == 'DIMENSIONDATA') {
                console.log('dim data: ', result.value)
                result.value.forEach ( function ( dimensionData ) {
                  if (dimensionData.name == 'ATTRIBUTESPERDIMENSION') DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTESPERDIMENSION')][dimensionData.value.index] = dimensionData.value.value;
                  else DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTERELATIONS')][dimensionData.value.index] = dimensionData.value.value;
                });
              } else DYNAMIC_ARRAY[STATIC_ARRAY.indexOf(result.name)] = result.value;
            });
            displayedDimension = 'Interdimensional view';
            showDimension(displayedDimension);
          }).catch ( function ( err ) {
            console.log('Error loading data from DB: ', err);
            handleErrorHelper(err);
          });
      } else {
        try {
          var dimensions = JSON.parse(STATIC_DIMENSIONS);
          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('DIMENSIONS')] = dimensions;
          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('MAJORATTRIBUTES')] = JSON.parse(STATIC_MAJORATTRIBUTES);
          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('INTERDIMENSIONALRELATIONS')] = JSON.parse(STATIC_INTERDIMENSIONALRELATIONS);
          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('RELATIONTYPES')] = JSON.parse(STATIC_RELATIONTYPES);
          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('CITATIONCOUNTS')] = JSON.parse(STATIC_CITATIONCOUNTS);

          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTESPERDIMENSION')] = [];
          DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTERELATIONS')] = [];

          var attributesPerDimension = JSON.parse(STATIC_ATTRIBUTESPERDIMENSION);
          var attributeRelations = JSON.parse(STATIC_ATTRIBUTERELATIONS);

          for ( var i = 0; i < dimensions.length; i++ ) {
            if (attributesPerDimension.length > i) DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTESPERDIMENSION')][i] = attributesPerDimension[i];
            if (attributeRelations.length > i) DYNAMIC_ARRAY[STATIC_ARRAY.indexOf('ATTRIBUTERELATIONS')][i] = attributeRelations[i];
          }
          displayedDimension = 'Interdimensional view';
          showDimension(displayedDimension);
        } catch ( err ) {
          handleErrorHelper('Error parsing STATIC data.');
        }
      }
    }