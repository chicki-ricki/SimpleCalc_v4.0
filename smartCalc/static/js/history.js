function showHistory(historyJson) {
    historyJson.forEach(item => {
      let elem = document.createElement('div');
      if (item.equation.length > 20) {
        equationText = item.equation.slice(0,20)+"...";
      } else {
        equationText = item.equation;
      }
      switch (item.mode) {
        case 'calc':
          elem.innerHTML = `
            <button class="historyButtonsClass"> 
                ${item.mode}:   ${equationText} = ${item.result}
            </button>
          `;
          break;
        case 'equal':
          elem.innerHTML = `
            <button class="historyButtonsClass"> 
                ${item.mode}: {x=${item.xEqual}} ${equationText} = ${item.result}
            </button>
          `;
          break;
        case 'graph':
            elem.innerHTML = `
            <button class="historyButtonsClass"> 
                ${item.mode}: X{${item.xFrom}...${item.xTo}} Y{${item.yFrom}...${item.yTo}}; y = ${equationText}
            </button>
          `;
          break;            
      }
      historyWindow.append(elem);
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
}