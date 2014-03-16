var websock_comm = (function($) {

    // private variables
    var _socket;
    var _handlers = {};
    var _host = window.location.host;
    var _url = "/comm"

    // public interface
    var WS_COMM = {};

    WS_COMM.AddHandler = function(tag, handle) {
        _handlers[tag] = handle;
    }

    WS_COMM.RemoveHandler = function(tag, handle) {
        delete _handlers[tag];
    }

    WS_COMM.Connect = function(host, url, secure) {
        if (host == undefined || host == null || host == "") {
            host = _host;
        }
        if (url == undefined || url == null || url == "") {
            url = _url;
        }

        if (secure === true) {
            _socket = new WebSocket('wss://' + host + url);
        } else {
            _socket = new WebSocket('ws://' + host + url);
        }

        _socket.onopen = function() {
            console.log("Connection established");
        };

        // Log errors
        _socket.onerror = function(error) {
            console.log('WebSocket Error ' + error);
            console.log(error);
        };

        _socket.onclose = function(e) {
            console.log('Closing: ' + e.data);
            console.log(e);
        };

        _socket.onmessage = function(e) {
            // console.log("Client <-- Server")
            handleMessage(e.data);
        }
    }


    WS_COMM.Send = function(msg) {
        // console.log("Client --> Server")
        _socket.send(msg); // Send a nice message to the server
    }

    // private functions

    function handleMessage(msg) {
        var len = msg.length;
        msg = msg.trim();

        var si = msg.indexOf(" ")
        var tag = msg.substr(0, si)

        if (tag in _handlers) {
            handle = _handlers[tag]
            var body = msg.substring(tag.length).trim();
            handle(body)
        } else {
            console.log('Error: message type \"' + tag + '\" unknown');
            console.log("msg: ", msg)
        }
    }

    return WS_COMM;
}(jQuery));