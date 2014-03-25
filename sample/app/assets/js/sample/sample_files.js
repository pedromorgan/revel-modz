
var fileview = fileview || {};//create new empty obj,
/*fileview.curr_tree_selection = ...,
fileview_current_base_dir = getroot
init key_str to fileview.curr_tree_selection*/

function init_dynatree() {


    $("#tree").dynatree({
        checkbox: true,
        // Override class name for checkbox icon:
        classNames: {
            checkbox: "dynatree-radio"
        },
        selectMode: 1,
        children: [],
        onSelect: function(select, node) {
            if (node.isSelected()) {
                // console.log("Activating", node.data.key, node.data.theFile.name);
                if (node.data.theFile != null) {
                    displayFileContents(node.data.theFile);
                }
                if (node.data.contents == null) {
                    getStoredFileContents(node);
                }
                iassic.curr_dataset = node.data;
            } else {
                // console.log("De-activating", node.data.key, node.data.theFile.name);
                $("#file-contents").text("");
                iassic.curr_dataset = null;
            }
            // Display list of selected nodes
        },
        onDblClick: function(node, event) {
            node.toggleSelect();
        },
        onKeydown: function(node, event) {
            if (event.which == 32) {
                node.toggleSelect();
                return false;
            }
        },

        /****************/
        // Kevin:  notice that you redeclared the onSelect function handler (3 above)
        onSelect: function(flag, node) {
            if (!flag) {
                console.log("You deselected node with data title " + node.data.title);

            }
            var selectedNodes = node.tree.getSelectedNodes();
            console.log(selectedNodes);

            var selectedKeys = $.map(selectedNodes, function(node) {
                return node.data.key;
            });
            console.log("Selected keys: " + selectedKeys.join(", "));




        },
        onExpand: function(flag, node) {
            var childNodes = node.childList;
            if (flag) {
                console.log("You expanded node with data title " + node.data.title);


                console.log(childNodes);

                //console.log(childNodes);
                console.log("node: ",node);
                console.log("Path of selected node: ",node.data.key);
                fileview.curr_tree_selection = node.tree
                fileview.curr_path = node.data.key;

                /*var selectedKeys = $.map(childNodes, function(node){
                return node.data.key;
            });*/
                //console.log("Expanded keys: " + selectedKeys.join(", "));
                //console.log(childNodes.length);
                //console.log("about to call renderFileRightPanel with childnodes " + childNodes.data.title);

                renderFileRightPanel(childNodes);

            }
            if (!flag) {
                console.log("You collapsed node with data title " + node.data.title);
                renderFileRightPanel(node.parent.childList);
            }


        },
        /****************/

        // onCreate: function(dtnode, nodeSpan) {
        //     // console.log(nodeSpan);
        //     nodeSpan.oncontextmenu = showContextMenu;
        // },

        // The following options are only required, if we have more than one tree on one page:
        //      initId: "treeData",
        cookieId: "dynatree-Cb1",
        idPrefix: "dynatree-Cb1-"
    });

    // console.log("CSRF_2: ", $("#csrf_token").val());


    getStoredFiles();
    fileview.current_base_dir = $("#tree").dynatree("getRoot");
    $("#tree").dynatree("getRoot").sortChildren(file_cmp, true);
}


var file_cmp = function(a, b) {
    fa = a.data.isFolder;
    fb = b.data.isFolder;
    if (fa && !fb) {
        return -1;
    }
    if (fb && !fa) {
        return 1;
    }
    a = a.data.title.toLowerCase();
    b = b.data.title.toLowerCase();
    return a > b ? 1 : a < b ? -1 : 0;
};

function displayFileContents(node) {

    var reader = new FileReader();

    reader.onload = (function(theFile) {
        return function(e) {
            dataSetFromMyFormat(theFile.name, e.target.results);
        };
    })(file);

    reader.readAsText(file);
}

function openFileOption() {
    document.getElementById("filedata").click();
}
function openAddFileOption(){
    document.getElementById("addfiledata").click();
}
function deleteFile(){

     var tree = fileview.curr_tree_selection;
    var node = $("#tree").dynatree("getActiveNode");

    node.remove();
}
function addDataFile(files) {

    // var files = document.getElementById('filedata').files;
    console.log("in addDataFile, files: ",files);
    var up_files = [];
    $("#progress-bar").css("display", "block");

    var meter = $("#upload-progress")
    meter.width("0%")

    var fCnt = 0;

    var total_sz = 0;
    for (var i = 0, numFiles = files.length; i < numFiles; i++) {
        var f = files[i];
        lfn = f.webkitRelativePath;
        if (lfn[lfn.length - 1] == ".") {
            continue;
        }
        fCnt += 1;
        // up_files.push(f);

        total_sz += f.size;
        // console.log(f.name, f.size, total_sz);
    }

    console.log("total size: ", total_sz);

    var runng_sz = 0;

    var worker = new Worker('/public/js/workers/dirupload.js');

    worker.onmessage = function(e) {
        console.log(e.data);
        runng_sz += parseInt(e.data);
        if (runng_sz >= total_sz) {
            $("#progress-bar").css("display", "none");
            return
        }
        console.log('Worker said: ', e.data, "   rsz", runng_sz);
        meter.width((100.0 * runng_sz / total_sz) + "%")
    }
    worker.onerror = werror;

    function werror(e) {
        console.log('ERROR: Line ', e.lineno, ' in ', e.filename, ': ', e.message);
    }

    //var tree = $("#tree").dynatree("getTree");
    var tree = fileview.curr_tree_selection;
    console.log("current tree selection: ", tree);

    var cnt = 0;
    //console.log("webkitRelativePath: ",files);
    for (var i = 0, numFiles = files.length; i < numFiles; i++) {
        var file = files[i];
        lfn = file.webkitRelativePath;
        lfn_parts = lfn.split("/");
        /*console.log("files[i]: " + i + " " + file);
        console.log("file.webkitrelativepath (lfn): " + lfn);
        console.log("lfn_parts: " + lfn_parts);*/
        if (lfn[lfn.length - 1] == ".") {
            continue;
        }

        //var key_str = "";
        var key_str = fileview.curr_path;

       // var last = $("#tree").dynatree("getRoot");
       var last = $("#tree").dynatree("getActiveNode");

        //var last = fileview.curr_tree_selection;
        console.log("last: ", last);
        for (var j = 0, numParts = lfn_parts.length; j < numParts; j++) {
            key_str += lfn_parts[j] + "/";
            var node = tree.getNodeByKey(key_str);
            if (node == null) {
                var dir = false;
                if (j + 1 < numParts) { // this is a dir
                    dir = true;
                }
                var n_title = lfn_parts[j];
                // if (dir == false) {
                //    n_title += "  " + file.size;
                // }
                var node_data = {
                    title: n_title,
                    key: key_str,
                    isFolder: dir,
                    hideCheckbox: dir,
                    theFile: file
                }


                var f_json = {
                    Name: n_title,
                    Path: key_str,
                    Folder: dir,
                    Size: 0,
                    Content: "",
                    Csrf: $("#csrf_token").val()
                }

                if (!dir) {
                    var reader = new FileReader();
                    reader.onload = (function(f_json, worker) {
                        return function(e) {
                            // console.log(e.target.result);
                            f_json.Content = e.target.result;
                            f_json.Size = f_json.Content.length;
                            worker.postMessage(JSON.stringify(f_json));
                        };
                    })(f_json, worker);

                    reader.readAsText(file);
                } else {
                    worker.postMessage(JSON.stringify(f_json));
                }

                node = last.addChild(node_data);
                tree.redraw();
            }
            last = node;
        }
    }
    //$("#tree").dynatree("getRoot").sortChildren(file_cmp, true);
    fileview.current_base_dir.sortChildren(file_cmp,true);
    // console.log(up_files);
    // worker.postMessage(up_files);
}

function uploadDataFiles(files) {
    // var files = document.getElementById('filedata').files;
    console.log("in uploadDataFiles, files: ", files);
    var up_files = [];
    $("#progress-bar").css("display", "block");

    var meter = $("#upload-progress")
    meter.width("0%")

    var fCnt = 0;

    var total_sz = 0;
    for (var i = 0, numFiles = files.length; i < numFiles; i++) {
        var f = files[i];
        lfn = f.webkitRelativePath;

        if (lfn[lfn.length - 1] == ".") {
            continue;
        }
        fCnt += 1;
        // up_files.push(f);

        total_sz += f.size;
        // console.log(f.name, f.size, total_sz);
    }

    console.log("total size: ", total_sz);

    var runng_sz = 0;

    var worker = new Worker('/public/js/workers/dirupload.js');

    worker.onmessage = function(e) {
        console.log(e.data);
        runng_sz += parseInt(e.data);
        if (runng_sz >= total_sz) {
            $("#progress-bar").css("display", "none");
            return
        }
        console.log('Worker said: ', e.data, "   rsz", runng_sz);
       // console.log("line 291 console test");
        meter.width((100.0 * runng_sz / total_sz) + "%")
    }
    worker.onerror = werror;

    function werror(e) {
        console.log('ERROR: Line ', e.lineno, ' in ', e.filename, ': ', e.message);
    }

    var tree = $("#tree").dynatree("getTree")

    var cnt = 0;

    for (var i = 0, numFiles = files.length; i < numFiles; i++) {
        var file = files[i];
        lfn = file.webkitRelativePath;
        lfn_parts = lfn.split("/");

        if (lfn[lfn.length - 1] == ".") {
            continue;
        }

        var key_str = "";
        var last = $("#tree").dynatree("getRoot");
        for (var j = 0, numParts = lfn_parts.length; j < numParts; j++) {
            key_str += lfn_parts[j] + "/";
            var node = tree.getNodeByKey(key_str);
            if (node == null) {
                var dir = false;
                if (j + 1 < numParts) { // this is a dir
                    dir = true;
                }
                var n_title = lfn_parts[j];
                // if (dir == false) {
                //    n_title += "  " + file.size;
                // }
                var node_data = {
                    title: n_title,
                    key: key_str,
                    isFolder: dir,
                    hideCheckbox: dir,
                    theFile: file
                }


                var f_json = {
                    Name: n_title,
                    Path: key_str,
                    Folder: dir,
                    Size: 0,
                    Content: "",
                    Csrf: $("#csrf_token").val()
                }

                if (!dir) {
                    var reader = new FileReader();
                    reader.onload = (function(f_json, worker) {
                        return function(e) {
                            // console.log(e.target.result);
                            f_json.Content = e.target.result;
                            f_json.Size = f_json.Content.length;
                            worker.postMessage(JSON.stringify(f_json));
                        };
                    })(f_json, worker);

                    reader.readAsText(file);
                } else {
                    worker.postMessage(JSON.stringify(f_json));
                }

                node = last.addChild(node_data);
                tree.redraw();
            }
            last = node;
        }
    }
    $("#tree").dynatree("getRoot").sortChildren(file_cmp, true);
    // console.log(up_files);
    // worker.postMessage(up_files);
}

function getStoredFiles() {
    console.log("Getting stored files")
    req = new XMLHttpRequest();
    req.onreadystatechange = function() {
        // console.log("req.state = ", req.readyState, "   req.status = ", req.status);
        if (req.readyState == 4 && req.status == 200) {
            displayStoredFiles(req.responseText); // Another callback here
        }
    }
    req.open("GET", "/files/query", true);
    req.send();
}

function displayStoredFiles(file_json) {

    // console.log(req.responseText);
    if (req.responseText === "NOFILES") {
        return
    }

    files = JSON.parse(req.responseText);
    if (files == null) {
        return;
    }
    // console.log(files);
    console.log(files.length);

    var tree = $("#tree").dynatree("getTree")
    for (var i = 0, numfs = files.length; i < numfs; i++) {
        var file = files[i];
        lfn = file.Path;
        // console.log(file.DataSetId);
        lfn_parts = lfn.split("/");

        var key_str = "";
        var last = $("#tree").dynatree("getRoot");
        for (var j = 0, numParts = lfn_parts.length; j < numParts; j++) {
            if (lfn_parts[j] == "") {
                continue;
            }
            key_str += lfn_parts[j] + "/";
            var node = tree.getNodeByKey(key_str);
            if (node == null) {
                var dir = file.Folder
                var n_title = lfn_parts[j];
                // if (dir == false) {
                //    n_title += "  " + file.size;
                // }
                var node_data = {
                    title: n_title,
                    key: key_str,
                    isFolder: dir,
                    hideCheckbox: dir,
                    theFile: null,
                    data_id: file.DataSetId
                }
                node = last.addChild(node_data);
            }
            last = node;
        }
        tree.redraw();
    }
}

function getStoredFileContents(node) {
    req = new XMLHttpRequest();
    console.log("node: ", node.data.data_id);
    req.onreadystatechange = function(theNode) {
        return function() {
            if (req.readyState == 4 && req.status == 200) {
                // console.log(req.responseText); // Another callback here
                theNode.data.content = req.responseText;
                dataSetFromMyFormat(theNode.data.title, theNode.data.content);
                // $("#file-contents").html("<p><b>" + theNode.data.title + "<br>========================================</b><br></p><pre>" + htmlEncode(req.responseText) + "</pre>");

            }
        }
    }(node);
    req.open("POST", "/files/content?data_id=" + node.data.data_id, false);
    req.send();
}

function renderFileRightPanel(fileNodes) {


   /* // clear any existing files in the DOM list
    $("#fileview-results").empty();

    // return if no fileNodes
    if (fileNodes === null) {
        return
    }

    // render the file Nodes
    var template = Hogan.compile(file_row_template_text, { delimiters: '<% %>' });
    if (fileNodes instanceof Array) {
        for (var i = 0; i < fileNodes.length; i++) {
            var output = template.render(fileNodes[i]);
            $("#fileview-results").append(output)
        }
    } else {  // should be a single element
            var output = template.render(fileNodes);
            $("#fileview-results").append(output)
    }*/

    // clear any existing files in the DOM list
    $("#file-panel").empty();

    // add header
    $("#file-panel").append(file_list_header);
    //console.log(fileNodes.length)
    var template = Hogan.compile(file_row_template_text, {
        delimiters: '<% %>'
    });
    for (var i = 0; i < fileNodes.length; i++) {
        var output = template.render(fileNodes[i]);
        $("#file-panel").append(output)
    }

}

/*var file_row_template_text = [
    '         <div class="row">',
    '             <div class="small-1 columns"> <%data.isFolder%>         </div>',
    '             <div class="small-4 columns"> <%data.title%>            </div>',
    '             <div class="small-2 columns"> <%data.parent.data.key%>  </div>',
    '             <div class="small-5 columns">',
    '               <ul class="button-group">',
    '                  <li><a href="#" class="tiny button success">View</a></li>',
    '                  <li><a href="#" class="tiny button warning">Edit</a></li>',
    '                  <li><a href="#" class="tiny button alert">Delete</a></li>',
    '               </ul>,',
    '             </div>',
    '         </div>',
].join("\n");
*/
    var file_list_header = '<div class="row">
                <div class="large-12 large-centered small-12 small-centered columns">
                    <div id="user-fileview" class="panel">

                        <div class="row">
                            <div class="large-1 small-1 columns">
                                Folder
                            </div>
                            <div class="large-3 small-3 columns">
                                File Name
                            </div>
                            <div class="large-3 small-3 columns">
                                Owner
                            </div>


                            <div class="large-5 small-5 columns">
                                Status
                            </div>
                        </div>


                    </div>
                </div>
            </div>';





var file_row_template_text = [
    '<div class="large-12 large-centered small-12 small-centered columns">',
    '         <div class="row">',
    '             <div class="large-1 small-1 columns">',
    '                 <%data.isFolder%>',
    '             </div>',
    '             <div class="large-3 small-3 columns">',
    '                 <%data.title%>',
    '             </div>',
    '             <div class="large-3 small-3 columns">',
    '                 <%parent.data.key%>',
    '             </div>',
    '             <div class="large-5 small-5 columns">',
    '                 <ul class="button-group">',
    '        <li><a href="#" class="tiny button success">View</a>',
    '        </li>',
    '        <li><a href="#" class="tiny button warning">Edit</a>',
    '        </li>',
    '        <li><a href="#" class="tiny button alert">Delete</a>',
    '            </li>',
    '   </ul>',
    '             </div>',
    '         </div>',
    '</div>',
].join("\n");
/*>>>>>>> Stashed changes*/
