function showHistoryElem(jsonElem) {
let elem = document.createElement('div');
      if (jsonElem.equation.length > 20) {
        equationText = jsonElem.equation.slice(0,20)+"...";
      } else {
        equationText = jsonElem.equation;
      }
      switch (jsonElem.mode) {
        case 'calc':
          elem.innerHTML = `
            <button class="historyButtonsClass"> 
                ${jsonElem.mode}:   ${equationText} = ${jsonElem.result}
            </button>
          `;
          break;
        case 'equal':
          elem.innerHTML = `
            <button class="historyButtonsClass"> 
                ${jsonElem.mode}: {x=${jsonElem.xEqual}} ${equationText} = ${jsonElem.result}
            </button>
          `;
          break;
        case 'graph':
            elem.innerHTML = `
            <button class="historyButtonsClass"> 
                ${jsonElem.mode}: X{${jsonElem.xFrom}...${jsonElem.xTo}} Y{${jsonElem.yFrom}...${jsonElem.yTo}}; y = ${equationText}
            </button>
          `;
          break;            
      }
      historyWindowData.append(elem);  
}

function showHistory(historyJson) {
    historyJson.forEach(item => {
      showHistoryElem(item);
    });
};


function showClickHistoryButton(historyButtons, historyJson) {
    historyButtons.map((hButton) => {
        hButton.addEventListener("click", (e) => {
          let index = historyButtons.indexOf(hButton);
          switch (historyJson[index].mode) {
            case "calc":
              select.selectedIndex = 0;
              showEntries(select, poleequal, polegraph, displayEquation);
              displaytext.innerText = historyJson[index].equation;
              break;
            case "equal":
              select.selectedIndex = 1;
              showEntries(select, poleequal, polegraph, displayEquation);
              displayEquation.innerText = historyJson[index].equation;
              x.value = historyJson[index].xEqual;
              displaytext.innerText = 0;
              break;
            case "graph":
              select.selectedIndex = 2;
              showEntries(select, poleequal, polegraph, displayEquation);
              displayEquation.innerText = historyJson[index].equation;
              xFrom.value = historyJson[index].xFrom;
              xTo.value = historyJson[index].xTo;
              yFrom.value = historyJson[index].yFrom;
              yTo.value = historyJson[index].yTo;
              displaytext.innerText = 0;
              break;
          }
        });
    });
};

clearHistoryButton.addEventListener("click", (e) => {
  try {
    socket.send('clearHistory');
  } catch (e) {
  pole.innerText = "Socket error!";
}
});
