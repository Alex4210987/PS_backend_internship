window.addEventListener('DOMContentLoaded', (event) => {
  const routePlanningForm = document.getElementById('routePlanningForm');
  const historyContainer = document.getElementById('historyDataContainer');

  routePlanningForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const originInput = document.getElementById('originInput');
    const destinationInput = document.getElementById('destinationInput');
    const modeSelect = document.getElementById('modeSelect');
    const outputModeSelect = document.getElementById('outputModeSelect');
    const tacticsSelect = document.getElementById('tacticsSelect');

    const origin = originInput.value;
    const destination = destinationInput.value;
    const mode = modeSelect.value;
    const outputMode = outputModeSelect.value;
    const tactics = tacticsSelect.value;

    const formData = new FormData();
    formData.append('origin', origin);
    formData.append('destination', destination);
    formData.append('mode', mode);
    formData.append('outputmode', outputMode);
    formData.append('tactics', tactics);

    fetch('/routeplanning', {
      method: 'POST',
      body: formData,
    })
    .then((response) => response.json())
    .then((data) => {
      const status = data.status;
      const msg = data.message;
      const output = data.result;

      displayResult(output, origin,destination,mode,outputMode,tactics);
    })
    .catch((error) => {
      console.error('路径规划失败:', error);
    });
  });

  function displayResult(output, origin,destination,mode,outputMode,tactics) {
    var historyContainer = document.getElementById("historyDataContainer");
    var inputHTML = `<p>起点: ${origin}</p>`+'<p>终点: '+destination+'</p>'+'<p>出行方式: '+mode+'</p>'+'<p>输出方式: '+outputMode+'</p>'+'<p>策略: '+tactics+'</p>';
    var outputHTML = '';
    output.msg.forEach((message) => {
      outputHTML += `<p>${message}</p>`;
    });

    historyHTML = inputHTML + outputHTML;
    historyContainer.innerHTML += historyHTML;
  }
});
