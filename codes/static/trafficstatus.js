window.addEventListener('DOMContentLoaded', (event) => {
  const trafficStatusForm = document.getElementById('trafficStatusForm');
  const historyContainer = document.getElementById('historyDataContainer');

  trafficStatusForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const locationInput = document.getElementById('locationInput');
    const location = locationInput.value;

    const formData = new FormData();
    formData.append('location', location);

    fetch('/trafficstatus', {
      method: 'POST',
      body: formData,
    })
    .then((response) => response.json())
    .then((data) => {
      const status = data.status;
      const msg = data.msg;
      const trafficData = data.data;
      const trafficDesc = data.desc;

      displayResult(status, msg, trafficData, trafficDesc,location);
    })
    .catch((error) => {
      console.error('路况查询失败:', error);
    });
  });

  function displayResult(status, msg, trafficData, trafficDesc,location) {
    var inputHTML = `
      <p>路况数据：${trafficData}</p>
      <p>描述信息：${trafficDesc}</p>
    `;
    var locHTML="<p>"+location+"的路况是："+"</p>";
    historyContainer.innerHTML += locHTML;
    historyContainer.innerHTML += inputHTML;
  }
});
