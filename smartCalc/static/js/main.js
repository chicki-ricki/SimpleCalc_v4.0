
// Set variable
let 
  displaytext = document.querySelector(".displaytext"),
  displayEqual = document.getElementById("dEqual"),
  displayEquation = document.getElementById("dEText"),
  displaycopy = document.querySelector(".displaycopy"),
  select = document.querySelector(".select"),
  poleequal = document.getElementById("equalentryes"),
  polegraph = document.getElementById("graphentryes"),
  buttons = Array.from(document.querySelectorAll(".button")),
  buttonX = document.querySelector(".bigbutton")
  graphWindow = document.getElementById("graphWindow"),
  historyWindow = document.getElementById("historyWindow"),
  // graphButton = document.getElementById("graphButton"),


  x = document.getElementById("x"),
  xFrom = document.getElementById("xfrom"),
  xTo = document.getElementById("xto"),
  yFrom = document.getElementById("yfrom"),
  yTo = document.getElementById("yto");

x.value = 1;
xFrom.value = -300;
xTo.value = 300;
yFrom.value = -300;
yTo.value = 300;

graphWindow.style.display = "none";
historyWindow.style.display = "none";

function showEntries (select, equal, graph, equation) {
  switch (select.value) {
    case 'calculate':
      equal.style.display = "none";		
      graph.style.display = "none";
      equation.style.display = "none";
      break;
  case "equal":
      equal.style.display = "block";		
      graph.style.display = "none";
      equation.style.display = "block";
      break;
  case "graph":
      equal.style.display = "none";		
      graph.style.display = "block";
      equation.style.display = "block";
  } 
}

function getEntries (select, x, xFrom, xTo, yFrom,yTo, space) {
  switch (select.value) {
    case 'calculate':
      return ("");
    case "equal":
      if (x.value == "") {x.value = "1"}
      return (x.value.replace(/\s/g,'') + space);
    case "graph":
      if (xFrom.value == "") {xFrom.value = "-300"}
      if (xTo.value == "") {xTo.value = "300"}
      if (yFrom.value == "") {yFrom.value = "-300"}
      if (yTo.value == "") {yTo.value = "300"}
      return (xFrom.value.replace(/\s/g,'') + space + 
          xTo.value.replace(/\s/g,'') + space + 
          yFrom.value.replace(/\s/g,'') + space + 
          yTo.value.replace(/\s/g,'') + space);
  } 
  return "";
}

function getRandomInRange(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

// Set entries to hidden
showEntries(select, poleequal,polegraph, displayEquation);

// Handle mode select with change entries
select.addEventListener("change", function() {
  // displaytext.innerText = `change select:` + select.value;
  showEntries(select, poleequal, polegraph, displayEquation);
});

uname = ""
fileDownload = ""
// Create WebSocket
let socket = new WebSocket("ws://localhost:8080/calculate/start");

// Assign "onmessage" for socket
socket.onmessage = function(event) {
  switch (event.data[0]){
    case "5":
      uname = event.data.slice(2);
      urlgraph = "/static/tmp/tempGraph" + uname + ".png"
      // displaytext.innerText = uname;
      break;
    case "0":
    case "1":
      displaytext.innerText = event.data.slice(2);
      break;
    case "2":
    displaytext.innerText = event.data.slice(2);
    if (displaytext.innerText != "Error") {
    graphWindow.style.display = "block";
    graphSpan.innerHTML = "<img src=\"" + urlgraph + "?dummy="+getRandomInRange(2, 500000)+"\" class=\"graphImage\" id=\"graphImage\">"
    fileDownload = displayEquation.innerText + "_" + getEntries(select, x, xFrom, xTo, yFrom, yTo, "_")
    }
  }
}

socket.onclose = function() {
  //reconnect
 socket = new WebSocket("ws://localhost:8080/calculate/start");
}

function download(url, fileDownload) {
  const a = document.createElement('a')
  a.href = url
  a.download = "CleverCalc_" + fileDownload + ".png"
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
}

downGraph.addEventListener("click", (e) => {
  download(urlgraph, fileDownload)
})

// Handle copying equation to clipboard by click on the pole
displaytext.addEventListener("click", (e) => {
      navigator.clipboard.writeText(document.getElementById('dtext').innerHTML);
      displaycopy.innerText = "copyed";
});

function clickHandle(val, pole, entries) {
switch (val) {
  case "<=":
    if (pole.innerText.length === 1){
      pole.innerText = "0";
    }else {
      pole.innerText = pole.innerText.slice(0, pole.innerText.length - 1)
    }
    break;
  case "C":
    pole.innerText = "0";
    displaycopy.innerText = "";
    break;
  case "=":
    try {
        socket.send(select.value + " " + entries + pole.innerText.replace(/\s/g,''));
      } catch (e) {
      pole.innerText = "Socket error!";
    }
    break;
  case "+/-":
    pole.innerText = "-";
    break;
  case "cos":
  case "acos":
  case "sin":
  case "asin":
  case "tan":
  case "atan":
  case "ln":
  case "log":
  case "sqrt":
  case "mod":
  case "^":
    if (pole.innerText != "0") {
    pole.innerText += val + "(";
  } else {
    pole.innerText = val + "(";
  }
    break;
  default:
    if (pole.innerText === "0" && val !== ".") {
      pole.innerText = val;
      displaycopy.innerText = "";;
    } else {
      pole.innerText += val;
      displaycopy.innerText = "";
    }
  }
}

historyButton.addEventListener("click", (e) => {
  if (historyWindow.style.display == "none") {
    historyWindow.style.display = "block"
  } else {
    historyWindow.style.display = "none"
  }
})

graphButton.addEventListener("click", (e) => {
  if (graphWindow.style.display == "none") {
    graphWindow.style.display = "block"
  } else {
    graphWindow.style.display = "none"
  }
})

// Handle buttons reaction
buttons.map((button) => {
  button.addEventListener("click", (e) => {
    if (select.value == "calculate") {
      clickHandle(e.target.innerText, displaytext, "");
    } else {
      clickHandle(e.target.innerText, displayEquation, getEntries(select, x, xFrom, xTo, yFrom, yTo, " "))
    }
  });
});

buttonX.addEventListener("click", (e) => {
  if (select.value != "calculate") {
    if (displayEquation.innerText != "0") {
      displayEquation.innerText += e.target.innerText;
    } else {
      displayEquation.innerText = e.target.innerText;
    }
  }
});