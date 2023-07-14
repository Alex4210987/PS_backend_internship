document.getElementById("aliasForm").addEventListener("submit", function(event) {
  event.preventDefault();

  var locationValue = document.getElementById("originalNameInput").value;
  var aliasValue = document.getElementById("aliasInput").value;

  var aliasObject = {
    location: locationValue,
    alias: aliasValue
  };
  console.log(aliasObject);

  fetch("/alias", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(aliasObject)
  })
    .then(response => response.json())
    .then(data => {
      console.log(data);
    })
    .catch(error => {
      console.error("Error:", error);
    });

  document.getElementById("originalNameInput").value = "";
  document.getElementById("aliasInput").value = "";
});
