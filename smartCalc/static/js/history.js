function showHistory(historyFromModel) {
    historyJson = JSON.parse(historyFromModel);
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

function showAlert(historyButtons) {
    historyButtons.map((hButton) => {
        hButton.addEventListener("click", (e) => {
        //   switch (select.value) {
            // case "calculate":
            //   console.log("e.target.data:", e.target.data);  
            alert(e.target.innerText);
            // displayEquation.innerText = "e.target.data.value";
            //   break;
        
            // default:
            //   break;
        //   }
        });
    });
}