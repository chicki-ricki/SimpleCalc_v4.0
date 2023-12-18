<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Calculator</title>
    <link rel="stylesheet" href="/static/css/startCalculate_style.css" />
  </head>
  <body>
    <div class="calculator">
      <div class="display" >
        <span class="displaytext" id="dtext">0</span>
        <div class="displaycopy" id="dcopy"></div>
      </div>
      <div class="displayEquation" id="dEqual">
        <span class="displayEqualText" id="dEText">0</span>
        <div class="displaycopy" id="dcopy"></div>
      </div>
    <div id="graphentryes">
      <button class="notapbutton">X =</button>
      <input type="text" class="notapbutton btn-lightgrey inputborder" placeholder="xfrom" inputmode="numeric" id="xfrom"></input>
      <input type="text" class="notapbutton btn-lightgrey inputborder" placeholder="xto" inputmode="numeric" id="xto"></input>
      <button class="notapbutton">Y =</button>
      <input type="text" class="notapbutton btn-lightgrey inputborder" placeholder="yfrom" inputmode="numeric" id="yfrom"></input>
      <input type="text" class="notapbutton btn-lightgrey inputborder" placeholder="yto" inputmode="numeric" id="yto"></input>
    </div>
    <div id="equalentryes">
      <button class="notapbutton">X =</button>
      <input type="text" class="notapbutton btn-lightgrey inputx" placeholder="enter X here" inputmode="numeric" id="x"></input>
    </div>
      <div class="buttons">
        <select name="mode" class="notapbutton up select">
          <option value="calculate" selected="selected">calculate</option>
          <option value="equal">equal</option>
          <option value="graph">graph</option>
        </select>
        <button class="button btn-orange">C</button>
        <button class="button btn-orange">&lt;=</button>
        <button class="button">x</button>
        <button class="button">e</button>
        <button class="button">cos</button>
        <button class="button">acos</button>
        <button class="button">(</button>
        <button class="button">)</button>
        <button class="button">&#37;</button>
        <button class="button">/</button>
        <button class="button">sin</button>
        <button class="button ">asin</button>
        <button class="button">7</button>
        <button class="button">8</button>
        <button class="button">9</button>
        <button class="button">*</button>
        <button class="button ">tan</button>
        <button class="button ">atan</button>
        <button class="button">4</button>
        <button class="button">5</button>
        <button class="button">6</button>
        <button class="button">-</button>
        <button class="button ">ln</button>
        <button class="button ">log</button>
        <button class="button">1</button>
        <button class="button">2</button>
        <button class="button">3</button>
        <button class="button">+</button>
        <button class="button">sqrt</button>
        <button class="button">mod</button>
        <button class="button">0</button>
        <button class="button">.</button>
        <button class="button">^</button>
        <button class="button btn-orange">=</button>
        <button class="notapbutton down">help</button>
        <button class="notapbutton down history">history</button>
        <button class="notapbutton down graph">graph</button>
      </div>
    </div>
    <script src="/static/js/main.js"></script>
  </body>
</html>
