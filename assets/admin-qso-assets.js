$(document).ready(function(){

//  at jquery object clicked

    $("#btn-search").on('click', function() {    //  When clicking on select button

        // data transfer process : POST

        var search_json = '';
         console.log(search_json);
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
          var bodyaddbtn = $('<input style="width: 20%; padding: 10px;" type="button" id="btn-upddel" value="上を確定" /><br />')
          var tabledef = $('<table class="table" id="down-table" bordercolor="#000000">');
          var thead = $('<thead color="blue">');
          var tbody = $('<tbody id="down-tbody">');
          var tr_text = '';
//          tabledef.addClass('table-striped'); //Bootstrapのテーブルを適用するときは必ずtbodyタグを使用する）
          thead.append('<tr><th >I</th><th >U</th><th >D</th><th >ID</th><th >CALLSIGN</th><th >DATETIME</th><th >FILES</th></tr>');

          console.log(result.results);
          let res = JSON.parse(result.results);
          for (let item in res) {
              console.log(res[item]);
              tr_text = ''
              tr_text = '<tr>';
              tr_text = tr_text +'<td ><input type="checkbox" name="insert" value="0"></td>';
              tr_text = tr_text +'<td ><input type="checkbox" name="update" value="1"></td>';
              tr_text = tr_text +'<td ><input type="checkbox" name="delete" value="2"></td>';
              tr_text = tr_text +'<td><input type="text" name="id" value="' + res[item].ID + '"></td>';
              tr_text = tr_text +'<td><input type="text" name="callsign" value="' + res[item].CALLSIGN + '"></td>';
              tr_text = tr_text +'<td><input type="text" name="datetime" value="' + res[item].DATETIME + '"></td>';
              tr_text = tr_text +'<td><input type="text" name="files" value="' + res[item].FILES + '"></td>';
              tr_text = tr_text + '</tr>';
              tbody.append(tr_text);
          }

          tabledef.append(thead);
          tabledef.append(tbody);

          $("#mini-tail").html('');
          $("#mini-tail").append(bodyhead);
          $("#mini-tail").append(tabledef);
          $("#mini-tail").append(bodytail);
          $("#mini-tail").append(bodyaddbtn);

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
        var bodyaddbtn2 = $('<input style="width: 20%; padding: 10px;" type="button" id="btn-inssub" value="追加分確定" /><br />')
        var tabledef2 = $('<table class="table" id="down-table2" bordercolor="#000000">');
        var thead2 = $('<thead color="blue">');
        var tbody2 = $('<tbody id="down-tbody2">');
        var tr_text2 = '';
        thead2.append('<tr><th >I</th><th >U</th><th >D</th><th >ID</th><th >CALLSIGN</th><th >DATETIME</th><th >FILES</th></tr>');
        tr_text2 = '<tr>';
        tr_text2 = tr_text2 + '<td ><input type="checkbox" name="insert" value="0"></td>';
        tr_text2 = tr_text2 + '<td ><input type="checkbox" name="update" value="1"></td>';
        tr_text2 = tr_text2 + '<td ><input type="checkbox" name="delete" value="2"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="id" value="     " maxlength="5"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="callsign" value="      "  maxlength="6"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="datetime" value="                     "  maxlength="25"></td>';
        tr_text2 = tr_text2 + '<td><input type="text" name="files" value="                     "  maxlength="25"></td>';
        tr_text2 = tr_text2 + '</tr>';
        tbody2.append(tr_text2);

        tabledef2.append(thead2);
        tabledef2.append(tbody2);
        $("#mini-tail2").html('');
        $("#mini-tail2").append(bodyhead2);
        $("#mini-tail2").append(tabledef2);
        $("#mini-tail2").append(bodytail2);
        $("#mini-tail2").append(bodyaddbtn2);

      });

      $("#btn-upddel").on('click', function() {    //  When clicking on update/delete button

          // collect data from screen
					console.log("Start upddel");
					var tary = downtableget();

          var record_json = JSON.stringify(tary);
           console.log(record_json);
          $.ajax({
        		url : '/qsl/qslupddel',
        		type : 'POST',
        		data : record_json,
            dataType: 'json',
        		async : true,
        		timeout : 10000,

            // data reception process

        	}).done(function(result, textStatus, xhr) {
        		// 正常系の処理を書く

            console.log(result.results);

            $("#sql-result").html('');
						$("#sql-result").append('Update/Delete operations are completed!');

        	}).fail(function(xhr, textStatus, error) {
        		// 異常系の処理を書く
						$("#sql-result").html('');
						$("#sql-result").append('Update/Delete operations are failed!');
        	}).always(function(xhr, settings) {
        		// 正常・異常系に関わらず行われる処理を書く
        	});
        });   // end of btn-upddel

				function downtableget(){
					var d=[];
					$('#down-table tr').each(function(i, e){
						var dd=[];
						if (i===0)
							$(this).find('th').each(function(j, el){dd.push($(this).text())});
						else
							$(this).find('td').each(function(j, el){dd.push($(this).text())});
						d.push(dd);
					});
					console.log(d);
					return d;

				}





});
