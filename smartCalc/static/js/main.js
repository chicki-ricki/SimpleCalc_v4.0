let displaytext = document.querySelector(".displaytext");
let displaycopy = document.querySelector(".displaycopy");
let select = document.querySelector(".select");

let buttons = Array.from(document.querySelectorAll(".button"));

displaytext.addEventListener("click", (e) => {
      navigator.clipboard.writeText(document.getElementById('dtext').innerHTML);
      displaycopy.innerText = "copyed"
});

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
          // displaytext.innerText = eval(displaytext.innerText);
           let socket = new WebSocket("ws://localhost:8080/calculate/start");
           socket.onopen = function() {
            socket.send(displaytext.innerText)
           }
           socket.onmessage = function(event) {
           displaytext.innerText = `[message] Данные получены с сервера: ${event.data}`;
          }
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