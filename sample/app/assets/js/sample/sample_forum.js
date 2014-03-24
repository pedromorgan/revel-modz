function initForumView() {
	console.log("initForumView");
	$("#forum-filter-button").on("click", handle_forum_filter_button_click);
}

function handle_forum_filter_button_click(e) {
	// dev debug printing
    console.log("filter button was clicked");
    console.log(e);

    // post filter values to server
    var id = $("#forum-filter-topicid").val();
    var csrf = $("#csrf_token").val();

    console.log("id", id);

    // actually send data
    dosend_forum_filter_update(id, csrf);
}

function dosend_forum_filter_update(id, csrf) {

    var post_query = "/forum/filter";
    post_query += "?id=" + encodeURIComponent(id);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", post_query, true);
    xhr.setRequestHeader('X-CSRF-Token', csrf);


    xhr.onload = function(e) {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                var results = JSON.parse(xhr.responseText);
                console.log(results);
                update_forum_topiclist_results_table(results);
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

function update_forum_topiclist_results_table(results) {
    // clear the current table results
    $("#forum-topiclist-results").empty();

    // render results in panel
    if (results === null) {
        return
    }

    var template = Hogan.compile(forum_topic_row_template_text, { delimiters: '<% %>' });
    if (results instanceof Array) {
        for (var i = 0; i < results.length; i++) {
            var output = template.render(results[i]);
            $("#forum-topiclist-results").append(output)
        }
    } else {  // should be a single element
            var output = template.render(results);
            $("#forum-topiclist-results").append(output)
    }


}

var forum_topic_row_template_text = [
'    <div class="row">',
'    	<div class="large-3 small-3 columns"> <%TopicId%> 	</div>',
'       <div class="large-3 small-3 columns"> <%TopicName%> </div>',
'       <div class="large-3 small-3 columns"> <%AuthorId%> 	</div>',
'       <div class="large-3 small-3 columns"> <%CreatedAt%> </div>',
'    </div>',
].join("\n");
