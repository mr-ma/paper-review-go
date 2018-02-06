	 	
var edit = ' glyphicon glyphicon-pencil';
var trash  = 'glyphicon glyphicon-trash';
var save   = 'glyphicon glyphicon-floppy-disk';
var cancel = 'glyphicon glyphicon-floppy-remove';

	function enableButtons() {
		$('#addRow').prop('disabled', false);
		$('.glyphicon-pencil').unbind().css('cursor','pointer').on('click', function (e) { disableButtons(); editRow(e); });
	    $('.glyphicon-trash').unbind().css('cursor','pointer').on('click', function  (e) {
	      e.preventDefault();
	      var row = $(e.target).parentsUntil('tbody');
	      row.remove();
	    });
	}

	function disableButtons () {
		$('#addRow').prop('disabled', true);
		$('.glyphicon-pencil').unbind();
		$('.glyphicon-trash').unbind();
	}

	function addRow () {
		$('#showTable').append('<tr>' +
		   '<td><input type="text" style="width:250px;" title="Synonym" /></td>' +
	       '<td><span  class="' + save + '"></span></td>' + 
	       '<td><span style="padding-left: 10px;" class="' + cancel + '"></span></td></tr>');
		$('.glyphicon-floppy-remove').unbind().css('cursor','pointer').on ( 'click', function (e) { e.preventDefault(); deleteRow(e); enableButtons(); });
	  $('.glyphicon-floppy-disk').unbind().css('cursor','pointer').on ( 'click', function (e) { 		    	
	    	e.preventDefault(); saveRow(e); });
	}

	function deleteRow(e) {  $(e.target).parentsUntil('tbody').remove(); }
	function saveRow(e) {
	  	e.preventDefault();
	  	var row = $(e.currentTarget).parentsUntil('tr');
		row.parent().find('td').each( function ( index, cell) {
		 	var cellText = $(cell).find('input').val();
		  switch(index) {
		  	case 1: $(cell).html('<span class="' + edit + '"></span>'); break;
		  	case 2: $(cell).html('<span class="' + trash + '"></span>'); break;
		  	default: $(cell).html('<span style="margin-right:10px;">' + cellText + '</span>');
				}
		});
	  	enableButtons(); 
	}

	function updateRow( e ) {
	  	e.preventDefault();
	  	var row = $(e.currentTarget).parentsUntil('tr');
		row.parent().find('td').each( function ( index, cell) {
		 	var cellText = $(cell).find('input').val();
		  switch(index) {
		  	case 1: $(cell).html('<span class="' + edit + '"></span>'); break;
		  	case 2: $(cell).html('<span class="' + trash + '"></span>'); break;
		  	default: $(cell).html('<span style="margin-right:10px;">' + cellText + '</span>');
				}
		});
	 	enableButtons();
	}

	function editRow( e ) {
    	e.preventDefault();
    	disableButtons();
	    var row = $(e.currentTarget).parentsUntil('tr');
	    var rowIndex = row.parent()[0].rowIndex;
	    row.parent().find('td').each( function ( index, cell) {
	     	var cellText = $(cell).text();
	      switch(index) {
	      	case 1: $(cell).html('<span class="' + save + '" id="saveButton_' + rowIndex + '"></span>'); break;
	      	case 2: $(cell).html('<span class="' + cancel + '" id="cancelButton_' + rowIndex + '"></span>'); break;
	      	default: $(cell).html('<input type="text" style="width:250px;" data-lastText="' + cellText + '" value="' + cellText + '"/>');
				}
	    });
	    $('#cancelButton_' + rowIndex).unbind().css('cursor','pointer').on ( 'click', function (e) {
	    	e.preventDefault();
		    row.parent().find('td').each( function ( index, cell) {
		     	var cellText = $(cell).find('input').attr('data-lastText');
		      switch(index) {
		      	case 1: $(cell).html('<span class="' + edit + '"></span>'); break;
		      	case 2: $(cell).html('<span class="' + trash + '"></span>'); break;
		      	default: $(cell).html('<span style="margin-right:10px;">' + cellText + '</span>');
					}
		    });
	      	enableButtons();
	   	});
		$('#saveButton_' + rowIndex).unbind().css('cursor','pointer').on ( 'click', function (e) { 		    	
			e.preventDefault();
			updateRow(e);
		});
	}

	function initSearchbar() {
	  $('.remove').unbind().click( function () {
	    var textField = $(this).parent().parent().find('input');
	    textField.val('');
	    textField.trigger('input');
	  });

	  $('#searchTableField').unbind().on('input', function () {
	    var text = $(this).val().toLowerCase();
	    var showRowWithIndex = -1;
	    $('#showTable').find('tr').each( function ( rowIndex, row ) {
	    	$(row).find('td').each( function ( cellIndex, cell ) {
					if ($(cell).text().toLowerCase().split(text).length > 1) {
		        $(row).show();
		        showRowWithIndex = rowIndex;
		      } else if (text.length > 1 && rowIndex != showRowWithIndex && cellIndex == 2) { $(row).hide();  } // cellIndex == 2 -> last column with text
		    });
	    });
	  });
	}

	function loadTableListeners () {
	    initSearchbar();
	    $('.glyphicon-pencil').unbind().css('cursor','pointer').on('click', function (e) { disableButtons(); editRow(e); });
	    $('.glyphicon-trash').unbind().css('cursor','pointer').on('click', function  (e) {
	      e.preventDefault();
	      var row = $(e.target).parentsUntil('tbody');
	      row.remove();
	    });
	    $('#addRow').unbind().on( 'click', function () { $(this).prop('disabled', true );  addRow(); });
 	};
	
	function initTable( tableHeader, rowData ) {
		$('#addRow').prop('disabled', false);
		var table = '<div class="table-responsive" style="max-height:300px;overflow-y:auto;"><table class="table table-striped table-hover" id="showTable">' + tableHeader + '<tbody>';
		rowData.forEach ( function (entry) { 
		  table = table + '<tr><td><span style="margin-right:10px;">' + entry + '</span></td><td><span class="' + edit + '"></span></td><td><span class="' + trash + '"></span></td></tr>';
		}); 
		table+='</tbody></table></div>';
		return table;
	}
