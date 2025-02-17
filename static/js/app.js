document.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById("eventBtn");
    if (btn) {
      btn.addEventListener("click", () => {
        fetch("/trigger-event")
          .then(response => response.json())
          .then(data => {
            alert(data.status);
          })
          .catch(err => {
            console.error(err);
            alert("Error triggering event");
          });
      });
    }
  });
  