// functions used to add, delete, change user accounts and permissions	 	
var edit = ' glyphicon glyphicon-pencil';
var trash  = 'glyphicon glyphicon-trash';
var save   = 'glyphicon glyphicon-floppy-disk';
var cancel = 'glyphicon glyphicon-floppy-remove';
var settings = 'glyphicon glyphicon-cog taxonomySettings';

	function enableButtons() {
		$('#add').prop('disabled', false);
		$('.glyphicon-pencil').css('cursor','pointer').on('click', function (e) { disableButtons(); editRow(e); });
	}

	function disableButtons () {
		$('#add').prop('disabled', true);
		$('.glyphicon-pencil').unbind();
	}

	function addRow () {
		$('#showTable').append('<tr>' +
	       '<td><input type="text" title="Email" placeholder ="Email" value="" /></td>' +
	       '<td><input type="text" title="Password" value="" /></td>' +
	       '<td></td>' +
	       '<td><input type="checkbox" title="Admin" value="admin" name="admin"></td>' +
	       '<td><span  class="' + save + '"></span></td>' + 
	       '<td><span style="padding-left: 10px;" class="' + cancel + '"></span></td></tr>');
		$('.glyphicon-floppy-remove').css('cursor','pointer').on ( 'click', function (e) { e.preventDefault(); deleteRow(e); enableButtons(); });
	  $('.glyphicon-floppy-disk').css('cursor','pointer').on ( 'click', function (e) { 		    	
	    	e.preventDefault(); saveRow(e); });
	}

	function deleteRow(e) {  $(e.target).parentsUntil('tbody').remove(); }
	function saveRow(e) { 
	  var email = $(e.target).parentsUntil('tbody').find('td:first').find('input').val();
	  var password = $(e.target).parentsUntil('tbody').find('td:nth(1)').find('input').val();
	  var admin = $(e.target).parentsUntil('tbody').find('td:nth(3)').find('input').is(':checked') ? 1 : 0;
	  if (email == '') { BootstrapDialog.show({message: 'The email field cannot be empty.'}); return; }
        var user = {};
        user.email = email;
        user.password = password;
        user.admin = admin;
        console.log('email: ' + user.email)
        if (!validateEmail(user.email)) {
        	handleError('Invalid email address');
        	return;
        }
        $.ajax
          ({
            type: "POST",
            url: 'createUser',
            dataType: 'json',
            contentType:'application/json',
            async: true,
            xhrFields: {
               withCredentials: true
            },
            data: JSON.stringify({email: user.email, password: user.password}),
            success: function ( result, status, xhr ) {
              if (!result || !result.success) {
                handleError('A user with this email already exists.');
                return;
              }
		      enableButtons();
		      window.location.reload(true);
            }
        });
	}

	function updateRow( e ) {
	  var email = $(e.target).parentsUntil('tbody').find('td:first').text();
	  var admin = $(e.target).parent().parent().find('td:nth(2)').find('input').is(':checked') ? 1 : 0;
	  if (email == '') { BootstrapDialog.show({message: 'The email field cannot be empty.'}); return; }
	  console.log('updating user: ' + email + ', admin: ' + admin)
      $.ajax
        ({
          type: "POST",
          url: 'updateUser',
          dataType: 'json',
          contentType:'application/json',
          async: true,
          xhrFields: {
             withCredentials: true
          },
          data: JSON.stringify({email: email, admin: admin}),
          success: function ( result, status, xhr ) {
            if (!result || !result.success) {
              handleError('Cannot update user with email: ' + email + '.');
              return;
            }
		    enableButtons();
		    window.location.reload(true);
          }
      });
	}

	function editRow( e ) {
    e.preventDefault();
    disableButtons();
	var email = $(e.target).parentsUntil('tbody').find('td:first').text();
	var row = $(e.currentTarget).parentsUntil('tr');
	var rowIndex = row.parent()[0].rowIndex;
	row.parent().find('td').each( function ( index, cell) {
	 	var cellText = $(cell).text();
	  switch(index) {
	  	case 2: $(cell).html(''); break;
	  	case 3: $(cell).html('<input type="checkbox" id="adminCheck_' + rowIndex + '" title="Admin" value="admin" name="admin">');
	      		$('#adminCheck_' + rowIndex).prop('checked', cellText == 'yes' ? true : false); break;
	  	case 4: $(cell).html('<span class="' + save + '" id="saveButton_' + rowIndex + '"></span>'); break;
	  	case 5: $(cell).html('<span class="' + cancel + '" id="cancelButton_' + rowIndex + '"></span>'); break;
	  	default: //$(cell).html('<input type="text" title="Email" value="' + cellText + '"/>');
			}
	});
	$('#cancelButton_' + rowIndex).unbind().css('cursor','pointer').on ( 'click', function (e) {
		e.preventDefault();
	  enableButtons();
	  window.location.reload(true);
	});
	$('#saveButton_' + rowIndex).unbind().css('cursor','pointer').on ( 'click', function (e) { 		    	
	e.preventDefault();
	updateRow(e);
	});
	}

	function initSearchbar() {
	  $('.remove').click( function () {
	    var textField = $(this).parent().parent().find('input');
	    textField.val('');
	    textField.trigger('input');
	  });

	  $('#searchFieldUsers').on('input', function () {
	    var text = $(this).val().toLowerCase();
	    var showRowWithIndex = -1;
	    $('#showTable').find('tr').each( function ( rowIndex, row ) {
	    	$(row).find('td').each( function ( cellIndex, cell ) {
					if ($(cell).text().toLowerCase().split(text).length > 1) {
		        $(row).show();
		        showRowWithIndex = rowIndex;
		      } else if (text.length > 1 && rowIndex != showRowWithIndex && cellIndex == 4) { $(row).hide();  } // cellIndex == 4 -> last column with text
		    });
	    });
	  });
	}

	function loadListener () {
	    initSearchbar();
	    $('.taxonomySettings').unbind().css('cursor','pointer').on('click', function (e) {
	      var email = $(e.target).parentsUntil('tbody').find('td:first').text();
	      if (!email || email == '') {
	      	handleError('Cannot change user taxonomy edit permissions for user with empty email.');
	      	return;
	      }
	      $.ajax
	        ({
	            type: "POST",
	            url: 'taxonomyPermissions',
	            dataType: 'json',
	            contentType:'application/json',
	            async: true,
	            data: JSON.stringify({email: email}),
	            success: function ( taxonomyPermissions ) {
	            	if (!taxonomyPermissions) taxonomyPermissions = [];
			      	var taxonomyIDs = [];
			      	taxonomyPermissions.forEach ( function ( taxonomyPermission ) {
			      		taxonomyIDs.push(taxonomyPermission.id);
			      	});
			      	console.log("permissions: ", taxonomyIDs);
				      $.get('taxonomy', function ( taxonomies ) {
				      	if (!taxonomies) {
				      		var msg = 'Error loading taxonomies.';
				      		//if (!!handleErrorHelper) handleErrorHelper(msg);
				      		handleError(msg);
				      		return;
				      	}
				      	if (!taxonomies.response) taxonomies.response = [];
				      	var taxonomyArray = [];
				      	var allTaxonomyIDs = [];
				      	taxonomies.response.forEach ( function ( taxonomy ) {
				      		taxonomyArray.push({id: taxonomy.id, title: taxonomy.text});
				      		allTaxonomyIDs.push(taxonomy.id);
				      	});
					    BootstrapDialog.show({
					    	title: 'Taxonomy Edit Permissions',
					    	message: '<div id="taxonomyTableContainer"></div>',
					    	closeable : true,
					    	onshown: function () {
						     var fields = [{
						     	field: 'id',
						     	title: 'ID',
						     	visible: false
						     },
						     {
						       field: 'title',
						       title: 'Taxonomy'
						     }];
						     var selected = taxonomyArray.filter( function ( taxonomy ) {
						     	return taxonomyIDs.indexOf(taxonomy.id) >= 0;
						     });
						     var notSelected = taxonomyArray.filter ( function ( taxonomy ) {
						     	return taxonomyIDs.indexOf(taxonomy.id) < 0;
						     })
						     var list = selected.concat(notSelected);
						     showTable(list, 'taxonomyTable', true, false, [], fields, 10);
						     $('#taxonomyTable').bootstrapTable('checkBy', { field : 'id', values : taxonomyIDs });
						   },
				            buttons: [{
				                label: 'Cancel',
				                cssClass: 'btn',
				                action: function ( dialogRef ) {
				                    dialogRef.close();
				                }
				            }, {
				                label: 'Save',
				                cssClass: 'btn-primary',
				                action: function ( dialogRef ) {
							   		var selections = '';
							   		$('#taxonomyTable').bootstrapTable('getSelections').forEach ( function ( entry ) {
							   			if (entry.id > 0) selections += (selections != '' ? ',' : '') + entry.id;
							   		});
								    $.ajax
								        ({
								          type: "POST",
								          url: 'updateTaxonomyPermissions',
								          dataType: 'json',
								          contentType:'application/json',
								          async: true,
								          data: JSON.stringify({email: email, permissions: selections}),
								          success: function ( result ) {
								          	if (!result || !result.success) {
								          		var msg = 'Cannot update taxonomy permissions for user with email: ' + email + '.';
								          		//if (!!handleErrorHelper) handleErrorHelper(msg);
								          		handleError(msg);
								          		return;
								          	}
								          	dialogRef.close();
								          }
								        });
				                }
				            }]
						  });
					    });
				    }
				});
	    	});
	    $('.glyphicon-pencil').css('cursor','pointer').on('click', function (e) { disableButtons(); editRow(e); });
	    $('.glyphicon-trash').css('cursor','pointer').on('click', function  (e) {
	      e.preventDefault();
	      var email = $(e.target).parentsUntil('tbody').find('td:first').text();
	      BootstrapDialog.confirm({  
	      	title : 'Deleting a user',
	      	message : 'Delete user: ' + email + '?',
	        callback : function(result){
	      if(result) {
	        $.ajax
	          ({
	            type: "POST",
	            url: 'deleteUser',
	            dataType: 'json',
	            contentType:'application/json',
	            async: true,
	            xhrFields: {
	               withCredentials: true
	            },
	            data: JSON.stringify({email: email}),
	            success: function ( result, status, xhr ) {
	              if (!result || !result.success) {
	                handleError('Cannot delete user with email: ' + email + '.');
	                return;
	              }
			      enableButtons();
			      window.location.reload(true);
	            }
	        });
	  		} } });
	    });
	    $('#add').on ( 'click', function () { $(this).prop('disabled', true );  addRow(); });
 	};
	
	function initUsers() {
		$('#add').prop('disabled', false);
		$.get("getUsers",function( data ) {
			console.log(data)
		 	if (!data) {
		 		handleError('Cannot get users from DB.');
		 		return;
		 	}
			var table = '<div class="table-responsive"><table class="table table-striped table-hover" id="showTable"><thead><tr><th style="width:20%;">Email</th><th style="width:20%;">Password</th><th>Permissions</th><th>Admin</th></tr></thead><tbody>';
			data.forEach ( function (entry) { 
			    table = table + '<tr><td style="width:20%;">' + entry.email + '</td><td style="width:20%;"></td><td><span class="' + settings + '" title="Edit taxonomy permissions"></span></td><td>' + (entry.admin == 1 ? 'yes' : 'no') + '</td><td><span class="' + edit + '"></span></td><td><span class="' + trash + '"></span></td></tr>';  }); 
				table+='</tbody></table></div>';
				$(table).appendTo('#tblSpan');
				loadListener();
		});
	}
