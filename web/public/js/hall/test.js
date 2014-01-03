/*
 * @author sheppard(ysf1026@gmail.com) 2014-01-02
 */

var socket;

$(document).ready(function() {
  connect();

  var btn_test = $('#btn-test');
  btn_test.text('haha');
  btn_test.click(function() {
    var content = btn_test.text(); 
    updatePanel(content);
    socket.send(content);
    return false;
  });
});

function updatePanel(msg) {
  $('#panel-body').text(msg);
}

function connect() {
  var userName = $('#userName').text();
  socket = new WebSocket('ws://'+window.location.host+'/hall/hall/socket?user=' + userName);

  socket.onmessage = function(event) {
    console.log(event);
  };

  socket.onclose = function() {
    console.log('disconnected');
  };
};
