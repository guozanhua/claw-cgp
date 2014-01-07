/*
 * @author sheppard(ysf1026@gmail.com) 2014-01-02
 */

var socket;
var msgs = [];

$(document).ready(function() {
  connect();

  var btn_test = $('#btn-test');
  btn_test.click(function() {
    var content = btn_test.text(); 
    socket.send(content);
    return false;
  });
});

function connect() {
  var userName = $('#userName').text();
  socket = new WebSocket('ws://'+window.location.host+'/hall/hall/socket?user=' + userName);

  socket.onmessage = function(event) {
    var msg = JSON.parse(event.data);
    addToPanel(msg.Name, msg.Content);
    refreshPanel();
  };

  socket.onclose = function() {
    console.log('disconnected');
  };
};

function addToPanel(prefix, msg) {
  msgs.push({prefix: prefix, msg: msg});
};

function refreshPanel() {
  var panel = $('#panel-chat');
  panel.empty();
  _.each(msgs, function(item) {
    panel.append('<strong>' + item.prefix + ':</strong> ');
    panel.append(item.msg).append('<br>');
  });
};

function refreshRoomList() {
  var rooms = $('#list-room');
}
