/*
 * @author sheppard(ysf1026@gmail.com) 2014-01-02
 */

window.onload = function() {
  var btn_test = document.getElementById('btn-test');
  btn_test.innerText = 'haha';
  btn_test.onclick = function() {
    updatePanel(btn_test.innerText);
    console.dir(btn_test);
    return false;
  };
};

function updatePanel(msg) {
  var panel = document.getElementById('panel-body');
  panel.innerText = msg;
}
