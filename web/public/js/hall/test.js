/*
 * @author sheppard(ysf1026@gmail.com) 2014-01-02
 */

var socket;
var msgs = [];

$(document).ready(function() {
  connect();

  var btn_test = $('#btn-test');
  btn_test.text('haha');
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
    addToPanel(event.data);
    refreshPanel();
  };

  socket.onclose = function() {
    console.log('disconnected');
  };
};

function addToPanel(msg) {
  msgs.push(msg);
};

function refreshPanel() {
  var panel = $('#panel-body');
  panel.empty();
  _.each(msgs, function(item) {
    panel.append(item).append('<br>');
  });
}

