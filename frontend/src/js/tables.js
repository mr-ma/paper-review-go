  function adjustTableLayout ( tableID, buttons ) {
      // adjust cell width
      $('#' + tableID).find('th:first').css('width','1%'); // Stats
      $('#' + tableID).find('th:nth(1)').css('text-align','center').css('width','1%'); // version
      $('#' + tableID).find('th:nth(2)').css('width','15%'); // date
      $('#' + tableID).find('th:nth(4)').css('width','15%'); // edited by
      $('#' + tableID).find('th:nth(8)').css('width','1%'); // buttons
      $('.ml10').css('margin-left','10px');
      
      $('#' + tableID).find('tr').each ( function (index,row) { 
        if ( index == 0 ) return;
        //$(row).find('td:nth(1)').css('text-align','center'); // version
        //$(row).find('td:nth(2)').css('text-align','center'); // date
        //$(row).find('td:nth(3)').css('text-align','center'); // edited by
        if (buttons.length > 0) $(row).find('td:last').css('text-align','center'); // button
      });
      
  }

 function initTable( data, tableID, hasCheckboxes, expandable, buttons, fields, pageSize ) {
    $('#' + tableID + 'Container').html('<table id="' + tableID + '"><thead></thead><tbody></tbody></table>');
    var columns = [{
          field: 'select',
          title: '',
          checkbox: true,
          searchable: false
        }];
      columns = columns.concat(fields);
      if (!hasCheckboxes) columns.shift();
      buttons.forEach ( function ( button ) {
        columns.push({field: button.name, title: button.title, searchable: false});
      });
      $('#' + tableID).bootstrapTable({
        columns: columns,
        data: data,
        detailView: !!expandable,
        search: true,
        searchAlign: 'left',
        clickToSelect: true,
        maintainSelected: true,
        pageSize: !!pageSize ? pageSize : 10,
        pagination: data.length > (!!pageSize ? pageSize : 10) ? true : false,
        onExpandRow: function(index, row, detail) {
          detail.html(row.nested.content);
        }
      });
  }

  function showTable ( tableData, tableID, hasCheckboxes, expandable, buttons, fields, pageSize ) {
    var data = [];
    var counter = 1;
    if (!buttons) var buttons = [];
    tableData.forEach ( function (entry) {
      var nestedColumns = {'content': entry.content};
      var dataRow = {'rowIndex': counter, 'nested': nestedColumns};
      fields.forEach ( function ( field ) {
        if (entry[field.field] || '' + entry[field.field] == '0') {
          if (!!entry[field.field].length && entry[field.field].length > 50) dataRow[field.field] = '<div class="limitedTextField">' + entry[field.field] + '</div>';
          else dataRow[field.field] = entry[field.field];
        }
      });
      if (!!dataRow.modifiedWhen) dataRow.modifiedWhen = getDateAndTime(entry.modifiedWhen);
      if (!!dataRow.loggedWhen) dataRow.loggedWhen = getDateAndTime(entry.loggedWhen);
      buttons.forEach ( function ( button ) {
        if (!dataRow[button.name]) {
          if (button['data-field'] && entry[button['data-field']]) dataRow[button.name] = '<span class="glyphicon ' + button.class + '" data-field="' + entry[button['data-field']] + '" style="cursor:pointer;vertical-align:middle;"' +
                                                                                          ' title="' + button.title + '"></span>';
          else dataRow[button.name] = '<span class="glyphicon ' + button.class + '" style="cursor:pointer;vertical-align:middle;" title="' + button.title + '"></span>';
        }
      });
      data.push(dataRow);
      counter++;
    });
    initTable(data, tableID, hasCheckboxes, expandable, buttons, fields, pageSize);
    if (buttons) {
      buttons.forEach ( function ( button ) {
        $('.' + button.class).unbind().on('click', function ( event ) {
          button.onClick(event, this);
        });
      });
    }
    adjustTableLayout(tableID, buttons);
  }