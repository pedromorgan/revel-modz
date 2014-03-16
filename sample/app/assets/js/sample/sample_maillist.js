function initMaillistView() {

    // function error_handler(msg) {
    //     console.log("Recv'd ERROR: " + msg)
    // }

    // function ack_handler(msg) {
    //     console.log("Recv'd ACK: " + msg)
    // }

    // $("#send_msg_btn").click(function() {
    //     websock_comm.Send($('#comm_input').val());
    // })

    // websock_comm.AddHandler("error", error_handler)
    // websock_comm.AddHandler("ack", ack_handler)
    // websock_comm.Connect(window.location.host, "/comm");



    // setup filter form submit button
    $("#maillist-listview-filter-button").on("click", handle_maillist_filter_button_click);
}

function handle_maillist_filter_button_click(e) {
    // dev debug printing
    console.log("filter button was clicked");
    console.log(e);

    // post filter values to server
    var group = $("#maillist-listview-field-list").val();
    var id = $("#maillist-listview-field-email").val();
    var csrf = $("#csrf_token").val();

    console.log("group", group);
    console.log("id", id);

    // actually send data
    dosend_maillist_filter_update(group, id, csrf);
}

function dosend_maillist_filter_update(group, id, csrf) {

    var post_query = "/a/maillist/filter";
    post_query += "?list=" + encodeURIComponent(group);
    post_query += "&email=" + encodeURIComponent(id);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", post_query, true);
    xhr.setRequestHeader('X-CSRF-Token', csrf);


    xhr.onload = function(e) {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var results = JSON.parse(xhr.responseText);
                console.log(results);
                update_maillist_results_table(results);
            } else {
                console.error(xhr.statusText);
            }
        }
    };



    xhr.onerror = function(e) {
        console.error(xhr.statusText);
    };
    xhr.send(null);
}



function update_maillist_results_table(results) {
    // clear the current table results
    $("#maillist-listview-results").empty();

    // render results in panel
    if (results === null) {
        return
    }

    var template = Hogan.compile(maillist_row_template_text, { delimiters: '<% %>' });
    if (results instanceof Array) {
        for (var i = 0; i < results.length; i++) {
            var output = template.render(results[i]);
            $("#maillist-listview-results").append(output)
        }
    } else {  // should be a single element
            var output = template.render(results);
            $("#maillist-listview-results").append(output)
    }


}

var maillist_row_template_text = [
'    <div class="row">',
'                <div class="large-2 small-2 columns"> <%UserId%>     </div>',
'                <div class="large-2 small-2 columns"> <%Email%>      </div>',
'                <div class="large-2 small-2 columns"> <%List%>       </div>',
'                <div class="large-2 small-2 columns"> <%Activated%>  </div>',
'                <div class="large-4 small-4 columns">',
'                    <ul class="button-group">',
'                        <li><a href="#" class="small button success">View</a>',
'                        </li>',
'                        <li><a href="#" class="small button warning">Edit</a>',
'                        </li>',
'                        <li><a href="#" class="small button alert">Delete</a>',
'                        </li>',
'                    </ul>',
'                </div>',
'            </div>',
].join("\n");


function initMaillistCompose() {
    // button handler
}

function handleMaillistComposeSaveDraft(event) {
    console.log("Saving Draft");
}




