
// Set variable
let displaytext = document.querySelector(".displaytext");
let displaycopy = document.querySelector(".displaycopy");
let select = document.querySelector(".select");
let poleequal = document.getElementById("equalentryes");
let polegraph = document.getElementById("graphentryes");
let buttons = Array.from(document.querySelectorAll(".button"));

// Set entries to hidden
  poleequal.style.display = "none";		
  polegraph.style.display = "none";

// Handle mode select with change entries
select.addEventListener("change", function() {
  displaytext.innerText = `change select:` + select.value;

  if (select.value == "calculate") {
    poleequal.style.display = "none";		
    polegraph.style.display = "none";
  } else if (select.value == "equal") {
    poleequal.style.display = "block";		
    polegraph.style.display = "none";
  } else if (select.value == "graph") {
    poleequal.style.display = "block";		
    polegraph.style.display = "block";
  } 

});

// Create WebSocket and assign "onmessage"
let socket = new WebSocket("ws://localhost:8080/calculate/start");
socket.onmessage = function(event) {
  displaytext.innerText = `SERV: ${event.data}`;
}

// Handle copying equation to clipboard by click on the pole
displaytext.addEventListener("click", (e) => {
      navigator.clipboard.writeText(document.getElementById('dtext').innerHTML);
      displaycopy.innerText = "copyed"
});

// Handle 
buttons.map((button) => {
  button.addEventListener("click", (e) => {
    switch (e.target.innerText) {
      case "<=":
        if (displaytext.innerText.length === 1){
          displaytext.innerText = "0";
        }else {
          displaytext.innerText = displaytext.innerText.slice(0, displaytext.innerText.length - 1)
        }
        break;
      case "C":
        displaytext.innerText = "0";
        displaycopy.innerText = ""
        break;
      case "=":
        try {
            socket.send(select.value + " " + displaytext.innerText)
        } catch (e) {
          displaytext.innerText = "Error!";
        }
        break;
      case "+/-":
        displaytext.innerText = "-";
        break;
      case "%":
        let passedText = displaytext.innerText + "/100";
        displaytext.innerText = eval(passedText);
        displaycopy.innerText = ""
        break;
      default:
        if (displaytext.innerText === "0" && e.target.innerText !== ".") {
          displaytext.innerText = e.target.innerText;
          displaycopy.innerText = ""
        } else {
          displaytext.innerText += e.target.innerText;
          displaycopy.innerText = ""
        }
    }
  });
});