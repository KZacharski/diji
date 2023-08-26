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

  const copyButtonLabel = "Copy Code";

// use a class selector if available
let blocks = document.querySelectorAll("pre");

blocks.forEach((block) => {
  // only add button if browser supports Clipboard API
  if (navigator.clipboard) {
    let button = document.createElement("button");

    button.innerText = copyButtonLabel;
    block.appendChild(button);

    button.addEventListener("click", async () => {
      await copyCode(block, button);
    });
  }
});

async function copyCode(block, button) {
  let code = block.querySelector("code");
  let text = code.innerText;

  await navigator.clipboard.writeText(text);

  // visual feedback that task is completed
  button.innerText = "Code Copied";

  setTimeout(() => {
    button.innerText = copyButtonLabel;
  }, 700);
}

  