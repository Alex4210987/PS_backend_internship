document.getElementById("trackMatchForm").addEventListener("submit", function(event) {
  event.preventDefault(); // 阻止表单的默认提交行为

  var trackInput = document.getElementById("trackInput").value;
  var standardTrackInput = document.getElementById("standardTrackInput").value;

  var requestData = {
    track: trackInput.split(";"), // 使用分号作为字符串分隔符
    standard_track: standardTrackInput.split(";")
  };

  // 发送POST请求
  fetch("/trackmatch", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(requestData)
  })
  .then(response => response.json())
  .then(data => {
    if (data.status === 200) {
      displayResult(data.data.data.similarity, requestData); // 显示相似度结果，并传递requestData
      console.log("轨迹匹配成功:", data.data.data.similarity);
      console.log(data.data.data.similarity);
    } else {
      console.error("轨迹匹配失败:", data.msg);
    }
  })
  .catch(error => {
    console.error("请求错误:", error);
  });
});

function displayResult(similarity, requestData) {
  var historyContainer = document.getElementById("historyDataContainer");

  // 显示输入数据
  var inputHTML ="<p>输入轨迹: " + requestData.track.join(";") + "</p>";
  inputHTML += "<p>标准轨迹: " + requestData.standard_track.join(";") + "</p>";
  historyContainer.innerHTML += inputHTML;

  // 显示输出数据
  var outputHTML = "<p>轨迹匹配成功，相似度: " + similarity + "</p>";
  historyContainer.innerHTML += outputHTML;
}

  