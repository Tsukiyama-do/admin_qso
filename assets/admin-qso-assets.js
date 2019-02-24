function file_upload(){
    // フォームデータを取得
    var formdata = new FormData($('#upload_form').get(0));
    console.log(formdata);

    // POSTでアップロード
    $.ajax({
        url  : "/qsl/fileupload",
        type : "POST",
        data : formdata,
        cache       : false,
        contentType : false,
        processData : false,
        dataType    : "json"
    })
    .done(function(result, textStatus, xhr){
//        alert(data);

        console.log(result);
//        let res = JSON.parse(result);

        $("#mini-ftail2").html('');
        $("#mini-ftail2").html('アップロードに成功しました。' + result.results);

    })
    .fail(function(xhr, textStatus, error){
      // 異常系の処理を書く
      $("#mini-ftail2").html('');
      $("#mini-ftail2").html('アップロードに失敗しました。');
    })
    .always(function(xhr, settings) {
    // 正常・異常系に関わらず行われる処理を書く
    });
}

$(document).ready(function(){

//  at jquery object clicked

    $("#btn-search").on('click', function() {    //  When clicking on select button

        // data transfer process : POST

        var search_json = '';
//         console.log(search_json);
        $.ajax({
      		url : '/qsl/qslselect',
      		type : 'POST',
      		data : search_json,
          dataType: 'json',
      		async : true,
      		timeout : 10000,

          // data reception process

      	}).done(function(result, textStatus, xhr) {
      		// 正常系の処理を書く
//          console.log(result);

          var bodyhead = $('<div id="tablebefore">以下に検索した結果を示します。</div>')
          var bodytail = $('<div id="tableafter">処理は終了です。</div>')

          var tabledef = $('<table class="table" id="down-table" bordercolor="#000000">');
          var thead = $('<thead color="blue">');
          var tbody = $('<tbody id="down-tbody">');
          var tr_text = '';
//          tabledef.addClass('table-striped'); //Bootstrapのテーブルを適用するときは必ずtbodyタグを使用する）
          thead.append('<tr><th >I</th><th >U</th><th >D</th><th >ID</th><th >CALLSIGN</th><th >DATETIME</th><th >FILES</th></tr>');

  //        console.log(result.results);
          let res = JSON.parse(result.results);
          for (let item in res) {
//              console.log(res[item]);
              tr_text = ''
              tr_text = '<tr>';
              tr_text = tr_text +'<td></td>';
              tr_text = tr_text +'<td class="cbu"><input type="checkbox" name="update"></td>';
              tr_text = tr_text +'<td class="cbd"><input type="checkbox" name="delete"></td>';
              tr_text = tr_text +'<td><label>' + res[item].ID + '    </label></td>';
              tr_text = tr_text +'<td><input type="text" name="callsign" value="' + res[item].CALLSIGN + '" maxlength="7"></td>';
              tr_text = tr_text +'<td><input type="text" name="datetime" value="' + res[item].DATETIME + ' "maxlength="30"></td>';
              tr_text = tr_text +'<td><input type="text" name="files" value="' + res[item].FILES + '" maxlength="40"></td>';
              tr_text = tr_text + '</tr>';
              tbody.append(tr_text);
          }

          tabledef.append(thead);
          tabledef.append(tbody);

          $("#mini-tail").html('');
          $("#mini-tail").append(bodyhead);
          $("#mini-tail").append(tabledef);
          $("#mini-tail").append(bodytail);

      	}).fail(function(xhr, textStatus, error) {
      		// 異常系の処理を書く
          $("#mini-tail").html('');
      	}).always(function(xhr, settings) {
      		// 正常・異常系に関わらず行われる処理を書く
      	});
      });   // end of btn-search

      $("#btn-src-rst").on('click', function() {    //  When clicking on clear-button

        $("#mini-tail").html('');


      });

      $("#btn-insert").on('click', function() {    //  When clicking on insert-button

        var bodyhead2 = $('<div id="tablebefore2"></div>')
        var bodytail2 = $('<div id="tableafter2"></div>')
        var tabledef2 = $('<table class="table" id="down-table2" bordercolor="#000000">');
        var thead2 = $('<thead color="blue">');
        var tbody2 = $('<tbody id="down-tbody2">');
        var tr_text2 = '';
        thead2.append('<tr><th >I</th><th >U</th><th >D</th><th >ID</th><th >CALLSIGN</th><th >DATETIME</th><th >FILES</th></tr>');
        tr_text2 = '<tr>';
        tr_text2 = tr_text2 + '<td class="cbi2"><input type="checkbox" name="insert"></td>';
        tr_text2 = tr_text2 + '<td ></td>';
        tr_text2 = tr_text2 + '<td ></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="id" value="" maxlength="6"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="callsign" value=""  maxlength="7"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="datetime" value=""  maxlength="30"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="files" value=""  maxlength="40"></td>';
        tr_text2 = tr_text2 + '</tr>';
        tbody2.append(tr_text2);

        tabledef2.append(thead2);
        tabledef2.append(tbody2);
        $("#mini-tail2").html('');
        $("#mini-tail2").append(bodyhead2);
        $("#mini-tail2").append(tabledef2);
        $("#mini-tail2").append(bodytail2);

      });

      $("#btn-upddel").on('click', function() {    //  When clicking on update/delete button

          // collect data from screen
					console.log("Start upddel");

          var s_obj = [];    //  Updated objects


          $('td.cbd').parents('tr').each(function() {
            if ($(this).find('input[name="delete"]').prop('checked')) {
              var s_wk = new Object();
              s_wk.mode = "D";
              s_wk.id = $(this).find('input[name="id"]').val();
              s_wk.callsign = $(this).find('input[name="callsign"]').val();
              s_wk.datetime = $(this).find('input[name="datetime"]').val();
              s_wk.files = $(this).find('input[name="files"]').val();
              s_obj.push(s_wk);
            }
          });
          $('td.cbu').parents('tr').each(function() {
            if ($(this).find('input[name="update"]').prop('checked')) {
              var s_wk = new Object();
              s_wk.mode = "U";
              s_wk.id = $(this).find('input[name="id"]').val();
              s_wk.callsign = $(this).find('input[name="callsign"]').val();
              s_wk.datetime = $(this).find('input[name="datetime"]').val();
              s_wk.files = $(this).find('input[name="files"]').val();
              s_obj.push(s_wk);
            }
          });

          // end of getting table
          var record_json = JSON.stringify(s_obj);
           console.log(record_json);
          $.ajax({
        		url : '/qsl/qslupddel',
        		type : 'POST',
        		data : record_json,
            dataType: 'text',
        		async : true,
        		timeout : 10000,

            // data reception process

        	}).done(function(result, textStatus, xhr) {
        		// 正常系の処理を書く

            console.log(result);

            $("#sql-result").html('');
						$("#sql-result").append('Update/Delete operations are completed!');

        	}).fail(function(xhr, textStatus, error) {
        		// 異常系の処理を書く
            alert( "Request failed: " + textStatus );
            alert( "Request failed error: " + error );
						$("#sql-result").html('');
						$("#sql-result").append('Update/Delete operations are failed!');
        	}).always(function(xhr, settings) {
        		// 正常・異常系に関わらず行われる処理を書く
        	});
        });   // end of btn-upddel


        $("#btn-inssub").on('click', function() {    //  When clicking on insert button

            // collect data from screen
  					console.log("Start inssub");

            var s_obj = [];    //  Updated objects


            $('td.cbi2').parents('tr').each(function() {
              if ($(this).find('input[name="insert"]').prop('checked')) {
                  var s_wk = new Object();
                  s_wk.mode = "I";
                  s_wk.id = $(this).find('input[name="id"]').val();
                  s_wk.callsign = $(this).find('input[name="callsign"]').val();
                  s_wk.datetime = $(this).find('input[name="datetime"]').val();
                  s_wk.files = $(this).find('input[name="files"]').val();
                  s_obj.push(s_wk);
                }
                });

            // end of getting table
            var record_json = JSON.stringify(s_obj);
             console.log(record_json);
            $.ajax({
          		url : '/qsl/qslinsert',
          		type : 'POST',
          		data : record_json,
              dataType: 'text',
          		async : true,
          		timeout : 10000,

              // data reception process

          	}).done(function(result, textStatus, xhr) {
          		// 正常系の処理を書く

            //  console.log(result);

              $("#sql-result").html('');
  						$("#sql-result").append('Insert operations are completed!');

          	}).fail(function(xhr, textStatus, error) {
          		// 異常系の処理を書く
              alert( "Request failed: " + textStatus );
              alert( "Request failed error: " + error );
  						$("#sql-result").html('');
  						$("#sql-result").append('Insert operations are failed!');
          	}).always(function(xhr, settings) {
          		// 正常・異常系に関わらず行われる処理を書く
          	});
          });   // end of btn-upddel


          $("#btn-fsearch").on('click', function() {    //  When clicking on select button

              // data transfer process : POST

              var search_json = '';
      //         console.log(search_json);
              $.ajax({
            		url : '/qsl/fileselect',
            		type : 'POST',
            		data : search_json,
                dataType: 'json',
            		async : true,
            		timeout : 10000,

                // data reception process

            	}).done(function(result, textStatus, xhr) {
            		// 正常系の処理を書く
      //          console.log(result);

                var bodyhead = $('<div id="tablebeforef"> ~/github.com/Tsukiyama-do/qso/uploads 配下のファイルは以下です。</div>')
                var bodytail = $('<div id="tableafterf">処理は終了です。</div>')

                var tabledef = $('<table class="table" id="down-tablef" bordercolor="#000000">');
                var thead = $('<thead color="blue">');
                var tbody = $('<tbody id="down-tbodyf">');
                var tr_text = '';
      //          tabledef.addClass('table-striped'); //Bootstrapのテーブルを適用するときは必ずtbodyタグを使用する）
                thead.append('<tr><th >File Name</th></tr>');

                console.log(result.results);
                let res = JSON.parse(result.results);
                for (let item in res) {
      //              console.log(res[item]);
                    tr_text = ''
                    tr_text = '<tr>';
                    tr_text = tr_text +'<td><label>' + res[item].filename + '    </label></td>';
                    tr_text = tr_text + '</tr>';
                    tbody.append(tr_text);
                }

                tabledef.append(thead);
                tabledef.append(tbody);

                $("#mini-ftail").html('');
                $("#mini-ftail").append(bodyhead);
                $("#mini-ftail").append(tabledef);
                $("#mini-ftail").append(bodytail);

            	}).fail(function(xhr, textStatus, error) {
            		// 異常系の処理を書く
                $("#mini-ftail").html('');
            	}).always(function(xhr, settings) {
            		// 正常・異常系に関わらず行われる処理を書く
            	});
            });   // end of btn-search


});
