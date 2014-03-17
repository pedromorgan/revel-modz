function initForumView() {
	console.log("initForumView");
	$("#forums-filter-button").on("click", handle_forums_filter_button_click);
}

function handle_forums_filter_button_click(e) {
	// dev debug printing
    console.log("filter button was clicked");
    console.log(e);

    // post filter values to server
    var id = $("#forums-filter-topicid").val();
    var csrf = $("#csrf_token").val();

    console.log("id", id);

    // actually send data
    dosend_forums_filter_update(id, csrf);
}

function dosend_forums_filter_update(id, csrf) {

    var post_query = "";
    post_query += "" + encodeURIComponent(id);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", post_query, true);
    xhr.setRequestHeader('X-CSRF-Token', csrf);


    xhr.onload = function(e) {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var results = JSON.parse(xhr.responseText);
                console.log(results);
                update_forums_results_table(results);
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

function update_forums_results_table(results) {
    // clear the current table results
    $("#forums-topiclist-results").empty();

    // render results in panel
    if (results === null) {
        return
    }

    var template = Hogan.compile(forums_row_template_text, { delimiters: '<% %>' });
    if (results instanceof Array) {
        for (var i = 0; i < results.length; i++) {
            var output = template.render(results[i]);
            $("#forums-topiclist-results").append(output)
        }
    } else {  // should be a single element
            var output = template.render(results);
            $("#forums-topiclist-results").append(output)
    }


}

var forums_row_template_text = [
'    <div class="row">',
'    	<div class="large-3 small-3 columns"> <%TopicId%> 	</div>',
'       <div class="large-3 small-3 columns"> <%TopicName%> </div>',
'       <div class="large-3 small-3 columns"> <%AuthorId%> 	</div>',
'       <div class="large-3 small-3 columns"> <%CreatedAt%> </div>',
'    </div>',
].join("\n");
