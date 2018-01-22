var admin = false;

function validateEmail(email) {
  var re = /^[\w-']+(\.[\w-']+)*@([a-z0-9-]+(\.[a-z0-9-]+)*?\.[a-z]{2,6}|(\d{1,3}\.){3}\d{1,3})(:\d{4})?$/i;
  return re.test(email);}

  function addTaxonomy () {
    if (!admin) return;
    //if (!!cy) cy.elements().trigger('qtiphide');
    $('#addTaxonomyModalText').val('');
    $('#addTaxonomyModalText1').val('');
    $('#addTaxonomyModalForm').on('submit', function ( evt ) {
      evt.preventDefault();
      var taxonomy = $('#addTaxonomyModalText1').val();
      var dimension = $('#addTaxonomyModalText2').val();
      $('#addTaxonomyModalButton').prop('disabled', true);
      $.ajax
        ({
            type: "POST",
            url: 'addTaxonomy',
            dataType: 'json',
            contentType:'application/json',
            async: true,
            data: JSON.stringify({taxonomy: taxonomy, dimension: dimension}),
            success: function ( result ) {
              $('#addTaxonomyModalButton').prop('disabled', false);
              if (!result || !result.success) {
                var msg = 'Cannot add taxonomy: ' + taxonomy;
                if (!!handleErrorHelper) handleErrorHelper(msg);
                else handleError(msg);
                return;
              }
              $('#addTaxonomy-modal').modal('hide');
              var url = window.location.origin;
              window.location.href = url + '/taxonomyRelations#' + taxonomy;
            }
        });
    });
    $('#addTaxonomy-modal').modal('show');
  }

  function getUser ( resolve, reject ){
    console.log('login')
    $.ajax
      ({
        type: "POST",
        url: 'login',
        dataType: 'json',
        contentType:'application/json',
        async: true,
        xhrFields: {
           withCredentials: true
        },
        data: JSON.stringify({email: '', password: ''}),
        success: function ( result ) {
          console.log(result)
          if (!result || !result.success || !result.user) {
            reject({});
            return;
          }
          resolve(result.user);
        }
    });
  }

  function loadBanner ( type ) {
    switch(type) {
      case 'ADMIN': $('.ribbonContainer').append('<div class="ribbon red"><span>ADMIN</span></div>'); break;
      case 'EDIT': $('.ribbonContainer').append('<div class="ribbon red"><span>EDIT</span></div>'); break;
      case 'VIEW': $('.ribbonContainer').append('<div class="ribbon blue"><span>VIEW</span></div>'); break;
      case 'STATIC': $('.ribbonContainer').append('<div class="ribbon green"><span>STATIC</span></div>'); break;
      default:
    }
    $('.ribbon').css('right','-15px')
  }

  function loadNavbar ( user, resolve, reject ) {
    if (!!user.admin && (user.admin - 0) == 1) var navbarPath = '/navbarAdmin.html';
    else var navbarPath = '/navbar.html';
    $.get(navbarPath, function (data ) {
      $('.navbar').replaceWith(data);
      loadNavbarFields(user);
      $('a').on('click', function ( e ) {
        e.preventDefault();
        var hash = window.location.hash;
        window.location = window.location.origin + $(this).attr('href') + hash;
      });
      if (!!user.admin) {
        $('#taxonomyDropdown').html('');
        $.get('taxonomy', function ( taxonomies ) {
          if (!!taxonomies && !!taxonomies.response) {
            taxonomies.response.forEach ( function ( taxonomy ) {
              $('#taxonomyDropdown').append('<li><a href="#' + taxonomy.text + '" name="' + taxonomy.id + '">' + taxonomy.text + '</a></li>')
            });
          }
          $('#taxonomyDropdown').append('<li><input type="button" class="btn btn-primary" id="addTaxonomy" value="Add Taxonomy" style="margin-left:20px;margin-top:10px;" onclick="addTaxonomy()"></li>');
          resolve();
        });
      } else resolve();
    });
  }

  function loadNavbarFields ( user ) {
    if (!!user && !!user.email) {
      $('#userName').val(user.email);
      $('#userName').prop('title', user.email);
      $('#userField').show();
    }
  }

  function loadModals ( user, resolve, reject ) {
    $.get('modals.html', function (data ) {
      $('.modals').replaceWith(data);
      $('.close-btn').unbind().click( function () {
        $(this).parent().parent().parent().modal('hide');
      });
      $('#loginForm').on('submit', function (e) {
        e.preventDefault();
        var user = {};
        user.email = $(this).find('input[name=email]').val();
        user.password = $(this).find('input[name=password]').val();
        if (!validateEmail(user.email)) {
          $(this).find('input[name=email]').css('color', 'red');
          return;
        } else {
          $(this).find('input[name=email]').css('color', 'black');
        }
        $('#loginSubmitButton').prop('disabled', true);
        $.ajax
          ({
            type: "POST",
            url: 'login',
            dataType: 'json',
            contentType:'application/json',
            async: true,
            xhrFields: {
               withCredentials: true
            },
            data: JSON.stringify({email: user.email, password: user.password}),
            success: function ( result, status, xhr ) {
              $('#loginSubmitButton').prop('disabled', false);
              if (!result || !result.success || !result.user) {
                var msg = 'A user with this name or password does not exist.';
                if (!!handleErrorHelper) handleErrorHelper(msg);
                else handleError(msg);
                return;
              }
              //loginUser(result.user, false);
              $('#login-modal').modal('hide');
              window.location.reload(false);
            }
        });
      });
      $('#login-modal').on('show.bs.modal', function () {
          $(this).find('form').trigger('reset');
      });
      $('#signupForm').on('submit', function (e) {
        e.preventDefault();
        var user = {};
        user.email = $(this).find('input[name=email]').val();
        user.password = $(this).find('input[name=password]').val();
        user.admin = 0;
        if (!validateEmail(user.email)) {
          $(this).find('input[name=email]').css('color', 'red');
          return;
        } else {
          $(this).find('input[name=email]').css('color', 'black');
        }
        $('#signupSubmitButton').prop('disabled', true);
        $.ajax
          ({
            type: "POST",
            url: 'saveUser',
            dataType: 'json',
            contentType:'application/json',
            async: true,
            xhrFields: {
               withCredentials: true
            },
            data: JSON.stringify({email: user.email, password: user.password}),
            success: function ( result, status, xhr ) {
              console.log("result: ", result)
              $('#signupSubmitButton').prop('disabled', false);
              if (!result || !result.success) {
                var msg = 'A user with this email already exists.';
                if (!!handleErrorHelper) handleErrorHelper(msg);
                else handleError(msg);
                return;
              }
              //loginUser(user, false);
              $('#signup-modal').modal('hide');
              window.location.reload(false);
            }
        });
      });
      $('#signup-modal').on('show.bs.modal', function () {
          $(this).find('form').trigger('reset');
      });
      resolve();
    });
  }

  function loginUser ( user, withModals, resolve, reject ) {
    if (!!user && user.admin == 1) admin = true;
    var promises = [];
    var loadNavbarPromise = new Promise ( function ( resolve, reject ) {
      loadNavbar(user, resolve, reject);
    });
    promises.push(loadNavbarPromise);
    if (!!withModals) {
      var loadModalsPromise = new Promise ( function ( resolve, reject ) {
        loadModals(user, resolve, reject);
      });
      promises.push(loadModalsPromise);
    }
    var getTaxonomyIDPromise = new Promise ( function ( resolve, reject ) {
      var url = window.location.href.split('#');
      var hash = window.location.hash.split('#').pop();
      if (!!hash && hash != '') {
        $.ajax
          ({
            type: "POST",
            url: 'getTaxonomyID',
            dataType: 'json',
            contentType:'application/json',
            async: true,
            data: JSON.stringify({text: unescape(hash)}),
            success: function ( taxonomy ) {
              if (!taxonomy || !taxonomy.response || taxonomy.response.length == 0) {
                console.log('Cannot get taxonomy ID From DB.');
                TAXONOMY_ID = DEFAULT_TAXONOMY_ID;
                window.location.href = url.shift();
                resolve();
                return;
              } else {
                var id = taxonomy.response[0].id - 0;
                console.log("taxonomy id: ", id)
                if (!isNaN(id)) TAXONOMY_ID = id;
                else TAXONOMY_ID = DEFAULT_TAXONOMY_ID;
              }
              resolve();
            }
        });
      } else {
        TAXONOMY_ID = DEFAULT_TAXONOMY_ID;
        resolve();
      }
    });
    promises.push(getTaxonomyIDPromise);
    Promise.all(promises)
      .then ( function ( results ) {
        if (user.email != '') {
          $('#loginField').hide();
          $('#logoutField').show();
        } else {
          $('#loginField').show();
          $('#logoutField').hide();
        }
        if (user.admin == 1) loadBanner('ADMIN');
        else {
          var taxonomies = (!!user.taxonomies && user.taxonomies != '') ? user.taxonomies.split(',') : [];
          if (taxonomies.indexOf('' + TAXONOMY_ID) > -1) loadBanner('EDIT');
          else loadBanner('VIEW');
        }
        $('#logoutField').unbind().on('click', function () {
          $.ajax
            ({
              type: "POST",
              url: 'logout',
              dataType: 'json',
              contentType:'application/json',
              async: true,
              xhrFields: {
                 withCredentials: true
              },
              data: JSON.stringify({}),
              success: function ( result, status, xhr ) {
                logoutUser();
                window.location.reload(false);
              }
          });
        });
        if (!!resolve) resolve(user);
      }).catch ( function ( err ) {
        console.log('Init user error: ', err);
        var msg = 'Init user error.';
        if (!!handleErrorHelper) handleErrorHelper(msg);
        else handleError(msg);
        if (!!resolve) resolve(user);
      });
  }

  function logoutUser() {
    loginUser({email: '', admin: 0}, false);
    admin = false;
  }

  function initUserManagement( resolve, reject ) {
    var loginPromise = new Promise( function ( resolve, reject ) {
      getUser(resolve, reject);
    }).then( function ( user ) {
      loginUser(user, true, resolve, reject);
    }).catch( function ( err ) {
      reject(err);
    });
  }