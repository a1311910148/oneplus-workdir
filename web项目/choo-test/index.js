// import choo
var choo = require("choo");

// import choo's template helper
var html = require("choo/html");

// initialize choo
var app = choo();

// create a template
var main = function () {
  return html`<div>sognzhxinxin<a href="/side">side</a></div>`;
};
// create a template
var side = function () {
  return html`<div>side${[1,2,3].map(main)}</div>`;
};



app.route('/', main)
app.route('/side', side)


// start app
app.mount("div");
