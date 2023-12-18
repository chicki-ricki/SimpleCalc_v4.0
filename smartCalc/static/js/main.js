
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

  x = document.getElementById("x"),
  xFrom = document.getElementById("xfrom"),
  xTo = document.getElementById("xto"),
  yFrom = document.getElementById("yfrom"),
  yTo = document.getElementById("yto");

x.value = 0;
xFrom.value = -300;
xTo.value = 300;
yFrom.value = -300;
yTo.value = 300;

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

function getEntries (select, x, xFrom, xTo, yFrom,yTo) {
  switch (select.value) {
    case 'calculate':
      return ("");
    case "equal":
      return (x.value.replace(/\s/g,'') + " ");
    case "graph":
      return (xFrom.value.replace(/\s/g,'') + " " + 
          xTo.value.replace(/\s/g,'') + " " + 
          yFrom.value.replace(/\s/g,'') + " " + 
          yTo.value.replace(/\s/g,'') + " ");
  } 
  return "";
}

// Set entries to hidden
showEntries(select, poleequal,polegraph, displayEquation);

// Handle mode select with change entries
select.addEventListener("change", function() {
  // displaytext.innerText = `change select:` + select.value;
  showEntries(select, poleequal, polegraph, displayEquation);
});

// Create WebSocket and assign "onmessage"
let socket = new WebSocket("ws://localhost:8080/calculate/start");
socket.onmessage = function(event) {
  displaytext.innerText = `SERV: ${event.data}`;
}
socket.onclose = function() {
  //reconnect
 socket = new WebSocket("ws://localhost:8080/calculate/start");
}

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

// Handle buttons reaction
buttons.map((button) => {
  button.addEventListener("click", (e) => {
    if (select.value == "calculate") {
      clickHandle(e.target.innerText, displaytext, "");
    } else {
      clickHandle(e.target.innerText, displayEquation, getEntries(select, x, xFrom, xTo, yFrom, yTo))
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