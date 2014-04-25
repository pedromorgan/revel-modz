function initAnalyticsView() {

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
    $("#analytics-pagereqs-filter-button").on("click", handle_filter_button_click);
}

function dosend_filter_update(group, id, csrf) {

    var post_query = "/a/analytics/filter";
    post_query += "?group=" + encodeURIComponent(group);
    post_query += "&id=" + encodeURIComponent(id);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", post_query, true);
    xhr.setRequestHeader('X-CSRF-Token', csrf);


    xhr.onload = function(e) {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var results = JSON.parse(xhr.responseText);
                console.log(results);
                update_results_table(results);
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

function handle_filter_button_click(e) {
    // dev debug printing
    console.log("filter button was clicked");
    console.log(e);

    // post filter values to server
    var group = $("#analytics-pagereqs-field-group").val();
    var id = $("#analytics-pagereqs-field-id").val();
    var csrf = $("#csrf_token").val();

    console.log("group", group);
    console.log("id", id);

    // actually send data
    dosend_filter_update(group, id, csrf);
}

function update_results_table(results) {
    if results == null || results.length < 1 {
        return
    }
    // clear the current table results

    // render results in panel
    var template = Hogan.compile(analytics_row_template_text, {
        delimiters: '<% %>'
    });
    for (var i = 0; i < results.length; i++) {
        var output = template.render(results[i]);
        $("#analytics-pagereqs-results").append(output)
    }
}

var analytics_row_template_text = [
    '<div class="row">',
    '     <div class="large-12 columns">',
    '         <div class="row">',
    '             <div class="small-2 columns">',
    '                 <%Time%>',
    '             </div>',
    '             <div class="small-2 columns">',
    '                 <%UserId%>',
    '             </div>',
    '             <div class="small-2 columns">',
    '                 <%RequestURI%>',
    '             </div>',
    '             <div class="small-2 columns">',
    '                 <%Host%>',
    '             </div>',
    '             <div class="small-2 columns">',
    '                 <%XRealIp%>',
    '             </div>',
    '             <div class="small-2 columns">',
    '                 <%Referer%>',
    '             </div>',
    '         </div>',
    '     </div>',
    ' </div>',
].join("\n");
