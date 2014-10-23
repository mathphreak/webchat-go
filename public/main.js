$(function(){

    var names = ["aisle", "bdellium", "czar", "djinn", "Euphrates", "fohn", "gnarly", "irk", "hour", "jalapenos", "knick-knack", "llama", "mnemonic", "ndomo", "ouija board", "pneumonia", "qat",
    "argyle", "Szr", "tsunami", "urn", "vraisamblance", "wren", "Xian", "yiperite", "Zed Zed Top"];

    window.conn = {};
    var msg = $("#msg");
    var log = $("#log");

    $("#name").val(names[Math.floor(Math.random() * (names.length))]);

    var knownUsers = [];

    function appendLog(safeName, safeMsg) {
        var d = log[0];
        var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
        var msg = $("<tr><td>" + safeName + "</td><td>" + safeMsg + "</td></tr>");
        msg.appendTo(log);
        if (doScroll) {
            d.scrollTop = d.scrollHeight - d.clientHeight;
        }
    }

    $("#form").on("submit", function(evt) {
        if (!conn) {
            return false;
        }
        if (!msg.val()) {
            return false;
        }
        var message = msg.val();
        var username = $("#name").val();
        conn.send(JSON.stringify({type: "message", msg: message, name: username}));
        msg.val("");
        evt.preventDefault();
    });

    if (window.WebSocket && window.JSON) {
        var connectTimeout;
        var connect = function() {
            conn = new WebSocket("ws://" + location.host + "/ws");
            function onMessage(evt) {
                var data = JSON.parse(evt.data);
                var safeName = $("<span/>").text(data.name).html();
                if (data.type === "message") {
                    var safeMsg = $("<span/>").text(data.msg).html();
                    appendLog(safeName, safeMsg);
                } else if (data.type === "join") {
                    if (knownUsers.indexOf(safeName) == -1) {
                        $("#users").append("<li>" + safeName + "</li>");
                        knownUsers.push(safeName);
                        conn.send(JSON.stringify({type: "join", name: $("#name").val()}));
                    }
                } else if (data.type === "leave") {
                    var oldHTML = $("#users").html();
                    var newHTML = oldHTML.replace("<li>" + safeName + "</li>", "");
                    $("#users").html(newHTML);
                    var index = knownUsers.indexOf(safeName);
                    if (index != -1) {
                        knownUsers = knownUsers.filter(function(x, i){return i != index;});
                    }
                }
            }
            function onOpen(evt) {
                appendLog("<b>Status</b>", "<b>Connection established!</b>");
                setTimeout(function(){conn.send(JSON.stringify({type: "join", name: $("#name").val()}));},10);
                clearTimeout(connectTimeout);
            }
            function onClose(evt) {
                appendLog("<b>Status</b>", "<b>Connection closed.</b>");
                connectTimeout = setTimeout(connect, 1000);
            }
            conn.onopen = onOpen;
            conn.onclose = onClose;
            conn.onmessage = onMessage;
            conn.onerror = function(){console.err(arguments);};
        };
        connect();

        window.onbeforeunload = window.onunload = function() {
            conn.send(JSON.stringify({type:"leave",name:$("#name").val()}));
        };
    } else {
        appendLog("<b>Status</b>", "<b>Your browser does not support WebSockets and/or JSON.</b></td></tr>");
    }
});
