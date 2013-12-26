(function() {
  window.onload = connect;
  function connect() {
    var ProtoBuf = dcodeIO.ProtoBuf;
    var CLPing = ProtoBuf.loadProtoFile('/public/proto/client_logic.proto').build('CLPing');

    var h1 = document.getElementsByTagName("h1")[0];
    var userName = (h1.getAttribute('user'));
    // Create a socket
    var socket = new WebSocket('ws://'+window.location.host+'/hall/hall/socket?user=' + userName)

    // Display a message
    var display = function(data) {
      alert(data)
      $('#thread').append(tmpl('message_tmpl', data));
      $('#thread').scrollTo('max')
    }

    // Message received on the socket
    socket.onmessage = function(event) {
      alert(event)
      display(event.data)
    }

    socket.onclose = function() {
      alert('disconnected')
    }

    $('#send').click(function(e) {
      var message = $('#message').val()
      $('#message').val('')

      socket.send(message)
    });

    $('#message').keypress(function(e) {
      if(e.charCode == 13 || e.keyCode == 13) {
        $('#send').click()
        e.preventDefault()
      }
    })
  }
})();
