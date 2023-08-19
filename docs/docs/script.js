var isOpen = false

function openNav() {
    document.getElementById("drawer").style.width = "250px";
    document.getElementById("drawerbg").style.width = "100%";

    isOpen = true
  }
  
  /* Set the width of the side navigation to 0 */
  function closeNav() {
    document.getElementById("drawer").style.width = "0";
    document.getElementById("drawerbg").style.width = "0";
    isOpen = false
  } 

  