document.addEventListener("DOMContentLoaded", () => {
  const btn = document.getElementById("eventBtn");
  if (btn) {
      btn.addEventListener("click", () => {
          fetch("/trigger-event", {
              method: "POST", // Change from GET to POST
              headers: {
                  "Content-Type": "application/json"
              }
          })
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

      // Toggle the display of the extra details container when the button is clicked
document.querySelectorAll('.toggle-section').forEach(function(button) {
  button.addEventListener('click', function() {
    var targetId = this.getAttribute('data-target');
    var section = document.getElementById(targetId);
    if (section.style.display === 'none') {
      section.style.display = 'block';
      this.innerText = 'Hide';
    } else {
      section.style.display = 'none';
      this.innerText = 'Show';
    }
  });
});
  